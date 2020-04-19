// httpdownload.go
package main

import (
	"context"
	"time"

	"github.com/gotask/gohttpserver"
	"github.com/gotask/gost/stconfig"
	"github.com/gotask/gost/stutil"
)

type HttpDownLoad struct {
	CfMd5 string
	HD    map[string]*gohttpserver.GoHttp
}

func HttpDownLoadStart() {
	go func() {
		hd := &HttpDownLoad{}
		for {
			hd.loadConfig()
			time.Sleep(time.Second * 3)
		}
	}()
}

func (h *HttpDownLoad) loadConfig() {
	md5, _ := stutil.FileMD5("sync.ini")
	if md5 == "" || md5 == h.CfMd5 {
		return
	}
	for _, v := range h.HD {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()
		if err := v.Stop(ctx); nil != err {
			LOG.Error("stop failed %s", err.Error())
			continue
		}
	}
	c, e := stconfig.LoadINI("sync.ini")
	if e != nil {
		LOG.Error("loadconfig failed %s", e.Error())
		return
	}
	h.CfMd5 = md5
	h.HD = make(map[string]*gohttpserver.GoHttp, 0)
	for i := 1; i <= 100; i++ {
		sec := "httpdownload" + stutil.IntToString(int64(i))
		tempD := c.StringSection(sec, "dir", "")
		if tempD == "" {
			continue
		}
		tempA := c.StringSection(sec, "address", "")
		if tempA == "" {
			continue
		}

		Http := &gohttpserver.GoHttp{}
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()
		if err := Http.Start(tempD, tempA, true, false, ctx); nil != err {
			LOG.Error("%s start failed %s", sec, err.Error())
			continue
		}
		LOG.Info("httpdownload start success %s@%s", tempD, tempA)
		h.HD[tempD] = Http
	}
}
