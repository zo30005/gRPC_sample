// Code generated by protoc-gen-go. DO NOT EDIT.
// source: send.proto

package send

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// The request message
type SendRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e4a6fa8e6deebe, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The response message
type SendReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendReply) Reset()         { *m = SendReply{} }
func (m *SendReply) String() string { return proto.CompactTextString(m) }
func (*SendReply) ProtoMessage()    {}
func (*SendReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e4a6fa8e6deebe, []int{1}
}

func (m *SendReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendReply.Unmarshal(m, b)
}
func (m *SendReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendReply.Marshal(b, m, deterministic)
}
func (m *SendReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendReply.Merge(m, src)
}
func (m *SendReply) XXX_Size() int {
	return xxx_messageInfo_SendReply.Size(m)
}
func (m *SendReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendReply proto.InternalMessageInfo

func (m *SendReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "send.SendRequest")
	proto.RegisterType((*SendReply)(nil), "send.SendReply")
}

func init() { proto.RegisterFile("send.proto", fileDescriptor_27e4a6fa8e6deebe) }

var fileDescriptor_27e4a6fa8e6deebe = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4e, 0xcd, 0x4b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xd4, 0xb9, 0xb8, 0x83, 0x53,
	0xf3, 0x52, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b,
	0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x55, 0x2e,
	0x4e, 0x88, 0xc2, 0x82, 0x9c, 0x4a, 0xdc, 0xca, 0x8c, 0x6c, 0xb9, 0xd8, 0x40, 0xca, 0x52, 0x8b,
	0x84, 0x8c, 0x21, 0x26, 0x87, 0xe4, 0x07, 0xe7, 0x24, 0x26, 0x67, 0x0b, 0x09, 0xea, 0x81, 0xed,
	0x46, 0xb2, 0x4c, 0x8a, 0x1f, 0x59, 0xa8, 0x20, 0xa7, 0x52, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x36,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x7e, 0xef, 0x95, 0xa9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SenderClient is the client API for Sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SenderClient interface {
	SendToSlack(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error)
}

type senderClient struct {
	cc *grpc.ClientConn
}

func NewSenderClient(cc *grpc.ClientConn) SenderClient {
	return &senderClient{cc}
}

func (c *senderClient) SendToSlack(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error) {
	out := new(SendReply)
	err := c.cc.Invoke(ctx, "/send.Sender/SendToSlack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SenderServer is the server API for Sender service.
type SenderServer interface {
	SendToSlack(context.Context, *SendRequest) (*SendReply, error)
}

func RegisterSenderServer(s *grpc.Server, srv SenderServer) {
	s.RegisterService(&_Sender_serviceDesc, srv)
}

func _Sender_SendToSlack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).SendToSlack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/send.Sender/SendToSlack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).SendToSlack(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "send.Sender",
	HandlerType: (*SenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToSlack",
			Handler:    _Sender_SendToSlack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "send.proto",
}