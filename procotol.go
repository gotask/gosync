// procotol.go
package main

const (
	CmdId_Net_Connect_CS = 10001
	CmdId_Net_Connect_SC = 10002
)

type Cmd_Net_Connect_CS struct {
	Key  string
	File string
	Md5  string
}

type Cmd_Net_Connect_SC struct {
	Key  string
	Need bool
}
