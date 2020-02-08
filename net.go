// net.go
package main

import (
	"strings"

	. "github.com/gotask/gost/stnet"
	"github.com/gotask/gost/stutil"
	"github.com/gotask/gpbrpc"
)

var (
	SVR     *gpbrpc.GpbServer
	SyncW   *Watch
	NET_LEN uint32 = 1024 * 1024
)

func NetStart(w *Watch) error {
	LOG.Info("listen@%s", w.Conf.Address)
	SVR = gpbrpc.NewGpbServer("sync-server", 100)
	_, e := SVR.AddGpbService("loop", "", 0, &SyncServer{}, 0)
	if e != nil {
		return e
	}
	e = SVR.AddRpcService("server", w.Conf.Address, 0, NewSyncFileServer(&SyncFileImp{}), 0)
	if e != nil {
		return e
	}
	e = SVR.AddRpcClient("rpcclient", 0)
	if e != nil {
		return e
	}
	SyncW = w
	return SVR.Start()
}

type SyncConn struct {
	conn   *SyncFileClient
	ev     *SyncEvent
	ok     bool
	failed bool
}
type SyncEvent struct {
	Event
	ProcessNum int
}
type SyncServer struct {
	Ev    map[string]*SyncEvent
	Conns map[string]*SyncConn
}

func (s *SyncServer) Init() bool {
	s.Ev = make(map[string]*SyncEvent)
	s.Conns = make(map[string]*SyncConn)
	return true
}
func dirKey(path string, Conf DirConfig) string {
	return path + "@" + Conf.RemoteAddr + "@" + Conf.RemoteDir
}
func (s *SyncServer) Loop() {
	n := 0
	for n < 1024 {
		n++
		select {
		case ev := <-SyncW.Ev:
			{
				for i, v := range ev {
					if v.Op != Init {
						LOG.Debug("file change %d %v", i, v)
					}
					key := dirKey(v.Path, v.Conf)
					if e, ok := s.Ev[key]; ok {
						if e.Info.ModTime().Before(v.Info.ModTime()) {
							s.Ev[key] = &SyncEvent{v, 0}
						}
					} else {
						s.Ev[key] = &SyncEvent{v, 0}
					}
				}
			}
		default:
		}
	}

	oldEvLen := len(s.Ev)

	for k, v := range s.Conns {
		if v.ok || (v.failed && v.ev.ProcessNum >= 3) {
			if e, o := s.Ev[k]; o {
				if e.Info.ModTime().UnixNano() <= v.ev.Info.ModTime().UnixNano() {
					delete(s.Ev, k)
				}
			} else {
				LOG.Error("no key %s", k)
			}
			v.conn.GetConn().Close()
			delete(s.Conns, k)
		} else if v.failed {
			LOG.Debug("sync failed %s", k)
			v.conn.GetConn().Close()
			delete(s.Conns, k)
		}
	}

	for k, v := range s.Ev {
		if len(s.Conns) > 10 {
			//LOG.Info("connection is too many %d", len(s.Conns))
			return
		}
		if _, ok := s.Conns[k]; !ok {

			info := &SyncConn{NewSyncFileClient(), v, false, false}
			s.Conns[k] = info
			go func(k string, v *SyncEvent) {
				v.ProcessNum++
				SVR.NewRpcConnector("rpcclient", v.Conf.RemoteAddr, info.conn)

				f := strings.Replace(v.Path, v.Conf.LocalDir, v.Conf.RemoteDir, -1)
				md5 := ""
				var e error
				if v.Op != Delete {
					md5, e = stutil.FileMD5(v.Path)
					if e != nil {
						info.failed = true
						LOG.Error("md5 read failed %s %s %s", k, e.Error(), v.Op)
						return
					}
				}
				r, i := info.conn.FileInit_Sync(&FileInitInfo_CS{File: f, Md5: md5})
				if i != 0 {
					info.failed = true
					LOG.Error("FileInit_Sync failed %s %d", k, i)
					return
				}
				if !r.Need {
					info.ok = true
					//LOG.Error("FileInit_Sync same md5 %s %d", k, i)
					if v.Op != Init {
						LOG.Info("FileInit_Sync success %s", k)
					}
					return
				}
				b, e := stutil.FileReadAll(v.Path)
				if e != nil {
					info.failed = true
					LOG.Error("FileInit_Sync readfile failed %s %s", k, e.Error())
					return
				}
				var n uint32
				t := uint32(len(b))
				for ; n < t; n += NET_LEN {
					fend := n + NET_LEN
					if fend > t {
						fend = t
					}
					_, i := info.conn.FileContent_Sync(&FileConent_CS{File: f, Index: n, Total: t, Content: b[n:fend]})
					if i != 0 {
						info.failed = true
						LOG.Error("FileContent_Sync failed %s %d", k, i)
						return
					} else {
						//LOG.Info("sync ok %s index %d total %d", k, r.Index, t)
					}
				}
				info.ok = true
				LOG.Info("FileContent_Sync success %s", k)
			}(k, v)

		}
	}

	newEvLen := len(s.Ev)
	if oldEvLen != newEvLen && newEvLen == 0 {
		LOG.Info("sync complete!!!!!!")
	}
}
func (s *SyncServer) Destroy() {

}
func (s *SyncServer) HandleMessage(current *CurrentContent, msg *gpbrpc.ProtocolRequest) {
}
func (s *SyncServer) HashProcessor(sess *Session, msg *gpbrpc.ProtocolRequest) (processorID int) {
	return int(sess.GetID() % uint64(ProcessorThreadsNum))
}
func (s *SyncServer) SessionOpen(sess *Session) {

}
func (s *SyncServer) SessionClose(sess *Session) {

}
