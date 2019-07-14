// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sellinfo.proto

package sellinfo

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

type SellInfoCreateResponse_Status int32

const (
	SellInfoCreateResponse_UNKNOWN       SellInfoCreateResponse_Status = 0
	SellInfoCreateResponse_INVALID_PARAM SellInfoCreateResponse_Status = -1
	SellInfoCreateResponse_SUCCESS       SellInfoCreateResponse_Status = 1
	SellInfoCreateResponse_INVALID_TOKEN SellInfoCreateResponse_Status = 2
)

var SellInfoCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
}

var SellInfoCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
}

func (x SellInfoCreateResponse_Status) String() string {
	return proto.EnumName(SellInfoCreateResponse_Status_name, int32(x))
}

func (SellInfoCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{3, 0}
}

type ContentCreateRequest_Type int32

const (
	ContentCreateRequest_UNKNOWN ContentCreateRequest_Type = 0
	ContentCreateRequest_PICTURE ContentCreateRequest_Type = 1
	ContentCreateRequest_VIDEO   ContentCreateRequest_Type = 2
)

var ContentCreateRequest_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PICTURE",
	2: "VIDEO",
}

var ContentCreateRequest_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"PICTURE": 1,
	"VIDEO":   2,
}

func (x ContentCreateRequest_Type) String() string {
	return proto.EnumName(ContentCreateRequest_Type_name, int32(x))
}

func (ContentCreateRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{4, 0}
}

type ContentCreateResponse_Status int32

const (
	ContentCreateResponse_UNKNOWN       ContentCreateResponse_Status = 0
	ContentCreateResponse_INVALID_PARAM ContentCreateResponse_Status = -1
	ContentCreateResponse_SUCCESS       ContentCreateResponse_Status = 1
	ContentCreateResponse_INVALID_TOKEN ContentCreateResponse_Status = 2
)

var ContentCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
}

var ContentCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
}

func (x ContentCreateResponse_Status) String() string {
	return proto.EnumName(ContentCreateResponse_Status_name, int32(x))
}

func (ContentCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{5, 0}
}

type SellInfoMsg struct {
	SellInfoId           int32    `protobuf:"varint,1,opt,name=sellInfoId,proto3" json:"sellInfoId,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	ReleaseTime          int64    `protobuf:"varint,3,opt,name=releaseTime,proto3" json:"releaseTime,omitempty"`
	ValidTime            int64    `protobuf:"varint,4,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string   `protobuf:"bytes,5,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Price                float64  `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string   `protobuf:"bytes,8,opt,name=contentId,proto3" json:"contentId,omitempty"`
	UserId               int32    `protobuf:"varint,9,opt,name=userId,proto3" json:"userId,omitempty"`
	ClearEmpty           bool     `protobuf:"varint,99,opt,name=clearEmpty,proto3" json:"clearEmpty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoMsg) Reset()         { *m = SellInfoMsg{} }
func (m *SellInfoMsg) String() string { return proto.CompactTextString(m) }
func (*SellInfoMsg) ProtoMessage()    {}
func (*SellInfoMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{0}
}

func (m *SellInfoMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoMsg.Unmarshal(m, b)
}
func (m *SellInfoMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoMsg.Marshal(b, m, deterministic)
}
func (m *SellInfoMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoMsg.Merge(m, src)
}
func (m *SellInfoMsg) XXX_Size() int {
	return xxx_messageInfo_SellInfoMsg.Size(m)
}
func (m *SellInfoMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoMsg proto.InternalMessageInfo

func (m *SellInfoMsg) GetSellInfoId() int32 {
	if m != nil {
		return m.SellInfoId
	}
	return 0
}

func (m *SellInfoMsg) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SellInfoMsg) GetReleaseTime() int64 {
	if m != nil {
		return m.ReleaseTime
	}
	return 0
}

func (m *SellInfoMsg) GetValidTime() int64 {
	if m != nil {
		return m.ValidTime
	}
	return 0
}

func (m *SellInfoMsg) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *SellInfoMsg) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SellInfoMsg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SellInfoMsg) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *SellInfoMsg) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SellInfoMsg) GetClearEmpty() bool {
	if m != nil {
		return m.ClearEmpty
	}
	return false
}

type SellInfoQueryRequest struct {
	SellInfoId           int32    `protobuf:"varint,1,opt,name=sellInfoId,proto3" json:"sellInfoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoQueryRequest) Reset()         { *m = SellInfoQueryRequest{} }
func (m *SellInfoQueryRequest) String() string { return proto.CompactTextString(m) }
func (*SellInfoQueryRequest) ProtoMessage()    {}
func (*SellInfoQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{1}
}

func (m *SellInfoQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoQueryRequest.Unmarshal(m, b)
}
func (m *SellInfoQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoQueryRequest.Marshal(b, m, deterministic)
}
func (m *SellInfoQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoQueryRequest.Merge(m, src)
}
func (m *SellInfoQueryRequest) XXX_Size() int {
	return xxx_messageInfo_SellInfoQueryRequest.Size(m)
}
func (m *SellInfoQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoQueryRequest proto.InternalMessageInfo

func (m *SellInfoQueryRequest) GetSellInfoId() int32 {
	if m != nil {
		return m.SellInfoId
	}
	return 0
}

type SellInfoCreateRequest struct {
	ValidTime            int64    `protobuf:"varint,1,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string   `protobuf:"bytes,2,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string   `protobuf:"bytes,5,opt,name=contentId,proto3" json:"contentId,omitempty"`
	ContentToken         string   `protobuf:"bytes,6,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	UserId               int32    `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoCreateRequest) Reset()         { *m = SellInfoCreateRequest{} }
func (m *SellInfoCreateRequest) String() string { return proto.CompactTextString(m) }
func (*SellInfoCreateRequest) ProtoMessage()    {}
func (*SellInfoCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{2}
}

func (m *SellInfoCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoCreateRequest.Unmarshal(m, b)
}
func (m *SellInfoCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoCreateRequest.Marshal(b, m, deterministic)
}
func (m *SellInfoCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoCreateRequest.Merge(m, src)
}
func (m *SellInfoCreateRequest) XXX_Size() int {
	return xxx_messageInfo_SellInfoCreateRequest.Size(m)
}
func (m *SellInfoCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoCreateRequest proto.InternalMessageInfo

func (m *SellInfoCreateRequest) GetValidTime() int64 {
	if m != nil {
		return m.ValidTime
	}
	return 0
}

func (m *SellInfoCreateRequest) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *SellInfoCreateRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SellInfoCreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SellInfoCreateRequest) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *SellInfoCreateRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func (m *SellInfoCreateRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type SellInfoCreateResponse struct {
	Status               SellInfoCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=SellInfoCreateResponse_Status" json:"status,omitempty"`
	SellInfoId           int32                         `protobuf:"varint,2,opt,name=sellInfoId,proto3" json:"sellInfoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SellInfoCreateResponse) Reset()         { *m = SellInfoCreateResponse{} }
func (m *SellInfoCreateResponse) String() string { return proto.CompactTextString(m) }
func (*SellInfoCreateResponse) ProtoMessage()    {}
func (*SellInfoCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{3}
}

func (m *SellInfoCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoCreateResponse.Unmarshal(m, b)
}
func (m *SellInfoCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoCreateResponse.Marshal(b, m, deterministic)
}
func (m *SellInfoCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoCreateResponse.Merge(m, src)
}
func (m *SellInfoCreateResponse) XXX_Size() int {
	return xxx_messageInfo_SellInfoCreateResponse.Size(m)
}
func (m *SellInfoCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoCreateResponse proto.InternalMessageInfo

func (m *SellInfoCreateResponse) GetStatus() SellInfoCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return SellInfoCreateResponse_UNKNOWN
}

func (m *SellInfoCreateResponse) GetSellInfoId() int32 {
	if m != nil {
		return m.SellInfoId
	}
	return 0
}

type ContentCreateRequest struct {
	ContentId            string                    `protobuf:"bytes,1,opt,name=contentId,proto3" json:"contentId,omitempty"`
	ContentToken         string                    `protobuf:"bytes,2,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	Content              []byte                    `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Type                 ContentCreateRequest_Type `protobuf:"varint,4,opt,name=type,proto3,enum=ContentCreateRequest_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ContentCreateRequest) Reset()         { *m = ContentCreateRequest{} }
func (m *ContentCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ContentCreateRequest) ProtoMessage()    {}
func (*ContentCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{4}
}

func (m *ContentCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentCreateRequest.Unmarshal(m, b)
}
func (m *ContentCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentCreateRequest.Marshal(b, m, deterministic)
}
func (m *ContentCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentCreateRequest.Merge(m, src)
}
func (m *ContentCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ContentCreateRequest.Size(m)
}
func (m *ContentCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentCreateRequest proto.InternalMessageInfo

func (m *ContentCreateRequest) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *ContentCreateRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func (m *ContentCreateRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ContentCreateRequest) GetType() ContentCreateRequest_Type {
	if m != nil {
		return m.Type
	}
	return ContentCreateRequest_UNKNOWN
}

type ContentCreateResponse struct {
	Status               ContentCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ContentCreateResponse_Status" json:"status,omitempty"`
	ContentId            string                       `protobuf:"bytes,2,opt,name=contentId,proto3" json:"contentId,omitempty"`
	ContentToken         string                       `protobuf:"bytes,3,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ContentCreateResponse) Reset()         { *m = ContentCreateResponse{} }
func (m *ContentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ContentCreateResponse) ProtoMessage()    {}
func (*ContentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{5}
}

func (m *ContentCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentCreateResponse.Unmarshal(m, b)
}
func (m *ContentCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentCreateResponse.Marshal(b, m, deterministic)
}
func (m *ContentCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentCreateResponse.Merge(m, src)
}
func (m *ContentCreateResponse) XXX_Size() int {
	return xxx_messageInfo_ContentCreateResponse.Size(m)
}
func (m *ContentCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentCreateResponse proto.InternalMessageInfo

func (m *ContentCreateResponse) GetStatus() ContentCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return ContentCreateResponse_UNKNOWN
}

func (m *ContentCreateResponse) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *ContentCreateResponse) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

type SellInfoFindRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoFindRequest) Reset()         { *m = SellInfoFindRequest{} }
func (m *SellInfoFindRequest) String() string { return proto.CompactTextString(m) }
func (*SellInfoFindRequest) ProtoMessage()    {}
func (*SellInfoFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{6}
}

func (m *SellInfoFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoFindRequest.Unmarshal(m, b)
}
func (m *SellInfoFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoFindRequest.Marshal(b, m, deterministic)
}
func (m *SellInfoFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoFindRequest.Merge(m, src)
}
func (m *SellInfoFindRequest) XXX_Size() int {
	return xxx_messageInfo_SellInfoFindRequest.Size(m)
}
func (m *SellInfoFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoFindRequest proto.InternalMessageInfo

func (m *SellInfoFindRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SellInfoFindRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SellInfoFindRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SellInfoFindResponse struct {
	SellInfo             []*SellInfoMsg `protobuf:"bytes,1,rep,name=sellInfo,proto3" json:"sellInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SellInfoFindResponse) Reset()         { *m = SellInfoFindResponse{} }
func (m *SellInfoFindResponse) String() string { return proto.CompactTextString(m) }
func (*SellInfoFindResponse) ProtoMessage()    {}
func (*SellInfoFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{7}
}

func (m *SellInfoFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoFindResponse.Unmarshal(m, b)
}
func (m *SellInfoFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoFindResponse.Marshal(b, m, deterministic)
}
func (m *SellInfoFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoFindResponse.Merge(m, src)
}
func (m *SellInfoFindResponse) XXX_Size() int {
	return xxx_messageInfo_SellInfoFindResponse.Size(m)
}
func (m *SellInfoFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoFindResponse proto.InternalMessageInfo

func (m *SellInfoFindResponse) GetSellInfo() []*SellInfoMsg {
	if m != nil {
		return m.SellInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("SellInfoCreateResponse_Status", SellInfoCreateResponse_Status_name, SellInfoCreateResponse_Status_value)
	proto.RegisterEnum("ContentCreateRequest_Type", ContentCreateRequest_Type_name, ContentCreateRequest_Type_value)
	proto.RegisterEnum("ContentCreateResponse_Status", ContentCreateResponse_Status_name, ContentCreateResponse_Status_value)
	proto.RegisterType((*SellInfoMsg)(nil), "SellInfoMsg")
	proto.RegisterType((*SellInfoQueryRequest)(nil), "SellInfoQueryRequest")
	proto.RegisterType((*SellInfoCreateRequest)(nil), "SellInfoCreateRequest")
	proto.RegisterType((*SellInfoCreateResponse)(nil), "SellInfoCreateResponse")
	proto.RegisterType((*ContentCreateRequest)(nil), "ContentCreateRequest")
	proto.RegisterType((*ContentCreateResponse)(nil), "ContentCreateResponse")
	proto.RegisterType((*SellInfoFindRequest)(nil), "SellInfoFindRequest")
	proto.RegisterType((*SellInfoFindResponse)(nil), "SellInfoFindResponse")
}

func init() { proto.RegisterFile("sellinfo.proto", fileDescriptor_9c322b32f573704b) }

var fileDescriptor_9c322b32f573704b = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x59, 0x27, 0xce, 0x9f, 0x49, 0x5b, 0x85, 0x25, 0x09, 0x96, 0x05, 0x95, 0xe5, 0x93,
	0x25, 0xd0, 0x1e, 0x52, 0x51, 0xc4, 0x8d, 0x92, 0x06, 0xc9, 0x2a, 0x4d, 0xdb, 0x4d, 0x5a, 0x0e,
	0x1c, 0x90, 0x89, 0x37, 0x95, 0x85, 0x63, 0x1b, 0xaf, 0x83, 0x94, 0x37, 0xe2, 0xc6, 0x43, 0xf0,
	0x00, 0x3c, 0x04, 0x0f, 0x02, 0xf2, 0xda, 0x8e, 0xed, 0xd4, 0x6d, 0x4e, 0xe4, 0xf6, 0xcd, 0xcc,
	0x7a, 0x77, 0x7e, 0xf3, 0xed, 0x06, 0x0e, 0x38, 0x73, 0x5d, 0xc7, 0x5b, 0xf8, 0x24, 0x08, 0xfd,
	0xc8, 0xd7, 0x7f, 0x4a, 0xd0, 0x99, 0x32, 0xd7, 0x35, 0xbd, 0x85, 0x7f, 0xce, 0x6f, 0xf1, 0x21,
	0x00, 0x4f, 0xa5, 0x69, 0x2b, 0x48, 0x43, 0x86, 0x4c, 0x0b, 0x11, 0x3c, 0x80, 0x06, 0x8f, 0xac,
	0x68, 0xc5, 0x15, 0x49, 0xe4, 0x52, 0x85, 0x35, 0xe8, 0x84, 0xcc, 0x65, 0x16, 0x67, 0x33, 0x67,
	0xc9, 0x94, 0x9a, 0x86, 0x8c, 0x1a, 0x2d, 0x86, 0xf0, 0x33, 0x68, 0x7f, 0xb7, 0x5c, 0xc7, 0x16,
	0xf9, 0xba, 0xc8, 0xe7, 0x01, 0xac, 0x42, 0xeb, 0xd6, 0xf7, 0xed, 0x89, 0xb5, 0x64, 0x8a, 0xac,
	0x21, 0xa3, 0x4d, 0x37, 0x1a, 0xf7, 0x40, 0x0e, 0x42, 0x67, 0xce, 0x94, 0x86, 0x86, 0x0c, 0x44,
	0x13, 0x11, 0xef, 0x68, 0x33, 0x3e, 0x0f, 0x9d, 0x20, 0x72, 0x7c, 0x4f, 0x69, 0x8a, 0x45, 0xc5,
	0x50, 0xbc, 0xe3, 0xdc, 0xf7, 0x22, 0xe6, 0x45, 0xa6, 0xad, 0xb4, 0x44, 0x3e, 0x0f, 0xc4, 0x9d,
	0xac, 0x38, 0x0b, 0x4d, 0x5b, 0x69, 0x27, 0x9d, 0x24, 0x2a, 0x26, 0x30, 0x77, 0x99, 0x15, 0x8e,
	0x97, 0x41, 0xb4, 0x56, 0xe6, 0x1a, 0x32, 0x5a, 0xb4, 0x10, 0xd1, 0x8f, 0xa1, 0x97, 0x01, 0xbb,
	0x5a, 0xb1, 0x70, 0x4d, 0xd9, 0xb7, 0x15, 0xe3, 0xd1, 0x2e, 0x72, 0xfa, 0x1f, 0x04, 0xfd, 0x6c,
	0xe1, 0x28, 0x64, 0x56, 0xc4, 0xb2, 0x95, 0x25, 0x32, 0xe8, 0x21, 0x32, 0xd2, 0x7d, 0x64, 0x6a,
	0x0f, 0x90, 0xa9, 0xef, 0x20, 0x23, 0x6f, 0x93, 0xd1, 0x61, 0x2f, 0x15, 0x33, 0xff, 0x2b, 0xf3,
	0x04, 0xf6, 0x36, 0x2d, 0xc5, 0x0a, 0xf4, 0x9a, 0x45, 0x7a, 0xfa, 0x2f, 0x04, 0x83, 0xed, 0x2e,
	0x79, 0xe0, 0x7b, 0x9c, 0xe1, 0xe3, 0x8d, 0x75, 0xe2, 0x1e, 0x0f, 0x86, 0x87, 0xa4, 0xba, 0x90,
	0x4c, 0x45, 0xd5, 0xc6, 0x5a, 0x65, 0xb0, 0xd2, 0x1d, 0xb0, 0x57, 0xd0, 0x48, 0x56, 0xe0, 0x0e,
	0x34, 0xaf, 0x27, 0x67, 0x93, 0x8b, 0x8f, 0x93, 0xee, 0x23, 0xac, 0xc2, 0xbe, 0x39, 0xb9, 0x39,
	0xf9, 0x60, 0x9e, 0x7e, 0xbe, 0x3c, 0xa1, 0x27, 0xe7, 0xdd, 0xbf, 0xd9, 0x0f, 0xc5, 0x85, 0xd3,
	0xeb, 0xd1, 0x68, 0x3c, 0x9d, 0x76, 0x11, 0x7e, 0x9c, 0x17, 0xce, 0x2e, 0xce, 0xc6, 0x93, 0xae,
	0xa4, 0xff, 0x46, 0xd0, 0x1b, 0x25, 0xed, 0xde, 0x19, 0x55, 0x0e, 0x0e, 0xed, 0x02, 0x27, 0x55,
	0x80, 0x53, 0xa0, 0x99, 0x6a, 0x31, 0xb4, 0x3d, 0x9a, 0x49, 0x4c, 0xa0, 0x1e, 0xad, 0x83, 0xe4,
	0x6e, 0x1c, 0x0c, 0x55, 0x52, 0x75, 0x00, 0x32, 0x5b, 0x07, 0x8c, 0x8a, 0x3a, 0xfd, 0x05, 0xd4,
	0x63, 0x55, 0xee, 0xba, 0x03, 0xcd, 0x4b, 0x73, 0x34, 0xbb, 0xa6, 0xe3, 0x2e, 0xc2, 0x6d, 0x90,
	0x6f, 0xcc, 0xd3, 0xf1, 0x45, 0x57, 0x12, 0xee, 0xdb, 0xfa, 0x60, 0x3a, 0x96, 0x57, 0x5b, 0x63,
	0x79, 0x4e, 0x2a, 0xeb, 0xb6, 0xa7, 0x52, 0x22, 0x21, 0xed, 0x22, 0x51, 0xbb, 0x4b, 0xe2, 0x7f,
	0xcc, 0xed, 0x13, 0x3c, 0xc9, 0x3c, 0xf5, 0xde, 0xf1, 0xec, 0x6c, 0x6a, 0xb9, 0x59, 0x51, 0xe9,
	0xaa, 0xf7, 0x40, 0x76, 0x9d, 0xa5, 0x13, 0x89, 0xf3, 0xef, 0xd3, 0x44, 0xc4, 0xd5, 0xfe, 0x62,
	0xc1, 0x59, 0x32, 0xa0, 0x7d, 0x9a, 0x2a, 0xfd, 0x6d, 0x7e, 0xf1, 0x93, 0x8f, 0xa7, 0x00, 0x0d,
	0x68, 0x65, 0x6e, 0x54, 0x90, 0x56, 0x33, 0x3a, 0xc3, 0x3d, 0x52, 0x78, 0x52, 0xe9, 0x26, 0x3b,
	0xfc, 0x81, 0xa0, 0x95, 0x65, 0xf0, 0x4b, 0x90, 0xc5, 0xfb, 0x81, 0xfb, 0xa4, 0xea, 0x3d, 0x51,
	0x4b, 0x1f, 0xc1, 0x6f, 0xa0, 0x91, 0xcc, 0x03, 0x0f, 0x48, 0xe5, 0x2b, 0xa2, 0x3e, 0xbd, 0xe7,
	0x3a, 0xe1, 0x23, 0xa8, 0xc7, 0xe7, 0xc5, 0x3d, 0x52, 0xc1, 0x46, 0xed, 0x93, 0xaa, 0xa6, 0x86,
	0xef, 0xa0, 0x99, 0xda, 0x00, 0xbf, 0xde, 0x6c, 0xdd, 0xaf, 0xf4, 0xa4, 0x3a, 0xa8, 0x76, 0xcc,
	0x97, 0x86, 0xf8, 0x8b, 0x39, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xee, 0xea, 0x99, 0x4d, 0x74,
	0x06, 0x00, 0x00,
}
