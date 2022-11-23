// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: express_service.proto

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

type ExpressProviderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*SExpressProvider `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
}

func (x *ExpressProviderListResponse) Reset() {
	*x = ExpressProviderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressProviderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressProviderListResponse) ProtoMessage() {}

func (x *ExpressProviderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressProviderListResponse.ProtoReflect.Descriptor instead.
func (*ExpressProviderListResponse) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExpressProviderListResponse) GetValue() []*SExpressProvider {
	if x != nil {
		return x.Value
	}
	return nil
}

// 快递服务商
type SExpressProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 快递公司编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 快递名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// 首字母，用于索引分组
	Letter string `protobuf:"bytes,3,opt,name=letter,proto3" json:"letter"`
	// 分组,多个组,用","隔开
	GroupFlag string `protobuf:"bytes,4,opt,name=groupFlag,proto3" json:"groupFlag"`
	// 快递公司编码
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code"`
	// 接口编码
	ApiCode string `protobuf:"bytes,6,opt,name=apiCode,proto3" json:"apiCode"`
	// 是否启用
	Enabled bool `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled"`
}

func (x *SExpressProvider) Reset() {
	*x = SExpressProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SExpressProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SExpressProvider) ProtoMessage() {}

func (x *SExpressProvider) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SExpressProvider.ProtoReflect.Descriptor instead.
func (*SExpressProvider) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{1}
}

func (x *SExpressProvider) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SExpressProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SExpressProvider) GetLetter() string {
	if x != nil {
		return x.Letter
	}
	return ""
}

func (x *SExpressProvider) GetGroupFlag() string {
	if x != nil {
		return x.GroupFlag
	}
	return ""
}

func (x *SExpressProvider) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SExpressProvider) GetApiCode() string {
	if x != nil {
		return x.ApiCode
	}
	return ""
}

func (x *SExpressProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// 快递模板
type SExpressTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 运营商编号
	SellerId int64 `protobuf:"varint,2,opt,name=sellerId,proto3" json:"sellerId"`
	// 运费模板名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// 是否卖价承担运费
	IsFree bool `protobuf:"varint,4,opt,name=isFree,proto3" json:"isFree"`
	// 运费计价依据
	Basis int32 `protobuf:"varint,5,opt,name=basis,proto3" json:"basis"`
	// 首次计价单位,如首重为2kg
	FirstUnit int32 `protobuf:"varint,6,opt,name=firstUnit,proto3" json:"firstUnit"`
	// 首次计价单价,如续重1kg
	FirstFee int64 `protobuf:"varint,7,opt,name=firstFee,proto3" json:"firstFee"`
	// 超过首次计价计算单位,如续重1kg
	AddUnit int32 `protobuf:"varint,8,opt,name=addUnit,proto3" json:"addUnit"`
	// 超过首次计价单价，如续重1kg
	AddFee int64 `protobuf:"varint,9,opt,name=addFee,proto3" json:"addFee"`
	// 是否启用
	Enabled bool `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled"`
}

func (x *SExpressTemplate) Reset() {
	*x = SExpressTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SExpressTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SExpressTemplate) ProtoMessage() {}

func (x *SExpressTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SExpressTemplate.ProtoReflect.Descriptor instead.
func (*SExpressTemplate) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{2}
}

func (x *SExpressTemplate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SExpressTemplate) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *SExpressTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SExpressTemplate) GetIsFree() bool {
	if x != nil {
		return x.IsFree
	}
	return false
}

func (x *SExpressTemplate) GetBasis() int32 {
	if x != nil {
		return x.Basis
	}
	return 0
}

func (x *SExpressTemplate) GetFirstUnit() int32 {
	if x != nil {
		return x.FirstUnit
	}
	return 0
}

func (x *SExpressTemplate) GetFirstFee() int64 {
	if x != nil {
		return x.FirstFee
	}
	return 0
}

func (x *SExpressTemplate) GetAddUnit() int32 {
	if x != nil {
		return x.AddUnit
	}
	return 0
}

func (x *SExpressTemplate) GetAddFee() int64 {
	if x != nil {
		return x.AddFee
	}
	return 0
}

func (x *SExpressTemplate) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type ExpressTemplateId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellerId   int64 `protobuf:"varint,1,opt,name=sellerId,proto3" json:"sellerId"`
	TemplateId int64 `protobuf:"varint,2,opt,name=templateId,proto3" json:"templateId"`
}

func (x *ExpressTemplateId) Reset() {
	*x = ExpressTemplateId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressTemplateId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressTemplateId) ProtoMessage() {}

func (x *ExpressTemplateId) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressTemplateId.ProtoReflect.Descriptor instead.
func (*ExpressTemplateId) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExpressTemplateId) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *ExpressTemplateId) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

type GetTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellerId int64 `protobuf:"varint,1,opt,name=sellerId,proto3" json:"sellerId"`
	// 仅返回已启用的模板
	OnlyEnabled bool `protobuf:"varint,2,opt,name=onlyEnabled,proto3" json:"onlyEnabled"`
}

func (x *GetTemplatesRequest) Reset() {
	*x = GetTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplatesRequest) ProtoMessage() {}

func (x *GetTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplatesRequest.ProtoReflect.Descriptor instead.
func (*GetTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTemplatesRequest) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *GetTemplatesRequest) GetOnlyEnabled() bool {
	if x != nil {
		return x.OnlyEnabled
	}
	return false
}

type ExpressTemplateListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*SExpressTemplate `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
}

func (x *ExpressTemplateListResponse) Reset() {
	*x = ExpressTemplateListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressTemplateListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressTemplateListResponse) ProtoMessage() {}

func (x *ExpressTemplateListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressTemplateListResponse.ProtoReflect.Descriptor instead.
func (*ExpressTemplateListResponse) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{5}
}

func (x *ExpressTemplateListResponse) GetValue() []*SExpressTemplate {
	if x != nil {
		return x.Value
	}
	return nil
}

// 快递地区模板
type SExpressAreaTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 模板编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 地区编号列表，通常精确到省即可
	CodeList string `protobuf:"bytes,2,opt,name=codeList,proto3" json:"codeList"`
	// 地区名称列表
	NameList string `protobuf:"bytes,3,opt,name=nameList,proto3" json:"nameList"`
	// 首次数值，如 首重为2kg
	FirstUnit int32 `protobuf:"varint,4,opt,name=firstUnit,proto3" json:"firstUnit"`
	// 首次金额，如首重10元
	FirstFee int64 `protobuf:"varint,5,opt,name=firstFee,proto3" json:"firstFee"`
	// 增加数值，如续重1kg
	AddUnit int32 `protobuf:"varint,6,opt,name=addUnit,proto3" json:"addUnit"`
	// 增加产生费用，如续重1kg 10元
	AddFee int64 `protobuf:"varint,7,opt,name=addFee,proto3" json:"addFee"`
}

func (x *SExpressAreaTemplate) Reset() {
	*x = SExpressAreaTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SExpressAreaTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SExpressAreaTemplate) ProtoMessage() {}

func (x *SExpressAreaTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SExpressAreaTemplate.ProtoReflect.Descriptor instead.
func (*SExpressAreaTemplate) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{6}
}

func (x *SExpressAreaTemplate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SExpressAreaTemplate) GetCodeList() string {
	if x != nil {
		return x.CodeList
	}
	return ""
}

func (x *SExpressAreaTemplate) GetNameList() string {
	if x != nil {
		return x.NameList
	}
	return ""
}

func (x *SExpressAreaTemplate) GetFirstUnit() int32 {
	if x != nil {
		return x.FirstUnit
	}
	return 0
}

func (x *SExpressAreaTemplate) GetFirstFee() int64 {
	if x != nil {
		return x.FirstFee
	}
	return 0
}

func (x *SExpressAreaTemplate) GetAddUnit() int32 {
	if x != nil {
		return x.AddUnit
	}
	return 0
}

func (x *SExpressAreaTemplate) GetAddFee() int64 {
	if x != nil {
		return x.AddFee
	}
	return 0
}

type SaveAreaExpTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellerId   int64                 `protobuf:"varint,1,opt,name=sellerId,proto3" json:"sellerId"`
	TemplateId int64                 `protobuf:"varint,2,opt,name=templateId,proto3" json:"templateId"`
	Value      *SExpressAreaTemplate `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
}

func (x *SaveAreaExpTemplateRequest) Reset() {
	*x = SaveAreaExpTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveAreaExpTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveAreaExpTemplateRequest) ProtoMessage() {}

func (x *SaveAreaExpTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveAreaExpTemplateRequest.ProtoReflect.Descriptor instead.
func (*SaveAreaExpTemplateRequest) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{7}
}

func (x *SaveAreaExpTemplateRequest) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *SaveAreaExpTemplateRequest) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *SaveAreaExpTemplateRequest) GetValue() *SExpressAreaTemplate {
	if x != nil {
		return x.Value
	}
	return nil
}

type AreaTemplateId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellerId       int64 `protobuf:"varint,1,opt,name=sellerId,proto3" json:"sellerId"`
	TemplateId     int64 `protobuf:"varint,2,opt,name=templateId,proto3" json:"templateId"`
	AreaTemplateId int64 `protobuf:"varint,3,opt,name=areaTemplateId,proto3" json:"areaTemplateId"`
}

func (x *AreaTemplateId) Reset() {
	*x = AreaTemplateId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaTemplateId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaTemplateId) ProtoMessage() {}

func (x *AreaTemplateId) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaTemplateId.ProtoReflect.Descriptor instead.
func (*AreaTemplateId) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{8}
}

func (x *AreaTemplateId) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *AreaTemplateId) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *AreaTemplateId) GetAreaTemplateId() int64 {
	if x != nil {
		return x.AreaTemplateId
	}
	return 0
}

type SaveTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode    int64  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode"`
	ErrMsg     string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg"`
	TemplateId int64  `protobuf:"varint,3,opt,name=templateId,proto3" json:"templateId"`
}

func (x *SaveTemplateResponse) Reset() {
	*x = SaveTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_express_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTemplateResponse) ProtoMessage() {}

func (x *SaveTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_express_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTemplateResponse.ProtoReflect.Descriptor instead.
func (*SaveTemplateResponse) Descriptor() ([]byte, []int) {
	return file_express_service_proto_rawDescGZIP(), []int{9}
}

func (x *SaveTemplateResponse) GetErrCode() int64 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *SaveTemplateResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SaveTemplateResponse) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

var File_express_service_proto protoreflect.FileDescriptor

var file_express_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x1b, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb4, 0x01,
	0x0a, 0x10, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x22, 0x86, 0x02, 0x0a, 0x10, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x46,
	0x72, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x72, 0x65,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x72, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x64, 0x64, 0x46, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x64, 0x64,
	0x46, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x4f, 0x0a,
	0x11, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x53,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c, 0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6f, 0x6e, 0x6c, 0x79, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x1b, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x14,
	0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x41, 0x72, 0x65, 0x61, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x46, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x46, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x55, 0x6e, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x64, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x46, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x64, 0x64, 0x46, 0x65, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x1a, 0x53, 0x61, 0x76,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x45, 0x78, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x41, 0x72, 0x65,
	0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x74, 0x0a, 0x0e, 0x41, 0x72, 0x65, 0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x61, 0x72, 0x65, 0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x72, 0x65, 0x61, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x32, 0x8c, 0x04, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x49, 0x64, 0x4f, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x11, 0x2e, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x13, 0x53, 0x61, 0x76,
	0x65, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x11, 0x2e, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x53, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x41, 0x72, 0x65, 0x61, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x41, 0x72, 0x65, 0x61,
	0x45, 0x78, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42,
	0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f,
	0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_express_service_proto_rawDescOnce sync.Once
	file_express_service_proto_rawDescData = file_express_service_proto_rawDesc
)

func file_express_service_proto_rawDescGZIP() []byte {
	file_express_service_proto_rawDescOnce.Do(func() {
		file_express_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_express_service_proto_rawDescData)
	})
	return file_express_service_proto_rawDescData
}

var file_express_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_express_service_proto_goTypes = []interface{}{
	(*ExpressProviderListResponse)(nil), // 0: ExpressProviderListResponse
	(*SExpressProvider)(nil),            // 1: SExpressProvider
	(*SExpressTemplate)(nil),            // 2: SExpressTemplate
	(*ExpressTemplateId)(nil),           // 3: ExpressTemplateId
	(*GetTemplatesRequest)(nil),         // 4: GetTemplatesRequest
	(*ExpressTemplateListResponse)(nil), // 5: ExpressTemplateListResponse
	(*SExpressAreaTemplate)(nil),        // 6: SExpressAreaTemplate
	(*SaveAreaExpTemplateRequest)(nil),  // 7: SaveAreaExpTemplateRequest
	(*AreaTemplateId)(nil),              // 8: AreaTemplateId
	(*SaveTemplateResponse)(nil),        // 9: SaveTemplateResponse
	(*IdOrName)(nil),                    // 10: IdOrName
	(*Empty)(nil),                       // 11: Empty
	(*Result)(nil),                      // 12: Result
}
var file_express_service_proto_depIdxs = []int32{
	1,  // 0: ExpressProviderListResponse.value:type_name -> SExpressProvider
	2,  // 1: ExpressTemplateListResponse.value:type_name -> SExpressTemplate
	6,  // 2: SaveAreaExpTemplateRequest.value:type_name -> SExpressAreaTemplate
	10, // 3: ExpressService.GetExpressProvider:input_type -> IdOrName
	1,  // 4: ExpressService.SaveExpressProvider:input_type -> SExpressProvider
	11, // 5: ExpressService.GetProviders:input_type -> Empty
	2,  // 6: ExpressService.SaveTemplate:input_type -> SExpressTemplate
	3,  // 7: ExpressService.GetTemplate:input_type -> ExpressTemplateId
	4,  // 8: ExpressService.GetTemplates:input_type -> GetTemplatesRequest
	3,  // 9: ExpressService.DeleteTemplate:input_type -> ExpressTemplateId
	7,  // 10: ExpressService.SaveAreaTemplate:input_type -> SaveAreaExpTemplateRequest
	8,  // 11: ExpressService.DeleteAreaTemplate:input_type -> AreaTemplateId
	1,  // 12: ExpressService.GetExpressProvider:output_type -> SExpressProvider
	12, // 13: ExpressService.SaveExpressProvider:output_type -> Result
	0,  // 14: ExpressService.GetProviders:output_type -> ExpressProviderListResponse
	9,  // 15: ExpressService.SaveTemplate:output_type -> SaveTemplateResponse
	2,  // 16: ExpressService.GetTemplate:output_type -> SExpressTemplate
	5,  // 17: ExpressService.GetTemplates:output_type -> ExpressTemplateListResponse
	12, // 18: ExpressService.DeleteTemplate:output_type -> Result
	12, // 19: ExpressService.SaveAreaTemplate:output_type -> Result
	12, // 20: ExpressService.DeleteAreaTemplate:output_type -> Result
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_express_service_proto_init() }
func file_express_service_proto_init() {
	if File_express_service_proto != nil {
		return
	}
	file_global_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_express_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressProviderListResponse); i {
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
		file_express_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SExpressProvider); i {
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
		file_express_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SExpressTemplate); i {
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
		file_express_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressTemplateId); i {
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
		file_express_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplatesRequest); i {
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
		file_express_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressTemplateListResponse); i {
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
		file_express_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SExpressAreaTemplate); i {
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
		file_express_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveAreaExpTemplateRequest); i {
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
		file_express_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaTemplateId); i {
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
		file_express_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTemplateResponse); i {
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
			RawDescriptor: file_express_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_express_service_proto_goTypes,
		DependencyIndexes: file_express_service_proto_depIdxs,
		MessageInfos:      file_express_service_proto_msgTypes,
	}.Build()
	File_express_service_proto = out.File
	file_express_service_proto_rawDesc = nil
	file_express_service_proto_goTypes = nil
	file_express_service_proto_depIdxs = nil
}
