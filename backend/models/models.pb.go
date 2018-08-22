// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

package models

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

type User_Role int32

const (
	User_ROLE_USER          User_Role = 0
	User_ROLE_CONTENT_ADMIN User_Role = 1
	User_ROLE_ADMIN         User_Role = 2
)

var User_Role_name = map[int32]string{
	0: "ROLE_USER",
	1: "ROLE_CONTENT_ADMIN",
	2: "ROLE_ADMIN",
}
var User_Role_value = map[string]int32{
	"ROLE_USER":          0,
	"ROLE_CONTENT_ADMIN": 1,
	"ROLE_ADMIN":         2,
}

func (x User_Role) String() string {
	return proto.EnumName(User_Role_name, int32(x))
}
func (User_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_models_3a8afb88d6637717, []int{0, 0}
}

type User struct {
	ID                   uint32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt            uint32    `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            uint32    `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FirstName            string    `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string    `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string    `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Password             string    `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Role                 User_Role `protobuf:"varint,8,opt,name=role,proto3,enum=models.User_Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_3a8afb88d6637717, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetCreatedAt() uint32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() uint32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetRole() User_Role {
	if m != nil {
		return m.Role
	}
	return User_ROLE_USER
}

func init() {
	proto.RegisterType((*User)(nil), "models.User")
	proto.RegisterEnum("models.User_Role", User_Role_name, User_Role_value)
}

func init() { proto.RegisterFile("models.proto", fileDescriptor_models_3a8afb88d6637717) }

var fileDescriptor_models_3a8afb88d6637717 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xd0, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x71, 0x13, 0xeb, 0xd9, 0x3c, 0xbc, 0x72, 0x66, 0x90, 0xa0, 0x08, 0xe5, 0x40, 0xe8,
	0xd4, 0x41, 0x67, 0x87, 0xc3, 0x76, 0x28, 0x68, 0x0f, 0xe2, 0xdd, 0x5c, 0xa2, 0x8d, 0x20, 0xb4,
	0xa6, 0x24, 0x39, 0xfc, 0xaf, 0xfc, 0x1b, 0x25, 0x2f, 0xed, 0xf8, 0xfb, 0x7e, 0xa6, 0xf7, 0xe0,
	0x6a, 0x34, 0xbd, 0x1e, 0x5c, 0x39, 0x59, 0xe3, 0x0d, 0x5f, 0xc5, 0xb5, 0xfd, 0xa3, 0x90, 0x1c,
	0x9d, 0xb6, 0x3c, 0x03, 0xda, 0x54, 0x82, 0xe4, 0xa4, 0x58, 0x4b, 0xda, 0x54, 0xfc, 0x1e, 0xe0,
	0xd3, 0x6a, 0xe5, 0x75, 0xdf, 0x29, 0x2f, 0x28, 0x76, 0x36, 0x97, 0x9d, 0x0f, 0x7c, 0x9a, 0xfa,
	0x85, 0xcf, 0x23, 0xcf, 0x25, 0xf2, 0xd7, 0xb7, 0x75, 0xbe, 0xfb, 0x51, 0xa3, 0x16, 0x49, 0x4e,
	0x0a, 0x26, 0x19, 0x96, 0x56, 0x8d, 0x9a, 0xdf, 0x01, 0x1b, 0xd4, 0xa2, 0x17, 0xa8, 0x69, 0x08,
	0x88, 0xb7, 0x90, 0x9e, 0x9c, 0xb6, 0x68, 0xab, 0x68, 0xcb, 0x0e, 0x36, 0x29, 0xe7, 0x7e, 0x8d,
	0xed, 0xc5, 0x65, 0xb4, 0x65, 0xf3, 0x07, 0x48, 0xac, 0x19, 0xb4, 0x48, 0x73, 0x52, 0x64, 0x8f,
	0xd7, 0xe5, 0x7c, 0x6f, 0xb8, 0xae, 0x94, 0x66, 0xd0, 0x12, 0x79, 0xfb, 0x0c, 0x49, 0x58, 0x7c,
	0x0d, 0x4c, 0xee, 0x5f, 0xeb, 0xee, 0xf8, 0x5e, 0xcb, 0xcd, 0x19, 0xbf, 0x01, 0x8e, 0xf3, 0x65,
	0xdf, 0x1e, 0xea, 0xf6, 0xd0, 0xed, 0xaa, 0xb7, 0xa6, 0xdd, 0x10, 0x9e, 0x01, 0x60, 0x8f, 0x9b,
	0x7e, 0xac, 0xf0, 0x7f, 0x4f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x59, 0x25, 0xf8, 0x4f,
	0x01, 0x00, 0x00,
}
