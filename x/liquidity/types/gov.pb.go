// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aether/liquidity/v1beta1/gov.proto

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

type UpdateGenericParamsProposal struct {
	AppId       uint64   `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Keys        []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	Values      []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	Title       string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
}

func (m *UpdateGenericParamsProposal) Reset()         { *m = UpdateGenericParamsProposal{} }
func (m *UpdateGenericParamsProposal) String() string { return proto.CompactTextString(m) }
func (*UpdateGenericParamsProposal) ProtoMessage()    {}
func (*UpdateGenericParamsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_117e1f5baeb7b742, []int{0}
}
func (m *UpdateGenericParamsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateGenericParamsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateGenericParamsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateGenericParamsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGenericParamsProposal.Merge(m, src)
}
func (m *UpdateGenericParamsProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateGenericParamsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGenericParamsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGenericParamsProposal proto.InternalMessageInfo

type CreateNewLiquidityPairProposal struct {
	From           string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	AppId          uint64 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	BaseCoinDenom  string `protobuf:"bytes,3,opt,name=base_coin_denom,json=baseCoinDenom,proto3" json:"base_coin_denom,omitempty"`
	QuoteCoinDenom string `protobuf:"bytes,4,opt,name=quote_coin_denom,json=quoteCoinDenom,proto3" json:"quote_coin_denom,omitempty"`
	Title          string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description    string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
}

func (m *CreateNewLiquidityPairProposal) Reset()         { *m = CreateNewLiquidityPairProposal{} }
func (m *CreateNewLiquidityPairProposal) String() string { return proto.CompactTextString(m) }
func (*CreateNewLiquidityPairProposal) ProtoMessage()    {}
func (*CreateNewLiquidityPairProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_117e1f5baeb7b742, []int{1}
}
func (m *CreateNewLiquidityPairProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateNewLiquidityPairProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateNewLiquidityPairProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateNewLiquidityPairProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNewLiquidityPairProposal.Merge(m, src)
}
func (m *CreateNewLiquidityPairProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateNewLiquidityPairProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNewLiquidityPairProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNewLiquidityPairProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateGenericParamsProposal)(nil), "aether.liquidity.v1beta1.UpdateGenericParamsProposal")
	proto.RegisterType((*CreateNewLiquidityPairProposal)(nil), "aether.liquidity.v1beta1.CreateNewLiquidityPairProposal")
}

func init() {
	proto.RegisterFile("aether/liquidity/v1beta1/gov.proto", fileDescriptor_117e1f5baeb7b742)
}

var fileDescriptor_117e1f5baeb7b742 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0xdb, 0x24, 0xd0, 0xf1, 0xdf, 0x32, 0xe8, 0x12, 0x14, 0x66, 0x4b, 0x0e, 0x4b,
	0x2e, 0x26, 0x2c, 0x7b, 0x11, 0x8f, 0xbb, 0x82, 0x08, 0x22, 0x25, 0xb0, 0x17, 0x2f, 0x65, 0x92,
	0x4c, 0xe3, 0x60, 0x92, 0x77, 0x3a, 0x99, 0x54, 0xf3, 0x2d, 0xfc, 0x18, 0x7e, 0x0e, 0x4f, 0x3d,
	0xf6, 0xe8, 0xa9, 0x68, 0xfa, 0x0d, 0xfa, 0x05, 0x94, 0x4e, 0x6a, 0x8d, 0x37, 0xf7, 0xf6, 0xce,
	0xf3, 0xfe, 0xe6, 0xf0, 0x7b, 0x79, 0xb0, 0x9f, 0x42, 0x99, 0xf1, 0xcf, 0x51, 0x21, 0x16, 0x8d,
	0xc8, 0x84, 0x6e, 0xa3, 0xe5, 0x65, 0xc2, 0x35, 0xbb, 0x8c, 0x72, 0x58, 0x86, 0x52, 0x81, 0x06,
	0xe2, 0xf5, 0x4c, 0x78, 0x64, 0xc2, 0x03, 0xf3, 0xf4, 0x71, 0x0e, 0x39, 0x18, 0x28, 0xda, 0x4f,
	0x3d, 0xef, 0x7f, 0x43, 0xf8, 0xd9, 0xad, 0xcc, 0x98, 0xe6, 0xaf, 0x79, 0xc5, 0x95, 0x48, 0xa7,
	0x4c, 0xb1, 0xb2, 0x9e, 0x2a, 0x90, 0x50, 0xb3, 0x82, 0x3c, 0xc1, 0x2e, 0x93, 0x72, 0x26, 0x32,
	0x0f, 0x4d, 0x50, 0x60, 0xc7, 0x0e, 0x93, 0xf2, 0x4d, 0x46, 0x08, 0xb6, 0x3f, 0xf2, 0xb6, 0xf6,
	0x4e, 0x26, 0xa3, 0x60, 0x1c, 0x9b, 0x99, 0x9c, 0x61, 0x77, 0xc9, 0x8a, 0x86, 0xd7, 0xde, 0xc8,
	0xa4, 0x87, 0x17, 0xb9, 0xc0, 0x8e, 0x16, 0xba, 0xe0, 0x9e, 0x3d, 0x41, 0xc1, 0xf8, 0xfa, 0x74,
	0xb7, 0x39, 0xbf, 0xdf, 0xb2, 0xb2, 0x78, 0xe9, 0x9b, 0xd8, 0x8f, 0xfb, 0x35, 0x79, 0x81, 0xef,
	0x65, 0xbc, 0x4e, 0x95, 0x90, 0x5a, 0x40, 0xe5, 0x39, 0x86, 0x3e, 0xdb, 0x6d, 0xce, 0x49, 0x4f,
	0x0f, 0x96, 0x7e, 0x3c, 0x44, 0xfd, 0x5f, 0x08, 0xd3, 0x1b, 0xc5, 0x99, 0xe6, 0xef, 0xf8, 0xa7,
	0xb7, 0x7f, 0xcc, 0xa7, 0x4c, 0xa8, 0xa3, 0x07, 0xc1, 0xf6, 0x5c, 0x41, 0x69, 0x2c, 0xc6, 0xb1,
	0x99, 0x07, 0x6e, 0x27, 0x43, 0xb7, 0x0b, 0xfc, 0x28, 0x61, 0x35, 0x9f, 0xa5, 0x20, 0xaa, 0x59,
	0xc6, 0x2b, 0x28, 0xbd, 0x91, 0xf9, 0xf5, 0x60, 0x1f, 0xdf, 0x80, 0xa8, 0x5e, 0xed, 0x43, 0x12,
	0xe0, 0xd3, 0x45, 0x03, 0xfa, 0x1f, 0xd0, 0x28, 0xc6, 0x0f, 0x4d, 0xfe, 0x97, 0x3c, 0x5e, 0xc0,
	0xb9, 0xd3, 0x05, 0xdc, 0xff, 0xbe, 0xc0, 0xf5, 0xed, 0xea, 0x27, 0xb5, 0xbe, 0x76, 0xd4, 0x5a,
	0x75, 0x14, 0xad, 0x3b, 0x8a, 0x7e, 0x74, 0x14, 0x7d, 0xd9, 0x52, 0x6b, 0xbd, 0xa5, 0xd6, 0xf7,
	0x2d, 0xb5, 0xde, 0x5f, 0xe5, 0x42, 0x7f, 0x68, 0x92, 0x30, 0x85, 0x32, 0xea, 0x3b, 0xf2, 0x1c,
	0xe6, 0x73, 0x91, 0x0a, 0x56, 0x1c, 0xde, 0xd1, 0xb0, 0x59, 0xba, 0x95, 0xbc, 0x4e, 0x5c, 0x53,
	0x92, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0xdd, 0x8a, 0xd7, 0x7a, 0x02, 0x00, 0x00,
}

func (m *UpdateGenericParamsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateGenericParamsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateGenericParamsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarintGov(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintGov(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.AppId != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateNewLiquidityPairProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateNewLiquidityPairProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateNewLiquidityPairProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.QuoteCoinDenom) > 0 {
		i -= len(m.QuoteCoinDenom)
		copy(dAtA[i:], m.QuoteCoinDenom)
		i = encodeVarintGov(dAtA, i, uint64(len(m.QuoteCoinDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BaseCoinDenom) > 0 {
		i -= len(m.BaseCoinDenom)
		copy(dAtA[i:], m.BaseCoinDenom)
		i = encodeVarintGov(dAtA, i, uint64(len(m.BaseCoinDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AppId != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintGov(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateGenericParamsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppId != 0 {
		n += 1 + sovGov(uint64(m.AppId))
	}
	if len(m.Keys) > 0 {
		for _, s := range m.Keys {
			l = len(s)
			n += 1 + l + sovGov(uint64(l))
		}
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *CreateNewLiquidityPairProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.AppId != 0 {
		n += 1 + sovGov(uint64(m.AppId))
	}
	l = len(m.BaseCoinDenom)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.QuoteCoinDenom)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateGenericParamsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateGenericParamsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateGenericParamsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *CreateNewLiquidityPairProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: CreateNewLiquidityPairProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateNewLiquidityPairProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
