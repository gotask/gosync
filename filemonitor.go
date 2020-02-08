// filemonitor.go
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gotask/gost/stconfig"
	"github.com/gotask/gost/stutil"
)

type Op uint32

// Ops
const (
	Create Op = iota
	Modify
	Delete
	Init
)

var ops = map[Op]string{
	Create: "CREATE",
	Modify: "MODIFY",
	Delete: "DELETE",
	Init:   "INIT",
}

// String prints the string version of the Op consts
func (e Op) String() string {
	if op, found := ops[e]; found {
		return op
	}
	return "???"
}

type Event struct {
	Op
	Path string
	Info os.FileInfo
	Conf DirConfig
}

func (e Event) String() string {
	return fmt.Sprintf("%s [%s]", e.Op, e.Path)
}

func allFiles(path string, include, exclude *regexp.Regexp, completex []string) map[string]os.FileInfo {
	fileList := make(map[string]os.FileInfo)
	stutil.FileReadDir(path, true, fileList)

	if include != nil || exclude != nil || len(completex) > 0 {
		for k, v := range fileList {
			if v.IsDir() {
				delete(fileList, k)
				continue
			}
			if len(completex) > 0 {
				con := false
				for _, c := range completex {
					if k == c {
						con = true
						break
					}
				}
				if !con {
					ds := strings.Split(k, string(os.PathSeparator))
					for _, d := range ds {
						for _, c := range completex {
							if d == c {
								con = true
								break
							}
						}
						if con {
							break
						}
					}
				}
				if con {
					delete(fileList, k)
					continue
				}
			}
			if include != nil {
				if include.MatchString(k) {
					continue
				}
			}
			if exclude != nil {
				if exclude.MatchString(k) {
					delete(fileList, k)
				}
			}
		}
	}
	return fileList
}

func compareFiles(new, old map[string]os.FileInfo, c DirConfig) []Event {
	var result []Event
	if old == nil { //first call
		for k, v := range new {
			result = append(result, Event{Init, k, v, c})
		}
		return result
	}
	for k, v := range new {
		if v2, ok := old[k]; ok {
			if v.ModTime() != v2.ModTime() {
				result = append(result, Event{Modify, k, v, c})
			}
			delete(old, k)
		} else {
			result = append(result, Event{Create, k, v, c})
		}
	}
	for k, v := range old {
		result = append(result, Event{Delete, k, v, c})
	}
	return result
}

type WatchDir struct {
	Path      string
	Include   *regexp.Regexp
	Exclude   *regexp.Regexp
	Completex []string
	Current   map[string]os.FileInfo
	Conf      DirConfig
}

type Watch struct {
	Ev    chan []Event
	Conf  WatchConfig
	CfMd5 string
	elems map[string]*WatchDir
	elock sync.Mutex
	close bool
	cchan chan int
}

type DirConfig struct {
	LocalDir   string
	RemoteDir  string
	RemoteAddr string
	Include    string
	Exclude    string
	Completex  string
}
type WatchConfig struct {
	LoopSec uint32
	Address string
	Dirs    map[string]DirConfig
}

func (w *Watch) Start() {
	w.Ev = make(chan []Event, 1024)
	//w.elems = make(map[string]*WatchDir)
	w.close = false
	w.cchan = make(chan int, 1)

	go func(w *Watch) {
		for !w.close {
			w.Reload(false)
			need := time.Duration(w.Conf.LoopSec) * time.Second
			n1 := time.Now()
			w.elock.Lock()
			for _, e := range w.elems {
				newf := allFiles(e.Path, e.Include, e.Exclude, e.Completex)
				ev := compareFiles(newf, e.Current, e.Conf)
				if len(ev) > 0 {
					w.Ev <- ev
				}
				e.Current = newf
			}
			w.elock.Unlock()
			n2 := time.Now()
			d := n2.Sub(n1)
			if d < need {
				time.Sleep(need - d)
			}
		}
		w.cchan <- 1
	}(w)
}

func (w *Watch) Stop() {
	w.close = true
	<-w.cchan
	close(w.cchan)
}

func (w *Watch) Reload(force bool) {
	if !force {
		md5, _ := stutil.FileMD5("sync.ini")
		if md5 == w.CfMd5 {
			return
		}
	}

	w.elock.Lock()
	w.loadConfig()
	w.elems = make(map[string]*WatchDir)
	w.elock.Unlock()
	for _, v := range w.Conf.Dirs {
		w.AddWatch(v)
	}
	LOG.Info("reload config~~~~~~")
}

func (w *Watch) AddWatch(conf DirConfig) error {
	var include, exclude *regexp.Regexp
	var err error
	if conf.Include != "" {
		include, err = regexp.Compile(conf.Include)
	}
	if conf.Exclude != "" {
		exclude, err = regexp.Compile(conf.Exclude)
	}
	if err != nil {
		LOG.Error("addwatch error %v %s", conf, err.Error())
		return err
	}

	LOG.Info("addwatch %v", conf)

	w.elock.Lock()
	defer w.elock.Unlock()
	if e, ok := w.elems[conf.LocalDir]; ok {
		e.Include = include
		e.Exclude = exclude
		e.Completex = strings.Split(conf.Completex, "|")
		e.Conf = conf
	} else {
		w.elems[conf.LocalDir] = &WatchDir{conf.LocalDir, include, exclude, strings.Split(conf.Completex, "|"), nil, conf}
	}
	return nil
}

func (w *Watch) loadConfig() {
	w.CfMd5, _ = stutil.FileMD5("sync.ini")
	c, e := stconfig.LoadINI("sync.ini")
	var conf WatchConfig
	conf.LoopSec = 1
	conf.Address = "0.0.0.0:3030"
	conf.Dirs = make(map[string]DirConfig)
	if e != nil {
		LOG.Info("no config ini,loopsec=1 listen@%s", conf.Address)
		w.Conf = conf
		return
	}
	conf.LoopSec = uint32(c.IntegerSection("system", "loopsec", 1))
	conf.Address = c.StringSection("system", "address", "0.0.0.0:3030")

	for i := 1; i <= 100; i++ {
		sec := "sync" + stutil.IntToString(int64(i))
		tempD := c.StringSection(sec, "dir", "")
		if tempD == "" {
			continue
		}
		var d DirConfig
		ds := strings.Split(tempD, "=>")
		if len(ds) == 2 {
			rs := strings.Split(ds[1], "@")
			if len(rs) == 2 {
				d.LocalDir = ds[0]
				d.RemoteDir = rs[0]
				d.RemoteAddr = rs[1]

				d.Include = c.StringSection(sec, "include", "")
				d.Exclude = c.StringSection(sec, "exclude", "")
				d.Completex = c.StringSection(sec, "completex", "")
			}
		}
		if d.LocalDir == "" {
			LOG.Error("%s dir config error %s", sec, tempD)
			continue
		}
		conf.Dirs[d.LocalDir] = d
	}

	w.Conf = conf
}
func NewWatcher() *Watch {
	w := &Watch{}
	w.Reload(true)
	return w
}
