// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: state.proto

package corepb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance              []byte   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce                uint64   `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Vesting              []byte   `protobuf:"bytes,11,opt,name=vesting,proto3" json:"vesting,omitempty"`
	VotedRootHash        []byte   `protobuf:"bytes,12,opt,name=voted_root_hash,json=votedRootHash,proto3" json:"voted_root_hash,omitempty"`
	Bandwidth            []byte   `protobuf:"bytes,13,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	LastBandwidthTs      int64    `protobuf:"varint,14,opt,name=last_bandwidth_ts,json=lastBandwidthTs,proto3" json:"last_bandwidth_ts,omitempty"`
	Unstaking            []byte   `protobuf:"bytes,15,opt,name=unstaking,proto3" json:"unstaking,omitempty"`
	LastUnstakingTs      int64    `protobuf:"varint,16,opt,name=last_unstaking_ts,json=lastUnstakingTs,proto3" json:"last_unstaking_ts,omitempty"`
	Collateral           []byte   `protobuf:"bytes,21,opt,name=collateral,proto3" json:"collateral,omitempty"`
	VotersRootHash       []byte   `protobuf:"bytes,22,opt,name=voters_root_hash,json=votersRootHash,proto3" json:"voters_root_hash,omitempty"`
	VotePower            []byte   `protobuf:"bytes,23,opt,name=vote_power,json=votePower,proto3" json:"vote_power,omitempty"`
	TxsFromRootHash      []byte   `protobuf:"bytes,31,opt,name=txs_from_root_hash,json=txsFromRootHash,proto3" json:"txs_from_root_hash,omitempty"`
	TxsToRootHash        []byte   `protobuf:"bytes,32,opt,name=txs_to_root_hash,json=txsToRootHash,proto3" json:"txs_to_root_hash,omitempty"`
	DataRootHash         []byte   `protobuf:"bytes,40,opt,name=data_root_hash,json=dataRootHash,proto3" json:"data_root_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_f93f06acce7b4c5b, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Account) GetVesting() []byte {
	if m != nil {
		return m.Vesting
	}
	return nil
}

func (m *Account) GetVotedRootHash() []byte {
	if m != nil {
		return m.VotedRootHash
	}
	return nil
}

func (m *Account) GetBandwidth() []byte {
	if m != nil {
		return m.Bandwidth
	}
	return nil
}

func (m *Account) GetLastBandwidthTs() int64 {
	if m != nil {
		return m.LastBandwidthTs
	}
	return 0
}

func (m *Account) GetUnstaking() []byte {
	if m != nil {
		return m.Unstaking
	}
	return nil
}

func (m *Account) GetLastUnstakingTs() int64 {
	if m != nil {
		return m.LastUnstakingTs
	}
	return 0
}

func (m *Account) GetCollateral() []byte {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func (m *Account) GetVotersRootHash() []byte {
	if m != nil {
		return m.VotersRootHash
	}
	return nil
}

func (m *Account) GetVotePower() []byte {
	if m != nil {
		return m.VotePower
	}
	return nil
}

func (m *Account) GetTxsFromRootHash() []byte {
	if m != nil {
		return m.TxsFromRootHash
	}
	return nil
}

func (m *Account) GetTxsToRootHash() []byte {
	if m != nil {
		return m.TxsToRootHash
	}
	return nil
}

func (m *Account) GetDataRootHash() []byte {
	if m != nil {
		return m.DataRootHash
	}
	return nil
}

type AliasAccount struct {
	Account              []byte   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliasAccount) Reset()         { *m = AliasAccount{} }
func (m *AliasAccount) String() string { return proto.CompactTextString(m) }
func (*AliasAccount) ProtoMessage()    {}
func (*AliasAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_f93f06acce7b4c5b, []int{1}
}
func (m *AliasAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliasAccount.Unmarshal(m, b)
}
func (m *AliasAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliasAccount.Marshal(b, m, deterministic)
}
func (dst *AliasAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliasAccount.Merge(dst, src)
}
func (m *AliasAccount) XXX_Size() int {
	return xxx_messageInfo_AliasAccount.Size(m)
}
func (m *AliasAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_AliasAccount.DiscardUnknown(m)
}

var xxx_messageInfo_AliasAccount proto.InternalMessageInfo

func (m *AliasAccount) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

type DataState struct {
	TxStateRootHash            []byte   `protobuf:"bytes,1,opt,name=tx_state_root_hash,json=txStateRootHash,proto3" json:"tx_state_root_hash,omitempty"`
	RecordStateRootHash        []byte   `protobuf:"bytes,2,opt,name=record_state_root_hash,json=recordStateRootHash,proto3" json:"record_state_root_hash,omitempty"`
	CertificationStateRootHash []byte   `protobuf:"bytes,3,opt,name=certification_state_root_hash,json=certificationStateRootHash,proto3" json:"certification_state_root_hash,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *DataState) Reset()         { *m = DataState{} }
func (m *DataState) String() string { return proto.CompactTextString(m) }
func (*DataState) ProtoMessage()    {}
func (*DataState) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_f93f06acce7b4c5b, []int{2}
}
func (m *DataState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataState.Unmarshal(m, b)
}
func (m *DataState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataState.Marshal(b, m, deterministic)
}
func (dst *DataState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataState.Merge(dst, src)
}
func (m *DataState) XXX_Size() int {
	return xxx_messageInfo_DataState.Size(m)
}
func (m *DataState) XXX_DiscardUnknown() {
	xxx_messageInfo_DataState.DiscardUnknown(m)
}

var xxx_messageInfo_DataState proto.InternalMessageInfo

func (m *DataState) GetTxStateRootHash() []byte {
	if m != nil {
		return m.TxStateRootHash
	}
	return nil
}

func (m *DataState) GetRecordStateRootHash() []byte {
	if m != nil {
		return m.RecordStateRootHash
	}
	return nil
}

func (m *DataState) GetCertificationStateRootHash() []byte {
	if m != nil {
		return m.CertificationStateRootHash
	}
	return nil
}

type Record struct {
	RecordHash           []byte   `protobuf:"bytes,1,opt,name=record_hash,json=recordHash,proto3" json:"record_hash,omitempty"`
	Owner                []byte   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_f93f06acce7b4c5b, []int{3}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (dst *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(dst, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetRecordHash() []byte {
	if m != nil {
		return m.RecordHash
	}
	return nil
}

func (m *Record) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Certification struct {
	CertificateHash      []byte   `protobuf:"bytes,1,opt,name=certificate_hash,json=certificateHash,proto3" json:"certificate_hash,omitempty"`
	Issuer               []byte   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Certified            []byte   `protobuf:"bytes,3,opt,name=certified,proto3" json:"certified,omitempty"`
	IssueTime            int64    `protobuf:"varint,4,opt,name=issue_time,json=issueTime,proto3" json:"issue_time,omitempty"`
	ExpirationTime       int64    `protobuf:"varint,5,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	RevocationTime       int64    `protobuf:"varint,6,opt,name=revocation_time,json=revocationTime,proto3" json:"revocation_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Certification) Reset()         { *m = Certification{} }
func (m *Certification) String() string { return proto.CompactTextString(m) }
func (*Certification) ProtoMessage()    {}
func (*Certification) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_f93f06acce7b4c5b, []int{4}
}
func (m *Certification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certification.Unmarshal(m, b)
}
func (m *Certification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certification.Marshal(b, m, deterministic)
}
func (dst *Certification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certification.Merge(dst, src)
}
func (m *Certification) XXX_Size() int {
	return xxx_messageInfo_Certification.Size(m)
}
func (m *Certification) XXX_DiscardUnknown() {
	xxx_messageInfo_Certification.DiscardUnknown(m)
}

var xxx_messageInfo_Certification proto.InternalMessageInfo

func (m *Certification) GetCertificateHash() []byte {
	if m != nil {
		return m.CertificateHash
	}
	return nil
}

func (m *Certification) GetIssuer() []byte {
	if m != nil {
		return m.Issuer
	}
	return nil
}

func (m *Certification) GetCertified() []byte {
	if m != nil {
		return m.Certified
	}
	return nil
}

func (m *Certification) GetIssueTime() int64 {
	if m != nil {
		return m.IssueTime
	}
	return 0
}

func (m *Certification) GetExpirationTime() int64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *Certification) GetRevocationTime() int64 {
	if m != nil {
		return m.RevocationTime
	}
	return 0
}

type Alias struct {
	AliasName            string   `protobuf:"bytes,1,opt,name=alias_name,json=aliasName,proto3" json:"alias_name,omitempty"`
	AliasCollateral      []byte   `protobuf:"bytes,2,opt,name=alias_collateral,json=aliasCollateral,proto3" json:"alias_collateral,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_f93f06acce7b4c5b, []int{5}
}
func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (dst *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(dst, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetAliasName() string {
	if m != nil {
		return m.AliasName
	}
	return ""
}

func (m *Alias) GetAliasCollateral() []byte {
	if m != nil {
		return m.AliasCollateral
	}
	return nil
}

func (m *Alias) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "corepb.Account")
	proto.RegisterType((*AliasAccount)(nil), "corepb.AliasAccount")
	proto.RegisterType((*DataState)(nil), "corepb.DataState")
	proto.RegisterType((*Record)(nil), "corepb.Record")
	proto.RegisterType((*Certification)(nil), "corepb.Certification")
	proto.RegisterType((*Alias)(nil), "corepb.Alias")
}

func init() { proto.RegisterFile("state.proto", fileDescriptor_state_f93f06acce7b4c5b) }

var fileDescriptor_state_f93f06acce7b4c5b = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdb, 0x8e, 0xd3, 0x30,
	0x10, 0x40, 0x15, 0x7a, 0x59, 0x75, 0x7a, 0x49, 0x31, 0x4b, 0x89, 0x10, 0xcb, 0x56, 0x11, 0x62,
	0x03, 0x48, 0xbc, 0xec, 0x17, 0x94, 0x45, 0x88, 0x27, 0x84, 0x42, 0x79, 0x44, 0x91, 0x9b, 0x78,
	0x69, 0x44, 0x12, 0x57, 0xf6, 0xf4, 0xf2, 0x15, 0x7c, 0x0b, 0xbf, 0xc4, 0x9f, 0x20, 0x8f, 0x73,
	0x71, 0xb5, 0x12, 0x6f, 0x9d, 0xe3, 0xe3, 0xb9, 0x58, 0x93, 0xc2, 0x58, 0x23, 0x47, 0xf1, 0x7e,
	0xa7, 0x24, 0x4a, 0x36, 0x4c, 0xa5, 0x12, 0xbb, 0x4d, 0xf8, 0xbb, 0x0f, 0x17, 0xab, 0x34, 0x95,
	0xfb, 0x0a, 0x59, 0x00, 0x17, 0x3c, 0xcb, 0x94, 0xd0, 0x3a, 0xf0, 0x96, 0x5e, 0x34, 0x89, 0x9b,
	0xd0, 0x9c, 0x6c, 0x78, 0xc1, 0xab, 0x54, 0x04, 0x8f, 0xec, 0x49, 0x1d, 0xb2, 0x4b, 0x18, 0x54,
	0xd2, 0xf0, 0xde, 0xd2, 0x8b, 0xfa, 0xb1, 0x0d, 0x8c, 0x7f, 0x10, 0x1a, 0xf3, 0xea, 0x67, 0x30,
	0xb6, 0x7e, 0x1d, 0xb2, 0xd7, 0xe0, 0x1f, 0x24, 0x8a, 0x2c, 0x51, 0x52, 0x62, 0xb2, 0xe5, 0x7a,
	0x1b, 0x4c, 0xc8, 0x98, 0x12, 0x8e, 0xa5, 0xc4, 0xcf, 0x5c, 0x6f, 0xd9, 0x0b, 0x18, 0x6d, 0x78,
	0x95, 0x1d, 0xf3, 0x0c, 0xb7, 0xc1, 0x94, 0x8c, 0x0e, 0xb0, 0xb7, 0xf0, 0xb8, 0xe0, 0x1a, 0x93,
	0x96, 0x24, 0xa8, 0x83, 0xd9, 0xd2, 0x8b, 0x7a, 0xb1, 0x6f, 0x0e, 0x3e, 0x34, 0x7c, 0xad, 0x4d,
	0xa6, 0x7d, 0xa5, 0x91, 0xff, 0x32, 0xdd, 0xf8, 0x36, 0x53, 0x0b, 0xda, 0x4c, 0x2d, 0x31, 0x99,
	0xe6, 0x5d, 0xa6, 0xef, 0x0d, 0x5f, 0x6b, 0xf6, 0x12, 0x20, 0x95, 0x45, 0xc1, 0x51, 0x28, 0x5e,
	0x04, 0x4f, 0x29, 0x95, 0x43, 0x58, 0x04, 0x73, 0x33, 0x84, 0xd2, 0xce, 0x70, 0x0b, 0xb2, 0x66,
	0x96, 0xb7, 0xd3, 0x5d, 0x01, 0x18, 0x92, 0xec, 0xe4, 0x51, 0xa8, 0xe0, 0x99, 0x6d, 0xca, 0x90,
	0xaf, 0x06, 0xb0, 0x77, 0xc0, 0xf0, 0xa4, 0x93, 0x7b, 0x25, 0x4b, 0x27, 0xd5, 0x35, 0x69, 0x3e,
	0x9e, 0xf4, 0x27, 0x25, 0xcb, 0x36, 0xd7, 0x0d, 0xcc, 0x8d, 0x8c, 0xd2, 0x51, 0x97, 0xf6, 0x49,
	0xf1, 0xa4, 0xd7, 0xb2, 0x15, 0x5f, 0xc1, 0x2c, 0xe3, 0xc8, 0x1d, 0x2d, 0x22, 0x6d, 0x62, 0x68,
	0x63, 0x85, 0x11, 0x4c, 0x56, 0x45, 0xce, 0xb5, 0xbb, 0x14, 0xf6, 0x67, 0xbb, 0x14, 0x36, 0x0c,
	0xff, 0x78, 0x30, 0xfa, 0xc8, 0x91, 0x7f, 0x33, 0x6b, 0x65, 0x7b, 0x4e, 0x68, 0xc5, 0x9c, 0x0a,
	0x5e, 0xd3, 0x33, 0x49, 0x6d, 0x2b, 0xb7, 0xb0, 0x50, 0x22, 0x95, 0x2a, 0x7b, 0x70, 0xc1, 0xae,
	0xd7, 0x13, 0x7b, 0x7a, 0x7e, 0x69, 0x05, 0x57, 0xa9, 0x50, 0x98, 0xdf, 0xe7, 0x29, 0xc7, 0x5c,
	0x56, 0x0f, 0xee, 0xf6, 0xe8, 0xee, 0xf3, 0x33, 0xe9, 0x2c, 0x45, 0xf8, 0x03, 0x86, 0x31, 0x65,
	0x66, 0xd7, 0x30, 0xae, 0x3b, 0x70, 0xfa, 0x04, 0x8b, 0xa8, 0xda, 0x25, 0x0c, 0xe4, 0xb1, 0x12,
	0xaa, 0xee, 0xc8, 0x06, 0x66, 0x99, 0x30, 0x2f, 0x85, 0x46, 0x5e, 0xee, 0xa8, 0x5e, 0x2f, 0xee,
	0x40, 0xf8, 0xd7, 0x83, 0xe9, 0x9d, 0x5b, 0x9d, 0xbd, 0x81, 0x79, 0xd7, 0x8e, 0x38, 0x7b, 0x13,
	0x87, 0x53, 0xc1, 0x05, 0x0c, 0x73, 0xad, 0xf7, 0x6d, 0xc5, 0x3a, 0x32, 0x25, 0x6b, 0x55, 0x64,
	0xf5, 0x88, 0x1d, 0x30, 0x9b, 0x44, 0x5e, 0x62, 0xba, 0x08, 0xfa, 0xb6, 0x23, 0x22, 0xeb, 0xbc,
	0x14, 0xec, 0x06, 0x7c, 0x71, 0xda, 0xe5, 0xca, 0x3e, 0x18, 0x39, 0x03, 0x72, 0x66, 0x1d, 0x6e,
	0x44, 0x25, 0x0e, 0x32, 0x75, 0xc4, 0xa1, 0x15, 0x3b, 0x6c, 0xc4, 0x50, 0xc2, 0x80, 0xf6, 0xc3,
	0x54, 0xe6, 0xe6, 0x47, 0x52, 0xf1, 0x52, 0xd0, 0x50, 0xa3, 0x78, 0x44, 0xe4, 0x0b, 0x2f, 0x85,
	0x99, 0xdc, 0x1e, 0x3b, 0x9f, 0x8c, 0x1d, 0xcc, 0x27, 0x7e, 0xd7, 0x7d, 0x37, 0xff, 0x7d, 0xd4,
	0xcd, 0x90, 0xfe, 0xb0, 0x6e, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xa3, 0xc9, 0xea, 0xbf,
	0x04, 0x00, 0x00,
}
