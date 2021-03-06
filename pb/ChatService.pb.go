// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ChatService.proto

package pb

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (Chat_RouteChatClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (Chat_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/proto.Chat/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRouteChatClient{stream}
	return x, nil
}

type Chat_RouteChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatRouteChatClient struct {
	grpc.ClientStream
}

func (x *chatRouteChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRouteChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	RouteChat(Chat_RouteChatServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).RouteChat(&chatRouteChatServer{stream})
}

type Chat_RouteChatServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatRouteChatServer struct {
	grpc.ServerStream
}

func (x *chatRouteChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRouteChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteChat",
			Handler:       _Chat_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ChatService.proto",
}

func init() { proto.RegisterFile("ChatService.proto", fileDescriptor_ChatService_fe443586a5fad162) }

var fileDescriptor_ChatService_fe443586a5fad162 = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x74, 0xce, 0x48, 0x2c,
	0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x53, 0x52, 0x60, 0x19, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0xa8, 0x8c, 0x91, 0x3d, 0x17, 0x0b,
	0x48, 0x50, 0xc8, 0x9c, 0x8b, 0x33, 0x28, 0xbf, 0xb4, 0x24, 0x15, 0xcc, 0x11, 0x82, 0x48, 0xea,
	0x21, 0x29, 0x97, 0xc2, 0x22, 0xa6, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x16, 0x36, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x91, 0x8f, 0xe7, 0xf9, 0x76, 0x00, 0x00, 0x00,
}
