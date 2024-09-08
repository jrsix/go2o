// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: express_service.proto

package proto

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

// ExpressServiceClient is the client API for ExpressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpressServiceClient interface {
	// 获取快递公司
	GetExpressProvider(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SExpressProvider, error)
	// 保存快递公司
	SaveExpressProvider(ctx context.Context, in *SExpressProvider, opts ...grpc.CallOption) (*Result, error)
	// 获取可用的快递公司
	GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderListResponse, error)
	// 获取可用的快递公司分组
	GetProviderGroup(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderGroupResponse, error)
	// 保存快递模板
	SaveExpressTemplate(ctx context.Context, in *SExpressTemplate, opts ...grpc.CallOption) (*SaveTemplateResponse, error)
	// 获取单个快递模板
	GetTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*SExpressTemplate, error)
	// 获取卖家的快递模板
	GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*ExpressTemplateListResponse, error)
	// 删除模板
	DeleteTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*Result, error)
	// 保存地区快递模板
	SaveAreaTemplate(ctx context.Context, in *SaveAreaExpTemplateRequest, opts ...grpc.CallOption) (*Result, error)
}

type expressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpressServiceClient(cc grpc.ClientConnInterface) ExpressServiceClient {
	return &expressServiceClient{cc}
}

func (c *expressServiceClient) GetExpressProvider(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SExpressProvider, error) {
	out := new(SExpressProvider)
	err := c.cc.Invoke(ctx, "/ExpressService/GetExpressProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveExpressProvider(ctx context.Context, in *SExpressProvider, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveExpressProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderListResponse, error) {
	out := new(ExpressProviderListResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/GetProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetProviderGroup(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderGroupResponse, error) {
	out := new(ExpressProviderGroupResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/GetProviderGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveExpressTemplate(ctx context.Context, in *SExpressTemplate, opts ...grpc.CallOption) (*SaveTemplateResponse, error) {
	out := new(SaveTemplateResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveExpressTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*SExpressTemplate, error) {
	out := new(SExpressTemplate)
	err := c.cc.Invoke(ctx, "/ExpressService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*ExpressTemplateListResponse, error) {
	out := new(ExpressTemplateListResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/GetTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) DeleteTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveAreaTemplate(ctx context.Context, in *SaveAreaExpTemplateRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveAreaTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpressServiceServer is the server API for ExpressService service.
// All implementations must embed UnimplementedExpressServiceServer
// for forward compatibility
type ExpressServiceServer interface {
	// 获取快递公司
	GetExpressProvider(context.Context, *IdOrName) (*SExpressProvider, error)
	// 保存快递公司
	SaveExpressProvider(context.Context, *SExpressProvider) (*Result, error)
	// 获取可用的快递公司
	GetProviders(context.Context, *Empty) (*ExpressProviderListResponse, error)
	// 获取可用的快递公司分组
	GetProviderGroup(context.Context, *Empty) (*ExpressProviderGroupResponse, error)
	// 保存快递模板
	SaveExpressTemplate(context.Context, *SExpressTemplate) (*SaveTemplateResponse, error)
	// 获取单个快递模板
	GetTemplate(context.Context, *ExpressTemplateId) (*SExpressTemplate, error)
	// 获取卖家的快递模板
	GetTemplates(context.Context, *GetTemplatesRequest) (*ExpressTemplateListResponse, error)
	// 删除模板
	DeleteTemplate(context.Context, *ExpressTemplateId) (*Result, error)
	// 保存地区快递模板
	SaveAreaTemplate(context.Context, *SaveAreaExpTemplateRequest) (*Result, error)
	mustEmbedUnimplementedExpressServiceServer()
}

// UnimplementedExpressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExpressServiceServer struct {
}

func (UnimplementedExpressServiceServer) GetExpressProvider(context.Context, *IdOrName) (*SExpressProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpressProvider not implemented")
}
func (UnimplementedExpressServiceServer) SaveExpressProvider(context.Context, *SExpressProvider) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveExpressProvider not implemented")
}
func (UnimplementedExpressServiceServer) GetProviders(context.Context, *Empty) (*ExpressProviderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProviders not implemented")
}
func (UnimplementedExpressServiceServer) GetProviderGroup(context.Context, *Empty) (*ExpressProviderGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProviderGroup not implemented")
}
func (UnimplementedExpressServiceServer) SaveExpressTemplate(context.Context, *SExpressTemplate) (*SaveTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveExpressTemplate not implemented")
}
func (UnimplementedExpressServiceServer) GetTemplate(context.Context, *ExpressTemplateId) (*SExpressTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedExpressServiceServer) GetTemplates(context.Context, *GetTemplatesRequest) (*ExpressTemplateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplates not implemented")
}
func (UnimplementedExpressServiceServer) DeleteTemplate(context.Context, *ExpressTemplateId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}
func (UnimplementedExpressServiceServer) SaveAreaTemplate(context.Context, *SaveAreaExpTemplateRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAreaTemplate not implemented")
}
func (UnimplementedExpressServiceServer) mustEmbedUnimplementedExpressServiceServer() {}

// UnsafeExpressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpressServiceServer will
// result in compilation errors.
type UnsafeExpressServiceServer interface {
	mustEmbedUnimplementedExpressServiceServer()
}

func RegisterExpressServiceServer(s grpc.ServiceRegistrar, srv ExpressServiceServer) {
	s.RegisterService(&ExpressService_ServiceDesc, srv)
}

func _ExpressService_GetExpressProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetExpressProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetExpressProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetExpressProvider(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveExpressProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SExpressProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveExpressProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveExpressProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveExpressProvider(ctx, req.(*SExpressProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetProviders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetProviderGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetProviderGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetProviderGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetProviderGroup(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveExpressTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SExpressTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveExpressTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveExpressTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveExpressTemplate(ctx, req.(*SExpressTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetTemplate(ctx, req.(*ExpressTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetTemplates(ctx, req.(*GetTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).DeleteTemplate(ctx, req.(*ExpressTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveAreaTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAreaExpTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveAreaTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveAreaTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveAreaTemplate(ctx, req.(*SaveAreaExpTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpressService_ServiceDesc is the grpc.ServiceDesc for ExpressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ExpressService",
	HandlerType: (*ExpressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExpressProvider",
			Handler:    _ExpressService_GetExpressProvider_Handler,
		},
		{
			MethodName: "SaveExpressProvider",
			Handler:    _ExpressService_SaveExpressProvider_Handler,
		},
		{
			MethodName: "GetProviders",
			Handler:    _ExpressService_GetProviders_Handler,
		},
		{
			MethodName: "GetProviderGroup",
			Handler:    _ExpressService_GetProviderGroup_Handler,
		},
		{
			MethodName: "SaveExpressTemplate",
			Handler:    _ExpressService_SaveExpressTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _ExpressService_GetTemplate_Handler,
		},
		{
			MethodName: "GetTemplates",
			Handler:    _ExpressService_GetTemplates_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _ExpressService_DeleteTemplate_Handler,
		},
		{
			MethodName: "SaveAreaTemplate",
			Handler:    _ExpressService_SaveAreaTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "express_service.proto",
}
