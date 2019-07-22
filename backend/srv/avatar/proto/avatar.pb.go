// Code generated by protoc-gen-go. DO NOT EDIT.
// source: avatar.proto

package avatar

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

type AvatarCreateResponse_Status int32

const (
	AvatarCreateResponse_UNKNOWN       AvatarCreateResponse_Status = 0
	AvatarCreateResponse_INVALID_PARAM AvatarCreateResponse_Status = -1
	AvatarCreateResponse_SUCCESS       AvatarCreateResponse_Status = 1
	AvatarCreateResponse_NOT_FOUND     AvatarCreateResponse_Status = 2
	AvatarCreateResponse_INVALID_TYPE  AvatarCreateResponse_Status = 3
)

var AvatarCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "NOT_FOUND",
	3:  "INVALID_TYPE",
}

var AvatarCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"NOT_FOUND":     2,
	"INVALID_TYPE":  3,
}

func (x AvatarCreateResponse_Status) String() string {
	return proto.EnumName(AvatarCreateResponse_Status_name, int32(x))
}

func (AvatarCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_082ad687f88f08fc, []int{1, 0}
}

type AvatarCreateRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	File                 []byte   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvatarCreateRequest) Reset()         { *m = AvatarCreateRequest{} }
func (m *AvatarCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AvatarCreateRequest) ProtoMessage()    {}
func (*AvatarCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_082ad687f88f08fc, []int{0}
}

func (m *AvatarCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvatarCreateRequest.Unmarshal(m, b)
}
func (m *AvatarCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvatarCreateRequest.Marshal(b, m, deterministic)
}
func (m *AvatarCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvatarCreateRequest.Merge(m, src)
}
func (m *AvatarCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AvatarCreateRequest.Size(m)
}
func (m *AvatarCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvatarCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvatarCreateRequest proto.InternalMessageInfo

func (m *AvatarCreateRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AvatarCreateRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type AvatarCreateResponse struct {
	Status               AvatarCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=AvatarCreateResponse_Status" json:"status,omitempty"`
	AvatarId             string                      `protobuf:"bytes,2,opt,name=avatarId,proto3" json:"avatarId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AvatarCreateResponse) Reset()         { *m = AvatarCreateResponse{} }
func (m *AvatarCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AvatarCreateResponse) ProtoMessage()    {}
func (*AvatarCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_082ad687f88f08fc, []int{1}
}

func (m *AvatarCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvatarCreateResponse.Unmarshal(m, b)
}
func (m *AvatarCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvatarCreateResponse.Marshal(b, m, deterministic)
}
func (m *AvatarCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvatarCreateResponse.Merge(m, src)
}
func (m *AvatarCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AvatarCreateResponse.Size(m)
}
func (m *AvatarCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvatarCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvatarCreateResponse proto.InternalMessageInfo

func (m *AvatarCreateResponse) GetStatus() AvatarCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return AvatarCreateResponse_UNKNOWN
}

func (m *AvatarCreateResponse) GetAvatarId() string {
	if m != nil {
		return m.AvatarId
	}
	return ""
}

func init() {
	proto.RegisterEnum("AvatarCreateResponse_Status", AvatarCreateResponse_Status_name, AvatarCreateResponse_Status_value)
	proto.RegisterType((*AvatarCreateRequest)(nil), "AvatarCreateRequest")
	proto.RegisterType((*AvatarCreateResponse)(nil), "AvatarCreateResponse")
}

func init() { proto.RegisterFile("avatar.proto", fileDescriptor_082ad687f88f08fc) }

var fileDescriptor_082ad687f88f08fc = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xcd, 0xd4, 0xe8, 0x9e, 0x9d, 0x84, 0xe7, 0x94, 0x51, 0x3c, 0x8c, 0x9e, 0x76, 0xea,
	0x61, 0xea, 0x59, 0x42, 0x37, 0xa1, 0xa8, 0xe9, 0x48, 0x57, 0xc5, 0x53, 0x89, 0x2c, 0x82, 0x20,
	0x76, 0x26, 0xa9, 0xff, 0xa3, 0x7f, 0x95, 0x62, 0xb2, 0x09, 0x42, 0x73, 0xca, 0x07, 0xbf, 0xf7,
	0xbd, 0xef, 0x7d, 0x10, 0xa9, 0x4f, 0xe5, 0x94, 0x49, 0xd7, 0xa6, 0x71, 0x4d, 0xc2, 0xe1, 0x84,
	0x7b, 0x9d, 0x19, 0xad, 0x9c, 0x96, 0xfa, 0xa3, 0xd5, 0xd6, 0xe1, 0x19, 0xd0, 0xd6, 0x6a, 0x93,
	0xaf, 0x46, 0x64, 0x4c, 0x26, 0xfb, 0x72, 0xa3, 0x10, 0x61, 0xef, 0xe5, 0xf5, 0x4d, 0x8f, 0x7a,
	0x63, 0x32, 0x89, 0xa4, 0xff, 0x27, 0x5f, 0x04, 0x86, 0xff, 0x3d, 0xec, 0xba, 0x79, 0xb7, 0x1a,
	0x2f, 0x81, 0x5a, 0xa7, 0x5c, 0x6b, 0xbd, 0xc9, 0xf1, 0xf4, 0x3c, 0xed, 0xc2, 0xd2, 0xd2, 0x33,
	0x72, 0xc3, 0x62, 0x0c, 0x87, 0x21, 0x61, 0xbe, 0xf2, 0x6b, 0xfa, 0xf2, 0x4f, 0x27, 0x35, 0xd0,
	0x40, 0xe3, 0x11, 0x1c, 0x54, 0xe2, 0x56, 0x14, 0x8f, 0x82, 0xed, 0x60, 0x0c, 0x83, 0x5c, 0x3c,
	0xf0, 0xbb, 0x7c, 0x56, 0x2f, 0xb8, 0xe4, 0xf7, 0xec, 0x7b, 0xfb, 0xc8, 0x2f, 0x58, 0x56, 0x59,
	0x36, 0x2f, 0x4b, 0x46, 0x70, 0x00, 0x7d, 0x51, 0x2c, 0xeb, 0x9b, 0xa2, 0x12, 0x33, 0xd6, 0x43,
	0x06, 0xd1, 0x76, 0x6e, 0xf9, 0xb4, 0x98, 0xb3, 0xdd, 0xe9, 0x35, 0xd0, 0x90, 0x11, 0xaf, 0x80,
	0x86, 0x9c, 0x38, 0x4c, 0x3b, 0x1a, 0x8a, 0x4f, 0x3b, 0x8f, 0x79, 0xa6, 0xbe, 0xd6, 0x8b, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x97, 0x7d, 0x83, 0x66, 0x01, 0x00, 0x00,
}
