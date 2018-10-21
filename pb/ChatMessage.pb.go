// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ChatMessage.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChatMessage struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver             string   `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ChatMessage_e6944ce291343236, []int{0}
}
func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (dst *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(dst, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ChatMessage) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "proto.ChatMessage")
}

func init() { proto.RegisterFile("ChatMessage.proto", fileDescriptor_ChatMessage_e6944ce291343236) }

var fileDescriptor_ChatMessage_e6944ce291343236 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x74, 0xce, 0x48, 0x2c,
	0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x53, 0x4a, 0xd1, 0x5c, 0xdc, 0x48, 0x72, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29, 0xa9,
	0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x14, 0x17, 0x47, 0x51, 0x6a,
	0x72, 0x6a, 0x66, 0x59, 0x6a, 0x91, 0x04, 0x13, 0x58, 0x06, 0xce, 0x17, 0x92, 0xe0, 0x62, 0xcf,
	0x85, 0x68, 0x97, 0x60, 0x06, 0x4b, 0xc1, 0xb8, 0x49, 0x6c, 0x60, 0x3b, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa3, 0xf4, 0x4d, 0x1d, 0x7f, 0x00, 0x00, 0x00,
}
