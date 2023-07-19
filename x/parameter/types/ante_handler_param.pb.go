// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/parameter/ante_handler_param.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type AnteHandlerParam struct {
	MinCommissionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=minCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minCommissionRate"`
	MaxVotingPower    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=maxVotingPower,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maxVotingPower"`
	MinSelfDelegation github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=minSelfDelegation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"minSelfDelegation"`
}

func (m *AnteHandlerParam) Reset()         { *m = AnteHandlerParam{} }
func (m *AnteHandlerParam) String() string { return proto.CompactTextString(m) }
func (*AnteHandlerParam) ProtoMessage()    {}
func (*AnteHandlerParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ae17167f40c52b, []int{0}
}
func (m *AnteHandlerParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnteHandlerParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnteHandlerParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnteHandlerParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnteHandlerParam.Merge(m, src)
}
func (m *AnteHandlerParam) XXX_Size() int {
	return m.Size()
}
func (m *AnteHandlerParam) XXX_DiscardUnknown() {
	xxx_messageInfo_AnteHandlerParam.DiscardUnknown(m)
}

var xxx_messageInfo_AnteHandlerParam proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AnteHandlerParam)(nil), "elys.parameter.AnteHandlerParam")
}

func init() {
	proto.RegisterFile("elys/parameter/ante_handler_param.proto", fileDescriptor_61ae17167f40c52b)
}

var fileDescriptor_61ae17167f40c52b = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0x77, 0xf3, 0xc1, 0x07, 0x6e, 0x11, 0x34, 0x58, 0x84, 0x14, 0x13, 0xb1, 0x50, 0x9b,
	0xcc, 0x10, 0x7c, 0x02, 0x63, 0x8a, 0xd8, 0x85, 0x08, 0x29, 0x44, 0x08, 0xb3, 0x9b, 0xeb, 0x66,
	0xc8, 0xce, 0xdc, 0x65, 0xe6, 0x6a, 0x92, 0xb7, 0xf0, 0x29, 0x7c, 0x96, 0x94, 0x29, 0xc5, 0x22,
	0xc8, 0xee, 0x8b, 0xc8, 0xee, 0x06, 0x09, 0xa6, 0xd2, 0x6a, 0xfe, 0x9c, 0xc3, 0xef, 0x5c, 0xee,
	0x09, 0x2e, 0x21, 0x59, 0x39, 0x91, 0x4a, 0x2b, 0x35, 0x10, 0x58, 0x21, 0x0d, 0xc1, 0x64, 0x26,
	0xcd, 0x34, 0x01, 0x3b, 0x29, 0xbf, 0x79, 0x6a, 0x91, 0xb0, 0x51, 0x2f, 0x8c, 0xfc, 0xdb, 0xd8,
	0x3a, 0x8d, 0x31, 0xc6, 0x52, 0x12, 0xc5, 0xad, 0x72, 0xb5, 0x58, 0x84, 0x4e, 0xa3, 0x13, 0xa1,
	0x74, 0x20, 0x5e, 0xba, 0x21, 0x90, 0xec, 0x8a, 0x08, 0x95, 0xa9, 0xf4, 0xf3, 0xb7, 0x5a, 0x70,
	0x7c, 0x63, 0x08, 0x06, 0x55, 0xc2, 0xb0, 0xc0, 0x35, 0x1e, 0x83, 0x13, 0xad, 0xcc, 0x2d, 0x6a,
	0xad, 0x9c, 0x53, 0x68, 0x46, 0x92, 0xa0, 0xe9, 0x9f, 0xf9, 0x57, 0x47, 0x3d, 0xbe, 0xde, 0xb6,
	0xbd, 0x8f, 0x6d, 0xfb, 0x22, 0x56, 0x34, 0x7b, 0x0e, 0x79, 0x84, 0x5a, 0xec, 0x22, 0xaa, 0xa3,
	0xe3, 0xa6, 0x73, 0x41, 0xab, 0x14, 0x1c, 0xef, 0x43, 0x34, 0x3a, 0x04, 0x35, 0xc6, 0x41, 0x5d,
	0xcb, 0xe5, 0x18, 0x49, 0x99, 0x78, 0x88, 0x0b, 0xb0, 0xcd, 0xda, 0x9f, 0xd0, 0x3f, 0x28, 0xbb,
	0xa9, 0xef, 0x21, 0x79, 0xea, 0x43, 0x02, 0xb1, 0x24, 0x85, 0xa6, 0xf9, 0xef, 0xd7, 0xe8, 0x3b,
	0x43, 0xa3, 0x43, 0x50, 0x6f, 0xb0, 0xce, 0x98, 0xbf, 0xc9, 0x98, 0xff, 0x99, 0x31, 0xff, 0x35,
	0x67, 0xde, 0x26, 0x67, 0xde, 0x7b, 0xce, 0xbc, 0x07, 0xbe, 0x07, 0x2d, 0x3a, 0xe9, 0x18, 0xa0,
	0x05, 0xda, 0x79, 0xf9, 0x10, 0xcb, 0xbd, 0x2e, 0xcb, 0x80, 0xf0, 0x7f, 0xb9, 0xf9, 0xeb, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xc5, 0xbc, 0x00, 0xea, 0x01, 0x00, 0x00,
}

func (m *AnteHandlerParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnteHandlerParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnteHandlerParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinSelfDelegation.Size()
		i -= size
		if _, err := m.MinSelfDelegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MaxVotingPower.Size()
		i -= size
		if _, err := m.MaxVotingPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinCommissionRate.Size()
		i -= size
		if _, err := m.MinCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAnteHandlerParam(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnteHandlerParam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnteHandlerParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinCommissionRate.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	l = m.MaxVotingPower.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	l = m.MinSelfDelegation.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	return n
}

func sovAnteHandlerParam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnteHandlerParam(x uint64) (n int) {
	return sovAnteHandlerParam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnteHandlerParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnteHandlerParam
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
			return fmt.Errorf("proto: AnteHandlerParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnteHandlerParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVotingPower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxVotingPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSelfDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSelfDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnteHandlerParam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnteHandlerParam
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
func skipAnteHandlerParam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnteHandlerParam
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
					return 0, ErrIntOverflowAnteHandlerParam
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
					return 0, ErrIntOverflowAnteHandlerParam
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
				return 0, ErrInvalidLengthAnteHandlerParam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnteHandlerParam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnteHandlerParam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnteHandlerParam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnteHandlerParam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnteHandlerParam = fmt.Errorf("proto: unexpected end of group")
)
