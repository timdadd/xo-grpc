// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: employee_territory.proto

package employee_territory

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
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

	EmployeeID  int32  `protobuf:"varint,1,opt,name=EmployeeID,proto3" json:"EmployeeID,omitempty"`
	TerritoryID string `protobuf:"bytes,2,opt,name=TerritoryID,proto3" json:"TerritoryID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_territory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_territory_proto_msgTypes[0]
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
	return file_employee_territory_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteRequest) GetEmployeeID() int32 {
	if x != nil {
		return x.EmployeeID
	}
	return 0
}

func (x *DeleteRequest) GetTerritoryID() string {
	if x != nil {
		return x.TerritoryID
	}
	return ""
}

type EmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeID int32 `protobuf:"varint,1,opt,name=EmployeeID,proto3" json:"EmployeeID,omitempty"`
}

func (x *EmployeeRequest) Reset() {
	*x = EmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_territory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeRequest) ProtoMessage() {}

func (x *EmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_territory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_territory_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeRequest) GetEmployeeID() int32 {
	if x != nil {
		return x.EmployeeID
	}
	return 0
}

type EmployeeTerritoryByEmployeeIDTerritoryIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeID  int32  `protobuf:"varint,1,opt,name=employeeID,proto3" json:"employeeID,omitempty"`
	TerritoryID string `protobuf:"bytes,2,opt,name=territoryID,proto3" json:"territoryID,omitempty"`
}

func (x *EmployeeTerritoryByEmployeeIDTerritoryIDRequest) Reset() {
	*x = EmployeeTerritoryByEmployeeIDTerritoryIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_territory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeTerritoryByEmployeeIDTerritoryIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeTerritoryByEmployeeIDTerritoryIDRequest) ProtoMessage() {}

func (x *EmployeeTerritoryByEmployeeIDTerritoryIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_territory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeTerritoryByEmployeeIDTerritoryIDRequest.ProtoReflect.Descriptor instead.
func (*EmployeeTerritoryByEmployeeIDTerritoryIDRequest) Descriptor() ([]byte, []int) {
	return file_employee_territory_proto_rawDescGZIP(), []int{2}
}

func (x *EmployeeTerritoryByEmployeeIDTerritoryIDRequest) GetEmployeeID() int32 {
	if x != nil {
		return x.EmployeeID
	}
	return 0
}

func (x *EmployeeTerritoryByEmployeeIDTerritoryIDRequest) GetTerritoryID() string {
	if x != nil {
		return x.TerritoryID
	}
	return ""
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeID  int32  `protobuf:"varint,1,opt,name=EmployeeID,proto3" json:"EmployeeID,omitempty"`
	TerritoryID string `protobuf:"bytes,2,opt,name=TerritoryID,proto3" json:"TerritoryID,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_territory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_territory_proto_msgTypes[3]
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
	return file_employee_territory_proto_rawDescGZIP(), []int{3}
}

func (x *InsertRequest) GetEmployeeID() int32 {
	if x != nil {
		return x.EmployeeID
	}
	return 0
}

func (x *InsertRequest) GetTerritoryID() string {
	if x != nil {
		return x.TerritoryID
	}
	return ""
}

type TerritoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerritoryID string `protobuf:"bytes,1,opt,name=TerritoryID,proto3" json:"TerritoryID,omitempty"`
}

func (x *TerritoryRequest) Reset() {
	*x = TerritoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_territory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerritoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerritoryRequest) ProtoMessage() {}

func (x *TerritoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_territory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerritoryRequest.ProtoReflect.Descriptor instead.
func (*TerritoryRequest) Descriptor() ([]byte, []int) {
	return file_employee_territory_proto_rawDescGZIP(), []int{4}
}

func (x *TerritoryRequest) GetTerritoryID() string {
	if x != nil {
		return x.TerritoryID
	}
	return ""
}

var File_employee_territory_proto protoreflect.FileDescriptor

var file_employee_territory_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x31, 0x0a, 0x0f,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x22,
	0x73, 0x0a, 0x2f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44,
	0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x22, 0x51, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x32, 0x8e, 0x05,
	0x0a, 0x18, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f,
	0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x78, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2f, 0x7b, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x7d, 0x12, 0xab, 0x01, 0x0a, 0x28, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x42,
	0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x54, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x43, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x5f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12,
	0x16, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x74, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x66, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x12, 0x21, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x72,
	0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2d, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12,
	0x7d, 0x0a, 0x09, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x72,
	0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x74, 0x65, 0x72,
	0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x7b, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x7d, 0x42, 0xf5,
	0x02, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61,
	0x6c, 0x74, 0x65, 0x72, 0x77, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x65, 0x79, 0x2f, 0x78, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x6e, 0x6f, 0x72, 0x74, 0x68, 0x77, 0x69, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x92, 0x41, 0xa0, 0x02, 0x12, 0x9d, 0x02, 0x0a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x77, 0x61, 0x6e, 0x64,
	0x65, 0x72, 0x6c, 0x65, 0x79, 0x2f, 0x78, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x5f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x77, 0x69, 0x6e,
	0x64, 0x12, 0x93, 0x01, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20,
	0x63, 0x6f, 0x64, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x62,
	0x79, 0x20, 0x78, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x20, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x20,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x6e, 0x20, 0x72, 0x75, 0x6e, 0x20, 0x27,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x27, 0x20, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x20, 0x6f, 0x6e, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x20, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x22, 0x48, 0x0a, 0x1a, 0x78, 0x6f, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x20, 0x2d, 0x20, 0x57, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x57, 0x61, 0x6e, 0x64,
	0x65, 0x72, 0x6c, 0x65, 0x79, 0x12, 0x2a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6c, 0x74, 0x65, 0x72,
	0x77, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x65, 0x79, 0x2f, 0x78, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_employee_territory_proto_rawDescOnce sync.Once
	file_employee_territory_proto_rawDescData = file_employee_territory_proto_rawDesc
)

func file_employee_territory_proto_rawDescGZIP() []byte {
	file_employee_territory_proto_rawDescOnce.Do(func() {
		file_employee_territory_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_territory_proto_rawDescData)
	})
	return file_employee_territory_proto_rawDescData
}

var file_employee_territory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_employee_territory_proto_goTypes = []interface{}{
	(*DeleteRequest)(nil),   // 0: employee_territory.DeleteRequest
	(*EmployeeRequest)(nil), // 1: employee_territory.EmployeeRequest
	(*EmployeeTerritoryByEmployeeIDTerritoryIDRequest)(nil), // 2: employee_territory.EmployeeTerritoryByEmployeeIDTerritoryIDRequest
	(*InsertRequest)(nil),             // 3: employee_territory.InsertRequest
	(*TerritoryRequest)(nil),          // 4: employee_territory.TerritoryRequest
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
	(*typespb.Employee)(nil),          // 6: typespb.Employee
	(*typespb.EmployeeTerritory)(nil), // 7: typespb.EmployeeTerritory
	(*typespb.Territory)(nil),         // 8: typespb.Territory
}
var file_employee_territory_proto_depIdxs = []int32{
	0, // 0: employee_territory.EmployeeTerritoryService.Delete:input_type -> employee_territory.DeleteRequest
	1, // 1: employee_territory.EmployeeTerritoryService.Employee:input_type -> employee_territory.EmployeeRequest
	2, // 2: employee_territory.EmployeeTerritoryService.EmployeeTerritoryByEmployeeIDTerritoryID:input_type -> employee_territory.EmployeeTerritoryByEmployeeIDTerritoryIDRequest
	3, // 3: employee_territory.EmployeeTerritoryService.Insert:input_type -> employee_territory.InsertRequest
	4, // 4: employee_territory.EmployeeTerritoryService.Territory:input_type -> employee_territory.TerritoryRequest
	5, // 5: employee_territory.EmployeeTerritoryService.Delete:output_type -> google.protobuf.Empty
	6, // 6: employee_territory.EmployeeTerritoryService.Employee:output_type -> typespb.Employee
	7, // 7: employee_territory.EmployeeTerritoryService.EmployeeTerritoryByEmployeeIDTerritoryID:output_type -> typespb.EmployeeTerritory
	5, // 8: employee_territory.EmployeeTerritoryService.Insert:output_type -> google.protobuf.Empty
	8, // 9: employee_territory.EmployeeTerritoryService.Territory:output_type -> typespb.Territory
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_employee_territory_proto_init() }
func file_employee_territory_proto_init() {
	if File_employee_territory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_territory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_employee_territory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeRequest); i {
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
		file_employee_territory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeTerritoryByEmployeeIDTerritoryIDRequest); i {
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
		file_employee_territory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_employee_territory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerritoryRequest); i {
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
			RawDescriptor: file_employee_territory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employee_territory_proto_goTypes,
		DependencyIndexes: file_employee_territory_proto_depIdxs,
		MessageInfos:      file_employee_territory_proto_msgTypes,
	}.Build()
	File_employee_territory_proto = out.File
	file_employee_territory_proto_rawDesc = nil
	file_employee_territory_proto_goTypes = nil
	file_employee_territory_proto_depIdxs = nil
}
