// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/feegrant/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	query "github.com/onomyprotocol/onomy-sdk/types/query"
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

// QueryFeeAllowanceRequest is the request type for the Query/FeeAllowance RPC method.
type QueryFeeAllowanceRequest struct {
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (m *QueryFeeAllowanceRequest) Reset()         { *m = QueryFeeAllowanceRequest{} }
func (m *QueryFeeAllowanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeeAllowanceRequest) ProtoMessage()    {}
func (*QueryFeeAllowanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{0}
}
func (m *QueryFeeAllowanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeAllowanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeAllowanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeAllowanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeAllowanceRequest.Merge(m, src)
}
func (m *QueryFeeAllowanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeAllowanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeAllowanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeAllowanceRequest proto.InternalMessageInfo

func (m *QueryFeeAllowanceRequest) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *QueryFeeAllowanceRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

// QueryFeeAllowanceResponse is the response type for the Query/FeeAllowance RPC method.
type QueryFeeAllowanceResponse struct {
	// fee_allowance is a fee_allowance granted for grantee by granter.
	FeeAllowance *FeeAllowanceGrant `protobuf:"bytes,1,opt,name=fee_allowance,json=feeAllowance,proto3" json:"fee_allowance,omitempty"`
}

func (m *QueryFeeAllowanceResponse) Reset()         { *m = QueryFeeAllowanceResponse{} }
func (m *QueryFeeAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeeAllowanceResponse) ProtoMessage()    {}
func (*QueryFeeAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{1}
}
func (m *QueryFeeAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeAllowanceResponse.Merge(m, src)
}
func (m *QueryFeeAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeAllowanceResponse proto.InternalMessageInfo

func (m *QueryFeeAllowanceResponse) GetFeeAllowance() *FeeAllowanceGrant {
	if m != nil {
		return m.FeeAllowance
	}
	return nil
}

// QueryFeeAllowancesRequest is the request type for the Query/FeeAllowances RPC method.
type QueryFeeAllowancesRequest struct {
	Grantee string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFeeAllowancesRequest) Reset()         { *m = QueryFeeAllowancesRequest{} }
func (m *QueryFeeAllowancesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeeAllowancesRequest) ProtoMessage()    {}
func (*QueryFeeAllowancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{2}
}
func (m *QueryFeeAllowancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeAllowancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeAllowancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeAllowancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeAllowancesRequest.Merge(m, src)
}
func (m *QueryFeeAllowancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeAllowancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeAllowancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeAllowancesRequest proto.InternalMessageInfo

func (m *QueryFeeAllowancesRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *QueryFeeAllowancesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryFeeAllowancesResponse is the response type for the Query/FeeAllowances RPC method.
type QueryFeeAllowancesResponse struct {
	// fee_allowances are fee_allowance's granted for grantee by granter.
	FeeAllowances []*FeeAllowanceGrant `protobuf:"bytes,1,rep,name=fee_allowances,json=feeAllowances,proto3" json:"fee_allowances,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFeeAllowancesResponse) Reset()         { *m = QueryFeeAllowancesResponse{} }
func (m *QueryFeeAllowancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeeAllowancesResponse) ProtoMessage()    {}
func (*QueryFeeAllowancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_59efc303945de53f, []int{3}
}
func (m *QueryFeeAllowancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeAllowancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeAllowancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeAllowancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeAllowancesResponse.Merge(m, src)
}
func (m *QueryFeeAllowancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeAllowancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeAllowancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeAllowancesResponse proto.InternalMessageInfo

func (m *QueryFeeAllowancesResponse) GetFeeAllowances() []*FeeAllowanceGrant {
	if m != nil {
		return m.FeeAllowances
	}
	return nil
}

func (m *QueryFeeAllowancesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryFeeAllowanceRequest)(nil), "cosmos.feegrant.v1beta1.QueryFeeAllowanceRequest")
	proto.RegisterType((*QueryFeeAllowanceResponse)(nil), "cosmos.feegrant.v1beta1.QueryFeeAllowanceResponse")
	proto.RegisterType((*QueryFeeAllowancesRequest)(nil), "cosmos.feegrant.v1beta1.QueryFeeAllowancesRequest")
	proto.RegisterType((*QueryFeeAllowancesResponse)(nil), "cosmos.feegrant.v1beta1.QueryFeeAllowancesResponse")
}

func init() {
	proto.RegisterFile("cosmos/feegrant/v1beta1/query.proto", fileDescriptor_59efc303945de53f)
}

var fileDescriptor_59efc303945de53f = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x22, 0x40, 0xf8, 0xae, 0x0c, 0x16, 0x12, 0x21, 0x42, 0xd1, 0x29, 0x48, 0x07,
	0x3a, 0xe9, 0x62, 0x35, 0x9d, 0x40, 0x2c, 0xdc, 0x70, 0xdd, 0x80, 0xcb, 0xc8, 0x82, 0x9c, 0xf2,
	0x6a, 0x22, 0x72, 0x71, 0x2e, 0x76, 0x81, 0x13, 0xea, 0xc2, 0x27, 0x40, 0xe2, 0xa3, 0xb0, 0xc0,
	0x37, 0x60, 0x3c, 0x89, 0x85, 0x11, 0xb5, 0x7c, 0x09, 0x36, 0x14, 0x3b, 0x6e, 0x52, 0x91, 0x00,
	0x99, 0xe2, 0xe4, 0xfd, 0xdf, 0x7b, 0xbf, 0xf7, 0xcf, 0x33, 0xbe, 0x33, 0x13, 0xf2, 0x54, 0x48,
	0x3a, 0x07, 0xe0, 0x05, 0xcb, 0x14, 0x7d, 0x3d, 0x8e, 0x41, 0xb1, 0x31, 0x3d, 0x5b, 0x40, 0x71,
	0x1e, 0xe4, 0x85, 0x50, 0x82, 0xdc, 0x34, 0xa2, 0xc0, 0x8a, 0x82, 0x4a, 0xe4, 0xde, 0xe0, 0x82,
	0x0b, 0xad, 0xa1, 0xe5, 0xc9, 0xc8, 0xdd, 0xfd, 0xae, 0x9a, 0x9b, 0x7c, 0xa3, 0x3b, 0xa8, 0x74,
	0x31, 0x93, 0x60, 0xfa, 0x6d, 0x94, 0x39, 0xe3, 0x49, 0xc6, 0x54, 0x22, 0xb2, 0x4a, 0x7b, 0x9b,
	0x0b, 0xc1, 0x53, 0xa0, 0x2c, 0x4f, 0x28, 0xcb, 0x32, 0xa1, 0x74, 0x50, 0x9a, 0xa8, 0xff, 0x18,
	0x3b, 0x27, 0x65, 0xfe, 0x31, 0xc0, 0xa3, 0x34, 0x15, 0x6f, 0x58, 0x36, 0x83, 0x08, 0xce, 0x16,
	0x20, 0x15, 0x71, 0xf0, 0x55, 0xdd, 0x14, 0x0a, 0x07, 0xed, 0xa1, 0x7b, 0xd7, 0x22, 0xfb, 0x5a,
	0x47, 0xc0, 0x19, 0x36, 0x23, 0xe0, 0xa7, 0xf8, 0x56, 0x4b, 0x3d, 0x99, 0x8b, 0x4c, 0x02, 0x79,
	0x82, 0x47, 0x73, 0x80, 0xe7, 0xcc, 0x06, 0x74, 0xd9, 0x9d, 0xf0, 0x20, 0xe8, 0x70, 0x29, 0x68,
	0x56, 0x99, 0x96, 0x91, 0x68, 0x77, 0xde, 0xf8, 0xe4, 0x2f, 0x5b, 0xba, 0xc9, 0x3f, 0xf0, 0x61,
	0x1b, 0x1f, 0xc8, 0x31, 0xc6, 0xb5, 0x4d, 0x7a, 0x82, 0x9d, 0x70, 0xdf, 0x42, 0x94, 0x9e, 0x06,
	0xe6, 0x1f, 0x5a, 0x8c, 0xa7, 0x8c, 0x5b, 0x53, 0xa2, 0x46, 0xa6, 0xff, 0x19, 0x61, 0xb7, 0xad,
	0x7f, 0x35, 0xee, 0x09, 0xbe, 0xbe, 0x35, 0xae, 0x74, 0xd0, 0xde, 0xa5, 0x9e, 0xf3, 0x8e, 0x9a,
	0xf3, 0x4a, 0x32, 0x6d, 0x21, 0xbf, 0xfb, 0x4f, 0x72, 0xc3, 0xd3, 0x44, 0x0f, 0x7f, 0x0d, 0xf1,
	0x65, 0x8d, 0x4e, 0xbe, 0x20, 0xbc, 0xdb, 0xec, 0x4b, 0xc6, 0x9d, 0x78, 0x5d, 0x9b, 0xe2, 0x86,
	0x7d, 0x52, 0x0c, 0x8d, 0x7f, 0xf4, 0xfe, 0xdb, 0xcf, 0x8f, 0xc3, 0x87, 0xe4, 0x01, 0xfd, 0xcb,
	0xd2, 0xd7, 0xe6, 0xd1, 0x77, 0xd5, 0xf2, 0x2d, 0xed, 0x09, 0x96, 0xe4, 0x13, 0xc2, 0xa3, 0x2d,
	0xef, 0x49, 0x0f, 0x12, 0xbb, 0x28, 0xee, 0xa4, 0x57, 0x4e, 0x85, 0x7f, 0x5f, 0xe3, 0x4f, 0xc8,
	0xf8, 0xff, 0xf0, 0x65, 0x4d, 0x7d, 0x34, 0xfd, 0xba, 0xf2, 0xd0, 0xc5, 0xca, 0x43, 0x3f, 0x56,
	0x1e, 0xfa, 0xb0, 0xf6, 0x06, 0x17, 0x6b, 0x6f, 0xf0, 0x7d, 0xed, 0x0d, 0x9e, 0x1d, 0xf2, 0x44,
	0xbd, 0x5c, 0xc4, 0xc1, 0x4c, 0x9c, 0xda, 0xb2, 0xe6, 0x71, 0x28, 0x5f, 0xbc, 0xa2, 0x6f, 0xeb,
	0x1e, 0xea, 0x3c, 0x07, 0x19, 0x5f, 0xd1, 0x77, 0x78, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb8,
	0x46, 0x4c, 0xba, 0x8b, 0x04, 0x00, 0x00,
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
	// FeeAllowance returns fee granted to the grantee by the granter.
	FeeAllowance(ctx context.Context, in *QueryFeeAllowanceRequest, opts ...grpc.CallOption) (*QueryFeeAllowanceResponse, error)
	// FeeAllowances returns all the grants for address.
	FeeAllowances(ctx context.Context, in *QueryFeeAllowancesRequest, opts ...grpc.CallOption) (*QueryFeeAllowancesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) FeeAllowance(ctx context.Context, in *QueryFeeAllowanceRequest, opts ...grpc.CallOption) (*QueryFeeAllowanceResponse, error) {
	out := new(QueryFeeAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Query/FeeAllowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FeeAllowances(ctx context.Context, in *QueryFeeAllowancesRequest, opts ...grpc.CallOption) (*QueryFeeAllowancesResponse, error) {
	out := new(QueryFeeAllowancesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Query/FeeAllowances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// FeeAllowance returns fee granted to the grantee by the granter.
	FeeAllowance(context.Context, *QueryFeeAllowanceRequest) (*QueryFeeAllowanceResponse, error)
	// FeeAllowances returns all the grants for address.
	FeeAllowances(context.Context, *QueryFeeAllowancesRequest) (*QueryFeeAllowancesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) FeeAllowance(ctx context.Context, req *QueryFeeAllowanceRequest) (*QueryFeeAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeAllowance not implemented")
}
func (*UnimplementedQueryServer) FeeAllowances(ctx context.Context, req *QueryFeeAllowancesRequest) (*QueryFeeAllowancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeAllowances not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_FeeAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeeAllowanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Query/FeeAllowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeAllowance(ctx, req.(*QueryFeeAllowanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FeeAllowances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeeAllowancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeAllowances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Query/FeeAllowances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeAllowances(ctx, req.(*QueryFeeAllowancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.feegrant.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FeeAllowance",
			Handler:    _Query_FeeAllowance_Handler,
		},
		{
			MethodName: "FeeAllowances",
			Handler:    _Query_FeeAllowances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/feegrant/v1beta1/query.proto",
}

func (m *QueryFeeAllowanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeAllowanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeAllowanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFeeAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeAllowance != nil {
		{
			size, err := m.FeeAllowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFeeAllowancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeAllowancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeAllowancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFeeAllowancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeAllowancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeAllowancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.FeeAllowances) > 0 {
		for iNdEx := len(m.FeeAllowances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeAllowances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryFeeAllowanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFeeAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FeeAllowance != nil {
		l = m.FeeAllowance.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFeeAllowancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFeeAllowancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeeAllowances) > 0 {
		for _, e := range m.FeeAllowances {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryFeeAllowanceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFeeAllowanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeAllowanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = string(dAtA[iNdEx:postIndex])
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
func (m *QueryFeeAllowanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFeeAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAllowance", wireType)
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
			if m.FeeAllowance == nil {
				m.FeeAllowance = &FeeAllowanceGrant{}
			}
			if err := m.FeeAllowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryFeeAllowancesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFeeAllowancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeAllowancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryFeeAllowancesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFeeAllowancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeAllowancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAllowances", wireType)
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
			m.FeeAllowances = append(m.FeeAllowances, &FeeAllowanceGrant{})
			if err := m.FeeAllowances[len(m.FeeAllowances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
