// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aether/bandoracle/v1beta1/fetch_price.proto

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

type FetchPriceCallData struct {
	Symbols    []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Multiplier uint64   `protobuf:"varint,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (m *FetchPriceCallData) Reset()         { *m = FetchPriceCallData{} }
func (m *FetchPriceCallData) String() string { return proto.CompactTextString(m) }
func (*FetchPriceCallData) ProtoMessage()    {}
func (*FetchPriceCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_419dc3984543857a, []int{0}
}
func (m *FetchPriceCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPriceCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPriceCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPriceCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPriceCallData.Merge(m, src)
}
func (m *FetchPriceCallData) XXX_Size() int {
	return m.Size()
}
func (m *FetchPriceCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPriceCallData.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPriceCallData proto.InternalMessageInfo

type FetchPriceResult struct {
	Rates []uint64 `protobuf:"varint,1,rep,packed,name=rates,proto3" json:"rates,omitempty"`
}

func (m *FetchPriceResult) Reset()         { *m = FetchPriceResult{} }
func (m *FetchPriceResult) String() string { return proto.CompactTextString(m) }
func (*FetchPriceResult) ProtoMessage()    {}
func (*FetchPriceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_419dc3984543857a, []int{1}
}
func (m *FetchPriceResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPriceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPriceResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPriceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPriceResult.Merge(m, src)
}
func (m *FetchPriceResult) XXX_Size() int {
	return m.Size()
}
func (m *FetchPriceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPriceResult.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPriceResult proto.InternalMessageInfo

type Market struct {
	Symbol   string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	ScriptID uint64 `protobuf:"varint,2,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_419dc3984543857a, []int{2}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

type DiscardData struct {
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty" yaml:"block_height"`
	DiscardBool bool  `protobuf:"varint,2,opt,name=discard_bool,json=discardBool,proto3" json:"discard_bool,omitempty" yaml:"discard_bool"`
}

func (m *DiscardData) Reset()         { *m = DiscardData{} }
func (m *DiscardData) String() string { return proto.CompactTextString(m) }
func (*DiscardData) ProtoMessage()    {}
func (*DiscardData) Descriptor() ([]byte, []int) {
	return fileDescriptor_419dc3984543857a, []int{3}
}
func (m *DiscardData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DiscardData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DiscardData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DiscardData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardData.Merge(m, src)
}
func (m *DiscardData) XXX_Size() int {
	return m.Size()
}
func (m *DiscardData) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardData.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FetchPriceCallData)(nil), "aether.bandoracle.v1beta1.FetchPriceCallData")
	proto.RegisterType((*FetchPriceResult)(nil), "aether.bandoracle.v1beta1.FetchPriceResult")
	proto.RegisterType((*Market)(nil), "aether.bandoracle.v1beta1.Market")
	proto.RegisterType((*DiscardData)(nil), "aether.bandoracle.v1beta1.DiscardData")
}

func init() {
	proto.RegisterFile("aether/bandoracle/v1beta1/fetch_price.proto", fileDescriptor_419dc3984543857a)
}

var fileDescriptor_419dc3984543857a = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x31, 0x8e, 0xd4, 0x30,
	0x14, 0x86, 0x63, 0x76, 0x19, 0x66, 0x3c, 0x20, 0x2d, 0x61, 0x25, 0x06, 0x0a, 0x27, 0x72, 0x15,
	0x84, 0x48, 0xb4, 0x82, 0x6a, 0x25, 0x9a, 0x30, 0x42, 0x6c, 0x01, 0x42, 0x46, 0xa2, 0xa0, 0x89,
	0x1c, 0xc7, 0x93, 0x58, 0xeb, 0xe0, 0xc8, 0xf1, 0x20, 0xe6, 0x00, 0xf4, 0x1c, 0x83, 0xa3, 0x6c,
	0xb9, 0x25, 0x55, 0x04, 0x99, 0x1b, 0xcc, 0x09, 0x50, 0xec, 0xac, 0x36, 0x74, 0xef, 0x7f, 0xef,
	0xff, 0xbf, 0x97, 0xe8, 0x19, 0x3e, 0x67, 0xaa, 0x2e, 0xf8, 0xf7, 0x24, 0xa7, 0x5f, 0x0b, 0xa5,
	0x29, 0x93, 0x3c, 0xf9, 0x76, 0x96, 0x73, 0x43, 0xcf, 0x92, 0x0d, 0x37, 0xac, 0xca, 0x1a, 0x2d,
	0x18, 0x8f, 0x1b, 0xad, 0x8c, 0xf2, 0x9f, 0x38, 0x73, 0x7c, 0x6b, 0x8e, 0x47, 0xf3, 0xd3, 0xd3,
	0x52, 0x95, 0xca, 0xba, 0x92, 0xa1, 0x72, 0x01, 0xfc, 0x01, 0xfa, 0x6f, 0x07, 0xca, 0xc7, 0x01,
	0xf2, 0x86, 0x4a, 0xb9, 0xa6, 0x86, 0xfa, 0x2b, 0x78, 0xaf, 0xdd, 0xd5, 0xb9, 0x92, 0xed, 0x0a,
	0x84, 0x47, 0xd1, 0x82, 0xdc, 0x48, 0x1f, 0x41, 0x58, 0x6f, 0xa5, 0x11, 0x8d, 0x14, 0x5c, 0xaf,
	0xee, 0x84, 0x20, 0x3a, 0x26, 0x93, 0x0e, 0x8e, 0xe0, 0xc9, 0x2d, 0x8f, 0xf0, 0x76, 0x2b, 0x8d,
	0x7f, 0x0a, 0xef, 0x6a, 0x6a, 0xb8, 0x63, 0x1d, 0x13, 0x27, 0xb0, 0x86, 0xb3, 0xf7, 0x54, 0x5f,
	0x72, 0xe3, 0x3f, 0x83, 0x33, 0x87, 0x5f, 0x81, 0x10, 0x44, 0x8b, 0xf4, 0xe1, 0xa1, 0x0b, 0x1e,
	0xec, 0x68, 0x2d, 0xcf, 0xb1, 0xeb, 0x63, 0x32, 0x1a, 0xfc, 0xd7, 0x70, 0xd1, 0x32, 0x2d, 0x1a,
	0x93, 0x89, 0xc2, 0x6d, 0x4f, 0xc3, 0xbe, 0x0b, 0xe6, 0x9f, 0x6c, 0xf3, 0x62, 0x7d, 0xe8, 0x82,
	0x93, 0x31, 0x79, 0x63, 0xc3, 0x64, 0xee, 0xea, 0x8b, 0x02, 0xff, 0x00, 0x70, 0xb9, 0x16, 0x2d,
	0xa3, 0xba, 0xb0, 0xff, 0x79, 0x0e, 0xef, 0xe7, 0x52, 0xb1, 0xcb, 0xac, 0xe2, 0xa2, 0xac, 0x8c,
	0xdd, 0x7f, 0x94, 0x3e, 0x3e, 0x74, 0xc1, 0x23, 0x47, 0x99, 0x4e, 0x31, 0x59, 0x5a, 0xf9, 0xce,
	0xaa, 0x21, 0x5b, 0x38, 0x54, 0x96, 0x2b, 0x25, 0xed, 0xd7, 0xcc, 0xa7, 0xd9, 0xe9, 0x14, 0x93,
	0xe5, 0x28, 0x53, 0xa5, 0x64, 0xfa, 0xf9, 0xea, 0x2f, 0xf2, 0x7e, 0xf5, 0xc8, 0xbb, 0xea, 0x11,
	0xb8, 0xee, 0x11, 0xf8, 0xd3, 0x23, 0xf0, 0x73, 0x8f, 0xbc, 0xeb, 0x3d, 0xf2, 0x7e, 0xef, 0x91,
	0xf7, 0xe5, 0x55, 0x29, 0x4c, 0xb5, 0xcd, 0x63, 0xa6, 0xea, 0xc4, 0xdd, 0xf4, 0x85, 0xda, 0x6c,
	0x04, 0x13, 0x54, 0x8e, 0x3a, 0xf9, 0xef, 0x49, 0x98, 0x5d, 0xc3, 0xdb, 0x7c, 0x66, 0x8f, 0xfa,
	0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x16, 0x4f, 0x0d, 0x34, 0x02, 0x00, 0x00,
}

func (m *FetchPriceCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPriceCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPriceCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintFetchPrice(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintFetchPrice(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FetchPriceResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPriceResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPriceResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rates) > 0 {
		dAtA2 := make([]byte, len(m.Rates)*10)
		var j1 int
		for _, num := range m.Rates {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintFetchPrice(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ScriptID != 0 {
		i = encodeVarintFetchPrice(dAtA, i, uint64(m.ScriptID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintFetchPrice(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DiscardData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiscardData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DiscardData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DiscardBool {
		i--
		if m.DiscardBool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.BlockHeight != 0 {
		i = encodeVarintFetchPrice(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFetchPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovFetchPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FetchPriceCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovFetchPrice(uint64(l))
		}
	}
	if m.Multiplier != 0 {
		n += 1 + sovFetchPrice(uint64(m.Multiplier))
	}
	return n
}

func (m *FetchPriceResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rates) > 0 {
		l = 0
		for _, e := range m.Rates {
			l += sovFetchPrice(uint64(e))
		}
		n += 1 + sovFetchPrice(uint64(l)) + l
	}
	return n
}

func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovFetchPrice(uint64(l))
	}
	if m.ScriptID != 0 {
		n += 1 + sovFetchPrice(uint64(m.ScriptID))
	}
	return n
}

func (m *DiscardData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovFetchPrice(uint64(m.BlockHeight))
	}
	if m.DiscardBool {
		n += 2
	}
	return n
}

func sovFetchPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFetchPrice(x uint64) (n int) {
	return sovFetchPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FetchPriceCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFetchPrice
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
			return fmt.Errorf("proto: FetchPriceCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPriceCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
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
				return ErrInvalidLengthFetchPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFetchPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFetchPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFetchPrice
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
func (m *FetchPriceResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFetchPrice
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
			return fmt.Errorf("proto: FetchPriceResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPriceResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFetchPrice
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
				m.Rates = append(m.Rates, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFetchPrice
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
					return ErrInvalidLengthFetchPrice
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthFetchPrice
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
				if elementCount != 0 && len(m.Rates) == 0 {
					m.Rates = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFetchPrice
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
					m.Rates = append(m.Rates, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFetchPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFetchPrice
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
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFetchPrice
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
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
				return ErrInvalidLengthFetchPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFetchPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptID", wireType)
			}
			m.ScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFetchPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFetchPrice
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
func (m *DiscardData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFetchPrice
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
			return fmt.Errorf("proto: DiscardData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiscardData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiscardBool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DiscardBool = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFetchPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFetchPrice
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
func skipFetchPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFetchPrice
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
					return 0, ErrIntOverflowFetchPrice
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
					return 0, ErrIntOverflowFetchPrice
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
				return 0, ErrInvalidLengthFetchPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFetchPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFetchPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFetchPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFetchPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFetchPrice = fmt.Errorf("proto: unexpected end of group")
)
