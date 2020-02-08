// log.go
package main

import (
	"github.com/gotask/gost/stlog"
)

var LOG *stlog.Logger

func init() {
	LOG = stlog.NewFileLogger("sync.log")
	//LOG.SetTermLevel(stlog.CRITICAL)
}
