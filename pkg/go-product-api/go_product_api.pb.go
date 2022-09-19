// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: alexeykirinyuk/go_product_api/v1/go_product_api.proto

package go_product_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Currency int32

const (
	Currency_CURRENCY_UNKNOWN Currency = 0
	Currency_CURRENCY_RUB     Currency = 1
	Currency_CURRENCY_USD     Currency = 2
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "CURRENCY_UNKNOWN",
		1: "CURRENCY_RUB",
		2: "CURRENCY_USD",
	}
	Currency_value = map[string]int32{
		"CURRENCY_UNKNOWN": 0,
		"CURRENCY_RUB":     1,
		"CURRENCY_USD":     2,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_enumTypes[0].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_enumTypes[0]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{0}
}

type CreateProductV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category    string   `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Brand       string   `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Cost        float32  `protobuf:"fixed32,5,opt,name=cost,proto3" json:"cost,omitempty"`
	Currency    Currency `protobuf:"varint,6,opt,name=currency,proto3,enum=ozonmp.go_product_api.v1.Currency" json:"currency,omitempty"`
}

func (x *CreateProductV1Request) Reset() {
	*x = CreateProductV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductV1Request) ProtoMessage() {}

func (x *CreateProductV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductV1Request.ProtoReflect.Descriptor instead.
func (*CreateProductV1Request) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductV1Request) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateProductV1Request) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductV1Request) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CreateProductV1Request) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *CreateProductV1Request) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_CURRENCY_UNKNOWN
}

type CreateProductV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *CreateProductV1Response) Reset() {
	*x = CreateProductV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductV1Response) ProtoMessage() {}

func (x *CreateProductV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductV1Response.ProtoReflect.Descriptor instead.
func (*CreateProductV1Response) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductV1Response) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type DescribeProductV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *DescribeProductV1Request) Reset() {
	*x = DescribeProductV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProductV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProductV1Request) ProtoMessage() {}

func (x *DescribeProductV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProductV1Request.ProtoReflect.Descriptor instead.
func (*DescribeProductV1Request) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeProductV1Request) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type DescribeProductV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *DescribeProductV1Response) Reset() {
	*x = DescribeProductV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProductV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProductV1Response) ProtoMessage() {}

func (x *DescribeProductV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProductV1Response.ProtoReflect.Descriptor instead.
func (*DescribeProductV1Response) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeProductV1Response) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ListProductsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProductsV1Request) Reset() {
	*x = ListProductsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsV1Request) ProtoMessage() {}

func (x *ListProductsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsV1Request.ProtoReflect.Descriptor instead.
func (*ListProductsV1Request) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{4}
}

type ListProductsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListProductsV1Response) Reset() {
	*x = ListProductsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsV1Response) ProtoMessage() {}

func (x *ListProductsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsV1Response.ProtoReflect.Descriptor instead.
func (*ListProductsV1Response) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListProductsV1Response) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type RemoveProductV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *RemoveProductV1Request) Reset() {
	*x = RemoveProductV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductV1Request) ProtoMessage() {}

func (x *RemoveProductV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductV1Request.ProtoReflect.Descriptor instead.
func (*RemoveProductV1Request) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveProductV1Request) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type RemoveProductV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *RemoveProductV1Response) Reset() {
	*x = RemoveProductV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductV1Response) ProtoMessage() {}

func (x *RemoveProductV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductV1Response.ProtoReflect.Descriptor instead.
func (*RemoveProductV1Response) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveProductV1Response) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category    string                 `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Brand       string                 `protobuf:"bytes,5,opt,name=brand,proto3" json:"brand,omitempty"`
	Cost        float32                `protobuf:"fixed32,6,opt,name=cost,proto3" json:"cost,omitempty"`
	Currency    Currency               `protobuf:"varint,7,opt,name=currency,proto3,enum=ozonmp.go_product_api.v1.Currency" json:"currency,omitempty"`
	Created     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP(), []int{8}
}

func (x *Product) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Product) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Product) GetCurrency() Currency {
	if x != nil {
		return x.Currency
	}
	return Currency_CURRENCY_UNKNOWN
}

func (x *Product) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

var File_alexeykirinyuk_go_product_api_v1_go_product_api_proto protoreflect.FileDescriptor

var file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x6c, 0x65, 0x78, 0x65, 0x79, 0x6b, 0x69, 0x72, 0x69, 0x6e, 0x79, 0x75, 0x6b,
	0x2f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01,
	0x18, 0x32, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x25, 0x00,
	0x00, 0x00, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x20, 0x00, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0x56, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x42, 0x0a, 0x18, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x58, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x57, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x16, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x56, 0x0a,
	0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x8b, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x3e, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x2a, 0x44, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x14, 0x0a, 0x10, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43,
	0x59, 0x5f, 0x52, 0x55, 0x42, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x55, 0x52, 0x52, 0x45,
	0x4e, 0x43, 0x59, 0x5f, 0x55, 0x53, 0x44, 0x10, 0x02, 0x32, 0xf3, 0x04, 0x0a, 0x13, 0x47, 0x6f,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x98, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x12, 0x32, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70,
	0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x95, 0x01, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31,
	0x12, 0x30, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x90, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x56, 0x31, 0x12, 0x2f, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70,
	0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d,
	0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x95, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x12, 0x30, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x42,
	0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x65, 0x78, 0x65, 0x79, 0x6b, 0x69, 0x72, 0x69, 0x6e, 0x79, 0x75, 0x6b, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescOnce sync.Once
	file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescData = file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDesc
)

func file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescGZIP() []byte {
	file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescOnce.Do(func() {
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescData)
	})
	return file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDescData
}

var file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_goTypes = []interface{}{
	(Currency)(0),                     // 0: ozonmp.go_product_api.v1.Currency
	(*CreateProductV1Request)(nil),    // 1: ozonmp.go_product_api.v1.CreateProductV1Request
	(*CreateProductV1Response)(nil),   // 2: ozonmp.go_product_api.v1.CreateProductV1Response
	(*DescribeProductV1Request)(nil),  // 3: ozonmp.go_product_api.v1.DescribeProductV1Request
	(*DescribeProductV1Response)(nil), // 4: ozonmp.go_product_api.v1.DescribeProductV1Response
	(*ListProductsV1Request)(nil),     // 5: ozonmp.go_product_api.v1.ListProductsV1Request
	(*ListProductsV1Response)(nil),    // 6: ozonmp.go_product_api.v1.ListProductsV1Response
	(*RemoveProductV1Request)(nil),    // 7: ozonmp.go_product_api.v1.RemoveProductV1Request
	(*RemoveProductV1Response)(nil),   // 8: ozonmp.go_product_api.v1.RemoveProductV1Response
	(*Product)(nil),                   // 9: ozonmp.go_product_api.v1.Product
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
}
var file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_depIdxs = []int32{
	0,  // 0: ozonmp.go_product_api.v1.CreateProductV1Request.currency:type_name -> ozonmp.go_product_api.v1.Currency
	9,  // 1: ozonmp.go_product_api.v1.CreateProductV1Response.product:type_name -> ozonmp.go_product_api.v1.Product
	9,  // 2: ozonmp.go_product_api.v1.DescribeProductV1Response.product:type_name -> ozonmp.go_product_api.v1.Product
	9,  // 3: ozonmp.go_product_api.v1.ListProductsV1Response.products:type_name -> ozonmp.go_product_api.v1.Product
	9,  // 4: ozonmp.go_product_api.v1.RemoveProductV1Response.product:type_name -> ozonmp.go_product_api.v1.Product
	0,  // 5: ozonmp.go_product_api.v1.Product.currency:type_name -> ozonmp.go_product_api.v1.Currency
	10, // 6: ozonmp.go_product_api.v1.Product.created:type_name -> google.protobuf.Timestamp
	3,  // 7: ozonmp.go_product_api.v1.GoProductApiService.DescribeProductV1:input_type -> ozonmp.go_product_api.v1.DescribeProductV1Request
	1,  // 8: ozonmp.go_product_api.v1.GoProductApiService.CreateProductV1:input_type -> ozonmp.go_product_api.v1.CreateProductV1Request
	5,  // 9: ozonmp.go_product_api.v1.GoProductApiService.ListProductsV1:input_type -> ozonmp.go_product_api.v1.ListProductsV1Request
	7,  // 10: ozonmp.go_product_api.v1.GoProductApiService.RemoveProductV1:input_type -> ozonmp.go_product_api.v1.RemoveProductV1Request
	4,  // 11: ozonmp.go_product_api.v1.GoProductApiService.DescribeProductV1:output_type -> ozonmp.go_product_api.v1.DescribeProductV1Response
	2,  // 12: ozonmp.go_product_api.v1.GoProductApiService.CreateProductV1:output_type -> ozonmp.go_product_api.v1.CreateProductV1Response
	6,  // 13: ozonmp.go_product_api.v1.GoProductApiService.ListProductsV1:output_type -> ozonmp.go_product_api.v1.ListProductsV1Response
	8,  // 14: ozonmp.go_product_api.v1.GoProductApiService.RemoveProductV1:output_type -> ozonmp.go_product_api.v1.RemoveProductV1Response
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_init() }
func file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_init() {
	if File_alexeykirinyuk_go_product_api_v1_go_product_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductV1Request); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductV1Response); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProductV1Request); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProductV1Response); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsV1Request); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsV1Response); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductV1Request); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductV1Response); i {
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
		file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_goTypes,
		DependencyIndexes: file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_depIdxs,
		EnumInfos:         file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_enumTypes,
		MessageInfos:      file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_msgTypes,
	}.Build()
	File_alexeykirinyuk_go_product_api_v1_go_product_api_proto = out.File
	file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_rawDesc = nil
	file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_goTypes = nil
	file_alexeykirinyuk_go_product_api_v1_go_product_api_proto_depIdxs = nil
}
