// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc-spend/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CommunityPoolIBCSpendProposal struct {
	Title         string                                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Recipient     string                                   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	SourceChannel string                                   `protobuf:"bytes,5,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	Timeout       uint64                                   `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (m *CommunityPoolIBCSpendProposal) Reset()      { *m = CommunityPoolIBCSpendProposal{} }
func (*CommunityPoolIBCSpendProposal) ProtoMessage() {}
func (*CommunityPoolIBCSpendProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a84a3a9fca2cd0, []int{0}
}
func (m *CommunityPoolIBCSpendProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolIBCSpendProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolIBCSpendProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolIBCSpendProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolIBCSpendProposal.Merge(m, src)
}
func (m *CommunityPoolIBCSpendProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolIBCSpendProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolIBCSpendProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolIBCSpendProposal proto.InternalMessageInfo

type CommunityPoolIBCSpendProposalWithDeposit struct {
	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description   string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	Recipient     string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
	Amount        string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty" yaml:"amount"`
	SourceChannel string `protobuf:"bytes,5,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	Timeout       uint64 `protobuf:"varint,6,opt,name=timeout,proto3" json:"timeout,omitempty" yaml:"timeout"`
	Deposit       string `protobuf:"bytes,7,opt,name=deposit,proto3" json:"deposit,omitempty" yaml:"deposit"`
}

func (m *CommunityPoolIBCSpendProposalWithDeposit) Reset() {
	*m = CommunityPoolIBCSpendProposalWithDeposit{}
}
func (m *CommunityPoolIBCSpendProposalWithDeposit) String() string { return proto.CompactTextString(m) }
func (*CommunityPoolIBCSpendProposalWithDeposit) ProtoMessage()    {}
func (*CommunityPoolIBCSpendProposalWithDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a84a3a9fca2cd0, []int{1}
}
func (m *CommunityPoolIBCSpendProposalWithDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolIBCSpendProposalWithDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolIBCSpendProposalWithDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolIBCSpendProposalWithDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolIBCSpendProposalWithDeposit.Merge(m, src)
}
func (m *CommunityPoolIBCSpendProposalWithDeposit) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolIBCSpendProposalWithDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolIBCSpendProposalWithDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolIBCSpendProposalWithDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommunityPoolIBCSpendProposal)(nil), "publicawesome.stargaze.ibcspend.CommunityPoolIBCSpendProposal")
	proto.RegisterType((*CommunityPoolIBCSpendProposalWithDeposit)(nil), "publicawesome.stargaze.ibcspend.CommunityPoolIBCSpendProposalWithDeposit")
}

func init() { proto.RegisterFile("ibc-spend/tx.proto", fileDescriptor_08a84a3a9fca2cd0) }

var fileDescriptor_08a84a3a9fca2cd0 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0xed, 0xa6, 0x49, 0x7f, 0x75, 0xff, 0xa8, 0xbf, 0x53, 0x40, 0x6e, 0x05, 0xbe, 0xc8,
	0x12, 0x28, 0x48, 0xc4, 0xa6, 0x81, 0x01, 0x65, 0x42, 0x0e, 0x0b, 0x03, 0x52, 0x15, 0x06, 0x24,
	0x16, 0x64, 0x5f, 0x4e, 0xc9, 0x09, 0xdb, 0x8f, 0xe5, 0x3b, 0x43, 0xc3, 0x2b, 0x60, 0x44, 0x62,
	0x61, 0xcc, 0x08, 0xbc, 0x92, 0x8e, 0x1d, 0x99, 0x0c, 0x4a, 0x16, 0x66, 0xbf, 0x02, 0xe4, 0x3b,
	0x27, 0x38, 0xa5, 0x62, 0xb2, 0xef, 0xf9, 0x7e, 0xbf, 0xe7, 0xc7, 0x9f, 0xbb, 0xc7, 0x40, 0x2c,
	0x20, 0x3d, 0x9e, 0xd0, 0x78, 0xec, 0x8a, 0x73, 0x27, 0x49, 0x41, 0x00, 0xc2, 0x49, 0x16, 0x84,
	0x8c, 0xf8, 0xef, 0x28, 0x87, 0x88, 0x3a, 0x5c, 0xf8, 0xe9, 0xc4, 0x7f, 0x4f, 0x1d, 0x16, 0x10,
	0xe9, 0x3c, 0x69, 0x4f, 0x60, 0x02, 0xd2, 0xeb, 0x96, 0x6f, 0x2a, 0x76, 0x62, 0x11, 0xe0, 0x11,
	0x70, 0x37, 0xf0, 0x39, 0x75, 0xdf, 0x9e, 0x06, 0x54, 0xf8, 0xa7, 0x2e, 0x01, 0x16, 0x2b, 0xdd,
	0xfe, 0xba, 0x65, 0xdc, 0x1e, 0x42, 0x14, 0x65, 0x31, 0x13, 0xb3, 0x33, 0x80, 0xf0, 0x99, 0x37,
	0x7c, 0x51, 0xee, 0x77, 0x96, 0x42, 0x02, 0xdc, 0x0f, 0x51, 0xdb, 0x68, 0x0a, 0x26, 0x42, 0x6a,
	0xea, 0x1d, 0xbd, 0xbb, 0x3b, 0x52, 0x0b, 0xd4, 0x31, 0xf6, 0xc6, 0x94, 0x93, 0x94, 0x25, 0x82,
	0x41, 0x6c, 0x6e, 0x49, 0xad, 0x5e, 0x42, 0xb7, 0x8c, 0xdd, 0x94, 0x12, 0x96, 0x30, 0x1a, 0x0b,
	0xb3, 0x21, 0xf5, 0x3f, 0x05, 0x44, 0x8c, 0x96, 0x1f, 0x41, 0x16, 0x0b, 0x73, 0xbb, 0xd3, 0xe8,
	0xee, 0xf5, 0x8f, 0x1d, 0xd5, 0xa8, 0x53, 0x36, 0xea, 0x54, 0x8d, 0x3a, 0x43, 0x60, 0xb1, 0xf7,
	0xe0, 0x22, 0xc7, 0xda, 0xb7, 0x1f, 0xb8, 0x3b, 0x61, 0x62, 0x9a, 0x05, 0x0e, 0x81, 0xc8, 0xad,
	0xfe, 0x4a, 0x3d, 0x7a, 0x7c, 0xfc, 0xc6, 0x15, 0xb3, 0x84, 0x72, 0x19, 0xe0, 0xa3, 0x6a, 0x6b,
	0x74, 0xc7, 0x38, 0xe4, 0x90, 0xa5, 0x84, 0xbe, 0x26, 0x53, 0x3f, 0x8e, 0x69, 0x68, 0x36, 0x65,
	0x1f, 0x07, 0xaa, 0x3a, 0x54, 0x45, 0x64, 0x1a, 0x3b, 0x82, 0x45, 0x14, 0x32, 0x61, 0xb6, 0x3a,
	0x7a, 0x77, 0x7b, 0xb4, 0x5a, 0x0e, 0xf6, 0x3f, 0xcc, 0xb1, 0xf6, 0x79, 0x8e, 0xb5, 0x5f, 0x73,
	0xac, 0xd9, 0x9f, 0x1a, 0x46, 0xf7, 0x9f, 0xac, 0x5e, 0x32, 0x31, 0x7d, 0x4a, 0x13, 0xe0, 0x4c,
	0xa0, 0xbb, 0x1b, 0xd8, 0xbc, 0xa3, 0x22, 0xc7, 0xfb, 0x33, 0x3f, 0x0a, 0x07, 0xb6, 0x2c, 0xdb,
	0x2b, 0x90, 0x8f, 0xaf, 0x01, 0xe9, 0xdd, 0x2c, 0x72, 0x8c, 0x94, 0xbb, 0x26, 0xda, 0x9b, 0x80,
	0xfb, 0x7f, 0x01, 0xf6, 0xda, 0x45, 0x8e, 0x8f, 0x54, 0x6e, 0x2d, 0xd9, 0x75, 0xec, 0xf7, 0x6a,
	0xd8, 0xcb, 0xc0, 0xff, 0x45, 0x8e, 0x0f, 0x54, 0x40, 0xd5, 0xed, 0x35, 0xbc, 0x27, 0xd7, 0xc3,
	0xf3, 0x8e, 0x8b, 0x1c, 0xdf, 0x50, 0x91, 0x4d, 0xdd, 0xbe, 0xca, 0xf5, 0xfe, 0x15, 0xae, 0x1e,
	0x2a, 0x72, 0x7c, 0xb8, 0x82, 0x20, 0x05, 0x7b, 0xcd, 0xba, 0x74, 0x8f, 0x15, 0x3b, 0x73, 0x47,
	0x7e, 0xa8, 0xe6, 0xae, 0x04, 0x7b, 0xb4, 0xb2, 0x0c, 0xfe, 0xab, 0x4e, 0x46, 0xef, 0x37, 0x8d,
	0xc6, 0x73, 0x3e, 0xf1, 0x46, 0x5f, 0x16, 0x96, 0x7e, 0xb1, 0xb0, 0xf4, 0xcb, 0x85, 0xa5, 0xff,
	0x5c, 0x58, 0xfa, 0xc7, 0xa5, 0xa5, 0x5d, 0x2e, 0x2d, 0xed, 0xfb, 0xd2, 0xd2, 0x5e, 0x3d, 0xaa,
	0xdd, 0x1d, 0x35, 0x48, 0xbd, 0x6a, 0x92, 0xdc, 0xd5, 0x24, 0xb9, 0xe7, 0x6e, 0x6d, 0xec, 0xca,
	0xdb, 0x14, 0xb4, 0xe4, 0x8c, 0x3c, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xff, 0x61, 0x65, 0xc3,
	0x90, 0x03, 0x00, 0x00,
}

func (this *CommunityPoolIBCSpendProposalWithDeposit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CommunityPoolIBCSpendProposalWithDeposit)
	if !ok {
		that2, ok := that.(CommunityPoolIBCSpendProposalWithDeposit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Recipient != that1.Recipient {
		return false
	}
	if this.Amount != that1.Amount {
		return false
	}
	if this.SourceChannel != that1.SourceChannel {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if this.Deposit != that1.Deposit {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publicawesome.stargaze.ibcspend.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "ibc-spend/tx.proto",
}

func (m *CommunityPoolIBCSpendProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolIBCSpendProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolIBCSpendProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timeout != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommunityPoolIBCSpendProposalWithDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolIBCSpendProposalWithDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolIBCSpendProposalWithDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timeout != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommunityPoolIBCSpendProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovTx(uint64(m.Timeout))
	}
	return n
}

func (m *CommunityPoolIBCSpendProposalWithDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovTx(uint64(m.Timeout))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommunityPoolIBCSpendProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommunityPoolIBCSpendProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolIBCSpendProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommunityPoolIBCSpendProposalWithDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommunityPoolIBCSpendProposalWithDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolIBCSpendProposalWithDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)