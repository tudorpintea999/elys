// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/amm/pool_params.proto

package types

import (
	fmt "fmt"
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

type PoolParams struct {
	SwapFee                  uint64                    `protobuf:"varint,1,opt,name=swapFee,proto3" json:"swapFee,omitempty"`
	ExitFee                  uint64                    `protobuf:"varint,2,opt,name=exitFee,proto3" json:"exitFee,omitempty"`
	SmoothWeightChangeParams *SmoothWeightChangeParams `protobuf:"bytes,3,opt,name=smoothWeightChangeParams,proto3" json:"smoothWeightChangeParams,omitempty"`
}

func (m *PoolParams) Reset()         { *m = PoolParams{} }
func (m *PoolParams) String() string { return proto.CompactTextString(m) }
func (*PoolParams) ProtoMessage()    {}
func (*PoolParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3500125990074bc9, []int{0}
}
func (m *PoolParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolParams.Merge(m, src)
}
func (m *PoolParams) XXX_Size() int {
	return m.Size()
}
func (m *PoolParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolParams.DiscardUnknown(m)
}

var xxx_messageInfo_PoolParams proto.InternalMessageInfo

func (m *PoolParams) GetSwapFee() uint64 {
	if m != nil {
		return m.SwapFee
	}
	return 0
}

func (m *PoolParams) GetExitFee() uint64 {
	if m != nil {
		return m.ExitFee
	}
	return 0
}

func (m *PoolParams) GetSmoothWeightChangeParams() *SmoothWeightChangeParams {
	if m != nil {
		return m.SmoothWeightChangeParams
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolParams)(nil), "elysnetwork.elys.amm.PoolParams")
}

func init() { proto.RegisterFile("elys/amm/pool_params.proto", fileDescriptor_3500125990074bc9) }

var fileDescriptor_3500125990074bc9 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xcd, 0xa9, 0x2c,
	0xd6, 0x4f, 0xcc, 0xcd, 0xd5, 0x2f, 0xc8, 0xcf, 0xcf, 0x89, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x01, 0xc9, 0xe5, 0xa5, 0x96, 0x94, 0xe7, 0x17,
	0x65, 0xeb, 0x81, 0xd8, 0x7a, 0x89, 0xb9, 0xb9, 0x52, 0x5a, 0x70, 0x1d, 0xc5, 0xb9, 0xf9, 0xf9,
	0x25, 0x19, 0xf1, 0xe5, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0xf1, 0xc9, 0x19, 0x89, 0x79, 0xe9, 0xa9,
	0x28, 0x26, 0x28, 0xad, 0x61, 0xe4, 0xe2, 0x0a, 0xc8, 0xcf, 0xcf, 0x09, 0x00, 0x0b, 0x0a, 0x49,
	0x70, 0xb1, 0x17, 0x97, 0x27, 0x16, 0xb8, 0xa5, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04,
	0xc1, 0xb8, 0x20, 0x99, 0xd4, 0x8a, 0xcc, 0x12, 0x90, 0x0c, 0x13, 0x44, 0x06, 0xca, 0x15, 0xca,
	0xe2, 0x92, 0x80, 0xd8, 0x13, 0x0e, 0xb6, 0xc6, 0x19, 0x6c, 0x0b, 0xc4, 0x3c, 0x09, 0x66, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x3d, 0x3d, 0x6c, 0xee, 0xd4, 0x0b, 0xc6, 0xa1, 0x2b, 0x08, 0xa7, 0x79,
	0x4e, 0x4e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x91, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xb2, 0x41, 0x17, 0x6a, 0x1d, 0x98, 0xa3,
	0x5f, 0x01, 0x0e, 0x8e, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xcf, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x31, 0x2f, 0xe9, 0x32, 0x59, 0x01, 0x00, 0x00,
}

func (m *PoolParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SmoothWeightChangeParams != nil {
		{
			size, err := m.SmoothWeightChangeParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPoolParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ExitFee != 0 {
		i = encodeVarintPoolParams(dAtA, i, uint64(m.ExitFee))
		i--
		dAtA[i] = 0x10
	}
	if m.SwapFee != 0 {
		i = encodeVarintPoolParams(dAtA, i, uint64(m.SwapFee))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SwapFee != 0 {
		n += 1 + sovPoolParams(uint64(m.SwapFee))
	}
	if m.ExitFee != 0 {
		n += 1 + sovPoolParams(uint64(m.ExitFee))
	}
	if m.SmoothWeightChangeParams != nil {
		l = m.SmoothWeightChangeParams.Size()
		n += 1 + l + sovPoolParams(uint64(l))
	}
	return n
}

func sovPoolParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolParams(x uint64) (n int) {
	return sovPoolParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolParams
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
			return fmt.Errorf("proto: PoolParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			m.SwapFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SwapFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitFee", wireType)
			}
			m.ExitFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmoothWeightChangeParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
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
				return ErrInvalidLengthPoolParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SmoothWeightChangeParams == nil {
				m.SmoothWeightChangeParams = &SmoothWeightChangeParams{}
			}
			if err := m.SmoothWeightChangeParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolParams
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
func skipPoolParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolParams
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
					return 0, ErrIntOverflowPoolParams
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
					return 0, ErrIntOverflowPoolParams
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
				return 0, ErrInvalidLengthPoolParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolParams = fmt.Errorf("proto: unexpected end of group")
)