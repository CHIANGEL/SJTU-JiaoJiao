// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package transaction

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

type TransStatus int32

const (
	TransStatus_UNKNOWN  TransStatus = 0
	TransStatus_ASKING   TransStatus = 1
	TransStatus_ACCEPTED TransStatus = 2
	TransStatus_REJECTED TransStatus = 3
	TransStatus_CLOSED   TransStatus = 4
	TransStatus_PENDING  TransStatus = 5
	TransStatus_DONE     TransStatus = 6
)

var TransStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "ASKING",
	2: "ACCEPTED",
	3: "REJECTED",
	4: "CLOSED",
	5: "PENDING",
	6: "DONE",
}

var TransStatus_value = map[string]int32{
	"UNKNOWN":  0,
	"ASKING":   1,
	"ACCEPTED": 2,
	"REJECTED": 3,
	"CLOSED":   4,
	"PENDING":  5,
	"DONE":     6,
}

func (x TransStatus) String() string {
	return proto.EnumName(TransStatus_name, int32(x))
}

func (TransStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

type TransactionCreateRequest_Category int32

const (
	TransactionCreateRequest_UNKNOWN TransactionCreateRequest_Category = 0
	TransactionCreateRequest_SELL    TransactionCreateRequest_Category = 1
	TransactionCreateRequest_BUY     TransactionCreateRequest_Category = 2
)

var TransactionCreateRequest_Category_name = map[int32]string{
	0: "UNKNOWN",
	1: "SELL",
	2: "BUY",
}

var TransactionCreateRequest_Category_value = map[string]int32{
	"UNKNOWN": 0,
	"SELL":    1,
	"BUY":     2,
}

func (x TransactionCreateRequest_Category) String() string {
	return proto.EnumName(TransactionCreateRequest_Category_name, int32(x))
}

func (TransactionCreateRequest_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0, 0}
}

type TransactionCreateResponse_Status int32

const (
	TransactionCreateResponse_UNKNOWN       TransactionCreateResponse_Status = 0
	TransactionCreateResponse_INVALID_PARAM TransactionCreateResponse_Status = -1
	TransactionCreateResponse_SUCCESS       TransactionCreateResponse_Status = 1
	TransactionCreateResponse_NOT_FOUND     TransactionCreateResponse_Status = 2
)

var TransactionCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "NOT_FOUND",
}

var TransactionCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"NOT_FOUND":     2,
}

func (x TransactionCreateResponse_Status) String() string {
	return proto.EnumName(TransactionCreateResponse_Status_name, int32(x))
}

func (TransactionCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1, 0}
}

type TransactionUpdateResponse_Status int32

const (
	TransactionUpdateResponse_UNKNOWN       TransactionUpdateResponse_Status = 0
	TransactionUpdateResponse_INVALID_PARAM TransactionUpdateResponse_Status = -1
	TransactionUpdateResponse_SUCCESS       TransactionUpdateResponse_Status = 1
	TransactionUpdateResponse_NOT_FOUND     TransactionUpdateResponse_Status = 2
)

var TransactionUpdateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "NOT_FOUND",
}

var TransactionUpdateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"NOT_FOUND":     2,
}

func (x TransactionUpdateResponse_Status) String() string {
	return proto.EnumName(TransactionUpdateResponse_Status_name, int32(x))
}

func (TransactionUpdateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{3, 0}
}

type TransactionMsg_Category int32

const (
	TransactionMsg_CATEGORY_UNKNOWN TransactionMsg_Category = 0
	TransactionMsg_SELL             TransactionMsg_Category = 1
	TransactionMsg_BUY              TransactionMsg_Category = 2
)

var TransactionMsg_Category_name = map[int32]string{
	0: "CATEGORY_UNKNOWN",
	1: "SELL",
	2: "BUY",
}

var TransactionMsg_Category_value = map[string]int32{
	"CATEGORY_UNKNOWN": 0,
	"SELL":             1,
	"BUY":              2,
}

func (x TransactionMsg_Category) String() string {
	return proto.EnumName(TransactionMsg_Category_name, int32(x))
}

func (TransactionMsg_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{4, 0}
}

type TransactionFindRequest_Category int32

const (
	TransactionFindRequest_CATEGORY_UNKNOWN TransactionFindRequest_Category = 0
	TransactionFindRequest_SELL             TransactionFindRequest_Category = 1
	TransactionFindRequest_BUY              TransactionFindRequest_Category = 2
)

var TransactionFindRequest_Category_name = map[int32]string{
	0: "CATEGORY_UNKNOWN",
	1: "SELL",
	2: "BUY",
}

var TransactionFindRequest_Category_value = map[string]int32{
	"CATEGORY_UNKNOWN": 0,
	"SELL":             1,
	"BUY":              2,
}

func (x TransactionFindRequest_Category) String() string {
	return proto.EnumName(TransactionFindRequest_Category_name, int32(x))
}

func (TransactionFindRequest_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{5, 0}
}

type TransactionFindResponse_Status int32

const (
	TransactionFindResponse_UNKNOWN       TransactionFindResponse_Status = 0
	TransactionFindResponse_INVALID_PARAM TransactionFindResponse_Status = -1
	TransactionFindResponse_SUCCESS       TransactionFindResponse_Status = 1
	TransactionFindResponse_NOT_FOUND     TransactionFindResponse_Status = 2
)

var TransactionFindResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "NOT_FOUND",
}

var TransactionFindResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"NOT_FOUND":     2,
}

func (x TransactionFindResponse_Status) String() string {
	return proto.EnumName(TransactionFindResponse_Status_name, int32(x))
}

func (TransactionFindResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{6, 0}
}

type TransactionCreateRequest struct {
	InfoID               int32                             `protobuf:"varint,1,opt,name=infoID,proto3" json:"infoID,omitempty"`
	Category             TransactionCreateRequest_Category `protobuf:"varint,2,opt,name=category,proto3,enum=TransactionCreateRequest_Category" json:"category,omitempty"`
	FromUserID           int32                             `protobuf:"varint,3,opt,name=fromUserID,proto3" json:"fromUserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TransactionCreateRequest) Reset()         { *m = TransactionCreateRequest{} }
func (m *TransactionCreateRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionCreateRequest) ProtoMessage()    {}
func (*TransactionCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *TransactionCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionCreateRequest.Unmarshal(m, b)
}
func (m *TransactionCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionCreateRequest.Marshal(b, m, deterministic)
}
func (m *TransactionCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionCreateRequest.Merge(m, src)
}
func (m *TransactionCreateRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionCreateRequest.Size(m)
}
func (m *TransactionCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionCreateRequest proto.InternalMessageInfo

func (m *TransactionCreateRequest) GetInfoID() int32 {
	if m != nil {
		return m.InfoID
	}
	return 0
}

func (m *TransactionCreateRequest) GetCategory() TransactionCreateRequest_Category {
	if m != nil {
		return m.Category
	}
	return TransactionCreateRequest_UNKNOWN
}

func (m *TransactionCreateRequest) GetFromUserID() int32 {
	if m != nil {
		return m.FromUserID
	}
	return 0
}

type TransactionCreateResponse struct {
	Status               TransactionCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=TransactionCreateResponse_Status" json:"status,omitempty"`
	TransactionID        int32                            `protobuf:"varint,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *TransactionCreateResponse) Reset()         { *m = TransactionCreateResponse{} }
func (m *TransactionCreateResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionCreateResponse) ProtoMessage()    {}
func (*TransactionCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *TransactionCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionCreateResponse.Unmarshal(m, b)
}
func (m *TransactionCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionCreateResponse.Marshal(b, m, deterministic)
}
func (m *TransactionCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionCreateResponse.Merge(m, src)
}
func (m *TransactionCreateResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionCreateResponse.Size(m)
}
func (m *TransactionCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionCreateResponse proto.InternalMessageInfo

func (m *TransactionCreateResponse) GetStatus() TransactionCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return TransactionCreateResponse_UNKNOWN
}

func (m *TransactionCreateResponse) GetTransactionID() int32 {
	if m != nil {
		return m.TransactionID
	}
	return 0
}

type TransactionUpdateRequest struct {
	TransactionID        int32       `protobuf:"varint,1,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	Status               TransStatus `protobuf:"varint,2,opt,name=status,proto3,enum=TransStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TransactionUpdateRequest) Reset()         { *m = TransactionUpdateRequest{} }
func (m *TransactionUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionUpdateRequest) ProtoMessage()    {}
func (*TransactionUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
}

func (m *TransactionUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionUpdateRequest.Unmarshal(m, b)
}
func (m *TransactionUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionUpdateRequest.Marshal(b, m, deterministic)
}
func (m *TransactionUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionUpdateRequest.Merge(m, src)
}
func (m *TransactionUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionUpdateRequest.Size(m)
}
func (m *TransactionUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionUpdateRequest proto.InternalMessageInfo

func (m *TransactionUpdateRequest) GetTransactionID() int32 {
	if m != nil {
		return m.TransactionID
	}
	return 0
}

func (m *TransactionUpdateRequest) GetStatus() TransStatus {
	if m != nil {
		return m.Status
	}
	return TransStatus_UNKNOWN
}

type TransactionUpdateResponse struct {
	Status               TransactionUpdateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=TransactionUpdateResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *TransactionUpdateResponse) Reset()         { *m = TransactionUpdateResponse{} }
func (m *TransactionUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionUpdateResponse) ProtoMessage()    {}
func (*TransactionUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{3}
}

func (m *TransactionUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionUpdateResponse.Unmarshal(m, b)
}
func (m *TransactionUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionUpdateResponse.Marshal(b, m, deterministic)
}
func (m *TransactionUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionUpdateResponse.Merge(m, src)
}
func (m *TransactionUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionUpdateResponse.Size(m)
}
func (m *TransactionUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionUpdateResponse proto.InternalMessageInfo

func (m *TransactionUpdateResponse) GetStatus() TransactionUpdateResponse_Status {
	if m != nil {
		return m.Status
	}
	return TransactionUpdateResponse_UNKNOWN
}

type TransactionMsg struct {
	TransactionID        int32                   `protobuf:"varint,1,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	InfoID               int32                   `protobuf:"varint,2,opt,name=infoID,proto3" json:"infoID,omitempty"`
	Category             TransactionMsg_Category `protobuf:"varint,3,opt,name=category,proto3,enum=TransactionMsg_Category" json:"category,omitempty"`
	FromUserID           int32                   `protobuf:"varint,4,opt,name=fromUserID,proto3" json:"fromUserID,omitempty"`
	ToUserID             int32                   `protobuf:"varint,5,opt,name=toUserID,proto3" json:"toUserID,omitempty"`
	CreateTime           int64                   `protobuf:"varint,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Status               TransStatus             `protobuf:"varint,7,opt,name=status,proto3,enum=TransStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TransactionMsg) Reset()         { *m = TransactionMsg{} }
func (m *TransactionMsg) String() string { return proto.CompactTextString(m) }
func (*TransactionMsg) ProtoMessage()    {}
func (*TransactionMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{4}
}

func (m *TransactionMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionMsg.Unmarshal(m, b)
}
func (m *TransactionMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionMsg.Marshal(b, m, deterministic)
}
func (m *TransactionMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionMsg.Merge(m, src)
}
func (m *TransactionMsg) XXX_Size() int {
	return xxx_messageInfo_TransactionMsg.Size(m)
}
func (m *TransactionMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionMsg proto.InternalMessageInfo

func (m *TransactionMsg) GetTransactionID() int32 {
	if m != nil {
		return m.TransactionID
	}
	return 0
}

func (m *TransactionMsg) GetInfoID() int32 {
	if m != nil {
		return m.InfoID
	}
	return 0
}

func (m *TransactionMsg) GetCategory() TransactionMsg_Category {
	if m != nil {
		return m.Category
	}
	return TransactionMsg_CATEGORY_UNKNOWN
}

func (m *TransactionMsg) GetFromUserID() int32 {
	if m != nil {
		return m.FromUserID
	}
	return 0
}

func (m *TransactionMsg) GetToUserID() int32 {
	if m != nil {
		return m.ToUserID
	}
	return 0
}

func (m *TransactionMsg) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *TransactionMsg) GetStatus() TransStatus {
	if m != nil {
		return m.Status
	}
	return TransStatus_UNKNOWN
}

type TransactionFindRequest struct {
	InfoID               int32                           `protobuf:"varint,1,opt,name=infoID,proto3" json:"infoID,omitempty"`
	Category             TransactionFindRequest_Category `protobuf:"varint,2,opt,name=category,proto3,enum=TransactionFindRequest_Category" json:"category,omitempty"`
	UserID               int32                           `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
	LowCreateTime        int64                           `protobuf:"varint,4,opt,name=lowCreateTime,proto3" json:"lowCreateTime,omitempty"`
	HighCreateTime       int64                           `protobuf:"varint,5,opt,name=highCreateTime,proto3" json:"highCreateTime,omitempty"`
	Status               TransStatus                     `protobuf:"varint,6,opt,name=status,proto3,enum=TransStatus" json:"status,omitempty"`
	Limit                uint32                          `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32                          `protobuf:"varint,8,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TransactionFindRequest) Reset()         { *m = TransactionFindRequest{} }
func (m *TransactionFindRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionFindRequest) ProtoMessage()    {}
func (*TransactionFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{5}
}

func (m *TransactionFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionFindRequest.Unmarshal(m, b)
}
func (m *TransactionFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionFindRequest.Marshal(b, m, deterministic)
}
func (m *TransactionFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionFindRequest.Merge(m, src)
}
func (m *TransactionFindRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionFindRequest.Size(m)
}
func (m *TransactionFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionFindRequest proto.InternalMessageInfo

func (m *TransactionFindRequest) GetInfoID() int32 {
	if m != nil {
		return m.InfoID
	}
	return 0
}

func (m *TransactionFindRequest) GetCategory() TransactionFindRequest_Category {
	if m != nil {
		return m.Category
	}
	return TransactionFindRequest_CATEGORY_UNKNOWN
}

func (m *TransactionFindRequest) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *TransactionFindRequest) GetLowCreateTime() int64 {
	if m != nil {
		return m.LowCreateTime
	}
	return 0
}

func (m *TransactionFindRequest) GetHighCreateTime() int64 {
	if m != nil {
		return m.HighCreateTime
	}
	return 0
}

func (m *TransactionFindRequest) GetStatus() TransStatus {
	if m != nil {
		return m.Status
	}
	return TransStatus_UNKNOWN
}

func (m *TransactionFindRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TransactionFindRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type TransactionFindResponse struct {
	Status               TransactionFindResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=TransactionFindResponse_Status" json:"status,omitempty"`
	Transactions         []*TransactionMsg              `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TransactionFindResponse) Reset()         { *m = TransactionFindResponse{} }
func (m *TransactionFindResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionFindResponse) ProtoMessage()    {}
func (*TransactionFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{6}
}

func (m *TransactionFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionFindResponse.Unmarshal(m, b)
}
func (m *TransactionFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionFindResponse.Marshal(b, m, deterministic)
}
func (m *TransactionFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionFindResponse.Merge(m, src)
}
func (m *TransactionFindResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionFindResponse.Size(m)
}
func (m *TransactionFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionFindResponse proto.InternalMessageInfo

func (m *TransactionFindResponse) GetStatus() TransactionFindResponse_Status {
	if m != nil {
		return m.Status
	}
	return TransactionFindResponse_UNKNOWN
}

func (m *TransactionFindResponse) GetTransactions() []*TransactionMsg {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterEnum("TransStatus", TransStatus_name, TransStatus_value)
	proto.RegisterEnum("TransactionCreateRequest_Category", TransactionCreateRequest_Category_name, TransactionCreateRequest_Category_value)
	proto.RegisterEnum("TransactionCreateResponse_Status", TransactionCreateResponse_Status_name, TransactionCreateResponse_Status_value)
	proto.RegisterEnum("TransactionUpdateResponse_Status", TransactionUpdateResponse_Status_name, TransactionUpdateResponse_Status_value)
	proto.RegisterEnum("TransactionMsg_Category", TransactionMsg_Category_name, TransactionMsg_Category_value)
	proto.RegisterEnum("TransactionFindRequest_Category", TransactionFindRequest_Category_name, TransactionFindRequest_Category_value)
	proto.RegisterEnum("TransactionFindResponse_Status", TransactionFindResponse_Status_name, TransactionFindResponse_Status_value)
	proto.RegisterType((*TransactionCreateRequest)(nil), "TransactionCreateRequest")
	proto.RegisterType((*TransactionCreateResponse)(nil), "TransactionCreateResponse")
	proto.RegisterType((*TransactionUpdateRequest)(nil), "TransactionUpdateRequest")
	proto.RegisterType((*TransactionUpdateResponse)(nil), "TransactionUpdateResponse")
	proto.RegisterType((*TransactionMsg)(nil), "TransactionMsg")
	proto.RegisterType((*TransactionFindRequest)(nil), "TransactionFindRequest")
	proto.RegisterType((*TransactionFindResponse)(nil), "TransactionFindResponse")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x41, 0x4f, 0xd4, 0x50,
	0x10, 0xf6, 0xb5, 0xdb, 0x52, 0x07, 0x76, 0xad, 0x2f, 0x04, 0xca, 0x1e, 0x74, 0x6d, 0x88, 0xd9,
	0x70, 0xe8, 0x61, 0x31, 0x31, 0x26, 0x46, 0x53, 0xdb, 0x42, 0x2a, 0x4b, 0x4b, 0xda, 0xad, 0x86,
	0x13, 0x59, 0xa1, 0x0b, 0x4d, 0x60, 0x8b, 0x6d, 0x89, 0xf1, 0x07, 0xf9, 0x2b, 0x4c, 0xfc, 0x01,
	0x7a, 0xf7, 0xe6, 0x6f, 0xd1, 0xbc, 0xd7, 0x02, 0xef, 0x75, 0xb7, 0x40, 0x0c, 0x7b, 0x9b, 0xd9,
	0x6f, 0x66, 0xde, 0x7c, 0x33, 0xdf, 0x14, 0x1e, 0x17, 0xd9, 0x78, 0x9a, 0x8f, 0x0f, 0x8b, 0x24,
	0x9d, 0x1a, 0xe7, 0x59, 0x5a, 0xa4, 0xfa, 0x0f, 0x04, 0xda, 0xe8, 0xda, 0x6b, 0x65, 0xf1, 0xb8,
	0x88, 0x83, 0xf8, 0xf3, 0x45, 0x9c, 0x17, 0x78, 0x05, 0xe4, 0x64, 0x3a, 0x49, 0x5d, 0x5b, 0x43,
	0x3d, 0xd4, 0x97, 0x82, 0xca, 0xc2, 0x6f, 0x40, 0x39, 0x1c, 0x17, 0xf1, 0x71, 0x9a, 0x7d, 0xd5,
	0x84, 0x1e, 0xea, 0x77, 0x06, 0xba, 0xd1, 0x94, 0xc4, 0xb0, 0x2a, 0x64, 0x70, 0x15, 0x83, 0x9f,
	0x00, 0x4c, 0xb2, 0xf4, 0x2c, 0xca, 0xe3, 0xcc, 0xb5, 0x35, 0x91, 0xe6, 0x66, 0x3c, 0xfa, 0x06,
	0x28, 0x97, 0x51, 0x78, 0x11, 0x16, 0x22, 0x6f, 0xc7, 0xf3, 0x3f, 0x7a, 0xea, 0x03, 0xac, 0x40,
	0x2b, 0x74, 0x86, 0x43, 0x15, 0xe1, 0x05, 0x10, 0xdf, 0x45, 0xfb, 0xaa, 0xa0, 0xff, 0x42, 0xb0,
	0x36, 0xa7, 0x76, 0x7e, 0x9e, 0x4e, 0xf3, 0x18, 0xbf, 0x02, 0x39, 0x2f, 0xc6, 0xc5, 0x45, 0x4e,
	0x3b, 0xe8, 0x0c, 0x9e, 0x19, 0x8d, 0x58, 0x23, 0xa4, 0xc0, 0xa0, 0x0a, 0xc0, 0xeb, 0xd0, 0x66,
	0xe8, 0x72, 0x6d, 0xda, 0xa9, 0x14, 0xf0, 0x4e, 0x7d, 0x17, 0xe4, 0x32, 0x8e, 0x7f, 0x68, 0x17,
	0xda, 0xae, 0xf7, 0xc1, 0x1c, 0xba, 0xf6, 0xc1, 0x9e, 0x19, 0x98, 0xbb, 0xea, 0xdf, 0xcb, 0x1f,
	0x22, 0xc0, 0x30, 0xb2, 0x2c, 0x27, 0x0c, 0x55, 0x84, 0xdb, 0xf0, 0xd0, 0xf3, 0x47, 0x07, 0x5b,
	0x7e, 0xe4, 0xd9, 0xaa, 0xa0, 0x4f, 0xb8, 0x69, 0x44, 0xe7, 0x47, 0xcc, 0x34, 0x66, 0x1e, 0x84,
	0xe6, 0x3c, 0x08, 0xaf, 0x5f, 0x75, 0x5c, 0x4e, 0x66, 0xa9, 0xec, 0x98, 0x6f, 0x4e, 0xff, 0xc6,
	0xb3, 0x76, 0x59, 0xe8, 0x2e, 0xac, 0xf1, 0xd8, 0x1a, 0x6b, 0xf7, 0xcd, 0xc7, 0x77, 0x01, 0x3a,
	0x4c, 0xed, 0xdd, 0xfc, 0xf8, 0x8e, 0x34, 0x5c, 0xaf, 0xae, 0xc0, 0xad, 0xee, 0x0b, 0x66, 0x75,
	0x45, 0xda, 0x9c, 0x66, 0xf0, 0x05, 0x6e, 0x5f, 0xd8, 0x56, 0x7d, 0x61, 0x71, 0x17, 0x94, 0x22,
	0xad, 0xfe, 0x95, 0xe8, 0xbf, 0x57, 0x36, 0x89, 0x3d, 0xa4, 0x8b, 0x36, 0x4a, 0xce, 0x62, 0x4d,
	0xee, 0xa1, 0xbe, 0x18, 0x30, 0x1e, 0x66, 0x60, 0x0b, 0x37, 0x0c, 0x6c, 0x93, 0x91, 0xc4, 0x32,
	0xa8, 0x96, 0x39, 0x72, 0xb6, 0xfd, 0x60, 0xff, 0xe0, 0x46, 0x6d, 0xfc, 0x11, 0x60, 0x85, 0x69,
	0x6e, 0x2b, 0x99, 0x1e, 0xdd, 0x26, 0xed, 0xd7, 0x33, 0xd2, 0xee, 0x19, 0xf3, 0x53, 0xcc, 0xe3,
	0x69, 0x05, 0xe4, 0x0b, 0x56, 0xd4, 0x95, 0x45, 0x66, 0x76, 0x9a, 0x7e, 0xb1, 0xae, 0x69, 0x68,
	0x51, 0x1a, 0x78, 0x27, 0x7e, 0x0e, 0x9d, 0x93, 0xe4, 0xf8, 0x84, 0x81, 0x49, 0x14, 0x56, 0xf3,
	0x32, 0x8c, 0xc9, 0xcd, 0x8c, 0xe1, 0x65, 0x90, 0x4e, 0x93, 0xb3, 0xa4, 0xa0, 0xb4, 0xb6, 0x83,
	0xd2, 0x20, 0x2f, 0x4c, 0x27, 0x93, 0x3c, 0x2e, 0x34, 0x85, 0xba, 0x2b, 0xeb, 0xff, 0xf8, 0xfd,
	0x8d, 0x60, 0x75, 0x86, 0x9c, 0x4a, 0x43, 0x2f, 0x6b, 0x1a, 0x7a, 0x6a, 0x34, 0x20, 0xeb, 0x77,
	0x67, 0x13, 0x96, 0x98, 0x55, 0x26, 0x32, 0x16, 0xfb, 0x8b, 0x83, 0x47, 0xb5, 0x2d, 0x0d, 0x38,
	0xd0, 0x3d, 0xcb, 0x6e, 0x23, 0x86, 0x45, 0x86, 0x52, 0x3e, 0x27, 0x80, 0x6c, 0x86, 0x3b, 0xae,
	0xb7, 0xad, 0x22, 0xbc, 0x04, 0x8a, 0x69, 0x59, 0xce, 0xde, 0xc8, 0xb1, 0x55, 0x81, 0x58, 0x81,
	0xf3, 0xde, 0xb1, 0x88, 0x25, 0x12, 0x9c, 0x35, 0xf4, 0x43, 0xc7, 0x56, 0x5b, 0x24, 0xc1, 0x9e,
	0xe3, 0xd9, 0x24, 0x48, 0x22, 0x44, 0xda, 0xbe, 0xe7, 0xa8, 0xf2, 0xe0, 0x27, 0xaa, 0xea, 0x94,
	0x6d, 0xe0, 0xb7, 0x20, 0x97, 0x63, 0xc6, 0x6b, 0x8d, 0xdf, 0x93, 0x6e, 0xb7, 0xf9, 0x84, 0x93,
	0x04, 0xe5, 0x79, 0xe2, 0x13, 0x70, 0x77, 0x94, 0x4f, 0x30, 0x73, 0xf9, 0x5a, 0x64, 0x36, 0x78,
	0xb5, 0x61, 0xe9, 0xbb, 0x5a, 0xd3, 0x18, 0x3f, 0xc9, 0xf4, 0x83, 0xba, 0xf9, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0x25, 0x89, 0x45, 0x65, 0x07, 0x00, 0x00,
}
