// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.2
// source: message_service.proto

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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	// 获取通知项,key
	GetNotifyItem(ctx context.Context, in *String, opts ...grpc.CallOption) (*SNotifyItem, error)
	// 发送短信
	SendPhoneMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*TxResult, error)
	// 获取所有通知项
	GetAllNotifyItem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NotifyItemListResponse, error)
	// 保存通知项设置
	SaveNotifyItem(ctx context.Context, in *SNotifyItem, opts ...grpc.CallOption) (*Result, error)
	// 获取邮件模版
	GetMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SMailTemplate, error)
	// 保存邮件模板
	SaveMailTemplate(ctx context.Context, in *SMailTemplate, opts ...grpc.CallOption) (*Result, error)
	// 获取邮件模板
	GetMailTemplates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MailTemplateListResponse, error)
	// 删除邮件模板
	DeleteMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取邮件绑定
	// rpc GetConfig() mss.Config
	// 保存邮件
	// rpc SaveConfig(conf *mss.Config) error
	// 发送站内信
	SendSiteMessage(ctx context.Context, in *SendSiteMessageRequest, opts ...grpc.CallOption) (*Result, error)
	// 创建聊天会话
	// rpc CreateChatSession(senderRole int, senderId int32, toRole int, toId int32) (mss.Message, error)
	// 保存系统通知模板
	SaveNotifyTemplate(ctx context.Context, in *SaveNotifyTemplateRequest, opts ...grpc.CallOption) (*TxResult, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) GetNotifyItem(ctx context.Context, in *String, opts ...grpc.CallOption) (*SNotifyItem, error) {
	out := new(SNotifyItem)
	err := c.cc.Invoke(ctx, "/MessageService/GetNotifyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendPhoneMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/MessageService/SendPhoneMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetAllNotifyItem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NotifyItemListResponse, error) {
	out := new(NotifyItemListResponse)
	err := c.cc.Invoke(ctx, "/MessageService/GetAllNotifyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SaveNotifyItem(ctx context.Context, in *SNotifyItem, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SaveNotifyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SMailTemplate, error) {
	out := new(SMailTemplate)
	err := c.cc.Invoke(ctx, "/MessageService/GetMailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SaveMailTemplate(ctx context.Context, in *SMailTemplate, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SaveMailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMailTemplates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MailTemplateListResponse, error) {
	out := new(MailTemplateListResponse)
	err := c.cc.Invoke(ctx, "/MessageService/GetMailTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) DeleteMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/DeleteMailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendSiteMessage(ctx context.Context, in *SendSiteMessageRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SendSiteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SaveNotifyTemplate(ctx context.Context, in *SaveNotifyTemplateRequest, opts ...grpc.CallOption) (*TxResult, error) {
	out := new(TxResult)
	err := c.cc.Invoke(ctx, "/MessageService/SaveNotifyTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	// 获取通知项,key
	GetNotifyItem(context.Context, *String) (*SNotifyItem, error)
	// 发送短信
	SendPhoneMessage(context.Context, *SendMessageRequest) (*TxResult, error)
	// 获取所有通知项
	GetAllNotifyItem(context.Context, *Empty) (*NotifyItemListResponse, error)
	// 保存通知项设置
	SaveNotifyItem(context.Context, *SNotifyItem) (*Result, error)
	// 获取邮件模版
	GetMailTemplate(context.Context, *Int64) (*SMailTemplate, error)
	// 保存邮件模板
	SaveMailTemplate(context.Context, *SMailTemplate) (*Result, error)
	// 获取邮件模板
	GetMailTemplates(context.Context, *Empty) (*MailTemplateListResponse, error)
	// 删除邮件模板
	DeleteMailTemplate(context.Context, *Int64) (*Result, error)
	// 获取邮件绑定
	// rpc GetConfig() mss.Config
	// 保存邮件
	// rpc SaveConfig(conf *mss.Config) error
	// 发送站内信
	SendSiteMessage(context.Context, *SendSiteMessageRequest) (*Result, error)
	// 创建聊天会话
	// rpc CreateChatSession(senderRole int, senderId int32, toRole int, toId int32) (mss.Message, error)
	// 保存系统通知模板
	SaveNotifyTemplate(context.Context, *SaveNotifyTemplateRequest) (*TxResult, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) GetNotifyItem(context.Context, *String) (*SNotifyItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifyItem not implemented")
}
func (UnimplementedMessageServiceServer) SendPhoneMessage(context.Context, *SendMessageRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneMessage not implemented")
}
func (UnimplementedMessageServiceServer) GetAllNotifyItem(context.Context, *Empty) (*NotifyItemListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotifyItem not implemented")
}
func (UnimplementedMessageServiceServer) SaveNotifyItem(context.Context, *SNotifyItem) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveNotifyItem not implemented")
}
func (UnimplementedMessageServiceServer) GetMailTemplate(context.Context, *Int64) (*SMailTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMailTemplate not implemented")
}
func (UnimplementedMessageServiceServer) SaveMailTemplate(context.Context, *SMailTemplate) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMailTemplate not implemented")
}
func (UnimplementedMessageServiceServer) GetMailTemplates(context.Context, *Empty) (*MailTemplateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMailTemplates not implemented")
}
func (UnimplementedMessageServiceServer) DeleteMailTemplate(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMailTemplate not implemented")
}
func (UnimplementedMessageServiceServer) SendSiteMessage(context.Context, *SendSiteMessageRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSiteMessage not implemented")
}
func (UnimplementedMessageServiceServer) SaveNotifyTemplate(context.Context, *SaveNotifyTemplateRequest) (*TxResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveNotifyTemplate not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_GetNotifyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetNotifyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetNotifyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetNotifyItem(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendPhoneMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendPhoneMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SendPhoneMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendPhoneMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetAllNotifyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetAllNotifyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetAllNotifyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetAllNotifyItem(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SaveNotifyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SNotifyItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SaveNotifyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SaveNotifyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SaveNotifyItem(ctx, req.(*SNotifyItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetMailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMailTemplate(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SaveMailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMailTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SaveMailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SaveMailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SaveMailTemplate(ctx, req.(*SMailTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetMailTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMailTemplates(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_DeleteMailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DeleteMailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/DeleteMailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DeleteMailTemplate(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendSiteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSiteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendSiteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SendSiteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendSiteMessage(ctx, req.(*SendSiteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SaveNotifyTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveNotifyTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SaveNotifyTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SaveNotifyTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SaveNotifyTemplate(ctx, req.(*SaveNotifyTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotifyItem",
			Handler:    _MessageService_GetNotifyItem_Handler,
		},
		{
			MethodName: "SendPhoneMessage",
			Handler:    _MessageService_SendPhoneMessage_Handler,
		},
		{
			MethodName: "GetAllNotifyItem",
			Handler:    _MessageService_GetAllNotifyItem_Handler,
		},
		{
			MethodName: "SaveNotifyItem",
			Handler:    _MessageService_SaveNotifyItem_Handler,
		},
		{
			MethodName: "GetMailTemplate",
			Handler:    _MessageService_GetMailTemplate_Handler,
		},
		{
			MethodName: "SaveMailTemplate",
			Handler:    _MessageService_SaveMailTemplate_Handler,
		},
		{
			MethodName: "GetMailTemplates",
			Handler:    _MessageService_GetMailTemplates_Handler,
		},
		{
			MethodName: "DeleteMailTemplate",
			Handler:    _MessageService_DeleteMailTemplate_Handler,
		},
		{
			MethodName: "SendSiteMessage",
			Handler:    _MessageService_SendSiteMessage_Handler,
		},
		{
			MethodName: "SaveNotifyTemplate",
			Handler:    _MessageService_SaveNotifyTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_service.proto",
}
