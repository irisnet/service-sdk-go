// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crypto/multisig/keys.proto

package multisig

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/irisnet/service-sdk-go/codec/types"
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

// LegacyAminoPubKey specifies a public key type
// which nests multiple public keys and a threshold,
// it uses legacy amino address rules.
type LegacyAminoPubKey struct {
	Threshold uint32       `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty" yaml:"threshold"`
	PubKeys   []*types.Any `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty" yaml:"pubkeys"`
}

func (m *LegacyAminoPubKey) Reset()         { *m = LegacyAminoPubKey{} }
func (m *LegacyAminoPubKey) String() string { return proto.CompactTextString(m) }
func (*LegacyAminoPubKey) ProtoMessage()    {}
func (*LegacyAminoPubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b57537e097d47d, []int{0}
}
func (m *LegacyAminoPubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyAminoPubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyAminoPubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyAminoPubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyAminoPubKey.Merge(m, src)
}
func (m *LegacyAminoPubKey) XXX_Size() int {
	return m.Size()
}
func (m *LegacyAminoPubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyAminoPubKey.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyAminoPubKey proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LegacyAminoPubKey)(nil), "cosmos.crypto.multisig.LegacyAminoPubKey")
}

func init() { proto.RegisterFile("cosmos/crypto/multisig/keys.proto", fileDescriptor_46b57537e097d47d) }

var fileDescriptor_46b57537e097d47d = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4a, 0x03, 0x31,
	0x1c, 0xc6, 0x2f, 0x2a, 0x8a, 0x57, 0x14, 0x2d, 0x45, 0x6a, 0xc1, 0x5c, 0xbd, 0xa9, 0x4b, 0x13,
	0xa8, 0xe0, 0xd0, 0xad, 0x5d, 0x75, 0x28, 0xc5, 0xc9, 0x45, 0x9a, 0x34, 0xa6, 0xa1, 0x77, 0xf7,
	0x3f, 0x2e, 0x89, 0x90, 0x37, 0x70, 0xf4, 0x11, 0x04, 0x5f, 0xc6, 0xb1, 0xa3, 0x53, 0x91, 0xeb,
	0x1b, 0xf4, 0x09, 0xe4, 0x1a, 0xae, 0x6e, 0x09, 0xdf, 0xef, 0xff, 0x7d, 0x1f, 0x5f, 0x78, 0xcb,
	0x41, 0xa7, 0xa0, 0x29, 0x2f, 0x5c, 0x6e, 0x80, 0xa6, 0x36, 0x31, 0x4a, 0x2b, 0x49, 0x97, 0xc2,
	0x69, 0x92, 0x17, 0x60, 0xa0, 0x79, 0xe5, 0x11, 0xe2, 0x11, 0x52, 0x23, 0x9d, 0x96, 0x04, 0x09,
	0x3b, 0x84, 0x56, 0x2f, 0x4f, 0x77, 0xae, 0x25, 0x80, 0x4c, 0x04, 0xdd, 0xfd, 0x98, 0x7d, 0xa5,
	0xb3, 0xcc, 0x79, 0x29, 0xfe, 0x42, 0xe1, 0xe5, 0xa3, 0x90, 0x33, 0xee, 0x46, 0xa9, 0xca, 0x60,
	0x62, 0xd9, 0x83, 0x70, 0xcd, 0x41, 0x78, 0x6a, 0x16, 0x85, 0xd0, 0x0b, 0x48, 0xe6, 0x6d, 0xd4,
	0x45, 0xbd, 0xb3, 0x71, 0x6b, 0xbb, 0x8e, 0x2e, 0xdc, 0x2c, 0x4d, 0x86, 0xf1, 0x5e, 0x8a, 0xa7,
	0xff, 0x58, 0xf3, 0x29, 0x6c, 0xe4, 0x96, 0x25, 0x8a, 0xbf, 0x54, 0x3d, 0xdb, 0x07, 0xdd, 0xc3,
	0x5e, 0x63, 0xd0, 0x22, 0x3e, 0x9a, 0xd4, 0xd1, 0x64, 0x94, 0xb9, 0xf1, 0x4d, 0xb9, 0x8e, 0x4e,
	0x7c, 0x94, 0xde, 0xae, 0xa3, 0x73, 0x6f, 0x9b, 0x5b, 0x56, 0x5d, 0xc6, 0xd3, 0xd0, 0xfb, 0x54,
	0xea, 0xf0, 0xe8, 0xfd, 0x33, 0x0a, 0xc6, 0x93, 0xef, 0x12, 0xa3, 0x55, 0x89, 0xd1, 0x6f, 0x89,
	0xd1, 0xc7, 0x06, 0x07, 0xab, 0x0d, 0x0e, 0x7e, 0x36, 0x38, 0x78, 0xbe, 0x97, 0xca, 0x2c, 0x2c,
	0x23, 0x1c, 0x52, 0xaa, 0x0a, 0xa5, 0x33, 0x61, 0xa8, 0x16, 0xc5, 0x9b, 0xe2, 0xa2, 0xaf, 0xe7,
	0xcb, 0xbe, 0x84, 0x7a, 0xc6, 0xca, 0x7b, 0xbf, 0x25, 0x3b, 0xde, 0x15, 0xba, 0xfb, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0x87, 0x4f, 0x76, 0x6c, 0x01, 0x00, 0x00,
}

func (m *LegacyAminoPubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyAminoPubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyAminoPubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKeys) > 0 {
		for iNdEx := len(m.PubKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PubKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintKeys(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Threshold != 0 {
		i = encodeVarintKeys(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LegacyAminoPubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Threshold != 0 {
		n += 1 + sovKeys(uint64(m.Threshold))
	}
	if len(m.PubKeys) > 0 {
		for _, e := range m.PubKeys {
			l = e.Size()
			n += 1 + l + sovKeys(uint64(l))
		}
	}
	return n
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeys(x uint64) (n int) {
	return sovKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LegacyAminoPubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: LegacyAminoPubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyAminoPubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeys = append(m.PubKeys, &types.Any{})
			if err := m.PubKeys[len(m.PubKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeys
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
func skipKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
				return 0, ErrInvalidLengthKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeys = fmt.Errorf("proto: unexpected end of group")
)
