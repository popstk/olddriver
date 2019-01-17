// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend/backend.proto

package backend

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type SearchRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Keyword              string   `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_1549b5b0671b72ea, []int{0}
}
func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (dst *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(dst, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SearchRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type SearchReply struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Href                 string   `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Baidu                []string `protobuf:"bytes,4,rep,name=baidu,proto3" json:"baidu,omitempty"`
	Magnet               []string `protobuf:"bytes,5,rep,name=magnet,proto3" json:"magnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReply) Reset()         { *m = SearchReply{} }
func (m *SearchReply) String() string { return proto.CompactTextString(m) }
func (*SearchReply) ProtoMessage()    {}
func (*SearchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_1549b5b0671b72ea, []int{1}
}
func (m *SearchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReply.Unmarshal(m, b)
}
func (m *SearchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReply.Marshal(b, m, deterministic)
}
func (dst *SearchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReply.Merge(dst, src)
}
func (m *SearchReply) XXX_Size() int {
	return xxx_messageInfo_SearchReply.Size(m)
}
func (m *SearchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReply.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReply proto.InternalMessageInfo

func (m *SearchReply) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SearchReply) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *SearchReply) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *SearchReply) GetBaidu() []string {
	if m != nil {
		return m.Baidu
	}
	return nil
}

func (m *SearchReply) GetMagnet() []string {
	if m != nil {
		return m.Magnet
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "backend.SearchRequest")
	proto.RegisterType((*SearchReply)(nil), "backend.SearchReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpiderClient is the client API for Spider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpiderClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
}

type spiderClient struct {
	cc *grpc.ClientConn
}

func NewSpiderClient(cc *grpc.ClientConn) SpiderClient {
	return &spiderClient{cc}
}

func (c *spiderClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, "/backend.Spider/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpiderServer is the server API for Spider service.
type SpiderServer interface {
	Search(context.Context, *SearchRequest) (*SearchReply, error)
}

func RegisterSpiderServer(s *grpc.Server, srv SpiderServer) {
	s.RegisterService(&_Spider_serviceDesc, srv)
}

func _Spider_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiderServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Spider/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiderServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Spider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.Spider",
	HandlerType: (*SpiderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Spider_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/backend.proto",
}

func init() { proto.RegisterFile("backend/backend.proto", fileDescriptor_backend_1549b5b0671b72ea) }

var fileDescriptor_backend_1549b5b0671b72ea = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3f, 0x4f, 0x87, 0x30,
	0x10, 0x86, 0xfd, 0x09, 0x94, 0x78, 0xc6, 0xe5, 0x82, 0xa4, 0x71, 0x22, 0x4c, 0x4e, 0x98, 0xe8,
	0xe2, 0xe2, 0xe2, 0x47, 0x80, 0x4f, 0x50, 0xe8, 0x29, 0x0d, 0xff, 0x6a, 0x2d, 0x31, 0xfd, 0xf6,
	0x86, 0xd2, 0x0e, 0xc6, 0xa9, 0xf7, 0xbc, 0xc9, 0xd3, 0xeb, 0x5b, 0xb8, 0xef, 0xc5, 0x30, 0xd1,
	0x2a, 0x9f, 0xc2, 0xd9, 0x68, 0xb3, 0xd9, 0x0d, 0xf3, 0x80, 0xf5, 0x1b, 0xdc, 0x75, 0x24, 0xcc,
	0x30, 0xb6, 0xf4, 0xb5, 0xd3, 0xb7, 0x45, 0x84, 0xd4, 0x3a, 0x4d, 0xfc, 0x52, 0x5d, 0x1e, 0x6f,
	0x5a, 0x3f, 0x23, 0x87, 0x7c, 0x22, 0xf7, 0xb3, 0x19, 0xc9, 0xaf, 0x7d, 0x1c, 0xb1, 0x76, 0x70,
	0x1b, 0x75, 0x3d, 0x3b, 0x2c, 0x20, 0xb3, 0xca, 0xce, 0xd1, 0x3e, 0xe1, 0xb8, 0x72, 0x34, 0xf4,
	0x11, 0x5c, 0x3f, 0xfb, 0x35, 0x6a, 0x21, 0x9e, 0x84, 0x35, 0x6a, 0xa1, 0xc3, 0xee, 0x85, 0x92,
	0x3b, 0x4f, 0xab, 0xe4, 0xb0, 0x3d, 0x60, 0x09, 0x6c, 0x11, 0x9f, 0x2b, 0x59, 0x9e, 0xf9, 0x38,
	0xd0, 0xf3, 0x3b, 0xb0, 0x4e, 0x2b, 0x49, 0x06, 0x5f, 0x81, 0x9d, 0x8f, 0xc0, 0xb2, 0x89, 0x35,
	0xff, 0x94, 0x7a, 0x28, 0xfe, 0xe5, 0x7a, 0x76, 0xf5, 0x55, 0xcf, 0xfc, 0x6f, 0xbc, 0xfc, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0xc1, 0xc3, 0x15, 0x26, 0x01, 0x00, 0x00,
}