// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/launchpad/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	InitialPrice       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=initial_price,json=initialPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"initial_price"`
	TotalReserve       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=total_reserve,json=totalReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_reserve"`
	SoldAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=sold_amount,json=soldAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sold_amount"`
	WithdrawAddress    string                                 `protobuf:"bytes,4,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
	WithdrawnAmount    []types.Coin                           `protobuf:"bytes,5,rep,name=withdrawn_amount,json=withdrawnAmount,proto3" json:"withdrawn_amount"`
	LaunchpadStarttime uint64                                 `protobuf:"varint,6,opt,name=launchpad_starttime,json=launchpadStarttime,proto3" json:"launchpad_starttime,omitempty"`
	LaunchpadDuration  uint64                                 `protobuf:"varint,7,opt,name=launchpad_duration,json=launchpadDuration,proto3" json:"launchpad_duration,omitempty"`
	ReturnDuration     uint64                                 `protobuf:"varint,8,opt,name=return_duration,json=returnDuration,proto3" json:"return_duration,omitempty"`
	MaxReturnPercent   uint64                                 `protobuf:"varint,9,opt,name=max_return_percent,json=maxReturnPercent,proto3" json:"max_return_percent,omitempty"`
	SpendingTokens     []string                               `protobuf:"bytes,10,rep,name=spending_tokens,json=spendingTokens,proto3" json:"spending_tokens,omitempty"`
	BonusInfo          Bonus                                  `protobuf:"bytes,11,opt,name=bonus_info,json=bonusInfo,proto3" json:"bonus_info"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_42cec57a253264db, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetWithdrawAddress() string {
	if m != nil {
		return m.WithdrawAddress
	}
	return ""
}

func (m *Params) GetWithdrawnAmount() []types.Coin {
	if m != nil {
		return m.WithdrawnAmount
	}
	return nil
}

func (m *Params) GetLaunchpadStarttime() uint64 {
	if m != nil {
		return m.LaunchpadStarttime
	}
	return 0
}

func (m *Params) GetLaunchpadDuration() uint64 {
	if m != nil {
		return m.LaunchpadDuration
	}
	return 0
}

func (m *Params) GetReturnDuration() uint64 {
	if m != nil {
		return m.ReturnDuration
	}
	return 0
}

func (m *Params) GetMaxReturnPercent() uint64 {
	if m != nil {
		return m.MaxReturnPercent
	}
	return 0
}

func (m *Params) GetSpendingTokens() []string {
	if m != nil {
		return m.SpendingTokens
	}
	return nil
}

func (m *Params) GetBonusInfo() Bonus {
	if m != nil {
		return m.BonusInfo
	}
	return Bonus{}
}

func init() {
	proto.RegisterType((*Params)(nil), "elys.launchpad.Params")
}

func init() { proto.RegisterFile("elys/launchpad/params.proto", fileDescriptor_42cec57a253264db) }

var fileDescriptor_42cec57a253264db = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0xda, 0x15, 0xea, 0x8e, 0x6e, 0x18, 0x90, 0xc2, 0x90, 0xd2, 0x8a, 0x03, 0x04,
	0x89, 0xda, 0xda, 0xb8, 0xed, 0xb6, 0xb2, 0x03, 0xe3, 0x42, 0x95, 0x71, 0xe2, 0x12, 0x39, 0x89,
	0xd7, 0x5a, 0x4d, 0xec, 0xc8, 0x76, 0xd6, 0xee, 0x5b, 0x70, 0xe4, 0xc0, 0x81, 0x8f, 0xb3, 0xe3,
	0x8e, 0x88, 0xc3, 0x84, 0xda, 0x2f, 0x82, 0xec, 0xfc, 0xe9, 0xe0, 0x06, 0xa7, 0x38, 0xcf, 0xf3,
	0xe8, 0xf7, 0x5a, 0xef, 0xfb, 0x1a, 0x3c, 0xa7, 0xe9, 0x95, 0xc2, 0x29, 0x29, 0x78, 0x3c, 0xcf,
	0x49, 0x82, 0x73, 0x22, 0x49, 0xa6, 0x50, 0x2e, 0x85, 0x16, 0x70, 0x60, 0x4c, 0xd4, 0x98, 0x07,
	0x4f, 0x66, 0x62, 0x26, 0xac, 0x85, 0xcd, 0xa9, 0x4c, 0x1d, 0x78, 0x7f, 0x21, 0x9a, 0x53, 0xed,
	0xc7, 0x42, 0x65, 0x42, 0xe1, 0x88, 0x28, 0x8a, 0x2f, 0x0f, 0x23, 0xaa, 0xc9, 0x21, 0x8e, 0x05,
	0xe3, 0xa5, 0xff, 0xe2, 0xdb, 0x0e, 0xe8, 0x4e, 0x6d, 0x59, 0x78, 0x0e, 0x1e, 0x32, 0xce, 0x34,
	0x23, 0x69, 0x98, 0x4b, 0x16, 0x53, 0xd7, 0x19, 0x39, 0x7e, 0x6f, 0x82, 0xae, 0x6f, 0x87, 0xad,
	0x9f, 0xb7, 0xc3, 0x97, 0x33, 0xa6, 0xe7, 0x45, 0x84, 0x62, 0x91, 0xe1, 0x0a, 0x5a, 0x7e, 0xc6,
	0x2a, 0x59, 0x60, 0x7d, 0x95, 0x53, 0x85, 0x4e, 0x69, 0x1c, 0xec, 0x56, 0x90, 0xa9, 0x61, 0x18,
	0xa8, 0x16, 0x9a, 0xa4, 0xa1, 0xa4, 0x8a, 0xca, 0x4b, 0xea, 0xde, 0xfb, 0x67, 0xe8, 0x19, 0xd7,
	0xc1, 0xae, 0x85, 0x04, 0x25, 0x03, 0x7e, 0x04, 0x7d, 0x25, 0xd2, 0x24, 0x24, 0x99, 0x28, 0xb8,
	0x76, 0xdb, 0xff, 0x85, 0x04, 0x06, 0x71, 0x62, 0x09, 0xf0, 0x35, 0xd8, 0x5f, 0x32, 0x3d, 0x4f,
	0x24, 0x59, 0x86, 0x24, 0x49, 0x24, 0x55, 0xca, 0xed, 0x18, 0x6a, 0xb0, 0x57, 0xeb, 0x27, 0xa5,
	0x0c, 0x3f, 0x6c, 0xa3, 0xbc, 0xbe, 0xc0, 0xce, 0xa8, 0xed, 0xf7, 0x8f, 0x9e, 0xa1, 0xb2, 0x0e,
	0x32, 0xbd, 0x46, 0x55, 0xaf, 0xd1, 0x3b, 0xc1, 0xf8, 0xa4, 0x63, 0xee, 0xb6, 0x65, 0xf1, 0xaa,
	0x2c, 0x06, 0x8f, 0x9b, 0x79, 0x85, 0x4a, 0x13, 0xa9, 0x35, 0xcb, 0xa8, 0xdb, 0x1d, 0x39, 0x7e,
	0x27, 0x80, 0x8d, 0x75, 0x5e, 0x3b, 0x70, 0x0c, 0xb6, 0x6a, 0x98, 0x14, 0x92, 0x68, 0x26, 0xb8,
	0x7b, 0xdf, 0xe6, 0x1f, 0x35, 0xce, 0x69, 0x65, 0xc0, 0x57, 0x60, 0x4f, 0x52, 0x5d, 0x48, 0xbe,
	0xcd, 0x3e, 0xb0, 0xd9, 0x41, 0x29, 0x37, 0xc1, 0x37, 0x00, 0x66, 0x64, 0x15, 0x56, 0xe1, 0x9c,
	0xca, 0x98, 0x72, 0xed, 0xf6, 0x6c, 0x76, 0x3f, 0x23, 0xab, 0xc0, 0x1a, 0xd3, 0x52, 0x37, 0x58,
	0x95, 0x53, 0x9e, 0x30, 0x3e, 0x0b, 0xb5, 0x58, 0x50, 0xae, 0x5c, 0x30, 0x6a, 0xfb, 0xbd, 0x60,
	0x50, 0xcb, 0x9f, 0xac, 0x0a, 0x8f, 0x01, 0x88, 0x04, 0x2f, 0x54, 0xc8, 0xf8, 0x85, 0x70, 0xfb,
	0x23, 0xc7, 0xef, 0x1f, 0x3d, 0x45, 0x7f, 0xee, 0x35, 0x9a, 0x98, 0x44, 0xd5, 0xa1, 0x9e, 0x8d,
	0x9f, 0xf1, 0x0b, 0x71, 0xdc, 0xf9, 0xfa, 0x7d, 0xd8, 0x9a, 0xbc, 0xbf, 0x5e, 0x7b, 0xce, 0xcd,
	0xda, 0x73, 0x7e, 0xad, 0x3d, 0xe7, 0xcb, 0xc6, 0x6b, 0xdd, 0x6c, 0xbc, 0xd6, 0x8f, 0x8d, 0xd7,
	0xfa, 0x8c, 0xee, 0x8c, 0xd9, 0x10, 0xc7, 0x9c, 0xea, 0xa5, 0x90, 0x0b, 0xfb, 0x83, 0x57, 0x77,
	0x9e, 0x84, 0x1d, 0x79, 0xd4, 0xb5, 0xfb, 0xfe, 0xf6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc,
	0x7d, 0xe4, 0xf0, 0x74, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BonusInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if len(m.SpendingTokens) > 0 {
		for iNdEx := len(m.SpendingTokens) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SpendingTokens[iNdEx])
			copy(dAtA[i:], m.SpendingTokens[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.SpendingTokens[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.MaxReturnPercent != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxReturnPercent))
		i--
		dAtA[i] = 0x48
	}
	if m.ReturnDuration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReturnDuration))
		i--
		dAtA[i] = 0x40
	}
	if m.LaunchpadDuration != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LaunchpadDuration))
		i--
		dAtA[i] = 0x38
	}
	if m.LaunchpadStarttime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LaunchpadStarttime))
		i--
		dAtA[i] = 0x30
	}
	if len(m.WithdrawnAmount) > 0 {
		for iNdEx := len(m.WithdrawnAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.SoldAmount.Size()
		i -= size
		if _, err := m.SoldAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TotalReserve.Size()
		i -= size
		if _, err := m.TotalReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.InitialPrice.Size()
		i -= size
		if _, err := m.InitialPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InitialPrice.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.TotalReserve.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SoldAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.WithdrawnAmount) > 0 {
		for _, e := range m.WithdrawnAmount {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.LaunchpadStarttime != 0 {
		n += 1 + sovParams(uint64(m.LaunchpadStarttime))
	}
	if m.LaunchpadDuration != 0 {
		n += 1 + sovParams(uint64(m.LaunchpadDuration))
	}
	if m.ReturnDuration != 0 {
		n += 1 + sovParams(uint64(m.ReturnDuration))
	}
	if m.MaxReturnPercent != 0 {
		n += 1 + sovParams(uint64(m.MaxReturnPercent))
	}
	if len(m.SpendingTokens) > 0 {
		for _, s := range m.SpendingTokens {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.BonusInfo.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoldAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SoldAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnAmount = append(m.WithdrawnAmount, types.Coin{})
			if err := m.WithdrawnAmount[len(m.WithdrawnAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchpadStarttime", wireType)
			}
			m.LaunchpadStarttime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchpadStarttime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchpadDuration", wireType)
			}
			m.LaunchpadDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchpadDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnDuration", wireType)
			}
			m.ReturnDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReturnDuration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxReturnPercent", wireType)
			}
			m.MaxReturnPercent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxReturnPercent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendingTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendingTokens = append(m.SpendingTokens, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BonusInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BonusInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)