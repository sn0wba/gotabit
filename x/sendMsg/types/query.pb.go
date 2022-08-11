// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gotabit/sendMsg/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

// SentMessagesRequest is request type for the Query/SentMessages RPC method
type SentMessagesRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *SentMessagesRequest) Reset()         { *m = SentMessagesRequest{} }
func (m *SentMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*SentMessagesRequest) ProtoMessage()    {}
func (*SentMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32a3f8b73d4cbf60, []int{0}
}
func (m *SentMessagesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SentMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SentMessagesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SentMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentMessagesRequest.Merge(m, src)
}
func (m *SentMessagesRequest) XXX_Size() int {
	return m.Size()
}
func (m *SentMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SentMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SentMessagesRequest proto.InternalMessageInfo

func (m *SentMessagesRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// SentMessagesResponse is response type for the Query/SentMessages RPC method
type SentMessagesResponse struct {
	Messages []*SendMsg `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *SentMessagesResponse) Reset()         { *m = SentMessagesResponse{} }
func (m *SentMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*SentMessagesResponse) ProtoMessage()    {}
func (*SentMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32a3f8b73d4cbf60, []int{1}
}
func (m *SentMessagesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SentMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SentMessagesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SentMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentMessagesResponse.Merge(m, src)
}
func (m *SentMessagesResponse) XXX_Size() int {
	return m.Size()
}
func (m *SentMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SentMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SentMessagesResponse proto.InternalMessageInfo

func (m *SentMessagesResponse) GetMessages() []*SendMsg {
	if m != nil {
		return m.Messages
	}
	return nil
}

// ReceivedMessagesRequest is request type for the Query/ReceivedMessages RPC method
type ReceivedMessagesRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ReceivedMessagesRequest) Reset()         { *m = ReceivedMessagesRequest{} }
func (m *ReceivedMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*ReceivedMessagesRequest) ProtoMessage()    {}
func (*ReceivedMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32a3f8b73d4cbf60, []int{2}
}
func (m *ReceivedMessagesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceivedMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceivedMessagesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceivedMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedMessagesRequest.Merge(m, src)
}
func (m *ReceivedMessagesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ReceivedMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedMessagesRequest proto.InternalMessageInfo

func (m *ReceivedMessagesRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// ReceivedMessagesResponse is response type for the Query/ReceivedMessages RPC method
type ReceivedMessagesResponse struct {
	Messages []*SendMsg `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *ReceivedMessagesResponse) Reset()         { *m = ReceivedMessagesResponse{} }
func (m *ReceivedMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*ReceivedMessagesResponse) ProtoMessage()    {}
func (*ReceivedMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32a3f8b73d4cbf60, []int{3}
}
func (m *ReceivedMessagesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceivedMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceivedMessagesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceivedMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedMessagesResponse.Merge(m, src)
}
func (m *ReceivedMessagesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ReceivedMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedMessagesResponse proto.InternalMessageInfo

func (m *ReceivedMessagesResponse) GetMessages() []*SendMsg {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*SentMessagesRequest)(nil), "gotabit.sendMsg.v1beta1.SentMessagesRequest")
	proto.RegisterType((*SentMessagesResponse)(nil), "gotabit.sendMsg.v1beta1.SentMessagesResponse")
	proto.RegisterType((*ReceivedMessagesRequest)(nil), "gotabit.sendMsg.v1beta1.ReceivedMessagesRequest")
	proto.RegisterType((*ReceivedMessagesResponse)(nil), "gotabit.sendMsg.v1beta1.ReceivedMessagesResponse")
}

func init() {
	proto.RegisterFile("gotabit/sendMsg/v1beta1/query.proto", fileDescriptor_32a3f8b73d4cbf60)
}

var fileDescriptor_32a3f8b73d4cbf60 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4b, 0xe3, 0x40,
	0x1c, 0xc6, 0x3b, 0x5d, 0xf6, 0x6d, 0x76, 0x0f, 0xcb, 0x6c, 0xa1, 0x21, 0x2c, 0xa1, 0x64, 0x59,
	0xb6, 0xdb, 0xdd, 0x66, 0xb6, 0xed, 0xd5, 0x93, 0x78, 0xed, 0xc1, 0xd4, 0x83, 0x78, 0x9b, 0x24,
	0x7f, 0xc6, 0x80, 0xcd, 0xa4, 0x99, 0x69, 0xb1, 0x88, 0x17, 0x3f, 0x81, 0xe0, 0xd5, 0x8f, 0xe0,
	0x07, 0xf1, 0x58, 0x10, 0xc1, 0xa3, 0xb4, 0x7e, 0x10, 0x69, 0x3a, 0x89, 0xb5, 0x92, 0x52, 0xc1,
	0x53, 0xe6, 0xe5, 0x79, 0x9e, 0xfc, 0xf2, 0xfc, 0x83, 0x7f, 0x72, 0xa1, 0x98, 0x17, 0x2a, 0x2a,
	0x21, 0x0a, 0xba, 0x92, 0xd3, 0x51, 0xcb, 0x03, 0xc5, 0x5a, 0x74, 0x30, 0x84, 0x64, 0xec, 0xc4,
	0x89, 0x50, 0x82, 0x54, 0xb5, 0xc8, 0xd1, 0x22, 0x47, 0x8b, 0x4c, 0xcb, 0x17, 0xb2, 0x2f, 0x24,
	0xf5, 0x98, 0x84, 0xdc, 0xe9, 0x8b, 0x30, 0x5a, 0x18, 0xcd, 0xc6, 0xf2, 0x7d, 0x9a, 0x98, 0xab,
	0x62, 0xc6, 0xc3, 0x88, 0xa9, 0x50, 0x64, 0xda, 0x0a, 0x17, 0x5c, 0xa4, 0x4b, 0x3a, 0x5f, 0xe9,
	0xd3, 0x1f, 0x5c, 0x08, 0x7e, 0x04, 0x94, 0xc5, 0x21, 0x65, 0x51, 0x24, 0x54, 0x6a, 0x91, 0xfa,
	0xf6, 0x57, 0x11, 0x7d, 0x06, 0x9a, 0xca, 0x6c, 0x8a, 0xbf, 0xf7, 0x20, 0x52, 0x5d, 0x90, 0x92,
	0x71, 0x90, 0x2e, 0x0c, 0x86, 0x20, 0x15, 0x31, 0xf0, 0x47, 0x16, 0x04, 0x09, 0x48, 0x69, 0xa0,
	0x1a, 0xaa, 0x7f, 0x76, 0xb3, 0xad, 0xbd, 0x87, 0x2b, 0xcf, 0x0d, 0x32, 0x16, 0x91, 0x04, 0xb2,
	0x85, 0x3f, 0xf5, 0xf5, 0x99, 0x81, 0x6a, 0xef, 0xea, 0x5f, 0xda, 0x35, 0xa7, 0xa0, 0x1b, 0xa7,
	0xb7, 0xd8, 0xbb, 0xb9, 0xc3, 0xee, 0xe0, 0xaa, 0x0b, 0x3e, 0x84, 0x23, 0x08, 0x36, 0x47, 0xd9,
	0xc7, 0xc6, 0x4b, 0xd3, 0x5b, 0xe0, 0xb4, 0x6f, 0xcb, 0xf8, 0xfd, 0xee, 0x7c, 0x26, 0xe4, 0x12,
	0xe1, 0xaf, 0xcb, 0xdf, 0x4b, 0xfe, 0xad, 0x8b, 0x59, 0xed, 0xd1, 0x6c, 0x6e, 0xa8, 0x5e, 0x50,
	0xdb, 0xf4, 0xec, 0xe6, 0xe1, 0xa2, 0xfc, 0x87, 0xfc, 0xa6, 0x6b, 0xa6, 0xa7, 0xe8, 0x89, 0x6e,
	0xe0, 0x94, 0x5c, 0x21, 0xfc, 0x6d, 0xb5, 0x03, 0xf2, 0xbf, 0xf0, 0xa5, 0x05, 0x1d, 0x9b, 0xad,
	0x57, 0x38, 0x34, 0x6a, 0x27, 0x45, 0x6d, 0x92, 0xbf, 0x85, 0xa8, 0x89, 0xb6, 0x3e, 0xe1, 0x6e,
	0xef, 0x5c, 0x4f, 0x2d, 0x34, 0x99, 0x5a, 0xe8, 0x7e, 0x6a, 0xa1, 0xf3, 0x99, 0x55, 0x9a, 0xcc,
	0xac, 0xd2, 0xdd, 0xcc, 0x2a, 0x1d, 0x34, 0x78, 0xa8, 0x0e, 0x87, 0x9e, 0xe3, 0x8b, 0x7e, 0x1e,
	0x98, 0x3d, 0x8f, 0xf3, 0x68, 0x35, 0x8e, 0x41, 0x7a, 0x1f, 0xd2, 0x5f, 0xb7, 0xf3, 0x18, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0x91, 0x0b, 0xec, 0xa1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// SentMessages returns messages sent from user
	SentMessages(ctx context.Context, in *SentMessagesRequest, opts ...grpc.CallOption) (*SentMessagesResponse, error)
	// ReceivedMessages returns messages received by user
	ReceivedMessages(ctx context.Context, in *ReceivedMessagesRequest, opts ...grpc.CallOption) (*ReceivedMessagesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) SentMessages(ctx context.Context, in *SentMessagesRequest, opts ...grpc.CallOption) (*SentMessagesResponse, error) {
	out := new(SentMessagesResponse)
	err := c.cc.Invoke(ctx, "/gotabit.sendMsg.v1beta1.Query/SentMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ReceivedMessages(ctx context.Context, in *ReceivedMessagesRequest, opts ...grpc.CallOption) (*ReceivedMessagesResponse, error) {
	out := new(ReceivedMessagesResponse)
	err := c.cc.Invoke(ctx, "/gotabit.sendMsg.v1beta1.Query/ReceivedMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// SentMessages returns messages sent from user
	SentMessages(context.Context, *SentMessagesRequest) (*SentMessagesResponse, error)
	// ReceivedMessages returns messages received by user
	ReceivedMessages(context.Context, *ReceivedMessagesRequest) (*ReceivedMessagesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) SentMessages(ctx context.Context, req *SentMessagesRequest) (*SentMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SentMessages not implemented")
}
func (*UnimplementedQueryServer) ReceivedMessages(ctx context.Context, req *ReceivedMessagesRequest) (*ReceivedMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceivedMessages not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_SentMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SentMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gotabit.sendMsg.v1beta1.Query/SentMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SentMessages(ctx, req.(*SentMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ReceivedMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceivedMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ReceivedMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gotabit.sendMsg.v1beta1.Query/ReceivedMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ReceivedMessages(ctx, req.(*ReceivedMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gotabit.sendMsg.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SentMessages",
			Handler:    _Query_SentMessages_Handler,
		},
		{
			MethodName: "ReceivedMessages",
			Handler:    _Query_ReceivedMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gotabit/sendMsg/v1beta1/query.proto",
}

func (m *SentMessagesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SentMessagesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SentMessagesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SentMessagesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SentMessagesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SentMessagesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReceivedMessagesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceivedMessagesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceivedMessagesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReceivedMessagesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceivedMessagesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceivedMessagesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SentMessagesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *SentMessagesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *ReceivedMessagesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ReceivedMessagesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SentMessagesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SentMessagesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SentMessagesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *SentMessagesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SentMessagesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SentMessagesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &SendMsg{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ReceivedMessagesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ReceivedMessagesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceivedMessagesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ReceivedMessagesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ReceivedMessagesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceivedMessagesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &SendMsg{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)