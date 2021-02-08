// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/genesis.proto

package token

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/irisnet/service-sdk-go/types"
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

// GenesisState defines the token module's genesis state.
type GenesisState struct {
	Params      Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Tokens      []Token      `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens"`
	BurnedCoins []types.Coin `protobuf:"bytes,3,rep,name=burned_coins,json=burnedCoins,proto3" json:"burned_coins"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d637ba3268cd6c3, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTokens() []Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *GenesisState) GetBurnedCoins() []types.Coin {
	if m != nil {
		return m.BurnedCoins
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "irismod.token.GenesisState")
}

func init() { proto.RegisterFile("token/genesis.proto", fileDescriptor_7d637ba3268cd6c3) }

var fileDescriptor_7d637ba3268cd6c3 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3d, 0x4e, 0xc3, 0x40,
	0x10, 0x85, 0xbd, 0x04, 0xa5, 0x70, 0x42, 0x81, 0x09, 0x92, 0x49, 0xb1, 0x44, 0x54, 0x69, 0xb2,
	0xab, 0x38, 0x9c, 0xc0, 0x14, 0x74, 0x08, 0x01, 0x15, 0x4d, 0xe4, 0x9f, 0x91, 0x59, 0x05, 0x7b,
	0x22, 0xcf, 0x3a, 0xe7, 0xe0, 0x3e, 0x5c, 0x20, 0x65, 0x4a, 0x2a, 0x84, 0xec, 0x8b, 0xa0, 0xdd,
	0x75, 0x0a, 0x68, 0x46, 0xa3, 0x79, 0xef, 0xd3, 0x3c, 0x3d, 0xff, 0x42, 0xe3, 0x06, 0x2a, 0x59,
	0x40, 0x05, 0xa4, 0x48, 0x6c, 0x6b, 0xd4, 0x18, 0x9c, 0xa9, 0x5a, 0x51, 0x89, 0xb9, 0xb0, 0xe2,
	0x74, 0x52, 0x60, 0x81, 0x56, 0x91, 0x66, 0x73, 0xa6, 0xe9, 0xb9, 0x23, 0xed, 0xec, 0x4f, 0x3c,
	0x43, 0x2a, 0x91, 0x64, 0x9a, 0x10, 0xc8, 0xdd, 0x32, 0x05, 0x9d, 0x2c, 0x65, 0x86, 0xaa, 0xd7,
	0x6f, 0x3e, 0x99, 0x3f, 0xbe, 0x77, 0x9f, 0x9e, 0x75, 0xa2, 0x21, 0x58, 0xf9, 0xc3, 0x6d, 0x52,
	0x27, 0x25, 0x85, 0x6c, 0xc6, 0xe6, 0xa3, 0xe8, 0x52, 0xfc, 0xf9, 0x2c, 0x1e, 0xad, 0x18, 0x9f,
	0xee, 0xbf, 0xaf, 0xbd, 0xa7, 0xde, 0x1a, 0x44, 0xfe, 0xd0, 0xaa, 0x14, 0x9e, 0xcc, 0x06, 0xf3,
	0x51, 0x34, 0xf9, 0x07, 0xbd, 0x98, 0x79, 0x64, 0x9c, 0x33, 0x88, 0xfd, 0x71, 0xda, 0xd4, 0x15,
	0xe4, 0x6b, 0x13, 0x87, 0xc2, 0x81, 0x25, 0xaf, 0x84, 0x0b, 0x2c, 0x4c, 0x60, 0xd1, 0x07, 0x16,
	0x77, 0xa8, 0x8e, 0xf8, 0xc8, 0x41, 0xe6, 0x42, 0xf1, 0xc3, 0xbe, 0xe5, 0xec, 0xd0, 0x72, 0xf6,
	0xd3, 0x72, 0xf6, 0xd1, 0x71, 0xef, 0xd0, 0x71, 0xef, 0xab, 0xe3, 0xde, 0xeb, 0x6d, 0xa1, 0xf4,
	0x5b, 0x93, 0x8a, 0x0c, 0x4b, 0x69, 0xb2, 0x54, 0xa0, 0x25, 0x41, 0xbd, 0x53, 0x19, 0x2c, 0x28,
	0xdf, 0x2c, 0x0a, 0xb4, 0x95, 0xac, 0x4b, 0xcc, 0x9b, 0x77, 0x20, 0xd7, 0x59, 0x3a, 0xb4, 0xa5,
	0xac, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xf8, 0x81, 0x91, 0x83, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BurnedCoins) > 0 {
		for iNdEx := len(m.BurnedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BurnedCoins) > 0 {
		for _, e := range m.BurnedCoins {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, Token{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnedCoins = append(m.BurnedCoins, types.Coin{})
			if err := m.BurnedCoins[len(m.BurnedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)