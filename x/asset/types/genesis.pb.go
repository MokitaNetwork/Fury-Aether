// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aether/asset/v1beta1/genesis.proto

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
	Assets            []Asset             `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets" yaml:"assets"`
	Pairs             []Pair              `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs" yaml:"pairs"`
	AppData           []AppData           `protobuf:"bytes,3,rep,name=appData,proto3" json:"appData" yaml:"appData"`
	ExtendedPairVault []ExtendedPairVault `protobuf:"bytes,4,rep,name=extendedPairVault,proto3" json:"extendedPairVault" yaml:"extendedPairVault"`
	Params            Params              `protobuf:"bytes,5,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_523709cb81a3898b, []int{0}
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

func init() {
	proto.RegisterType((*GenesisState)(nil), "aether.asset.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("aether/asset/v1beta1/genesis.proto", fileDescriptor_523709cb81a3898b)
}

var fileDescriptor_523709cb81a3898b = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0x87, 0xdb, 0xcb, 0x85, 0x9b, 0x14, 0xae, 0x89, 0x0d, 0x9a, 0x06, 0x75, 0xa8, 0xb3, 0x91,
	0x85, 0xb6, 0x01, 0x77, 0xee, 0x6c, 0xfc, 0x93, 0xe8, 0x02, 0x53, 0x13, 0x17, 0xee, 0x06, 0x99,
	0x94, 0x06, 0xca, 0x4c, 0x3a, 0x83, 0xc0, 0x5b, 0xf8, 0x18, 0xbe, 0x82, 0x6f, 0xc0, 0x92, 0xa5,
	0x2b, 0xa2, 0xe5, 0x0d, 0x7c, 0x02, 0xd3, 0x99, 0x61, 0x03, 0xc3, 0x0e, 0x3a, 0xdf, 0xf9, 0xce,
	0x39, 0xbf, 0x1c, 0x0b, 0x22, 0xcc, 0x7b, 0x38, 0xf5, 0x11, 0x63, 0x98, 0xfb, 0xaf, 0xcd, 0x0e,
	0xe6, 0xa8, 0xe9, 0x47, 0x78, 0x88, 0x59, 0xcc, 0x3c, 0x9a, 0x12, 0x4e, 0xec, 0xaa, 0x64, 0x3c,
	0xc1, 0x78, 0x8a, 0xa9, 0x55, 0x23, 0x12, 0x11, 0x01, 0xf8, 0xf9, 0x2f, 0xc9, 0xd6, 0x5c, 0xad,
	0x4f, 0x56, 0x4a, 0xa2, 0xae, 0x25, 0x28, 0x8a, 0x53, 0x05, 0x00, 0xbd, 0x82, 0x52, 0xf5, 0x7e,
	0xaa, 0x7d, 0xc7, 0x13, 0x8e, 0x87, 0x5d, 0xdc, 0x7d, 0x40, 0x71, 0xfa, 0x84, 0x46, 0x83, 0x55,
	0xbb, 0xe3, 0x2d, 0xed, 0x52, 0x94, 0xa8, 0xfd, 0xe0, 0x47, 0xc1, 0xaa, 0xdc, 0xca, 0x8d, 0x1f,
	0x39, 0xe2, 0xd8, 0xbe, 0xb3, 0x4a, 0x02, 0x67, 0x8e, 0xe9, 0x16, 0x1a, 0xe5, 0xd6, 0x81, 0xa7,
	0x4b, 0xc0, 0xbb, 0xcc, 0xff, 0x05, 0x7b, 0xb3, 0x45, 0xdd, 0xf8, 0x59, 0xd4, 0xff, 0x4f, 0x51,
	0x32, 0xb8, 0x80, 0xb2, 0x10, 0x86, 0xca, 0x60, 0xdf, 0x58, 0xc5, 0x7c, 0x37, 0xe6, 0xfc, 0x11,
	0xaa, 0x9a, 0x5e, 0x95, 0x4f, 0x1d, 0x54, 0x95, 0xa9, 0x22, 0x4d, 0xa2, 0x0c, 0x86, 0xb2, 0xdc,
	0x6e, 0x5b, 0xff, 0x10, 0xa5, 0x57, 0x88, 0x23, 0xa7, 0x20, 0x4c, 0x47, 0x5b, 0x86, 0x92, 0x50,
	0xb0, 0xaf, 0x64, 0x3b, 0x6a, 0x2c, 0xf9, 0x19, 0x86, 0x2b, 0x8b, 0x3d, 0xb6, 0x76, 0x37, 0x32,
	0x73, 0xfe, 0x0a, 0xf5, 0x89, 0x5e, 0x7d, 0xbd, 0x8e, 0x07, 0xae, 0x6a, 0xe2, 0xc8, 0x26, 0x1b,
	0x3e, 0x18, 0x6e, 0xf6, 0xb0, 0xef, 0xad, 0x92, 0x8c, 0xdf, 0x29, 0xba, 0x66, 0xa3, 0xdc, 0x3a,
	0xdc, 0x16, 0x49, 0xce, 0xac, 0xc7, 0x2b, 0x2b, 0x61, 0xa8, 0x14, 0x41, 0x7b, 0xf6, 0x0d, 0x8c,
	0xf7, 0x0c, 0x18, 0xb3, 0x0c, 0x98, 0xf3, 0x0c, 0x98, 0x5f, 0x19, 0x30, 0xdf, 0x96, 0xc0, 0x98,
	0x2f, 0x81, 0xf1, 0xb9, 0x04, 0xc6, 0xf3, 0x59, 0x14, 0xf3, 0xde, 0xa8, 0xe3, 0xbd, 0x90, 0xc4,
	0x4f, 0x48, 0x3f, 0xe6, 0x68, 0x88, 0xf9, 0x98, 0xa4, 0x7d, 0x5f, 0x5d, 0xc6, 0x44, 0xdd, 0x06,
	0x9f, 0x52, 0xcc, 0x3a, 0x25, 0x71, 0x13, 0xe7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x42,
	0x31, 0x92, 0x19, 0x03, 0x00, 0x00,
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
	dAtA[i] = 0x2a
	if len(m.ExtendedPairVault) > 0 {
		for iNdEx := len(m.ExtendedPairVault) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExtendedPairVault[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AppData) > 0 {
		for iNdEx := len(m.AppData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AppData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Pairs) > 0 {
		for iNdEx := len(m.Pairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Pairs) > 0 {
		for _, e := range m.Pairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AppData) > 0 {
		for _, e := range m.AppData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ExtendedPairVault) > 0 {
		for _, e := range m.ExtendedPairVault {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
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
			m.Assets = append(m.Assets, Asset{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pairs", wireType)
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
			m.Pairs = append(m.Pairs, Pair{})
			if err := m.Pairs[len(m.Pairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppData", wireType)
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
			m.AppData = append(m.AppData, AppData{})
			if err := m.AppData[len(m.AppData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtendedPairVault", wireType)
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
			m.ExtendedPairVault = append(m.ExtendedPairVault, ExtendedPairVault{})
			if err := m.ExtendedPairVault[len(m.ExtendedPairVault)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
