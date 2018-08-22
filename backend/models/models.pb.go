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

type User_AuthRole int32

const (
	User_ROLE_USER          User_AuthRole = 0
	User_ROLE_CONTENT_ADMIN User_AuthRole = 1
	User_ROLE_ADMIN         User_AuthRole = 2
)

var User_AuthRole_name = map[int32]string{
	0: "ROLE_USER",
	1: "ROLE_CONTENT_ADMIN",
	2: "ROLE_ADMIN",
}
var User_AuthRole_value = map[string]int32{
	"ROLE_USER":          0,
	"ROLE_CONTENT_ADMIN": 1,
	"ROLE_ADMIN":         2,
}

func (x User_AuthRole) String() string {
	return proto.EnumName(User_AuthRole_name, int32(x))
}
func (User_AuthRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_models_21b25800083d68d8, []int{0, 0}
}

type User struct {
	ID                   uint32        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt            uint32        `protobuf:"varint,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            uint32        `protobuf:"varint,3,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	FirstName            string        `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string        `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Username             string        `protobuf:"bytes,6,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string        `protobuf:"bytes,7,opt,name=Password,proto3" json:"Password,omitempty"`
	Role                 User_AuthRole `protobuf:"varint,8,opt,name=Role,proto3,enum=models.User_AuthRole" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_21b25800083d68d8, []int{0}
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

func (m *User) GetRole() User_AuthRole {
	if m != nil {
		return m.Role
	}
	return User_ROLE_USER
}

func init() {
	proto.RegisterType((*User)(nil), "models.User")
	proto.RegisterEnum("models.User_AuthRole", User_AuthRole_name, User_AuthRole_value)
}

func init() { proto.RegisterFile("models.proto", fileDescriptor_models_21b25800083d68d8) }

var fileDescriptor_models_21b25800083d68d8 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x71, 0x13, 0x6b, 0x6d, 0xff, 0xb8, 0x32, 0xfe, 0xa0, 0x04, 0xf1, 0x50, 0x76, 0xaa,
	0x97, 0x1e, 0xf4, 0x09, 0xca, 0x5a, 0xa1, 0x30, 0x3b, 0x89, 0xeb, 0x79, 0x54, 0x12, 0x50, 0xd8,
	0xcc, 0x48, 0x22, 0xbe, 0x97, 0x4f, 0x28, 0x49, 0xda, 0xf5, 0xf8, 0xfb, 0x7e, 0x5a, 0x48, 0x02,
	0x37, 0x47, 0x25, 0xe4, 0xc1, 0x94, 0x27, 0xad, 0xac, 0xc2, 0x38, 0xac, 0xd5, 0x1f, 0x85, 0xa8,
	0x37, 0x52, 0x63, 0x06, 0xb4, 0xad, 0x19, 0xc9, 0x49, 0xb1, 0xe0, 0xb4, 0xad, 0xf1, 0x01, 0xd2,
	0xb5, 0x96, 0x83, 0x95, 0xa2, 0xb2, 0x8c, 0xfa, 0x3c, 0x07, 0xa7, 0xfd, 0x49, 0x8c, 0x7a, 0x19,
	0xf4, 0x1c, 0x9c, 0xbe, 0x7c, 0x69, 0x63, 0xbb, 0xe1, 0x28, 0x59, 0x94, 0x93, 0x22, 0xe5, 0x73,
	0xc0, 0x7b, 0x48, 0x36, 0xc3, 0x88, 0x57, 0x1e, 0xcf, 0xdb, 0x99, 0x3b, 0xcd, 0xb7, 0xb3, 0x38,
	0xd8, 0xb4, 0x9d, 0xbd, 0x0d, 0xc6, 0xfc, 0x2a, 0x2d, 0xd8, 0x75, 0xb0, 0x69, 0xe3, 0x23, 0x44,
	0x5c, 0x1d, 0x24, 0x4b, 0x72, 0x52, 0x64, 0x4f, 0xb7, 0xe5, 0x78, 0x57, 0xf7, 0x6f, 0x59, 0xfd,
	0xd8, 0x4f, 0x87, 0xdc, 0x7f, 0xb2, 0xaa, 0x20, 0x99, 0x0a, 0x2e, 0x20, 0xe5, 0xdb, 0x4d, 0xb3,
	0xef, 0xdf, 0x1b, 0xbe, 0xbc, 0xc0, 0x3b, 0x40, 0x3f, 0xd7, 0xdb, 0x6e, 0xd7, 0x74, 0xbb, 0x7d,
	0x55, 0xbf, 0xb6, 0xdd, 0x92, 0x60, 0x06, 0xe0, 0x7b, 0xd8, 0xf4, 0x23, 0xf6, 0x6f, 0xf8, 0xfc,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xbc, 0xc0, 0x62, 0x53, 0x01, 0x00, 0x00,
}
