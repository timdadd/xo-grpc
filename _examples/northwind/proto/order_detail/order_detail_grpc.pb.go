// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package order_detail

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	typespb "northwind/proto/typespb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderDetailClient is the client API for OrderDetail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderDetailClient interface {
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*typespb.Order, error)
	OrderDetailByOrderIDProductID(ctx context.Context, in *OrderDetailByOrderIDProductIDRequest, opts ...grpc.CallOption) (*typespb.OrderDetail, error)
	Product(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*typespb.Product, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderDetailClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderDetailClient(cc grpc.ClientConnInterface) OrderDetailClient {
	return &orderDetailClient{cc}
}

func (c *orderDetailClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/order_detail.OrderDetail/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderDetailClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/order_detail.OrderDetail/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderDetailClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*typespb.Order, error) {
	out := new(typespb.Order)
	err := c.cc.Invoke(ctx, "/order_detail.OrderDetail/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderDetailClient) OrderDetailByOrderIDProductID(ctx context.Context, in *OrderDetailByOrderIDProductIDRequest, opts ...grpc.CallOption) (*typespb.OrderDetail, error) {
	out := new(typespb.OrderDetail)
	err := c.cc.Invoke(ctx, "/order_detail.OrderDetail/OrderDetailByOrderIDProductID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderDetailClient) Product(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*typespb.Product, error) {
	out := new(typespb.Product)
	err := c.cc.Invoke(ctx, "/order_detail.OrderDetail/Product", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderDetailClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/order_detail.OrderDetail/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderDetailClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/order_detail.OrderDetail/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderDetailServer is the server API for OrderDetail service.
// All implementations must embed UnimplementedOrderDetailServer
// for forward compatibility
type OrderDetailServer interface {
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	Insert(context.Context, *InsertRequest) (*emptypb.Empty, error)
	Order(context.Context, *OrderRequest) (*typespb.Order, error)
	OrderDetailByOrderIDProductID(context.Context, *OrderDetailByOrderIDProductIDRequest) (*typespb.OrderDetail, error)
	Product(context.Context, *ProductRequest) (*typespb.Product, error)
	Update(context.Context, *UpdateRequest) (*emptypb.Empty, error)
	Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrderDetailServer()
}

// UnimplementedOrderDetailServer must be embedded to have forward compatible implementations.
type UnimplementedOrderDetailServer struct {
}

func (UnimplementedOrderDetailServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOrderDetailServer) Insert(context.Context, *InsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedOrderDetailServer) Order(context.Context, *OrderRequest) (*typespb.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}
func (UnimplementedOrderDetailServer) OrderDetailByOrderIDProductID(context.Context, *OrderDetailByOrderIDProductIDRequest) (*typespb.OrderDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetailByOrderIDProductID not implemented")
}
func (UnimplementedOrderDetailServer) Product(context.Context, *ProductRequest) (*typespb.Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Product not implemented")
}
func (UnimplementedOrderDetailServer) Update(context.Context, *UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderDetailServer) Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedOrderDetailServer) mustEmbedUnimplementedOrderDetailServer() {}

// UnsafeOrderDetailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderDetailServer will
// result in compilation errors.
type UnsafeOrderDetailServer interface {
	mustEmbedUnimplementedOrderDetailServer()
}

func RegisterOrderDetailServer(s grpc.ServiceRegistrar, srv OrderDetailServer) {
	s.RegisterService(&OrderDetail_ServiceDesc, srv)
}

func _OrderDetail_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderDetailServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_detail.OrderDetail/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderDetailServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderDetail_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderDetailServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_detail.OrderDetail/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderDetailServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderDetail_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderDetailServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_detail.OrderDetail/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderDetailServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderDetail_OrderDetailByOrderIDProductID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDetailByOrderIDProductIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderDetailServer).OrderDetailByOrderIDProductID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_detail.OrderDetail/OrderDetailByOrderIDProductID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderDetailServer).OrderDetailByOrderIDProductID(ctx, req.(*OrderDetailByOrderIDProductIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderDetail_Product_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderDetailServer).Product(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_detail.OrderDetail/Product",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderDetailServer).Product(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderDetail_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderDetailServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_detail.OrderDetail/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderDetailServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderDetail_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderDetailServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_detail.OrderDetail/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderDetailServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderDetail_ServiceDesc is the grpc.ServiceDesc for OrderDetail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderDetail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_detail.OrderDetail",
	HandlerType: (*OrderDetailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _OrderDetail_Delete_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _OrderDetail_Insert_Handler,
		},
		{
			MethodName: "Order",
			Handler:    _OrderDetail_Order_Handler,
		},
		{
			MethodName: "OrderDetailByOrderIDProductID",
			Handler:    _OrderDetail_OrderDetailByOrderIDProductID_Handler,
		},
		{
			MethodName: "Product",
			Handler:    _OrderDetail_Product_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderDetail_Update_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _OrderDetail_Upsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_detail.proto",
}
