// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/main/proto/org/kai/example/helloworld.proto

package example

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GreetingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingRequest) Reset()         { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()    {}
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36eb56c8231d77dc, []int{0}
}

func (m *GreetingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingRequest.Unmarshal(m, b)
}
func (m *GreetingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingRequest.Marshal(b, m, deterministic)
}
func (m *GreetingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingRequest.Merge(m, src)
}
func (m *GreetingRequest) XXX_Size() int {
	return xxx_messageInfo_GreetingRequest.Size(m)
}
func (m *GreetingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingRequest proto.InternalMessageInfo

func (m *GreetingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetingResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingResponse) Reset()         { *m = GreetingResponse{} }
func (m *GreetingResponse) String() string { return proto.CompactTextString(m) }
func (*GreetingResponse) ProtoMessage()    {}
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36eb56c8231d77dc, []int{1}
}

func (m *GreetingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingResponse.Unmarshal(m, b)
}
func (m *GreetingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingResponse.Marshal(b, m, deterministic)
}
func (m *GreetingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingResponse.Merge(m, src)
}
func (m *GreetingResponse) XXX_Size() int {
	return xxx_messageInfo_GreetingResponse.Size(m)
}
func (m *GreetingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingResponse proto.InternalMessageInfo

func (m *GreetingResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetingRequest)(nil), "example.GreetingRequest")
	proto.RegisterType((*GreetingResponse)(nil), "example.GreetingResponse")
}

func init() {
	proto.RegisterFile("src/main/proto/org/kai/example/helloworld.proto", fileDescriptor_36eb56c8231d77dc)
}

var fileDescriptor_36eb56c8231d77dc = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x2e, 0x4a, 0xd6,
	0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x2f, 0x4a, 0xd7, 0xcf,
	0x4e, 0xcc, 0xd4, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0xcf, 0x48, 0xcd, 0xc9, 0xc9,
	0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x03, 0x2b, 0x10, 0x62, 0x87, 0xca, 0x48, 0xc9, 0xa4, 0xe7,
	0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96,
	0x64, 0xe6, 0xe7, 0x15, 0x43, 0x94, 0x29, 0xa9, 0x72, 0xf1, 0xbb, 0x17, 0xa5, 0xa6, 0x96, 0x64,
	0xe6, 0xa5, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6,
	0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x7a, 0x5c, 0x02, 0x08, 0x65,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0xe9, 0x50, 0x31, 0xa8, 0x5a, 0x38,
	0xdf, 0x28, 0x97, 0x4b, 0xd0, 0x03, 0xe4, 0xa2, 0x70, 0x90, 0x8b, 0x82, 0x53, 0x8b, 0xca, 0x32,
	0x93, 0x53, 0x85, 0x22, 0x10, 0x1a, 0x84, 0x24, 0xf4, 0xa0, 0xee, 0xd3, 0x43, 0xb3, 0x5e, 0x4a,
	0x12, 0x8b, 0x0c, 0xc4, 0x46, 0x25, 0xf1, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x09, 0x2a, 0xf1, 0xe8,
	0x97, 0x19, 0xea, 0xc3, 0x8c, 0xb2, 0x62, 0xd4, 0x72, 0xd2, 0xe6, 0xe2, 0xcf, 0x2f, 0x4a, 0xd7,
	0xcb, 0x4e, 0xcc, 0x84, 0x69, 0x76, 0x12, 0xc3, 0xb0, 0x3f, 0x00, 0xe4, 0xe1, 0x00, 0xc6, 0x24,
	0x36, 0xb0, 0xcf, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x3a, 0x82, 0x41, 0x53, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error)
}

type helloWorldServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldServiceClient(cc *grpc.ClientConn) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, "/example.HelloWorldService/greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
type HelloWorldServiceServer interface {
	Greeting(context.Context, *GreetingRequest) (*GreetingResponse, error)
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.HelloWorldService/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).Greeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "greeting",
			Handler:    _HelloWorldService_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/main/proto/org/kai/example/helloworld.proto",
}
