// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package http is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	ServerConfig
	ClientConfig
*/
package http

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

// Config for HTTP proxy server.
type ServerConfig struct {
	Timeout uint32 `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// ClientConfig for HTTP proxy client.
type ClientConfig struct {
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ServerConfig)(nil), "com.v2ray.core.proxy.http.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "com.v2ray.core.proxy.http.ClientConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0xce, 0xcf, 0xd5, 0x2b, 0x33,
	0x2a, 0x4a, 0xac, 0xd4, 0x4b, 0xce, 0x2f, 0x4a, 0x05, 0x89, 0x56, 0x54, 0xea, 0x65, 0x94, 0x94,
	0x14, 0x28, 0x69, 0x70, 0xf1, 0x04, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x39, 0x83, 0x35, 0x08, 0x49,
	0x70, 0xb1, 0x97, 0x64, 0xe6, 0xa6, 0xe6, 0x97, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06,
	0xc1, 0xb8, 0x4a, 0x7c, 0x5c, 0x3c, 0xce, 0x39, 0x99, 0xa9, 0x79, 0x25, 0x10, 0x95, 0x4e, 0x6c,
	0x51, 0x2c, 0x20, 0x13, 0x92, 0xd8, 0xc0, 0x76, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0xea, 0x59, 0xf8, 0x73, 0x00, 0x00, 0x00,
}
