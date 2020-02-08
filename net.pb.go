// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net.proto

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	"time"
	"github.com/gotask/gpbrpc"
	"github.com/gotask/gost/stnet"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileInitInfo_CS struct {
	File                 string   `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	Md5                  string   `protobuf:"bytes,2,opt,name=Md5,proto3" json:"Md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInitInfo_CS) Reset()         { *m = FileInitInfo_CS{} }
func (m *FileInitInfo_CS) String() string { return proto.CompactTextString(m) }
func (*FileInitInfo_CS) ProtoMessage()    {}
func (*FileInitInfo_CS) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_58665dbf8d997fad, []int{0}
}
func (m *FileInitInfo_CS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInitInfo_CS.Unmarshal(m, b)
}
func (m *FileInitInfo_CS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInitInfo_CS.Marshal(b, m, deterministic)
}
func (dst *FileInitInfo_CS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInitInfo_CS.Merge(dst, src)
}
func (m *FileInitInfo_CS) XXX_Size() int {
	return xxx_messageInfo_FileInitInfo_CS.Size(m)
}
func (m *FileInitInfo_CS) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInitInfo_CS.DiscardUnknown(m)
}

var xxx_messageInfo_FileInitInfo_CS proto.InternalMessageInfo

func (m *FileInitInfo_CS) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *FileInitInfo_CS) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

type FileInitInfo_SC struct {
	Need                 bool     `protobuf:"varint,1,opt,name=Need,proto3" json:"Need,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInitInfo_SC) Reset()         { *m = FileInitInfo_SC{} }
func (m *FileInitInfo_SC) String() string { return proto.CompactTextString(m) }
func (*FileInitInfo_SC) ProtoMessage()    {}
func (*FileInitInfo_SC) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_58665dbf8d997fad, []int{1}
}
func (m *FileInitInfo_SC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInitInfo_SC.Unmarshal(m, b)
}
func (m *FileInitInfo_SC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInitInfo_SC.Marshal(b, m, deterministic)
}
func (dst *FileInitInfo_SC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInitInfo_SC.Merge(dst, src)
}
func (m *FileInitInfo_SC) XXX_Size() int {
	return xxx_messageInfo_FileInitInfo_SC.Size(m)
}
func (m *FileInitInfo_SC) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInitInfo_SC.DiscardUnknown(m)
}

var xxx_messageInfo_FileInitInfo_SC proto.InternalMessageInfo

func (m *FileInitInfo_SC) GetNeed() bool {
	if m != nil {
		return m.Need
	}
	return false
}

type FileConent_CS struct {
	File                 string   `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Total                uint32   `protobuf:"varint,3,opt,name=Total,proto3" json:"Total,omitempty"`
	Content              []byte   `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileConent_CS) Reset()         { *m = FileConent_CS{} }
func (m *FileConent_CS) String() string { return proto.CompactTextString(m) }
func (*FileConent_CS) ProtoMessage()    {}
func (*FileConent_CS) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_58665dbf8d997fad, []int{2}
}
func (m *FileConent_CS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileConent_CS.Unmarshal(m, b)
}
func (m *FileConent_CS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileConent_CS.Marshal(b, m, deterministic)
}
func (dst *FileConent_CS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileConent_CS.Merge(dst, src)
}
func (m *FileConent_CS) XXX_Size() int {
	return xxx_messageInfo_FileConent_CS.Size(m)
}
func (m *FileConent_CS) XXX_DiscardUnknown() {
	xxx_messageInfo_FileConent_CS.DiscardUnknown(m)
}

var xxx_messageInfo_FileConent_CS proto.InternalMessageInfo

func (m *FileConent_CS) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *FileConent_CS) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *FileConent_CS) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *FileConent_CS) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type FileConent_SC struct {
	Index                uint32   `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileConent_SC) Reset()         { *m = FileConent_SC{} }
func (m *FileConent_SC) String() string { return proto.CompactTextString(m) }
func (*FileConent_SC) ProtoMessage()    {}
func (*FileConent_SC) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_58665dbf8d997fad, []int{3}
}
func (m *FileConent_SC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileConent_SC.Unmarshal(m, b)
}
func (m *FileConent_SC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileConent_SC.Marshal(b, m, deterministic)
}
func (dst *FileConent_SC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileConent_SC.Merge(dst, src)
}
func (m *FileConent_SC) XXX_Size() int {
	return xxx_messageInfo_FileConent_SC.Size(m)
}
func (m *FileConent_SC) XXX_DiscardUnknown() {
	xxx_messageInfo_FileConent_SC.DiscardUnknown(m)
}

var xxx_messageInfo_FileConent_SC proto.InternalMessageInfo

func (m *FileConent_SC) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*FileInitInfo_CS)(nil), "main.FileInitInfo_CS")
	proto.RegisterType((*FileInitInfo_SC)(nil), "main.FileInitInfo_SC")
	proto.RegisterType((*FileConent_CS)(nil), "main.FileConent_CS")
	proto.RegisterType((*FileConent_SC)(nil), "main.FileConent_SC")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ gpbrpc.RPCClient
var _ stnet.Connect

// SyncFileClient is the client API for SyncFile service.
type SyncFile_FileInit_Exception = func(int32)
type SyncFile_FileInit_Callback = func(*FileInitInfo_SC)
type SyncFile_FileContent_Exception = func(int32)
type SyncFile_FileContent_Callback = func(*FileConent_SC)

type SyncFileClient struct {
	conn      *stnet.Connect
	rpcclient *gpbrpc.RPCClient
	content   map[string]string
	timeout   uint32
	hashfunc  func(*gpbrpc.RPCResponsePacket) int
}

// new client for SyncFile service.
func NewSyncFileClient() *SyncFileClient {
	return &SyncFileClient{nil, nil, make(map[string]string), 5000, nil}
}

// Change the content of this client which is send to service.
func (c *SyncFileClient) SetContent(k, v string) {
	c.content[k] = v
}

func (c *SyncFileClient) DelContent(k string) {
	delete(c.content, k)
}

func (c *SyncFileClient) GetContent() map[string]string {
	return c.content
}

// set timeout of the response from service.
func (c *SyncFileClient) SetTimeout(t uint32) {
	c.timeout = t
}

func (c *SyncFileClient) GetTimeout() uint32 {
	return c.timeout
}

// set rpcclient.
func (c *SyncFileClient) SetRPCClient(rpc *gpbrpc.RPCClient) {
	c.rpcclient = rpc
}

// set connection which is connecting the service.
func (c *SyncFileClient) SetConn(conn *stnet.Connect) {
	c.conn = conn
}

func (c *SyncFileClient) GetConn() *stnet.Connect {
	return c.conn
}

// decide which thread is used to handle the response.
func (c *SyncFileClient) SetHashFunc(f func(*gpbrpc.RPCResponsePacket) int) {
	c.hashfunc = f
}

func (c *SyncFileClient) HashProcessor(rsp *gpbrpc.RPCResponsePacket) int {
	if c.hashfunc != nil {
		return c.hashfunc(rsp)
	}
	return -1
}

// async call functions
func (_c *SyncFileClient) FileInit(in *FileInitInfo_CS, cb SyncFile_FileInit_Callback, exp SyncFile_FileInit_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileInit", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.rpcclient.PushRequest(req)
}

func (_c *SyncFileClient) FileContent(in *FileConent_CS, cb SyncFile_FileContent_Callback, exp SyncFile_FileContent_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileContent", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.rpcclient.PushRequest(req)
}

// sync call functions
func (_c *SyncFileClient) FileInit_Sync(in *FileInitInfo_CS) (out *FileInitInfo_SC, ret int32) {
	out = &FileInitInfo_SC{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileInit", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.rpcclient.PushRequest(req)
	to := time.NewTimer(time.Duration(_c.GetTimeout()) * time.Millisecond)
	select {
	case s := <-sg:
		ret = s.RPCRetCode
		if len(s.RspPayload) > 0 {
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				ret = gpbrpc.GPBUnmarshalFailed
			}
		}
	case <-to.C:
		ret = gpbrpc.SyncCallTimeout
	}
	to.Stop()
	_c.rpcclient.DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *SyncFileClient) FileContent_Sync(in *FileConent_CS) (out *FileConent_SC, ret int32) {
	out = &FileConent_SC{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileContent", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.rpcclient.PushRequest(req)
	to := time.NewTimer(time.Duration(_c.GetTimeout()) * time.Millisecond)
	select {
	case s := <-sg:
		ret = s.RPCRetCode
		if len(s.RspPayload) > 0 {
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				ret = gpbrpc.GPBUnmarshalFailed
			}
		}
	case <-to.C:
		ret = gpbrpc.SyncCallTimeout
	}
	to.Stop()
	_c.rpcclient.DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *SyncFileClient) HandleRSP(r *gpbrpc.RPCRequest, s *gpbrpc.RPCResponsePacket) {
	if r.Req.FuncName == "FileInit" {
		if s.RPCRetCode != 0 {
			if r.Exception.(SyncFile_FileInit_Exception) != nil {
				r.Exception.(SyncFile_FileInit_Exception)(s.RPCRetCode)
			}
		} else {
			out := &FileInitInfo_SC{}
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				if r.Exception.(SyncFile_FileInit_Exception) != nil {
					r.Exception.(SyncFile_FileInit_Exception)(gpbrpc.GPBUnmarshalFailed)
				}
			} else {
				if r.Callback.(SyncFile_FileInit_Callback) != nil {
					r.Callback.(SyncFile_FileInit_Callback)(out)
				}
			}
		}
	} else if r.Req.FuncName == "FileContent" {
		if s.RPCRetCode != 0 {
			if r.Exception.(SyncFile_FileContent_Exception) != nil {
				r.Exception.(SyncFile_FileContent_Exception)(s.RPCRetCode)
			}
		} else {
			out := &FileConent_SC{}
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				if r.Exception.(SyncFile_FileContent_Exception) != nil {
					r.Exception.(SyncFile_FileContent_Exception)(gpbrpc.GPBUnmarshalFailed)
				}
			} else {
				if r.Callback.(SyncFile_FileContent_Callback) != nil {
					r.Callback.(SyncFile_FileContent_Callback)(out)
				}
			}
		}
	}
}

// SyncFileServer is the server API for SyncFile service.
type SyncFileServer interface {
	FileInit(*FileInitInfo_CS) *FileInitInfo_SC
	FileContent(*FileConent_CS) *FileConent_SC
	//Get msg processor by hash
	HashProcessor(req *gpbrpc.RPCRequestPacket) int
	//each thread has one handle
	NewHandle(*gpbrpc.RPCHelper) SyncFileServer
}

type SyncFile struct {
	*gpbrpc.RPCHelper
	SyncFileServer
}

func NewSyncFileServer(s SyncFileServer) *SyncFile {
	return &SyncFile{&gpbrpc.RPCHelper{}, s}
}

func (s *SyncFile) NewHandle() gpbrpc.ServerInterface {
	h := &gpbrpc.RPCHelper{}
	return &SyncFile{h, s.SyncFileServer.NewHandle(h)}
}

func (s *SyncFile) HandleReq(req *gpbrpc.RPCRequestPacket) (msg []byte, ret int32, err error) {
	if req.FuncName == "FileInit" {
		in := &FileInitInfo_CS{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.FileInit(in))
		}
	} else if req.FuncName == "FileContent" {
		in := &FileConent_CS{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.FileContent(in))
		}
	} else {
		ret = gpbrpc.CallRemoteNullFunc
		err = gpbrpc.ErrCallRemoteNullFunc
	}
	return msg, ret, err
}

func init() { proto.RegisterFile("net.proto", fileDescriptor_net_58665dbf8d997fad) }

var fileDescriptor_net_58665dbf8d997fad = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4b, 0x2d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x32, 0xe7, 0xe2, 0x77,
	0xcb, 0xcc, 0x49, 0xf5, 0xcc, 0xcb, 0x2c, 0xf1, 0xcc, 0x4b, 0xcb, 0x8f, 0x77, 0x0e, 0x16, 0x12,
	0xe2, 0x62, 0x01, 0x09, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x02, 0x5c,
	0xcc, 0xbe, 0x29, 0xa6, 0x12, 0x4c, 0x60, 0x21, 0x10, 0x53, 0x49, 0x15, 0x4d, 0x63, 0xb0, 0x33,
	0x48, 0xa3, 0x5f, 0x6a, 0x6a, 0x0a, 0x58, 0x23, 0x47, 0x10, 0x98, 0xad, 0x94, 0xc9, 0xc5, 0x0b,
	0x52, 0xe6, 0x9c, 0x9f, 0x97, 0x9a, 0x57, 0x82, 0xcb, 0x74, 0x11, 0x2e, 0x56, 0xcf, 0xbc, 0x94,
	0xd4, 0x0a, 0xb0, 0xf9, 0xbc, 0x41, 0x10, 0x0e, 0x48, 0x34, 0x24, 0xbf, 0x24, 0x31, 0x47, 0x82,
	0x19, 0x22, 0x0a, 0xe6, 0x08, 0x49, 0x70, 0xb1, 0x3b, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48,
	0xb0, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x4a, 0xaa, 0x28, 0x56, 0x05, 0x3b, 0x23, 0x8c,
	0x65, 0x44, 0x32, 0xd6, 0xa8, 0x96, 0x8b, 0x23, 0xb8, 0x32, 0x2f, 0x19, 0x6c, 0xb1, 0x05, 0x17,
	0x07, 0xcc, 0x13, 0x42, 0xa2, 0x7a, 0xa0, 0x00, 0xd1, 0x43, 0x0b, 0x0d, 0x29, 0x6c, 0xc2, 0xc1,
	0xce, 0x42, 0xe6, 0x5c, 0xdc, 0x50, 0xcb, 0x40, 0x76, 0x0b, 0x09, 0x23, 0x54, 0xc1, 0xbd, 0x2a,
	0x85, 0x29, 0x18, 0xec, 0x9c, 0xc4, 0x06, 0x0e, 0x7d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0x8b, 0x1e, 0x9e, 0x8a, 0x01, 0x00, 0x00,
}
