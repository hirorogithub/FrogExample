// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package echoTest is a generated protocol buffer package.

It is generated from these files:
	echo.proto
	echo_service.proto

It has these top-level messages:
	ProtoEchoRequest
	ProtoEchoResponse
*/
package echoTest

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

type ProtoEchoRequest struct {
	Text             *string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoEchoRequest) Reset()                    { *m = ProtoEchoRequest{} }
func (m *ProtoEchoRequest) String() string            { return proto.CompactTextString(m) }
func (*ProtoEchoRequest) ProtoMessage()               {}
func (*ProtoEchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoEchoRequest) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type ProtoEchoResponse struct {
	Text             *string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoEchoResponse) Reset()                    { *m = ProtoEchoResponse{} }
func (m *ProtoEchoResponse) String() string            { return proto.CompactTextString(m) }
func (*ProtoEchoResponse) ProtoMessage()               {}
func (*ProtoEchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProtoEchoResponse) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoEchoRequest)(nil), "echoTest.ProtoEchoRequest")
	proto.RegisterType((*ProtoEchoResponse)(nil), "echoTest.ProtoEchoResponse")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xb1, 0x43, 0x52, 0x8b, 0x4b, 0x94, 0xd4,
	0xb8, 0x04, 0x02, 0x40, 0x42, 0xae, 0xc9, 0x19, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60,
	0xb6, 0x92, 0x3a, 0x97, 0x20, 0x92, 0xba, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x6c, 0x0a, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x13, 0x17, 0xa8, 0x67, 0x00, 0x00, 0x00,
}
