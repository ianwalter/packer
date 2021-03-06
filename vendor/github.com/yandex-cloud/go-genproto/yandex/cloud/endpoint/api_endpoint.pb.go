// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/endpoint/api_endpoint.proto

package endpoint

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

type ApiEndpoint struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiEndpoint) Reset()         { *m = ApiEndpoint{} }
func (m *ApiEndpoint) String() string { return proto.CompactTextString(m) }
func (*ApiEndpoint) ProtoMessage()    {}
func (*ApiEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f047dfc88bc45bd0, []int{0}
}

func (m *ApiEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiEndpoint.Unmarshal(m, b)
}
func (m *ApiEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiEndpoint.Marshal(b, m, deterministic)
}
func (m *ApiEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiEndpoint.Merge(m, src)
}
func (m *ApiEndpoint) XXX_Size() int {
	return xxx_messageInfo_ApiEndpoint.Size(m)
}
func (m *ApiEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ApiEndpoint proto.InternalMessageInfo

func (m *ApiEndpoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApiEndpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*ApiEndpoint)(nil), "yandex.cloud.endpoint.ApiEndpoint")
}

func init() {
	proto.RegisterFile("yandex/cloud/endpoint/api_endpoint.proto", fileDescriptor_f047dfc88bc45bd0)
}

var fileDescriptor_f047dfc88bc45bd0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xa8, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc,
	0x2b, 0xd1, 0x4f, 0x2c, 0xc8, 0x8c, 0x87, 0x71, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44,
	0x21, 0x2a, 0xf5, 0xc0, 0x2a, 0xf5, 0x60, 0x92, 0x4a, 0xe6, 0x5c, 0xdc, 0x8e, 0x05, 0x99, 0xae,
	0x50, 0xae, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53,
	0x66, 0x8a, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0x13, 0x58,
	0x10, 0xc6, 0x75, 0x72, 0x89, 0x72, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x87, 0x18, 0xae, 0x0b, 0x71, 0x46, 0x7a, 0xbe, 0x6e, 0x7a, 0x6a, 0x1e, 0xd8, 0x5a, 0x7d,
	0xac, 0xee, 0xb3, 0x86, 0x31, 0x92, 0xd8, 0xc0, 0xaa, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x42, 0x6c, 0x3b, 0xb8, 0xc8, 0x00, 0x00, 0x00,
}
