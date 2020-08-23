// gosync project main.go
package main

import (
	"os"
	"os/signal"

	"github.com/gotask/gost/stconfig"
)

func main() {
	go Display()
	cfg, e := stconfig.LoadINI("sync.ini")
	uiaddr := ""
	if e == nil {
		uiaddr = cfg.StringSection("system", "uiaddr", "")
	}
	go deamon("GoSync", "config http at: "+uiaddr)
	HttpDownLoadStart()
	w := NewWatcher()
	e = NetStart(w)
	if e != nil {
		LOG.Error("start failed %s", e.Error())
		LOG.Close()
		return
	}
	w.Start()

	c := make(chan os.Signal, 0)
	signal.Notify(c)
	<-c

	SVR.Stop()
}
