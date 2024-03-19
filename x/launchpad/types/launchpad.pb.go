// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/launchpad/launchpad.proto

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

// 0-20% raise 100% bonus
// 20-30% raise 90% bonus
// 30-40% raise bonus 80%
// 40-50% raise bonus 70%
// 50-60% raise bonus 60%
// 60-70% raise bonus 50%
// 70-80% raise bonus 40%
// 80-100% raise bonus 30%
type Bonus struct {
	RaisePercents   []uint64 `protobuf:"varint,1,rep,packed,name=raise_percents,json=raisePercents,proto3" json:"raise_percents,omitempty"`
	BonusPercents   []uint64 `protobuf:"varint,2,rep,packed,name=bonus_percents,json=bonusPercents,proto3" json:"bonus_percents,omitempty"`
	LockDuration    uint64   `protobuf:"varint,3,opt,name=lock_duration,json=lockDuration,proto3" json:"lock_duration,omitempty"`
	VestingDuration uint64   `protobuf:"varint,4,opt,name=vesting_duration,json=vestingDuration,proto3" json:"vesting_duration,omitempty"`
}

func (m *Bonus) Reset()         { *m = Bonus{} }
func (m *Bonus) String() string { return proto.CompactTextString(m) }
func (*Bonus) ProtoMessage()    {}
func (*Bonus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffa86e749b31956, []int{0}
}
func (m *Bonus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bonus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bonus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bonus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bonus.Merge(m, src)
}
func (m *Bonus) XXX_Size() int {
	return m.Size()
}
func (m *Bonus) XXX_DiscardUnknown() {
	xxx_messageInfo_Bonus.DiscardUnknown(m)
}

var xxx_messageInfo_Bonus proto.InternalMessageInfo

func (m *Bonus) GetRaisePercents() []uint64 {
	if m != nil {
		return m.RaisePercents
	}
	return nil
}

func (m *Bonus) GetBonusPercents() []uint64 {
	if m != nil {
		return m.BonusPercents
	}
	return nil
}

func (m *Bonus) GetLockDuration() uint64 {
	if m != nil {
		return m.LockDuration
	}
	return 0
}

func (m *Bonus) GetVestingDuration() uint64 {
	if m != nil {
		return m.VestingDuration
	}
	return 0
}

type Purchase struct {
	OrderId            uint64                                 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderMaker         string                                 `protobuf:"bytes,2,opt,name=order_maker,json=orderMaker,proto3" json:"order_maker,omitempty"`
	SpendingToken      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=spending_token,json=spendingToken,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"spending_token"`
	TokenAmount        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=token_amount,json=tokenAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_amount"`
	ElysAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=elys_amount,json=elysAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"elys_amount"`
	ReturnedElysAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=returned_elys_amount,json=returnedElysAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"returned_elys_amount"`
	BonusAmount        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=bonus_amount,json=bonusAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bonus_amount"`
}

func (m *Purchase) Reset()         { *m = Purchase{} }
func (m *Purchase) String() string { return proto.CompactTextString(m) }
func (*Purchase) ProtoMessage()    {}
func (*Purchase) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffa86e749b31956, []int{1}
}
func (m *Purchase) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Purchase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Purchase.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Purchase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Purchase.Merge(m, src)
}
func (m *Purchase) XXX_Size() int {
	return m.Size()
}
func (m *Purchase) XXX_DiscardUnknown() {
	xxx_messageInfo_Purchase.DiscardUnknown(m)
}

var xxx_messageInfo_Purchase proto.InternalMessageInfo

func (m *Purchase) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Purchase) GetOrderMaker() string {
	if m != nil {
		return m.OrderMaker
	}
	return ""
}

func init() {
	proto.RegisterType((*Bonus)(nil), "elys.launchpad.Bonus")
	proto.RegisterType((*Purchase)(nil), "elys.launchpad.Purchase")
}

func init() { proto.RegisterFile("elys/launchpad/launchpad.proto", fileDescriptor_5ffa86e749b31956) }

var fileDescriptor_5ffa86e749b31956 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbd, 0xce, 0xd3, 0x30,
	0x14, 0x86, 0xe3, 0xfe, 0xd7, 0xfd, 0x01, 0x59, 0x1d, 0x02, 0x43, 0x5a, 0x15, 0x81, 0xca, 0xd0,
	0x64, 0xe0, 0x0a, 0xa8, 0x40, 0xa2, 0x03, 0xa2, 0x8d, 0x60, 0x61, 0x09, 0x69, 0x6c, 0xa5, 0x51,
	0x1a, 0x3b, 0xb2, 0x1d, 0xa0, 0x77, 0xc1, 0x4d, 0x70, 0x2f, 0x1d, 0x3b, 0x22, 0x86, 0x0a, 0xb5,
	0xf7, 0xc0, 0xfc, 0xc9, 0x4e, 0xd2, 0x74, 0xfd, 0x3a, 0xc5, 0x7e, 0xcf, 0xa3, 0xc7, 0x39, 0xc9,
	0x31, 0xb4, 0xc8, 0x6e, 0x2f, 0x9c, 0x9d, 0x9f, 0xd1, 0x60, 0x9b, 0xfa, 0xb8, 0x5a, 0xd9, 0x29,
	0x67, 0x92, 0xa1, 0xa1, 0xaa, 0xdb, 0xd7, 0xf4, 0xf9, 0x28, 0x64, 0x21, 0xd3, 0x25, 0x47, 0xad,
	0x72, 0x6a, 0xfa, 0x1b, 0xc0, 0xe6, 0x82, 0xd1, 0x4c, 0xa0, 0x97, 0x70, 0xc8, 0xfd, 0x48, 0x10,
	0x2f, 0x25, 0x3c, 0x20, 0x54, 0x0a, 0x13, 0x4c, 0xea, 0xb3, 0x86, 0x3b, 0xd0, 0xe9, 0xaa, 0x08,
	0x15, 0xb6, 0x51, 0x7c, 0x85, 0xd5, 0x72, 0x4c, 0xa7, 0x57, 0xec, 0x05, 0x1c, 0xec, 0x58, 0x10,
	0x7b, 0x38, 0xe3, 0xbe, 0x8c, 0x18, 0x35, 0xeb, 0x13, 0x30, 0x6b, 0xb8, 0x7d, 0x15, 0xbe, 0x2b,
	0x32, 0xf4, 0x1a, 0x3e, 0xfd, 0x4e, 0x84, 0x8c, 0x68, 0x58, 0x71, 0x0d, 0xcd, 0x3d, 0x29, 0xf2,
	0x12, 0x9d, 0xfe, 0xaf, 0xc3, 0xce, 0x2a, 0xe3, 0xc1, 0xd6, 0x17, 0x04, 0x3d, 0x83, 0x1d, 0xc6,
	0x31, 0xe1, 0x5e, 0x84, 0x4d, 0xa0, 0xf9, 0xb6, 0xde, 0x2f, 0x31, 0x1a, 0xc3, 0x5e, 0x5e, 0x4a,
	0xfc, 0x98, 0x70, 0xb3, 0x36, 0x01, 0xb3, 0xae, 0x0b, 0x75, 0xf4, 0x51, 0x25, 0xe8, 0x0b, 0x1c,
	0x8a, 0x94, 0x50, 0xac, 0x0e, 0x95, 0x2c, 0x26, 0xf9, 0x9b, 0x75, 0x17, 0xf6, 0xe1, 0x34, 0x36,
	0xfe, 0x9e, 0xc6, 0xaf, 0xc2, 0x48, 0x6e, 0xb3, 0x8d, 0x1d, 0xb0, 0xc4, 0x09, 0x98, 0x48, 0x98,
	0x28, 0x1e, 0x73, 0x81, 0x63, 0x47, 0xee, 0x53, 0x22, 0xec, 0x25, 0x95, 0xee, 0xa0, 0xb4, 0x7c,
	0x56, 0x12, 0xb4, 0x86, 0x7d, 0x6d, 0xf3, 0xfc, 0x84, 0x65, 0x54, 0xea, 0x36, 0x1e, 0x2f, 0xed,
	0x69, 0xc7, 0x5b, 0xad, 0x40, 0x9f, 0x60, 0x4f, 0xfd, 0xc2, 0xd2, 0xd8, 0xbc, 0xcb, 0x08, 0x95,
	0xa2, 0x10, 0x7e, 0x83, 0x23, 0x4e, 0x64, 0xc6, 0x29, 0xc1, 0xde, 0xad, 0xb9, 0x75, 0x97, 0x19,
	0x95, 0xae, 0xf7, 0xd5, 0x09, 0x6b, 0xd8, 0xcf, 0x87, 0xa3, 0x30, 0xb7, 0xef, 0xfb, 0x0a, 0xda,
	0x91, 0x2b, 0x17, 0x1f, 0x0e, 0x67, 0x0b, 0x1c, 0xcf, 0x16, 0xf8, 0x77, 0xb6, 0xc0, 0xaf, 0x8b,
	0x65, 0x1c, 0x2f, 0x96, 0xf1, 0xe7, 0x62, 0x19, 0x5f, 0xed, 0x1b, 0x9d, 0x6a, 0x67, 0x4e, 0x89,
	0xfc, 0xc1, 0x78, 0xac, 0x37, 0xce, 0xcf, 0x9b, 0xab, 0xa1, 0xd5, 0x9b, 0x96, 0x9e, 0xf8, 0x37,
	0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x29, 0x1b, 0x50, 0x39, 0x03, 0x00, 0x00,
}

func (m *Bonus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bonus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bonus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VestingDuration != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.VestingDuration))
		i--
		dAtA[i] = 0x20
	}
	if m.LockDuration != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.LockDuration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BonusPercents) > 0 {
		dAtA2 := make([]byte, len(m.BonusPercents)*10)
		var j1 int
		for _, num := range m.BonusPercents {
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
		i = encodeVarintLaunchpad(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RaisePercents) > 0 {
		dAtA4 := make([]byte, len(m.RaisePercents)*10)
		var j3 int
		for _, num := range m.RaisePercents {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintLaunchpad(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Purchase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Purchase) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Purchase) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BonusAmount.Size()
		i -= size
		if _, err := m.BonusAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLaunchpad(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.ReturnedElysAmount.Size()
		i -= size
		if _, err := m.ReturnedElysAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLaunchpad(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ElysAmount.Size()
		i -= size
		if _, err := m.ElysAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLaunchpad(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TokenAmount.Size()
		i -= size
		if _, err := m.TokenAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLaunchpad(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SpendingToken.Size()
		i -= size
		if _, err := m.SpendingToken.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLaunchpad(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.OrderMaker) > 0 {
		i -= len(m.OrderMaker)
		copy(dAtA[i:], m.OrderMaker)
		i = encodeVarintLaunchpad(dAtA, i, uint64(len(m.OrderMaker)))
		i--
		dAtA[i] = 0x12
	}
	if m.OrderId != 0 {
		i = encodeVarintLaunchpad(dAtA, i, uint64(m.OrderId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLaunchpad(dAtA []byte, offset int, v uint64) int {
	offset -= sovLaunchpad(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bonus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RaisePercents) > 0 {
		l = 0
		for _, e := range m.RaisePercents {
			l += sovLaunchpad(uint64(e))
		}
		n += 1 + sovLaunchpad(uint64(l)) + l
	}
	if len(m.BonusPercents) > 0 {
		l = 0
		for _, e := range m.BonusPercents {
			l += sovLaunchpad(uint64(e))
		}
		n += 1 + sovLaunchpad(uint64(l)) + l
	}
	if m.LockDuration != 0 {
		n += 1 + sovLaunchpad(uint64(m.LockDuration))
	}
	if m.VestingDuration != 0 {
		n += 1 + sovLaunchpad(uint64(m.VestingDuration))
	}
	return n
}

func (m *Purchase) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderId != 0 {
		n += 1 + sovLaunchpad(uint64(m.OrderId))
	}
	l = len(m.OrderMaker)
	if l > 0 {
		n += 1 + l + sovLaunchpad(uint64(l))
	}
	l = m.SpendingToken.Size()
	n += 1 + l + sovLaunchpad(uint64(l))
	l = m.TokenAmount.Size()
	n += 1 + l + sovLaunchpad(uint64(l))
	l = m.ElysAmount.Size()
	n += 1 + l + sovLaunchpad(uint64(l))
	l = m.ReturnedElysAmount.Size()
	n += 1 + l + sovLaunchpad(uint64(l))
	l = m.BonusAmount.Size()
	n += 1 + l + sovLaunchpad(uint64(l))
	return n
}

func sovLaunchpad(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLaunchpad(x uint64) (n int) {
	return sovLaunchpad(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bonus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLaunchpad
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
			return fmt.Errorf("proto: Bonus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bonus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLaunchpad
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
				m.RaisePercents = append(m.RaisePercents, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLaunchpad
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
					return ErrInvalidLengthLaunchpad
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLaunchpad
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
				if elementCount != 0 && len(m.RaisePercents) == 0 {
					m.RaisePercents = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLaunchpad
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
					m.RaisePercents = append(m.RaisePercents, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RaisePercents", wireType)
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLaunchpad
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
				m.BonusPercents = append(m.BonusPercents, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLaunchpad
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
					return ErrInvalidLengthLaunchpad
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthLaunchpad
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
				if elementCount != 0 && len(m.BonusPercents) == 0 {
					m.BonusPercents = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLaunchpad
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
					m.BonusPercents = append(m.BonusPercents, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusPercents", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockDuration", wireType)
			}
			m.LockDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingDuration", wireType)
			}
			m.VestingDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLaunchpad(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLaunchpad
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
func (m *Purchase) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLaunchpad
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
			return fmt.Errorf("proto: Purchase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Purchase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			m.OrderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderMaker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderMaker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendingToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SpendingToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElysAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ElysAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnedElysAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReturnedElysAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLaunchpad
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
				return ErrInvalidLengthLaunchpad
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLaunchpad
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BonusAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLaunchpad(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLaunchpad
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
func skipLaunchpad(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLaunchpad
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
					return 0, ErrIntOverflowLaunchpad
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
					return 0, ErrIntOverflowLaunchpad
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
				return 0, ErrInvalidLengthLaunchpad
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLaunchpad
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLaunchpad
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLaunchpad        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLaunchpad          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLaunchpad = fmt.Errorf("proto: unexpected end of group")
)
