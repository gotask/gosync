// ui.go
package main

import (
	"strings"

	"github.com/gotask/gost/stconfig"
	"github.com/gotask/gost/stutil"
	"github.com/gotask/ssui"
)

func Display() error {
	c, e := stconfig.LoadINI("sync.ini")
	if e != nil {
		stutil.FileCreateAndWrite("sync.ini", `[system]
loopsec = 1
address = 0.0.0.0:3030
token = 
uiaddr = 0.0.0.0:2020

[httpdownload1]
dir = .
address = 0.0.0.0:2021

[sync1]
dir = D:\myproject=>/home/xxx/myproject@192.168.1.111:3030
include = .*\.(cpp|h|go)
exclude = .*
completex = build
`)
		c, e = stconfig.LoadINI("sync.ini")
		if e != nil {
			return e
		}
	}

	uiaddr := c.StringSection("system", "uiaddr", "")
	if uiaddr == "" {
		return nil
	}

	app := ssui.NewApp(uiaddr)
	login := ssui.NewFrame("/login", "gosync", "gosync", "gosync", func(token string) {
	}, func(token string) (success bool, tokenFailedUrl string) {
		return true, ""
	})
	login.AddElem(ssui.NewLabel("Token")).AddElem(ssui.NewLineEdit("loginToken", "token", "", false)).AddElem(ssui.NewButton("loginReq", "验证", func(param map[string]string) *ssui.HResponse {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return ssui.ResponseURL("/", "")
		}
		t := c.StringSection("system", "token", "")
		token := ssui.Value("loginToken", param)
		if t != token {
			return ssui.ResponseError("验证失败", "")
		}
		return ssui.ResponseURL("/", token)
	}))
	f := ssui.NewFrame("/", "gosync", "gosync", "gosync", func(token string) {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return
		}
		app.GetElem(token, "/", "loopsec").(*ssui.HLineEdit).Text = c.StringSection("system", "loopsec", "1")
		app.GetElem(token, "/", "address").(*ssui.HLineEdit).Text = c.StringSection("system", "address", ":3030")
		httptable := app.GetElem(token, "/", "httpdownload").(*ssui.HTable)
		httptable.Reset()
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
			httptable.AddRow(sec, []string{sec, tempD, tempA})
		}

		synctable := app.GetElem(token, "/", "sync").(*ssui.HTable)
		synctable.Reset()
		for i := 1; i <= 100; i++ {
			sec := "sync" + stutil.IntToString(int64(i))
			di := c.StringSection(sec, "dir", "")
			if di == "" {
				continue
			}
			in := c.StringSection(sec, "include", "")
			ex := c.StringSection(sec, "exclude", "")
			cex := c.StringSection(sec, "completex", "")
			synctable.AddRow(sec, []string{sec, di, in, ex, cex})
		}
	}, func(token string) (success bool, tokenFailedUrl string) {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return true, ""
		}
		t := c.StringSection("system", "token", "")
		if t == token {
			return true, ""
		}
		return false, "/login"
	})
	f.AddElem(ssui.NewLegend("HTTP"))
	f.AddElem(ssui.NewLabel("HTTP下载配置表"))
	f.AddElem(ssui.NewTable("httpdownload", []string{"ID(格式不能变)", "下载目录", "http地址"}, true, true, func(t *ssui.HTable, rowid string) *ssui.HResponse {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return ssui.ResponseError(e.Error(), "")
		}
		c.DelSection(rowid)
		c.Save()
		return ssui.ResponseURL("/", "")
	}, func(t *ssui.HTable, cols []string) *ssui.HResponse {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return ssui.ResponseError(e.Error(), "")
		}
		if !strings.HasPrefix(cols[0], "httpdownload") || cols[1] == "" || cols[2] == "" {
			return ssui.ResponseError("error param", "")
		}
		c.SectionSet(cols[0], "dir", cols[1], "")
		c.SectionSet(cols[0], "address", cols[2], "")
		c.Save()
		return ssui.ResponseURL("/", "")
	}))
	f.AddElem(ssui.NewLegend("SYNC"))
	f.AddElem(ssui.NewRow().AddElem(ssui.NewLabel("sync同步配置表")).AddElem(ssui.NewButton("resync", "重新同步", func(param map[string]string) *ssui.HResponse {
		SyncW.Reload(true)
		return ssui.ResponseMsg("开始同步，请观察日志...", "")
	})).AddElem(ssui.NewLabel("循环间隔")).AddElem(ssui.NewLineEdit("loopsec", "s", c.StringSection("system", "loopsec", "1"), false)).AddElem(
		ssui.NewLabel("监听地址")).AddElem(ssui.NewLineEdit("address", ":3030", c.StringSection("system", "address", ":3030"), false)).AddElem(ssui.NewButton("change", "提交", func(param map[string]string) *ssui.HResponse {
		loopsec := ssui.Value("loopsec", param)
		address := ssui.Value("address", param)
		if loopsec == "" || address == "" {
			return ssui.ResponseError("error param", "")
		}
		c.SectionSet("system", "loopsec", loopsec, "")
		c.SectionSet("system", "address", address, "")
		c.Save()
		return ssui.ResponseMsg("保存成功，重启生效", "")
	})))
	f.AddElem(ssui.NewTable("sync", []string{"ID(格式不能变)", "同步地址", "Include", "Exclude", "CompleteExclude"}, true, true, func(t *ssui.HTable, rowid string) *ssui.HResponse {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return ssui.ResponseError(e.Error(), "")
		}
		c.DelSection(rowid)
		c.Save()
		return ssui.ResponseURL("/", "")
	}, func(t *ssui.HTable, cols []string) *ssui.HResponse {
		c, e := stconfig.LoadINI("sync.ini")
		if e != nil {
			return ssui.ResponseError(e.Error(), "")
		}
		if !strings.HasPrefix(cols[0], "sync") || cols[1] == "" {
			return ssui.ResponseError("error param", "")
		}
		c.SectionSet(cols[0], "dir", cols[1], "")
		c.SectionSet(cols[0], "include", cols[2], "")
		c.SectionSet(cols[0], "exclude", cols[3], "")
		c.SectionSet(cols[0], "completex", cols[4], "")
		c.Save()
		return ssui.ResponseURL("/", "")
	}))
	f.AddElem(ssui.NewLegend("LOG"))
	f.AddElem(ssui.NewButton("log", "日志", func(param map[string]string) *ssui.HResponse {
		buf, e := stutil.FileReadAll("sync.log")
		if e != nil {
			return ssui.ResponseError(e.Error(), "")
		}
		ret := ""
		if len(buf) < 1024 {
			ret = string(buf)
		} else {
			ret = string(buf[len(buf)-1024:])
		}
		return ssui.ResponseText(ret, "")
	}))
	return app.AddFrame(f).AddFrame(login).Run()
}
