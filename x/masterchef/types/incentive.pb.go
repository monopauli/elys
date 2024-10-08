// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/masterchef/incentive.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Incentive Info
type LegacyIncentiveInfo struct {
	// reward amount in eden for 1 year
	EdenAmountPerYear github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=eden_amount_per_year,json=edenAmountPerYear,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"eden_amount_per_year"`
	// starting block height of the distribution
	DistributionStartBlock github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=distribution_start_block,json=distributionStartBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"distribution_start_block"`
	// distribution duration - block number per year
	TotalBlocksPerYear github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=total_blocks_per_year,json=totalBlocksPerYear,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_blocks_per_year"`
	// blocks distributed
	BlocksDistributed github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=blocks_distributed,json=blocksDistributed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"blocks_distributed"`
}

func (m *LegacyIncentiveInfo) Reset()         { *m = LegacyIncentiveInfo{} }
func (m *LegacyIncentiveInfo) String() string { return proto.CompactTextString(m) }
func (*LegacyIncentiveInfo) ProtoMessage()    {}
func (*LegacyIncentiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbfbd92cca5ccf94, []int{0}
}
func (m *LegacyIncentiveInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyIncentiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyIncentiveInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyIncentiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyIncentiveInfo.Merge(m, src)
}
func (m *LegacyIncentiveInfo) XXX_Size() int {
	return m.Size()
}
func (m *LegacyIncentiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyIncentiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyIncentiveInfo proto.InternalMessageInfo

type IncentiveInfo struct {
	// reward amount in eden for 1 year
	EdenAmountPerYear github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=eden_amount_per_year,json=edenAmountPerYear,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"eden_amount_per_year"`
	// blocks distributed
	BlocksDistributed int64 `protobuf:"varint,2,opt,name=blocks_distributed,json=blocksDistributed,proto3" json:"blocks_distributed,omitempty"`
}

func (m *IncentiveInfo) Reset()         { *m = IncentiveInfo{} }
func (m *IncentiveInfo) String() string { return proto.CompactTextString(m) }
func (*IncentiveInfo) ProtoMessage()    {}
func (*IncentiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbfbd92cca5ccf94, []int{1}
}
func (m *IncentiveInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveInfo.Merge(m, src)
}
func (m *IncentiveInfo) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveInfo proto.InternalMessageInfo

func (m *IncentiveInfo) GetBlocksDistributed() int64 {
	if m != nil {
		return m.BlocksDistributed
	}
	return 0
}

func init() {
	proto.RegisterType((*LegacyIncentiveInfo)(nil), "elys.masterchef.LegacyIncentiveInfo")
	proto.RegisterType((*IncentiveInfo)(nil), "elys.masterchef.IncentiveInfo")
}

func init() { proto.RegisterFile("elys/masterchef/incentive.proto", fileDescriptor_fbfbd92cca5ccf94) }

var fileDescriptor_fbfbd92cca5ccf94 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x93, 0xdb, 0xcb, 0x85, 0x1b, 0xb8, 0x5c, 0x8c, 0x55, 0x42, 0x17, 0x89, 0x74, 0x21,
	0x6e, 0x9a, 0x59, 0xf8, 0x04, 0x16, 0x37, 0x01, 0x17, 0xa2, 0x2b, 0x05, 0x09, 0x93, 0xe4, 0x34,
	0x0d, 0x4d, 0xe6, 0x84, 0x99, 0x13, 0x35, 0x6f, 0xe1, 0x1b, 0x08, 0x3e, 0x4d, 0x97, 0x5d, 0x8a,
	0x8b, 0x22, 0xed, 0x8b, 0x48, 0x26, 0xad, 0xed, 0xa2, 0xab, 0x2e, 0x5c, 0xcd, 0x1c, 0xe6, 0xe7,
	0x3b, 0xff, 0x99, 0xf3, 0x5b, 0x1e, 0xe4, 0xb5, 0x62, 0x05, 0x57, 0x04, 0x32, 0x1e, 0xc3, 0x88,
	0x65, 0x22, 0x06, 0x41, 0xd9, 0x23, 0xf8, 0xa5, 0x44, 0x42, 0xfb, 0x7f, 0x23, 0xf0, 0x37, 0x82,
	0x5e, 0x37, 0xc5, 0x14, 0xf5, 0x1b, 0x6b, 0x6e, 0xad, 0xac, 0xe7, 0xa5, 0x88, 0x69, 0x0e, 0x4c,
	0x57, 0x51, 0x35, 0x62, 0x94, 0x15, 0xa0, 0x88, 0x17, 0x65, 0x2b, 0xe8, 0xbf, 0x75, 0xac, 0xc3,
	0x2b, 0x48, 0x79, 0x5c, 0x07, 0xeb, 0x0e, 0x81, 0x18, 0xa1, 0x1d, 0x5a, 0x5d, 0x48, 0x40, 0x84,
	0xbc, 0xc0, 0x4a, 0x50, 0x58, 0x82, 0x0c, 0x6b, 0xe0, 0xd2, 0x31, 0x4f, 0xcc, 0xb3, 0xbf, 0x43,
	0x7f, 0x3a, 0xf7, 0x8c, 0x8f, 0xb9, 0x77, 0x9a, 0x66, 0x34, 0xae, 0x22, 0x3f, 0xc6, 0x82, 0xc5,
	0xa8, 0x0a, 0x54, 0xab, 0x63, 0xa0, 0x92, 0x09, 0xa3, 0xba, 0x04, 0xe5, 0x07, 0x82, 0x6e, 0x0e,
	0x1a, 0xd6, 0x85, 0x46, 0x5d, 0x83, 0xbc, 0x03, 0x2e, 0xed, 0xb1, 0xe5, 0x24, 0x99, 0x22, 0x99,
	0x45, 0x15, 0x65, 0x28, 0x42, 0x45, 0x5c, 0x52, 0x18, 0xe5, 0x18, 0x4f, 0x9c, 0x5f, 0x7b, 0x35,
	0x39, 0xde, 0xe6, 0xdd, 0x36, 0xb8, 0x61, 0x43, 0xb3, 0xb9, 0x75, 0x44, 0x48, 0x3c, 0x6f, 0xe1,
	0x6a, 0x33, 0x4b, 0x67, 0xaf, 0x36, 0xb6, 0x86, 0x69, 0xb4, 0x5a, 0x0f, 0xf3, 0x60, 0xd9, 0x2b,
	0xf8, 0xb7, 0x07, 0x48, 0x9c, 0xdf, 0xfb, 0xfd, 0x55, 0x4b, 0xba, 0xdc, 0x80, 0xfa, 0xaf, 0xa6,
	0xf5, 0xef, 0x87, 0xd7, 0x33, 0xd8, 0x39, 0x51, 0xb3, 0x98, 0xce, 0x0e, 0x87, 0xc3, 0x60, 0xba,
	0x70, 0xcd, 0xd9, 0xc2, 0x35, 0x3f, 0x17, 0xae, 0xf9, 0xb2, 0x74, 0x8d, 0xd9, 0xd2, 0x35, 0xde,
	0x97, 0xae, 0x71, 0xcf, 0xb6, 0x3c, 0x34, 0x99, 0x1d, 0x08, 0xa0, 0x27, 0x94, 0x13, 0x5d, 0xb0,
	0xe7, 0xed, 0x8c, 0x6b, 0x43, 0xd1, 0x1f, 0x1d, 0xcc, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0xd1, 0x4b, 0x41, 0x03, 0x03, 0x00, 0x00,
}

func (m *LegacyIncentiveInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyIncentiveInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyIncentiveInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BlocksDistributed.Size()
		i -= size
		if _, err := m.BlocksDistributed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TotalBlocksPerYear.Size()
		i -= size
		if _, err := m.TotalBlocksPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DistributionStartBlock.Size()
		i -= size
		if _, err := m.DistributionStartBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.EdenAmountPerYear.Size()
		i -= size
		if _, err := m.EdenAmountPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IncentiveInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlocksDistributed != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.BlocksDistributed))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.EdenAmountPerYear.Size()
		i -= size
		if _, err := m.EdenAmountPerYear.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LegacyIncentiveInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EdenAmountPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.DistributionStartBlock.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.TotalBlocksPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.BlocksDistributed.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func (m *IncentiveInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EdenAmountPerYear.Size()
	n += 1 + l + sovIncentive(uint64(l))
	if m.BlocksDistributed != 0 {
		n += 1 + sovIncentive(uint64(m.BlocksDistributed))
	}
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LegacyIncentiveInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: LegacyIncentiveInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyIncentiveInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenAmountPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenAmountPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionStartBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionStartBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBlocksPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBlocksPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksDistributed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlocksDistributed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *IncentiveInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentiveInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdenAmountPerYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EdenAmountPerYear.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksDistributed", wireType)
			}
			m.BlocksDistributed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksDistributed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
