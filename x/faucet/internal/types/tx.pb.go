// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/faucet/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgMintResponse struct {
}

func (m *MsgMintResponse) Reset()         { *m = MsgMintResponse{} }
func (m *MsgMintResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintResponse) ProtoMessage()    {}
func (*MsgMintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e9d9c3aef384d7, []int{0}
}
func (m *MsgMintResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintResponse.Merge(m, src)
}
func (m *MsgMintResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintResponse proto.InternalMessageInfo

type MsgFaucetKeyResponse struct {
}

func (m *MsgFaucetKeyResponse) Reset()         { *m = MsgFaucetKeyResponse{} }
func (m *MsgFaucetKeyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFaucetKeyResponse) ProtoMessage()    {}
func (*MsgFaucetKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e9d9c3aef384d7, []int{1}
}
func (m *MsgFaucetKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFaucetKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFaucetKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFaucetKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFaucetKeyResponse.Merge(m, src)
}
func (m *MsgFaucetKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFaucetKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFaucetKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFaucetKeyResponse proto.InternalMessageInfo

type MsgMint struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender" yaml:"sender"`
	Minter string `protobuf:"bytes,2,opt,name=minter,proto3" json:"minter" yaml:"minter"`
	Time   int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time" yaml:"time"`
	Denom  string `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom" yaml:"denom"`
}

func (m *MsgMint) Reset()         { *m = MsgMint{} }
func (m *MsgMint) String() string { return proto.CompactTextString(m) }
func (*MsgMint) ProtoMessage()    {}
func (*MsgMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e9d9c3aef384d7, []int{2}
}
func (m *MsgMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMint.Merge(m, src)
}
func (m *MsgMint) XXX_Size() int {
	return m.Size()
}
func (m *MsgMint) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMint.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMint proto.InternalMessageInfo

func (m *MsgMint) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgMint) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

func (m *MsgMint) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MsgMint) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type MsgFaucetKey struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender" yaml:"sender"`
	Armor  string `protobuf:"bytes,2,opt,name=armor,proto3" json:"armor" yaml:"armor"`
}

func (m *MsgFaucetKey) Reset()         { *m = MsgFaucetKey{} }
func (m *MsgFaucetKey) String() string { return proto.CompactTextString(m) }
func (*MsgFaucetKey) ProtoMessage()    {}
func (*MsgFaucetKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e9d9c3aef384d7, []int{3}
}
func (m *MsgFaucetKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFaucetKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFaucetKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFaucetKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFaucetKey.Merge(m, src)
}
func (m *MsgFaucetKey) XXX_Size() int {
	return m.Size()
}
func (m *MsgFaucetKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFaucetKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFaucetKey proto.InternalMessageInfo

func (m *MsgFaucetKey) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgFaucetKey) GetArmor() string {
	if m != nil {
		return m.Armor
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgMintResponse)(nil), "stargaze.faucet.v1beta1.MsgMintResponse")
	proto.RegisterType((*MsgFaucetKeyResponse)(nil), "stargaze.faucet.v1beta1.MsgFaucetKeyResponse")
	proto.RegisterType((*MsgMint)(nil), "stargaze.faucet.v1beta1.MsgMint")
	proto.RegisterType((*MsgFaucetKey)(nil), "stargaze.faucet.v1beta1.MsgFaucetKey")
}

func init() { proto.RegisterFile("stargaze/faucet/v1beta1/tx.proto", fileDescriptor_c6e9d9c3aef384d7) }

var fileDescriptor_c6e9d9c3aef384d7 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x6b, 0xe2, 0x40,
	0x18, 0xc6, 0x9d, 0x35, 0xba, 0x38, 0xeb, 0xb2, 0x6c, 0x90, 0x35, 0xeb, 0x42, 0x12, 0x06, 0x16,
	0x84, 0x62, 0x06, 0xeb, 0xcd, 0xa3, 0x87, 0x5e, 0x4a, 0x2e, 0xa1, 0xa7, 0xde, 0x26, 0x3a, 0x4d,
	0x03, 0x26, 0x13, 0x32, 0x63, 0xab, 0x7e, 0x8a, 0x7e, 0x84, 0x7e, 0x90, 0x7e, 0x80, 0x1e, 0xbd,
	0xb5, 0xa7, 0x50, 0xf4, 0x52, 0x3c, 0xfa, 0x09, 0x8a, 0x33, 0x89, 0xe9, 0xa5, 0x7f, 0xe8, 0x2d,
	0xcf, 0x93, 0xdf, 0xc3, 0xc3, 0x3b, 0xef, 0x0b, 0x6d, 0x2e, 0x48, 0x1a, 0x90, 0x25, 0xc5, 0x17,
	0x64, 0x36, 0xa6, 0x02, 0x5f, 0xf5, 0x7d, 0x2a, 0x48, 0x1f, 0x8b, 0xb9, 0x93, 0xa4, 0x4c, 0x30,
	0xbd, 0x5d, 0x10, 0x8e, 0x22, 0x9c, 0x9c, 0xe8, 0xb4, 0x02, 0x16, 0x30, 0xc9, 0xe0, 0xfd, 0x97,
	0xc2, 0xd1, 0x6f, 0xf8, 0xcb, 0xe5, 0x81, 0x1b, 0xc6, 0xc2, 0xa3, 0x3c, 0x61, 0x31, 0xa7, 0xe8,
	0x0f, 0x6c, 0xb9, 0x3c, 0x38, 0x91, 0xe9, 0x53, 0xba, 0x38, 0xf8, 0x0f, 0x00, 0x7e, 0xcf, 0x59,
	0x7d, 0x00, 0xeb, 0x9c, 0xc6, 0x13, 0x9a, 0x1a, 0xc0, 0x06, 0xdd, 0xc6, 0xe8, 0xdf, 0x36, 0xb3,
	0x72, 0x67, 0x97, 0x59, 0x3f, 0x17, 0x24, 0x9a, 0x0e, 0x91, 0xd2, 0xc8, 0xcb, 0x7f, 0xec, 0x43,
	0x51, 0x18, 0x0b, 0x9a, 0x1a, 0xdf, 0xca, 0x90, 0x72, 0xca, 0x90, 0xd2, 0xc8, 0xcb, 0x7f, 0xe8,
	0x47, 0x50, 0x13, 0x61, 0x44, 0x8d, 0xaa, 0x0d, 0xba, 0xd5, 0x51, 0x7b, 0x9b, 0x59, 0x52, 0xef,
	0x32, 0xeb, 0x87, 0x0a, 0xec, 0x15, 0xf2, 0xa4, 0xa9, 0x63, 0x58, 0x9b, 0xd0, 0x98, 0x45, 0x86,
	0x26, 0x0b, 0xfe, 0x6e, 0x33, 0x4b, 0x19, 0xbb, 0xcc, 0x6a, 0x2a, 0x5c, 0x4a, 0xe4, 0x29, 0x7b,
	0xa8, 0x3d, 0xdf, 0x5a, 0x00, 0x2d, 0x61, 0xf3, 0xf5, 0xc4, 0x5f, 0x9b, 0x0e, 0xc3, 0x1a, 0x49,
	0x23, 0x56, 0x0c, 0x27, 0xbb, 0xa5, 0x51, 0x76, 0x4b, 0x89, 0x3c, 0x65, 0xab, 0xee, 0xe3, 0x3b,
	0x00, 0xab, 0x2e, 0x0f, 0x74, 0x0f, 0x6a, 0xf2, 0x65, 0x6d, 0xe7, 0x8d, 0x05, 0x3a, 0xf9, 0xdb,
	0x77, 0xba, 0x1f, 0x11, 0xc5, 0xc6, 0x74, 0x02, 0x1b, 0xe5, 0x50, 0xff, 0xdf, 0x8b, 0x1d, 0xb0,
	0x4e, 0xef, 0x53, 0x58, 0x51, 0x31, 0x3a, 0xbb, 0x5f, 0x9b, 0x60, 0xb5, 0x36, 0xc1, 0xd3, 0xda,
	0x04, 0x37, 0x1b, 0xb3, 0xb2, 0xda, 0x98, 0x95, 0xc7, 0x8d, 0x59, 0x39, 0x1f, 0x06, 0xa1, 0xb8,
	0x9c, 0xf9, 0xce, 0x98, 0x45, 0x38, 0x99, 0xf9, 0xd3, 0x70, 0xdc, 0x23, 0xd7, 0x94, 0xb3, 0x88,
	0xe2, 0xc3, 0x11, 0xcf, 0x8b, 0x33, 0x96, 0xbb, 0x8e, 0xc9, 0x14, 0x8b, 0x45, 0x42, 0xb9, 0x5f,
	0x97, 0xc7, 0x39, 0x78, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x45, 0xb6, 0xfc, 0xef, 0x02, 0x00,
	0x00,
}

func (this *MsgMint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgMint)
	if !ok {
		that2, ok := that.(MsgMint)
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
	if this.Sender != that1.Sender {
		return false
	}
	if this.Minter != that1.Minter {
		return false
	}
	if this.Time != that1.Time {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	return true
}
func (this *MsgFaucetKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgFaucetKey)
	if !ok {
		that2, ok := that.(MsgFaucetKey)
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
	if this.Sender != that1.Sender {
		return false
	}
	if this.Armor != that1.Armor {
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
	// Mint defines a method for minting coins
	Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error)
	// FaucetKey defines a method for publishing a faucet key
	FaucetKey(ctx context.Context, in *MsgFaucetKey, opts ...grpc.CallOption) (*MsgFaucetKeyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error) {
	out := new(MsgMintResponse)
	err := c.cc.Invoke(ctx, "/stargaze.faucet.v1beta1.Msg/Mint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FaucetKey(ctx context.Context, in *MsgFaucetKey, opts ...grpc.CallOption) (*MsgFaucetKeyResponse, error) {
	out := new(MsgFaucetKeyResponse)
	err := c.cc.Invoke(ctx, "/stargaze.faucet.v1beta1.Msg/FaucetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Mint defines a method for minting coins
	Mint(context.Context, *MsgMint) (*MsgMintResponse, error)
	// FaucetKey defines a method for publishing a faucet key
	FaucetKey(context.Context, *MsgFaucetKey) (*MsgFaucetKeyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Mint(ctx context.Context, req *MsgMint) (*MsgMintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (*UnimplementedMsgServer) FaucetKey(ctx context.Context, req *MsgFaucetKey) (*MsgFaucetKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FaucetKey not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stargaze.faucet.v1beta1.Msg/Mint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Mint(ctx, req.(*MsgMint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FaucetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFaucetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FaucetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stargaze.faucet.v1beta1.Msg/FaucetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FaucetKey(ctx, req.(*MsgFaucetKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stargaze.faucet.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mint",
			Handler:    _Msg_Mint_Handler,
		},
		{
			MethodName: "FaucetKey",
			Handler:    _Msg_FaucetKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stargaze/faucet/v1beta1/tx.proto",
}

func (m *MsgMintResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgFaucetKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFaucetKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFaucetKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if m.Time != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFaucetKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFaucetKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFaucetKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Armor) > 0 {
		i -= len(m.Armor)
		copy(dAtA[i:], m.Armor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Armor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
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
func (m *MsgMintResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgFaucetKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Time != 0 {
		n += 1 + sovTx(uint64(m.Time))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFaucetKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Armor)
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
func (m *MsgMintResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMintResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgFaucetKeyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFaucetKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFaucetKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgMint) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
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
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *MsgFaucetKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFaucetKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFaucetKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Armor", wireType)
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
			m.Armor = string(dAtA[iNdEx:postIndex])
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
