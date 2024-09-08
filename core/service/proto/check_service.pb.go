//*
// Copyright (C) 2021 56X.NET, All rights reserved.
//
// name : template_service.proto
// author : jarrysix
// date : 2024/06/07 23:37:45
// description : 校验服务
// history :

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.26.1
// source: check_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 发送验证码请求
type SendCheckCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	// 会员编号
	UserId int64 `protobuf:"zigzag64,2,opt,name=userId,proto3" json:"userId"`
	// 接收账号（手机号、邮箱）
	ReceptAccount string `protobuf:"bytes,3,opt,name=receptAccount,proto3" json:"receptAccount"`
	// 操作
	Operation string `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation"`
	// 发送消息模板编号
	TemplateCode string `protobuf:"bytes,5,opt,name=templateCode,proto3" json:"templateCode"`
	// 有效时间（分钟）
	Effective int32 `protobuf:"zigzag32,6,opt,name=effective,proto3" json:"effective"`
}

func (x *SendCheckCodeRequest) Reset() {
	*x = SendCheckCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCheckCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCheckCodeRequest) ProtoMessage() {}

func (x *SendCheckCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCheckCodeRequest.ProtoReflect.Descriptor instead.
func (*SendCheckCodeRequest) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendCheckCodeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendCheckCodeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SendCheckCodeRequest) GetReceptAccount() string {
	if x != nil {
		return x.ReceptAccount
	}
	return ""
}

func (x *SendCheckCodeRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *SendCheckCodeRequest) GetTemplateCode() string {
	if x != nil {
		return x.TemplateCode
	}
	return ""
}

func (x *SendCheckCodeRequest) GetEffective() int32 {
	if x != nil {
		return x.Effective
	}
	return 0
}

// 发送验证码响应
type SendCheckCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// 错误消息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// 验证码
	CheckCode string `protobuf:"bytes,3,opt,name=checkCode,proto3" json:"checkCode"`
}

func (x *SendCheckCodeResponse) Reset() {
	*x = SendCheckCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCheckCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCheckCodeResponse) ProtoMessage() {}

func (x *SendCheckCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCheckCodeResponse.ProtoReflect.Descriptor instead.
func (*SendCheckCodeResponse) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendCheckCodeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SendCheckCodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendCheckCodeResponse) GetCheckCode() string {
	if x != nil {
		return x.CheckCode
	}
	return ""
}

// 比较验证码请求
type CompareCheckCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	// 接收账号（手机号、邮箱）
	ReceptAccount string `protobuf:"bytes,2,opt,name=receptAccount,proto3" json:"receptAccount"`
	// 校验码
	CheckCode string `protobuf:"bytes,3,opt,name=checkCode,proto3" json:"checkCode"`
	// 验证成功后是否重置
	ResetIfOk bool `protobuf:"varint,4,opt,name=resetIfOk,proto3" json:"resetIfOk"`
}

func (x *CompareCheckCodeRequest) Reset() {
	*x = CompareCheckCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareCheckCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareCheckCodeRequest) ProtoMessage() {}

func (x *CompareCheckCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareCheckCodeRequest.ProtoReflect.Descriptor instead.
func (*CompareCheckCodeRequest) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{2}
}

func (x *CompareCheckCodeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CompareCheckCodeRequest) GetReceptAccount() string {
	if x != nil {
		return x.ReceptAccount
	}
	return ""
}

func (x *CompareCheckCodeRequest) GetCheckCode() string {
	if x != nil {
		return x.CheckCode
	}
	return ""
}

func (x *CompareCheckCodeRequest) GetResetIfOk() bool {
	if x != nil {
		return x.ResetIfOk
	}
	return false
}

// 比较验证码响应
type CompareCheckCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// 错误消息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// 会员编号
	UserId int64 `protobuf:"zigzag64,3,opt,name=userId,proto3" json:"userId"`
}

func (x *CompareCheckCodeResponse) Reset() {
	*x = CompareCheckCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareCheckCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareCheckCodeResponse) ProtoMessage() {}

func (x *CompareCheckCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareCheckCodeResponse.ProtoReflect.Descriptor instead.
func (*CompareCheckCodeResponse) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{3}
}

func (x *CompareCheckCodeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CompareCheckCodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CompareCheckCodeResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 请求访问令牌
type GrantAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户类型 1:会员 2:商户 3:系统用户
	UserType int32 `protobuf:"zigzag32,1,opt,name=userType,proto3" json:"userType"`
	// 会员编号
	UserId int64 `protobuf:"zigzag64,2,opt,name=userId,proto3" json:"userId"`
	// 令牌过期时间(单位:s)
	ExpiresTime int64 `protobuf:"zigzag64,3,opt,name=expiresTime,proto3" json:"expiresTime"`
	// 令牌主题
	Sub string `protobuf:"bytes,4,opt,name=sub,proto3" json:"sub"`
}

func (x *GrantAccessTokenRequest) Reset() {
	*x = GrantAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantAccessTokenRequest) ProtoMessage() {}

func (x *GrantAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*GrantAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{4}
}

func (x *GrantAccessTokenRequest) GetUserType() int32 {
	if x != nil {
		return x.UserType
	}
	return 0
}

func (x *GrantAccessTokenRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GrantAccessTokenRequest) GetExpiresTime() int64 {
	if x != nil {
		return x.ExpiresTime
	}
	return 0
}

func (x *GrantAccessTokenRequest) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

// 发放访问令牌响应
type GrantAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// 错误信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// 访问令牌
	AccessToken string `protobuf:"bytes,3,opt,name=accessToken,proto3" json:"accessToken"`
}

func (x *GrantAccessTokenResponse) Reset() {
	*x = GrantAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantAccessTokenResponse) ProtoMessage() {}

func (x *GrantAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*GrantAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{5}
}

func (x *GrantAccessTokenResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GrantAccessTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GrantAccessTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// 检查令牌请求
type CheckAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 访问令牌
	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken"`
	// 检测过期时间,在此时间内过期,将生成返回新的令牌
	CheckExpireTime int64 `protobuf:"zigzag64,2,opt,name=checkExpireTime,proto3" json:"checkExpireTime"`
	// 新生成令牌的有效时间
	RenewExpiresTime int64 `protobuf:"zigzag64,3,opt,name=renewExpiresTime,proto3" json:"renewExpiresTime"`
	// 令牌主题
	Sub string `protobuf:"bytes,4,opt,name=sub,proto3" json:"sub"`
}

func (x *CheckAccessTokenRequest) Reset() {
	*x = CheckAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessTokenRequest) ProtoMessage() {}

func (x *CheckAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*CheckAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{6}
}

func (x *CheckAccessTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CheckAccessTokenRequest) GetCheckExpireTime() int64 {
	if x != nil {
		return x.CheckExpireTime
	}
	return 0
}

func (x *CheckAccessTokenRequest) GetRenewExpiresTime() int64 {
	if x != nil {
		return x.RenewExpiresTime
	}
	return 0
}

func (x *CheckAccessTokenRequest) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

// 检查令牌响应
type CheckAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// 错误信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// 用户类型 1:会员 2:商户 3:系统用户
	UserType int32 `protobuf:"zigzag32,3,opt,name=userType,proto3" json:"userType"`
	// 会员编号
	UserId int64 `protobuf:"zigzag64,4,opt,name=userId,proto3" json:"userId"`
	// 是否过期
	IsExpires bool `protobuf:"varint,5,opt,name=isExpires,proto3" json:"isExpires"`
	// 令牌过期时间
	TokenExpiresTime int64 `protobuf:"zigzag64,6,opt,name=tokenExpiresTime,proto3" json:"tokenExpiresTime"`
	// 重新生成的令牌
	RenewAccessToken string `protobuf:"bytes,7,opt,name=renewAccessToken,proto3" json:"renewAccessToken"`
}

func (x *CheckAccessTokenResponse) Reset() {
	*x = CheckAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessTokenResponse) ProtoMessage() {}

func (x *CheckAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*CheckAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_check_service_proto_rawDescGZIP(), []int{7}
}

func (x *CheckAccessTokenResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CheckAccessTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CheckAccessTokenResponse) GetUserType() int32 {
	if x != nil {
		return x.UserType
	}
	return 0
}

func (x *CheckAccessTokenResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CheckAccessTokenResponse) GetIsExpires() bool {
	if x != nil {
		return x.IsExpires
	}
	return false
}

func (x *CheckAccessTokenResponse) GetTokenExpiresTime() int64 {
	if x != nil {
		return x.TokenExpiresTime
	}
	return 0
}

func (x *CheckAccessTokenResponse) GetRenewAccessToken() string {
	if x != nil {
		return x.RenewAccessToken
	}
	return ""
}

var File_check_service_proto protoreflect.FileDescriptor

var file_check_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x22, 0x63, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63,
	0x65, 0x70, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x65, 0x74, 0x49, 0x66, 0x4f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x65, 0x74, 0x49, 0x66, 0x4f, 0x6b, 0x22, 0x60, 0x0a, 0x18, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x81, 0x01,
	0x0a, 0x17, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75,
	0x62, 0x22, 0x6a, 0x0a, 0x18, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa3, 0x01,
	0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x12, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x10, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x75, 0x62, 0x22, 0xf2, 0x01, 0x0a, 0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x12, 0x2a, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x10, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x72, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa7, 0x02, 0x0a, 0x0c, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_service_proto_rawDescOnce sync.Once
	file_check_service_proto_rawDescData = file_check_service_proto_rawDesc
)

func file_check_service_proto_rawDescGZIP() []byte {
	file_check_service_proto_rawDescOnce.Do(func() {
		file_check_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_service_proto_rawDescData)
	})
	return file_check_service_proto_rawDescData
}

var file_check_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_check_service_proto_goTypes = []interface{}{
	(*SendCheckCodeRequest)(nil),     // 0: SendCheckCodeRequest
	(*SendCheckCodeResponse)(nil),    // 1: SendCheckCodeResponse
	(*CompareCheckCodeRequest)(nil),  // 2: CompareCheckCodeRequest
	(*CompareCheckCodeResponse)(nil), // 3: CompareCheckCodeResponse
	(*GrantAccessTokenRequest)(nil),  // 4: GrantAccessTokenRequest
	(*GrantAccessTokenResponse)(nil), // 5: GrantAccessTokenResponse
	(*CheckAccessTokenRequest)(nil),  // 6: CheckAccessTokenRequest
	(*CheckAccessTokenResponse)(nil), // 7: CheckAccessTokenResponse
}
var file_check_service_proto_depIdxs = []int32{
	0, // 0: CheckService.sendCode:input_type -> SendCheckCodeRequest
	2, // 1: CheckService.compareCode:input_type -> CompareCheckCodeRequest
	4, // 2: CheckService.grantAccessToken:input_type -> GrantAccessTokenRequest
	6, // 3: CheckService.checkAccessToken:input_type -> CheckAccessTokenRequest
	1, // 4: CheckService.sendCode:output_type -> SendCheckCodeResponse
	3, // 5: CheckService.compareCode:output_type -> CompareCheckCodeResponse
	5, // 6: CheckService.grantAccessToken:output_type -> GrantAccessTokenResponse
	7, // 7: CheckService.checkAccessToken:output_type -> CheckAccessTokenResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_service_proto_init() }
func file_check_service_proto_init() {
	if File_check_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCheckCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_check_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCheckCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_check_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareCheckCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_check_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareCheckCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_check_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantAccessTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_check_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantAccessTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_check_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccessTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_check_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAccessTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_check_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_check_service_proto_goTypes,
		DependencyIndexes: file_check_service_proto_depIdxs,
		MessageInfos:      file_check_service_proto_msgTypes,
	}.Build()
	File_check_service_proto = out.File
	file_check_service_proto_rawDesc = nil
	file_check_service_proto_goTypes = nil
	file_check_service_proto_depIdxs = nil
}
