// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aether/liquidation/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

type GenesisState struct {
	LockedVault     []LockedVault `protobuf:"bytes,1,rep,name=lockedVault,proto3" json:"lockedVault" yaml:"lockedVault"`
	WhitelistedApps []uint64      `protobuf:"varint,2,rep,packed,name=whitelistedApps,proto3" json:"whitelistedApps,omitempty" yaml:"whitelistedApps"`
	Params          Params        `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9c2514f06c02d3, []int{0}
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

func (m *GenesisState) GetLockedVault() []LockedVault {
	if m != nil {
		return m.LockedVault
	}
	return nil
}

func (m *GenesisState) GetWhitelistedApps() []uint64 {
	if m != nil {
		return m.WhitelistedApps
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "aether.liquidation.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("aether/liquidation/v1beta1/genesis.proto", fileDescriptor_1c9c2514f06c02d3)
}

var fileDescriptor_1c9c2514f06c02d3 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0x31, 0x6b, 0xc2, 0x40,
	0x18, 0x06, 0xe0, 0x5c, 0x15, 0x87, 0x58, 0x28, 0x84, 0x52, 0xe4, 0x86, 0x53, 0xb2, 0x98, 0xa5,
	0x77, 0xa8, 0x5b, 0xa7, 0x36, 0x14, 0xba, 0xb4, 0x50, 0x2c, 0x74, 0xe8, 0x52, 0x4e, 0xfd, 0x88,
	0x87, 0x89, 0x97, 0x26, 0x9f, 0x5a, 0xff, 0x45, 0x7f, 0x96, 0xa3, 0x63, 0x27, 0x29, 0xfa, 0x0f,
	0xdc, 0x0b, 0xc5, 0x5c, 0xa0, 0xa9, 0xd0, 0x6c, 0x37, 0x3c, 0xef, 0xcb, 0x7b, 0x9f, 0xed, 0x49,
	0xc0, 0x31, 0x24, 0x22, 0x54, 0x6f, 0x33, 0x35, 0x92, 0xa8, 0xf4, 0x54, 0xcc, 0x3b, 0x03, 0x40,
	0xd9, 0x11, 0x01, 0x4c, 0x21, 0x55, 0x29, 0x8f, 0x13, 0x8d, 0xda, 0xa1, 0x46, 0xf2, 0x82, 0xe4,
	0xb9, 0xa4, 0xe7, 0x81, 0x0e, 0x74, 0xc6, 0xc4, 0xe1, 0x65, 0x12, 0xb4, 0x5d, 0xd2, 0x1d, 0xcb,
	0x44, 0x46, 0x79, 0x35, 0xbd, 0x2c, 0x81, 0xa1, 0x1e, 0x4e, 0x60, 0xf4, 0x3a, 0x97, 0xb3, 0x10,
	0x0d, 0x77, 0xbf, 0x89, 0x7d, 0x7a, 0x67, 0xb6, 0x3d, 0xa1, 0x44, 0x70, 0xc0, 0xae, 0x1b, 0xf6,
	0x7c, 0x50, 0x0d, 0xd2, 0xaa, 0x78, 0xf5, 0x6e, 0x9b, 0xff, 0x3f, 0x98, 0xdf, 0xff, 0x72, 0x9f,
	0xae, 0x36, 0x4d, 0x6b, 0xbf, 0x69, 0x3a, 0x4b, 0x19, 0x85, 0x57, 0x6e, 0xa1, 0xc9, 0xed, 0x17,
	0x7b, 0x9d, 0x5b, 0xfb, 0x6c, 0x31, 0x56, 0x08, 0xa1, 0x4a, 0x11, 0x46, 0x37, 0x71, 0x9c, 0x36,
	0x4e, 0x5a, 0x15, 0xaf, 0xea, 0xd3, 0xfd, 0xa6, 0x79, 0x61, 0xd2, 0x47, 0xc0, 0xed, 0x1f, 0x47,
	0x9c, 0x6b, 0xbb, 0x66, 0x3e, 0xdf, 0xa8, 0xb6, 0x88, 0x57, 0xef, 0xba, 0x65, 0x3b, 0x1f, 0x33,
	0xe9, 0x57, 0x0f, 0x13, 0xfb, 0x79, 0xce, 0x7f, 0x58, 0x6d, 0x19, 0x59, 0x6f, 0x19, 0xf9, 0xda,
	0x32, 0xf2, 0xb1, 0x63, 0xd6, 0x7a, 0xc7, 0xac, 0xcf, 0x1d, 0xb3, 0x5e, 0x7a, 0x81, 0xc2, 0xf1,
	0x6c, 0xc0, 0x87, 0x3a, 0x12, 0x91, 0x9e, 0x28, 0x94, 0x53, 0xc0, 0x85, 0x4e, 0x26, 0x22, 0xbf,
	0xf0, 0xfb, 0x9f, 0x1b, 0xe3, 0x32, 0x86, 0x74, 0x50, 0xcb, 0xae, 0xda, 0xfb, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xd1, 0x78, 0x45, 0x7a, 0x0b, 0x02, 0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.WhitelistedApps) > 0 {
		dAtA3 := make([]byte, len(m.WhitelistedApps)*10)
		var j2 int
		for _, num := range m.WhitelistedApps {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintGenesis(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x12
	}
	if len(m.LockedVault) > 0 {
		for iNdEx := len(m.LockedVault) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockedVault[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
	if len(m.LockedVault) > 0 {
		for _, e := range m.LockedVault {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WhitelistedApps) > 0 {
		l = 0
		for _, e := range m.WhitelistedApps {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field LockedVault", wireType)
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
			m.LockedVault = append(m.LockedVault, LockedVault{})
			if err := m.LockedVault[len(m.LockedVault)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.WhitelistedApps = append(m.WhitelistedApps, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.WhitelistedApps) == 0 {
					m.WhitelistedApps = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.WhitelistedApps = append(m.WhitelistedApps, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedApps", wireType)
			}
		case 4:
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
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
