// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	"github.com/gotask/gpbrpc"
	"time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Connect_CS struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect_CS) Reset()         { *m = Connect_CS{} }
func (m *Connect_CS) String() string { return proto.CompactTextString(m) }
func (*Connect_CS) ProtoMessage()    {}
func (*Connect_CS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5b10ce944527a32, []int{0}
}

func (m *Connect_CS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect_CS.Unmarshal(m, b)
}
func (m *Connect_CS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect_CS.Marshal(b, m, deterministic)
}
func (m *Connect_CS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect_CS.Merge(m, src)
}
func (m *Connect_CS) XXX_Size() int {
	return xxx_messageInfo_Connect_CS.Size(m)
}
func (m *Connect_CS) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect_CS.DiscardUnknown(m)
}

var xxx_messageInfo_Connect_CS proto.InternalMessageInfo

func (m *Connect_CS) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

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
	return fileDescriptor_a5b10ce944527a32, []int{1}
}

func (m *FileInitInfo_CS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInitInfo_CS.Unmarshal(m, b)
}
func (m *FileInitInfo_CS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInitInfo_CS.Marshal(b, m, deterministic)
}
func (m *FileInitInfo_CS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInitInfo_CS.Merge(m, src)
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
	return fileDescriptor_a5b10ce944527a32, []int{2}
}

func (m *FileInitInfo_SC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInitInfo_SC.Unmarshal(m, b)
}
func (m *FileInitInfo_SC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInitInfo_SC.Marshal(b, m, deterministic)
}
func (m *FileInitInfo_SC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInitInfo_SC.Merge(m, src)
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
	return fileDescriptor_a5b10ce944527a32, []int{3}
}

func (m *FileConent_CS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileConent_CS.Unmarshal(m, b)
}
func (m *FileConent_CS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileConent_CS.Marshal(b, m, deterministic)
}
func (m *FileConent_CS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileConent_CS.Merge(m, src)
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
	return fileDescriptor_a5b10ce944527a32, []int{4}
}

func (m *FileConent_SC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileConent_SC.Unmarshal(m, b)
}
func (m *FileConent_SC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileConent_SC.Marshal(b, m, deterministic)
}
func (m *FileConent_SC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileConent_SC.Merge(m, src)
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
	proto.RegisterType((*Connect_CS)(nil), "main.Connect_CS")
	proto.RegisterType((*FileInitInfo_CS)(nil), "main.FileInitInfo_CS")
	proto.RegisterType((*FileInitInfo_SC)(nil), "main.FileInitInfo_SC")
	proto.RegisterType((*FileConent_CS)(nil), "main.FileConent_CS")
	proto.RegisterType((*FileConent_SC)(nil), "main.FileConent_SC")
}

func init() { proto.RegisterFile("net.proto", fileDescriptor_a5b10ce944527a32) }

var fileDescriptor_a5b10ce944527a32 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x65, 0x6d, 0xa3, 0xe9, 0x68, 0xb1, 0x8c, 0x15, 0x96, 0x9c, 0xca, 0x42, 0xa1, 0x17, 0x73,
	0x50, 0xa4, 0xde, 0x17, 0x84, 0x1c, 0xf4, 0x90, 0xf5, 0x2e, 0xb5, 0x19, 0x21, 0x58, 0x67, 0x45,
	0xf6, 0xa0, 0xff, 0xe4, 0x47, 0xca, 0xce, 0xb6, 0x46, 0x5b, 0x7b, 0x9b, 0xf7, 0xf2, 0xe6, 0xcd,
	0xcb, 0x5b, 0x18, 0x30, 0x85, 0xf2, 0xed, 0xdd, 0x07, 0x8f, 0xfd, 0xd7, 0x45, 0xcb, 0xc6, 0x00,
	0x58, 0xcf, 0x4c, 0xcb, 0xf0, 0x68, 0x1d, 0x8e, 0x21, 0x7b, 0xf0, 0x2f, 0xc4, 0x5a, 0x4d, 0xd4,
	0x6c, 0x50, 0x27, 0x60, 0xe6, 0x70, 0x7a, 0xdb, 0xae, 0xa8, 0xe2, 0x36, 0x54, 0xfc, 0xec, 0xa3,
	0x10, 0xa1, 0x1f, 0xa9, 0xb5, 0x4e, 0x66, 0x1c, 0x41, 0xef, 0xae, 0xb9, 0xd6, 0x07, 0x42, 0xc5,
	0xd1, 0x4c, 0xb7, 0x16, 0x9d, 0x8d, 0x8b, 0xf7, 0x44, 0x8d, 0x2c, 0xe6, 0xb5, 0xcc, 0xa6, 0x85,
	0x61, 0x94, 0x59, 0xcf, 0xc4, 0x61, 0x9f, 0xfb, 0x18, 0xb2, 0x8a, 0x1b, 0xfa, 0x10, 0xff, 0x61,
	0x9d, 0x40, 0x0a, 0x1c, 0x16, 0x2b, 0xdd, 0x4b, 0xac, 0x00, 0xd4, 0x70, 0x64, 0x3d, 0x07, 0xe2,
	0xa0, 0xb3, 0x89, 0x9a, 0x9d, 0xd4, 0x1b, 0x68, 0xa6, 0x7f, 0x4e, 0x39, 0xdb, 0xd9, 0xaa, 0x5f,
	0xb6, 0x97, 0x5f, 0x0a, 0x72, 0xf7, 0xc9, 0x4b, 0xb9, 0x7c, 0x21, 0x6e, 0xb1, 0x22, 0x1c, 0x95,
	0xb1, 0xb4, 0xb2, 0x6b, 0xac, 0xd8, 0x61, 0xf0, 0x06, 0xf2, 0xcd, 0x4f, 0xe3, 0x79, 0xfa, 0xba,
	0xd5, 0x5e, 0xf1, 0x1f, 0xed, 0x2c, 0xce, 0xe1, 0x78, 0x1d, 0x2e, 0x66, 0xc5, 0xb3, 0x4e, 0xf5,
	0x53, 0x4d, 0xb1, 0x4b, 0x3a, 0xfb, 0x74, 0x28, 0x2f, 0x7a, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0xa8, 0x2a, 0x90, 0xa7, 0xde, 0x01, 0x00, 0x00,
}

// SyncFileClient is the client API for SyncFile service.
type SyncFile_Connect_Exception = func(int32)
type SyncFile_Connect_Callback = func(*Connect_CS)
type SyncFile_FileInit_Exception = func(int32)
type SyncFile_FileInit_Callback = func(*FileInitInfo_SC)
type SyncFile_FileContent_Exception = func(int32)
type SyncFile_FileContent_Callback = func(*FileConent_SC)

type SyncFileClient struct {
	*gpbrpc.RPCBaseClient
}

// new client for SyncFile service.
func NewSyncFileClient() *SyncFileClient {
	return &SyncFileClient{gpbrpc.NewRPCBaseClient()}
}

// async call functions
func (_c *SyncFileClient) Connect(in *Connect_CS, cb SyncFile_Connect_Callback, exp SyncFile_Connect_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "Connect", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.GetRPCClient().PushRequest(req)
}

func (_c *SyncFileClient) FileInit(in *FileInitInfo_CS, cb SyncFile_FileInit_Callback, exp SyncFile_FileInit_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileInit", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.GetRPCClient().PushRequest(req)
}

func (_c *SyncFileClient) FileContent(in *FileConent_CS, cb SyncFile_FileContent_Callback, exp SyncFile_FileContent_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileContent", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.GetRPCClient().PushRequest(req)
}

// sync call functions
func (_c *SyncFileClient) Connect_Sync(in *Connect_CS) (out *Connect_CS, ret int32) {
	out = &Connect_CS{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "Connect", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.GetRPCClient().PushRequest(req)
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
	_c.GetRPCClient().DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *SyncFileClient) FileInit_Sync(in *FileInitInfo_CS) (out *FileInitInfo_SC, ret int32) {
	out = &FileInitInfo_SC{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileInit", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.GetRPCClient().PushRequest(req)
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
	_c.GetRPCClient().DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *SyncFileClient) FileContent_Sync(in *FileConent_CS) (out *FileConent_SC, ret int32) {
	out = &FileConent_SC{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "SyncFile", FuncName: "FileContent", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.GetRPCClient().PushRequest(req)
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
	_c.GetRPCClient().DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *SyncFileClient) HandleRSP(r *gpbrpc.RPCRequest, s *gpbrpc.RPCResponsePacket) {
	if r.Req.FuncName == "Connect" {
		if s.RPCRetCode != 0 {
			if r.Exception.(SyncFile_Connect_Exception) != nil {
				r.Exception.(SyncFile_Connect_Exception)(s.RPCRetCode)
			}
		} else {
			out := &Connect_CS{}
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				if r.Exception.(SyncFile_Connect_Exception) != nil {
					r.Exception.(SyncFile_Connect_Exception)(gpbrpc.GPBUnmarshalFailed)
				}
			} else {
				if r.Callback.(SyncFile_Connect_Callback) != nil {
					r.Callback.(SyncFile_Connect_Callback)(out)
				}
			}
		}
	} else if r.Req.FuncName == "FileInit" {
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

// SyncFileService is the server API for SyncFile service.
type SyncFileService interface {
	Connect(*gpbrpc.RPCHelper, *Connect_CS) *Connect_CS
	FileInit(*gpbrpc.RPCHelper, *FileInitInfo_CS) *FileInitInfo_SC
	FileContent(*gpbrpc.RPCHelper, *FileConent_CS) *FileConent_SC
	//Get msg processor by hash
	HashProcessor(req *gpbrpc.RPCRequestPacket) int
}

type SyncFile struct {
	*gpbrpc.RPCHelper
	SyncFileService
}

func NewSyncFileService(s SyncFileService) gpbrpc.ServerInterface {
	return &SyncFile{&gpbrpc.RPCHelper{}, s}
}

func (s *SyncFile) NewHandle() gpbrpc.ServerInterface {
	return &SyncFile{&gpbrpc.RPCHelper{}, gpbrpc.NewStruct(s.SyncFileService).(SyncFileService)}
}

func (s *SyncFile) HandleReq(req *gpbrpc.RPCRequestPacket) (msg []byte, ret int32, err error) {
	if req.FuncName == "Connect" {
		in := &Connect_CS{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.Connect(s.RPCHelper, in))
		}
	} else if req.FuncName == "FileInit" {
		in := &FileInitInfo_CS{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.FileInit(s.RPCHelper, in))
		}
	} else if req.FuncName == "FileContent" {
		in := &FileConent_CS{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.FileContent(s.RPCHelper, in))
		}
	} else {
		ret = gpbrpc.CallRemoteNullFunc
		err = gpbrpc.ErrCallRemoteNullFunc
	}
	return msg, ret, err
}
