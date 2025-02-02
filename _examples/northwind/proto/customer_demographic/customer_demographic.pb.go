// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: customer_demographic.proto

package customer_demographic

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

type CustomerDemographicByCustomerTypeIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerTypeID string `protobuf:"bytes,1,opt,name=customerTypeID,proto3" json:"customerTypeID,omitempty"`
}

func (x *CustomerDemographicByCustomerTypeIDRequest) Reset() {
	*x = CustomerDemographicByCustomerTypeIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_demographic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerDemographicByCustomerTypeIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerDemographicByCustomerTypeIDRequest) ProtoMessage() {}

func (x *CustomerDemographicByCustomerTypeIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_demographic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerDemographicByCustomerTypeIDRequest.ProtoReflect.Descriptor instead.
func (*CustomerDemographicByCustomerTypeIDRequest) Descriptor() ([]byte, []int) {
	return file_customer_demographic_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerDemographicByCustomerTypeIDRequest) GetCustomerTypeID() string {
	if x != nil {
		return x.CustomerTypeID
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerTypeID string `protobuf:"bytes,1,opt,name=customerTypeID,proto3" json:"customerTypeID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_demographic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_demographic_proto_msgTypes[1]
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
	return file_customer_demographic_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteRequest) GetCustomerTypeID() string {
	if x != nil {
		return x.CustomerTypeID
	}
	return ""
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerDesc   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=customerDesc,proto3" json:"customerDesc,omitempty"`
	CustomerTypeID string                  `protobuf:"bytes,2,opt,name=customerTypeID,proto3" json:"customerTypeID,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_demographic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_demographic_proto_msgTypes[2]
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
	return file_customer_demographic_proto_rawDescGZIP(), []int{2}
}

func (x *InsertRequest) GetCustomerDesc() *wrapperspb.StringValue {
	if x != nil {
		return x.CustomerDesc
	}
	return nil
}

func (x *InsertRequest) GetCustomerTypeID() string {
	if x != nil {
		return x.CustomerTypeID
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerDesc *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=customerDesc,proto3" json:"customerDesc,omitempty"`
	// Output only.
	CustomerTypeID string `protobuf:"bytes,2,opt,name=customerTypeID,proto3" json:"customerTypeID,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_demographic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_demographic_proto_msgTypes[3]
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
	return file_customer_demographic_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetCustomerDesc() *wrapperspb.StringValue {
	if x != nil {
		return x.CustomerDesc
	}
	return nil
}

func (x *UpdateRequest) GetCustomerTypeID() string {
	if x != nil {
		return x.CustomerTypeID
	}
	return ""
}

type UpsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerDesc   *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=customerDesc,proto3" json:"customerDesc,omitempty"`
	CustomerTypeID string                  `protobuf:"bytes,2,opt,name=customerTypeID,proto3" json:"customerTypeID,omitempty"`
}

func (x *UpsertRequest) Reset() {
	*x = UpsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_demographic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRequest) ProtoMessage() {}

func (x *UpsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_demographic_proto_msgTypes[4]
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
	return file_customer_demographic_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertRequest) GetCustomerDesc() *wrapperspb.StringValue {
	if x != nil {
		return x.CustomerDesc
	}
	return nil
}

func (x *UpsertRequest) GetCustomerTypeID() string {
	if x != nil {
		return x.CustomerTypeID
	}
	return ""
}

var File_customer_demographic_proto protoreflect.FileDescriptor

var file_customer_demographic_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x2a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x42, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x22, 0x37, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0x79, 0x0a, 0x0d, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x26,
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0x79, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x22, 0x79, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x32, 0xa6, 0x05, 0x0a,
	0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x69, 0x63, 0x12, 0xb8, 0x01, 0x0a, 0x23, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x42, 0x79, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x40, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6d, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x42, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x7d, 0x12,
	0x78, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x2a, 0x29,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x6d,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x7d, 0x12, 0x6a, 0x0a, 0x06, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64,
	0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x23, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2e, 0x1a, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x71, 0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x23, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x24, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2d,
	0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x2f, 0x75, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x26, 0x5a, 0x24, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x77, 0x69,
	0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_demographic_proto_rawDescOnce sync.Once
	file_customer_demographic_proto_rawDescData = file_customer_demographic_proto_rawDesc
)

func file_customer_demographic_proto_rawDescGZIP() []byte {
	file_customer_demographic_proto_rawDescOnce.Do(func() {
		file_customer_demographic_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_demographic_proto_rawDescData)
	})
	return file_customer_demographic_proto_rawDescData
}

var file_customer_demographic_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_customer_demographic_proto_goTypes = []interface{}{
	(*CustomerDemographicByCustomerTypeIDRequest)(nil), // 0: customer_demographic.CustomerDemographicByCustomerTypeIDRequest
	(*DeleteRequest)(nil),                              // 1: customer_demographic.DeleteRequest
	(*InsertRequest)(nil),                              // 2: customer_demographic.InsertRequest
	(*UpdateRequest)(nil),                              // 3: customer_demographic.UpdateRequest
	(*UpsertRequest)(nil),                              // 4: customer_demographic.UpsertRequest
	(*wrapperspb.StringValue)(nil),                     // 5: google.protobuf.StringValue
	(*typespb.CustomerDemographic)(nil),                // 6: typespb.CustomerDemographic
	(*emptypb.Empty)(nil),                              // 7: google.protobuf.Empty
}
var file_customer_demographic_proto_depIdxs = []int32{
	5, // 0: customer_demographic.InsertRequest.customerDesc:type_name -> google.protobuf.StringValue
	5, // 1: customer_demographic.UpdateRequest.customerDesc:type_name -> google.protobuf.StringValue
	5, // 2: customer_demographic.UpsertRequest.customerDesc:type_name -> google.protobuf.StringValue
	0, // 3: customer_demographic.CustomerDemographic.CustomerDemographicByCustomerTypeID:input_type -> customer_demographic.CustomerDemographicByCustomerTypeIDRequest
	1, // 4: customer_demographic.CustomerDemographic.Delete:input_type -> customer_demographic.DeleteRequest
	2, // 5: customer_demographic.CustomerDemographic.Insert:input_type -> customer_demographic.InsertRequest
	3, // 6: customer_demographic.CustomerDemographic.Update:input_type -> customer_demographic.UpdateRequest
	4, // 7: customer_demographic.CustomerDemographic.Upsert:input_type -> customer_demographic.UpsertRequest
	6, // 8: customer_demographic.CustomerDemographic.CustomerDemographicByCustomerTypeID:output_type -> typespb.CustomerDemographic
	7, // 9: customer_demographic.CustomerDemographic.Delete:output_type -> google.protobuf.Empty
	7, // 10: customer_demographic.CustomerDemographic.Insert:output_type -> google.protobuf.Empty
	7, // 11: customer_demographic.CustomerDemographic.Update:output_type -> google.protobuf.Empty
	7, // 12: customer_demographic.CustomerDemographic.Upsert:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customer_demographic_proto_init() }
func file_customer_demographic_proto_init() {
	if File_customer_demographic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_demographic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerDemographicByCustomerTypeIDRequest); i {
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
		file_customer_demographic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_customer_demographic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_customer_demographic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_customer_demographic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_customer_demographic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_demographic_proto_goTypes,
		DependencyIndexes: file_customer_demographic_proto_depIdxs,
		MessageInfos:      file_customer_demographic_proto_msgTypes,
	}.Build()
	File_customer_demographic_proto = out.File
	file_customer_demographic_proto_rawDesc = nil
	file_customer_demographic_proto_goTypes = nil
	file_customer_demographic_proto_depIdxs = nil
}
