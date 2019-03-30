// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Signal_Control int32

const (
	Signal_NOT_USED  Signal_Control = 0
	Signal_START     Signal_Control = 1
	Signal_TERMINATE Signal_Control = 2
	Signal_LINKED    Signal_Control = 3
	Signal_MESSAGE   Signal_Control = 4
	Signal_BLOCKED   Signal_Control = 5
)

var Signal_Control_name = map[int32]string{
	0: "NOT_USED",
	1: "START",
	2: "TERMINATE",
	3: "LINKED",
	4: "MESSAGE",
	5: "BLOCKED",
}

var Signal_Control_value = map[string]int32{
	"NOT_USED":  0,
	"START":     1,
	"TERMINATE": 2,
	"LINKED":    3,
	"MESSAGE":   4,
	"BLOCKED":   5,
}

func (x Signal_Control) String() string {
	return proto.EnumName(Signal_Control_name, int32(x))
}

func (Signal_Control) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5, 0}
}

type Token struct {
	// container id
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LinkRequest struct {
	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// required(id, name)
	Src *Station `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	// required(name, image)
	Dst                  *Station `protobuf:"bytes,3,opt,name=dst,proto3" json:"dst,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkRequest) Reset()         { *m = LinkRequest{} }
func (m *LinkRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRequest) ProtoMessage()    {}
func (*LinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *LinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkRequest.Unmarshal(m, b)
}
func (m *LinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkRequest.Marshal(b, m, deterministic)
}
func (m *LinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkRequest.Merge(m, src)
}
func (m *LinkRequest) XXX_Size() int {
	return xxx_messageInfo_LinkRequest.Size(m)
}
func (m *LinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LinkRequest proto.InternalMessageInfo

func (m *LinkRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *LinkRequest) GetSrc() *Station {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *LinkRequest) GetDst() *Station {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *LinkRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type BlockRequest struct {
	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// required(id, name)
	Src *Station `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	// required(name, image)
	Dst                  *Station `protobuf:"bytes,3,opt,name=dst,proto3" json:"dst,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockRequest) Reset()         { *m = BlockRequest{} }
func (m *BlockRequest) String() string { return proto.CompactTextString(m) }
func (*BlockRequest) ProtoMessage()    {}
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *BlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRequest.Unmarshal(m, b)
}
func (m *BlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRequest.Marshal(b, m, deterministic)
}
func (m *BlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRequest.Merge(m, src)
}
func (m *BlockRequest) XXX_Size() int {
	return xxx_messageInfo_BlockRequest.Size(m)
}
func (m *BlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRequest proto.InternalMessageInfo

func (m *BlockRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *BlockRequest) GetSrc() *Station {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *BlockRequest) GetDst() *Station {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *BlockRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TransmitRequest struct {
	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// required(id, name)
	Src *Station `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	// required(name, image)
	Dst                  *Station `protobuf:"bytes,3,opt,name=dst,proto3" json:"dst,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransmitRequest) Reset()         { *m = TransmitRequest{} }
func (m *TransmitRequest) String() string { return proto.CompactTextString(m) }
func (*TransmitRequest) ProtoMessage()    {}
func (*TransmitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *TransmitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransmitRequest.Unmarshal(m, b)
}
func (m *TransmitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransmitRequest.Marshal(b, m, deterministic)
}
func (m *TransmitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransmitRequest.Merge(m, src)
}
func (m *TransmitRequest) XXX_Size() int {
	return xxx_messageInfo_TransmitRequest.Size(m)
}
func (m *TransmitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransmitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransmitRequest proto.InternalMessageInfo

func (m *TransmitRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *TransmitRequest) GetSrc() *Station {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *TransmitRequest) GetDst() *Station {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *TransmitRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListenRequest struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenRequest) Reset()         { *m = ListenRequest{} }
func (m *ListenRequest) String() string { return proto.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()    {}
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{4}
}

func (m *ListenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenRequest.Unmarshal(m, b)
}
func (m *ListenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenRequest.Marshal(b, m, deterministic)
}
func (m *ListenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenRequest.Merge(m, src)
}
func (m *ListenRequest) XXX_Size() int {
	return xxx_messageInfo_ListenRequest.Size(m)
}
func (m *ListenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenRequest proto.InternalMessageInfo

func (m *ListenRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

// required                 src dst
// for control=START,        o   y
// for control=LINKED,       y   y
// for control=MESSAGE       y   y
type Signal struct {
	// required(id, name, image)
	Src *Station `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	// required(id, name, image)
	Dst *Station `protobuf:"bytes,3,opt,name=dst,proto3" json:"dst,omitempty"`
	// always optional
	Message              string         `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Control              Signal_Control `protobuf:"varint,5,opt,name=control,proto3,enum=loko.metro.api.Signal_Control" json:"control,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Signal) Reset()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) ProtoMessage()    {}
func (*Signal) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5}
}

func (m *Signal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signal.Unmarshal(m, b)
}
func (m *Signal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signal.Marshal(b, m, deterministic)
}
func (m *Signal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signal.Merge(m, src)
}
func (m *Signal) XXX_Size() int {
	return xxx_messageInfo_Signal.Size(m)
}
func (m *Signal) XXX_DiscardUnknown() {
	xxx_messageInfo_Signal.DiscardUnknown(m)
}

var xxx_messageInfo_Signal proto.InternalMessageInfo

func (m *Signal) GetSrc() *Station {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *Signal) GetDst() *Station {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *Signal) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Signal) GetControl() Signal_Control {
	if m != nil {
		return m.Control
	}
	return Signal_NOT_USED
}

func init() {
	proto.RegisterEnum("loko.metro.api.Signal_Control", Signal_Control_name, Signal_Control_value)
	proto.RegisterType((*Token)(nil), "loko.metro.api.Token")
	proto.RegisterType((*LinkRequest)(nil), "loko.metro.api.LinkRequest")
	proto.RegisterType((*BlockRequest)(nil), "loko.metro.api.BlockRequest")
	proto.RegisterType((*TransmitRequest)(nil), "loko.metro.api.TransmitRequest")
	proto.RegisterType((*ListenRequest)(nil), "loko.metro.api.ListenRequest")
	proto.RegisterType((*Signal)(nil), "loko.metro.api.Signal")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x9b, 0xb4, 0x49, 0xb7, 0xa7, 0xdd, 0x1a, 0x0e, 0xe8, 0x86, 0xaa, 0x6b, 0xc9, 0xd5,
	0x8a, 0x90, 0x4a, 0xbd, 0xf1, 0x42, 0x90, 0xb4, 0x1b, 0x96, 0xc5, 0x6e, 0x17, 0x26, 0x11, 0xc1,
	0x1b, 0xc9, 0xa6, 0x43, 0x1d, 0xda, 0xcc, 0xd4, 0x99, 0xe9, 0x53, 0xf8, 0x26, 0x0a, 0xbe, 0x91,
	0xef, 0x22, 0x99, 0x5a, 0xe8, 0x36, 0x4b, 0x05, 0x2f, 0x16, 0xf6, 0xf2, 0xe4, 0x7c, 0xfc, 0xfc,
	0x87, 0xfc, 0xff, 0x40, 0x47, 0x8a, 0xb5, 0xa6, 0x32, 0x5c, 0x49, 0xa1, 0x05, 0x76, 0x97, 0x62,
	0x21, 0xc2, 0x82, 0x6a, 0x29, 0xc2, 0x6c, 0xc5, 0x7a, 0x9d, 0x5c, 0x14, 0x85, 0xe0, 0x9b, 0x6d,
	0x70, 0x02, 0x4e, 0x2a, 0x16, 0x94, 0x63, 0x17, 0x6c, 0x36, 0xf3, 0xad, 0xbe, 0x75, 0xd6, 0x22,
	0x36, 0x9b, 0x05, 0x3f, 0x2c, 0x68, 0x4f, 0x18, 0x5f, 0x10, 0xfa, 0x6d, 0x4d, 0x95, 0xc6, 0x57,
	0xe0, 0xe8, 0x12, 0x34, 0x48, 0x7b, 0xf8, 0x38, 0xbc, 0x2d, 0x1b, 0x1a, 0x15, 0xb2, 0x61, 0xf0,
	0x25, 0xd4, 0x95, 0xcc, 0x7d, 0xdb, 0xa0, 0x27, 0xfb, 0x68, 0xa2, 0x33, 0xcd, 0x04, 0x27, 0x25,
	0x53, 0xa2, 0x33, 0xa5, 0xfd, 0xfa, 0x3f, 0xd0, 0x99, 0xd2, 0xe8, 0x43, 0xb3, 0xa0, 0x4a, 0x65,
	0x73, 0xea, 0x37, 0x8c, 0xcf, 0xed, 0x18, 0xfc, 0xb4, 0xa0, 0x33, 0x5a, 0x8a, 0xfc, 0x61, 0xb8,
	0xfd, 0x65, 0xc1, 0xa3, 0x54, 0x66, 0x5c, 0x15, 0x4c, 0x3f, 0x08, 0xc3, 0xef, 0xe0, 0x78, 0xc2,
	0x94, 0xa6, 0xfc, 0x7f, 0xdc, 0x06, 0xdf, 0x6d, 0x70, 0x13, 0x36, 0xe7, 0xd9, 0xf2, 0xbe, 0x8d,
	0xe3, 0x5b, 0x68, 0xe6, 0x82, 0x6b, 0x29, 0x96, 0xbe, 0xd3, 0xb7, 0xce, 0xba, 0xc3, 0xd3, 0x8a,
	0x90, 0x31, 0x16, 0x8e, 0x37, 0x14, 0xd9, 0xe2, 0xc1, 0x27, 0x68, 0xfe, 0xfd, 0x86, 0x1d, 0x38,
	0x9a, 0x5e, 0xa7, 0x5f, 0x3e, 0x26, 0xf1, 0xb9, 0x57, 0xc3, 0x16, 0x38, 0x49, 0x1a, 0x91, 0xd4,
	0xb3, 0xf0, 0x18, 0x5a, 0x69, 0x4c, 0xae, 0x2e, 0xa7, 0x51, 0x1a, 0x7b, 0x36, 0x02, 0xb8, 0x93,
	0xcb, 0xe9, 0x87, 0xf8, 0xdc, 0xab, 0x63, 0x1b, 0x9a, 0x57, 0x71, 0x92, 0x44, 0x17, 0xb1, 0xd7,
	0x28, 0x87, 0xd1, 0xe4, 0x7a, 0x5c, 0x6e, 0x9c, 0xe1, 0x6f, 0x1b, 0x5c, 0x62, 0xfa, 0x89, 0x11,
	0x38, 0x89, 0xce, 0xa4, 0xc6, 0x67, 0x77, 0x9c, 0x27, 0xb7, 0xd1, 0xe8, 0xf9, 0xfb, 0x5b, 0x42,
	0xd5, 0x4a, 0x70, 0x45, 0x83, 0x1a, 0xbe, 0x87, 0x46, 0x59, 0x52, 0x7c, 0xba, 0xcf, 0xec, 0x54,
	0xf7, 0xa0, 0x40, 0x04, 0x8e, 0x29, 0x4e, 0xd5, 0xc3, 0x6e, 0x9f, 0x0e, 0x4a, 0x5c, 0xc0, 0xd1,
	0x36, 0xcd, 0xf8, 0xa2, 0x92, 0x84, 0xdb, 0x39, 0x3f, 0x28, 0x34, 0x06, 0x77, 0x13, 0x33, 0x7c,
	0x5e, 0x3d, 0x67, 0x27, 0x7e, 0xbd, 0x27, 0x77, 0xff, 0xc5, 0xa0, 0xf6, 0xda, 0x1a, 0xf5, 0x3f,
	0x9f, 0xce, 0x99, 0xfe, 0xba, 0xbe, 0x09, 0x73, 0x51, 0x0c, 0x4a, 0xae, 0x10, 0x9a, 0xaa, 0x81,
	0x81, 0x07, 0xd9, 0x8a, 0xdd, 0xb8, 0xe6, 0xe5, 0x7b, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0xf6, 0xb6, 0xac, 0x27, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*Response, error)
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*Response, error)
	Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*Response, error)
	Transmit(ctx context.Context, in *TransmitRequest, opts ...grpc.CallOption) (*Response, error)
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (Router_ListenClient, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/loko.metro.api.Router/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/loko.metro.api.Router/Link", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/loko.metro.api.Router/Block", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Transmit(ctx context.Context, in *TransmitRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/loko.metro.api.Router/Transmit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (Router_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Router_serviceDesc.Streams[0], "/loko.metro.api.Router/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_ListenClient interface {
	Recv() (*Signal, error)
	grpc.ClientStream
}

type routerListenClient struct {
	grpc.ClientStream
}

func (x *routerListenClient) Recv() (*Signal, error) {
	m := new(Signal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Start(context.Context, *StartRequest) (*Response, error)
	Link(context.Context, *LinkRequest) (*Response, error)
	Block(context.Context, *BlockRequest) (*Response, error)
	Transmit(context.Context, *TransmitRequest) (*Response, error)
	Listen(*ListenRequest, Router_ListenServer) error
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) Start(ctx context.Context, req *StartRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedRouterServer) Link(ctx context.Context, req *LinkRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (*UnimplementedRouterServer) Block(ctx context.Context, req *BlockRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}
func (*UnimplementedRouterServer) Transmit(ctx context.Context, req *TransmitRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transmit not implemented")
}
func (*UnimplementedRouterServer) Listen(req *ListenRequest, srv Router_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loko.metro.api.Router/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loko.metro.api.Router/Link",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Link(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loko.metro.api.Router/Block",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Block(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Transmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Transmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loko.metro.api.Router/Transmit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Transmit(ctx, req.(*TransmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).Listen(m, &routerListenServer{stream})
}

type Router_ListenServer interface {
	Send(*Signal) error
	grpc.ServerStream
}

type routerListenServer struct {
	grpc.ServerStream
}

func (x *routerListenServer) Send(m *Signal) error {
	return x.ServerStream.SendMsg(m)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loko.metro.api.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Router_Start_Handler,
		},
		{
			MethodName: "Link",
			Handler:    _Router_Link_Handler,
		},
		{
			MethodName: "Block",
			Handler:    _Router_Block_Handler,
		},
		{
			MethodName: "Transmit",
			Handler:    _Router_Transmit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _Router_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "router.proto",
}
