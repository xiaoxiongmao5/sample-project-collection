// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v4.24.2
// source: api.proto

package api

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// Generator1Client is the client API for Generator1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Generator1Client interface {
	GetID(ctx context.Context, in *GenReq, opts ...grpc_go.CallOption) (*GenResp, common.ErrorWithAttachment)
}

type generator1Client struct {
	cc *triple.TripleConn
}

type Generator1ClientImpl struct {
	GetID func(ctx context.Context, in *GenReq) (*GenResp, error)
}

func (c *Generator1ClientImpl) GetDubboStub(cc *triple.TripleConn) Generator1Client {
	return NewGenerator1Client(cc)
}

func (c *Generator1ClientImpl) XXX_InterfaceName() string {
	return "api.Generator1"
}

func NewGenerator1Client(cc *triple.TripleConn) Generator1Client {
	return &generator1Client{cc}
}

func (c *generator1Client) GetID(ctx context.Context, in *GenReq, opts ...grpc_go.CallOption) (*GenResp, common.ErrorWithAttachment) {
	out := new(GenResp)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetID", in, out)
}

// Generator1Server is the server API for Generator1 service.
// All implementations must embed UnimplementedGenerator1Server
// for forward compatibility
type Generator1Server interface {
	GetID(context.Context, *GenReq) (*GenResp, error)
	mustEmbedUnimplementedGenerator1Server()
}

// UnimplementedGenerator1Server must be embedded to have forward compatible implementations.
type UnimplementedGenerator1Server struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedGenerator1Server) GetID(context.Context, *GenReq) (*GenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetID not implemented")
}
func (s *UnimplementedGenerator1Server) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedGenerator1Server) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedGenerator1Server) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &Generator1_ServiceDesc
}
func (s *UnimplementedGenerator1Server) XXX_InterfaceName() string {
	return "api.Generator1"
}

func (UnimplementedGenerator1Server) mustEmbedUnimplementedGenerator1Server() {}

// UnsafeGenerator1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Generator1Server will
// result in compilation errors.
type UnsafeGenerator1Server interface {
	mustEmbedUnimplementedGenerator1Server()
}

func RegisterGenerator1Server(s grpc_go.ServiceRegistrar, srv Generator1Server) {
	s.RegisterService(&Generator1_ServiceDesc, srv)
}

func _Generator1_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetID", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// Generator1_ServiceDesc is the grpc_go.ServiceDesc for Generator1 service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generator1_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "api.Generator1",
	HandlerType: (*Generator1Server)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "GetID",
			Handler:    _Generator1_GetID_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "api.proto",
}

// XtestClient is the client API for Xtest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XtestClient interface {
	GetUser(ctx context.Context, in *XtestReq, opts ...grpc_go.CallOption) (*XtestResp, common.ErrorWithAttachment)
}

type xtestClient struct {
	cc *triple.TripleConn
}

type XtestClientImpl struct {
	GetUser func(ctx context.Context, in *XtestReq) (*XtestResp, error)
}

func (c *XtestClientImpl) GetDubboStub(cc *triple.TripleConn) XtestClient {
	return NewXtestClient(cc)
}

func (c *XtestClientImpl) XXX_InterfaceName() string {
	return "api.Xtest"
}

func NewXtestClient(cc *triple.TripleConn) XtestClient {
	return &xtestClient{cc}
}

func (c *xtestClient) GetUser(ctx context.Context, in *XtestReq, opts ...grpc_go.CallOption) (*XtestResp, common.ErrorWithAttachment) {
	out := new(XtestResp)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetUser", in, out)
}

// XtestServer is the server API for Xtest service.
// All implementations must embed UnimplementedXtestServer
// for forward compatibility
type XtestServer interface {
	GetUser(context.Context, *XtestReq) (*XtestResp, error)
	mustEmbedUnimplementedXtestServer()
}

// UnimplementedXtestServer must be embedded to have forward compatible implementations.
type UnimplementedXtestServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedXtestServer) GetUser(context.Context, *XtestReq) (*XtestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (s *UnimplementedXtestServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedXtestServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedXtestServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &Xtest_ServiceDesc
}
func (s *UnimplementedXtestServer) XXX_InterfaceName() string {
	return "api.Xtest"
}

func (UnimplementedXtestServer) mustEmbedUnimplementedXtestServer() {}

// UnsafeXtestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XtestServer will
// result in compilation errors.
type UnsafeXtestServer interface {
	mustEmbedUnimplementedXtestServer()
}

func RegisterXtestServer(s grpc_go.ServiceRegistrar, srv XtestServer) {
	s.RegisterService(&Xtest_ServiceDesc, srv)
}

func _Xtest_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(XtestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetUser", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// Xtest_ServiceDesc is the grpc_go.ServiceDesc for Xtest service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var Xtest_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "api.Xtest",
	HandlerType: (*XtestServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Xtest_GetUser_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "api.proto",
}