// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/spine.proto

package model

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

// RequestType used to sign a node administration request
type SpinePublicKeyAction int32

const (
	SpinePublicKeyAction_AddKey    SpinePublicKeyAction = 0
	SpinePublicKeyAction_RemoveKey SpinePublicKeyAction = 1
)

var SpinePublicKeyAction_name = map[int32]string{
	0: "AddKey",
	1: "RemoveKey",
}

var SpinePublicKeyAction_value = map[string]int32{
	"AddKey":    0,
	"RemoveKey": 1,
}

func (x SpinePublicKeyAction) String() string {
	return proto.EnumName(SpinePublicKeyAction_name, int32(x))
}

func (SpinePublicKeyAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59a5e7c6458ac090, []int{0}
}

// Block represent the block data structure stored in the database
type SpinePublicKey struct {
	NodePublicKey        []byte               `protobuf:"bytes,1,opt,name=NodePublicKey,proto3" json:"NodePublicKey,omitempty"`
	BlockID              int64                `protobuf:"varint,2,opt,name=BlockID,proto3" json:"BlockID,omitempty"`
	PublicKeyAction      SpinePublicKeyAction `protobuf:"varint,3,opt,name=PublicKeyAction,proto3,enum=model.SpinePublicKeyAction" json:"PublicKeyAction,omitempty"`
	Latest               bool                 `protobuf:"varint,4,opt,name=Latest,proto3" json:"Latest,omitempty"`
	Height               uint32               `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SpinePublicKey) Reset()         { *m = SpinePublicKey{} }
func (m *SpinePublicKey) String() string { return proto.CompactTextString(m) }
func (*SpinePublicKey) ProtoMessage()    {}
func (*SpinePublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_59a5e7c6458ac090, []int{0}
}

func (m *SpinePublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpinePublicKey.Unmarshal(m, b)
}
func (m *SpinePublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpinePublicKey.Marshal(b, m, deterministic)
}
func (m *SpinePublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpinePublicKey.Merge(m, src)
}
func (m *SpinePublicKey) XXX_Size() int {
	return xxx_messageInfo_SpinePublicKey.Size(m)
}
func (m *SpinePublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SpinePublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_SpinePublicKey proto.InternalMessageInfo

func (m *SpinePublicKey) GetNodePublicKey() []byte {
	if m != nil {
		return m.NodePublicKey
	}
	return nil
}

func (m *SpinePublicKey) GetBlockID() int64 {
	if m != nil {
		return m.BlockID
	}
	return 0
}

func (m *SpinePublicKey) GetPublicKeyAction() SpinePublicKeyAction {
	if m != nil {
		return m.PublicKeyAction
	}
	return SpinePublicKeyAction_AddKey
}

func (m *SpinePublicKey) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
}

func (m *SpinePublicKey) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("model.SpinePublicKeyAction", SpinePublicKeyAction_name, SpinePublicKeyAction_value)
	proto.RegisterType((*SpinePublicKey)(nil), "model.SpinePublicKey")
}

func init() { proto.RegisterFile("model/spine.proto", fileDescriptor_59a5e7c6458ac090) }

var fileDescriptor_59a5e7c6458ac090 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2f, 0x2e, 0xc8, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x0b, 0x29, 0x1d, 0x67, 0xe4, 0xe2, 0x0b, 0x06, 0x09, 0x07, 0x94, 0x26, 0xe5, 0x64, 0x26, 0x7b,
	0xa7, 0x56, 0x0a, 0xa9, 0x70, 0xf1, 0xfa, 0xe5, 0xa7, 0x20, 0x04, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x78, 0x82, 0x50, 0x05, 0x85, 0x24, 0xb8, 0xd8, 0x9d, 0x72, 0xf2, 0x93, 0xb3, 0x3d, 0x5d, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x60, 0x5c, 0x21, 0x57, 0x2e, 0x7e, 0xb8, 0x32, 0xc7, 0xe4,
	0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x69, 0x3d, 0xb0, 0x9d, 0x7a,
	0xa8, 0xf6, 0x41, 0x94, 0x04, 0xa1, 0xeb, 0x11, 0x12, 0xe3, 0x62, 0xf3, 0x49, 0x2c, 0x49, 0x2d,
	0x2e, 0x91, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0xf2, 0x40, 0xe2, 0x1e, 0xa9, 0x99, 0xe9,
	0x19, 0x25, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x50, 0x9e, 0x96, 0x21, 0x97, 0x08, 0x36,
	0x83, 0x85, 0xb8, 0xb8, 0xd8, 0x1c, 0x53, 0x52, 0xbc, 0x53, 0x2b, 0x05, 0x18, 0x84, 0x78, 0xb9,
	0x38, 0x83, 0x52, 0x73, 0xf3, 0xcb, 0x52, 0x41, 0x5c, 0x46, 0x27, 0xad, 0x28, 0x8d, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xaa, 0xfc, 0xfc, 0xa4, 0x64, 0x08, 0xa9,
	0x9b, 0x9c, 0x5f, 0x94, 0xaa, 0x9f, 0x9c, 0x9f, 0x9b, 0x9b, 0x9f, 0xa7, 0x0f, 0x76, 0x74, 0x12,
	0x1b, 0x38, 0xd8, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x1b, 0x45, 0xd6, 0x4b, 0x01,
	0x00, 0x00,
}