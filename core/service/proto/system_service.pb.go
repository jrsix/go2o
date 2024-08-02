// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: system_service.proto

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

// 系统信息
type SSystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 配置最后更新时间
	LastUpdateTime uint64 `protobuf:"varint,1,opt,name=lastUpdateTime,proto3" json:"lastUpdateTime"`
}

func (x *SSystemInfo) Reset() {
	*x = SSystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSystemInfo) ProtoMessage() {}

func (x *SSystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSystemInfo.ProtoReflect.Descriptor instead.
func (*SSystemInfo) Descriptor() ([]byte, []int) {
	return file_system_service_proto_rawDescGZIP(), []int{0}
}

func (x *SSystemInfo) GetLastUpdateTime() uint64 {
	if x != nil {
		return x.LastUpdateTime
	}
	return 0
}

// * 替换敏感词请求
type ReplaceSensitiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text        string `protobuf:"bytes,1,opt,name=text,proto3" json:"text"`
	Replacement string `protobuf:"bytes,2,opt,name=replacement,proto3" json:"replacement"`
}

func (x *ReplaceSensitiveRequest) Reset() {
	*x = ReplaceSensitiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceSensitiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceSensitiveRequest) ProtoMessage() {}

func (x *ReplaceSensitiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceSensitiveRequest.ProtoReflect.Descriptor instead.
func (*ReplaceSensitiveRequest) Descriptor() ([]byte, []int) {
	return file_system_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReplaceSensitiveRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ReplaceSensitiveRequest) GetReplacement() string {
	if x != nil {
		return x.Replacement
	}
	return ""
}

// 获取下级选项请求
type OptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 上级ID
	ParentId uint64 `protobuf:"varint,1,opt,name=parentId,proto3" json:"parentId"`
	// 类型名称,当上级ID为0时,筛选根节点数据
	TypeName string `protobuf:"bytes,2,opt,name=typeName,proto3" json:"typeName"`
}

func (x *OptionsRequest) Reset() {
	*x = OptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionsRequest) ProtoMessage() {}

func (x *OptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionsRequest.ProtoReflect.Descriptor instead.
func (*OptionsRequest) Descriptor() ([]byte, []int) {
	return file_system_service_proto_rawDescGZIP(), []int{2}
}

func (x *OptionsRequest) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *OptionsRequest) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

// 选项响应
type OptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 选项
	Value []*SOption `protobuf:"bytes,1,rep,name=value,proto3" json:"value"`
}

func (x *OptionsResponse) Reset() {
	*x = OptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionsResponse) ProtoMessage() {}

func (x *OptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionsResponse.ProtoReflect.Descriptor instead.
func (*OptionsResponse) Descriptor() ([]byte, []int) {
	return file_system_service_proto_rawDescGZIP(), []int{3}
}

func (x *OptionsResponse) GetValue() []*SOption {
	if x != nil {
		return x.Value
	}
	return nil
}

// 选项
type SOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// 值
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	// 是否为叶子节点
	IsLeaf bool `protobuf:"varint,4,opt,name=isLeaf,proto3" json:"isLeaf"`
}

func (x *SOption) Reset() {
	*x = SOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SOption) ProtoMessage() {}

func (x *SOption) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SOption.ProtoReflect.Descriptor instead.
func (*SOption) Descriptor() ([]byte, []int) {
	return file_system_service_proto_rawDescGZIP(), []int{4}
}

func (x *SOption) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SOption) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SOption) GetIsLeaf() bool {
	if x != nil {
		return x.IsLeaf
	}
	return false
}

// 获取地区名称请求
type FindAreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// 编号
	Code int64 `protobuf:"varint,2,opt,name=code,proto3" json:"code"`
}

func (x *FindAreaRequest) Reset() {
	*x = FindAreaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAreaRequest) ProtoMessage() {}

func (x *FindAreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAreaRequest.ProtoReflect.Descriptor instead.
func (*FindAreaRequest) Descriptor() ([]byte, []int) {
	return file_system_service_proto_rawDescGZIP(), []int{5}
}

func (x *FindAreaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindAreaRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

// 查询地区
type SArea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编码
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// 站点编号
	StationId int64 `protobuf:"varint,3,opt,name=stationId,proto3" json:"stationId"`
}

func (x *SArea) Reset() {
	*x = SArea{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SArea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SArea) ProtoMessage() {}

func (x *SArea) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SArea.ProtoReflect.Descriptor instead.
func (*SArea) Descriptor() ([]byte, []int) {
	return file_system_service_proto_rawDescGZIP(), []int{6}
}

func (x *SArea) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SArea) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SArea) GetStationId() int64 {
	if x != nil {
		return x.StationId
	}
	return 0
}

var File_system_service_proto protoreflect.FileDescriptor

var file_system_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x0b, 0x53, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x31, 0x0a, 0x0f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x5b, 0x0a, 0x07, 0x53, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4c, 0x65,
	0x61, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x66,
	0x22, 0x39, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x53,
	0x41, 0x72, 0x65, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xb7, 0x0a, 0x0a, 0x0d, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x07,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x53, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0a, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x49, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x0f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53,
	0x6d, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x6d, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x53, 0x53, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65,
	0x53, 0x6d, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x53, 0x53, 0x6d,
	0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x43,
	0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x6c, 0x65, 0x61,
	0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x15, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x48, 0x6f, 0x6f,
	0x6b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a,
	0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0b, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x12, 0x08, 0x2e, 0x53, 0x53, 0x73, 0x6f,
	0x41, 0x70, 0x70, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x1d,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x1a, 0x08, 0x2e, 0x53, 0x53, 0x73, 0x6f, 0x41, 0x70, 0x70, 0x22, 0x00, 0x12, 0x2d, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x73, 0x6f, 0x41, 0x70, 0x70, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0d,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x1a, 0x13, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x24,
	0x0a, 0x0d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x53, 0x75, 0x70, 0x65, 0x72, 0x50, 0x77, 0x64, 0x12,
	0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x10, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x49, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x65, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x41, 0x72,
	0x65, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x06, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x1a, 0x11, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x53, 0x41, 0x72, 0x65, 0x61, 0x22, 0x00,
	0x12, 0x2b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x53, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x2d, 0x0a,
	0x0d, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x6f, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x11,
	0x2e, 0x53, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x57, 0x78, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x53, 0x57, 0x78, 0x41, 0x70, 0x69, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x57,
	0x78, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0d, 0x2e, 0x53, 0x57, 0x78,
	0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x47, 0x6c, 0x6f, 0x62, 0x4d, 0x63, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x5f, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x53, 0x47, 0x6c, 0x6f,
	0x62, 0x4d, 0x63, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x14, 0x53, 0x61, 0x76, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x4d, 0x63, 0x68, 0x53, 0x61, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x5f, 0x12, 0x11, 0x2e, 0x53, 0x47, 0x6c, 0x6f, 0x62, 0x4d, 0x63,
	0x68, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_service_proto_rawDescOnce sync.Once
	file_system_service_proto_rawDescData = file_system_service_proto_rawDesc
)

func file_system_service_proto_rawDescGZIP() []byte {
	file_system_service_proto_rawDescOnce.Do(func() {
		file_system_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_service_proto_rawDescData)
	})
	return file_system_service_proto_rawDescData
}

var file_system_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_system_service_proto_goTypes = []interface{}{
	(*SSystemInfo)(nil),             // 0: SSystemInfo
	(*ReplaceSensitiveRequest)(nil), // 1: ReplaceSensitiveRequest
	(*OptionsRequest)(nil),          // 2: OptionsRequest
	(*OptionsResponse)(nil),         // 3: OptionsResponse
	(*SOption)(nil),                 // 4: SOption
	(*FindAreaRequest)(nil),         // 5: FindAreaRequest
	(*SArea)(nil),                   // 6: SArea
	(*String)(nil),                  // 7: String
	(*Empty)(nil),                   // 8: Empty
	(*GetNamesRequest)(nil),         // 9: GetNamesRequest
	(*GetSmsSettingRequest)(nil),    // 10: GetSmsSettingRequest
	(*SSmsProviderSetting)(nil),     // 11: SSmsProviderSetting
	(*CleanCacheRequest)(nil),       // 12: CleanCacheRequest
	(*BoardHookSaveRequest)(nil),    // 13: BoardHookSaveRequest
	(*SSsoApp)(nil),                 // 14: SSsoApp
	(*UserPwd)(nil),                 // 15: UserPwd
	(*AreaStringRequest)(nil),       // 16: AreaStringRequest
	(*Int32)(nil),                   // 17: Int32
	(*SMobileAppConfig)(nil),        // 18: SMobileAppConfig
	(*SWxApiConfig)(nil),            // 19: SWxApiConfig
	(*SGlobMchSaleConf)(nil),        // 20: SGlobMchSaleConf
	(*Bool)(nil),                    // 21: Bool
	(*Result)(nil),                  // 22: Result
	(*IntStringMapResponse)(nil),    // 23: IntStringMapResponse
	(*CleanCacheResponse)(nil),      // 24: CleanCacheResponse
	(*StringListResponse)(nil),      // 25: StringListResponse
	(*SuperLoginResponse)(nil),      // 26: SuperLoginResponse
	(*AreaListResponse)(nil),        // 27: AreaListResponse
	(*PaymentPlatformResponse)(nil), // 28: PaymentPlatformResponse
}
var file_system_service_proto_depIdxs = []int32{
	4,  // 0: OptionsResponse.value:type_name -> SOption
	7,  // 1: SystemService.CheckSensitive:input_type -> String
	1,  // 2: SystemService.ReplaceSensitive:input_type -> ReplaceSensitiveRequest
	8,  // 3: SystemService.GetSystemInfo:input_type -> Empty
	8,  // 4: SystemService.FlushCache:input_type -> Empty
	9,  // 5: SystemService.GetOptionNames:input_type -> GetNamesRequest
	2,  // 6: SystemService.GetChildOptions:input_type -> OptionsRequest
	10, // 7: SystemService.GetSmsSetting:input_type -> GetSmsSettingRequest
	11, // 8: SystemService.SaveSmsSetting:input_type -> SSmsProviderSetting
	12, // 9: SystemService.CleanCache:input_type -> CleanCacheRequest
	13, // 10: SystemService.SaveBoardHook:input_type -> BoardHookSaveRequest
	7,  // 11: SystemService.ResourceUrl:input_type -> String
	14, // 12: SystemService.RegisterApp:input_type -> SSsoApp
	7,  // 13: SystemService.GetApp:input_type -> String
	8,  // 14: SystemService.GetAllSsoApp:input_type -> Empty
	15, // 15: SystemService.SuperValidate:input_type -> UserPwd
	15, // 16: SystemService.FlushSuperPwd:input_type -> UserPwd
	7,  // 17: SystemService.GetSyncLoginUrl:input_type -> String
	9,  // 18: SystemService.GetDistrictNames:input_type -> GetNamesRequest
	16, // 19: SystemService.GetAreaString:input_type -> AreaStringRequest
	17, // 20: SystemService.GetChildAreas:input_type -> Int32
	5,  // 21: SystemService.FindCity:input_type -> FindAreaRequest
	8,  // 22: SystemService.GetMoAppConf:input_type -> Empty
	18, // 23: SystemService.SaveMoAppConf:input_type -> SMobileAppConfig
	8,  // 24: SystemService.GetWxApiConfig:input_type -> Empty
	19, // 25: SystemService.SaveWxApiConfig:input_type -> SWxApiConfig
	8,  // 26: SystemService.GetPayPlatform:input_type -> Empty
	8,  // 27: SystemService.GetGlobMchSaleConf_:input_type -> Empty
	20, // 28: SystemService.SaveGlobMchSaleConf_:input_type -> SGlobMchSaleConf
	21, // 29: SystemService.CheckSensitive:output_type -> Bool
	7,  // 30: SystemService.ReplaceSensitive:output_type -> String
	0,  // 31: SystemService.GetSystemInfo:output_type -> SSystemInfo
	22, // 32: SystemService.FlushCache:output_type -> Result
	23, // 33: SystemService.GetOptionNames:output_type -> IntStringMapResponse
	3,  // 34: SystemService.GetChildOptions:output_type -> OptionsResponse
	11, // 35: SystemService.GetSmsSetting:output_type -> SSmsProviderSetting
	22, // 36: SystemService.SaveSmsSetting:output_type -> Result
	24, // 37: SystemService.CleanCache:output_type -> CleanCacheResponse
	22, // 38: SystemService.SaveBoardHook:output_type -> Result
	7,  // 39: SystemService.ResourceUrl:output_type -> String
	7,  // 40: SystemService.RegisterApp:output_type -> String
	14, // 41: SystemService.GetApp:output_type -> SSsoApp
	25, // 42: SystemService.GetAllSsoApp:output_type -> StringListResponse
	26, // 43: SystemService.SuperValidate:output_type -> SuperLoginResponse
	22, // 44: SystemService.FlushSuperPwd:output_type -> Result
	7,  // 45: SystemService.GetSyncLoginUrl:output_type -> String
	23, // 46: SystemService.GetDistrictNames:output_type -> IntStringMapResponse
	7,  // 47: SystemService.GetAreaString:output_type -> String
	27, // 48: SystemService.GetChildAreas:output_type -> AreaListResponse
	6,  // 49: SystemService.FindCity:output_type -> SArea
	18, // 50: SystemService.GetMoAppConf:output_type -> SMobileAppConfig
	22, // 51: SystemService.SaveMoAppConf:output_type -> Result
	19, // 52: SystemService.GetWxApiConfig:output_type -> SWxApiConfig
	22, // 53: SystemService.SaveWxApiConfig:output_type -> Result
	28, // 54: SystemService.GetPayPlatform:output_type -> PaymentPlatformResponse
	20, // 55: SystemService.GetGlobMchSaleConf_:output_type -> SGlobMchSaleConf
	22, // 56: SystemService.SaveGlobMchSaleConf_:output_type -> Result
	29, // [29:57] is the sub-list for method output_type
	1,  // [1:29] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_system_service_proto_init() }
func file_system_service_proto_init() {
	if File_system_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_system_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_system_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSystemInfo); i {
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
		file_system_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceSensitiveRequest); i {
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
		file_system_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionsRequest); i {
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
		file_system_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionsResponse); i {
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
		file_system_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SOption); i {
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
		file_system_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAreaRequest); i {
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
		file_system_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SArea); i {
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
			RawDescriptor: file_system_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_service_proto_goTypes,
		DependencyIndexes: file_system_service_proto_depIdxs,
		MessageInfos:      file_system_service_proto_msgTypes,
	}.Build()
	File_system_service_proto = out.File
	file_system_service_proto_rawDesc = nil
	file_system_service_proto_goTypes = nil
	file_system_service_proto_depIdxs = nil
}
