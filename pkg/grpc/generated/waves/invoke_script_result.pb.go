// Code generated by protoc-gen-go. DO NOT EDIT.
// source: waves/invoke_script_result.proto

package waves

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

type InvokeScriptResult struct {
	Data                 []*DataTransactionData_DataEntry `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Transfers            []*InvokeScriptResult_Payment    `protobuf:"bytes,2,rep,name=transfers,proto3" json:"transfers,omitempty"`
	Issues               []*InvokeScriptResult_Issue      `protobuf:"bytes,3,rep,name=issues,proto3" json:"issues,omitempty"`
	Reissues             []*InvokeScriptResult_Reissue    `protobuf:"bytes,4,rep,name=reissues,proto3" json:"reissues,omitempty"`
	Burns                []*InvokeScriptResult_Burn       `protobuf:"bytes,5,rep,name=burns,proto3" json:"burns,omitempty"`
	ErrorMessage         *InvokeScriptResult_ErrorMessage `protobuf:"bytes,6,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *InvokeScriptResult) Reset()         { *m = InvokeScriptResult{} }
func (m *InvokeScriptResult) String() string { return proto.CompactTextString(m) }
func (*InvokeScriptResult) ProtoMessage()    {}
func (*InvokeScriptResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7c55166e11cf76, []int{0}
}

func (m *InvokeScriptResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeScriptResult.Unmarshal(m, b)
}
func (m *InvokeScriptResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeScriptResult.Marshal(b, m, deterministic)
}
func (m *InvokeScriptResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeScriptResult.Merge(m, src)
}
func (m *InvokeScriptResult) XXX_Size() int {
	return xxx_messageInfo_InvokeScriptResult.Size(m)
}
func (m *InvokeScriptResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeScriptResult.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeScriptResult proto.InternalMessageInfo

func (m *InvokeScriptResult) GetData() []*DataTransactionData_DataEntry {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeScriptResult) GetTransfers() []*InvokeScriptResult_Payment {
	if m != nil {
		return m.Transfers
	}
	return nil
}

func (m *InvokeScriptResult) GetIssues() []*InvokeScriptResult_Issue {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *InvokeScriptResult) GetReissues() []*InvokeScriptResult_Reissue {
	if m != nil {
		return m.Reissues
	}
	return nil
}

func (m *InvokeScriptResult) GetBurns() []*InvokeScriptResult_Burn {
	if m != nil {
		return m.Burns
	}
	return nil
}

func (m *InvokeScriptResult) GetErrorMessage() *InvokeScriptResult_ErrorMessage {
	if m != nil {
		return m.ErrorMessage
	}
	return nil
}

type InvokeScriptResult_Payment struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount               *Amount  `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeScriptResult_Payment) Reset()         { *m = InvokeScriptResult_Payment{} }
func (m *InvokeScriptResult_Payment) String() string { return proto.CompactTextString(m) }
func (*InvokeScriptResult_Payment) ProtoMessage()    {}
func (*InvokeScriptResult_Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7c55166e11cf76, []int{0, 0}
}

func (m *InvokeScriptResult_Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeScriptResult_Payment.Unmarshal(m, b)
}
func (m *InvokeScriptResult_Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeScriptResult_Payment.Marshal(b, m, deterministic)
}
func (m *InvokeScriptResult_Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeScriptResult_Payment.Merge(m, src)
}
func (m *InvokeScriptResult_Payment) XXX_Size() int {
	return xxx_messageInfo_InvokeScriptResult_Payment.Size(m)
}
func (m *InvokeScriptResult_Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeScriptResult_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeScriptResult_Payment proto.InternalMessageInfo

func (m *InvokeScriptResult_Payment) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *InvokeScriptResult_Payment) GetAmount() *Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

type InvokeScriptResult_Issue struct {
	AssetId              []byte   `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount               int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Decimals             int32    `protobuf:"varint,5,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Reissuable           bool     `protobuf:"varint,6,opt,name=reissuable,proto3" json:"reissuable,omitempty"`
	Script               []byte   `protobuf:"bytes,7,opt,name=script,proto3" json:"script,omitempty"`
	Nonce                int64    `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeScriptResult_Issue) Reset()         { *m = InvokeScriptResult_Issue{} }
func (m *InvokeScriptResult_Issue) String() string { return proto.CompactTextString(m) }
func (*InvokeScriptResult_Issue) ProtoMessage()    {}
func (*InvokeScriptResult_Issue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7c55166e11cf76, []int{0, 1}
}

func (m *InvokeScriptResult_Issue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeScriptResult_Issue.Unmarshal(m, b)
}
func (m *InvokeScriptResult_Issue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeScriptResult_Issue.Marshal(b, m, deterministic)
}
func (m *InvokeScriptResult_Issue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeScriptResult_Issue.Merge(m, src)
}
func (m *InvokeScriptResult_Issue) XXX_Size() int {
	return xxx_messageInfo_InvokeScriptResult_Issue.Size(m)
}
func (m *InvokeScriptResult_Issue) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeScriptResult_Issue.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeScriptResult_Issue proto.InternalMessageInfo

func (m *InvokeScriptResult_Issue) GetAssetId() []byte {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *InvokeScriptResult_Issue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvokeScriptResult_Issue) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *InvokeScriptResult_Issue) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InvokeScriptResult_Issue) GetDecimals() int32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *InvokeScriptResult_Issue) GetReissuable() bool {
	if m != nil {
		return m.Reissuable
	}
	return false
}

func (m *InvokeScriptResult_Issue) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *InvokeScriptResult_Issue) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type InvokeScriptResult_Reissue struct {
	AssetId              []byte   `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	IsReissuable         bool     `protobuf:"varint,3,opt,name=is_reissuable,json=isReissuable,proto3" json:"is_reissuable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeScriptResult_Reissue) Reset()         { *m = InvokeScriptResult_Reissue{} }
func (m *InvokeScriptResult_Reissue) String() string { return proto.CompactTextString(m) }
func (*InvokeScriptResult_Reissue) ProtoMessage()    {}
func (*InvokeScriptResult_Reissue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7c55166e11cf76, []int{0, 2}
}

func (m *InvokeScriptResult_Reissue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeScriptResult_Reissue.Unmarshal(m, b)
}
func (m *InvokeScriptResult_Reissue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeScriptResult_Reissue.Marshal(b, m, deterministic)
}
func (m *InvokeScriptResult_Reissue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeScriptResult_Reissue.Merge(m, src)
}
func (m *InvokeScriptResult_Reissue) XXX_Size() int {
	return xxx_messageInfo_InvokeScriptResult_Reissue.Size(m)
}
func (m *InvokeScriptResult_Reissue) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeScriptResult_Reissue.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeScriptResult_Reissue proto.InternalMessageInfo

func (m *InvokeScriptResult_Reissue) GetAssetId() []byte {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *InvokeScriptResult_Reissue) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InvokeScriptResult_Reissue) GetIsReissuable() bool {
	if m != nil {
		return m.IsReissuable
	}
	return false
}

type InvokeScriptResult_Burn struct {
	AssetId              []byte   `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeScriptResult_Burn) Reset()         { *m = InvokeScriptResult_Burn{} }
func (m *InvokeScriptResult_Burn) String() string { return proto.CompactTextString(m) }
func (*InvokeScriptResult_Burn) ProtoMessage()    {}
func (*InvokeScriptResult_Burn) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7c55166e11cf76, []int{0, 3}
}

func (m *InvokeScriptResult_Burn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeScriptResult_Burn.Unmarshal(m, b)
}
func (m *InvokeScriptResult_Burn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeScriptResult_Burn.Marshal(b, m, deterministic)
}
func (m *InvokeScriptResult_Burn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeScriptResult_Burn.Merge(m, src)
}
func (m *InvokeScriptResult_Burn) XXX_Size() int {
	return xxx_messageInfo_InvokeScriptResult_Burn.Size(m)
}
func (m *InvokeScriptResult_Burn) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeScriptResult_Burn.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeScriptResult_Burn proto.InternalMessageInfo

func (m *InvokeScriptResult_Burn) GetAssetId() []byte {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *InvokeScriptResult_Burn) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type InvokeScriptResult_ErrorMessage struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeScriptResult_ErrorMessage) Reset()         { *m = InvokeScriptResult_ErrorMessage{} }
func (m *InvokeScriptResult_ErrorMessage) String() string { return proto.CompactTextString(m) }
func (*InvokeScriptResult_ErrorMessage) ProtoMessage()    {}
func (*InvokeScriptResult_ErrorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7c55166e11cf76, []int{0, 4}
}

func (m *InvokeScriptResult_ErrorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeScriptResult_ErrorMessage.Unmarshal(m, b)
}
func (m *InvokeScriptResult_ErrorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeScriptResult_ErrorMessage.Marshal(b, m, deterministic)
}
func (m *InvokeScriptResult_ErrorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeScriptResult_ErrorMessage.Merge(m, src)
}
func (m *InvokeScriptResult_ErrorMessage) XXX_Size() int {
	return xxx_messageInfo_InvokeScriptResult_ErrorMessage.Size(m)
}
func (m *InvokeScriptResult_ErrorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeScriptResult_ErrorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeScriptResult_ErrorMessage proto.InternalMessageInfo

func (m *InvokeScriptResult_ErrorMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InvokeScriptResult_ErrorMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*InvokeScriptResult)(nil), "waves.InvokeScriptResult")
	proto.RegisterType((*InvokeScriptResult_Payment)(nil), "waves.InvokeScriptResult.Payment")
	proto.RegisterType((*InvokeScriptResult_Issue)(nil), "waves.InvokeScriptResult.Issue")
	proto.RegisterType((*InvokeScriptResult_Reissue)(nil), "waves.InvokeScriptResult.Reissue")
	proto.RegisterType((*InvokeScriptResult_Burn)(nil), "waves.InvokeScriptResult.Burn")
	proto.RegisterType((*InvokeScriptResult_ErrorMessage)(nil), "waves.InvokeScriptResult.ErrorMessage")
}

func init() { proto.RegisterFile("waves/invoke_script_result.proto", fileDescriptor_4d7c55166e11cf76) }

var fileDescriptor_4d7c55166e11cf76 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x8e, 0xd3, 0x3e,
	0x10, 0x57, 0xb6, 0x4d, 0xdb, 0x9d, 0xb6, 0x17, 0xeb, 0xaf, 0xfd, 0x87, 0x1c, 0x96, 0xf0, 0xb5,
	0xea, 0x29, 0x95, 0x16, 0x04, 0xec, 0x01, 0x21, 0x56, 0xec, 0xa1, 0x20, 0x24, 0x64, 0x90, 0x90,
	0xb8, 0x54, 0x6e, 0x32, 0x0d, 0x51, 0x1b, 0xbb, 0xb2, 0x9d, 0x85, 0x7d, 0x04, 0x5e, 0x85, 0x47,
	0xe2, 0x69, 0x90, 0xc7, 0x69, 0x9b, 0x15, 0x2a, 0x48, 0x9c, 0xea, 0x19, 0xff, 0x3e, 0xa6, 0xbf,
	0x89, 0x21, 0xf9, 0x2a, 0xae, 0xd1, 0x4c, 0x4b, 0x79, 0xad, 0x56, 0x38, 0x37, 0x99, 0x2e, 0x37,
	0x76, 0xae, 0xd1, 0xd4, 0x6b, 0x9b, 0x6e, 0xb4, 0xb2, 0x8a, 0x85, 0x84, 0x88, 0xff, 0xf7, 0x40,
	0xab, 0x85, 0x34, 0x22, 0xb3, 0xa5, 0x92, 0xfe, 0x3e, 0x66, 0xfe, 0x42, 0x54, 0xaa, 0x96, 0x0d,
	0xe7, 0xfe, 0xf7, 0x3e, 0xb0, 0x19, 0x49, 0x7e, 0x20, 0x45, 0x4e, 0x82, 0xec, 0x39, 0x74, 0x73,
	0x61, 0x45, 0x14, 0x24, 0x9d, 0xc9, 0xf0, 0xfc, 0x61, 0x4a, 0xcc, 0xf4, 0xb5, 0xb0, 0xe2, 0xe3,
	0x5e, 0xd6, 0x95, 0xd4, 0xbb, 0x92, 0x56, 0xdf, 0x70, 0x62, 0xb0, 0x97, 0x70, 0x4c, 0xce, 0x4b,
	0xd4, 0x26, 0x3a, 0x22, 0xfa, 0xbd, 0x86, 0xfe, 0xbb, 0x4f, 0xfa, 0x5e, 0xdc, 0x54, 0x28, 0x2d,
	0xdf, 0x73, 0xd8, 0x33, 0xe8, 0x95, 0xc6, 0xd4, 0x68, 0xa2, 0x0e, 0xb1, 0xef, 0x1e, 0x66, 0xcf,
	0x1c, 0x8e, 0x37, 0x70, 0xf6, 0x02, 0x06, 0x1a, 0x1b, 0x6a, 0xf7, 0x6f, 0xc6, 0xdc, 0x23, 0xf9,
	0x8e, 0xc2, 0x9e, 0x40, 0xb8, 0xa8, 0xb5, 0x34, 0x51, 0x48, 0xdc, 0xd3, 0xc3, 0xdc, 0xcb, 0x5a,
	0x4b, 0xee, 0xc1, 0xec, 0x2d, 0x8c, 0x51, 0x6b, 0xa5, 0xe7, 0x15, 0x1a, 0x23, 0x0a, 0x8c, 0x7a,
	0x49, 0x30, 0x19, 0x9e, 0x9f, 0x1d, 0x66, 0x5f, 0x39, 0xf8, 0x3b, 0x8f, 0xe6, 0x23, 0x6c, 0x55,
	0xf1, 0x1b, 0xe8, 0x37, 0x81, 0xb0, 0x08, 0xfa, 0x22, 0xcf, 0x35, 0x1a, 0x13, 0x05, 0x49, 0x30,
	0x19, 0xf1, 0x6d, 0xc9, 0x1e, 0x41, 0xcf, 0x6f, 0x30, 0x3a, 0x22, 0xab, 0x71, 0x63, 0xf5, 0x8a,
	0x9a, 0xbc, 0xb9, 0x8c, 0x7f, 0x06, 0x10, 0x52, 0x3e, 0xec, 0x0e, 0x0c, 0x84, 0x31, 0x68, 0xe7,
	0x65, 0xbe, 0xd3, 0x72, 0xf5, 0x2c, 0x67, 0x0c, 0xba, 0x52, 0x54, 0x48, 0x4a, 0xc7, 0x9c, 0xce,
	0x2c, 0x81, 0x61, 0x8e, 0xfe, 0xf3, 0x2a, 0x95, 0x8c, 0x3a, 0x74, 0xd5, 0x6e, 0xb1, 0x93, 0xdd,
	0x04, 0xdd, 0x24, 0x98, 0x74, 0xb6, 0x96, 0x2c, 0x86, 0x41, 0x8e, 0x59, 0x59, 0x89, 0xb5, 0x0b,
	0x31, 0x98, 0x84, 0x7c, 0x57, 0xb3, 0x53, 0x00, 0x9f, 0xb4, 0x58, 0xac, 0x7d, 0x48, 0x03, 0xde,
	0xea, 0x38, 0x4d, 0x6f, 0x10, 0xf5, 0x69, 0xc4, 0xa6, 0x62, 0xff, 0x41, 0x28, 0x95, 0xcc, 0x30,
	0x1a, 0x90, 0x95, 0x2f, 0x62, 0x01, 0xfd, 0x66, 0x81, 0x7f, 0xfa, 0x77, 0x27, 0xb7, 0x92, 0xda,
	0xcf, 0xf9, 0x00, 0xc6, 0xa5, 0x99, 0xb7, 0xc6, 0xe9, 0xd0, 0x38, 0xa3, 0xd2, 0xf0, 0x5d, 0x2f,
	0xbe, 0x80, 0xae, 0xdb, 0xf3, 0x3f, 0xe8, 0xc7, 0x4f, 0x61, 0xd4, 0x5e, 0xb2, 0x4b, 0x39, 0x53,
	0x39, 0x12, 0x3d, 0xe4, 0x74, 0x76, 0x3d, 0x8b, 0xdf, 0xec, 0x36, 0x79, 0x77, 0xbe, 0x5c, 0xc1,
	0x59, 0xa6, 0x2a, 0xbf, 0xce, 0xcd, 0x5a, 0xd8, 0xa5, 0xd2, 0x95, 0x7f, 0xa6, 0x8b, 0x7a, 0x99,
	0xb6, 0xde, 0xf3, 0xe7, 0x8b, 0xa2, 0xb4, 0x5f, 0xea, 0x45, 0x9a, 0xa9, 0x6a, 0x7a, 0x0b, 0x3e,
	0x2d, 0x94, 0x7f, 0xe4, 0x9b, 0x55, 0x31, 0x2d, 0xf4, 0x26, 0x9b, 0x16, 0x28, 0x51, 0x0b, 0x8b,
	0xb9, 0x07, 0xfe, 0x38, 0x0a, 0x3f, 0xb9, 0xdf, 0x45, 0x8f, 0x84, 0x1f, 0xff, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0xd9, 0x3a, 0x13, 0x57, 0x04, 0x00, 0x00,
}