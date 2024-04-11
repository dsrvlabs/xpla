// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xpla/reward/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgFundFeeCollector allows an account to directly
// fund the fee collector.
type MsgFundFeeCollector struct {
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Depositor string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
}

func (m *MsgFundFeeCollector) Reset()         { *m = MsgFundFeeCollector{} }
func (m *MsgFundFeeCollector) String() string { return proto.CompactTextString(m) }
func (*MsgFundFeeCollector) ProtoMessage()    {}
func (*MsgFundFeeCollector) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab6178765eaaaf2, []int{0}
}
func (m *MsgFundFeeCollector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundFeeCollector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundFeeCollector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundFeeCollector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundFeeCollector.Merge(m, src)
}
func (m *MsgFundFeeCollector) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundFeeCollector) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundFeeCollector.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundFeeCollector proto.InternalMessageInfo

// MsgFundFeeCollectorResponse defines the Msg/FundFeeCollector response type.
type MsgFundFeeCollectorResponse struct {
}

func (m *MsgFundFeeCollectorResponse) Reset()         { *m = MsgFundFeeCollectorResponse{} }
func (m *MsgFundFeeCollectorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFundFeeCollectorResponse) ProtoMessage()    {}
func (*MsgFundFeeCollectorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab6178765eaaaf2, []int{1}
}
func (m *MsgFundFeeCollectorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundFeeCollectorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundFeeCollectorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundFeeCollectorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundFeeCollectorResponse.Merge(m, src)
}
func (m *MsgFundFeeCollectorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundFeeCollectorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundFeeCollectorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundFeeCollectorResponse proto.InternalMessageInfo

// MsgUpdateParams is the Msg/UpdateParams request type for reward parameters.
// Since: cosmos-sdk 0.47
type MsgUpdateParams struct {
	// authority is the address of the governance account.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/evm parameters to update.
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab6178765eaaaf2, []int{2}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
// Since: cosmos-sdk 0.47
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab6178765eaaaf2, []int{3}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgFundFeeCollector)(nil), "xpla.reward.v1beta1.MsgFundFeeCollector")
	proto.RegisterType((*MsgFundFeeCollectorResponse)(nil), "xpla.reward.v1beta1.MsgFundFeeCollectorResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "xpla.reward.v1beta1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "xpla.reward.v1beta1.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("xpla/reward/v1beta1/tx.proto", fileDescriptor_9ab6178765eaaaf2) }

var fileDescriptor_9ab6178765eaaaf2 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0x51, 0x14, 0x29, 0x57, 0xc4, 0x0f, 0xb7, 0x52, 0xf3, 0xa3, 0x38, 0x51, 0x84, 0x90,
	0x55, 0x51, 0x5f, 0x13, 0x24, 0x24, 0x3a, 0x41, 0x2a, 0x75, 0x8b, 0x84, 0x8c, 0x58, 0x58, 0xd0,
	0xd9, 0x3e, 0xb9, 0x16, 0xb1, 0x9f, 0x75, 0xef, 0x12, 0xd2, 0x95, 0x09, 0x31, 0x31, 0xf0, 0x07,
	0x74, 0x66, 0xea, 0xc0, 0x5f, 0xc0, 0xd4, 0xb1, 0x62, 0x82, 0x05, 0x50, 0x32, 0x94, 0x3f, 0x03,
	0xd9, 0x3e, 0x93, 0x02, 0x46, 0xca, 0x92, 0x8b, 0xdf, 0xf7, 0xdd, 0xf7, 0x7d, 0xef, 0xde, 0xa3,
	0xdb, 0xb3, 0x74, 0xcc, 0x99, 0x14, 0xaf, 0xb8, 0x0c, 0xd8, 0xb4, 0xef, 0x09, 0xc5, 0xfb, 0x4c,
	0xcd, 0x9c, 0x54, 0x82, 0x02, 0x73, 0x23, 0x43, 0x9d, 0x02, 0x75, 0x34, 0xda, 0xda, 0x0c, 0x21,
	0x84, 0x1c, 0x67, 0xd9, 0xbf, 0x82, 0xda, 0xb2, 0x7c, 0xc0, 0x18, 0x90, 0x79, 0x1c, 0xc5, 0x6f,
	0x21, 0x1f, 0xa2, 0x44, 0xe3, 0x5b, 0x1a, 0x8f, 0x31, 0x64, 0xd3, 0x7e, 0x76, 0x68, 0xa0, 0x59,
	0x00, 0x2f, 0x0a, 0xc5, 0xe2, 0x43, 0x43, 0xdd, 0xaa, 0x70, 0x3a, 0x4d, 0xc1, 0xb8, 0xc5, 0xe3,
	0x28, 0x01, 0x96, 0xff, 0x16, 0xa5, 0xde, 0x27, 0x42, 0x37, 0x46, 0x18, 0x1e, 0x4e, 0x92, 0xe0,
	0x50, 0x88, 0x03, 0x18, 0x8f, 0x85, 0xaf, 0x40, 0x9a, 0x3e, 0xad, 0xf1, 0x18, 0x26, 0x89, 0x6a,
	0x90, 0xee, 0x9a, 0xbd, 0x3e, 0x68, 0x3a, 0xda, 0x2b, 0x4b, 0x5c, 0x36, 0xe7, 0x1c, 0x40, 0x94,
	0x0c, 0xf7, 0xce, 0xbe, 0x75, 0x8c, 0x0f, 0xdf, 0x3b, 0x76, 0x18, 0xa9, 0xa3, 0x89, 0xe7, 0xf8,
	0x10, 0xeb, 0x60, 0xfa, 0xd8, 0xc5, 0xe0, 0x25, 0x53, 0xc7, 0xa9, 0xc0, 0xfc, 0x02, 0xba, 0x5a,
	0xda, 0xdc, 0xa6, 0xf5, 0x40, 0xa4, 0x80, 0x91, 0x02, 0xd9, 0xb8, 0xd2, 0x25, 0x76, 0xdd, 0x5d,
	0x16, 0xf6, 0xed, 0x37, 0x27, 0x1d, 0xe3, 0xe7, 0x49, 0xc7, 0x78, 0x7b, 0x71, 0xba, 0xd3, 0xce,
	0x9a, 0x0b, 0xc4, 0x94, 0x55, 0x84, 0xed, 0xdd, 0xa6, 0xed, 0x8a, 0xb2, 0x2b, 0x30, 0x85, 0x04,
	0x45, 0xef, 0x3d, 0xa1, 0x37, 0x46, 0x18, 0x3e, 0x4b, 0x03, 0xae, 0xc4, 0x13, 0x2e, 0x79, 0x8c,
	0xe6, 0x03, 0x5a, 0xe7, 0x13, 0x75, 0x04, 0x32, 0x52, 0xc7, 0x0d, 0x92, 0x59, 0x0f, 0x1b, 0x9f,
	0x3f, 0xee, 0x6e, 0xea, 0x2e, 0x1f, 0x07, 0x81, 0x14, 0x88, 0x4f, 0x95, 0x8c, 0x92, 0xd0, 0x5d,
	0x52, 0xcd, 0x87, 0xb4, 0x96, 0xe6, 0x0a, 0x79, 0xde, 0xf5, 0x41, 0xdb, 0xa9, 0x18, 0xba, 0x53,
	0x98, 0x0c, 0xaf, 0x66, 0x2f, 0xe3, 0xea, 0x0b, 0xfb, 0xd7, 0x5f, 0x5f, 0x9c, 0xee, 0x2c, 0xa5,
	0x7a, 0x4d, 0xba, 0xf5, 0x57, 0xaa, 0x32, 0xf1, 0xe0, 0x2b, 0xa1, 0x6b, 0x23, 0x0c, 0xcd, 0x84,
	0xde, 0xfc, 0x67, 0x32, 0x76, 0xa5, 0x63, 0x45, 0xff, 0xad, 0xbd, 0x55, 0x99, 0xa5, 0xaf, 0xe9,
	0xd1, 0x6b, 0x7f, 0xbc, 0xd2, 0x9d, 0xff, 0x29, 0x5c, 0x66, 0xb5, 0xee, 0xad, 0xc2, 0x2a, 0x3d,
	0x86, 0x8f, 0xce, 0xe6, 0x16, 0x39, 0x9f, 0x5b, 0xe4, 0xc7, 0xdc, 0x22, 0xef, 0x16, 0x96, 0x71,
	0xbe, 0xb0, 0x8c, 0x2f, 0x0b, 0xcb, 0x78, 0x7e, 0xf7, 0xd2, 0x02, 0x95, 0xe3, 0xce, 0x77, 0x7a,
	0x56, 0x6e, 0x75, 0xbe, 0x44, 0x5e, 0x2d, 0x5f, 0xdd, 0xfb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x69, 0x2f, 0x59, 0x50, 0x8e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// FundFeeCollector defines a method to allow an account to directly
	// fund the fee collector.
	FundFeeCollector(ctx context.Context, in *MsgFundFeeCollector, opts ...grpc.CallOption) (*MsgFundFeeCollectorResponse, error)
	// UpdateParams defined a governance operation for updating the x/reward
	// module parameters. The authority is hard-coded to the Cosmos SDK x/gov
	// module account
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) FundFeeCollector(ctx context.Context, in *MsgFundFeeCollector, opts ...grpc.CallOption) (*MsgFundFeeCollectorResponse, error) {
	out := new(MsgFundFeeCollectorResponse)
	err := c.cc.Invoke(ctx, "/xpla.reward.v1beta1.Msg/FundFeeCollector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/xpla.reward.v1beta1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// FundFeeCollector defines a method to allow an account to directly
	// fund the fee collector.
	FundFeeCollector(context.Context, *MsgFundFeeCollector) (*MsgFundFeeCollectorResponse, error)
	// UpdateParams defined a governance operation for updating the x/reward
	// module parameters. The authority is hard-coded to the Cosmos SDK x/gov
	// module account
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) FundFeeCollector(ctx context.Context, req *MsgFundFeeCollector) (*MsgFundFeeCollectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundFeeCollector not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_FundFeeCollector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundFeeCollector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundFeeCollector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xpla.reward.v1beta1.Msg/FundFeeCollector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundFeeCollector(ctx, req.(*MsgFundFeeCollector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xpla.reward.v1beta1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xpla.reward.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FundFeeCollector",
			Handler:    _Msg_FundFeeCollector_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xpla/reward/v1beta1/tx.proto",
}

func (m *MsgFundFeeCollector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundFeeCollector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundFeeCollector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgFundFeeCollectorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundFeeCollectorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundFeeCollectorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgFundFeeCollector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFundFeeCollectorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgFundFeeCollector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFundFeeCollector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundFeeCollector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgFundFeeCollectorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFundFeeCollectorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundFeeCollectorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
