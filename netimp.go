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
func (s *SyncFileImp) FileInit(cs *FileInitInfo_CS) *FileInitInfo_SC {
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
func (s *SyncFileImp) FileContent(cs *FileConent_CS) *FileConent_SC {
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
func (s *SyncFileImp) NewHandle(*gpbrpc.RPCHelper) SyncFileServer {
	return &SyncFileImp{}
}
