// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/faucet/v1beta1/faucet.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Mining struct {
	Minter   string     `protobuf:"bytes,1,opt,name=minter,proto3" json:"minter,omitempty" yaml:"minter"`
	LastTime int64      `protobuf:"varint,2,opt,name=last_time,json=lastTime,proto3" json:"last_time" yaml:"last_time"`
	Total    types.Coin `protobuf:"bytes,3,opt,name=total,proto3" json:"total"`
}

func (m *Mining) Reset()         { *m = Mining{} }
func (m *Mining) String() string { return proto.CompactTextString(m) }
func (*Mining) ProtoMessage()    {}
func (*Mining) Descriptor() ([]byte, []int) {
	return fileDescriptor_33aa9892445cb9d2, []int{0}
}
func (m *Mining) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mining) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mining.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mining) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mining.Merge(m, src)
}
func (m *Mining) XXX_Size() int {
	return m.Size()
}
func (m *Mining) XXX_DiscardUnknown() {
	xxx_messageInfo_Mining.DiscardUnknown(m)
}

var xxx_messageInfo_Mining proto.InternalMessageInfo

func (m *Mining) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

func (m *Mining) GetLastTime() int64 {
	if m != nil {
		return m.LastTime
	}
	return 0
}

func (m *Mining) GetTotal() types.Coin {
	if m != nil {
		return m.Total
	}
	return types.Coin{}
}

type FaucetKey struct {
	Armor string `protobuf:"bytes,1,opt,name=armor,proto3" json:"armor,omitempty"`
}

func (m *FaucetKey) Reset()         { *m = FaucetKey{} }
func (m *FaucetKey) String() string { return proto.CompactTextString(m) }
func (*FaucetKey) ProtoMessage()    {}
func (*FaucetKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_33aa9892445cb9d2, []int{1}
}
func (m *FaucetKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaucetKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaucetKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaucetKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaucetKey.Merge(m, src)
}
func (m *FaucetKey) XXX_Size() int {
	return m.Size()
}
func (m *FaucetKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FaucetKey.DiscardUnknown(m)
}

var xxx_messageInfo_FaucetKey proto.InternalMessageInfo

func (m *FaucetKey) GetArmor() string {
	if m != nil {
		return m.Armor
	}
	return ""
}

func init() {
	proto.RegisterType((*Mining)(nil), "stargaze.faucet.v1beta1.Mining")
	proto.RegisterType((*FaucetKey)(nil), "stargaze.faucet.v1beta1.FaucetKey")
}

func init() {
	proto.RegisterFile("stargaze/faucet/v1beta1/faucet.proto", fileDescriptor_33aa9892445cb9d2)
}

var fileDescriptor_33aa9892445cb9d2 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0xaf, 0x5f, 0xab, 0xd6, 0x08, 0x09, 0xa2, 0x4a, 0x94, 0x0e, 0x4e, 0x89, 0x18,
	0xca, 0xd0, 0x58, 0x05, 0xb1, 0x74, 0x0c, 0x12, 0x03, 0x7f, 0x96, 0xa8, 0x13, 0x0b, 0x72, 0x22,
	0x13, 0x2c, 0xc5, 0x71, 0x15, 0xdf, 0x02, 0xe5, 0x29, 0x78, 0x1a, 0x9e, 0xa1, 0x63, 0x47, 0xa6,
	0x08, 0xa5, 0x1b, 0x63, 0x9f, 0x00, 0x35, 0x4e, 0xc3, 0xe6, 0x7b, 0xfc, 0xd3, 0xd5, 0xb9, 0xe7,
	0xe0, 0x53, 0x0d, 0x2c, 0x8b, 0xd9, 0x3b, 0xa7, 0x4f, 0x6c, 0x1e, 0x71, 0xa0, 0x2f, 0xe3, 0x90,
	0x03, 0x1b, 0x57, 0xa3, 0x37, 0xcb, 0x14, 0x28, 0xfb, 0x68, 0x47, 0x79, 0x95, 0x5c, 0x51, 0xfd,
	0x6e, 0xac, 0x62, 0x55, 0x32, 0x74, 0xfb, 0x32, 0x78, 0x9f, 0x44, 0x4a, 0x4b, 0xa5, 0x69, 0xc8,
	0x34, 0xaf, 0x17, 0x46, 0x4a, 0xa4, 0xe6, 0xdf, 0xfd, 0x44, 0xb8, 0x75, 0x2f, 0x52, 0x91, 0xc6,
	0xf6, 0x19, 0x6e, 0x49, 0x91, 0x02, 0xcf, 0x7a, 0x68, 0x80, 0x86, 0x1d, 0xff, 0x70, 0x93, 0x3b,
	0xfb, 0x0b, 0x26, 0x93, 0x89, 0x6b, 0x74, 0x37, 0xa8, 0x00, 0xfb, 0x06, 0x77, 0x12, 0xa6, 0xe1,
	0x11, 0x84, 0xe4, 0xbd, 0x7f, 0x03, 0x34, 0x6c, 0xf8, 0xa3, 0x22, 0x77, 0xda, 0x77, 0x4c, 0xc3,
	0x54, 0x48, 0xfe, 0x93, 0x3b, 0x7f, 0xc0, 0x26, 0x77, 0x0e, 0xcc, 0x9a, 0x5a, 0x72, 0x83, 0x76,
	0x52, 0xa1, 0xf6, 0x25, 0x6e, 0x82, 0x02, 0x96, 0xf4, 0x1a, 0x03, 0x34, 0xdc, 0x3b, 0x3f, 0xf6,
	0x8c, 0x63, 0x6f, 0xeb, 0x78, 0x77, 0x9c, 0x77, 0xa5, 0x44, 0xea, 0xff, 0x5f, 0xe6, 0x8e, 0x15,
	0x18, 0xda, 0x3d, 0xc1, 0x9d, 0xeb, 0x32, 0x80, 0x5b, 0xbe, 0xb0, 0xbb, 0xb8, 0xc9, 0x32, 0xa9,
	0x2a, 0xe7, 0x81, 0x19, 0xfc, 0xe9, 0xb2, 0x20, 0x68, 0x55, 0x10, 0xf4, 0x5d, 0x10, 0xf4, 0xb1,
	0x26, 0xd6, 0x6a, 0x4d, 0xac, 0xaf, 0x35, 0xb1, 0x1e, 0x26, 0xb1, 0x80, 0xe7, 0x79, 0xe8, 0x45,
	0x4a, 0xd2, 0xd9, 0x3c, 0x4c, 0x44, 0x34, 0x62, 0xaf, 0x5c, 0x2b, 0xc9, 0x69, 0x5d, 0xc2, 0xdb,
	0xae, 0x86, 0xf2, 0xe4, 0x94, 0x25, 0x14, 0x16, 0x33, 0xae, 0xc3, 0x56, 0x19, 0xdc, 0xc5, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xf0, 0x95, 0x47, 0xaf, 0x01, 0x00, 0x00,
}

func (m *Mining) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mining) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mining) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Total.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFaucet(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.LastTime != 0 {
		i = encodeVarintFaucet(dAtA, i, uint64(m.LastTime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FaucetKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaucetKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FaucetKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Armor) > 0 {
		i -= len(m.Armor)
		copy(dAtA[i:], m.Armor)
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.Armor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFaucet(dAtA []byte, offset int, v uint64) int {
	offset -= sovFaucet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mining) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	if m.LastTime != 0 {
		n += 1 + sovFaucet(uint64(m.LastTime))
	}
	l = m.Total.Size()
	n += 1 + l + sovFaucet(uint64(l))
	return n
}

func (m *FaucetKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Armor)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	return n
}

func sovFaucet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFaucet(x uint64) (n int) {
	return sovFaucet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mining) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFaucet
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
			return fmt.Errorf("proto: Mining: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mining: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
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
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTime", wireType)
			}
			m.LastTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
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
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFaucet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFaucet
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
func (m *FaucetKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFaucet
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
			return fmt.Errorf("proto: FaucetKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaucetKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Armor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
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
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Armor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFaucet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFaucet
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
func skipFaucet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFaucet
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
					return 0, ErrIntOverflowFaucet
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
					return 0, ErrIntOverflowFaucet
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
				return 0, ErrInvalidLengthFaucet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFaucet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFaucet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFaucet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFaucet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFaucet = fmt.Errorf("proto: unexpected end of group")
)
