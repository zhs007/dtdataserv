// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dtdata.proto

package dtdatapb

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

// DTDataType - dtdata type
type DTDataType int32

const (
	DTDataType_DTDT_GAMEDAYREPORT DTDataType = 0
)

var DTDataType_name = map[int32]string{
	0: "DTDT_GAMEDAYREPORT",
}
var DTDataType_value = map[string]int32{
	"DTDT_GAMEDAYREPORT": 0,
}

func (x DTDataType) String() string {
	return proto.EnumName(DTDataType_name, int32(x))
}
func (DTDataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dtdata_77af717fb14ef449, []int{0}
}

// DTDataServCtrlInfo - DTData server ctrl info
type DTDataServCtrlInfo struct {
	Type                 DTDataType `protobuf:"varint,1,opt,name=type,proto3,enum=dtdatapb.DTDataType" json:"type,omitempty"`
	StartTime            string     `protobuf:"bytes,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              string     `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Env                  string     `protobuf:"bytes,4,opt,name=env,proto3" json:"env,omitempty"`
	Currency             string     `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	ScaleMoney           int32      `protobuf:"varint,6,opt,name=scaleMoney,proto3" json:"scaleMoney,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DTDataServCtrlInfo) Reset()         { *m = DTDataServCtrlInfo{} }
func (m *DTDataServCtrlInfo) String() string { return proto.CompactTextString(m) }
func (*DTDataServCtrlInfo) ProtoMessage()    {}
func (*DTDataServCtrlInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dtdata_77af717fb14ef449, []int{0}
}
func (m *DTDataServCtrlInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DTDataServCtrlInfo.Unmarshal(m, b)
}
func (m *DTDataServCtrlInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DTDataServCtrlInfo.Marshal(b, m, deterministic)
}
func (dst *DTDataServCtrlInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTDataServCtrlInfo.Merge(dst, src)
}
func (m *DTDataServCtrlInfo) XXX_Size() int {
	return xxx_messageInfo_DTDataServCtrlInfo.Size(m)
}
func (m *DTDataServCtrlInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DTDataServCtrlInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DTDataServCtrlInfo proto.InternalMessageInfo

func (m *DTDataServCtrlInfo) GetType() DTDataType {
	if m != nil {
		return m.Type
	}
	return DTDataType_DTDT_GAMEDAYREPORT
}

func (m *DTDataServCtrlInfo) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *DTDataServCtrlInfo) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *DTDataServCtrlInfo) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *DTDataServCtrlInfo) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *DTDataServCtrlInfo) GetScaleMoney() int32 {
	if m != nil {
		return m.ScaleMoney
	}
	return 0
}

// DTDataServCtrlReply - DTData server ctrl reply
type DTDataServCtrlReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DTDataServCtrlReply) Reset()         { *m = DTDataServCtrlReply{} }
func (m *DTDataServCtrlReply) String() string { return proto.CompactTextString(m) }
func (*DTDataServCtrlReply) ProtoMessage()    {}
func (*DTDataServCtrlReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dtdata_77af717fb14ef449, []int{1}
}
func (m *DTDataServCtrlReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DTDataServCtrlReply.Unmarshal(m, b)
}
func (m *DTDataServCtrlReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DTDataServCtrlReply.Marshal(b, m, deterministic)
}
func (dst *DTDataServCtrlReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTDataServCtrlReply.Merge(dst, src)
}
func (m *DTDataServCtrlReply) XXX_Size() int {
	return xxx_messageInfo_DTDataServCtrlReply.Size(m)
}
func (m *DTDataServCtrlReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DTDataServCtrlReply.DiscardUnknown(m)
}

var xxx_messageInfo_DTDataServCtrlReply proto.InternalMessageInfo

func (m *DTDataServCtrlReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// DTGameReport - dt game report
type DTGameReport struct {
	Rtp                  float32             `protobuf:"fixed32,1,opt,name=rtp,proto3" json:"rtp,omitempty"`
	TotalBet             float32             `protobuf:"fixed32,2,opt,name=totalBet,proto3" json:"totalBet,omitempty"`
	TotalWin             float32             `protobuf:"fixed32,3,opt,name=totalWin,proto3" json:"totalWin,omitempty"`
	SpinNums             int64               `protobuf:"varint,4,opt,name=spinNums,proto3" json:"spinNums,omitempty"`
	CurrencyNums         int32               `protobuf:"varint,5,opt,name=currencyNums,proto3" json:"currencyNums,omitempty"`
	MainCurrency         string              `protobuf:"bytes,6,opt,name=mainCurrency,proto3" json:"mainCurrency,omitempty"`
	BusinessNums         int32               `protobuf:"varint,7,opt,name=businessNums,proto3" json:"businessNums,omitempty"`
	Businessid           []string            `protobuf:"bytes,20,rep,name=businessid,proto3" json:"businessid,omitempty"`
	BusinessReport       []*DTBusinessReport `protobuf:"bytes,21,rep,name=businessReport,proto3" json:"businessReport,omitempty"`
	GameCode             string              `protobuf:"bytes,100,opt,name=gameCode,proto3" json:"gameCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DTGameReport) Reset()         { *m = DTGameReport{} }
func (m *DTGameReport) String() string { return proto.CompactTextString(m) }
func (*DTGameReport) ProtoMessage()    {}
func (*DTGameReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_dtdata_77af717fb14ef449, []int{2}
}
func (m *DTGameReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DTGameReport.Unmarshal(m, b)
}
func (m *DTGameReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DTGameReport.Marshal(b, m, deterministic)
}
func (dst *DTGameReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTGameReport.Merge(dst, src)
}
func (m *DTGameReport) XXX_Size() int {
	return xxx_messageInfo_DTGameReport.Size(m)
}
func (m *DTGameReport) XXX_DiscardUnknown() {
	xxx_messageInfo_DTGameReport.DiscardUnknown(m)
}

var xxx_messageInfo_DTGameReport proto.InternalMessageInfo

func (m *DTGameReport) GetRtp() float32 {
	if m != nil {
		return m.Rtp
	}
	return 0
}

func (m *DTGameReport) GetTotalBet() float32 {
	if m != nil {
		return m.TotalBet
	}
	return 0
}

func (m *DTGameReport) GetTotalWin() float32 {
	if m != nil {
		return m.TotalWin
	}
	return 0
}

func (m *DTGameReport) GetSpinNums() int64 {
	if m != nil {
		return m.SpinNums
	}
	return 0
}

func (m *DTGameReport) GetCurrencyNums() int32 {
	if m != nil {
		return m.CurrencyNums
	}
	return 0
}

func (m *DTGameReport) GetMainCurrency() string {
	if m != nil {
		return m.MainCurrency
	}
	return ""
}

func (m *DTGameReport) GetBusinessNums() int32 {
	if m != nil {
		return m.BusinessNums
	}
	return 0
}

func (m *DTGameReport) GetBusinessid() []string {
	if m != nil {
		return m.Businessid
	}
	return nil
}

func (m *DTGameReport) GetBusinessReport() []*DTBusinessReport {
	if m != nil {
		return m.BusinessReport
	}
	return nil
}

func (m *DTGameReport) GetGameCode() string {
	if m != nil {
		return m.GameCode
	}
	return ""
}

// DTBusinessReport - dt business report
type DTBusinessReport struct {
	Rtp                  float32         `protobuf:"fixed32,1,opt,name=rtp,proto3" json:"rtp,omitempty"`
	TotalBet             float32         `protobuf:"fixed32,2,opt,name=totalBet,proto3" json:"totalBet,omitempty"`
	TotalWin             float32         `protobuf:"fixed32,3,opt,name=totalWin,proto3" json:"totalWin,omitempty"`
	SpinNums             int64           `protobuf:"varint,4,opt,name=spinNums,proto3" json:"spinNums,omitempty"`
	CurrencyNums         int32           `protobuf:"varint,5,opt,name=currencyNums,proto3" json:"currencyNums,omitempty"`
	GameNums             int32           `protobuf:"varint,6,opt,name=gameNums,proto3" json:"gameNums,omitempty"`
	MainCurrency         string          `protobuf:"bytes,7,opt,name=mainCurrency,proto3" json:"mainCurrency,omitempty"`
	Gamecode             []string        `protobuf:"bytes,20,rep,name=gamecode,proto3" json:"gamecode,omitempty"`
	GameReport           []*DTGameReport `protobuf:"bytes,21,rep,name=gameReport,proto3" json:"gameReport,omitempty"`
	BusinessID           string          `protobuf:"bytes,100,opt,name=businessID,proto3" json:"businessID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DTBusinessReport) Reset()         { *m = DTBusinessReport{} }
func (m *DTBusinessReport) String() string { return proto.CompactTextString(m) }
func (*DTBusinessReport) ProtoMessage()    {}
func (*DTBusinessReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_dtdata_77af717fb14ef449, []int{3}
}
func (m *DTBusinessReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DTBusinessReport.Unmarshal(m, b)
}
func (m *DTBusinessReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DTBusinessReport.Marshal(b, m, deterministic)
}
func (dst *DTBusinessReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTBusinessReport.Merge(dst, src)
}
func (m *DTBusinessReport) XXX_Size() int {
	return xxx_messageInfo_DTBusinessReport.Size(m)
}
func (m *DTBusinessReport) XXX_DiscardUnknown() {
	xxx_messageInfo_DTBusinessReport.DiscardUnknown(m)
}

var xxx_messageInfo_DTBusinessReport proto.InternalMessageInfo

func (m *DTBusinessReport) GetRtp() float32 {
	if m != nil {
		return m.Rtp
	}
	return 0
}

func (m *DTBusinessReport) GetTotalBet() float32 {
	if m != nil {
		return m.TotalBet
	}
	return 0
}

func (m *DTBusinessReport) GetTotalWin() float32 {
	if m != nil {
		return m.TotalWin
	}
	return 0
}

func (m *DTBusinessReport) GetSpinNums() int64 {
	if m != nil {
		return m.SpinNums
	}
	return 0
}

func (m *DTBusinessReport) GetCurrencyNums() int32 {
	if m != nil {
		return m.CurrencyNums
	}
	return 0
}

func (m *DTBusinessReport) GetGameNums() int32 {
	if m != nil {
		return m.GameNums
	}
	return 0
}

func (m *DTBusinessReport) GetMainCurrency() string {
	if m != nil {
		return m.MainCurrency
	}
	return ""
}

func (m *DTBusinessReport) GetGamecode() []string {
	if m != nil {
		return m.Gamecode
	}
	return nil
}

func (m *DTBusinessReport) GetGameReport() []*DTGameReport {
	if m != nil {
		return m.GameReport
	}
	return nil
}

func (m *DTBusinessReport) GetBusinessID() string {
	if m != nil {
		return m.BusinessID
	}
	return ""
}

// DTReport - dt report
type DTReport struct {
	Rtp                  float32             `protobuf:"fixed32,1,opt,name=rtp,proto3" json:"rtp,omitempty"`
	TotalBet             float32             `protobuf:"fixed32,2,opt,name=totalBet,proto3" json:"totalBet,omitempty"`
	TotalWin             float32             `protobuf:"fixed32,3,opt,name=totalWin,proto3" json:"totalWin,omitempty"`
	SpinNums             int64               `protobuf:"varint,4,opt,name=spinNums,proto3" json:"spinNums,omitempty"`
	CurrencyNums         int32               `protobuf:"varint,5,opt,name=currencyNums,proto3" json:"currencyNums,omitempty"`
	GameNums             int32               `protobuf:"varint,6,opt,name=gameNums,proto3" json:"gameNums,omitempty"`
	MainCurrency         string              `protobuf:"bytes,7,opt,name=mainCurrency,proto3" json:"mainCurrency,omitempty"`
	BusinessNums         int32               `protobuf:"varint,8,opt,name=businessNums,proto3" json:"businessNums,omitempty"`
	Err                  string              `protobuf:"bytes,30,opt,name=err,proto3" json:"err,omitempty"`
	TopGames             []*DTGameReport     `protobuf:"bytes,100,rep,name=topGames,proto3" json:"topGames,omitempty"`
	TopBusiness          []*DTBusinessReport `protobuf:"bytes,101,rep,name=topBusiness,proto3" json:"topBusiness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DTReport) Reset()         { *m = DTReport{} }
func (m *DTReport) String() string { return proto.CompactTextString(m) }
func (*DTReport) ProtoMessage()    {}
func (*DTReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_dtdata_77af717fb14ef449, []int{4}
}
func (m *DTReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DTReport.Unmarshal(m, b)
}
func (m *DTReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DTReport.Marshal(b, m, deterministic)
}
func (dst *DTReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTReport.Merge(dst, src)
}
func (m *DTReport) XXX_Size() int {
	return xxx_messageInfo_DTReport.Size(m)
}
func (m *DTReport) XXX_DiscardUnknown() {
	xxx_messageInfo_DTReport.DiscardUnknown(m)
}

var xxx_messageInfo_DTReport proto.InternalMessageInfo

func (m *DTReport) GetRtp() float32 {
	if m != nil {
		return m.Rtp
	}
	return 0
}

func (m *DTReport) GetTotalBet() float32 {
	if m != nil {
		return m.TotalBet
	}
	return 0
}

func (m *DTReport) GetTotalWin() float32 {
	if m != nil {
		return m.TotalWin
	}
	return 0
}

func (m *DTReport) GetSpinNums() int64 {
	if m != nil {
		return m.SpinNums
	}
	return 0
}

func (m *DTReport) GetCurrencyNums() int32 {
	if m != nil {
		return m.CurrencyNums
	}
	return 0
}

func (m *DTReport) GetGameNums() int32 {
	if m != nil {
		return m.GameNums
	}
	return 0
}

func (m *DTReport) GetMainCurrency() string {
	if m != nil {
		return m.MainCurrency
	}
	return ""
}

func (m *DTReport) GetBusinessNums() int32 {
	if m != nil {
		return m.BusinessNums
	}
	return 0
}

func (m *DTReport) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *DTReport) GetTopGames() []*DTGameReport {
	if m != nil {
		return m.TopGames
	}
	return nil
}

func (m *DTReport) GetTopBusiness() []*DTBusinessReport {
	if m != nil {
		return m.TopBusiness
	}
	return nil
}

func init() {
	proto.RegisterType((*DTDataServCtrlInfo)(nil), "dtdatapb.DTDataServCtrlInfo")
	proto.RegisterType((*DTDataServCtrlReply)(nil), "dtdatapb.DTDataServCtrlReply")
	proto.RegisterType((*DTGameReport)(nil), "dtdatapb.DTGameReport")
	proto.RegisterType((*DTBusinessReport)(nil), "dtdatapb.DTBusinessReport")
	proto.RegisterType((*DTReport)(nil), "dtdatapb.DTReport")
	proto.RegisterEnum("dtdatapb.DTDataType", DTDataType_name, DTDataType_value)
}

func init() { proto.RegisterFile("dtdata.proto", fileDescriptor_dtdata_77af717fb14ef449) }

var fileDescriptor_dtdata_77af717fb14ef449 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0xa4, 0xf8, 0x6b, 0x62, 0x82, 0xd9, 0xba, 0x61, 0x09, 0x25, 0x08, 0xd1, 0x83, 0x68,
	0xc1, 0x07, 0x17, 0x7a, 0xea, 0x25, 0xb6, 0x42, 0xc8, 0x21, 0x6d, 0xd9, 0x0a, 0x42, 0x4f, 0x65,
	0x6d, 0x4d, 0x83, 0xa8, 0xb5, 0x5a, 0x56, 0xeb, 0x80, 0xfe, 0x5f, 0x7f, 0x40, 0x7f, 0x42, 0x7f,
	0x48, 0x0f, 0x65, 0x57, 0x1f, 0x96, 0x5d, 0xd3, 0x5e, 0x4b, 0x6f, 0x7a, 0x6f, 0x9e, 0x96, 0x79,
	0x6f, 0x86, 0x81, 0x71, 0xa2, 0x13, 0xae, 0xf9, 0x4c, 0xaa, 0x5c, 0xe7, 0x64, 0x58, 0x21, 0xb9,
	0x0a, 0xbe, 0x39, 0x40, 0xa2, 0x38, 0xe2, 0x9a, 0x7f, 0x44, 0xf5, 0xb8, 0xd4, 0x6a, 0x73, 0x2b,
	0xbe, 0xe4, 0x24, 0x84, 0x13, 0x5d, 0x4a, 0xa4, 0x8e, 0xef, 0x84, 0x67, 0xf3, 0xe9, 0xac, 0xd1,
	0xcf, 0x2a, 0x6d, 0x5c, 0x4a, 0x64, 0x56, 0x41, 0x9e, 0xc3, 0xa8, 0xd0, 0x5c, 0xe9, 0x38, 0xcd,
	0x90, 0xba, 0xbe, 0x13, 0x8e, 0xd8, 0x8e, 0x20, 0x14, 0x06, 0x28, 0x12, 0x5b, 0xf3, 0x6c, 0xad,
	0x81, 0x64, 0x02, 0x1e, 0x8a, 0x47, 0x7a, 0x62, 0x59, 0xf3, 0x49, 0x2e, 0x60, 0xb8, 0xde, 0x2a,
	0x85, 0x62, 0x5d, 0xd2, 0x9e, 0xa5, 0x5b, 0x4c, 0x2e, 0x01, 0x8a, 0x35, 0xdf, 0xe0, 0x5d, 0x2e,
	0xb0, 0xa4, 0x7d, 0xdf, 0x09, 0x7b, 0xac, 0xc3, 0x04, 0xaf, 0xe0, 0xe9, 0xbe, 0x0b, 0x86, 0x72,
	0x53, 0x92, 0x29, 0xf4, 0x74, 0xfe, 0x15, 0x85, 0xf5, 0x31, 0x62, 0x15, 0x08, 0x7e, 0xb8, 0x30,
	0x8e, 0xe2, 0x1b, 0x9e, 0x21, 0x43, 0x99, 0x2b, 0x6d, 0x7a, 0x51, 0x5a, 0x5a, 0x91, 0xcb, 0xcc,
	0xa7, 0xe9, 0x45, 0xe7, 0x9a, 0x6f, 0x16, 0xa8, 0xad, 0x29, 0x97, 0xb5, 0xb8, 0xad, 0xdd, 0xa7,
	0xc2, 0x9a, 0x6a, 0x6a, 0xf7, 0xa9, 0x30, 0xb5, 0x42, 0xa6, 0xe2, 0xdd, 0x36, 0x2b, 0xac, 0x35,
	0x8f, 0xb5, 0x98, 0x04, 0x30, 0x6e, 0xfc, 0xd8, 0x7a, 0xcf, 0xba, 0xd8, 0xe3, 0x8c, 0x26, 0xe3,
	0xa9, 0x58, 0x36, 0x39, 0xf4, 0x6d, 0xdf, 0x7b, 0x9c, 0xd1, 0xac, 0xb6, 0x45, 0x2a, 0xb0, 0x28,
	0xec, 0x3b, 0x83, 0xea, 0x9d, 0x2e, 0x67, 0xf2, 0x6a, 0x70, 0x9a, 0xd0, 0xa9, 0xef, 0x85, 0x23,
	0xd6, 0x61, 0xc8, 0x02, 0xce, 0x1a, 0x54, 0x65, 0x40, 0x9f, 0xf9, 0x5e, 0x78, 0x3a, 0xbf, 0xe8,
	0x4e, 0x7a, 0xb1, 0xa7, 0x60, 0x07, 0x7f, 0x18, 0xaf, 0x0f, 0x3c, 0xc3, 0x65, 0x9e, 0x20, 0x4d,
	0xaa, 0x79, 0x35, 0x38, 0xf8, 0xee, 0xc2, 0xe4, 0xf0, 0x81, 0x7f, 0x28, 0xe6, 0xba, 0x75, 0x5b,
	0xaf, 0x96, 0xa9, 0xc5, 0xbf, 0x8d, 0x60, 0x70, 0x64, 0x04, 0xf5, 0xff, 0x6b, 0x63, 0xbd, 0x0a,
	0xb7, 0xc5, 0xe4, 0x0d, 0xc0, 0x43, 0xbb, 0x5a, 0x75, 0xac, 0xe7, 0xdd, 0x58, 0x77, 0x8b, 0xc7,
	0x3a, 0xca, 0xee, 0xc8, 0x6e, 0xa3, 0x3a, 0xd0, 0x0e, 0x13, 0xfc, 0x74, 0x61, 0x18, 0xc5, 0xff,
	0x5d, 0x94, 0x87, 0xdb, 0x3c, 0x3c, 0xb2, 0xcd, 0xe6, 0x56, 0x28, 0x45, 0x2f, 0xeb, 0x5b, 0xa1,
	0x14, 0x99, 0x1b, 0x47, 0xd2, 0x24, 0x59, 0xd0, 0xe4, 0x8f, 0x11, 0xb7, 0x3a, 0xf2, 0x16, 0x4e,
	0x75, 0x2e, 0x9b, 0x9d, 0xa4, 0xf8, 0xd7, 0x85, 0xef, 0xca, 0x5f, 0xbe, 0x00, 0xd8, 0xdd, 0x3e,
	0x72, 0x6e, 0xaf, 0x66, 0xfc, 0xf9, 0xe6, 0xea, 0xee, 0x3a, 0xba, 0xfa, 0xc4, 0xae, 0x3f, 0xbc,
	0x67, 0xf1, 0xe4, 0xc9, 0xaa, 0x6f, 0xef, 0xeb, 0xeb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae,
	0x63, 0xce, 0x4f, 0x6f, 0x05, 0x00, 0x00,
}
