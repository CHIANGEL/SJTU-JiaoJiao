// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content.proto

package content

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
	return fileDescriptor_61cc9617ce0cf609, []int{0, 0}
}

type ContentCreateResponse_Status int32

const (
	ContentCreateResponse_UNKNOWN       ContentCreateResponse_Status = 0
	ContentCreateResponse_INVALID_PARAM ContentCreateResponse_Status = -1
	ContentCreateResponse_SUCCESS       ContentCreateResponse_Status = 1
	ContentCreateResponse_INVALID_TOKEN ContentCreateResponse_Status = 2
	ContentCreateResponse_INVALID_TYPE  ContentCreateResponse_Status = 3
)

var ContentCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
	3:  "INVALID_TYPE",
}

var ContentCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
	"INVALID_TYPE":  3,
}

func (x ContentCreateResponse_Status) String() string {
	return proto.EnumName(ContentCreateResponse_Status_name, int32(x))
}

func (ContentCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{1, 0}
}

type ContentUpdateRequest_Type int32

const (
	ContentUpdateRequest_UNKNOWN ContentUpdateRequest_Type = 0
	ContentUpdateRequest_PICTURE ContentUpdateRequest_Type = 1
	ContentUpdateRequest_VIDEO   ContentUpdateRequest_Type = 2
)

var ContentUpdateRequest_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PICTURE",
	2: "VIDEO",
}

var ContentUpdateRequest_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"PICTURE": 1,
	"VIDEO":   2,
}

func (x ContentUpdateRequest_Type) String() string {
	return proto.EnumName(ContentUpdateRequest_Type_name, int32(x))
}

func (ContentUpdateRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{2, 0}
}

type ContentUpdateResponse_Status int32

const (
	ContentUpdateResponse_UNKNOWN       ContentUpdateResponse_Status = 0
	ContentUpdateResponse_INVALID_PARAM ContentUpdateResponse_Status = -1
	ContentUpdateResponse_SUCCESS       ContentUpdateResponse_Status = 1
	ContentUpdateResponse_INVALID_TOKEN ContentUpdateResponse_Status = 2
	ContentUpdateResponse_NOT_FOUND     ContentUpdateResponse_Status = 3
	ContentUpdateResponse_FAILED        ContentUpdateResponse_Status = 4
	ContentUpdateResponse_INVALID_TYPE  ContentUpdateResponse_Status = 5
)

var ContentUpdateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
	3:  "NOT_FOUND",
	4:  "FAILED",
	5:  "INVALID_TYPE",
}

var ContentUpdateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
	"NOT_FOUND":     3,
	"FAILED":        4,
	"INVALID_TYPE":  5,
}

func (x ContentUpdateResponse_Status) String() string {
	return proto.EnumName(ContentUpdateResponse_Status_name, int32(x))
}

func (ContentUpdateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{3, 0}
}

type ContentDeleteResponse_Status int32

const (
	ContentDeleteResponse_UNKNOWN       ContentDeleteResponse_Status = 0
	ContentDeleteResponse_INVALID_PARAM ContentDeleteResponse_Status = -1
	ContentDeleteResponse_SUCCESS       ContentDeleteResponse_Status = 1
	ContentDeleteResponse_INVALID_TOKEN ContentDeleteResponse_Status = 2
)

var ContentDeleteResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
}

var ContentDeleteResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
}

func (x ContentDeleteResponse_Status) String() string {
	return proto.EnumName(ContentDeleteResponse_Status_name, int32(x))
}

func (ContentDeleteResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{5, 0}
}

type ContentMsg_Type int32

const (
	ContentMsg_UNKNOWN ContentMsg_Type = 0
	ContentMsg_PICTURE ContentMsg_Type = 1
	ContentMsg_VIDEO   ContentMsg_Type = 2
)

var ContentMsg_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PICTURE",
	2: "VIDEO",
}

var ContentMsg_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"PICTURE": 1,
	"VIDEO":   2,
}

func (x ContentMsg_Type) String() string {
	return proto.EnumName(ContentMsg_Type_name, int32(x))
}

func (ContentMsg_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{6, 0}
}

type ContentQueryResponse_Status int32

const (
	ContentQueryResponse_UNKNOWN       ContentQueryResponse_Status = 0
	ContentQueryResponse_INVALID_PARAM ContentQueryResponse_Status = -1
	ContentQueryResponse_SUCCESS       ContentQueryResponse_Status = 1
	ContentQueryResponse_NOT_FOUND     ContentQueryResponse_Status = 2
)

var ContentQueryResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "NOT_FOUND",
}

var ContentQueryResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"NOT_FOUND":     2,
}

func (x ContentQueryResponse_Status) String() string {
	return proto.EnumName(ContentQueryResponse_Status_name, int32(x))
}

func (ContentQueryResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{8, 0}
}

type ContentCheckResponse_Status int32

const (
	ContentCheckResponse_UNKNOWN       ContentCheckResponse_Status = 0
	ContentCheckResponse_INVALID_PARAM ContentCheckResponse_Status = -1
	ContentCheckResponse_VALID         ContentCheckResponse_Status = 1
	ContentCheckResponse_INVALID       ContentCheckResponse_Status = 2
)

var ContentCheckResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "VALID",
	2:  "INVALID",
}

var ContentCheckResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"VALID":         1,
	"INVALID":       2,
}

func (x ContentCheckResponse_Status) String() string {
	return proto.EnumName(ContentCheckResponse_Status_name, int32(x))
}

func (ContentCheckResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{10, 0}
}

type ContentCreateRequest struct {
	ContentID            string                    `protobuf:"bytes,1,opt,name=contentID,proto3" json:"contentID,omitempty"`
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
	return fileDescriptor_61cc9617ce0cf609, []int{0}
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

func (m *ContentCreateRequest) GetContentID() string {
	if m != nil {
		return m.ContentID
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
	ContentID            string                       `protobuf:"bytes,2,opt,name=contentID,proto3" json:"contentID,omitempty"`
	ContentToken         string                       `protobuf:"bytes,3,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	FileID               string                       `protobuf:"bytes,4,opt,name=fileID,proto3" json:"fileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ContentCreateResponse) Reset()         { *m = ContentCreateResponse{} }
func (m *ContentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ContentCreateResponse) ProtoMessage()    {}
func (*ContentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{1}
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

func (m *ContentCreateResponse) GetContentID() string {
	if m != nil {
		return m.ContentID
	}
	return ""
}

func (m *ContentCreateResponse) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func (m *ContentCreateResponse) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

type ContentUpdateRequest struct {
	ContentID            string                    `protobuf:"bytes,1,opt,name=contentID,proto3" json:"contentID,omitempty"`
	ContentToken         string                    `protobuf:"bytes,2,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	FileID               string                    `protobuf:"bytes,3,opt,name=fileID,proto3" json:"fileID,omitempty"`
	Content              []byte                    `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Type                 ContentUpdateRequest_Type `protobuf:"varint,5,opt,name=type,proto3,enum=ContentUpdateRequest_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ContentUpdateRequest) Reset()         { *m = ContentUpdateRequest{} }
func (m *ContentUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ContentUpdateRequest) ProtoMessage()    {}
func (*ContentUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{2}
}

func (m *ContentUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentUpdateRequest.Unmarshal(m, b)
}
func (m *ContentUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentUpdateRequest.Marshal(b, m, deterministic)
}
func (m *ContentUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentUpdateRequest.Merge(m, src)
}
func (m *ContentUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ContentUpdateRequest.Size(m)
}
func (m *ContentUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentUpdateRequest proto.InternalMessageInfo

func (m *ContentUpdateRequest) GetContentID() string {
	if m != nil {
		return m.ContentID
	}
	return ""
}

func (m *ContentUpdateRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func (m *ContentUpdateRequest) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *ContentUpdateRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ContentUpdateRequest) GetType() ContentUpdateRequest_Type {
	if m != nil {
		return m.Type
	}
	return ContentUpdateRequest_UNKNOWN
}

type ContentUpdateResponse struct {
	Status               ContentUpdateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ContentUpdateResponse_Status" json:"status,omitempty"`
	FileID               string                       `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ContentUpdateResponse) Reset()         { *m = ContentUpdateResponse{} }
func (m *ContentUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*ContentUpdateResponse) ProtoMessage()    {}
func (*ContentUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{3}
}

func (m *ContentUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentUpdateResponse.Unmarshal(m, b)
}
func (m *ContentUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentUpdateResponse.Marshal(b, m, deterministic)
}
func (m *ContentUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentUpdateResponse.Merge(m, src)
}
func (m *ContentUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_ContentUpdateResponse.Size(m)
}
func (m *ContentUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentUpdateResponse proto.InternalMessageInfo

func (m *ContentUpdateResponse) GetStatus() ContentUpdateResponse_Status {
	if m != nil {
		return m.Status
	}
	return ContentUpdateResponse_UNKNOWN
}

func (m *ContentUpdateResponse) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

type ContentDeleteRequest struct {
	ContentID            string   `protobuf:"bytes,1,opt,name=contentID,proto3" json:"contentID,omitempty"`
	ContentToken         string   `protobuf:"bytes,2,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentDeleteRequest) Reset()         { *m = ContentDeleteRequest{} }
func (m *ContentDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ContentDeleteRequest) ProtoMessage()    {}
func (*ContentDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{4}
}

func (m *ContentDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentDeleteRequest.Unmarshal(m, b)
}
func (m *ContentDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentDeleteRequest.Marshal(b, m, deterministic)
}
func (m *ContentDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentDeleteRequest.Merge(m, src)
}
func (m *ContentDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ContentDeleteRequest.Size(m)
}
func (m *ContentDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentDeleteRequest proto.InternalMessageInfo

func (m *ContentDeleteRequest) GetContentID() string {
	if m != nil {
		return m.ContentID
	}
	return ""
}

func (m *ContentDeleteRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

type ContentDeleteResponse struct {
	Status               ContentDeleteResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ContentDeleteResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ContentDeleteResponse) Reset()         { *m = ContentDeleteResponse{} }
func (m *ContentDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ContentDeleteResponse) ProtoMessage()    {}
func (*ContentDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{5}
}

func (m *ContentDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentDeleteResponse.Unmarshal(m, b)
}
func (m *ContentDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentDeleteResponse.Marshal(b, m, deterministic)
}
func (m *ContentDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentDeleteResponse.Merge(m, src)
}
func (m *ContentDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ContentDeleteResponse.Size(m)
}
func (m *ContentDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentDeleteResponse proto.InternalMessageInfo

func (m *ContentDeleteResponse) GetStatus() ContentDeleteResponse_Status {
	if m != nil {
		return m.Status
	}
	return ContentDeleteResponse_UNKNOWN
}

type ContentMsg struct {
	FileID               string          `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	Type                 ContentMsg_Type `protobuf:"varint,2,opt,name=type,proto3,enum=ContentMsg_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ContentMsg) Reset()         { *m = ContentMsg{} }
func (m *ContentMsg) String() string { return proto.CompactTextString(m) }
func (*ContentMsg) ProtoMessage()    {}
func (*ContentMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{6}
}

func (m *ContentMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentMsg.Unmarshal(m, b)
}
func (m *ContentMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentMsg.Marshal(b, m, deterministic)
}
func (m *ContentMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentMsg.Merge(m, src)
}
func (m *ContentMsg) XXX_Size() int {
	return xxx_messageInfo_ContentMsg.Size(m)
}
func (m *ContentMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ContentMsg proto.InternalMessageInfo

func (m *ContentMsg) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *ContentMsg) GetType() ContentMsg_Type {
	if m != nil {
		return m.Type
	}
	return ContentMsg_UNKNOWN
}

type ContentQueryRequest struct {
	ContentID            string   `protobuf:"bytes,1,opt,name=contentID,proto3" json:"contentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentQueryRequest) Reset()         { *m = ContentQueryRequest{} }
func (m *ContentQueryRequest) String() string { return proto.CompactTextString(m) }
func (*ContentQueryRequest) ProtoMessage()    {}
func (*ContentQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{7}
}

func (m *ContentQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentQueryRequest.Unmarshal(m, b)
}
func (m *ContentQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentQueryRequest.Marshal(b, m, deterministic)
}
func (m *ContentQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentQueryRequest.Merge(m, src)
}
func (m *ContentQueryRequest) XXX_Size() int {
	return xxx_messageInfo_ContentQueryRequest.Size(m)
}
func (m *ContentQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentQueryRequest proto.InternalMessageInfo

func (m *ContentQueryRequest) GetContentID() string {
	if m != nil {
		return m.ContentID
	}
	return ""
}

type ContentQueryResponse struct {
	Status               ContentQueryResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ContentQueryResponse_Status" json:"status,omitempty"`
	Files                []*ContentMsg               `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ContentQueryResponse) Reset()         { *m = ContentQueryResponse{} }
func (m *ContentQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ContentQueryResponse) ProtoMessage()    {}
func (*ContentQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{8}
}

func (m *ContentQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentQueryResponse.Unmarshal(m, b)
}
func (m *ContentQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentQueryResponse.Marshal(b, m, deterministic)
}
func (m *ContentQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentQueryResponse.Merge(m, src)
}
func (m *ContentQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ContentQueryResponse.Size(m)
}
func (m *ContentQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentQueryResponse proto.InternalMessageInfo

func (m *ContentQueryResponse) GetStatus() ContentQueryResponse_Status {
	if m != nil {
		return m.Status
	}
	return ContentQueryResponse_UNKNOWN
}

func (m *ContentQueryResponse) GetFiles() []*ContentMsg {
	if m != nil {
		return m.Files
	}
	return nil
}

type ContentCheckRequest struct {
	ContentID            string   `protobuf:"bytes,1,opt,name=contentID,proto3" json:"contentID,omitempty"`
	ContentToken         string   `protobuf:"bytes,2,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentCheckRequest) Reset()         { *m = ContentCheckRequest{} }
func (m *ContentCheckRequest) String() string { return proto.CompactTextString(m) }
func (*ContentCheckRequest) ProtoMessage()    {}
func (*ContentCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{9}
}

func (m *ContentCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentCheckRequest.Unmarshal(m, b)
}
func (m *ContentCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentCheckRequest.Marshal(b, m, deterministic)
}
func (m *ContentCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentCheckRequest.Merge(m, src)
}
func (m *ContentCheckRequest) XXX_Size() int {
	return xxx_messageInfo_ContentCheckRequest.Size(m)
}
func (m *ContentCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentCheckRequest proto.InternalMessageInfo

func (m *ContentCheckRequest) GetContentID() string {
	if m != nil {
		return m.ContentID
	}
	return ""
}

func (m *ContentCheckRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

type ContentCheckResponse struct {
	Status               ContentCheckResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ContentCheckResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ContentCheckResponse) Reset()         { *m = ContentCheckResponse{} }
func (m *ContentCheckResponse) String() string { return proto.CompactTextString(m) }
func (*ContentCheckResponse) ProtoMessage()    {}
func (*ContentCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61cc9617ce0cf609, []int{10}
}

func (m *ContentCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentCheckResponse.Unmarshal(m, b)
}
func (m *ContentCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentCheckResponse.Marshal(b, m, deterministic)
}
func (m *ContentCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentCheckResponse.Merge(m, src)
}
func (m *ContentCheckResponse) XXX_Size() int {
	return xxx_messageInfo_ContentCheckResponse.Size(m)
}
func (m *ContentCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentCheckResponse proto.InternalMessageInfo

func (m *ContentCheckResponse) GetStatus() ContentCheckResponse_Status {
	if m != nil {
		return m.Status
	}
	return ContentCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("ContentCreateRequest_Type", ContentCreateRequest_Type_name, ContentCreateRequest_Type_value)
	proto.RegisterEnum("ContentCreateResponse_Status", ContentCreateResponse_Status_name, ContentCreateResponse_Status_value)
	proto.RegisterEnum("ContentUpdateRequest_Type", ContentUpdateRequest_Type_name, ContentUpdateRequest_Type_value)
	proto.RegisterEnum("ContentUpdateResponse_Status", ContentUpdateResponse_Status_name, ContentUpdateResponse_Status_value)
	proto.RegisterEnum("ContentDeleteResponse_Status", ContentDeleteResponse_Status_name, ContentDeleteResponse_Status_value)
	proto.RegisterEnum("ContentMsg_Type", ContentMsg_Type_name, ContentMsg_Type_value)
	proto.RegisterEnum("ContentQueryResponse_Status", ContentQueryResponse_Status_name, ContentQueryResponse_Status_value)
	proto.RegisterEnum("ContentCheckResponse_Status", ContentCheckResponse_Status_name, ContentCheckResponse_Status_value)
	proto.RegisterType((*ContentCreateRequest)(nil), "ContentCreateRequest")
	proto.RegisterType((*ContentCreateResponse)(nil), "ContentCreateResponse")
	proto.RegisterType((*ContentUpdateRequest)(nil), "ContentUpdateRequest")
	proto.RegisterType((*ContentUpdateResponse)(nil), "ContentUpdateResponse")
	proto.RegisterType((*ContentDeleteRequest)(nil), "ContentDeleteRequest")
	proto.RegisterType((*ContentDeleteResponse)(nil), "ContentDeleteResponse")
	proto.RegisterType((*ContentMsg)(nil), "ContentMsg")
	proto.RegisterType((*ContentQueryRequest)(nil), "ContentQueryRequest")
	proto.RegisterType((*ContentQueryResponse)(nil), "ContentQueryResponse")
	proto.RegisterType((*ContentCheckRequest)(nil), "ContentCheckRequest")
	proto.RegisterType((*ContentCheckResponse)(nil), "ContentCheckResponse")
}

func init() { proto.RegisterFile("content.proto", fileDescriptor_61cc9617ce0cf609) }

var fileDescriptor_61cc9617ce0cf609 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x27, 0x4e, 0x93, 0xa9, 0x5f, 0xdb, 0xc9, 0x98, 0x75, 0x8a, 0xaa, 0x21, 0x15, 0x8b, 0x43,
	0x25, 0xa4, 0x1c, 0xba, 0xa1, 0x9d, 0xab, 0x24, 0x93, 0xa2, 0xad, 0x69, 0x97, 0xb6, 0x1b, 0x9c,
	0xa6, 0x51, 0xcc, 0x1f, 0x6d, 0x6a, 0x4b, 0x93, 0x0a, 0xf5, 0xc2, 0x1b, 0x70, 0xe6, 0x0a, 0x2f,
	0xc2, 0x2b, 0xf0, 0x16, 0xf0, 0x18, 0xa0, 0x38, 0x6e, 0x1d, 0x87, 0xa0, 0xc2, 0x68, 0x6f, 0x9f,
	0xf3, 0xfb, 0x6c, 0xff, 0xfe, 0xd8, 0x2e, 0xd4, 0xc6, 0xd3, 0x49, 0xcc, 0x26, 0xb1, 0x3d, 0x9b,
	0x4f, 0xe3, 0x29, 0xfd, 0xa6, 0xc1, 0x9e, 0x93, 0x8e, 0x38, 0x73, 0x76, 0x1d, 0xb3, 0x90, 0xbd,
	0x5b, 0xb0, 0x28, 0x26, 0x07, 0x50, 0x16, 0x48, 0xdf, 0xb5, 0xb4, 0xa6, 0xd6, 0x2a, 0x87, 0x72,
	0x80, 0x50, 0xa8, 0x8a, 0x62, 0x38, 0xbd, 0x61, 0x13, 0x0b, 0x71, 0x80, 0x32, 0x46, 0x2c, 0xd8,
	0x11, 0xb5, 0xa5, 0x37, 0xb5, 0x56, 0x35, 0x5c, 0x95, 0xc4, 0x86, 0x52, 0xbc, 0x9c, 0x31, 0xab,
	0xd4, 0xd4, 0x5a, 0xbb, 0xed, 0x86, 0x5d, 0xb4, 0x01, 0x7b, 0xb8, 0x9c, 0xb1, 0x90, 0xe3, 0xe8,
	0x13, 0x28, 0x25, 0x15, 0xa9, 0xc0, 0xce, 0x28, 0x38, 0x0d, 0x7a, 0x97, 0x01, 0xbe, 0x97, 0x14,
	0x7d, 0xdf, 0x19, 0x8e, 0x42, 0x0f, 0x6b, 0xa4, 0x0c, 0xc6, 0x85, 0xef, 0x7a, 0x3d, 0x8c, 0xe8,
	0x47, 0x04, 0xf5, 0xdc, 0x84, 0xd1, 0x6c, 0x3a, 0x89, 0x18, 0x79, 0x0a, 0x66, 0x14, 0x5f, 0xc7,
	0x8b, 0x88, 0xf3, 0xd9, 0x6d, 0x3f, 0xb4, 0x0b, 0x71, 0xf6, 0x80, 0x83, 0x42, 0x01, 0x56, 0x95,
	0x40, 0x9b, 0x94, 0xd0, 0x0b, 0x94, 0xd8, 0x07, 0xf3, 0xd5, 0xdb, 0x5b, 0xe6, 0xbb, 0x9c, 0x71,
	0x39, 0x14, 0x15, 0x1d, 0x83, 0x99, 0xae, 0xa5, 0x32, 0x6b, 0x40, 0xcd, 0x0f, 0x2e, 0x3a, 0x67,
	0xbe, 0x7b, 0xd5, 0xef, 0x84, 0x9d, 0x2e, 0xfe, 0xb9, 0xfa, 0x69, 0x09, 0x70, 0x30, 0x72, 0x1c,
	0x6f, 0x30, 0xc0, 0x1a, 0xb9, 0x2f, 0x81, 0xc3, 0xde, 0xa9, 0x17, 0x60, 0x44, 0x30, 0x54, 0xd7,
	0x43, 0xcf, 0xfb, 0x1e, 0xd6, 0xe9, 0x0f, 0xe9, 0xf0, 0x68, 0xf6, 0x72, 0xab, 0x0e, 0x4b, 0x5e,
	0x7a, 0x96, 0x57, 0xd6, 0xf9, 0x52, 0xb1, 0xf3, 0x86, 0xea, 0xbc, 0xb2, 0xb1, 0x3b, 0x3b, 0xff,
	0x5d, 0x5b, 0x3b, 0xbf, 0x9a, 0x70, 0x93, 0xf3, 0x2a, 0x2e, 0xef, 0xbc, 0xe4, 0x87, 0x14, 0xdf,
	0x3e, 0x6c, 0xdf, 0xb7, 0x1a, 0x94, 0x83, 0xde, 0xf0, 0xea, 0xa4, 0x37, 0x0a, 0x5c, 0xac, 0x13,
	0x00, 0xf3, 0xa4, 0xe3, 0x9f, 0x79, 0x2e, 0x2e, 0xfd, 0x66, 0xa9, 0x41, 0x9f, 0xad, 0x1d, 0x75,
	0xd9, 0x2d, 0xdb, 0xa2, 0xa3, 0xf4, 0x8b, 0x94, 0x70, 0x35, 0xf5, 0x26, 0x09, 0x55, 0x5c, 0x4e,
	0x42, 0x7a, 0xbe, 0x75, 0xa9, 0xe8, 0x7b, 0x00, 0xb1, 0x74, 0x37, 0x7a, 0x9d, 0xf1, 0x48, 0x53,
	0x32, 0xf8, 0x58, 0x24, 0x0d, 0xf1, 0xdd, 0x62, 0x5b, 0xb6, 0xdc, 0x39, 0x5f, 0x87, 0xf0, 0x40,
	0xcc, 0x72, 0xbe, 0x60, 0xf3, 0xe5, 0x5f, 0xa9, 0x4e, 0xbf, 0xca, 0xe3, 0x27, 0xba, 0x84, 0xa0,
	0x47, 0x39, 0x41, 0x0f, 0xec, 0x22, 0x58, 0x3e, 0x92, 0x8f, 0xc0, 0x48, 0x08, 0x46, 0x16, 0x6a,
	0xea, 0xad, 0x4a, 0xbb, 0x92, 0xe1, 0x15, 0xa6, 0x5f, 0x68, 0xf7, 0x3f, 0x25, 0x57, 0xa2, 0x88,
	0xe8, 0xe5, 0x9a, 0xb5, 0xf3, 0x86, 0x8d, 0x6f, 0xb6, 0x97, 0xb5, 0x4f, 0x99, 0xa7, 0x27, 0x9d,
	0x79, 0x93, 0x32, 0x0a, 0x2c, 0x9f, 0x34, 0xff, 0xdf, 0x69, 0x27, 0xde, 0x26, 0x5f, 0x30, 0x57,
	0x40, 0xc0, 0x30, 0x6a, 0x7f, 0x46, 0xb0, 0x23, 0x96, 0x24, 0xc7, 0x60, 0xa6, 0xcf, 0x03, 0xa9,
	0x17, 0xbe, 0x53, 0x8d, 0xfd, 0xe2, 0x57, 0x24, 0x69, 0x4c, 0x6f, 0x17, 0xd9, 0xa8, 0x5c, 0x73,
	0xb2, 0x31, 0x77, 0x59, 0x1d, 0x83, 0x99, 0x9e, 0x29, 0xd9, 0xa8, 0x1c, 0x73, 0xd9, 0x98, 0x3b,
	0xa2, 0x47, 0x60, 0xf0, 0xec, 0x90, 0x3d, 0xbb, 0x20, 0xa7, 0x8d, 0xba, 0xfd, 0x87, 0x1c, 0x1a,
	0x5c, 0x57, 0xd9, 0x95, 0xf5, 0x59, 0x76, 0x29, 0xe2, 0xbf, 0x30, 0xf9, 0xdf, 0x87, 0xc3, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x31, 0xdb, 0x36, 0x4f, 0x08, 0x00, 0x00,
}