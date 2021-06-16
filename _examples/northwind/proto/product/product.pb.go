// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: product.proto

package product

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	typespb "northwind/proto/typespb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int32 `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *CategoryRequest) Reset() {
	*x = CategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryRequest) ProtoMessage() {}

func (x *CategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryRequest.ProtoReflect.Descriptor instead.
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int32 `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID      *wrapperspb.Int64Value  `protobuf:"bytes,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
	Discontinued    int64                   `protobuf:"varint,2,opt,name=discontinued,proto3" json:"discontinued,omitempty"`
	ProductID       int32                   `protobuf:"varint,3,opt,name=productID,proto3" json:"productID,omitempty"`
	ProductName     string                  `protobuf:"bytes,4,opt,name=productName,proto3" json:"productName,omitempty"`
	QuantityPerUnit *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=quantityPerUnit,proto3" json:"quantityPerUnit,omitempty"`
	ReorderLevel    *wrapperspb.Int64Value  `protobuf:"bytes,6,opt,name=reorderLevel,proto3" json:"reorderLevel,omitempty"`
	SupplierID      *wrapperspb.Int64Value  `protobuf:"bytes,7,opt,name=supplierID,proto3" json:"supplierID,omitempty"`
	UnitPrice       *wrapperspb.DoubleValue `protobuf:"bytes,8,opt,name=unitPrice,proto3" json:"unitPrice,omitempty"`
	UnitsInStock    *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=unitsInStock,proto3" json:"unitsInStock,omitempty"`
	UnitsOnOrder    *wrapperspb.Int64Value  `protobuf:"bytes,10,opt,name=unitsOnOrder,proto3" json:"unitsOnOrder,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *InsertRequest) GetCategoryID() *wrapperspb.Int64Value {
	if x != nil {
		return x.CategoryID
	}
	return nil
}

func (x *InsertRequest) GetDiscontinued() int64 {
	if x != nil {
		return x.Discontinued
	}
	return 0
}

func (x *InsertRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *InsertRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *InsertRequest) GetQuantityPerUnit() *wrapperspb.StringValue {
	if x != nil {
		return x.QuantityPerUnit
	}
	return nil
}

func (x *InsertRequest) GetReorderLevel() *wrapperspb.Int64Value {
	if x != nil {
		return x.ReorderLevel
	}
	return nil
}

func (x *InsertRequest) GetSupplierID() *wrapperspb.Int64Value {
	if x != nil {
		return x.SupplierID
	}
	return nil
}

func (x *InsertRequest) GetUnitPrice() *wrapperspb.DoubleValue {
	if x != nil {
		return x.UnitPrice
	}
	return nil
}

func (x *InsertRequest) GetUnitsInStock() *wrapperspb.Int64Value {
	if x != nil {
		return x.UnitsInStock
	}
	return nil
}

func (x *InsertRequest) GetUnitsOnOrder() *wrapperspb.Int64Value {
	if x != nil {
		return x.UnitsOnOrder
	}
	return nil
}

type ProductByProductIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int32 `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *ProductByProductIDRequest) Reset() {
	*x = ProductByProductIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductByProductIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductByProductIDRequest) ProtoMessage() {}

func (x *ProductByProductIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductByProductIDRequest.ProtoReflect.Descriptor instead.
func (*ProductByProductIDRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductByProductIDRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type SupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int32 `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *SupplierRequest) Reset() {
	*x = SupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplierRequest) ProtoMessage() {}

func (x *SupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplierRequest.ProtoReflect.Descriptor instead.
func (*SupplierRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *SupplierRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID   *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
	Discontinued int64                  `protobuf:"varint,2,opt,name=discontinued,proto3" json:"discontinued,omitempty"`
	// Output only.
	ProductID       int32                   `protobuf:"varint,3,opt,name=productID,proto3" json:"productID,omitempty"`
	ProductName     string                  `protobuf:"bytes,4,opt,name=productName,proto3" json:"productName,omitempty"`
	QuantityPerUnit *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=quantityPerUnit,proto3" json:"quantityPerUnit,omitempty"`
	ReorderLevel    *wrapperspb.Int64Value  `protobuf:"bytes,6,opt,name=reorderLevel,proto3" json:"reorderLevel,omitempty"`
	SupplierID      *wrapperspb.Int64Value  `protobuf:"bytes,7,opt,name=supplierID,proto3" json:"supplierID,omitempty"`
	UnitPrice       *wrapperspb.DoubleValue `protobuf:"bytes,8,opt,name=unitPrice,proto3" json:"unitPrice,omitempty"`
	UnitsInStock    *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=unitsInStock,proto3" json:"unitsInStock,omitempty"`
	UnitsOnOrder    *wrapperspb.Int64Value  `protobuf:"bytes,10,opt,name=unitsOnOrder,proto3" json:"unitsOnOrder,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetCategoryID() *wrapperspb.Int64Value {
	if x != nil {
		return x.CategoryID
	}
	return nil
}

func (x *UpdateRequest) GetDiscontinued() int64 {
	if x != nil {
		return x.Discontinued
	}
	return 0
}

func (x *UpdateRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *UpdateRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *UpdateRequest) GetQuantityPerUnit() *wrapperspb.StringValue {
	if x != nil {
		return x.QuantityPerUnit
	}
	return nil
}

func (x *UpdateRequest) GetReorderLevel() *wrapperspb.Int64Value {
	if x != nil {
		return x.ReorderLevel
	}
	return nil
}

func (x *UpdateRequest) GetSupplierID() *wrapperspb.Int64Value {
	if x != nil {
		return x.SupplierID
	}
	return nil
}

func (x *UpdateRequest) GetUnitPrice() *wrapperspb.DoubleValue {
	if x != nil {
		return x.UnitPrice
	}
	return nil
}

func (x *UpdateRequest) GetUnitsInStock() *wrapperspb.Int64Value {
	if x != nil {
		return x.UnitsInStock
	}
	return nil
}

func (x *UpdateRequest) GetUnitsOnOrder() *wrapperspb.Int64Value {
	if x != nil {
		return x.UnitsOnOrder
	}
	return nil
}

type UpsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryID      *wrapperspb.Int64Value  `protobuf:"bytes,1,opt,name=categoryID,proto3" json:"categoryID,omitempty"`
	Discontinued    int64                   `protobuf:"varint,2,opt,name=discontinued,proto3" json:"discontinued,omitempty"`
	ProductID       int32                   `protobuf:"varint,3,opt,name=productID,proto3" json:"productID,omitempty"`
	ProductName     string                  `protobuf:"bytes,4,opt,name=productName,proto3" json:"productName,omitempty"`
	QuantityPerUnit *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=quantityPerUnit,proto3" json:"quantityPerUnit,omitempty"`
	ReorderLevel    *wrapperspb.Int64Value  `protobuf:"bytes,6,opt,name=reorderLevel,proto3" json:"reorderLevel,omitempty"`
	SupplierID      *wrapperspb.Int64Value  `protobuf:"bytes,7,opt,name=supplierID,proto3" json:"supplierID,omitempty"`
	UnitPrice       *wrapperspb.DoubleValue `protobuf:"bytes,8,opt,name=unitPrice,proto3" json:"unitPrice,omitempty"`
	UnitsInStock    *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=unitsInStock,proto3" json:"unitsInStock,omitempty"`
	UnitsOnOrder    *wrapperspb.Int64Value  `protobuf:"bytes,10,opt,name=unitsOnOrder,proto3" json:"unitsOnOrder,omitempty"`
}

func (x *UpsertRequest) Reset() {
	*x = UpsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRequest) ProtoMessage() {}

func (x *UpsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertRequest.ProtoReflect.Descriptor instead.
func (*UpsertRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{6}
}

func (x *UpsertRequest) GetCategoryID() *wrapperspb.Int64Value {
	if x != nil {
		return x.CategoryID
	}
	return nil
}

func (x *UpsertRequest) GetDiscontinued() int64 {
	if x != nil {
		return x.Discontinued
	}
	return 0
}

func (x *UpsertRequest) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *UpsertRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *UpsertRequest) GetQuantityPerUnit() *wrapperspb.StringValue {
	if x != nil {
		return x.QuantityPerUnit
	}
	return nil
}

func (x *UpsertRequest) GetReorderLevel() *wrapperspb.Int64Value {
	if x != nil {
		return x.ReorderLevel
	}
	return nil
}

func (x *UpsertRequest) GetSupplierID() *wrapperspb.Int64Value {
	if x != nil {
		return x.SupplierID
	}
	return nil
}

func (x *UpsertRequest) GetUnitPrice() *wrapperspb.DoubleValue {
	if x != nil {
		return x.UnitPrice
	}
	return nil
}

func (x *UpsertRequest) GetUnitsInStock() *wrapperspb.Int64Value {
	if x != nil {
		return x.UnitsInStock
	}
	return nil
}

func (x *UpsertRequest) GetUnitsOnOrder() *wrapperspb.Int64Value {
	if x != nil {
		return x.UnitsOnOrder
	}
	return nil
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x22, 0xb4, 0x04, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x74,
	0x69, 0x6e, 0x75, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x3f, 0x0a,
	0x0c, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3b,
	0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x09, 0x75,
	0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x75, 0x6e,
	0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x4f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x75, 0x6e, 0x69,
	0x74, 0x73, 0x4f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x19, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x22, 0x2f, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0xb4, 0x04, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x6e, 0x75, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x3f, 0x0a, 0x0c, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3a,
	0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x75,
	0x6e, 0x69, 0x74, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x0c, 0x75,
	0x6e, 0x69, 0x74, 0x73, 0x4f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x4f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xb4, 0x04, 0x0a,
	0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x46, 0x0a, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x55, 0x6e,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x72, 0x65, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x49, 0x6e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x4f, 0x6e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x4f, 0x6e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x32, 0xa0, 0x05, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x61, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x12, 0x50, 0x0a,
	0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22,
	0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0x6b, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x12, 0x61, 0x0a, 0x08,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12,
	0x5c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a,
	0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x75, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x19, 0x5a, 0x17, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x77,
	0x69, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_product_proto_goTypes = []interface{}{
	(*CategoryRequest)(nil),           // 0: product.CategoryRequest
	(*DeleteRequest)(nil),             // 1: product.DeleteRequest
	(*InsertRequest)(nil),             // 2: product.InsertRequest
	(*ProductByProductIDRequest)(nil), // 3: product.ProductByProductIDRequest
	(*SupplierRequest)(nil),           // 4: product.SupplierRequest
	(*UpdateRequest)(nil),             // 5: product.UpdateRequest
	(*UpsertRequest)(nil),             // 6: product.UpsertRequest
	(*wrapperspb.Int64Value)(nil),     // 7: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil),    // 8: google.protobuf.StringValue
	(*wrapperspb.DoubleValue)(nil),    // 9: google.protobuf.DoubleValue
	(*typespb.Category)(nil),          // 10: typespb.Category
	(*emptypb.Empty)(nil),             // 11: google.protobuf.Empty
	(*typespb.Product)(nil),           // 12: typespb.Product
	(*typespb.Supplier)(nil),          // 13: typespb.Supplier
}
var file_product_proto_depIdxs = []int32{
	7,  // 0: product.InsertRequest.categoryID:type_name -> google.protobuf.Int64Value
	8,  // 1: product.InsertRequest.quantityPerUnit:type_name -> google.protobuf.StringValue
	7,  // 2: product.InsertRequest.reorderLevel:type_name -> google.protobuf.Int64Value
	7,  // 3: product.InsertRequest.supplierID:type_name -> google.protobuf.Int64Value
	9,  // 4: product.InsertRequest.unitPrice:type_name -> google.protobuf.DoubleValue
	7,  // 5: product.InsertRequest.unitsInStock:type_name -> google.protobuf.Int64Value
	7,  // 6: product.InsertRequest.unitsOnOrder:type_name -> google.protobuf.Int64Value
	7,  // 7: product.UpdateRequest.categoryID:type_name -> google.protobuf.Int64Value
	8,  // 8: product.UpdateRequest.quantityPerUnit:type_name -> google.protobuf.StringValue
	7,  // 9: product.UpdateRequest.reorderLevel:type_name -> google.protobuf.Int64Value
	7,  // 10: product.UpdateRequest.supplierID:type_name -> google.protobuf.Int64Value
	9,  // 11: product.UpdateRequest.unitPrice:type_name -> google.protobuf.DoubleValue
	7,  // 12: product.UpdateRequest.unitsInStock:type_name -> google.protobuf.Int64Value
	7,  // 13: product.UpdateRequest.unitsOnOrder:type_name -> google.protobuf.Int64Value
	7,  // 14: product.UpsertRequest.categoryID:type_name -> google.protobuf.Int64Value
	8,  // 15: product.UpsertRequest.quantityPerUnit:type_name -> google.protobuf.StringValue
	7,  // 16: product.UpsertRequest.reorderLevel:type_name -> google.protobuf.Int64Value
	7,  // 17: product.UpsertRequest.supplierID:type_name -> google.protobuf.Int64Value
	9,  // 18: product.UpsertRequest.unitPrice:type_name -> google.protobuf.DoubleValue
	7,  // 19: product.UpsertRequest.unitsInStock:type_name -> google.protobuf.Int64Value
	7,  // 20: product.UpsertRequest.unitsOnOrder:type_name -> google.protobuf.Int64Value
	0,  // 21: product.Product.Category:input_type -> product.CategoryRequest
	1,  // 22: product.Product.Delete:input_type -> product.DeleteRequest
	2,  // 23: product.Product.Insert:input_type -> product.InsertRequest
	3,  // 24: product.Product.ProductByProductID:input_type -> product.ProductByProductIDRequest
	4,  // 25: product.Product.Supplier:input_type -> product.SupplierRequest
	5,  // 26: product.Product.Update:input_type -> product.UpdateRequest
	6,  // 27: product.Product.Upsert:input_type -> product.UpsertRequest
	10, // 28: product.Product.Category:output_type -> typespb.Category
	11, // 29: product.Product.Delete:output_type -> google.protobuf.Empty
	11, // 30: product.Product.Insert:output_type -> google.protobuf.Empty
	12, // 31: product.Product.ProductByProductID:output_type -> typespb.Product
	13, // 32: product.Product.Supplier:output_type -> typespb.Supplier
	11, // 33: product.Product.Update:output_type -> google.protobuf.Empty
	11, // 34: product.Product.Upsert:output_type -> google.protobuf.Empty
	28, // [28:35] is the sub-list for method output_type
	21, // [21:28] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryRequest); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductByProductIDRequest); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplierRequest); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertRequest); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
