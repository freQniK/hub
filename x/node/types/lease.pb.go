// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v2/lease.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Lease struct {
	ID             uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeAddress    string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	AccountAddress string     `protobuf:"bytes,3,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	Hours          int64      `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	Price          types.Coin `protobuf:"bytes,5,opt,name=price,proto3" json:"price"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_6057f83a070f02c0, []int{0}
}
func (m *Lease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lease.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(m, src)
}
func (m *Lease) XXX_Size() int {
	return m.Size()
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Lease)(nil), "sentinel.node.v2.Lease")
}

func init() { proto.RegisterFile("sentinel/node/v2/lease.proto", fileDescriptor_6057f83a070f02c0) }

var fileDescriptor_6057f83a070f02c0 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4e, 0x32, 0x31,
	0x18, 0x86, 0xa7, 0x03, 0x43, 0xf2, 0x97, 0x3f, 0x6a, 0x26, 0xc4, 0x8c, 0xc4, 0x94, 0xd1, 0x8d,
	0xb3, 0xa1, 0x0d, 0x18, 0x0f, 0x20, 0xba, 0x31, 0x31, 0x31, 0x99, 0xa5, 0x1b, 0x33, 0xd3, 0x29,
	0xd0, 0x04, 0xfa, 0x91, 0x69, 0x87, 0xe8, 0x2d, 0x3c, 0x86, 0x47, 0xf0, 0x08, 0x2c, 0x59, 0xba,
	0x22, 0x3a, 0x5c, 0xc4, 0x94, 0x82, 0xbb, 0xaf, 0xcf, 0xf7, 0xb4, 0xcd, 0xfb, 0xe2, 0x73, 0x2d,
	0x94, 0x91, 0x4a, 0xcc, 0x98, 0x82, 0x42, 0xb0, 0xe5, 0x90, 0xcd, 0x44, 0xa6, 0x05, 0x5d, 0x94,
	0x60, 0x20, 0x3c, 0x39, 0x6c, 0xa9, 0xdd, 0xd2, 0xe5, 0xb0, 0x4b, 0x38, 0xe8, 0x39, 0x68, 0x96,
	0x67, 0x5a, 0xb0, 0xe5, 0x20, 0x17, 0x26, 0x1b, 0x30, 0x0e, 0x52, 0xb9, 0x1b, 0xdd, 0xce, 0x04,
	0x26, 0xb0, 0x1b, 0x99, 0x9d, 0x1c, 0xbd, 0xfc, 0x44, 0x38, 0x78, 0xb4, 0xef, 0x86, 0xa7, 0xd8,
	0x97, 0x45, 0x84, 0x62, 0x94, 0x34, 0x47, 0xad, 0x7a, 0xd3, 0xf3, 0x1f, 0xee, 0x53, 0x5f, 0x16,
	0xe1, 0x05, 0xfe, 0x6f, 0xbf, 0x78, 0xc9, 0x8a, 0xa2, 0x14, 0x5a, 0x47, 0x7e, 0x8c, 0x92, 0x7f,
	0x69, 0xdb, 0xb2, 0x5b, 0x87, 0xc2, 0x2b, 0x7c, 0x9c, 0x71, 0x0e, 0x95, 0x32, 0x7f, 0x56, 0x63,
	0x67, 0x1d, 0xed, 0xf1, 0x41, 0xec, 0xe0, 0x60, 0x0a, 0x55, 0xa9, 0xa3, 0x66, 0x8c, 0x92, 0x46,
	0xea, 0x0e, 0xe1, 0x0d, 0x0e, 0x16, 0xa5, 0xe4, 0x22, 0x0a, 0x62, 0x94, 0xb4, 0x87, 0x67, 0xd4,
	0x25, 0xa1, 0x36, 0x09, 0xdd, 0x27, 0xa1, 0x77, 0x20, 0xd5, 0xa8, 0xb9, 0xda, 0xf4, 0xbc, 0xd4,
	0xd9, 0xa3, 0xa7, 0xd5, 0x0f, 0xf1, 0x3e, 0x6a, 0xe2, 0xad, 0x6a, 0x82, 0xd6, 0x35, 0x41, 0xdf,
	0x35, 0x41, 0xef, 0x5b, 0xe2, 0xad, 0xb7, 0xc4, 0xfb, 0xda, 0x12, 0xef, 0xb9, 0x3f, 0x91, 0x66,
	0x5a, 0xe5, 0x94, 0xc3, 0x9c, 0x1d, 0xfa, 0xea, 0xc3, 0x78, 0x2c, 0xb9, 0xcc, 0x66, 0x6c, 0x5a,
	0xe5, 0xec, 0xd5, 0x95, 0x6b, 0xde, 0x16, 0x42, 0xe7, 0xad, 0x5d, 0x25, 0xd7, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x4b, 0xfa, 0x03, 0xe2, 0x7a, 0x01, 0x00, 0x00,
}

func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLease(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Hours != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AccountAddress) > 0 {
		i -= len(m.AccountAddress)
		copy(dAtA[i:], m.AccountAddress)
		i = encodeVarintLease(dAtA, i, uint64(len(m.AccountAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintLease(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLease(dAtA []byte, offset int, v uint64) int {
	offset -= sovLease(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLease(uint64(m.ID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	l = len(m.AccountAddress)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	if m.Hours != 0 {
		n += 1 + sovLease(uint64(m.Hours))
	}
	l = m.Price.Size()
	n += 1 + l + sovLease(uint64(l))
	return n
}

func sovLease(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLease(x uint64) (n int) {
	return sovLease(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
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
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
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
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
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
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
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
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLease
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
func skipLease(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLease
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
					return 0, ErrIntOverflowLease
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
					return 0, ErrIntOverflowLease
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
				return 0, ErrInvalidLengthLease
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLease
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLease
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLease        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLease          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLease = fmt.Errorf("proto: unexpected end of group")
)