// gosync project main.go
package main

import (
	"github.com/gotask/gost/stconfig"
	"github.com/gotask/gost/stutil"
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

	stutil.SysWaitSignal()

	SVR.Stop()
}
