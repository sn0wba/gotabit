// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gotabit/sendMsg/v1beta1/sendMsg.proto

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

type SendMsg struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From    string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty" yaml:"from_address"`
	To      string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty" yaml:"to_address"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty" yaml:"msg"`
}

func (m *SendMsg) Reset()      { *m = SendMsg{} }
func (*SendMsg) ProtoMessage() {}
func (*SendMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbea64439c561480, []int{0}
}
func (m *SendMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsg.Merge(m, src)
}
func (m *SendMsg) XXX_Size() int {
	return m.Size()
}
func (m *SendMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsg proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SendMsg)(nil), "gotabit.sendMsg.v1beta1.SendMsg")
}

func init() {
	proto.RegisterFile("gotabit/sendMsg/v1beta1/sendMsg.proto", fileDescriptor_bbea64439c561480)
}

var fileDescriptor_bbea64439c561480 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xcf, 0x2f, 0x49,
	0x4c, 0xca, 0x2c, 0xd1, 0x2f, 0x4e, 0xcd, 0x4b, 0xf1, 0x2d, 0x4e, 0xd7, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0x84, 0xf1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xa1, 0xca, 0xf4,
	0x60, 0xc2, 0x50, 0x65, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x35, 0xfa, 0x20, 0x16, 0x44,
	0xb9, 0xd2, 0x42, 0x46, 0x2e, 0xf6, 0x60, 0x88, 0x4a, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x6d, 0x2e, 0x96, 0xb4, 0xa2, 0xfc,
	0x5c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xf1, 0x4f, 0xf7, 0xe4, 0x85, 0x2b, 0x13, 0x73,
	0x73, 0xac, 0x94, 0x40, 0xa2, 0xf1, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x4a, 0x41, 0x60,
	0x45, 0x42, 0xaa, 0x5c, 0x4c, 0x25, 0xf9, 0x12, 0xcc, 0x60, 0xa5, 0xa2, 0x9f, 0xee, 0xc9, 0x0b,
	0x42, 0x94, 0x96, 0xe4, 0x23, 0x14, 0x32, 0x95, 0xe4, 0x0b, 0x69, 0x70, 0xb1, 0xe7, 0xa6, 0x16,
	0x17, 0x27, 0xa6, 0xa7, 0x4a, 0xb0, 0x80, 0xd5, 0xf2, 0x7d, 0xba, 0x27, 0xcf, 0x05, 0x51, 0x9b,
	0x5b, 0x9c, 0xae, 0x14, 0x04, 0x93, 0xb6, 0xe2, 0xe8, 0x58, 0x20, 0xcf, 0x30, 0x63, 0x81, 0x3c,
	0x83, 0x93, 0xc7, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x0b, 0x22,
	0x18, 0x5d, 0x01, 0x0f, 0xac, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xa7, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x45, 0xae, 0x1a, 0x4c, 0x01, 0x00, 0x00,
}

func (m *SendMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintSendMsg(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintSendMsg(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintSendMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSendMsg(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSendMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovSendMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSendMsg(uint64(m.Id))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovSendMsg(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovSendMsg(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovSendMsg(uint64(l))
	}
	return n
}

func sovSendMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSendMsg(x uint64) (n int) {
	return sovSendMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSendMsg
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
			return fmt.Errorf("proto: SendMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSendMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSendMsg
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
				return ErrInvalidLengthSendMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSendMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSendMsg
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
				return ErrInvalidLengthSendMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSendMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSendMsg
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
				return ErrInvalidLengthSendMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSendMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSendMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSendMsg
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
func skipSendMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSendMsg
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
					return 0, ErrIntOverflowSendMsg
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
					return 0, ErrIntOverflowSendMsg
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
				return 0, ErrInvalidLengthSendMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSendMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSendMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSendMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSendMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSendMsg = fmt.Errorf("proto: unexpected end of group")
)