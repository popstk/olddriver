// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend/backend.proto

package backend

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type Item struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Href                 string               `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Baidu                []string             `protobuf:"bytes,4,rep,name=baidu,proto3" json:"baidu,omitempty"`
	Magnet               []string             `protobuf:"bytes,5,rep,name=magnet,proto3" json:"magnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_b81549028379a959, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Item) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *Item) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Item) GetBaidu() []string {
	if m != nil {
		return m.Baidu
	}
	return nil
}

func (m *Item) GetMagnet() []string {
	if m != nil {
		return m.Magnet
	}
	return nil
}

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
	return fileDescriptor_b81549028379a959, []int{1}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
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
	Data                 []*Item  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReply) Reset()         { *m = SearchReply{} }
func (m *SearchReply) String() string { return proto.CompactTextString(m) }
func (*SearchReply) ProtoMessage()    {}
func (*SearchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b81549028379a959, []int{2}
}

func (m *SearchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReply.Unmarshal(m, b)
}
func (m *SearchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReply.Marshal(b, m, deterministic)
}
func (m *SearchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReply.Merge(m, src)
}
func (m *SearchReply) XXX_Size() int {
	return xxx_messageInfo_SearchReply.Size(m)
}
func (m *SearchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReply.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReply proto.InternalMessageInfo

func (m *SearchReply) GetData() []*Item {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "backend.Item")
	proto.RegisterType((*SearchRequest)(nil), "backend.SearchRequest")
	proto.RegisterType((*SearchReply)(nil), "backend.SearchReply")
}

func init() { proto.RegisterFile("backend/backend.proto", fileDescriptor_b81549028379a959) }

var fileDescriptor_b81549028379a959 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0x55, 0xda, 0x34, 0xd5, 0x9f, 0xaa, 0xd2, 0x97, 0xd5, 0x56, 0x56, 0x40, 0x22, 0x64, 0x55,
	0xb1, 0x88, 0xa1, 0xec, 0x90, 0x38, 0x00, 0xdb, 0x94, 0x2d, 0x0b, 0xa7, 0x99, 0xb6, 0x56, 0x93,
	0xd8, 0x24, 0x13, 0x50, 0xb6, 0x1c, 0x80, 0x0d, 0x47, 0xe3, 0x0a, 0x1c, 0x04, 0xc5, 0x49, 0x90,
	0x10, 0x2b, 0xcf, 0x7b, 0xf3, 0xe6, 0xf9, 0xcd, 0xc0, 0x32, 0x91, 0xbb, 0x13, 0x16, 0xa9, 0xe8,
	0xdf, 0xc8, 0x94, 0x9a, 0x34, 0x9b, 0xf6, 0xd0, 0xbf, 0x38, 0x68, 0x7d, 0xc8, 0x50, 0x58, 0x3a,
	0xa9, 0xf7, 0x82, 0x54, 0x8e, 0x15, 0xc9, 0xdc, 0x74, 0x4a, 0xff, 0xbc, 0x17, 0x48, 0xa3, 0x84,
	0x2c, 0x0a, 0x4d, 0x92, 0x94, 0x2e, 0xaa, 0xae, 0x1b, 0xbe, 0x3b, 0xe0, 0x3e, 0x10, 0xe6, 0x6c,
	0x01, 0x13, 0x52, 0x94, 0x21, 0x77, 0x02, 0x67, 0xfd, 0x2f, 0xee, 0x00, 0x63, 0xe0, 0x1e, 0x4b,
	0xdc, 0xf3, 0x91, 0x25, 0x6d, 0xcd, 0x22, 0x70, 0xdb, 0x3f, 0xf8, 0x38, 0x70, 0xd6, 0xb3, 0x8d,
	0x1f, 0x75, 0xfe, 0xd1, 0x10, 0x20, 0x7a, 0x1c, 0x02, 0xc4, 0x56, 0xd7, 0x3a, 0x27, 0x52, 0xa5,
	0x35, 0x77, 0x83, 0x71, 0xeb, 0x6c, 0x01, 0x5b, 0x81, 0x97, 0xcb, 0x43, 0x81, 0xc4, 0x27, 0x96,
	0xee, 0x51, 0x78, 0x0f, 0xf3, 0x2d, 0xca, 0x72, 0x77, 0x8c, 0xf1, 0xb9, 0xc6, 0x8a, 0xda, 0x08,
	0xd4, 0x98, 0x21, 0x97, 0xad, 0x19, 0x87, 0xe9, 0x09, 0x9b, 0x57, 0x5d, 0xa6, 0x7d, 0xb2, 0x01,
	0x86, 0xd7, 0x30, 0x1b, 0xc6, 0x4d, 0xd6, 0xb0, 0x4b, 0x70, 0x53, 0x49, 0x92, 0x3b, 0xc1, 0x78,
	0x3d, 0xdb, 0xcc, 0xa3, 0xe1, 0x88, 0xed, 0xca, 0xb1, 0x6d, 0x6d, 0x9e, 0xc0, 0xdb, 0x1a, 0x95,
	0x62, 0xc9, 0xb6, 0xe0, 0x75, 0xb3, 0x6c, 0xf5, 0x23, 0xfc, 0x95, 0xc5, 0x5f, 0xfc, 0xe1, 0x4d,
	0xd6, 0x84, 0x67, 0x6f, 0x9f, 0x5f, 0x1f, 0xa3, 0x65, 0xf8, 0x5f, 0xbc, 0xdc, 0x88, 0xca, 0x7a,
	0x09, 0x45, 0x98, 0x57, 0x77, 0xce, 0x55, 0xe2, 0xd9, 0xbb, 0xdc, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x36, 0x0d, 0x0f, 0xc8, 0x01, 0x00, 0x00,
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
