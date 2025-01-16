// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.29.2
// source: message_service.proto

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

type NotifyType int32

const (
	// 未知
	NotifyType_NotifyTypeUnknown NotifyType = 0
	// 站内信
	NotifyType_SiteMessage NotifyType = 1
	// 短信
	NotifyType_SMS NotifyType = 2
	// 邮件
	NotifyType_Email NotifyType = 3
)

// Enum value maps for NotifyType.
var (
	NotifyType_name = map[int32]string{
		0: "NotifyTypeUnknown",
		1: "SiteMessage",
		2: "SMS",
		3: "Email",
	}
	NotifyType_value = map[string]int32{
		"NotifyTypeUnknown": 0,
		"SiteMessage":       1,
		"SMS":               2,
		"Email":             3,
	}
)

func (x NotifyType) Enum() *NotifyType {
	p := new(NotifyType)
	*p = x
	return p
}

func (x NotifyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_service_proto_enumTypes[0].Descriptor()
}

func (NotifyType) Type() protoreflect.EnumType {
	return &file_message_service_proto_enumTypes[0]
}

func (x NotifyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyType.Descriptor instead.
func (NotifyType) EnumDescriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{0}
}

// 保存系统通知模板请求
type SaveNotifyTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 模板编号
	TplCode string `protobuf:"bytes,2,opt,name=tplCode,proto3" json:"tplCode"`
	// 模板类型,1:站内信 2:短信 3:邮件
	TplType int32 `protobuf:"varint,3,opt,name=tplType,proto3" json:"tplType"`
	// 模板名称
	TplName string `protobuf:"bytes,4,opt,name=tplName,proto3" json:"tplName"`
	// 模板内容
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content"`
	// 模板标签, 多个用,隔开
	Labels string `protobuf:"bytes,6,opt,name=labels,proto3" json:"labels"`
	// 短信服务商代码
	SpCode string `protobuf:"bytes,7,opt,name=spCode,proto3" json:"spCode"`
	// 短信服务商模板编号
	SpTid string `protobuf:"bytes,8,opt,name=spTid,proto3" json:"spTid"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,9,opt,name=createTime,proto3" json:"createTime"`
	// UpdateTime
	UpdateTime int64 `protobuf:"varint,10,opt,name=updateTime,proto3" json:"updateTime"`
	// 是否删除,0:否 1:是
	IsDeleted int32 `protobuf:"varint,11,opt,name=isDeleted,proto3" json:"isDeleted"`
}

func (x *SaveNotifyTemplateRequest) Reset() {
	*x = SaveNotifyTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveNotifyTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveNotifyTemplateRequest) ProtoMessage() {}

func (x *SaveNotifyTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveNotifyTemplateRequest.ProtoReflect.Descriptor instead.
func (*SaveNotifyTemplateRequest) Descriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *SaveNotifyTemplateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveNotifyTemplateRequest) GetTplCode() string {
	if x != nil {
		return x.TplCode
	}
	return ""
}

func (x *SaveNotifyTemplateRequest) GetTplType() int32 {
	if x != nil {
		return x.TplType
	}
	return 0
}

func (x *SaveNotifyTemplateRequest) GetTplName() string {
	if x != nil {
		return x.TplName
	}
	return ""
}

func (x *SaveNotifyTemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SaveNotifyTemplateRequest) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *SaveNotifyTemplateRequest) GetSpCode() string {
	if x != nil {
		return x.SpCode
	}
	return ""
}

func (x *SaveNotifyTemplateRequest) GetSpTid() string {
	if x != nil {
		return x.SpTid
	}
	return ""
}

func (x *SaveNotifyTemplateRequest) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *SaveNotifyTemplateRequest) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *SaveNotifyTemplateRequest) GetIsDeleted() int32 {
	if x != nil {
		return x.IsDeleted
	}
	return 0
}

var File_message_service_proto protoreflect.FileDescriptor

var file_message_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb7, 0x02, 0x0a, 0x19, 0x53, 0x61, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x70, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x70, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x54, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70, 0x54, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2a, 0x48, 0x0a, 0x0a, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x4d, 0x53, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x10, 0x03, 0x32, 0xf2, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x1a, 0x0c, 0x2e, 0x53, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x32, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x17, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0c, 0x2e, 0x53, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x0e, 0x2e, 0x53,
	0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x10,
	0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x2e, 0x53, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x12,
	0x53, 0x61, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09,
	0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_message_service_proto_rawDescOnce sync.Once
	file_message_service_proto_rawDescData = file_message_service_proto_rawDesc
)

func file_message_service_proto_rawDescGZIP() []byte {
	file_message_service_proto_rawDescOnce.Do(func() {
		file_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_service_proto_rawDescData)
	})
	return file_message_service_proto_rawDescData
}

var file_message_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_service_proto_goTypes = []interface{}{
	(NotifyType)(0),                   // 0: NotifyType
	(*SaveNotifyTemplateRequest)(nil), // 1: SaveNotifyTemplateRequest
	(*String)(nil),                    // 2: String
	(*SendMessageRequest)(nil),        // 3: SendMessageRequest
	(*Empty)(nil),                     // 4: Empty
	(*SNotifyItem)(nil),               // 5: SNotifyItem
	(*Int64)(nil),                     // 6: Int64
	(*SMailTemplate)(nil),             // 7: SMailTemplate
	(*SendSiteMessageRequest)(nil),    // 8: SendSiteMessageRequest
	(*TxResult)(nil),                  // 9: TxResult
	(*NotifyItemListResponse)(nil),    // 10: NotifyItemListResponse
	(*Result)(nil),                    // 11: Result
	(*MailTemplateListResponse)(nil),  // 12: MailTemplateListResponse
}
var file_message_service_proto_depIdxs = []int32{
	2,  // 0: MessageService.GetNotifyItem:input_type -> String
	3,  // 1: MessageService.SendPhoneMessage:input_type -> SendMessageRequest
	4,  // 2: MessageService.GetAllNotifyItem:input_type -> Empty
	5,  // 3: MessageService.SaveNotifyItem:input_type -> SNotifyItem
	6,  // 4: MessageService.GetMailTemplate:input_type -> Int64
	7,  // 5: MessageService.SaveMailTemplate:input_type -> SMailTemplate
	4,  // 6: MessageService.GetMailTemplates:input_type -> Empty
	6,  // 7: MessageService.DeleteMailTemplate:input_type -> Int64
	8,  // 8: MessageService.SendSiteMessage:input_type -> SendSiteMessageRequest
	1,  // 9: MessageService.SaveNotifyTemplate:input_type -> SaveNotifyTemplateRequest
	5,  // 10: MessageService.GetNotifyItem:output_type -> SNotifyItem
	9,  // 11: MessageService.SendPhoneMessage:output_type -> TxResult
	10, // 12: MessageService.GetAllNotifyItem:output_type -> NotifyItemListResponse
	11, // 13: MessageService.SaveNotifyItem:output_type -> Result
	7,  // 14: MessageService.GetMailTemplate:output_type -> SMailTemplate
	11, // 15: MessageService.SaveMailTemplate:output_type -> Result
	12, // 16: MessageService.GetMailTemplates:output_type -> MailTemplateListResponse
	11, // 17: MessageService.DeleteMailTemplate:output_type -> Result
	11, // 18: MessageService.SendSiteMessage:output_type -> Result
	9,  // 19: MessageService.SaveNotifyTemplate:output_type -> TxResult
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_message_service_proto_init() }
func file_message_service_proto_init() {
	if File_message_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_message_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveNotifyTemplateRequest); i {
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
			RawDescriptor: file_message_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_service_proto_goTypes,
		DependencyIndexes: file_message_service_proto_depIdxs,
		EnumInfos:         file_message_service_proto_enumTypes,
		MessageInfos:      file_message_service_proto_msgTypes,
	}.Build()
	File_message_service_proto = out.File
	file_message_service_proto_rawDesc = nil
	file_message_service_proto_goTypes = nil
	file_message_service_proto_depIdxs = nil
}
