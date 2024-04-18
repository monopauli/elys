// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/incentive/dex_rewards_traker.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// DexRewardsTracker is used for tracking rewards for stakers and LPs, all
// amount here is in USDC
type DexRewardsTracker struct {
	// Number of blocks since start of epoch (distribution epoch)
	NumBlocks github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=num_blocks,json=numBlocks,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"num_blocks"`
	// Accumulated amount at distribution epoch - recalculated at every
	// distribution epoch
	Amount github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"amount"`
	// Accumulated rewards tracked by other (when it's for staking, from lp, if
	// it's for lp, from staking)
	AmountCollectedByOtherTracker github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=amount_collected_by_other_tracker,json=amountCollectedByOtherTracker,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"amount_collected_by_other_tracker"`
}

func (m *DexRewardsTracker) Reset()         { *m = DexRewardsTracker{} }
func (m *DexRewardsTracker) String() string { return proto.CompactTextString(m) }
func (*DexRewardsTracker) ProtoMessage()    {}
func (*DexRewardsTracker) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a7f3218c9553b55, []int{0}
}
func (m *DexRewardsTracker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DexRewardsTracker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DexRewardsTracker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DexRewardsTracker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DexRewardsTracker.Merge(m, src)
}
func (m *DexRewardsTracker) XXX_Size() int {
	return m.Size()
}
func (m *DexRewardsTracker) XXX_DiscardUnknown() {
	xxx_messageInfo_DexRewardsTracker.DiscardUnknown(m)
}

var xxx_messageInfo_DexRewardsTracker proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DexRewardsTracker)(nil), "elys.incentive.DexRewardsTracker")
}

func init() {
	proto.RegisterFile("elys/incentive/dex_rewards_traker.proto", fileDescriptor_6a7f3218c9553b55)
}

var fileDescriptor_6a7f3218c9553b55 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcd, 0xa9, 0x2c,
	0xd6, 0xcf, 0xcc, 0x4b, 0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0x4f, 0x49, 0xad, 0x88, 0x2f,
	0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0x29, 0x8e, 0x2f, 0x29, 0x4a, 0xcc, 0x4e, 0x2d, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x29, 0xd4, 0x83, 0x2b, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x4b, 0xe9, 0x83, 0x58, 0x10, 0x55, 0x4a, 0x4b, 0x98, 0xb8, 0x04, 0x5d, 0x52, 0x2b, 0x82,
	0x20, 0x26, 0x84, 0x14, 0x25, 0x26, 0x67, 0xa7, 0x16, 0x09, 0xf9, 0x72, 0x71, 0xe5, 0x95, 0xe6,
	0xc6, 0x27, 0xe5, 0xe4, 0x27, 0x67, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xe9, 0x9d,
	0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x5a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6, 0x2f,
	0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0xf3, 0xcc, 0x2b, 0x09, 0xe2, 0xcc, 0x2b, 0xcd, 0x75, 0x02, 0x1b,
	0x20, 0xe4, 0xc6, 0xc5, 0x96, 0x98, 0x9b, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0x44, 0xb2, 0x51, 0x2e,
	0xa9, 0xc9, 0x41, 0x50, 0xdd, 0x42, 0x15, 0x5c, 0x8a, 0x10, 0x56, 0x7c, 0x72, 0x7e, 0x4e, 0x4e,
	0x6a, 0x72, 0x49, 0x6a, 0x4a, 0x7c, 0x52, 0x65, 0x7c, 0x7e, 0x49, 0x46, 0x6a, 0x11, 0xc8, 0xf3,
	0x20, 0xb7, 0x4b, 0x30, 0x93, 0x65, 0x85, 0x2c, 0xc4, 0x60, 0x67, 0x98, 0xb9, 0x4e, 0x95, 0xfe,
	0x20, 0x53, 0xa1, 0x01, 0xe2, 0xe4, 0xb3, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0x21, 0x59, 0x02, 0x0a, 0x75, 0xdd, 0xbc, 0xd4, 0x92, 0xf2,
	0xfc, 0xa2, 0x6c, 0x30, 0x47, 0xbf, 0x02, 0x29, 0xb6, 0xc0, 0x16, 0x26, 0xb1, 0x81, 0xc3, 0xde,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xd6, 0xd4, 0xd8, 0xcc, 0x01, 0x00, 0x00,
}

func (this *DexRewardsTracker) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DexRewardsTracker)
	if !ok {
		that2, ok := that.(DexRewardsTracker)
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
	if !this.NumBlocks.Equal(that1.NumBlocks) {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	if !this.AmountCollectedByOtherTracker.Equal(that1.AmountCollectedByOtherTracker) {
		return false
	}
	return true
}
func (m *DexRewardsTracker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DexRewardsTracker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DexRewardsTracker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AmountCollectedByOtherTracker.Size()
		i -= size
		if _, err := m.AmountCollectedByOtherTracker.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDexRewardsTraker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDexRewardsTraker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.NumBlocks.Size()
		i -= size
		if _, err := m.NumBlocks.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDexRewardsTraker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDexRewardsTraker(dAtA []byte, offset int, v uint64) int {
	offset -= sovDexRewardsTraker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DexRewardsTracker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NumBlocks.Size()
	n += 1 + l + sovDexRewardsTraker(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovDexRewardsTraker(uint64(l))
	l = m.AmountCollectedByOtherTracker.Size()
	n += 1 + l + sovDexRewardsTraker(uint64(l))
	return n
}

func sovDexRewardsTraker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDexRewardsTraker(x uint64) (n int) {
	return sovDexRewardsTraker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DexRewardsTracker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDexRewardsTraker
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
			return fmt.Errorf("proto: DexRewardsTracker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DexRewardsTracker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDexRewardsTraker
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
				return ErrInvalidLengthDexRewardsTraker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDexRewardsTraker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumBlocks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDexRewardsTraker
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
				return ErrInvalidLengthDexRewardsTraker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDexRewardsTraker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountCollectedByOtherTracker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDexRewardsTraker
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
				return ErrInvalidLengthDexRewardsTraker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDexRewardsTraker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountCollectedByOtherTracker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDexRewardsTraker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDexRewardsTraker
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
func skipDexRewardsTraker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDexRewardsTraker
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
					return 0, ErrIntOverflowDexRewardsTraker
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
					return 0, ErrIntOverflowDexRewardsTraker
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
				return 0, ErrInvalidLengthDexRewardsTraker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDexRewardsTraker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDexRewardsTraker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDexRewardsTraker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDexRewardsTraker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDexRewardsTraker = fmt.Errorf("proto: unexpected end of group")
)
