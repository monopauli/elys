// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/margin/pool.proto

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

type Pool struct {
	AmmPoolId                    uint64                                 `protobuf:"varint,1,opt,name=ammPoolId,proto3" json:"ammPoolId,omitempty"`
	Health                       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=health,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"health"`
	Enabled                      bool                                   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Closed                       bool                                   `protobuf:"varint,4,opt,name=closed,proto3" json:"closed,omitempty"`
	ExternalLiabilities          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=external_liabilities,json=externalLiabilities,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"external_liabilities"`
	ExternalCustody              github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=external_custody,json=externalCustody,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"external_custody"`
	NativeLiabilities            github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=native_liabilities,json=nativeLiabilities,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"native_liabilities"`
	NativeCustody                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=native_custody,json=nativeCustody,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"native_custody"`
	InterestRate                 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=interest_rate,json=interestRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"interest_rate"`
	NativeAssetBalance           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=native_asset_balance,json=nativeAssetBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"native_asset_balance" yaml:"native_asset_balance"`
	ExternalAssetBalance         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=external_asset_balance,json=externalAssetBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"external_asset_balance" yaml:"external_asset_balance"`
	UnsettledExternalLiabilities github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=unsettled_external_liabilities,json=unsettledExternalLiabilities,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"unsettled_external_liabilities"`
	BlockInterestNative          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=block_interest_native,json=blockInterestNative,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"block_interest_native"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_030dd771f91c2bb4, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetAmmPoolId() uint64 {
	if m != nil {
		return m.AmmPoolId
	}
	return 0
}

func (m *Pool) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Pool) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func init() {
	proto.RegisterType((*Pool)(nil), "elys.margin.Pool")
}

func init() { proto.RegisterFile("elys/margin/pool.proto", fileDescriptor_030dd771f91c2bb4) }

var fileDescriptor_030dd771f91c2bb4 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xc7, 0x1b, 0xad, 0xdd, 0x76, 0x76, 0xeb, 0xcb, 0x58, 0xcb, 0xa0, 0x6b, 0x5a, 0x7a, 0x90,
	0x82, 0x6c, 0x72, 0xf0, 0xe6, 0xcd, 0xea, 0x0a, 0x05, 0xdf, 0x88, 0x78, 0x50, 0x90, 0x30, 0x49,
	0x1e, 0xda, 0xd0, 0x49, 0xa6, 0x64, 0x9e, 0xea, 0xf6, 0xe4, 0xc9, 0xbb, 0x1f, 0x6b, 0xf1, 0xb4,
	0x47, 0xf1, 0x50, 0xa4, 0xfd, 0x06, 0x7e, 0x82, 0x25, 0x93, 0x97, 0xdd, 0x42, 0x2f, 0xcd, 0x29,
	0x33, 0xff, 0x24, 0xbf, 0xfc, 0x98, 0xe7, 0xc9, 0x43, 0xba, 0x20, 0x96, 0xca, 0x8e, 0x78, 0x32,
	0x09, 0x63, 0x7b, 0x2e, 0xa5, 0xb0, 0xe6, 0x89, 0x44, 0x49, 0x0f, 0xd3, 0xdc, 0xca, 0xf2, 0x87,
	0x9d, 0x89, 0x9c, 0x48, 0x9d, 0xdb, 0xe9, 0x2a, 0x7b, 0x64, 0xf0, 0xbb, 0x49, 0xea, 0x1f, 0xa4,
	0x14, 0xf4, 0x98, 0xb4, 0x78, 0x14, 0xa5, 0xcb, 0x71, 0xc0, 0x8c, 0xbe, 0x31, 0xac, 0x3b, 0x57,
	0x01, 0x7d, 0x4d, 0x1a, 0x53, 0xe0, 0x02, 0xa7, 0xec, 0x46, 0xdf, 0x18, 0xb6, 0x46, 0xd6, 0xf9,
	0xaa, 0x57, 0xfb, 0xbb, 0xea, 0x3d, 0x99, 0x84, 0x38, 0x5d, 0x78, 0x96, 0x2f, 0x23, 0xdb, 0x97,
	0x2a, 0x92, 0x2a, 0xbf, 0x9c, 0xa8, 0x60, 0x66, 0xe3, 0x72, 0x0e, 0xca, 0x7a, 0x05, 0xbe, 0x93,
	0xbf, 0x4d, 0x19, 0x39, 0x80, 0x98, 0x7b, 0x02, 0x02, 0x76, 0xb3, 0x6f, 0x0c, 0x9b, 0x4e, 0xb1,
	0xa5, 0x5d, 0xd2, 0xf0, 0x85, 0x54, 0x10, 0xb0, 0xba, 0xbe, 0x91, 0xef, 0x28, 0x27, 0x1d, 0x38,
	0x43, 0x48, 0x62, 0x2e, 0x5c, 0x11, 0x72, 0x2f, 0x14, 0x21, 0x86, 0xa0, 0xd8, 0xad, 0xbd, 0x3d,
	0xc6, 0x31, 0x3a, 0xf7, 0x0b, 0xd6, 0x9b, 0x2b, 0x14, 0xfd, 0x4c, 0xee, 0x96, 0x9f, 0xf0, 0x17,
	0x0a, 0x65, 0xb0, 0x64, 0x8d, 0x4a, 0xf8, 0x3b, 0x05, 0xe7, 0x65, 0x86, 0xa1, 0x5f, 0x09, 0x8d,
	0x39, 0x86, 0xdf, 0x60, 0xcb, 0xfd, 0xa0, 0x12, 0xfc, 0x5e, 0x46, 0xba, 0x6e, 0xfe, 0x89, 0xdc,
	0xce, 0xf1, 0x85, 0x77, 0xb3, 0x12, 0xba, 0x9d, 0x51, 0x0a, 0xeb, 0x8f, 0xa4, 0x1d, 0xc6, 0x08,
	0x09, 0x28, 0x74, 0x13, 0x8e, 0xc0, 0x5a, 0x95, 0x8a, 0x7e, 0x54, 0x40, 0x1c, 0x8e, 0x40, 0x7f,
	0x90, 0x4e, 0xee, 0xca, 0x95, 0x02, 0x74, 0x3d, 0x2e, 0x78, 0xec, 0x03, 0x23, 0x9a, 0xfd, 0x76,
	0x3f, 0xe3, 0xff, 0xab, 0xde, 0xa3, 0x25, 0x8f, 0xc4, 0xf3, 0xc1, 0x2e, 0xe6, 0xc0, 0xc9, 0x4f,
	0xfd, 0x45, 0x9a, 0x8e, 0xb2, 0x90, 0xfe, 0x34, 0x48, 0xb7, 0xac, 0xf3, 0xb6, 0xc3, 0xa1, 0x76,
	0x78, 0xbf, 0xb7, 0xc3, 0xe3, 0xcc, 0x61, 0x37, 0x75, 0xe0, 0x94, 0x9d, 0xbb, 0xe5, 0x81, 0xc4,
	0x5c, 0xc4, 0x0a, 0x10, 0x05, 0x04, 0xee, 0xce, 0xde, 0x3e, 0xaa, 0x54, 0xc4, 0xe3, 0x92, 0x7a,
	0xba, 0xa3, 0xc9, 0x3d, 0xf2, 0xc0, 0x13, 0xd2, 0x9f, 0xb9, 0x65, 0x65, 0xb3, 0x23, 0x62, 0xed,
	0x6a, 0x3f, 0x92, 0x86, 0x8d, 0x73, 0xd6, 0x3b, 0x8d, 0x1a, 0x9d, 0x9e, 0xaf, 0x4d, 0xe3, 0x62,
	0x6d, 0x1a, 0xff, 0xd6, 0xa6, 0xf1, 0x6b, 0x63, 0xd6, 0x2e, 0x36, 0x66, 0xed, 0xcf, 0xc6, 0xac,
	0x7d, 0x79, 0x7a, 0x0d, 0x9b, 0x0e, 0xa5, 0x93, 0x18, 0xf0, 0xbb, 0x4c, 0x66, 0x7a, 0x63, 0x9f,
	0x15, 0xb3, 0x4b, 0xf3, 0xbd, 0x86, 0x1e, 0x4d, 0xcf, 0x2e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaf,
	0xcc, 0xe9, 0x32, 0xd7, 0x04, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BlockInterestNative.Size()
		i -= size
		if _, err := m.BlockInterestNative.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.UnsettledExternalLiabilities.Size()
		i -= size
		if _, err := m.UnsettledExternalLiabilities.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.ExternalAssetBalance.Size()
		i -= size
		if _, err := m.ExternalAssetBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.NativeAssetBalance.Size()
		i -= size
		if _, err := m.NativeAssetBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.InterestRate.Size()
		i -= size
		if _, err := m.InterestRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.NativeCustody.Size()
		i -= size
		if _, err := m.NativeCustody.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.NativeLiabilities.Size()
		i -= size
		if _, err := m.NativeLiabilities.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.ExternalCustody.Size()
		i -= size
		if _, err := m.ExternalCustody.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ExternalLiabilities.Size()
		i -= size
		if _, err := m.ExternalLiabilities.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Closed {
		i--
		if m.Closed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Health.Size()
		i -= size
		if _, err := m.Health.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AmmPoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.AmmPoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AmmPoolId != 0 {
		n += 1 + sovPool(uint64(m.AmmPoolId))
	}
	l = m.Health.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.Enabled {
		n += 2
	}
	if m.Closed {
		n += 2
	}
	l = m.ExternalLiabilities.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.ExternalCustody.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.NativeLiabilities.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.NativeCustody.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.InterestRate.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.NativeAssetBalance.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.ExternalAssetBalance.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.UnsettledExternalLiabilities.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.BlockInterestNative.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmmPoolId", wireType)
			}
			m.AmmPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AmmPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Health", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Health.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Closed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
			m.Closed = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalLiabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExternalLiabilities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalCustody", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExternalCustody.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeLiabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NativeLiabilities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeCustody", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NativeCustody.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeAssetBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NativeAssetBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAssetBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExternalAssetBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnsettledExternalLiabilities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnsettledExternalLiabilities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInterestNative", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockInterestNative.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)