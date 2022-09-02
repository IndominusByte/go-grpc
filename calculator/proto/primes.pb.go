// Code generated by protoc-gen-go. DO NOT EDIT.
// source: primes.proto

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

type PrimeRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeRequest) Reset()         { *m = PrimeRequest{} }
func (m *PrimeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeRequest) ProtoMessage()    {}
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c884635c897ca360, []int{0}
}

func (m *PrimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeRequest.Unmarshal(m, b)
}
func (m *PrimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeRequest.Merge(m, src)
}
func (m *PrimeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeRequest.Size(m)
}
func (m *PrimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeRequest proto.InternalMessageInfo

func (m *PrimeRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PrimeResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeResponse) Reset()         { *m = PrimeResponse{} }
func (m *PrimeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeResponse) ProtoMessage()    {}
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c884635c897ca360, []int{1}
}

func (m *PrimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeResponse.Unmarshal(m, b)
}
func (m *PrimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeResponse.Merge(m, src)
}
func (m *PrimeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeResponse.Size(m)
}
func (m *PrimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeResponse proto.InternalMessageInfo

func (m *PrimeResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*PrimeRequest)(nil), "calculator.PrimeRequest")
	proto.RegisterType((*PrimeResponse)(nil), "calculator.PrimeResponse")
}

func init() {
	proto.RegisterFile("primes.proto", fileDescriptor_c884635c897ca360)
}

var fileDescriptor_c884635c897ca360 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0xcc,
	0x4d, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4e, 0xcc, 0x49, 0x2e, 0xcd,
	0x49, 0x2c, 0xc9, 0x2f, 0x52, 0x52, 0xe3, 0xe2, 0x09, 0x00, 0xc9, 0x05, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x89, 0x71, 0xb1, 0xe5, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0x49, 0x30, 0x2a, 0x30,
	0x6a, 0x30, 0x07, 0x41, 0x79, 0x4a, 0xea, 0x5c, 0xbc, 0x50, 0x75, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x20, 0x85, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x30, 0x85, 0x10, 0x9e, 0x93, 0x71, 0x94,
	0x61, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x67, 0x5e, 0x4a, 0x7e,
	0x6e, 0x66, 0x5e, 0x69, 0xb1, 0x53, 0x65, 0x49, 0xaa, 0x7e, 0x7a, 0xbe, 0x6e, 0x7a, 0x51, 0x41,
	0xb2, 0x3e, 0xc2, 0x7e, 0x7d, 0xb0, 0x8b, 0x92, 0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x38, 0xee, 0x54, 0xdd, 0xa8, 0x00, 0x00, 0x00,
}