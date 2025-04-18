// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/cron/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryListPrivilegedRequest is request type for the Query/ListPrivileged RPC
// method.
type QueryListPrivilegedRequest struct {
}

func (m *QueryListPrivilegedRequest) Reset()         { *m = QueryListPrivilegedRequest{} }
func (m *QueryListPrivilegedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListPrivilegedRequest) ProtoMessage()    {}
func (*QueryListPrivilegedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f13a70ad4f0f33e, []int{0}
}
func (m *QueryListPrivilegedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListPrivilegedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListPrivilegedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListPrivilegedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListPrivilegedRequest.Merge(m, src)
}
func (m *QueryListPrivilegedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListPrivilegedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListPrivilegedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListPrivilegedRequest proto.InternalMessageInfo

// QueryListPrivilegedResponse is response type for the Query/ListPrivileged RPC
// method.
type QueryListPrivilegedResponse struct {
	// contract_addresses holds all the smart contract addresses which have
	// privilege status.
	ContractAddresses []string `protobuf:"bytes,1,rep,name=contract_addresses,json=contractAddresses,proto3" json:"privileged_contract_addresses,omitempty"`
}

func (m *QueryListPrivilegedResponse) Reset()         { *m = QueryListPrivilegedResponse{} }
func (m *QueryListPrivilegedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryListPrivilegedResponse) ProtoMessage()    {}
func (*QueryListPrivilegedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f13a70ad4f0f33e, []int{1}
}
func (m *QueryListPrivilegedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListPrivilegedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListPrivilegedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListPrivilegedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListPrivilegedResponse.Merge(m, src)
}
func (m *QueryListPrivilegedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryListPrivilegedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListPrivilegedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListPrivilegedResponse proto.InternalMessageInfo

func (m *QueryListPrivilegedResponse) GetContractAddresses() []string {
	if m != nil {
		return m.ContractAddresses
	}
	return nil
}

// QueryParamsRequest is request type for the Query/Params RPC
// method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f13a70ad4f0f33e, []int{2}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC
// method.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f13a70ad4f0f33e, []int{3}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryListPrivilegedRequest)(nil), "publicawesome.stargaze.cron.v1.QueryListPrivilegedRequest")
	proto.RegisterType((*QueryListPrivilegedResponse)(nil), "publicawesome.stargaze.cron.v1.QueryListPrivilegedResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "publicawesome.stargaze.cron.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "publicawesome.stargaze.cron.v1.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("publicawesome/stargaze/cron/v1/query.proto", fileDescriptor_5f13a70ad4f0f33e)
}

var fileDescriptor_5f13a70ad4f0f33e = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0xcd, 0xac, 0x58, 0x70, 0x04, 0xd1, 0x71, 0x0f, 0x21, 0x2e, 0x69, 0x8d, 0xa0, 0xbb, 0xea,
	0x66, 0x68, 0x7b, 0x10, 0xf4, 0x64, 0xcf, 0x0a, 0xeb, 0x9e, 0x64, 0x2f, 0x65, 0x9a, 0x0e, 0x71,
	0x20, 0xc9, 0x4c, 0x67, 0x26, 0xd1, 0x78, 0xf4, 0x17, 0x08, 0xfd, 0x2b, 0xfe, 0x88, 0x1e, 0x0b,
	0x7a, 0xf0, 0x54, 0xa4, 0xf5, 0xd4, 0x5f, 0x21, 0x9d, 0x49, 0xd4, 0x10, 0xb5, 0xea, 0x2d, 0xcc,
	0x7b, 0xef, 0x7b, 0xef, 0x7b, 0xf9, 0xe0, 0x7d, 0x91, 0x4f, 0x12, 0x16, 0x91, 0xd7, 0x54, 0xf1,
	0x94, 0x62, 0xa5, 0x89, 0x8c, 0xc9, 0x5b, 0x8a, 0x23, 0xc9, 0x33, 0x5c, 0xf4, 0xf1, 0x2c, 0xa7,
	0xb2, 0x0c, 0x85, 0xe4, 0x9a, 0x23, 0xbf, 0xc1, 0x0d, 0x6b, 0x6e, 0xb8, 0xe3, 0x86, 0x45, 0xdf,
	0x3b, 0x8c, 0x79, 0xcc, 0x0d, 0x15, 0xef, 0xbe, 0xac, 0xca, 0x3b, 0x8a, 0x39, 0x8f, 0x13, 0x8a,
	0x89, 0x60, 0x98, 0x64, 0x19, 0xd7, 0x44, 0x33, 0x9e, 0xa9, 0x0a, 0x3d, 0xd9, 0xe3, 0x6f, 0x66,
	0x1b, 0x6a, 0x70, 0x04, 0xbd, 0x17, 0xbb, 0x34, 0xcf, 0x98, 0xd2, 0x67, 0x92, 0x15, 0x2c, 0xa1,
	0x31, 0x9d, 0x9e, 0xd3, 0x59, 0x4e, 0x95, 0x0e, 0x4a, 0x78, 0xeb, 0x97, 0xa8, 0x12, 0x3c, 0x53,
	0x14, 0x5d, 0x40, 0x14, 0xf1, 0x4c, 0x4b, 0x12, 0xe9, 0x31, 0x99, 0x4e, 0x25, 0x55, 0x8a, 0x2a,
	0x17, 0xf4, 0x2e, 0x1d, 0x5f, 0x19, 0x3d, 0xd8, 0xae, 0xba, 0xf7, 0xc4, 0x77, 0xcd, 0xb8, 0x4d,
	0x7c, 0xc8, 0x53, 0xa6, 0x69, 0x2a, 0x74, 0x79, 0x7e, 0xa3, 0x46, 0x9f, 0xd6, 0x60, 0x70, 0x08,
	0x91, 0xb1, 0x3e, 0x23, 0x92, 0xa4, 0xaa, 0x0e, 0xc4, 0xe1, 0xcd, 0xc6, 0x6b, 0x15, 0xe4, 0x25,
	0xec, 0x08, 0xf3, 0xe2, 0x82, 0x1e, 0x38, 0xbe, 0x3a, 0xb8, 0x1b, 0xfe, 0xb9, 0xd5, 0xd0, 0xea,
	0x47, 0xee, 0x62, 0xd5, 0x75, 0xb6, 0xab, 0xee, 0x75, 0xab, 0xfe, 0x29, 0x51, 0x35, 0x6f, 0xf0,
	0xe9, 0x00, 0x5e, 0x36, 0x8e, 0xe8, 0x03, 0x80, 0xd7, 0x9a, 0x3d, 0xa0, 0xc7, 0xfb, 0x6c, 0x7e,
	0x5f, 0xad, 0xf7, 0xe4, 0xbf, 0xb4, 0x76, 0xdf, 0xe0, 0xe4, 0xdd, 0xc7, 0xaf, 0xf3, 0x83, 0x3b,
	0xe8, 0x76, 0xfb, 0xdf, 0x26, 0x4c, 0xe9, 0xd3, 0x1f, 0xbd, 0xa3, 0x39, 0x80, 0x1d, 0xbb, 0x2d,
	0x1a, 0xfc, 0x95, 0x65, 0xa3, 0x70, 0x6f, 0xf8, 0x4f, 0x9a, 0x2a, 0x5e, 0xcf, 0xc4, 0xf3, 0x90,
	0xdb, 0x8e, 0x67, 0x6b, 0x1d, 0x3d, 0x5f, 0xac, 0x7d, 0xb0, 0x5c, 0xfb, 0xe0, 0xcb, 0xda, 0x07,
	0xef, 0x37, 0xbe, 0xb3, 0xdc, 0xf8, 0xce, 0xe7, 0x8d, 0xef, 0x5c, 0x0c, 0x63, 0xa6, 0x5f, 0xe5,
	0x93, 0x30, 0xe2, 0x29, 0xb6, 0xd6, 0xa7, 0xad, 0x3b, 0x2e, 0xfa, 0x8f, 0xf0, 0x1b, 0x3b, 0x52,
	0x97, 0x82, 0xaa, 0x49, 0xc7, 0x1c, 0xf3, 0xf0, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x69,
	0xcd, 0x59, 0x79, 0x03, 0x00, 0x00,
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
	// ListPrivileged queries the contracts which have the priviledge status
	ListPrivileged(ctx context.Context, in *QueryListPrivilegedRequest, opts ...grpc.CallOption) (*QueryListPrivilegedResponse, error)
	// Params returns cron's module params.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ListPrivileged(ctx context.Context, in *QueryListPrivilegedRequest, opts ...grpc.CallOption) (*QueryListPrivilegedResponse, error) {
	out := new(QueryListPrivilegedResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.cron.v1.Query/ListPrivileged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.cron.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ListPrivileged queries the contracts which have the priviledge status
	ListPrivileged(context.Context, *QueryListPrivilegedRequest) (*QueryListPrivilegedResponse, error)
	// Params returns cron's module params.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ListPrivileged(ctx context.Context, req *QueryListPrivilegedRequest) (*QueryListPrivilegedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrivileged not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ListPrivileged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListPrivilegedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListPrivileged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.cron.v1.Query/ListPrivileged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListPrivileged(ctx, req.(*QueryListPrivilegedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.cron.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publicawesome.stargaze.cron.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPrivileged",
			Handler:    _Query_ListPrivileged_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publicawesome/stargaze/cron/v1/query.proto",
}

func (m *QueryListPrivilegedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListPrivilegedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListPrivilegedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryListPrivilegedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListPrivilegedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListPrivilegedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddresses) > 0 {
		for iNdEx := len(m.ContractAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ContractAddresses[iNdEx])
			copy(dAtA[i:], m.ContractAddresses[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QueryListPrivilegedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryListPrivilegedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ContractAddresses) > 0 {
		for _, s := range m.ContractAddresses {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryListPrivilegedRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryListPrivilegedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListPrivilegedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryListPrivilegedResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryListPrivilegedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListPrivilegedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddresses", wireType)
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
			m.ContractAddresses = append(m.ContractAddresses, string(dAtA[iNdEx:postIndex]))
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
