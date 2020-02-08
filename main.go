// gosync project main.go
package main

import (
	"os"
	"os/signal"
)

func main() {
	go Display()
	HttpDownLoadStart()
	w := NewWatcher()
	e := NetStart(w)
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
