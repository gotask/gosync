// gosync project main.go
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gotask/gost/stconfig"
)

var (
	EXIT = make(chan os.Signal, 1)
)

func main() {
	cfg, e := stconfig.LoadINI("sync.ini")
	uiaddr := ""
	if e == nil {
		uiaddr = cfg.StringSection("system", "uiaddr", "")
	}
	go deamon("GoSync", "config http at: "+uiaddr)
	go Display()
	HttpDownLoadStart()
	w := NewWatcher()
	e = NetStart(w)
	if e != nil {
		LOG.Error("start failed %s", e.Error())
		LOG.Close()
		return
	}
	w.Start()

	signal.Notify(EXIT, syscall.SIGINT, syscall.SIGTERM)
	signal.Ignore(syscall.SIGPIPE)

	select {
	case <-EXIT:
	}

	SVR.Stop()
}
