// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.2
// source: invoice_service.proto

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

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceServiceClient interface {
	// 获取发票租户
	GetTenant(ctx context.Context, in *InvoiceTenantRequest, opts ...grpc.CallOption) (*SInvoiceTenant, error)
	// 新增发票抬头
	CreateInvoiceTitle(ctx context.Context, in *CreateInvoiceTitleRequest, opts ...grpc.CallOption) (*CreateInvoiceTitleResponse, error)
	// 保存发票
	RequestInvoice(ctx context.Context, in *InvoiceRequest, opts ...grpc.CallOption) (*RequestInvoiceResponse, error)
	// 获取发票
	GetInvoice(ctx context.Context, in *InvoiceId, opts ...grpc.CallOption) (*SInvoice, error)
	// Issue 开具发票,更新发票图片
	Issue(ctx context.Context, in *InvoiceIssueRequest, opts ...grpc.CallOption) (*TxResult, error)
	// Issue 发票开具失败
	IssueFail(ctx context.Context, in *InvoiceIssueFailRequest, opts ...grpc.CallOption) (*TxResult, error)
	// SendMail 发送发票到邮件中
	SendMail(ctx context.Context, in *InvoiceSendMailRequest, opts ...grpc.CallOption) (*TxResult, error)
	// Revert 撤销发票
	Revert(ctx context.Context, in *InvoiceRevertRequest, opts ...grpc.CallOption) (*TxResult, error)
}

type invoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceServiceClient(cc grpc.ClientConnInterface) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) GetTenant(ctx context.Context, in *InvoiceTenantRequest, opts ...grpc.CallOption) (*SInvoiceTenant, error) {
	out := new(SInvoiceTenant)
	err := c.cc.Invoke(ctx, "/InvoiceService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) CreateInvoiceTitle(ctx context.Context, in *CreateInvoiceTitleRequest, opts ...grpc.CallOption) (*CreateInvoiceTitleResponse, error) {
	out := new(CreateInvoiceTitleResponse)
	err := c.cc.Invoke(ctx, "/InvoiceService/CreateInvoiceTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) RequestInvoice(ctx context.Context, in *InvoiceRequest, opts ...grpc.CallOption) (*RequestInvoiceResponse, error) {
	out := new(RequestInvoiceResponse)
	err := c.cc.Invoke(ctx, "/InvoiceService/RequestInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) GetInvoice(ctx context.Context, in *InvoiceId, opts ...grpc.CallOption) (*SInvoice, error) {
	out := new(SInvoice)
	err := c.cc.Invoke(ctx, "/InvoiceService/GetInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) Issue(ctx context.Context, in *InvoiceIssueRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/InvoiceService/Issue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) IssueFail(ctx context.Context, in *InvoiceIssueFailRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/InvoiceService/IssueFail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) SendMail(ctx context.Context, in *InvoiceSendMailRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/InvoiceService/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceServiceClient) Revert(ctx context.Context, in *InvoiceRevertRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/InvoiceService/Revert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
// All implementations must embed UnimplementedInvoiceServiceServer
// for forward compatibility
type InvoiceServiceServer interface {
	// 获取发票租户
	GetTenant(context.Context, *InvoiceTenantRequest) (*SInvoiceTenant, error)
	// 新增发票抬头
	CreateInvoiceTitle(context.Context, *CreateInvoiceTitleRequest) (*CreateInvoiceTitleResponse, error)
	// 保存发票
	RequestInvoice(context.Context, *InvoiceRequest) (*RequestInvoiceResponse, error)
	// 获取发票
	GetInvoice(context.Context, *InvoiceId) (*SInvoice, error)
	// Issue 开具发票,更新发票图片
	Issue(context.Context, *InvoiceIssueRequest) (*TxResult, error)
	// Issue 发票开具失败
	IssueFail(context.Context, *InvoiceIssueFailRequest) (*TxResult, error)
	// SendMail 发送发票到邮件中
	SendMail(context.Context, *InvoiceSendMailRequest) (*TxResult, error)
	// Revert 撤销发票
	Revert(context.Context, *InvoiceRevertRequest) (*TxResult, error)
	mustEmbedUnimplementedInvoiceServiceServer()
}

// UnimplementedInvoiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInvoiceServiceServer struct {
}

func (UnimplementedInvoiceServiceServer) GetTenant(context.Context, *InvoiceTenantRequest) (*SInvoiceTenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedInvoiceServiceServer) CreateInvoiceTitle(context.Context, *CreateInvoiceTitleRequest) (*CreateInvoiceTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoiceTitle not implemented")
}
func (UnimplementedInvoiceServiceServer) RequestInvoice(context.Context, *InvoiceRequest) (*RequestInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) GetInvoice(context.Context, *InvoiceId) (*SInvoice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedInvoiceServiceServer) Issue(context.Context, *InvoiceIssueRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedInvoiceServiceServer) IssueFail(context.Context, *InvoiceIssueFailRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueFail not implemented")
}
func (UnimplementedInvoiceServiceServer) SendMail(context.Context, *InvoiceSendMailRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedInvoiceServiceServer) Revert(context.Context, *InvoiceRevertRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revert not implemented")
}
func (UnimplementedInvoiceServiceServer) mustEmbedUnimplementedInvoiceServiceServer() {}

// UnsafeInvoiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceServiceServer will
// result in compilation errors.
type UnsafeInvoiceServiceServer interface {
	mustEmbedUnimplementedInvoiceServiceServer()
}

func RegisterInvoiceServiceServer(s grpc.ServiceRegistrar, srv InvoiceServiceServer) {
	s.RegisterService(&InvoiceService_ServiceDesc, srv)
}

func _InvoiceService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetTenant(ctx, req.(*InvoiceTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_CreateInvoiceTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).CreateInvoiceTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/CreateInvoiceTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).CreateInvoiceTitle(ctx, req.(*CreateInvoiceTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_RequestInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).RequestInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/RequestInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).RequestInvoice(ctx, req.(*InvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/GetInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).GetInvoice(ctx, req.(*InvoiceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/Issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).Issue(ctx, req.(*InvoiceIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_IssueFail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceIssueFailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).IssueFail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/IssueFail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).IssueFail(ctx, req.(*InvoiceIssueFailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceSendMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).SendMail(ctx, req.(*InvoiceSendMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceService_Revert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceRevertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).Revert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InvoiceService/Revert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).Revert(ctx, req.(*InvoiceRevertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InvoiceService_ServiceDesc is the grpc.ServiceDesc for InvoiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvoiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTenant",
			Handler:    _InvoiceService_GetTenant_Handler,
		},
		{
			MethodName: "CreateInvoiceTitle",
			Handler:    _InvoiceService_CreateInvoiceTitle_Handler,
		},
		{
			MethodName: "RequestInvoice",
			Handler:    _InvoiceService_RequestInvoice_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _InvoiceService_GetInvoice_Handler,
		},
		{
			MethodName: "Issue",
			Handler:    _InvoiceService_Issue_Handler,
		},
		{
			MethodName: "IssueFail",
			Handler:    _InvoiceService_IssueFail_Handler,
		},
		{
			MethodName: "SendMail",
			Handler:    _InvoiceService_SendMail_Handler,
		},
		{
			MethodName: "Revert",
			Handler:    _InvoiceService_Revert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice_service.proto",
}
