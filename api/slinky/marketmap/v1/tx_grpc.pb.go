// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: slinky/marketmap/v1/tx.proto

package marketmapv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateMarketMap_FullMethodName = "/slinky.marketmap.v1.Msg/UpdateMarketMap"
	Msg_Params_FullMethodName          = "/slinky.marketmap.v1.Msg/Params"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateMarketMap creates markets from the given message.
	UpdateMarketMap(ctx context.Context, in *MsgUpdateMarketMap, opts ...grpc.CallOption) (*MsgUpdateMarketMapResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateMarketMap(ctx context.Context, in *MsgUpdateMarketMap, opts ...grpc.CallOption) (*MsgUpdateMarketMapResponse, error) {
	out := new(MsgUpdateMarketMapResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateMarketMap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error) {
	out := new(MsgParamsResponse)
	err := c.cc.Invoke(ctx, Msg_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateMarketMap creates markets from the given message.
	UpdateMarketMap(context.Context, *MsgUpdateMarketMap) (*MsgUpdateMarketMapResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(context.Context, *MsgParams) (*MsgParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateMarketMap(context.Context, *MsgUpdateMarketMap) (*MsgUpdateMarketMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarketMap not implemented")
}
func (UnimplementedMsgServer) Params(context.Context, *MsgParams) (*MsgParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateMarketMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateMarketMap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMarketMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateMarketMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMarketMap(ctx, req.(*MsgUpdateMarketMap))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Params(ctx, req.(*MsgParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.marketmap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMarketMap",
			Handler:    _Msg_UpdateMarketMap_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Msg_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/marketmap/v1/tx.proto",
}
