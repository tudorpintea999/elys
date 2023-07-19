// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/incentive/params.proto

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

// Params defines the parameters for the module.
type Params struct {
	LpIncentives        []IncentiveInfo                        `protobuf:"bytes,1,rep,name=lp_incentives,json=lpIncentives,proto3" json:"lp_incentives"`
	StakeIncentives     []IncentiveInfo                        `protobuf:"bytes,2,rep,name=stake_incentives,json=stakeIncentives,proto3" json:"stake_incentives"`
	CommunityTax        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=community_tax,json=communityTax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_tax"`
	WithdrawAddrEnabled bool                                   `protobuf:"varint,4,opt,name=withdraw_addr_enabled,json=withdrawAddrEnabled,proto3" json:"withdraw_addr_enabled,omitempty"`
	// Dex revenue percent for lps, 100 - reward_portion_for_lps = revenue percent for stakers.
	RewardPortionForLps github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reward_portion_for_lps,json=rewardPortionForLps,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_portion_for_lps"`
	// Pool information
	// poolId, reward wallet, mulitplier
	PoolInfos []PoolInfo `protobuf:"bytes,6,rep,name=pool_infos,json=poolInfos,proto3" json:"pool_infos"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca0267cb466fec, []int{0}
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

func (m *Params) GetLpIncentives() []IncentiveInfo {
	if m != nil {
		return m.LpIncentives
	}
	return nil
}

func (m *Params) GetStakeIncentives() []IncentiveInfo {
	if m != nil {
		return m.StakeIncentives
	}
	return nil
}

func (m *Params) GetWithdrawAddrEnabled() bool {
	if m != nil {
		return m.WithdrawAddrEnabled
	}
	return false
}

func (m *Params) GetPoolInfos() []PoolInfo {
	if m != nil {
		return m.PoolInfos
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "elys.incentive.Params")
}

func init() { proto.RegisterFile("elys/incentive/params.proto", fileDescriptor_3bca0267cb466fec) }

var fileDescriptor_3bca0267cb466fec = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0xef, 0xd2, 0x30,
	0x14, 0xc7, 0x37, 0x41, 0x22, 0x15, 0xd4, 0x0c, 0x35, 0x13, 0xe3, 0x20, 0x1e, 0x0c, 0x17, 0xba,
	0x04, 0x6f, 0x26, 0x1e, 0x24, 0x6a, 0x20, 0x31, 0x86, 0xa0, 0x27, 0x2f, 0x4d, 0x59, 0x0b, 0x2c,
	0x6c, 0x7d, 0x4d, 0x5b, 0x1c, 0xfc, 0x17, 0x1e, 0x3d, 0xfa, 0xe7, 0x70, 0xe4, 0x62, 0x62, 0x3c,
	0x10, 0x03, 0xff, 0x88, 0xd9, 0x06, 0xcb, 0xe4, 0x64, 0x7e, 0xa7, 0xbe, 0xd7, 0xef, 0xeb, 0xe7,
	0x7d, 0xf3, 0xfa, 0xd0, 0x53, 0x1e, 0x6d, 0xb5, 0x1f, 0x8a, 0x80, 0x0b, 0x13, 0x7e, 0xe5, 0xbe,
	0xa4, 0x8a, 0xc6, 0x1a, 0x4b, 0x05, 0x06, 0x9c, 0x7b, 0xa9, 0x88, 0x0b, 0xb1, 0xfd, 0x70, 0x01,
	0x0b, 0xc8, 0x24, 0x3f, 0x8d, 0xf2, 0xaa, 0xb6, 0x77, 0x85, 0x28, 0xa2, 0xb3, 0xfe, 0xe4, 0xba,
	0x05, 0x40, 0x94, 0x4b, 0xcf, 0x7f, 0x56, 0x50, 0x6d, 0x92, 0x75, 0x74, 0x46, 0xa8, 0x19, 0x49,
	0x52, 0x54, 0x69, 0xd7, 0xee, 0x56, 0x7a, 0x77, 0x07, 0xcf, 0xf0, 0xbf, 0x1e, 0xf0, 0xf8, 0x12,
	0x8d, 0xc5, 0x1c, 0x86, 0xd5, 0xdd, 0xa1, 0x63, 0x4d, 0x1b, 0x91, 0x2c, 0xae, 0xb5, 0xf3, 0x11,
	0x3d, 0xd0, 0x86, 0xae, 0x78, 0x19, 0x76, 0xeb, 0xff, 0x61, 0xf7, 0xb3, 0xc7, 0x25, 0xde, 0x27,
	0xd4, 0x0c, 0x20, 0x8e, 0xd7, 0x22, 0x34, 0x5b, 0x62, 0xe8, 0xc6, 0xad, 0x74, 0xed, 0x5e, 0x7d,
	0x88, 0xd3, 0xea, 0xdf, 0x87, 0xce, 0x8b, 0x45, 0x68, 0x96, 0xeb, 0x19, 0x0e, 0x20, 0xf6, 0x03,
	0xd0, 0x31, 0xe8, 0xf3, 0xd1, 0xd7, 0x6c, 0xe5, 0x9b, 0xad, 0xe4, 0x1a, 0xbf, 0xe5, 0xc1, 0xb4,
	0x51, 0x40, 0x3e, 0xd3, 0x8d, 0x33, 0x40, 0x8f, 0x92, 0xd0, 0x2c, 0x99, 0xa2, 0x09, 0xa1, 0x8c,
	0x29, 0xc2, 0x05, 0x9d, 0x45, 0x9c, 0xb9, 0xd5, 0xae, 0xdd, 0xbb, 0x33, 0x6d, 0x5d, 0xc4, 0x37,
	0x8c, 0xa9, 0x77, 0xb9, 0xe4, 0x04, 0xe8, 0xb1, 0xe2, 0x09, 0x55, 0x8c, 0x48, 0x50, 0x26, 0x04,
	0x41, 0xe6, 0xa0, 0x48, 0x24, 0xb5, 0x7b, 0xfb, 0x46, 0x8e, 0x5a, 0x39, 0x6d, 0x92, 0xc3, 0xde,
	0x83, 0xfa, 0x20, 0xb5, 0xf3, 0x1a, 0xa1, 0xf4, 0x83, 0x48, 0x28, 0xe6, 0xa0, 0xdd, 0x5a, 0x36,
	0x37, 0xf7, 0x7a, 0x6e, 0x13, 0x80, 0xa8, 0x34, 0xb2, 0xba, 0x3c, 0xe7, 0xfa, 0x55, 0xf5, 0xfb,
	0x8f, 0x8e, 0x35, 0x1c, 0xed, 0x8e, 0x9e, 0xbd, 0x3f, 0x7a, 0xf6, 0x9f, 0xa3, 0x67, 0x7f, 0x3b,
	0x79, 0xd6, 0xfe, 0xe4, 0x59, 0xbf, 0x4e, 0x9e, 0xf5, 0x05, 0x97, 0xbc, 0xa5, 0xd0, 0xbe, 0xe0,
	0x26, 0x01, 0xb5, 0xca, 0x12, 0x7f, 0x53, 0x5a, 0x93, 0xcc, 0xe7, 0xac, 0x96, 0x2d, 0xca, 0xcb,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x38, 0x21, 0x0a, 0xa8, 0x02, 0x00, 0x00,
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
	if len(m.PoolInfos) > 0 {
		for iNdEx := len(m.PoolInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.RewardPortionForLps.Size()
		i -= size
		if _, err := m.RewardPortionForLps.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.WithdrawAddrEnabled {
		i--
		if m.WithdrawAddrEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.CommunityTax.Size()
		i -= size
		if _, err := m.CommunityTax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.StakeIncentives) > 0 {
		for iNdEx := len(m.StakeIncentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakeIncentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.LpIncentives) > 0 {
		for iNdEx := len(m.LpIncentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LpIncentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
	if len(m.LpIncentives) > 0 {
		for _, e := range m.LpIncentives {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.StakeIncentives) > 0 {
		for _, e := range m.StakeIncentives {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.CommunityTax.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.WithdrawAddrEnabled {
		n += 2
	}
	l = m.RewardPortionForLps.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.PoolInfos) > 0 {
		for _, e := range m.PoolInfos {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field LpIncentives", wireType)
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
			m.LpIncentives = append(m.LpIncentives, IncentiveInfo{})
			if err := m.LpIncentives[len(m.LpIncentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeIncentives", wireType)
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
			m.StakeIncentives = append(m.StakeIncentives, IncentiveInfo{})
			if err := m.StakeIncentives[len(m.StakeIncentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityTax", wireType)
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
			if err := m.CommunityTax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddrEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.WithdrawAddrEnabled = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPortionForLps", wireType)
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
			if err := m.RewardPortionForLps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolInfos", wireType)
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
			m.PoolInfos = append(m.PoolInfos, PoolInfo{})
			if err := m.PoolInfos[len(m.PoolInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
