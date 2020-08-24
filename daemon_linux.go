// deamon_linux.go
package main

import (
	"github.com/gotask/gost/stutil"
)

func deamon(title, info string) {
	stutil.SysDaemon()
}
