// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: craft/exp/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryWhiteListRequest struct {
}

func (m *QueryWhiteListRequest) Reset()         { *m = QueryWhiteListRequest{} }
func (m *QueryWhiteListRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWhiteListRequest) ProtoMessage()    {}
func (*QueryWhiteListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b68876adaf3112f, []int{0}
}
func (m *QueryWhiteListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhiteListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhiteListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhiteListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhiteListRequest.Merge(m, src)
}
func (m *QueryWhiteListRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhiteListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhiteListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhiteListRequest proto.InternalMessageInfo

type QueryWhiteListResponse struct {
	AccountRecord []AccountRecord `protobuf:"bytes,1,rep,name=accountRecord,proto3" json:"accountRecord"`
}

func (m *QueryWhiteListResponse) Reset()         { *m = QueryWhiteListResponse{} }
func (m *QueryWhiteListResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWhiteListResponse) ProtoMessage()    {}
func (*QueryWhiteListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b68876adaf3112f, []int{1}
}
func (m *QueryWhiteListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhiteListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhiteListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhiteListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhiteListResponse.Merge(m, src)
}
func (m *QueryWhiteListResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhiteListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhiteListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhiteListResponse proto.InternalMessageInfo

func (m *QueryWhiteListResponse) GetAccountRecord() []AccountRecord {
	if m != nil {
		return m.AccountRecord
	}
	return nil
}

type QueryDaoAssetRequest struct {
}

func (m *QueryDaoAssetRequest) Reset()         { *m = QueryDaoAssetRequest{} }
func (m *QueryDaoAssetRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDaoAssetRequest) ProtoMessage()    {}
func (*QueryDaoAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b68876adaf3112f, []int{2}
}
func (m *QueryDaoAssetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDaoAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDaoAssetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDaoAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDaoAssetRequest.Merge(m, src)
}
func (m *QueryDaoAssetRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDaoAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDaoAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDaoAssetRequest proto.InternalMessageInfo

type QueryDaoAssetResponse struct {
	DaoAsset *DaoAssetInfo `protobuf:"bytes,1,opt,name=dao_asset,json=daoAsset,proto3" json:"dao_asset,omitempty"`
}

func (m *QueryDaoAssetResponse) Reset()         { *m = QueryDaoAssetResponse{} }
func (m *QueryDaoAssetResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDaoAssetResponse) ProtoMessage()    {}
func (*QueryDaoAssetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b68876adaf3112f, []int{3}
}
func (m *QueryDaoAssetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDaoAssetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDaoAssetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDaoAssetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDaoAssetResponse.Merge(m, src)
}
func (m *QueryDaoAssetResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDaoAssetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDaoAssetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDaoAssetResponse proto.InternalMessageInfo

func (m *QueryDaoAssetResponse) GetDaoAsset() *DaoAssetInfo {
	if m != nil {
		return m.DaoAsset
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryWhiteListRequest)(nil), "craft.exp.v1beta1.QueryWhiteListRequest")
	proto.RegisterType((*QueryWhiteListResponse)(nil), "craft.exp.v1beta1.QueryWhiteListResponse")
	proto.RegisterType((*QueryDaoAssetRequest)(nil), "craft.exp.v1beta1.QueryDaoAssetRequest")
	proto.RegisterType((*QueryDaoAssetResponse)(nil), "craft.exp.v1beta1.QueryDaoAssetResponse")
}

func init() { proto.RegisterFile("craft/exp/v1beta1/query.proto", fileDescriptor_6b68876adaf3112f) }

var fileDescriptor_6b68876adaf3112f = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0xcf, 0xd2, 0x40,
	0x18, 0xef, 0xe1, 0x9f, 0xc0, 0x11, 0x07, 0x2f, 0x88, 0xa4, 0x4a, 0x8b, 0x35, 0x86, 0xea, 0xd0,
	0x0b, 0xb8, 0xba, 0x40, 0x5c, 0x4c, 0x58, 0x6c, 0x62, 0x4c, 0x5c, 0xcc, 0xd1, 0x5e, 0x6b, 0x13,
	0xbc, 0xa7, 0xf4, 0x0e, 0x85, 0x95, 0xcd, 0xcd, 0xc4, 0xc5, 0x8f, 0xc4, 0x48, 0xe2, 0xe2, 0x64,
	0x0c, 0xf8, 0x41, 0x4c, 0xaf, 0x05, 0x7d, 0xa1, 0x6f, 0xde, 0x77, 0xbb, 0x7b, 0x7e, 0x7f, 0xef,
	0xc9, 0xe1, 0x6e, 0x90, 0xb1, 0x48, 0x51, 0xbe, 0x4c, 0xe9, 0xa7, 0xc1, 0x94, 0x2b, 0x36, 0xa0,
	0xf3, 0x05, 0xcf, 0x56, 0x5e, 0x9a, 0x81, 0x02, 0x72, 0x57, 0xc3, 0x1e, 0x5f, 0xa6, 0x5e, 0x09,
	0x9b, 0xf6, 0xb9, 0x22, 0x64, 0x90, 0x88, 0x08, 0x0a, 0x8d, 0xf9, 0x30, 0x06, 0x88, 0x67, 0x9c,
	0xb2, 0x34, 0xa1, 0x4c, 0x08, 0x50, 0x4c, 0x25, 0x20, 0x64, 0x89, 0xb6, 0x62, 0x88, 0x41, 0x1f,
	0x69, 0x7e, 0x2a, 0xa7, 0xcf, 0x02, 0x90, 0x1f, 0x41, 0xd2, 0x29, 0x93, 0xbc, 0x28, 0x70, 0x34,
	0x4f, 0x59, 0x9c, 0x08, 0x6d, 0x51, 0x70, 0x9d, 0xfb, 0xf8, 0xde, 0xeb, 0x9c, 0xf1, 0xf6, 0x43,
	0xa2, 0xf8, 0x24, 0x91, 0xca, 0xe7, 0xf3, 0x05, 0x97, 0xca, 0x89, 0x70, 0xfb, 0x14, 0x90, 0x29,
	0x08, 0xc9, 0xc9, 0x04, 0xdf, 0x61, 0x41, 0x00, 0x0b, 0xa1, 0x7c, 0x1e, 0x40, 0x16, 0x76, 0x50,
	0xef, 0x86, 0xdb, 0x1c, 0xf6, 0xbc, 0xb3, 0xe7, 0x79, 0xa3, 0xff, 0x79, 0xe3, 0x9b, 0x9b, 0x5f,
	0xb6, 0xe1, 0x5f, 0x14, 0x3b, 0x6d, 0xdc, 0xd2, 0x39, 0x2f, 0x19, 0x8c, 0xa4, 0xe4, 0xc7, 0xfc,
	0x37, 0x65, 0xb1, 0x7f, 0xf3, 0x32, 0xfe, 0x05, 0x6e, 0x84, 0x0c, 0xde, 0xb3, 0x7c, 0xd8, 0x41,
	0x3d, 0xe4, 0x36, 0x87, 0x76, 0x45, 0xf4, 0x41, 0xf7, 0x4a, 0x44, 0xe0, 0xd7, 0xc3, 0xf2, 0x36,
	0xfc, 0x5e, 0xc3, 0xb7, 0xb4, 0x2f, 0xf9, 0x82, 0x70, 0xe3, 0xf8, 0x38, 0xe2, 0x56, 0x58, 0x54,
	0x2e, 0xc6, 0x7c, 0x7a, 0x0d, 0x66, 0x51, 0xd5, 0xe9, 0xaf, 0x7f, 0xfc, 0xf9, 0x56, 0x7b, 0x44,
	0x6c, 0x7a, 0x26, 0xa1, 0x9f, 0x73, 0xf6, 0x2c, 0x91, 0x2a, 0x64, 0x40, 0xd6, 0x08, 0xd7, 0x0f,
	0x85, 0x49, 0xff, 0xb2, 0x80, 0x93, 0x15, 0x99, 0xee, 0xd5, 0xc4, 0xb2, 0xc8, 0x63, 0x5d, 0xa4,
	0x4b, 0x1e, 0x54, 0x14, 0x09, 0x19, 0xe8, 0x5d, 0x8e, 0x9f, 0x6c, 0x76, 0x16, 0xda, 0xee, 0x2c,
	0xf4, 0x7b, 0x67, 0xa1, 0xaf, 0x7b, 0xcb, 0xd8, 0xee, 0x2d, 0xe3, 0xe7, 0xde, 0x32, 0xde, 0x35,
	0x97, 0xfa, 0x87, 0xaa, 0x55, 0xca, 0xe5, 0xf4, 0xb6, 0xfe, 0x38, 0xcf, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x98, 0xf9, 0x07, 0x5e, 0xed, 0x02, 0x00, 0x00,
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
	// returns coins that is going to be distributed
	WhiteList(ctx context.Context, in *QueryWhiteListRequest, opts ...grpc.CallOption) (*QueryWhiteListResponse, error)
	DaoAsset(ctx context.Context, in *QueryDaoAssetRequest, opts ...grpc.CallOption) (*QueryDaoAssetResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) WhiteList(ctx context.Context, in *QueryWhiteListRequest, opts ...grpc.CallOption) (*QueryWhiteListResponse, error) {
	out := new(QueryWhiteListResponse)
	err := c.cc.Invoke(ctx, "/craft.exp.v1beta1.Query/WhiteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DaoAsset(ctx context.Context, in *QueryDaoAssetRequest, opts ...grpc.CallOption) (*QueryDaoAssetResponse, error) {
	out := new(QueryDaoAssetResponse)
	err := c.cc.Invoke(ctx, "/craft.exp.v1beta1.Query/DaoAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// returns coins that is going to be distributed
	WhiteList(context.Context, *QueryWhiteListRequest) (*QueryWhiteListResponse, error)
	DaoAsset(context.Context, *QueryDaoAssetRequest) (*QueryDaoAssetResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) WhiteList(ctx context.Context, req *QueryWhiteListRequest) (*QueryWhiteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhiteList not implemented")
}
func (*UnimplementedQueryServer) DaoAsset(ctx context.Context, req *QueryDaoAssetRequest) (*QueryDaoAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DaoAsset not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_WhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWhiteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/craft.exp.v1beta1.Query/WhiteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WhiteList(ctx, req.(*QueryWhiteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DaoAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDaoAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DaoAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/craft.exp.v1beta1.Query/DaoAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DaoAsset(ctx, req.(*QueryDaoAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "craft.exp.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhiteList",
			Handler:    _Query_WhiteList_Handler,
		},
		{
			MethodName: "DaoAsset",
			Handler:    _Query_DaoAsset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "craft/exp/v1beta1/query.proto",
}

func (m *QueryWhiteListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhiteListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhiteListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryWhiteListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhiteListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhiteListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountRecord) > 0 {
		for iNdEx := len(m.AccountRecord) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountRecord[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryDaoAssetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDaoAssetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDaoAssetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryDaoAssetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDaoAssetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDaoAssetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DaoAsset != nil {
		{
			size, err := m.DaoAsset.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryWhiteListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryWhiteListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AccountRecord) > 0 {
		for _, e := range m.AccountRecord {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryDaoAssetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryDaoAssetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DaoAsset != nil {
		l = m.DaoAsset.Size()
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
func (m *QueryWhiteListRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWhiteListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhiteListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryWhiteListResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWhiteListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhiteListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountRecord", wireType)
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
			m.AccountRecord = append(m.AccountRecord, AccountRecord{})
			if err := m.AccountRecord[len(m.AccountRecord)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDaoAssetRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDaoAssetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDaoAssetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryDaoAssetResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDaoAssetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDaoAssetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoAsset", wireType)
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
			if m.DaoAsset == nil {
				m.DaoAsset = &DaoAssetInfo{}
			}
			if err := m.DaoAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
