// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: workorder_service.proto

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

// 分配工单客服请求
type AllocateWorkorderAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 工单编号
	WorkorderId int64 `protobuf:"varint,1,opt,name=workorderId,proto3" json:"workorderId"`
	// 用户编号
	UserId int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId"`
}

func (x *AllocateWorkorderAgentRequest) Reset() {
	*x = AllocateWorkorderAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workorder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocateWorkorderAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocateWorkorderAgentRequest) ProtoMessage() {}

func (x *AllocateWorkorderAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workorder_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocateWorkorderAgentRequest.ProtoReflect.Descriptor instead.
func (*AllocateWorkorderAgentRequest) Descriptor() ([]byte, []int) {
	return file_workorder_service_proto_rawDescGZIP(), []int{0}
}

func (x *AllocateWorkorderAgentRequest) GetWorkorderId() int64 {
	if x != nil {
		return x.WorkorderId
	}
	return 0
}

func (x *AllocateWorkorderAgentRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 工单评价请求
type WorkorderAppriseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 工单编号
	WorkorderId int64 `protobuf:"varint,1,opt,name=workorderId,proto3" json:"workorderId"`
	// 是否有用
	IsUsefully bool `protobuf:"varint,2,opt,name=isUsefully,proto3" json:"isUsefully"`
	// 评分
	Rank int32 `protobuf:"varint,3,opt,name=rank,proto3" json:"rank"`
	// 评价
	Apprise string `protobuf:"bytes,4,opt,name=apprise,proto3" json:"apprise"`
}

func (x *WorkorderAppriseRequest) Reset() {
	*x = WorkorderAppriseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workorder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkorderAppriseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkorderAppriseRequest) ProtoMessage() {}

func (x *WorkorderAppriseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workorder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkorderAppriseRequest.ProtoReflect.Descriptor instead.
func (*WorkorderAppriseRequest) Descriptor() ([]byte, []int) {
	return file_workorder_service_proto_rawDescGZIP(), []int{1}
}

func (x *WorkorderAppriseRequest) GetWorkorderId() int64 {
	if x != nil {
		return x.WorkorderId
	}
	return 0
}

func (x *WorkorderAppriseRequest) GetIsUsefully() bool {
	if x != nil {
		return x.IsUsefully
	}
	return false
}

func (x *WorkorderAppriseRequest) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *WorkorderAppriseRequest) GetApprise() string {
	if x != nil {
		return x.Apprise
	}
	return ""
}

// 提交工单评论请求
type SubmitWorkorderCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 工单编号
	WorkorderId int64 `protobuf:"varint,1,opt,name=workorderId,proto3" json:"workorderId"`
	// 评论内容
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content"`
	// 是否为回复
	IsReplay bool `protobuf:"varint,3,opt,name=isReplay,proto3" json:"isReplay"`
	// 回复的评论编号
	RefCommentId int32 `protobuf:"varint,4,opt,name=refCommentId,proto3" json:"refCommentId"`
}

func (x *SubmitWorkorderCommentRequest) Reset() {
	*x = SubmitWorkorderCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workorder_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitWorkorderCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitWorkorderCommentRequest) ProtoMessage() {}

func (x *SubmitWorkorderCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workorder_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitWorkorderCommentRequest.ProtoReflect.Descriptor instead.
func (*SubmitWorkorderCommentRequest) Descriptor() ([]byte, []int) {
	return file_workorder_service_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitWorkorderCommentRequest) GetWorkorderId() int64 {
	if x != nil {
		return x.WorkorderId
	}
	return 0
}

func (x *SubmitWorkorderCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SubmitWorkorderCommentRequest) GetIsReplay() bool {
	if x != nil {
		return x.IsReplay
	}
	return false
}

func (x *SubmitWorkorderCommentRequest) GetRefCommentId() int32 {
	if x != nil {
		return x.RefCommentId
	}
	return 0
}

// 提交工单请求
type SubmitWorkorderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 会员编号
	MemberId int64 `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId"`
	// 类型, 1: 建议 2:申诉
	ClassId int32 `protobuf:"varint,2,opt,name=classId,proto3" json:"classId"`
	// 关联商户
	MchId int64 `protobuf:"varint,3,opt,name=mchId,proto3" json:"mchId"`
	// 关联业务, 如:CHARGE:2014050060
	Wip string `protobuf:"bytes,4,opt,name=wip,proto3" json:"wip"`
	// Subject
	Subject string `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject"`
	// 投诉内容
	Content string `protobuf:"bytes,6,opt,name=content,proto3" json:"content"`
	// 是否开放评论
	IsOpened int32 `protobuf:"varint,7,opt,name=isOpened,proto3" json:"isOpened"`
	// 诉求描述
	HopeDesc string `protobuf:"bytes,8,opt,name=hopeDesc,proto3" json:"hopeDesc"`
	// 图片列表
	PhotoList string `protobuf:"bytes,9,opt,name=photoList,proto3" json:"photoList"`
}

func (x *SubmitWorkorderRequest) Reset() {
	*x = SubmitWorkorderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workorder_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitWorkorderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitWorkorderRequest) ProtoMessage() {}

func (x *SubmitWorkorderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workorder_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitWorkorderRequest.ProtoReflect.Descriptor instead.
func (*SubmitWorkorderRequest) Descriptor() ([]byte, []int) {
	return file_workorder_service_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitWorkorderRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SubmitWorkorderRequest) GetClassId() int32 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *SubmitWorkorderRequest) GetMchId() int64 {
	if x != nil {
		return x.MchId
	}
	return 0
}

func (x *SubmitWorkorderRequest) GetWip() string {
	if x != nil {
		return x.Wip
	}
	return ""
}

func (x *SubmitWorkorderRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SubmitWorkorderRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SubmitWorkorderRequest) GetIsOpened() int32 {
	if x != nil {
		return x.IsOpened
	}
	return 0
}

func (x *SubmitWorkorderRequest) GetHopeDesc() string {
	if x != nil {
		return x.HopeDesc
	}
	return ""
}

func (x *SubmitWorkorderRequest) GetPhotoList() string {
	if x != nil {
		return x.PhotoList
	}
	return ""
}

// 保存工单响应
type SubmitWorkorderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Msg         string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	WorkorderId int64  `protobuf:"varint,3,opt,name=workorderId,proto3" json:"workorderId"`
}

func (x *SubmitWorkorderResponse) Reset() {
	*x = SubmitWorkorderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workorder_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitWorkorderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitWorkorderResponse) ProtoMessage() {}

func (x *SubmitWorkorderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workorder_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitWorkorderResponse.ProtoReflect.Descriptor instead.
func (*SubmitWorkorderResponse) Descriptor() ([]byte, []int) {
	return file_workorder_service_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitWorkorderResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitWorkorderResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SubmitWorkorderResponse) GetWorkorderId() int64 {
	if x != nil {
		return x.WorkorderId
	}
	return 0
}

// 工单编号
type WorkorderId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 工单编号
	WorkorderId int64 `protobuf:"varint,1,opt,name=workorderId,proto3" json:"workorderId"`
}

func (x *WorkorderId) Reset() {
	*x = WorkorderId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workorder_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkorderId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkorderId) ProtoMessage() {}

func (x *WorkorderId) ProtoReflect() protoreflect.Message {
	mi := &file_workorder_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkorderId.ProtoReflect.Descriptor instead.
func (*WorkorderId) Descriptor() ([]byte, []int) {
	return file_workorder_service_proto_rawDescGZIP(), []int{5}
}

func (x *WorkorderId) GetWorkorderId() int64 {
	if x != nil {
		return x.WorkorderId
	}
	return 0
}

// 工单
type SWorkorder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 工单号
	OrderNo string `protobuf:"bytes,2,opt,name=orderNo,proto3" json:"orderNo"`
	// 会员编号
	MemberId int64 `protobuf:"varint,3,opt,name=memberId,proto3" json:"memberId"`
	// 类型, 1: 建议 2:申诉
	ClassId int32 `protobuf:"varint,4,opt,name=classId,proto3" json:"classId"`
	// 关联商户
	MchId int64 `protobuf:"varint,5,opt,name=mchId,proto3" json:"mchId"`
	// 标志, 1:用户关闭
	Flag int32 `protobuf:"varint,6,opt,name=flag,proto3" json:"flag"`
	// 关联业务, 如:CHARGE:2014050060
	Wip string `protobuf:"bytes,7,opt,name=wip,proto3" json:"wip"`
	// Subject
	Subject string `protobuf:"bytes,8,opt,name=subject,proto3" json:"subject"`
	// 投诉内容
	Content string `protobuf:"bytes,9,opt,name=content,proto3" json:"content"`
	// 是否开放评论
	IsOpened int32 `protobuf:"varint,10,opt,name=isOpened,proto3" json:"isOpened"`
	// 诉求描述
	HopeDesc string `protobuf:"bytes,11,opt,name=hopeDesc,proto3" json:"hopeDesc"`
	// 图片
	FirstPhoto string `protobuf:"bytes,12,opt,name=firstPhoto,proto3" json:"firstPhoto"`
	// 图片列表
	PhotoList string `protobuf:"bytes,13,opt,name=photoList,proto3" json:"photoList"`
	// 状态,1:待处理 2:处理中 3:已完结
	Status int32 `protobuf:"varint,14,opt,name=status,proto3" json:"status"`
	// 分配的客服编号
	AllocateAid int64 `protobuf:"varint,15,opt,name=allocateAid,proto3" json:"allocateAid"`
	// 服务评分
	ServiceRank int32 `protobuf:"varint,16,opt,name=serviceRank,proto3" json:"serviceRank"`
	// 服务评价
	ServiceApprise string `protobuf:"bytes,17,opt,name=serviceApprise,proto3" json:"serviceApprise"`
	// 是否有用 0:未评价 1:是 2:否
	IsUsefully int32 `protobuf:"varint,18,opt,name=isUsefully,proto3" json:"isUsefully"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,19,opt,name=createTime,proto3" json:"createTime"`
	// 更新时间
	UpdateTime int64 `protobuf:"varint,20,opt,name=updateTime,proto3" json:"updateTime"`
}

func (x *SWorkorder) Reset() {
	*x = SWorkorder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workorder_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SWorkorder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SWorkorder) ProtoMessage() {}

func (x *SWorkorder) ProtoReflect() protoreflect.Message {
	mi := &file_workorder_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SWorkorder.ProtoReflect.Descriptor instead.
func (*SWorkorder) Descriptor() ([]byte, []int) {
	return file_workorder_service_proto_rawDescGZIP(), []int{6}
}

func (x *SWorkorder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SWorkorder) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *SWorkorder) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SWorkorder) GetClassId() int32 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *SWorkorder) GetMchId() int64 {
	if x != nil {
		return x.MchId
	}
	return 0
}

func (x *SWorkorder) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *SWorkorder) GetWip() string {
	if x != nil {
		return x.Wip
	}
	return ""
}

func (x *SWorkorder) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SWorkorder) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SWorkorder) GetIsOpened() int32 {
	if x != nil {
		return x.IsOpened
	}
	return 0
}

func (x *SWorkorder) GetHopeDesc() string {
	if x != nil {
		return x.HopeDesc
	}
	return ""
}

func (x *SWorkorder) GetFirstPhoto() string {
	if x != nil {
		return x.FirstPhoto
	}
	return ""
}

func (x *SWorkorder) GetPhotoList() string {
	if x != nil {
		return x.PhotoList
	}
	return ""
}

func (x *SWorkorder) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SWorkorder) GetAllocateAid() int64 {
	if x != nil {
		return x.AllocateAid
	}
	return 0
}

func (x *SWorkorder) GetServiceRank() int32 {
	if x != nil {
		return x.ServiceRank
	}
	return 0
}

func (x *SWorkorder) GetServiceApprise() string {
	if x != nil {
		return x.ServiceApprise
	}
	return ""
}

func (x *SWorkorder) GetIsUsefully() int32 {
	if x != nil {
		return x.IsUsefully
	}
	return 0
}

func (x *SWorkorder) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *SWorkorder) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_workorder_service_proto protoreflect.FileDescriptor

var file_workorder_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x1d, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x72, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x55, 0x73, 0x65, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x55, 0x73, 0x65, 0x66, 0x75, 0x6c, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x72, 0x69, 0x73, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x72, 0x69, 0x73, 0x65, 0x22, 0x9b,
	0x01, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x80, 0x02, 0x0a,
	0x16, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x63, 0x68, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x63,
	0x68, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x77, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4f,
	0x70, 0x65, 0x6e, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x4f,
	0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x70, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x70, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x61, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xb6, 0x04, 0x0a, 0x0a, 0x53, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x63, 0x68, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6d, 0x63, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x77, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x69, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x41,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e,
	0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x61, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x70, 0x70, 0x72, 0x69, 0x73, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x70, 0x72, 0x69, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x73, 0x55, 0x73, 0x65, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x73, 0x55, 0x73, 0x65, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xae, 0x03, 0x0a,
	0x10, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x53, 0x57, 0x6f, 0x72, 0x6b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x56, 0x32, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x56, 0x32, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x0c,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x32, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x05, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x0c, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x32, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x72, 0x69, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x70, 0x70, 0x72, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x32, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1e, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x32, 0x22, 0x00, 0x42, 0x1f, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f,
	0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workorder_service_proto_rawDescOnce sync.Once
	file_workorder_service_proto_rawDescData = file_workorder_service_proto_rawDesc
)

func file_workorder_service_proto_rawDescGZIP() []byte {
	file_workorder_service_proto_rawDescOnce.Do(func() {
		file_workorder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_workorder_service_proto_rawDescData)
	})
	return file_workorder_service_proto_rawDescData
}

var file_workorder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_workorder_service_proto_goTypes = []interface{}{
	(*AllocateWorkorderAgentRequest)(nil), // 0: AllocateWorkorderAgentRequest
	(*WorkorderAppriseRequest)(nil),       // 1: WorkorderAppriseRequest
	(*SubmitWorkorderCommentRequest)(nil), // 2: SubmitWorkorderCommentRequest
	(*SubmitWorkorderRequest)(nil),        // 3: SubmitWorkorderRequest
	(*SubmitWorkorderResponse)(nil),       // 4: SubmitWorkorderResponse
	(*WorkorderId)(nil),                   // 5: WorkorderId
	(*SWorkorder)(nil),                    // 6: SWorkorder
	(*ResultV2)(nil),                      // 7: ResultV2
}
var file_workorder_service_proto_depIdxs = []int32{
	3, // 0: WorkorderService.SubmitWorkorder:input_type -> SubmitWorkorderRequest
	5, // 1: WorkorderService.GetWorkorder:input_type -> WorkorderId
	5, // 2: WorkorderService.DeleteWorkorder:input_type -> WorkorderId
	0, // 3: WorkorderService.AllocateAgentId:input_type -> AllocateWorkorderAgentRequest
	5, // 4: WorkorderService.Finish:input_type -> WorkorderId
	5, // 5: WorkorderService.Close:input_type -> WorkorderId
	1, // 6: WorkorderService.Apprise:input_type -> WorkorderAppriseRequest
	2, // 7: WorkorderService.SubmitComment:input_type -> SubmitWorkorderCommentRequest
	4, // 8: WorkorderService.SubmitWorkorder:output_type -> SubmitWorkorderResponse
	6, // 9: WorkorderService.GetWorkorder:output_type -> SWorkorder
	7, // 10: WorkorderService.DeleteWorkorder:output_type -> ResultV2
	7, // 11: WorkorderService.AllocateAgentId:output_type -> ResultV2
	7, // 12: WorkorderService.Finish:output_type -> ResultV2
	7, // 13: WorkorderService.Close:output_type -> ResultV2
	7, // 14: WorkorderService.Apprise:output_type -> ResultV2
	7, // 15: WorkorderService.SubmitComment:output_type -> ResultV2
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_workorder_service_proto_init() }
func file_workorder_service_proto_init() {
	if File_workorder_service_proto != nil {
		return
	}
	file_global_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_workorder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocateWorkorderAgentRequest); i {
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
		file_workorder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkorderAppriseRequest); i {
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
		file_workorder_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitWorkorderCommentRequest); i {
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
		file_workorder_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitWorkorderRequest); i {
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
		file_workorder_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitWorkorderResponse); i {
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
		file_workorder_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkorderId); i {
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
		file_workorder_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SWorkorder); i {
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
			RawDescriptor: file_workorder_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workorder_service_proto_goTypes,
		DependencyIndexes: file_workorder_service_proto_depIdxs,
		MessageInfos:      file_workorder_service_proto_msgTypes,
	}.Build()
	File_workorder_service_proto = out.File
	file_workorder_service_proto_rawDesc = nil
	file_workorder_service_proto_goTypes = nil
	file_workorder_service_proto_depIdxs = nil
}
