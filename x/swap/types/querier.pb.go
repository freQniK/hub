// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/swap/v1/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QuerySwapsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySwapsRequest) Reset()         { *m = QuerySwapsRequest{} }
func (m *QuerySwapsRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySwapsRequest) ProtoMessage()    {}
func (*QuerySwapsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0914180c1c187a7, []int{0}
}
func (m *QuerySwapsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapsRequest.Merge(m, src)
}
func (m *QuerySwapsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapsRequest proto.InternalMessageInfo

type QuerySwapRequest struct {
	TxHash []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (m *QuerySwapRequest) Reset()         { *m = QuerySwapRequest{} }
func (m *QuerySwapRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySwapRequest) ProtoMessage()    {}
func (*QuerySwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0914180c1c187a7, []int{1}
}
func (m *QuerySwapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapRequest.Merge(m, src)
}
func (m *QuerySwapRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapRequest proto.InternalMessageInfo

type QuerySwapsResponse struct {
	Swaps      []Swap              `protobuf:"bytes,1,rep,name=swaps,proto3" json:"swaps"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySwapsResponse) Reset()         { *m = QuerySwapsResponse{} }
func (m *QuerySwapsResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySwapsResponse) ProtoMessage()    {}
func (*QuerySwapsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0914180c1c187a7, []int{2}
}
func (m *QuerySwapsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapsResponse.Merge(m, src)
}
func (m *QuerySwapsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapsResponse proto.InternalMessageInfo

type QuerySwapResponse struct {
	Swap Swap `protobuf:"bytes,1,opt,name=swap,proto3" json:"swap"`
}

func (m *QuerySwapResponse) Reset()         { *m = QuerySwapResponse{} }
func (m *QuerySwapResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySwapResponse) ProtoMessage()    {}
func (*QuerySwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0914180c1c187a7, []int{3}
}
func (m *QuerySwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapResponse.Merge(m, src)
}
func (m *QuerySwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QuerySwapsRequest)(nil), "sentinel.swap.v1.QuerySwapsRequest")
	proto.RegisterType((*QuerySwapRequest)(nil), "sentinel.swap.v1.QuerySwapRequest")
	proto.RegisterType((*QuerySwapsResponse)(nil), "sentinel.swap.v1.QuerySwapsResponse")
	proto.RegisterType((*QuerySwapResponse)(nil), "sentinel.swap.v1.QuerySwapResponse")
}

func init() { proto.RegisterFile("sentinel/swap/v1/querier.proto", fileDescriptor_c0914180c1c187a7) }

var fileDescriptor_c0914180c1c187a7 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xe3, 0x31, 0x0a, 0x78, 0x13, 0x2a, 0x16, 0x82, 0xaa, 0x20, 0x33, 0x65, 0x08, 0x26,
	0xd0, 0x6c, 0x5a, 0xde, 0x60, 0x12, 0x7f, 0xee, 0x80, 0x70, 0x07, 0x17, 0xc8, 0x89, 0x3c, 0xc7,
	0x52, 0x67, 0x67, 0xb1, 0x93, 0x75, 0x42, 0xdc, 0xf0, 0x04, 0x20, 0x5e, 0x82, 0x47, 0xe9, 0xe5,
	0x24, 0x6e, 0xb8, 0x42, 0x90, 0x22, 0xf1, 0x1a, 0x28, 0x76, 0xd2, 0xb5, 0x54, 0x0a, 0x77, 0x96,
	0xcf, 0x77, 0xbe, 0xef, 0xe7, 0xe3, 0x03, 0xb1, 0xe1, 0xca, 0x4a, 0xc5, 0x27, 0xd4, 0x9c, 0xb0,
	0x8c, 0x96, 0x23, 0x7a, 0x5c, 0xf0, 0x5c, 0xf2, 0x9c, 0x64, 0xb9, 0xb6, 0x1a, 0xf5, 0xdb, 0x3a,
	0xa9, 0xeb, 0xa4, 0x1c, 0x0d, 0x1f, 0x24, 0xda, 0x1c, 0x69, 0x43, 0x63, 0x66, 0xb8, 0x13, 0x9f,
	0xd2, 0x72, 0x14, 0x73, 0xcb, 0x46, 0x34, 0x63, 0x42, 0x2a, 0x66, 0xa5, 0x56, 0xbe, 0x7b, 0x78,
	0x5d, 0x68, 0xa1, 0xdd, 0x91, 0xd6, 0xa7, 0xe6, 0xf6, 0xb6, 0xd0, 0x5a, 0x4c, 0x38, 0x65, 0x99,
	0xa4, 0x4c, 0x29, 0x6d, 0x5d, 0x8b, 0x69, 0xaa, 0xb7, 0xd6, 0x88, 0x5c, 0xb2, 0x2b, 0x86, 0x6f,
	0xe1, 0xb5, 0x57, 0x75, 0xe4, 0xeb, 0x13, 0x96, 0x99, 0x88, 0x1f, 0x17, 0xdc, 0x58, 0xf4, 0x14,
	0xc2, 0xf3, 0xe4, 0x01, 0xd8, 0x01, 0x7b, 0x5b, 0xe3, 0x7b, 0xc4, 0x63, 0x92, 0x1a, 0x93, 0x38,
	0x4c, 0xd2, 0x60, 0x92, 0x97, 0x4c, 0xf0, 0xa6, 0x37, 0x5a, 0xea, 0x0c, 0x1f, 0xc2, 0xfe, 0xc2,
	0xbc, 0xf5, 0xbe, 0x09, 0x2f, 0xd9, 0xe9, 0xbb, 0x94, 0x99, 0xd4, 0x19, 0x6f, 0x47, 0x3d, 0x3b,
	0x7d, 0xce, 0x4c, 0x1a, 0x7e, 0x06, 0x10, 0x2d, 0xa3, 0x98, 0x4c, 0x2b, 0xc3, 0xd1, 0x18, 0x5e,
	0xac, 0x71, 0xcd, 0x00, 0xec, 0x5c, 0xd8, 0xdb, 0x1a, 0xdf, 0x20, 0xff, 0xce, 0x8f, 0xd4, 0xfa,
	0x83, 0xcd, 0xd9, 0x8f, 0x3b, 0x41, 0xe4, 0xa5, 0xe8, 0xd9, 0x0a, 0xff, 0x86, 0xe3, 0xbf, 0xff,
	0x5f, 0x7e, 0x1f, 0xb8, 0xf2, 0x80, 0x27, 0x4b, 0xd3, 0x59, 0x10, 0x3d, 0x82, 0x9b, 0x75, 0x4c,
	0x33, 0x97, 0x6e, 0x20, 0xa7, 0x1c, 0xff, 0x01, 0x70, 0xdb, 0xfb, 0xf0, 0xbc, 0x94, 0x09, 0x47,
	0x02, 0xc2, 0xf3, 0xa7, 0xa2, 0xdd, 0x75, 0x8b, 0xb5, 0x3f, 0x19, 0xde, 0xed, 0x16, 0x79, 0xb6,
	0xf0, 0xea, 0xc7, 0x6f, 0xbf, 0xbf, 0x6c, 0x5c, 0x46, 0x3d, 0xea, 0x27, 0xa1, 0xe0, 0x95, 0x85,
	0x0a, 0x85, 0x1d, 0x16, 0x6d, 0xcc, 0x6e, 0xa7, 0xa6, 0x49, 0x19, 0xb8, 0x14, 0x84, 0xfa, 0x3e,
	0x85, 0xbe, 0x6f, 0x7e, 0xf4, 0xc3, 0xc1, 0x8b, 0xd9, 0x2f, 0x1c, 0x7c, 0xad, 0x70, 0x30, 0xab,
	0x30, 0x38, 0xab, 0x30, 0xf8, 0x59, 0x61, 0xf0, 0x69, 0x8e, 0x83, 0xb3, 0x39, 0x0e, 0xbe, 0xcf,
	0x71, 0xf0, 0x66, 0x5f, 0x48, 0x9b, 0x16, 0x31, 0x49, 0xf4, 0x11, 0x6d, 0xa3, 0xf6, 0xf5, 0xe1,
	0xa1, 0x4c, 0x24, 0x9b, 0xd0, 0xb4, 0x88, 0xe9, 0xd4, 0xef, 0xa9, 0x3d, 0xcd, 0xb8, 0x89, 0x7b,
	0x6e, 0x4d, 0x1f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xa4, 0xe2, 0xef, 0x57, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	QuerySwaps(ctx context.Context, in *QuerySwapsRequest, opts ...grpc.CallOption) (*QuerySwapsResponse, error)
	QuerySwap(ctx context.Context, in *QuerySwapRequest, opts ...grpc.CallOption) (*QuerySwapResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QuerySwaps(ctx context.Context, in *QuerySwapsRequest, opts ...grpc.CallOption) (*QuerySwapsResponse, error) {
	out := new(QuerySwapsResponse)
	err := c.cc.Invoke(ctx, "/sentinel.swap.v1.QueryService/QuerySwaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QuerySwap(ctx context.Context, in *QuerySwapRequest, opts ...grpc.CallOption) (*QuerySwapResponse, error) {
	out := new(QuerySwapResponse)
	err := c.cc.Invoke(ctx, "/sentinel.swap.v1.QueryService/QuerySwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	QuerySwaps(context.Context, *QuerySwapsRequest) (*QuerySwapsResponse, error)
	QuerySwap(context.Context, *QuerySwapRequest) (*QuerySwapResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QuerySwaps(ctx context.Context, req *QuerySwapsRequest) (*QuerySwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySwaps not implemented")
}
func (*UnimplementedQueryServiceServer) QuerySwap(ctx context.Context, req *QuerySwapRequest) (*QuerySwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySwap not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QuerySwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QuerySwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.swap.v1.QueryService/QuerySwaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QuerySwaps(ctx, req.(*QuerySwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QuerySwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QuerySwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.swap.v1.QueryService/QuerySwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QuerySwap(ctx, req.(*QuerySwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.swap.v1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QuerySwaps",
			Handler:    _QueryService_QuerySwaps_Handler,
		},
		{
			MethodName: "QuerySwap",
			Handler:    _QueryService_QuerySwap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/swap/v1/querier.proto",
}

func (m *QuerySwapsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySwapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySwapsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Swaps) > 0 {
		for iNdEx := len(m.Swaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Swaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuerier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QuerySwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Swap.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuerier(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuerySwapsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QuerySwapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QuerySwapsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Swaps) > 0 {
		for _, e := range m.Swaps {
			l = e.Size()
			n += 1 + l + sovQuerier(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QuerySwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Swap.Size()
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func sovQuerier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerier(x uint64) (n int) {
	return sovQuerier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySwapsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QuerySwapsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QuerySwapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QuerySwapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QuerySwapsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QuerySwapsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Swaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Swaps = append(m.Swaps, Swap{})
			if err := m.Swaps[len(m.Swaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QuerySwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QuerySwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Swap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Swap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func skipQuerier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
				return 0, ErrInvalidLengthQuerier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerier = fmt.Errorf("proto: unexpected end of group")
)