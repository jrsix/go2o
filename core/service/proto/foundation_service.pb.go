// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: foundation_service.proto

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
		mi := &file_foundation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceSensitiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceSensitiveRequest) ProtoMessage() {}

func (x *ReplaceSensitiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_service_proto_msgTypes[0]
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
	return file_foundation_service_proto_rawDescGZIP(), []int{0}
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

var File_foundation_service_proto protoreflect.FileDescriptor

var file_foundation_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xd4, 0x08, 0x0a, 0x11, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a,
	0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x53, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0e, 0x53, 0x61,
	0x76, 0x65, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x53,
	0x53, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0a, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x6c,
	0x65, 0x61, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x15, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x48,
	0x6f, 0x6f, 0x6b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0b,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x12, 0x08, 0x2e, 0x53, 0x53,
	0x73, 0x6f, 0x41, 0x70, 0x70, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x1d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x1a, 0x08, 0x2e, 0x53, 0x53, 0x73, 0x6f, 0x41, 0x70, 0x70, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x73, 0x6f, 0x41, 0x70, 0x70, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x0d, 0x53, 0x75, 0x70, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x1a, 0x13, 0x2e, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x53, 0x75, 0x70, 0x65, 0x72, 0x50, 0x77,
	0x64, 0x12, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x1a, 0x07, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e,
	0x63, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x49, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e,
	0x41, 0x72, 0x65, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x06, 0x2e,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x1a, 0x11, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x11, 0x2e, 0x53, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x4d,
	0x6f, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x11, 0x2e, 0x53, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x07, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x78, 0x41,
	0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0d, 0x2e, 0x53, 0x57, 0x78, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x00, 0x12, 0x2b, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x57, 0x78, 0x41, 0x70, 0x69, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x0d, 0x2e, 0x53, 0x57, 0x78, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x4d,
	0x63, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x5f, 0x12, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x53, 0x47, 0x6c, 0x6f, 0x62, 0x4d, 0x63, 0x68, 0x53, 0x61,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65,
	0x47, 0x6c, 0x6f, 0x62, 0x4d, 0x63, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x5f,
	0x12, 0x11, 0x2e, 0x53, 0x47, 0x6c, 0x6f, 0x62, 0x4d, 0x63, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32,
	0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foundation_service_proto_rawDescOnce sync.Once
	file_foundation_service_proto_rawDescData = file_foundation_service_proto_rawDesc
)

func file_foundation_service_proto_rawDescGZIP() []byte {
	file_foundation_service_proto_rawDescOnce.Do(func() {
		file_foundation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_foundation_service_proto_rawDescData)
	})
	return file_foundation_service_proto_rawDescData
}

var file_foundation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_foundation_service_proto_goTypes = []interface{}{
	(*ReplaceSensitiveRequest)(nil), // 0: ReplaceSensitiveRequest
	(*String)(nil),                  // 1: String
	(*GetSmsSettingRequest)(nil),    // 2: GetSmsSettingRequest
	(*SSmsProviderSetting)(nil),     // 3: SSmsProviderSetting
	(*CleanCacheRequest)(nil),       // 4: CleanCacheRequest
	(*BoardHookSaveRequest)(nil),    // 5: BoardHookSaveRequest
	(*SSsoApp)(nil),                 // 6: SSsoApp
	(*Empty)(nil),                   // 7: Empty
	(*UserPwd)(nil),                 // 8: UserPwd
	(*GetAreaNamesRequest)(nil),     // 9: GetAreaNamesRequest
	(*AreaStringRequest)(nil),       // 10: AreaStringRequest
	(*Int32)(nil),                   // 11: Int32
	(*SMobileAppConfig)(nil),        // 12: SMobileAppConfig
	(*SWxApiConfig)(nil),            // 13: SWxApiConfig
	(*SGlobMchSaleConf)(nil),        // 14: SGlobMchSaleConf
	(*Bool)(nil),                    // 15: Bool
	(*Result)(nil),                  // 16: Result
	(*CleanCacheResponse)(nil),      // 17: CleanCacheResponse
	(*StringListResponse)(nil),      // 18: StringListResponse
	(*SuperLoginResponse)(nil),      // 19: SuperLoginResponse
	(*IntStringMapResponse)(nil),    // 20: IntStringMapResponse
	(*AreaListResponse)(nil),        // 21: AreaListResponse
	(*PaymentPlatformResponse)(nil), // 22: PaymentPlatformResponse
}
var file_foundation_service_proto_depIdxs = []int32{
	1,  // 0: FoundationService.CheckSensitive:input_type -> String
	0,  // 1: FoundationService.ReplaceSensitive:input_type -> ReplaceSensitiveRequest
	2,  // 2: FoundationService.GetSmsSetting:input_type -> GetSmsSettingRequest
	3,  // 3: FoundationService.SaveSmsSetting:input_type -> SSmsProviderSetting
	4,  // 4: FoundationService.CleanCache:input_type -> CleanCacheRequest
	5,  // 5: FoundationService.SaveBoardHook:input_type -> BoardHookSaveRequest
	1,  // 6: FoundationService.ResourceUrl:input_type -> String
	6,  // 7: FoundationService.RegisterApp:input_type -> SSsoApp
	1,  // 8: FoundationService.GetApp:input_type -> String
	7,  // 9: FoundationService.GetAllSsoApp:input_type -> Empty
	8,  // 10: FoundationService.SuperValidate:input_type -> UserPwd
	8,  // 11: FoundationService.FlushSuperPwd:input_type -> UserPwd
	1,  // 12: FoundationService.GetSyncLoginUrl:input_type -> String
	9,  // 13: FoundationService.GetAreaNames:input_type -> GetAreaNamesRequest
	10, // 14: FoundationService.GetAreaString:input_type -> AreaStringRequest
	11, // 15: FoundationService.GetChildAreas:input_type -> Int32
	7,  // 16: FoundationService.GetMoAppConf:input_type -> Empty
	12, // 17: FoundationService.SaveMoAppConf:input_type -> SMobileAppConfig
	7,  // 18: FoundationService.GetWxApiConfig:input_type -> Empty
	13, // 19: FoundationService.SaveWxApiConfig:input_type -> SWxApiConfig
	7,  // 20: FoundationService.GetPayPlatform:input_type -> Empty
	7,  // 21: FoundationService.GetGlobMchSaleConf_:input_type -> Empty
	14, // 22: FoundationService.SaveGlobMchSaleConf_:input_type -> SGlobMchSaleConf
	15, // 23: FoundationService.CheckSensitive:output_type -> Bool
	1,  // 24: FoundationService.ReplaceSensitive:output_type -> String
	3,  // 25: FoundationService.GetSmsSetting:output_type -> SSmsProviderSetting
	16, // 26: FoundationService.SaveSmsSetting:output_type -> Result
	17, // 27: FoundationService.CleanCache:output_type -> CleanCacheResponse
	16, // 28: FoundationService.SaveBoardHook:output_type -> Result
	1,  // 29: FoundationService.ResourceUrl:output_type -> String
	1,  // 30: FoundationService.RegisterApp:output_type -> String
	6,  // 31: FoundationService.GetApp:output_type -> SSsoApp
	18, // 32: FoundationService.GetAllSsoApp:output_type -> StringListResponse
	19, // 33: FoundationService.SuperValidate:output_type -> SuperLoginResponse
	16, // 34: FoundationService.FlushSuperPwd:output_type -> Result
	1,  // 35: FoundationService.GetSyncLoginUrl:output_type -> String
	20, // 36: FoundationService.GetAreaNames:output_type -> IntStringMapResponse
	1,  // 37: FoundationService.GetAreaString:output_type -> String
	21, // 38: FoundationService.GetChildAreas:output_type -> AreaListResponse
	12, // 39: FoundationService.GetMoAppConf:output_type -> SMobileAppConfig
	16, // 40: FoundationService.SaveMoAppConf:output_type -> Result
	13, // 41: FoundationService.GetWxApiConfig:output_type -> SWxApiConfig
	16, // 42: FoundationService.SaveWxApiConfig:output_type -> Result
	22, // 43: FoundationService.GetPayPlatform:output_type -> PaymentPlatformResponse
	14, // 44: FoundationService.GetGlobMchSaleConf_:output_type -> SGlobMchSaleConf
	16, // 45: FoundationService.SaveGlobMchSaleConf_:output_type -> Result
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_foundation_service_proto_init() }
func file_foundation_service_proto_init() {
	if File_foundation_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_foundation_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_foundation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foundation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foundation_service_proto_goTypes,
		DependencyIndexes: file_foundation_service_proto_depIdxs,
		MessageInfos:      file_foundation_service_proto_msgTypes,
	}.Build()
	File_foundation_service_proto = out.File
	file_foundation_service_proto_rawDesc = nil
	file_foundation_service_proto_goTypes = nil
	file_foundation_service_proto_depIdxs = nil
}
