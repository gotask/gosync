// netimp.go
package main

import (
	"os"
	"strings"

	"github.com/gotask/gost/stutil"

	"github.com/gotask/gpbrpc"
)

var (
	FileTemp = make(map[string][]byte)
)

type SyncFileImp struct {
}

func getFile(s string) string {
	r := strings.Replace(s, "/", string(os.PathSeparator), -1)
	r = strings.Replace(r, "\\", string(os.PathSeparator), -1)
	return r
}

func (s *SyncFileImp) Connect(h *gpbrpc.RPCHelper, c *Connect_CS) *Connect_CS {
	if c.GetToken() != Token {
		h.GetCurrent().Sess.Close()
		LOG.Error("error token: %s", c.GetToken())
		return &Connect_CS{}
	} else {
		h.GetCurrent().Sess.UserData = true
	}
	return c
}
func (s *SyncFileImp) FileInit(h *gpbrpc.RPCHelper, cs *FileInitInfo_CS) *FileInitInfo_SC {
	if Token != "" && !h.GetCurrent().Sess.UserData.(bool) {
		h.GetCurrent().Sess.Close()
		LOG.Error("first proto is not Connect")
		return nil
	}
	f := getFile(cs.File)
	if cs.Md5 == "" { //delete
		stutil.FileDelete(f)
	} else {
		md5, _ := stutil.FileMD5(f)
		if md5 != cs.Md5 {
			return &FileInitInfo_SC{Need: true}
		}
	}
	return &FileInitInfo_SC{Need: false}
}
func (s *SyncFileImp) FileContent(h *gpbrpc.RPCHelper, cs *FileConent_CS) *FileConent_SC {
	if Token != "" && !h.GetCurrent().Sess.UserData.(bool) {
		h.GetCurrent().Sess.Close()
		LOG.Error("first proto is not Connect")
		return nil
	}
	f := getFile(cs.File)
	if cs.Index == 0 {
		FileTemp[f] = make([]byte, cs.Total, cs.Total)
	}
	copy(FileTemp[f][cs.Index:], cs.Content)
	if cs.Index+NET_LEN >= cs.Total {
		stutil.FileCreateAndWrite(f, string(FileTemp[f]))
		delete(FileTemp, f)
	}
	return &FileConent_SC{Index: cs.Index}
}

func (s *SyncFileImp) HashProcessor(req *gpbrpc.RPCRequestPacket) int {
	return 0
}
