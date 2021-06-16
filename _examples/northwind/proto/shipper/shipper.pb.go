// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: shipper.proto

package shipper

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

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperID int32 `protobuf:"varint,1,opt,name=shipperID,proto3" json:"shipperID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipper_proto_msgTypes[0]
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
	return file_shipper_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteRequest) GetShipperID() int32 {
	if x != nil {
		return x.ShipperID
	}
	return 0
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName string                  `protobuf:"bytes,1,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Phone       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	ShipperID   int32                   `protobuf:"varint,3,opt,name=shipperID,proto3" json:"shipperID,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipper_proto_msgTypes[1]
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
	return file_shipper_proto_rawDescGZIP(), []int{1}
}

func (x *InsertRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *InsertRequest) GetPhone() *wrapperspb.StringValue {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *InsertRequest) GetShipperID() int32 {
	if x != nil {
		return x.ShipperID
	}
	return 0
}

type ShipperByShipperIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperID int32 `protobuf:"varint,1,opt,name=shipperID,proto3" json:"shipperID,omitempty"`
}

func (x *ShipperByShipperIDRequest) Reset() {
	*x = ShipperByShipperIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipperByShipperIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipperByShipperIDRequest) ProtoMessage() {}

func (x *ShipperByShipperIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipperByShipperIDRequest.ProtoReflect.Descriptor instead.
func (*ShipperByShipperIDRequest) Descriptor() ([]byte, []int) {
	return file_shipper_proto_rawDescGZIP(), []int{2}
}

func (x *ShipperByShipperIDRequest) GetShipperID() int32 {
	if x != nil {
		return x.ShipperID
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName string                  `protobuf:"bytes,1,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Phone       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	// Output only.
	ShipperID int32 `protobuf:"varint,3,opt,name=shipperID,proto3" json:"shipperID,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipper_proto_msgTypes[3]
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
	return file_shipper_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *UpdateRequest) GetPhone() *wrapperspb.StringValue {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *UpdateRequest) GetShipperID() int32 {
	if x != nil {
		return x.ShipperID
	}
	return 0
}

type UpsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName string                  `protobuf:"bytes,1,opt,name=companyName,proto3" json:"companyName,omitempty"`
	Phone       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	ShipperID   int32                   `protobuf:"varint,3,opt,name=shipperID,proto3" json:"shipperID,omitempty"`
}

func (x *UpsertRequest) Reset() {
	*x = UpsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRequest) ProtoMessage() {}

func (x *UpsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipper_proto_msgTypes[4]
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
	return file_shipper_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *UpsertRequest) GetPhone() *wrapperspb.StringValue {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *UpsertRequest) GetShipperID() int32 {
	if x != nil {
		return x.ShipperID
	}
	return 0
}

var File_shipper_proto protoreflect.FileDescriptor

var file_shipper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x19, 0x53, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x79, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x32, 0xda,
	0x03, 0x0a, 0x07, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x49, 0x44, 0x7d, 0x12, 0x50, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12,
	0x16, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x6b, 0x0a, 0x12, 0x53, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x79, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x2e,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x42,
	0x79, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x49, 0x44, 0x7d, 0x12, 0x5c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x44, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x57, 0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x2f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x19, 0x5a, 0x17, 0x6e,
	0x6f, 0x72, 0x74, 0x68, 0x77, 0x69, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipper_proto_rawDescOnce sync.Once
	file_shipper_proto_rawDescData = file_shipper_proto_rawDesc
)

func file_shipper_proto_rawDescGZIP() []byte {
	file_shipper_proto_rawDescOnce.Do(func() {
		file_shipper_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipper_proto_rawDescData)
	})
	return file_shipper_proto_rawDescData
}

var file_shipper_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shipper_proto_goTypes = []interface{}{
	(*DeleteRequest)(nil),             // 0: shipper.DeleteRequest
	(*InsertRequest)(nil),             // 1: shipper.InsertRequest
	(*ShipperByShipperIDRequest)(nil), // 2: shipper.ShipperByShipperIDRequest
	(*UpdateRequest)(nil),             // 3: shipper.UpdateRequest
	(*UpsertRequest)(nil),             // 4: shipper.UpsertRequest
	(*wrapperspb.StringValue)(nil),    // 5: google.protobuf.StringValue
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
	(*typespb.Shipper)(nil),           // 7: typespb.Shipper
}
var file_shipper_proto_depIdxs = []int32{
	5, // 0: shipper.InsertRequest.phone:type_name -> google.protobuf.StringValue
	5, // 1: shipper.UpdateRequest.phone:type_name -> google.protobuf.StringValue
	5, // 2: shipper.UpsertRequest.phone:type_name -> google.protobuf.StringValue
	0, // 3: shipper.Shipper.Delete:input_type -> shipper.DeleteRequest
	1, // 4: shipper.Shipper.Insert:input_type -> shipper.InsertRequest
	2, // 5: shipper.Shipper.ShipperByShipperID:input_type -> shipper.ShipperByShipperIDRequest
	3, // 6: shipper.Shipper.Update:input_type -> shipper.UpdateRequest
	4, // 7: shipper.Shipper.Upsert:input_type -> shipper.UpsertRequest
	6, // 8: shipper.Shipper.Delete:output_type -> google.protobuf.Empty
	6, // 9: shipper.Shipper.Insert:output_type -> google.protobuf.Empty
	7, // 10: shipper.Shipper.ShipperByShipperID:output_type -> typespb.Shipper
	6, // 11: shipper.Shipper.Update:output_type -> google.protobuf.Empty
	6, // 12: shipper.Shipper.Upsert:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shipper_proto_init() }
func file_shipper_proto_init() {
	if File_shipper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_shipper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_shipper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipperByShipperIDRequest); i {
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
		file_shipper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_shipper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_shipper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipper_proto_goTypes,
		DependencyIndexes: file_shipper_proto_depIdxs,
		MessageInfos:      file_shipper_proto_msgTypes,
	}.Build()
	File_shipper_proto = out.File
	file_shipper_proto_rawDesc = nil
	file_shipper_proto_goTypes = nil
	file_shipper_proto_depIdxs = nil
}
