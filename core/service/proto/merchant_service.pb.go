// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: merchant_service.proto

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

type MerchantOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商户编号
	MerchantId int64 `protobuf:"varint,1,opt,name=merchantId,proto3" json:"merchantId"`
	// 是否分页
	Pagination bool `protobuf:"varint,2,opt,name=pagination,proto3" json:"pagination"`
	// 分页参数
	Params *SPagingParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (x *MerchantOrderRequest) Reset() {
	*x = MerchantOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantOrderRequest) ProtoMessage() {}

func (x *MerchantOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantOrderRequest.ProtoReflect.Descriptor instead.
func (*MerchantOrderRequest) Descriptor() ([]byte, []int) {
	return file_merchant_service_proto_rawDescGZIP(), []int{0}
}

func (x *MerchantOrderRequest) GetMerchantId() int64 {
	if x != nil {
		return x.MerchantId
	}
	return 0
}

func (x *MerchantOrderRequest) GetPagination() bool {
	if x != nil {
		return x.Pagination
	}
	return false
}

func (x *MerchantOrderRequest) GetParams() *SPagingParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type PagingMerchantOrderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64             `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Data  []*SMerchantOrder `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *PagingMerchantOrderListResponse) Reset() {
	*x = PagingMerchantOrderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingMerchantOrderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingMerchantOrderListResponse) ProtoMessage() {}

func (x *PagingMerchantOrderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingMerchantOrderListResponse.ProtoReflect.Descriptor instead.
func (*PagingMerchantOrderListResponse) Descriptor() ([]byte, []int) {
	return file_merchant_service_proto_rawDescGZIP(), []int{1}
}

func (x *PagingMerchantOrderListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagingMerchantOrderListResponse) GetData() []*SMerchantOrder {
	if x != nil {
		return x.Data
	}
	return nil
}

type SMerchantOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   int64  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId"`
	OrderNo   string `protobuf:"bytes,2,opt,name=orderNo,proto3" json:"orderNo"`
	ParentNo  string `protobuf:"bytes,3,opt,name=parentNo,proto3" json:"parentNo"`
	BuyerId   int64  `protobuf:"varint,4,opt,name=buyerId,proto3" json:"buyerId"`
	BuyerName string `protobuf:"bytes,5,opt,name=buyerName,proto3" json:"buyerName"`
	// 订单详情,主要描述订单的内容
	Details        string            `protobuf:"bytes,6,opt,name=details,proto3" json:"details"`
	ItemAmount     float64           `protobuf:"fixed64,7,opt,name=itemAmount,proto3" json:"itemAmount"`
	DiscountAmount float64           `protobuf:"fixed64,8,opt,name=discountAmount,proto3" json:"discountAmount"`
	ExpressFee     float64           `protobuf:"fixed64,9,opt,name=expressFee,proto3" json:"expressFee"`
	PackageFee     float64           `protobuf:"fixed64,10,opt,name=packageFee,proto3" json:"packageFee"`
	IsPaid         bool              `protobuf:"varint,11,opt,name=isPaid,proto3" json:"isPaid"`
	FinalAmount    float64           `protobuf:"fixed64,12,opt,name=finalAmount,proto3" json:"finalAmount"`
	State          int32             `protobuf:"varint,13,opt,name=state,proto3" json:"state"`
	StatusText     string            `protobuf:"bytes,14,opt,name=statusText,proto3" json:"statusText"`
	CreateTime     int64             `protobuf:"varint,15,opt,name=createTime,proto3" json:"createTime"`
	Items          []*SOrderItem     `protobuf:"bytes,16,rep,name=items,proto3" json:"items"`
	Data           map[string]string `protobuf:"bytes,17,rep,name=data,proto3" json:"data" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SMerchantOrder) Reset() {
	*x = SMerchantOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMerchantOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMerchantOrder) ProtoMessage() {}

func (x *SMerchantOrder) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMerchantOrder.ProtoReflect.Descriptor instead.
func (*SMerchantOrder) Descriptor() ([]byte, []int) {
	return file_merchant_service_proto_rawDescGZIP(), []int{2}
}

func (x *SMerchantOrder) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *SMerchantOrder) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *SMerchantOrder) GetParentNo() string {
	if x != nil {
		return x.ParentNo
	}
	return ""
}

func (x *SMerchantOrder) GetBuyerId() int64 {
	if x != nil {
		return x.BuyerId
	}
	return 0
}

func (x *SMerchantOrder) GetBuyerName() string {
	if x != nil {
		return x.BuyerName
	}
	return ""
}

func (x *SMerchantOrder) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *SMerchantOrder) GetItemAmount() float64 {
	if x != nil {
		return x.ItemAmount
	}
	return 0
}

func (x *SMerchantOrder) GetDiscountAmount() float64 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *SMerchantOrder) GetExpressFee() float64 {
	if x != nil {
		return x.ExpressFee
	}
	return 0
}

func (x *SMerchantOrder) GetPackageFee() float64 {
	if x != nil {
		return x.PackageFee
	}
	return 0
}

func (x *SMerchantOrder) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *SMerchantOrder) GetFinalAmount() float64 {
	if x != nil {
		return x.FinalAmount
	}
	return 0
}

func (x *SMerchantOrder) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SMerchantOrder) GetStatusText() string {
	if x != nil {
		return x.StatusText
	}
	return ""
}

func (x *SMerchantOrder) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *SMerchantOrder) GetItems() []*SOrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SMerchantOrder) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_merchant_service_proto protoreflect.FileDescriptor

var file_merchant_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x14, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x5c, 0x0a, 0x1f, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x53, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd5, 0x04, 0x0a, 0x0e, 0x53, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62,
	0x75, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x46, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x46, 0x65, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x46, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x46, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x53, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x32, 0xc6, 0x0f, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x0a, 0x2e, 0x51,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x06,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x0c, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12,
	0x2e, 0x4d, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4d, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x19, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73,
	0x61, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a,
	0x14, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x57, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x5f, 0x12, 0x06, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x1a, 0x17, 0x2e, 0x53, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x11,
	0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x53, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x5f, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x12, 0x15, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x09, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x53, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x53,
	0x65, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x28, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x79,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x06, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x07, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x53,
	0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x1c, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12,
	0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x53,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x22, 0x00, 0x12, 0x22, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12,
	0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0b, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x11,
	0x2e, 0x53, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0d, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69,
	0x50, 0x65, 0x72, 0x6d, 0x12, 0x17, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x41,
	0x70, 0x69, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x79, 0x41, 0x70, 0x69, 0x49, 0x64,
	0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x18, 0x50, 0x61, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x15, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x1b, 0x50, 0x61,
	0x67, 0x65, 0x64, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4f, 0x66, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x15, 0x2e, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x17, 0x50, 0x61, 0x67, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x66, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x15, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x17, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x6f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x54, 0x6f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x63, 0x68, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x12, 0x15, 0x2e, 0x4d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x1a, 0x14, 0x2e, 0x53, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x75, 0x79,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x53, 0x61, 0x76,
	0x65, 0x4d, 0x63, 0x68, 0x42, 0x75, 0x79, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x75, 0x79,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x42, 0x75, 0x79, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x0b, 0x2e, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x1f, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x75, 0x79, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x1a, 0x20, 0x2e, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x52,
	0x65, 0x62, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x52, 0x65,
	0x62, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70,
	0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_merchant_service_proto_rawDescOnce sync.Once
	file_merchant_service_proto_rawDescData = file_merchant_service_proto_rawDesc
)

func file_merchant_service_proto_rawDescGZIP() []byte {
	file_merchant_service_proto_rawDescOnce.Do(func() {
		file_merchant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_merchant_service_proto_rawDescData)
	})
	return file_merchant_service_proto_rawDescData
}

var file_merchant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_merchant_service_proto_goTypes = []interface{}{
	(*MerchantOrderRequest)(nil),            // 0: MerchantOrderRequest
	(*PagingMerchantOrderListResponse)(nil), // 1: PagingMerchantOrderListResponse
	(*SMerchantOrder)(nil),                  // 2: SMerchantOrder
	nil,                                     // 3: SMerchantOrder.DataEntry
	(*SPagingParams)(nil),                   // 4: SPagingParams
	(*SOrderItem)(nil),                      // 5: SOrderItem
	(*Int64)(nil),                           // 6: Int64
	(*String)(nil),                          // 7: String
	(*CreateMerchantRequest)(nil),           // 8: CreateMerchantRequest
	(*SaveMerchantRequest)(nil),             // 9: SaveMerchantRequest
	(*SaveAuthenticateRequest)(nil),         // 10: SaveAuthenticateRequest
	(*MerchantReviewRequest)(nil),           // 11: MerchantReviewRequest
	(*MchUserPwdRequest)(nil),               // 12: MchUserPwdRequest
	(*TradeConfRequest)(nil),                // 13: TradeConfRequest
	(*TradeConfSaveRequest)(nil),            // 14: TradeConfSaveRequest
	(*MemberId)(nil),                        // 15: MemberId
	(*ChangeMemberBindRequest)(nil),         // 16: ChangeMemberBindRequest
	(*MerchantId)(nil),                      // 17: MerchantId
	(*MerchantDisableRequest)(nil),          // 18: MerchantDisableRequest
	(*SaveMerchantSaleConfRequest)(nil),     // 19: SaveMerchantSaleConfRequest
	(*ModifyMerchantPasswordRequest)(nil),   // 20: ModifyMerchantPasswordRequest
	(*MerchantApiPermRequest)(nil),          // 21: MerchantApiPermRequest
	(*WithdrawToMemberAccountRequest)(nil),  // 22: WithdrawToMemberAccountRequest
	(*MerchantChargeRequest)(nil),           // 23: MerchantChargeRequest
	(*MerchantBuyerGroupId)(nil),            // 24: MerchantBuyerGroupId
	(*SaveMerchantBuyerGroupRequest)(nil),   // 25: SaveMerchantBuyerGroupRequest
	(*SaveWholesaleRebateRateRequest)(nil),  // 26: SaveWholesaleRebateRateRequest
	(*QMerchant)(nil),                       // 27: QMerchant
	(*MerchantCreateResponse)(nil),          // 28: MerchantCreateResponse
	(*Result)(nil),                          // 29: Result
	(*MchLoginResponse)(nil),                // 30: MchLoginResponse
	(*SyncWSItemsResponse)(nil),             // 31: SyncWSItemsResponse
	(*STradeConfListResponse)(nil),          // 32: STradeConfListResponse
	(*STradeConf_)(nil),                     // 33: STradeConf_
	(*SMerchantAccount)(nil),                // 34: SMerchantAccount
	(*SMerchantSaleConf)(nil),               // 35: SMerchantSaleConf
	(*SMerchantApiInfo)(nil),                // 36: SMerchantApiInfo
	(*SMerchantBuyerGroup)(nil),             // 37: SMerchantBuyerGroup
	(*MerchantBuyerGroupListResponse)(nil),  // 38: MerchantBuyerGroupListResponse
	(*WholesaleRebateRateListResponse)(nil), // 39: WholesaleRebateRateListResponse
}
var file_merchant_service_proto_depIdxs = []int32{
	4,  // 0: MerchantOrderRequest.params:type_name -> SPagingParams
	2,  // 1: PagingMerchantOrderListResponse.data:type_name -> SMerchantOrder
	5,  // 2: SMerchantOrder.items:type_name -> SOrderItem
	3,  // 3: SMerchantOrder.data:type_name -> SMerchantOrder.DataEntry
	6,  // 4: MerchantService.GetMerchant:input_type -> Int64
	7,  // 5: MerchantService.GetMerchantIdByUsername:input_type -> String
	8,  // 6: MerchantService.CreateMerchant:input_type -> CreateMerchantRequest
	9,  // 7: MerchantService.SaveMerchant:input_type -> SaveMerchantRequest
	10, // 8: MerchantService.SaveAuthenticate:input_type -> SaveAuthenticateRequest
	11, // 9: MerchantService.ReviewAuthenticate:input_type -> MerchantReviewRequest
	12, // 10: MerchantService.CheckLogin:input_type -> MchUserPwdRequest
	6,  // 11: MerchantService.Stat:input_type -> Int64
	6,  // 12: MerchantService.SyncWholesaleItem:input_type -> Int64
	6,  // 13: MerchantService.GetAllTradeConf_:input_type -> Int64
	13, // 14: MerchantService.GetTradeConf:input_type -> TradeConfRequest
	14, // 15: MerchantService.SaveTradeConf:input_type -> TradeConfSaveRequest
	15, // 16: MerchantService.GetMerchantIdByMember:input_type -> MemberId
	16, // 17: MerchantService.ChangeMemberBind:input_type -> ChangeMemberBindRequest
	17, // 18: MerchantService.GetAccount:input_type -> MerchantId
	18, // 19: MerchantService.SetEnabled:input_type -> MerchantDisableRequest
	7,  // 20: MerchantService.GetMerchantIdByHost:input_type -> String
	17, // 21: MerchantService.GetMerchantMajorHost:input_type -> MerchantId
	19, // 22: MerchantService.SaveSaleConf:input_type -> SaveMerchantSaleConfRequest
	17, // 23: MerchantService.GetSaleConf:input_type -> MerchantId
	17, // 24: MerchantService.GetShopId:input_type -> MerchantId
	20, // 25: MerchantService.ChangePassword:input_type -> ModifyMerchantPasswordRequest
	17, // 26: MerchantService.GetApiInfo:input_type -> MerchantId
	21, // 27: MerchantService.ToggleApiPerm:input_type -> MerchantApiPermRequest
	7,  // 28: MerchantService.GetMerchantIdByApiId:input_type -> String
	0,  // 29: MerchantService.PagedNormalOrderOfVendor:input_type -> MerchantOrderRequest
	0,  // 30: MerchantService.PagedWholesaleOrderOfVendor:input_type -> MerchantOrderRequest
	0,  // 31: MerchantService.PagedTradeOrderOfVendor:input_type -> MerchantOrderRequest
	22, // 32: MerchantService.WithdrawToMemberAccount:input_type -> WithdrawToMemberAccountRequest
	23, // 33: MerchantService.ChargeAccount:input_type -> MerchantChargeRequest
	24, // 34: MerchantService.GetMchBuyerGroup_:input_type -> MerchantBuyerGroupId
	25, // 35: MerchantService.SaveMchBuyerGroup:input_type -> SaveMerchantBuyerGroupRequest
	17, // 36: MerchantService.GetBuyerGroups:input_type -> MerchantId
	24, // 37: MerchantService.GetRebateRate:input_type -> MerchantBuyerGroupId
	26, // 38: MerchantService.SaveGroupRebateRate:input_type -> SaveWholesaleRebateRateRequest
	27, // 39: MerchantService.GetMerchant:output_type -> QMerchant
	6,  // 40: MerchantService.GetMerchantIdByUsername:output_type -> Int64
	28, // 41: MerchantService.CreateMerchant:output_type -> MerchantCreateResponse
	29, // 42: MerchantService.SaveMerchant:output_type -> Result
	29, // 43: MerchantService.SaveAuthenticate:output_type -> Result
	29, // 44: MerchantService.ReviewAuthenticate:output_type -> Result
	30, // 45: MerchantService.CheckLogin:output_type -> MchLoginResponse
	29, // 46: MerchantService.Stat:output_type -> Result
	31, // 47: MerchantService.SyncWholesaleItem:output_type -> SyncWSItemsResponse
	32, // 48: MerchantService.GetAllTradeConf_:output_type -> STradeConfListResponse
	33, // 49: MerchantService.GetTradeConf:output_type -> STradeConf_
	29, // 50: MerchantService.SaveTradeConf:output_type -> Result
	6,  // 51: MerchantService.GetMerchantIdByMember:output_type -> Int64
	29, // 52: MerchantService.ChangeMemberBind:output_type -> Result
	34, // 53: MerchantService.GetAccount:output_type -> SMerchantAccount
	29, // 54: MerchantService.SetEnabled:output_type -> Result
	6,  // 55: MerchantService.GetMerchantIdByHost:output_type -> Int64
	7,  // 56: MerchantService.GetMerchantMajorHost:output_type -> String
	29, // 57: MerchantService.SaveSaleConf:output_type -> Result
	35, // 58: MerchantService.GetSaleConf:output_type -> SMerchantSaleConf
	6,  // 59: MerchantService.GetShopId:output_type -> Int64
	29, // 60: MerchantService.ChangePassword:output_type -> Result
	36, // 61: MerchantService.GetApiInfo:output_type -> SMerchantApiInfo
	29, // 62: MerchantService.ToggleApiPerm:output_type -> Result
	6,  // 63: MerchantService.GetMerchantIdByApiId:output_type -> Int64
	1,  // 64: MerchantService.PagedNormalOrderOfVendor:output_type -> PagingMerchantOrderListResponse
	1,  // 65: MerchantService.PagedWholesaleOrderOfVendor:output_type -> PagingMerchantOrderListResponse
	1,  // 66: MerchantService.PagedTradeOrderOfVendor:output_type -> PagingMerchantOrderListResponse
	29, // 67: MerchantService.WithdrawToMemberAccount:output_type -> Result
	29, // 68: MerchantService.ChargeAccount:output_type -> Result
	37, // 69: MerchantService.GetMchBuyerGroup_:output_type -> SMerchantBuyerGroup
	29, // 70: MerchantService.SaveMchBuyerGroup:output_type -> Result
	38, // 71: MerchantService.GetBuyerGroups:output_type -> MerchantBuyerGroupListResponse
	39, // 72: MerchantService.GetRebateRate:output_type -> WholesaleRebateRateListResponse
	29, // 73: MerchantService.SaveGroupRebateRate:output_type -> Result
	39, // [39:74] is the sub-list for method output_type
	4,  // [4:39] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_merchant_service_proto_init() }
func file_merchant_service_proto_init() {
	if File_merchant_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_merchant_dto_proto_init()
	file_message_order_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_merchant_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantOrderRequest); i {
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
		file_merchant_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingMerchantOrderListResponse); i {
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
		file_merchant_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMerchantOrder); i {
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
			RawDescriptor: file_merchant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_merchant_service_proto_goTypes,
		DependencyIndexes: file_merchant_service_proto_depIdxs,
		MessageInfos:      file_merchant_service_proto_msgTypes,
	}.Build()
	File_merchant_service_proto = out.File
	file_merchant_service_proto_rawDesc = nil
	file_merchant_service_proto_goTypes = nil
	file_merchant_service_proto_depIdxs = nil
}
