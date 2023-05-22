// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/amm/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_763da04a7298bbac, []int{0}
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

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_763da04a7298bbac, []int{1}
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

type QueryGetPoolRequest struct {
	PoolId uint64 `protobuf:"varint,1,opt,name=poolId,proto3" json:"poolId,omitempty"`
}

func (m *QueryGetPoolRequest) Reset()         { *m = QueryGetPoolRequest{} }
func (m *QueryGetPoolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetPoolRequest) ProtoMessage()    {}
func (*QueryGetPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_763da04a7298bbac, []int{2}
}
func (m *QueryGetPoolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPoolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPoolRequest.Merge(m, src)
}
func (m *QueryGetPoolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPoolRequest proto.InternalMessageInfo

func (m *QueryGetPoolRequest) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

type QueryGetPoolResponse struct {
	Pool Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool"`
}

func (m *QueryGetPoolResponse) Reset()         { *m = QueryGetPoolResponse{} }
func (m *QueryGetPoolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetPoolResponse) ProtoMessage()    {}
func (*QueryGetPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_763da04a7298bbac, []int{3}
}
func (m *QueryGetPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPoolResponse.Merge(m, src)
}
func (m *QueryGetPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPoolResponse proto.InternalMessageInfo

func (m *QueryGetPoolResponse) GetPool() Pool {
	if m != nil {
		return m.Pool
	}
	return Pool{}
}

type QueryAllPoolRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPoolRequest) Reset()         { *m = QueryAllPoolRequest{} }
func (m *QueryAllPoolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPoolRequest) ProtoMessage()    {}
func (*QueryAllPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_763da04a7298bbac, []int{4}
}
func (m *QueryAllPoolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPoolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPoolRequest.Merge(m, src)
}
func (m *QueryAllPoolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPoolRequest proto.InternalMessageInfo

func (m *QueryAllPoolRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPoolResponse struct {
	Pool       []Pool              `protobuf:"bytes,1,rep,name=pool,proto3" json:"pool"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPoolResponse) Reset()         { *m = QueryAllPoolResponse{} }
func (m *QueryAllPoolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPoolResponse) ProtoMessage()    {}
func (*QueryAllPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_763da04a7298bbac, []int{5}
}
func (m *QueryAllPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPoolResponse.Merge(m, src)
}
func (m *QueryAllPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPoolResponse proto.InternalMessageInfo

func (m *QueryAllPoolResponse) GetPool() []Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

func (m *QueryAllPoolResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "elysnetwork.elys.amm.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "elysnetwork.elys.amm.QueryParamsResponse")
	proto.RegisterType((*QueryGetPoolRequest)(nil), "elysnetwork.elys.amm.QueryGetPoolRequest")
	proto.RegisterType((*QueryGetPoolResponse)(nil), "elysnetwork.elys.amm.QueryGetPoolResponse")
	proto.RegisterType((*QueryAllPoolRequest)(nil), "elysnetwork.elys.amm.QueryAllPoolRequest")
	proto.RegisterType((*QueryAllPoolResponse)(nil), "elysnetwork.elys.amm.QueryAllPoolResponse")
}

func init() { proto.RegisterFile("elys/amm/query.proto", fileDescriptor_763da04a7298bbac) }

var fileDescriptor_763da04a7298bbac = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0x13, 0x3f,
	0x10, 0xc6, 0xe3, 0x36, 0xff, 0xfc, 0x25, 0x73, 0x73, 0x03, 0x82, 0xa5, 0xdd, 0xa2, 0x05, 0x4a,
	0xa8, 0xa8, 0xad, 0x16, 0x4e, 0xdc, 0x9a, 0x03, 0x15, 0x12, 0x87, 0x34, 0x47, 0x24, 0x84, 0x9c,
	0x60, 0x2d, 0x2b, 0xbc, 0x3b, 0xdb, 0xd8, 0x01, 0x22, 0xc4, 0x05, 0xc1, 0x89, 0x0b, 0x12, 0xe2,
	0x0d, 0x78, 0x98, 0x1e, 0x2b, 0x71, 0xe1, 0x84, 0x50, 0xc2, 0x83, 0x20, 0xdb, 0xd3, 0x36, 0xdb,
	0xa6, 0x21, 0x37, 0x67, 0xfc, 0xcd, 0x37, 0xbf, 0x8c, 0xbf, 0xa5, 0x4d, 0xa5, 0x47, 0x46, 0xc8,
	0x3c, 0x17, 0x07, 0x43, 0x35, 0x18, 0xf1, 0x72, 0x00, 0x16, 0x98, 0xaf, 0x16, 0xca, 0xbe, 0x81,
	0xc1, 0x2b, 0xee, 0xce, 0x5c, 0xe6, 0x79, 0xd4, 0x4c, 0x21, 0x05, 0x2f, 0x10, 0xee, 0x14, 0xb4,
	0xd1, 0x6a, 0x0a, 0x90, 0x6a, 0x25, 0x64, 0x99, 0x09, 0x59, 0x14, 0x60, 0xa5, 0xcd, 0xa0, 0x30,
	0x78, 0xbb, 0xd9, 0x07, 0x93, 0x83, 0x11, 0x3d, 0x69, 0x54, 0x18, 0x21, 0x5e, 0x6f, 0xf7, 0x94,
	0x95, 0xdb, 0xa2, 0x94, 0x69, 0x56, 0x78, 0x31, 0x6a, 0x2f, 0x9f, 0xb0, 0x94, 0x72, 0x20, 0xf3,
	0x63, 0x8b, 0x95, 0xd3, 0x32, 0x80, 0xc6, 0xe2, 0xb5, 0x4a, 0xf1, 0xb9, 0x34, 0x46, 0x59, 0xbc,
	0x8a, 0xaa, 0x57, 0x15, 0xaf, 0x78, 0x1a, 0xe7, 0x18, 0xa4, 0x0f, 0x19, 0x22, 0x24, 0x4d, 0xca,
	0xf6, 0x1d, 0x64, 0xc7, 0x37, 0x75, 0xd5, 0xc1, 0x50, 0x19, 0x9b, 0xec, 0xd3, 0x95, 0x4a, 0xd5,
	0x94, 0x50, 0x18, 0xc5, 0x1e, 0xd2, 0x46, 0x30, 0xbf, 0x4a, 0x6e, 0x90, 0xd6, 0xa5, 0x9d, 0x55,
	0x3e, 0x6b, 0x6d, 0x3c, 0x74, 0xb5, 0xeb, 0x87, 0xbf, 0xd6, 0x6b, 0x5d, 0xec, 0x48, 0xb6, 0xd0,
	0x72, 0x4f, 0xd9, 0x0e, 0x80, 0xc6, 0x49, 0xec, 0x0a, 0x6d, 0x38, 0xe8, 0xc7, 0x2f, 0xbc, 0x65,
	0xbd, 0x8b, 0xbf, 0x92, 0x27, 0xb4, 0x59, 0x95, 0x23, 0xc2, 0x03, 0x5a, 0x77, 0x0a, 0x04, 0x88,
	0x2e, 0x00, 0x00, 0xd0, 0x38, 0xde, 0xab, 0x93, 0x67, 0x38, 0x7c, 0x57, 0xeb, 0xe9, 0xe1, 0x8f,
	0x28, 0x3d, 0x7d, 0x13, 0xb4, 0xdc, 0xe0, 0x61, 0x63, 0xdc, 0x6d, 0x8c, 0x87, 0x8c, 0xe0, 0xde,
	0x78, 0x47, 0xa6, 0x0a, 0x7b, 0xbb, 0x53, 0x9d, 0xc9, 0x37, 0x82, 0xb4, 0x27, 0xfe, 0xe7, 0x68,
	0x97, 0x17, 0xa7, 0x65, 0x7b, 0x15, 0xac, 0x25, 0x8f, 0x75, 0xe7, 0x9f, 0x58, 0x61, 0xe4, 0x34,
	0xd7, 0xce, 0xf7, 0x65, 0xfa, 0x9f, 0xe7, 0x62, 0x1f, 0x09, 0x6d, 0x84, 0x67, 0x61, 0xad, 0xd9,
	0x14, 0xe7, 0x53, 0x10, 0xdd, 0x5d, 0x40, 0x19, 0xa6, 0x26, 0xb7, 0x3f, 0xfc, 0xf8, 0xf3, 0x75,
	0x69, 0x9d, 0xad, 0x09, 0x27, 0xdb, 0xc2, 0x1e, 0x71, 0x26, 0xdf, 0xec, 0x33, 0xa1, 0x75, 0xf7,
	0x77, 0xd9, 0x3c, 0xeb, 0x6a, 0x42, 0xa2, 0xcd, 0x45, 0xa4, 0x88, 0x71, 0xcf, 0x63, 0x6c, 0xb0,
	0x5b, 0x17, 0x61, 0x00, 0x68, 0xf1, 0x2e, 0x44, 0xec, 0x3d, 0xfb, 0x44, 0xe8, 0xff, 0xae, 0x7d,
	0x57, 0xcf, 0x07, 0xaa, 0xa6, 0x66, 0x2e, 0xd0, 0x99, 0x00, 0x24, 0x37, 0x3d, 0xd0, 0x1a, 0xbb,
	0x3e, 0x07, 0xa8, 0xdd, 0x3e, 0x1c, 0xc7, 0xe4, 0x68, 0x1c, 0x93, 0xdf, 0xe3, 0x98, 0x7c, 0x99,
	0xc4, 0xb5, 0xa3, 0x49, 0x5c, 0xfb, 0x39, 0x89, 0x6b, 0x4f, 0x5b, 0x69, 0x66, 0x5f, 0x0e, 0x7b,
	0xbc, 0x0f, 0xf9, 0x0c, 0x83, 0xb7, 0xde, 0xc2, 0x8e, 0x4a, 0x65, 0x7a, 0x0d, 0xff, 0x39, 0xdf,
	0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xef, 0xbf, 0xf5, 0xa1, 0xdf, 0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Pool items.
	Pool(ctx context.Context, in *QueryGetPoolRequest, opts ...grpc.CallOption) (*QueryGetPoolResponse, error)
	PoolAll(ctx context.Context, in *QueryAllPoolRequest, opts ...grpc.CallOption) (*QueryAllPoolResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/elysnetwork.elys.amm.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pool(ctx context.Context, in *QueryGetPoolRequest, opts ...grpc.CallOption) (*QueryGetPoolResponse, error) {
	out := new(QueryGetPoolResponse)
	err := c.cc.Invoke(ctx, "/elysnetwork.elys.amm.Query/Pool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PoolAll(ctx context.Context, in *QueryAllPoolRequest, opts ...grpc.CallOption) (*QueryAllPoolResponse, error) {
	out := new(QueryAllPoolResponse)
	err := c.cc.Invoke(ctx, "/elysnetwork.elys.amm.Query/PoolAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Pool items.
	Pool(context.Context, *QueryGetPoolRequest) (*QueryGetPoolResponse, error)
	PoolAll(context.Context, *QueryAllPoolRequest) (*QueryAllPoolResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Pool(ctx context.Context, req *QueryGetPoolRequest) (*QueryGetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (*UnimplementedQueryServer) PoolAll(ctx context.Context, req *QueryAllPoolRequest) (*QueryAllPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PoolAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
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
		FullMethod: "/elysnetwork.elys.amm.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elysnetwork.elys.amm.Query/Pool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryGetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PoolAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PoolAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elysnetwork.elys.amm.Query/PoolAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PoolAll(ctx, req.(*QueryAllPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elysnetwork.elys.amm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "PoolAll",
			Handler:    _Query_PoolAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elys/amm/query.proto",
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

func (m *QueryGetPoolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPoolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPoolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllPoolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPoolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPoolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Pool) > 0 {
		for iNdEx := len(m.Pool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGetPoolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovQuery(uint64(m.PoolId))
	}
	return n
}

func (m *QueryGetPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pool.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllPoolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pool) > 0 {
		for _, e := range m.Pool {
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
func (m *QueryGetPoolRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetPoolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryGetPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
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
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllPoolRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPoolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
func (m *QueryAllPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
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
			m.Pool = append(m.Pool, Pool{})
			if err := m.Pool[len(m.Pool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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