// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GreetRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
}

func init() {
	proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579)
}

var fileDescriptor_32c0044392f32579 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0xb0, 0x41, 0xdf, 0x54, 0x30, 0x8a, 0xc8, 0x40, 0x90, 0x5d, 0xdc, 0xc1, 0xb5,
	0x63, 0x5e, 0xbc, 0xe8, 0x61, 0x28, 0x22, 0x4c, 0x0f, 0x53, 0x10, 0xbc, 0x48, 0xd7, 0x3d, 0xb3,
	0x40, 0x93, 0x57, 0x93, 0x74, 0xd0, 0x4f, 0xed, 0x57, 0x90, 0xbe, 0x06, 0x51, 0x6f, 0x3d, 0x85,
	0xff, 0x8f, 0xf7, 0x7b, 0x2f, 0xc9, 0x83, 0x81, 0xb4, 0x88, 0x3e, 0x29, 0x2d, 0x79, 0x12, 0x3d,
	0x0e, 0xc3, 0x41, 0x41, 0x52, 0x99, 0x96, 0x8d, 0x26, 0xb0, 0x7b, 0xdf, 0xd0, 0x25, 0x7e, 0x56,
	0xe8, 0xbc, 0x38, 0x05, 0xf8, 0x50, 0xd6, 0xf9, 0x77, 0x93, 0x69, 0x3c, 0x89, 0xce, 0xa2, 0x71,
	0xbc, 0x8c, 0x99, 0x3c, 0x65, 0x1a, 0x47, 0xe7, 0xb0, 0x17, 0xca, 0x5d, 0x49, 0xc6, 0xa1, 0x38,
	0x86, 0xbe, 0x45, 0x57, 0x15, 0x3e, 0xd4, 0x86, 0x34, 0xfb, 0xda, 0x09, 0x8d, 0x9f, 0xd1, 0x6e,
	0x55, 0x8e, 0x62, 0x06, 0x3d, 0xce, 0xe2, 0x30, 0x69, 0xef, 0xf4, 0x7b, 0xec, 0xf0, 0xe8, 0x2f,
	0x0c, 0xcd, 0xaf, 0x61, 0x9f, 0xc1, 0x63, 0x66, 0xea, 0x17, 0xa5, 0xd1, 0x75, 0x90, 0xa7, 0x91,
	0xb8, 0x82, 0x78, 0x41, 0x46, 0x76, 0x1d, 0x3b, 0x8e, 0xc4, 0x4d, 0x78, 0xe6, 0xdd, 0x16, 0x6d,
	0x4d, 0x06, 0x3b, 0xd9, 0xd3, 0xc6, 0x3f, 0x60, 0xf8, 0xaa, 0xfc, 0xe6, 0x16, 0xb3, 0x75, 0xa1,
	0x3a, 0xf5, 0x68, 0x3e, 0x6b, 0xd1, 0x2c, 0xe9, 0xc7, 0xe1, 0xf4, 0xdf, 0x09, 0xb0, 0x75, 0xe6,
	0xc9, 0xdb, 0x85, 0x54, 0x7e, 0x53, 0xad, 0x92, 0x9c, 0x74, 0xfa, 0x60, 0xd6, 0xa4, 0x95, 0xa9,
	0xdc, 0xbc, 0xf6, 0x98, 0x4a, 0x9a, 0x48, 0x5b, 0xe6, 0x29, 0x7b, 0x29, 0x6f, 0x7e, 0xd5, 0xe7,
	0xe3, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x47, 0x04, 0x31, 0x23, 0x02, 0x00, 0x00,
}
