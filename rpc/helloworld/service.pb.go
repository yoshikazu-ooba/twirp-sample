// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/helloworld/service.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	rpc/helloworld/service.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package helloworld

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

type HelloRequest struct {
	Subject string `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type HelloResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "twirp.example.helloworld.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "twirp.example.helloworld.HelloResponse")
}

func init() { proto.RegisterFile("rpc/helloworld/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2a, 0x48, 0xd6,
	0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x29, 0xcf, 0x2c, 0x2a, 0xd0,
	0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x43, 0xa8, 0x53, 0xd2, 0xe0, 0xe2, 0xf1, 0x00,
	0xf1, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x4b, 0x93, 0xb2,
	0x52, 0x93, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x65, 0x2e, 0x5e,
	0xa8, 0xca, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xd4, 0x0a, 0x98,
	0x3a, 0x30, 0xdb, 0x28, 0x8d, 0x8b, 0x0b, 0xac, 0x28, 0x1c, 0x64, 0xb8, 0x50, 0x04, 0x17, 0x2b,
	0x98, 0x27, 0xa4, 0xa6, 0x87, 0xcb, 0x01, 0x7a, 0xc8, 0xb6, 0x4b, 0xa9, 0x13, 0x54, 0x07, 0xb1,
	0xdb, 0x89, 0x27, 0x8a, 0x0b, 0x21, 0x97, 0xc4, 0x06, 0xf6, 0xa5, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x57, 0x3e, 0x7f, 0x2b, 0x05, 0x01, 0x00, 0x00,
}