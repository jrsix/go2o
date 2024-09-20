// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: approval_service.proto

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

// ApprovalServiceClient is the client API for ApprovalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApprovalServiceClient interface {
	// 获取审批流
	GetFlow(ctx context.Context, in *ApprovalFlowRequest, opts ...grpc.CallOption) (*SApprovalFlow, error)
	// 审批通过
	Approve(ctx context.Context, in *ApprovalApproveRequest, opts ...grpc.CallOption) (*TxResult, error)
	// 审批拒绝
	Reject(ctx context.Context, in *ApprovalRejectRequest, opts ...grpc.CallOption) (*TxResult, error)
	// 分配审批人,当节点审批后切换到下个节点, 需分配审批人
	Assign(ctx context.Context, in *ApprovalAssignRequest, opts ...grpc.CallOption) (*TxResult, error)
}

type approvalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApprovalServiceClient(cc grpc.ClientConnInterface) ApprovalServiceClient {
	return &approvalServiceClient{cc}
}

func (c *approvalServiceClient) GetFlow(ctx context.Context, in *ApprovalFlowRequest, opts ...grpc.CallOption) (*SApprovalFlow, error) {
	out := new(SApprovalFlow)
	err := c.cc.Invoke(ctx, "/ApprovalService/GetFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) Approve(ctx context.Context, in *ApprovalApproveRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/ApprovalService/Approve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) Reject(ctx context.Context, in *ApprovalRejectRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/ApprovalService/Reject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *approvalServiceClient) Assign(ctx context.Context, in *ApprovalAssignRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/ApprovalService/Assign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApprovalServiceServer is the server API for ApprovalService service.
// All implementations must embed UnimplementedApprovalServiceServer
// for forward compatibility
type ApprovalServiceServer interface {
	// 获取审批流
	GetFlow(context.Context, *ApprovalFlowRequest) (*SApprovalFlow, error)
	// 审批通过
	Approve(context.Context, *ApprovalApproveRequest) (*TxResult, error)
	// 审批拒绝
	Reject(context.Context, *ApprovalRejectRequest) (*TxResult, error)
	// 分配审批人,当节点审批后切换到下个节点, 需分配审批人
	Assign(context.Context, *ApprovalAssignRequest) (*TxResult, error)
	mustEmbedUnimplementedApprovalServiceServer()
}

// UnimplementedApprovalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApprovalServiceServer struct {
}

func (UnimplementedApprovalServiceServer) GetFlow(context.Context, *ApprovalFlowRequest) (*SApprovalFlow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlow not implemented")
}
func (UnimplementedApprovalServiceServer) Approve(context.Context, *ApprovalApproveRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (UnimplementedApprovalServiceServer) Reject(context.Context, *ApprovalRejectRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reject not implemented")
}
func (UnimplementedApprovalServiceServer) Assign(context.Context, *ApprovalAssignRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assign not implemented")
}
func (UnimplementedApprovalServiceServer) mustEmbedUnimplementedApprovalServiceServer() {}

// UnsafeApprovalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApprovalServiceServer will
// result in compilation errors.
type UnsafeApprovalServiceServer interface {
	mustEmbedUnimplementedApprovalServiceServer()
}

func RegisterApprovalServiceServer(s grpc.ServiceRegistrar, srv ApprovalServiceServer) {
	s.RegisterService(&ApprovalService_ServiceDesc, srv)
}

func _ApprovalService_GetFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApprovalFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).GetFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalService/GetFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).GetFlow(ctx, req.(*ApprovalFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApprovalApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalService/Approve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).Approve(ctx, req.(*ApprovalApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApprovalRejectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalService/Reject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).Reject(ctx, req.(*ApprovalRejectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprovalService_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApprovalAssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprovalServiceServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApprovalService/Assign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprovalServiceServer).Assign(ctx, req.(*ApprovalAssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApprovalService_ServiceDesc is the grpc.ServiceDesc for ApprovalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApprovalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ApprovalService",
	HandlerType: (*ApprovalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlow",
			Handler:    _ApprovalService_GetFlow_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _ApprovalService_Approve_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _ApprovalService_Reject_Handler,
		},
		{
			MethodName: "Assign",
			Handler:    _ApprovalService_Assign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "approval_service.proto",
}
