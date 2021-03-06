// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package proto

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

type RespType int32

const (
	RespType_NONE    RespType = 0
	RespType_ASCEND  RespType = 1
	RespType_DESCEND RespType = 2
)

var RespType_name = map[int32]string{
	0: "NONE",
	1: "ASCEND",
	2: "DESCEND",
}
var RespType_value = map[string]int32{
	"NONE":    0,
	"ASCEND":  1,
	"DESCEND": 2,
}

func (x RespType) String() string {
	return proto.EnumName(RespType_name, int32(x))
}
func (RespType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_0e7d481af33caa89, []int{0}
}

type SayParam struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayParam) Reset()         { *m = SayParam{} }
func (m *SayParam) String() string { return proto.CompactTextString(m) }
func (*SayParam) ProtoMessage()    {}
func (*SayParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_0e7d481af33caa89, []int{0}
}
func (m *SayParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayParam.Unmarshal(m, b)
}
func (m *SayParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayParam.Marshal(b, m, deterministic)
}
func (dst *SayParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayParam.Merge(dst, src)
}
func (m *SayParam) XXX_Size() int {
	return xxx_messageInfo_SayParam.Size(m)
}
func (m *SayParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SayParam.DiscardUnknown(m)
}

var xxx_messageInfo_SayParam proto.InternalMessageInfo

func (m *SayParam) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Pair struct {
	Key                  int32    `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	Values               string   `protobuf:"bytes,2,opt,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_0e7d481af33caa89, []int{1}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (dst *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(dst, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Pair) GetValues() string {
	if m != nil {
		return m.Values
	}
	return ""
}

type SayResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	// 数组
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	// map
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type                 RespType         `protobuf:"varint,4,opt,name=type,enum=proto.RespType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SayResponse) Reset()         { *m = SayResponse{} }
func (m *SayResponse) String() string { return proto.CompactTextString(m) }
func (*SayResponse) ProtoMessage()    {}
func (*SayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_0e7d481af33caa89, []int{2}
}
func (m *SayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayResponse.Unmarshal(m, b)
}
func (m *SayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayResponse.Marshal(b, m, deterministic)
}
func (dst *SayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayResponse.Merge(dst, src)
}
func (m *SayResponse) XXX_Size() int {
	return xxx_messageInfo_SayResponse.Size(m)
}
func (m *SayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayResponse proto.InternalMessageInfo

func (m *SayResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SayResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SayResponse) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SayResponse) GetType() RespType {
	if m != nil {
		return m.Type
	}
	return RespType_NONE
}

func init() {
	proto.RegisterType((*SayParam)(nil), "proto.SayParam")
	proto.RegisterType((*Pair)(nil), "proto.Pair")
	proto.RegisterType((*SayResponse)(nil), "proto.SayResponse")
	proto.RegisterMapType((map[string]*Pair)(nil), "proto.SayResponse.HeaderEntry")
	proto.RegisterEnum("proto.RespType", RespType_name, RespType_value)
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_common_0e7d481af33caa89) }

var fileDescriptor_common_0e7d481af33caa89 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xdd, 0x4a, 0x84, 0x40,
	0x14, 0xc7, 0x77, 0xfc, 0xca, 0x3d, 0x03, 0xad, 0x9c, 0x8b, 0x90, 0x25, 0xc2, 0xec, 0x46, 0x82,
	0x2c, 0x8c, 0x22, 0xba, 0x8b, 0xd6, 0xd8, 0x2b, 0x5b, 0xc6, 0x5e, 0x60, 0xaa, 0xa1, 0x22, 0x75,
	0x44, 0x2d, 0x98, 0x77, 0xed, 0x61, 0xc2, 0xd1, 0x58, 0x63, 0xaf, 0x3c, 0xc7, 0xff, 0x07, 0xe7,
	0x37, 0x80, 0x75, 0x23, 0x3b, 0x79, 0xfe, 0x22, 0xcb, 0x52, 0x56, 0xb1, 0x5e, 0xd0, 0xd6, 0x9f,
	0xf0, 0x10, 0xdc, 0x9c, 0xab, 0x0d, 0x6f, 0x78, 0x89, 0x1e, 0x98, 0x65, 0xfb, 0xe6, 0x93, 0x80,
	0x44, 0x73, 0xd6, 0x8f, 0xe1, 0x05, 0x58, 0x1b, 0xfe, 0xd1, 0xf4, 0xca, 0xa7, 0x50, 0x5a, 0xb1,
	0x59, 0x3f, 0xe2, 0x01, 0x38, 0xdf, 0xbc, 0xf8, 0x12, 0xad, 0x6f, 0x68, 0xfb, 0xb8, 0x85, 0x3f,
	0x04, 0x68, 0xce, 0x15, 0x13, 0x6d, 0x2d, 0xab, 0x56, 0xec, 0x76, 0xfe, 0x4b, 0x9a, 0xdb, 0x24,
	0x5e, 0x83, 0xf3, 0x2e, 0xf8, 0xab, 0x68, 0x7c, 0x33, 0x30, 0x23, 0x9a, 0x1c, 0x0d, 0x87, 0xc6,
	0x93, 0xb6, 0x78, 0xad, 0x0d, 0x69, 0xd5, 0x35, 0x8a, 0x8d, 0x6e, 0x3c, 0x01, 0xab, 0x53, 0xb5,
	0xf0, 0xad, 0x80, 0x44, 0xfb, 0xc9, 0x62, 0x4c, 0xf5, 0x91, 0x27, 0x55, 0x0b, 0xa6, 0xc5, 0xe5,
	0x03, 0xd0, 0x49, 0x76, 0xca, 0x33, 0x1f, 0x78, 0x8e, 0xc1, 0xd6, 0x77, 0x68, 0x1c, 0x9a, 0xd0,
	0xb1, 0xa6, 0xa7, 0x67, 0x83, 0x72, 0x6b, 0xdc, 0x90, 0xd3, 0x33, 0x70, 0xff, 0x9a, 0xd1, 0x05,
	0x2b, 0x7b, 0xcc, 0x52, 0x6f, 0x86, 0x00, 0xce, 0x5d, 0x7e, 0x9f, 0x66, 0x2b, 0x8f, 0x20, 0x85,
	0xbd, 0x55, 0x3a, 0x2c, 0x46, 0x72, 0x05, 0x66, 0xce, 0x15, 0xc6, 0x60, 0xaf, 0x45, 0x51, 0x48,
	0x5c, 0x6c, 0x99, 0xf4, 0x93, 0x2f, 0x71, 0x17, 0x32, 0x9c, 0x3d, 0x3b, 0xfa, 0xe7, 0xe5, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x72, 0x08, 0x88, 0xb8, 0x01, 0x00, 0x00,
}
