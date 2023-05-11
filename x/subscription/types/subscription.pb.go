// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v2/subscription.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/sentinel-official/hub/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SubscriptionType int32

const (
	TypeUnspecified SubscriptionType = 0
	TypeNode        SubscriptionType = 1
	TypePlan        SubscriptionType = 2
)

var SubscriptionType_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_NODE",
	2: "TYPE_PLAN",
}

var SubscriptionType_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"TYPE_NODE":        1,
	"TYPE_PLAN":        2,
}

func (x SubscriptionType) String() string {
	return proto.EnumName(SubscriptionType_name, int32(x))
}

func (SubscriptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{0}
}

type BaseSubscription struct {
	ID       uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address  string       `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ExpiryAt time.Time    `protobuf:"bytes,3,opt,name=expiry_at,json=expiryAt,proto3,stdtime" json:"expiry_at"`
	Status   types.Status `protobuf:"varint,4,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt time.Time    `protobuf:"bytes,5,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *BaseSubscription) Reset()         { *m = BaseSubscription{} }
func (m *BaseSubscription) String() string { return proto.CompactTextString(m) }
func (*BaseSubscription) ProtoMessage()    {}
func (*BaseSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{0}
}
func (m *BaseSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseSubscription.Merge(m, src)
}
func (m *BaseSubscription) XXX_Size() int {
	return m.Size()
}
func (m *BaseSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_BaseSubscription proto.InternalMessageInfo

func (m *BaseSubscription) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *BaseSubscription) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BaseSubscription) GetExpiryAt() time.Time {
	if m != nil {
		return m.ExpiryAt
	}
	return time.Time{}
}

func (m *BaseSubscription) GetStatus() types.Status {
	if m != nil {
		return m.Status
	}
	return types.StatusUnspecified
}

func (m *BaseSubscription) GetStatusAt() time.Time {
	if m != nil {
		return m.StatusAt
	}
	return time.Time{}
}

type NodeSubscription struct {
	*BaseSubscription `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base,omitempty"`
	NodeAddress       string      `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	Gigabytes         int64       `protobuf:"varint,3,opt,name=gigabytes,proto3" json:"gigabytes,omitempty"`
	Hours             int64       `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	Deposit           types1.Coin `protobuf:"bytes,5,opt,name=deposit,proto3" json:"deposit"`
}

func (m *NodeSubscription) Reset()         { *m = NodeSubscription{} }
func (m *NodeSubscription) String() string { return proto.CompactTextString(m) }
func (*NodeSubscription) ProtoMessage()    {}
func (*NodeSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{1}
}
func (m *NodeSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeSubscription.Merge(m, src)
}
func (m *NodeSubscription) XXX_Size() int {
	return m.Size()
}
func (m *NodeSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_NodeSubscription proto.InternalMessageInfo

type PlanSubscription struct {
	*BaseSubscription `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base,omitempty"`
	PlanID            uint64 `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Denom             string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *PlanSubscription) Reset()         { *m = PlanSubscription{} }
func (m *PlanSubscription) String() string { return proto.CompactTextString(m) }
func (*PlanSubscription) ProtoMessage()    {}
func (*PlanSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6350e663da1ca66, []int{2}
}
func (m *PlanSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PlanSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PlanSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PlanSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanSubscription.Merge(m, src)
}
func (m *PlanSubscription) XXX_Size() int {
	return m.Size()
}
func (m *PlanSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_PlanSubscription proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("sentinel.subscription.v2.SubscriptionType", SubscriptionType_name, SubscriptionType_value)
	proto.RegisterType((*BaseSubscription)(nil), "sentinel.subscription.v2.BaseSubscription")
	proto.RegisterType((*NodeSubscription)(nil), "sentinel.subscription.v2.NodeSubscription")
	proto.RegisterType((*PlanSubscription)(nil), "sentinel.subscription.v2.PlanSubscription")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v2/subscription.proto", fileDescriptor_f6350e663da1ca66)
}

var fileDescriptor_f6350e663da1ca66 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xb1, 0x6f, 0xd3, 0x4e,
	0x18, 0xf5, 0xa5, 0x6e, 0xda, 0x5c, 0xab, 0xdf, 0xcf, 0x32, 0x15, 0x4a, 0x03, 0xb2, 0x43, 0x59,
	0x42, 0x11, 0x77, 0x4a, 0x10, 0x03, 0x6c, 0x71, 0x13, 0xa4, 0x48, 0x28, 0x44, 0x6e, 0x2b, 0x04,
	0x4b, 0x64, 0xc7, 0x17, 0xf7, 0xa4, 0xc4, 0x67, 0xe5, 0x2e, 0x51, 0x33, 0x30, 0xb1, 0x54, 0x9d,
	0x2a, 0x31, 0x57, 0x42, 0x62, 0xe1, 0x4f, 0xe9, 0xd8, 0x91, 0x29, 0x80, 0xbb, 0xf1, 0x57, 0xa0,
	0x3b, 0xdb, 0x6d, 0x83, 0xc4, 0xc0, 0xc0, 0x76, 0xdf, 0xbd, 0xf7, 0xf9, 0x7b, 0xef, 0xdd, 0x27,
	0xc3, 0xc7, 0x9c, 0x44, 0x82, 0x46, 0x64, 0x84, 0xf9, 0xd4, 0xe7, 0x83, 0x09, 0x8d, 0x05, 0x65,
	0x11, 0x9e, 0x35, 0x96, 0x6a, 0x14, 0x4f, 0x98, 0x60, 0x66, 0x39, 0x27, 0xa3, 0x25, 0x70, 0xd6,
	0xa8, 0x58, 0x03, 0xc6, 0xc7, 0x8c, 0x63, 0xdf, 0xe3, 0x04, 0xcf, 0xea, 0x3e, 0x11, 0x5e, 0x1d,
	0x0f, 0x18, 0xcd, 0x3a, 0x2b, 0x5b, 0x21, 0x0b, 0x99, 0x3a, 0x62, 0x79, 0xca, 0x6e, 0xed, 0x90,
	0xb1, 0x70, 0x44, 0xb0, 0xaa, 0xfc, 0xe9, 0x10, 0x0b, 0x3a, 0x26, 0x5c, 0x78, 0xe3, 0x38, 0x23,
	0x58, 0xd7, 0xea, 0xc4, 0x3c, 0x26, 0x1c, 0xcf, 0xea, 0x98, 0x0b, 0x4f, 0x4c, 0x79, 0x8a, 0xef,
	0x7c, 0x28, 0x40, 0xc3, 0xf1, 0x38, 0xd9, 0xbf, 0x25, 0xc7, 0xbc, 0x0b, 0x0b, 0x34, 0x28, 0x83,
	0x2a, 0xa8, 0xe9, 0x4e, 0x31, 0x59, 0xd8, 0x85, 0x4e, 0xcb, 0x2d, 0xd0, 0xc0, 0x2c, 0xc3, 0x35,
	0x2f, 0x08, 0x26, 0x84, 0xf3, 0x72, 0xa1, 0x0a, 0x6a, 0x25, 0x37, 0x2f, 0xcd, 0x26, 0x2c, 0x91,
	0xe3, 0x98, 0x4e, 0xe6, 0x7d, 0x4f, 0x94, 0x57, 0xaa, 0xa0, 0xb6, 0xd1, 0xa8, 0xa0, 0x54, 0x1b,
	0xca, 0xb5, 0xa1, 0x83, 0x5c, 0x9b, 0xb3, 0x7e, 0xb1, 0xb0, 0xb5, 0xb3, 0x6f, 0x36, 0x70, 0xd7,
	0xd3, 0xb6, 0xa6, 0x30, 0xeb, 0xb0, 0x98, 0x2a, 0x2b, 0xeb, 0x55, 0x50, 0xfb, 0xaf, 0xb1, 0x8d,
	0xae, 0xb3, 0x52, 0xd2, 0xd1, 0xac, 0x8e, 0xf6, 0x15, 0xc1, 0xcd, 0x88, 0x72, 0x6a, 0x7a, 0x92,
	0x53, 0x57, 0xff, 0x66, 0x6a, 0xda, 0xd6, 0x14, 0x2f, 0xf4, 0x93, 0x4f, 0x36, 0xd8, 0xf9, 0x09,
	0xa0, 0xd1, 0x65, 0xc1, 0x72, 0x0a, 0x2d, 0xa8, 0xcb, 0xc7, 0x50, 0x39, 0x6c, 0x34, 0x76, 0xd1,
	0x9f, 0x9e, 0x0e, 0xfd, 0x9e, 0x9f, 0xa3, 0x5f, 0x2e, 0x6c, 0xe0, 0xaa, 0x6e, 0xf3, 0x01, 0xdc,
	0x8c, 0x58, 0x40, 0xfa, 0xcb, 0xc1, 0x6d, 0xc8, 0xbb, 0x66, 0x16, 0xde, 0x7d, 0x58, 0x0a, 0x69,
	0xe8, 0xf9, 0x73, 0x41, 0xb8, 0x0a, 0x6f, 0xc5, 0xbd, 0xb9, 0x30, 0xb7, 0xe0, 0xea, 0x11, 0x9b,
	0x4e, 0xd2, 0x58, 0x56, 0xdc, 0xb4, 0x30, 0x9f, 0xc3, 0xb5, 0x80, 0xc4, 0x8c, 0xd3, 0xdc, 0xf8,
	0x36, 0x4a, 0x17, 0x08, 0xc9, 0xa9, 0x28, 0x5b, 0x20, 0xb4, 0xc7, 0x68, 0xe4, 0xe8, 0xd2, 0xb7,
	0x9b, 0xf3, 0x77, 0x3e, 0x02, 0x68, 0xf4, 0x46, 0x5e, 0xf4, 0x0f, 0xcc, 0x3e, 0x84, 0x6b, 0xf1,
	0xc8, 0x8b, 0xfa, 0x34, 0x50, 0x3e, 0x75, 0x07, 0x26, 0x0b, 0xbb, 0x28, 0x87, 0x75, 0x5a, 0x6e,
	0x51, 0x42, 0x9d, 0x40, 0x1a, 0x0a, 0x48, 0xc4, 0xc6, 0xca, 0x6a, 0xc9, 0x4d, 0x8b, 0xdd, 0xf7,
	0xd0, 0xb8, 0xfd, 0xd9, 0x83, 0x79, 0x4c, 0xcc, 0x47, 0xd0, 0x38, 0x78, 0xdb, 0x6b, 0xf7, 0x0f,
	0xbb, 0xfb, 0xbd, 0xf6, 0x5e, 0xe7, 0x65, 0xa7, 0xdd, 0x32, 0xb4, 0xca, 0x9d, 0xd3, 0xf3, 0xea,
	0xff, 0x12, 0x3f, 0x8c, 0x78, 0x4c, 0x06, 0x74, 0x48, 0x49, 0x60, 0xde, 0x83, 0x25, 0x45, 0xed,
	0xbe, 0x6e, 0xb5, 0x0d, 0x50, 0xd9, 0x3c, 0x3d, 0xaf, 0xae, 0x4b, 0x8e, 0x7c, 0xd5, 0x6b, 0xb0,
	0xf7, 0xaa, 0xd9, 0x35, 0x0a, 0x37, 0xa0, 0x14, 0x56, 0xd1, 0x4f, 0x3e, 0x5b, 0x9a, 0xf3, 0xe6,
	0xe2, 0x87, 0xa5, 0x7d, 0x49, 0x2c, 0xed, 0x22, 0xb1, 0xc0, 0x65, 0x62, 0x81, 0xef, 0x89, 0x05,
	0xce, 0xae, 0x2c, 0xed, 0xf2, 0xca, 0xd2, 0xbe, 0x5e, 0x59, 0xda, 0xbb, 0x67, 0x21, 0x15, 0x47,
	0x53, 0x1f, 0x0d, 0xd8, 0x18, 0xe7, 0xe9, 0x3c, 0x61, 0xc3, 0x21, 0x1d, 0x50, 0x6f, 0x84, 0x8f,
	0xa6, 0x3e, 0x3e, 0x5e, 0xfe, 0x03, 0xa8, 0xad, 0xf5, 0x8b, 0x6a, 0x11, 0x9f, 0xfe, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0x06, 0x4b, 0x99, 0x27, 0x04, 0x00, 0x00,
}

func (m *BaseSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSubscription(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.Status != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpiryAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiryAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSubscription(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NodeSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSubscription(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Hours != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x20
	}
	if m.Gigabytes != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Gigabytes))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseSubscription != nil {
		{
			size, err := m.BaseSubscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubscription(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PlanSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlanSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PlanSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PlanID != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PlanID))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseSubscription != nil {
		{
			size, err := m.BaseSubscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubscription(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubscription(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubscription(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovSubscription(uint64(m.ID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiryAt)
	n += 1 + l + sovSubscription(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSubscription(uint64(m.Status))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovSubscription(uint64(l))
	return n
}

func (m *NodeSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseSubscription != nil {
		l = m.BaseSubscription.Size()
		n += 1 + l + sovSubscription(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.Gigabytes != 0 {
		n += 1 + sovSubscription(uint64(m.Gigabytes))
	}
	if m.Hours != 0 {
		n += 1 + sovSubscription(uint64(m.Hours))
	}
	l = m.Deposit.Size()
	n += 1 + l + sovSubscription(uint64(l))
	return n
}

func (m *PlanSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseSubscription != nil {
		l = m.BaseSubscription.Size()
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.PlanID != 0 {
		n += 1 + sovSubscription(uint64(m.PlanID))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	return n
}

func sovSubscription(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubscription(x uint64) (n int) {
	return sovSubscription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: BaseSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpiryAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func (m *NodeSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: NodeSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseSubscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseSubscription == nil {
				m.BaseSubscription = &BaseSubscription{}
			}
			if err := m.BaseSubscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gigabytes", wireType)
			}
			m.Gigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func (m *PlanSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: PlanSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlanSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseSubscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseSubscription == nil {
				m.BaseSubscription = &BaseSubscription{}
			}
			if err := m.BaseSubscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanID", wireType)
			}
			m.PlanID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func skipSubscription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
				return 0, ErrInvalidLengthSubscription
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubscription
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubscription
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubscription        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubscription          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubscription = fmt.Errorf("proto: unexpected end of group")
)
