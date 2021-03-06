// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello_world.proto

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_2153d3bc7c0685ef, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	// 用户ID
	Id                   int64      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserReply            *UserReply `protobuf:"bytes,2,opt,name=user_reply,json=userReply" json:"user_reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_2153d3bc7c0685ef, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HelloResponse) GetUserReply() *UserReply {
	if m != nil {
		return m.UserReply
	}
	return nil
}

type UserReply struct {
	// 用户ID
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// 用户名
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// 证件类型
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	// 真实姓名
	Truename string `protobuf:"bytes,4,opt,name=truename" json:"truename,omitempty"`
	// 邮箱
	Email string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	// 手机号
	Mobile string `protobuf:"bytes,6,opt,name=mobile" json:"mobile,omitempty"`
	// 是否开启支付验证
	PayPassword string `protobuf:"bytes,7,opt,name=payPassword" json:"payPassword,omitempty"`
	// 认证等级
	AuthLevel int32 `protobuf:"varint,8,opt,name=authLevel" json:"authLevel,omitempty"`
	// openId
	OpenId               string   `protobuf:"bytes,10,opt,name=openId" json:"openId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReply) Reset()         { *m = UserReply{} }
func (m *UserReply) String() string { return proto.CompactTextString(m) }
func (*UserReply) ProtoMessage()    {}
func (*UserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_2153d3bc7c0685ef, []int{2}
}
func (m *UserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReply.Unmarshal(m, b)
}
func (m *UserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReply.Marshal(b, m, deterministic)
}
func (dst *UserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReply.Merge(dst, src)
}
func (m *UserReply) XXX_Size() int {
	return xxx_messageInfo_UserReply.Size(m)
}
func (m *UserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserReply proto.InternalMessageInfo

func (m *UserReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserReply) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UserReply) GetTruename() string {
	if m != nil {
		return m.Truename
	}
	return ""
}

func (m *UserReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserReply) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserReply) GetPayPassword() string {
	if m != nil {
		return m.PayPassword
	}
	return ""
}

func (m *UserReply) GetAuthLevel() int32 {
	if m != nil {
		return m.AuthLevel
	}
	return 0
}

func (m *UserReply) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "proto.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "proto.HelloResponse")
	proto.RegisterType((*UserReply)(nil), "proto.UserReply")
}

func init() {
	proto.RegisterFile("proto/hello_world.proto", fileDescriptor_hello_world_2153d3bc7c0685ef)
}

var fileDescriptor_hello_world_2153d3bc7c0685ef = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0xda, 0xc4, 0x66, 0xaa, 0x22, 0x6b, 0xd1, 0xa5, 0x78, 0x08, 0x39, 0xe5, 0xd4,
	0x42, 0xf5, 0x17, 0x78, 0x52, 0xf0, 0x50, 0x16, 0xc4, 0x63, 0x49, 0xc9, 0x40, 0x03, 0x9b, 0xec,
	0xba, 0x9b, 0x58, 0xf2, 0xa3, 0xfd, 0x0f, 0xb2, 0xb3, 0x69, 0xad, 0xf4, 0x94, 0x79, 0xdf, 0x7b,
	0x4c, 0x1e, 0xb3, 0xf0, 0xa0, 0x8d, 0x6a, 0xd5, 0x72, 0x87, 0x52, 0xaa, 0xcd, 0x5e, 0x19, 0x59,
	0x2e, 0x88, 0xb0, 0x88, 0x3e, 0x59, 0x06, 0x57, 0xaf, 0xce, 0x13, 0xf8, 0xd5, 0xa1, 0x6d, 0x19,
	0x83, 0x71, 0x53, 0xd4, 0xc8, 0x83, 0x34, 0xc8, 0x13, 0x41, 0x73, 0xb6, 0x86, 0xeb, 0x21, 0x63,
	0xb5, 0x6a, 0x2c, 0xb2, 0x1b, 0x08, 0xab, 0x92, 0x22, 0x23, 0x11, 0x56, 0x25, 0x5b, 0x02, 0x74,
	0x16, 0xcd, 0xc6, 0xa0, 0x96, 0x3d, 0x0f, 0xd3, 0x20, 0x9f, 0xae, 0x6e, 0xfd, 0x7f, 0x16, 0x1f,
	0x16, 0x8d, 0x70, 0x5c, 0x24, 0xdd, 0x61, 0xcc, 0x7e, 0x02, 0x48, 0x8e, 0xc6, 0xd9, 0xba, 0x39,
	0x4c, 0x5c, 0x94, 0x7a, 0x84, 0xd4, 0xe3, 0xa8, 0x5d, 0xbf, 0xb6, 0xd7, 0xc8, 0x47, 0xbe, 0x9f,
	0x9b, 0x5d, 0xbe, 0x35, 0x1d, 0x52, 0x7e, 0xec, 0xf3, 0x07, 0xcd, 0x66, 0x10, 0x61, 0x5d, 0x54,
	0x92, 0x47, 0x64, 0x78, 0xc1, 0xee, 0x21, 0xae, 0xd5, 0xb6, 0x92, 0xc8, 0x63, 0xc2, 0x83, 0x62,
	0x29, 0x4c, 0x75, 0xd1, 0xaf, 0x0b, 0x6b, 0xf7, 0xca, 0x94, 0xfc, 0x92, 0xcc, 0x53, 0xc4, 0x1e,
	0x21, 0x29, 0xba, 0x76, 0xf7, 0x8e, 0xdf, 0x28, 0xf9, 0x24, 0x0d, 0xf2, 0x48, 0xfc, 0x01, 0xb7,
	0x57, 0x69, 0x6c, 0xde, 0x4a, 0x0e, 0x7e, 0xaf, 0x57, 0xab, 0x17, 0x00, 0xba, 0xe0, 0xa7, 0x7b,
	0x00, 0xf6, 0x0c, 0x11, 0x29, 0x76, 0x37, 0xdc, 0xe8, 0xf4, 0x05, 0xe6, 0xb3, 0xff, 0xd0, 0x9f,
	0x3c, 0xbb, 0xd8, 0xc6, 0x84, 0x9f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x72, 0x3f, 0xdc, 0x52,
	0xd2, 0x01, 0x00, 0x00,
}
