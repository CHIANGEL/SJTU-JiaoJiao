// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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

type MessageInfo_Type int32

const (
	MessageInfo_UNKNOWN MessageInfo_Type = 0
	MessageInfo_TEXT    MessageInfo_Type = 1
	MessageInfo_PICTURE MessageInfo_Type = 2
	MessageInfo_VIDEO   MessageInfo_Type = 3
)

var MessageInfo_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "TEXT",
	2: "PICTURE",
	3: "VIDEO",
}

var MessageInfo_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"TEXT":    1,
	"PICTURE": 2,
	"VIDEO":   3,
}

func (x MessageInfo_Type) String() string {
	return proto.EnumName(MessageInfo_Type_name, int32(x))
}

func (MessageInfo_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type MessageCreateRequest_Type int32

const (
	MessageCreateRequest_UNKNOWN MessageCreateRequest_Type = 0
	MessageCreateRequest_TEXT    MessageCreateRequest_Type = 1
	MessageCreateRequest_PICTURE MessageCreateRequest_Type = 2
	MessageCreateRequest_VIDEO   MessageCreateRequest_Type = 3
)

var MessageCreateRequest_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "TEXT",
	2: "PICTURE",
	3: "VIDEO",
}

var MessageCreateRequest_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"TEXT":    1,
	"PICTURE": 2,
	"VIDEO":   3,
}

func (x MessageCreateRequest_Type) String() string {
	return proto.EnumName(MessageCreateRequest_Type_name, int32(x))
}

func (MessageCreateRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1, 0}
}

type MessageCreateResponse_Status int32

const (
	MessageCreateResponse_UNKNOWN       MessageCreateResponse_Status = 0
	MessageCreateResponse_INVALID_PARAM MessageCreateResponse_Status = -1
	MessageCreateResponse_SUCCESS       MessageCreateResponse_Status = 1
)

var MessageCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
}

var MessageCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
}

func (x MessageCreateResponse_Status) String() string {
	return proto.EnumName(MessageCreateResponse_Status_name, int32(x))
}

func (MessageCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2, 0}
}

type MessageFindRequest_Way int32

const (
	MessageFindRequest_UNKNOWN      MessageFindRequest_Way = 0
	MessageFindRequest_ONLY_PULL    MessageFindRequest_Way = 1
	MessageFindRequest_READ_MESSAGE MessageFindRequest_Way = 2
	MessageFindRequest_HISTORY      MessageFindRequest_Way = 3
)

var MessageFindRequest_Way_name = map[int32]string{
	0: "UNKNOWN",
	1: "ONLY_PULL",
	2: "READ_MESSAGE",
	3: "HISTORY",
}

var MessageFindRequest_Way_value = map[string]int32{
	"UNKNOWN":      0,
	"ONLY_PULL":    1,
	"READ_MESSAGE": 2,
	"HISTORY":      3,
}

func (x MessageFindRequest_Way) String() string {
	return proto.EnumName(MessageFindRequest_Way_name, int32(x))
}

func (MessageFindRequest_Way) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3, 0}
}

type MessageFindResponse_Status int32

const (
	MessageFindResponse_UNKNOWN       MessageFindResponse_Status = 0
	MessageFindResponse_INVALID_PARAM MessageFindResponse_Status = -1
	MessageFindResponse_SUCCESS       MessageFindResponse_Status = 1
	MessageFindResponse_NOT_FOUND     MessageFindResponse_Status = 2
)

var MessageFindResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "NOT_FOUND",
}

var MessageFindResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"NOT_FOUND":     2,
}

func (x MessageFindResponse_Status) String() string {
	return proto.EnumName(MessageFindResponse_Status_name, int32(x))
}

func (MessageFindResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4, 0}
}

type MessageInfo struct {
	Time                 int64            `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Forward              bool             `protobuf:"varint,2,opt,name=forward,proto3" json:"forward,omitempty"`
	Type                 MessageInfo_Type `protobuf:"varint,3,opt,name=type,proto3,enum=MessageInfo_Type" json:"type,omitempty"`
	Text                 string           `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Unread               bool             `protobuf:"varint,5,opt,name=unread,proto3" json:"unread,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MessageInfo) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

func (m *MessageInfo) GetType() MessageInfo_Type {
	if m != nil {
		return m.Type
	}
	return MessageInfo_UNKNOWN
}

func (m *MessageInfo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *MessageInfo) GetUnread() bool {
	if m != nil {
		return m.Unread
	}
	return false
}

type MessageCreateRequest struct {
	FromUser             int32                     `protobuf:"varint,1,opt,name=fromUser,proto3" json:"fromUser,omitempty"`
	ToUser               int32                     `protobuf:"varint,2,opt,name=toUser,proto3" json:"toUser,omitempty"`
	Type                 MessageCreateRequest_Type `protobuf:"varint,3,opt,name=type,proto3,enum=MessageCreateRequest_Type" json:"type,omitempty"`
	Text                 string                    `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	File                 []byte                    `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MessageCreateRequest) Reset()         { *m = MessageCreateRequest{} }
func (m *MessageCreateRequest) String() string { return proto.CompactTextString(m) }
func (*MessageCreateRequest) ProtoMessage()    {}
func (*MessageCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *MessageCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageCreateRequest.Unmarshal(m, b)
}
func (m *MessageCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageCreateRequest.Marshal(b, m, deterministic)
}
func (m *MessageCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageCreateRequest.Merge(m, src)
}
func (m *MessageCreateRequest) XXX_Size() int {
	return xxx_messageInfo_MessageCreateRequest.Size(m)
}
func (m *MessageCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageCreateRequest proto.InternalMessageInfo

func (m *MessageCreateRequest) GetFromUser() int32 {
	if m != nil {
		return m.FromUser
	}
	return 0
}

func (m *MessageCreateRequest) GetToUser() int32 {
	if m != nil {
		return m.ToUser
	}
	return 0
}

func (m *MessageCreateRequest) GetType() MessageCreateRequest_Type {
	if m != nil {
		return m.Type
	}
	return MessageCreateRequest_UNKNOWN
}

func (m *MessageCreateRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *MessageCreateRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type MessageCreateResponse struct {
	Status               MessageCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=MessageCreateResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MessageCreateResponse) Reset()         { *m = MessageCreateResponse{} }
func (m *MessageCreateResponse) String() string { return proto.CompactTextString(m) }
func (*MessageCreateResponse) ProtoMessage()    {}
func (*MessageCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *MessageCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageCreateResponse.Unmarshal(m, b)
}
func (m *MessageCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageCreateResponse.Marshal(b, m, deterministic)
}
func (m *MessageCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageCreateResponse.Merge(m, src)
}
func (m *MessageCreateResponse) XXX_Size() int {
	return xxx_messageInfo_MessageCreateResponse.Size(m)
}
func (m *MessageCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageCreateResponse proto.InternalMessageInfo

func (m *MessageCreateResponse) GetStatus() MessageCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return MessageCreateResponse_UNKNOWN
}

type MessageFindRequest struct {
	FromUser             int32                  `protobuf:"varint,1,opt,name=fromUser,proto3" json:"fromUser,omitempty"`
	ToUser               int32                  `protobuf:"varint,2,opt,name=toUser,proto3" json:"toUser,omitempty"`
	Way                  MessageFindRequest_Way `protobuf:"varint,3,opt,name=way,proto3,enum=MessageFindRequest_Way" json:"way,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MessageFindRequest) Reset()         { *m = MessageFindRequest{} }
func (m *MessageFindRequest) String() string { return proto.CompactTextString(m) }
func (*MessageFindRequest) ProtoMessage()    {}
func (*MessageFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *MessageFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageFindRequest.Unmarshal(m, b)
}
func (m *MessageFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageFindRequest.Marshal(b, m, deterministic)
}
func (m *MessageFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageFindRequest.Merge(m, src)
}
func (m *MessageFindRequest) XXX_Size() int {
	return xxx_messageInfo_MessageFindRequest.Size(m)
}
func (m *MessageFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageFindRequest proto.InternalMessageInfo

func (m *MessageFindRequest) GetFromUser() int32 {
	if m != nil {
		return m.FromUser
	}
	return 0
}

func (m *MessageFindRequest) GetToUser() int32 {
	if m != nil {
		return m.ToUser
	}
	return 0
}

func (m *MessageFindRequest) GetWay() MessageFindRequest_Way {
	if m != nil {
		return m.Way
	}
	return MessageFindRequest_UNKNOWN
}

type MessageFindResponse struct {
	Status               MessageFindResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=MessageFindResponse_Status" json:"status,omitempty"`
	FromUser             int32                      `protobuf:"varint,2,opt,name=fromUser,proto3" json:"fromUser,omitempty"`
	ToUser               int32                      `protobuf:"varint,3,opt,name=toUser,proto3" json:"toUser,omitempty"`
	Badge                int32                      `protobuf:"varint,4,opt,name=badge,proto3" json:"badge,omitempty"`
	Infos                []*MessageInfo             `protobuf:"bytes,5,rep,name=infos,proto3" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MessageFindResponse) Reset()         { *m = MessageFindResponse{} }
func (m *MessageFindResponse) String() string { return proto.CompactTextString(m) }
func (*MessageFindResponse) ProtoMessage()    {}
func (*MessageFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *MessageFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageFindResponse.Unmarshal(m, b)
}
func (m *MessageFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageFindResponse.Marshal(b, m, deterministic)
}
func (m *MessageFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageFindResponse.Merge(m, src)
}
func (m *MessageFindResponse) XXX_Size() int {
	return xxx_messageInfo_MessageFindResponse.Size(m)
}
func (m *MessageFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageFindResponse proto.InternalMessageInfo

func (m *MessageFindResponse) GetStatus() MessageFindResponse_Status {
	if m != nil {
		return m.Status
	}
	return MessageFindResponse_UNKNOWN
}

func (m *MessageFindResponse) GetFromUser() int32 {
	if m != nil {
		return m.FromUser
	}
	return 0
}

func (m *MessageFindResponse) GetToUser() int32 {
	if m != nil {
		return m.ToUser
	}
	return 0
}

func (m *MessageFindResponse) GetBadge() int32 {
	if m != nil {
		return m.Badge
	}
	return 0
}

func (m *MessageFindResponse) GetInfos() []*MessageInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

func init() {
	proto.RegisterEnum("MessageInfo_Type", MessageInfo_Type_name, MessageInfo_Type_value)
	proto.RegisterEnum("MessageCreateRequest_Type", MessageCreateRequest_Type_name, MessageCreateRequest_Type_value)
	proto.RegisterEnum("MessageCreateResponse_Status", MessageCreateResponse_Status_name, MessageCreateResponse_Status_value)
	proto.RegisterEnum("MessageFindRequest_Way", MessageFindRequest_Way_name, MessageFindRequest_Way_value)
	proto.RegisterEnum("MessageFindResponse_Status", MessageFindResponse_Status_name, MessageFindResponse_Status_value)
	proto.RegisterType((*MessageInfo)(nil), "MessageInfo")
	proto.RegisterType((*MessageCreateRequest)(nil), "MessageCreateRequest")
	proto.RegisterType((*MessageCreateResponse)(nil), "MessageCreateResponse")
	proto.RegisterType((*MessageFindRequest)(nil), "MessageFindRequest")
	proto.RegisterType((*MessageFindResponse)(nil), "MessageFindResponse")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x76, 0xf2, 0xd3, 0x9f, 0xd3, 0x76, 0x19, 0x67, 0xbb, 0x6b, 0xa8, 0x08, 0x25, 0x20, 0xd4,
	0x9b, 0x01, 0xbb, 0x2c, 0xde, 0x89, 0xa5, 0xcd, 0x6a, 0xb0, 0x4d, 0xca, 0x24, 0xd9, 0xba, 0x57,
	0x25, 0x4b, 0xa7, 0x4b, 0xc1, 0x26, 0x35, 0x49, 0x59, 0xf3, 0x04, 0xde, 0xf9, 0x36, 0x3e, 0x82,
	0x0f, 0xe1, 0xd3, 0x28, 0x99, 0x64, 0x65, 0x53, 0x23, 0x82, 0x9b, 0xab, 0x39, 0x3f, 0xf9, 0xce,
	0xf7, 0x7d, 0x73, 0x12, 0xe8, 0x6c, 0x79, 0x1c, 0xfb, 0x37, 0x9c, 0xee, 0xa2, 0x30, 0x09, 0xf5,
	0xef, 0x08, 0x5a, 0xb3, 0x3c, 0x63, 0x06, 0xeb, 0x90, 0x10, 0x50, 0x92, 0xcd, 0x96, 0x6b, 0xa8,
	0x8f, 0x06, 0x32, 0x13, 0x67, 0xa2, 0x41, 0x7d, 0x1d, 0x46, 0xb7, 0x7e, 0xb4, 0xd2, 0xa4, 0x3e,
	0x1a, 0x34, 0xd8, 0x5d, 0x48, 0x9e, 0x83, 0x92, 0xa4, 0x3b, 0xae, 0xc9, 0x7d, 0x34, 0x38, 0x1a,
	0x3e, 0xa6, 0xf7, 0x90, 0xa8, 0x9b, 0xee, 0x38, 0x13, 0x65, 0x01, 0xca, 0x3f, 0x27, 0x9a, 0xd2,
	0x47, 0x83, 0x26, 0x13, 0x67, 0x72, 0x0a, 0xb5, 0x7d, 0x10, 0x71, 0x7f, 0xa5, 0xa9, 0x02, 0xb3,
	0x88, 0xf4, 0x73, 0x50, 0xb2, 0x37, 0x49, 0x0b, 0xea, 0x9e, 0xf5, 0xde, 0xb2, 0x17, 0x16, 0x7e,
	0x44, 0x1a, 0xa0, 0xb8, 0xc6, 0x07, 0x17, 0xa3, 0x2c, 0x3d, 0x37, 0xc7, 0xae, 0xc7, 0x0c, 0x2c,
	0x91, 0x26, 0xa8, 0x97, 0xe6, 0xc4, 0xb0, 0xb1, 0xac, 0xff, 0x40, 0xd0, 0x2d, 0xa6, 0x8f, 0x23,
	0xee, 0x27, 0x9c, 0xf1, 0x4f, 0x7b, 0x1e, 0x27, 0xa4, 0x07, 0x8d, 0x75, 0x14, 0x6e, 0xbd, 0x98,
	0x47, 0x42, 0x94, 0xca, 0x7e, 0xc7, 0x19, 0x87, 0x24, 0x14, 0x15, 0x49, 0x54, 0x8a, 0x88, 0xd0,
	0x92, 0xac, 0x1e, 0xad, 0x02, 0xfe, 0x97, 0x3e, 0x02, 0xca, 0x7a, 0xf3, 0x91, 0x0b, 0x75, 0x6d,
	0x26, 0xce, 0xff, 0xab, 0xed, 0x2b, 0x82, 0x93, 0x03, 0x0a, 0xf1, 0x2e, 0x0c, 0x62, 0x4e, 0xce,
	0xa1, 0x16, 0x27, 0x7e, 0xb2, 0x8f, 0x85, 0xb4, 0xa3, 0xe1, 0x33, 0x5a, 0xd9, 0x47, 0x1d, 0xd1,
	0xc4, 0x8a, 0x66, 0xfd, 0x35, 0xd4, 0xf2, 0x4c, 0x99, 0x49, 0x0f, 0x3a, 0xa6, 0x75, 0x39, 0x9a,
	0x9a, 0x93, 0xe5, 0x7c, 0xc4, 0x46, 0x33, 0xfc, 0xf3, 0xee, 0x11, 0xdc, 0x1c, 0x6f, 0x3c, 0x36,
	0x1c, 0x07, 0x23, 0xfd, 0x1b, 0x02, 0x52, 0x0c, 0xba, 0xd8, 0x04, 0xab, 0x87, 0x58, 0xfd, 0x02,
	0xe4, 0x5b, 0x3f, 0x2d, 0x9c, 0x7e, 0x42, 0xff, 0x44, 0xa5, 0x0b, 0x3f, 0x65, 0x59, 0x8f, 0xfe,
	0x06, 0xe4, 0x85, 0x9f, 0x96, 0x29, 0x77, 0xa0, 0x69, 0x5b, 0xd3, 0xab, 0xe5, 0xdc, 0x9b, 0x4e,
	0x31, 0x22, 0x18, 0xda, 0xcc, 0x18, 0x4d, 0x96, 0x33, 0xc3, 0x71, 0x46, 0x6f, 0x33, 0x1b, 0x5b,
	0x50, 0x7f, 0x67, 0x3a, 0xae, 0xcd, 0xae, 0xb0, 0xac, 0x7f, 0x91, 0xe0, 0xb8, 0x34, 0xa1, 0xb0,
	0xf1, 0xec, 0xc0, 0xc6, 0xa7, 0xb4, 0xa2, 0xeb, 0xc0, 0xc4, 0x92, 0x5a, 0xe9, 0xaf, 0x6a, 0xe5,
	0x92, 0xda, 0x2e, 0xa8, 0xd7, 0xfe, 0xea, 0x86, 0x8b, 0x4d, 0x51, 0x59, 0x1e, 0x10, 0x1d, 0xd4,
	0x4d, 0xb0, 0x0e, 0x63, 0x4d, 0xed, 0xcb, 0x83, 0xd6, 0xb0, 0x7d, 0xff, 0x33, 0x62, 0x79, 0x49,
	0x9f, 0x3d, 0xec, 0xca, 0x32, 0xa3, 0x2c, 0xdb, 0x5d, 0x5e, 0xd8, 0x9e, 0x35, 0xc1, 0xd2, 0x70,
	0x0f, 0xf5, 0x62, 0x08, 0x79, 0x05, 0xb5, 0x7c, 0x5b, 0xc8, 0x49, 0xe5, 0xa2, 0xf7, 0x4e, 0xab,
	0x97, 0x8a, 0xbc, 0x04, 0x25, 0xf3, 0x87, 0x1c, 0x57, 0xdc, 0x5a, 0xaf, 0x5b, 0x65, 0xe1, 0x75,
	0x4d, 0xfc, 0x74, 0xce, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x58, 0xf9, 0x5f, 0x85, 0x04,
	0x00, 0x00,
}
