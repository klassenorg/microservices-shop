// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	GetAllProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Products, error)
	GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*Product, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) GetAllProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/microservicesshop.CatalogService/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/microservicesshop.CatalogService/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations must embed UnimplementedCatalogServiceServer
// for forward compatibility
type CatalogServiceServer interface {
	GetAllProducts(context.Context, *Empty) (*Products, error)
	GetProductByID(context.Context, *GetProductByIDRequest) (*Product, error)
	mustEmbedUnimplementedCatalogServiceServer()
}

// UnimplementedCatalogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (UnimplementedCatalogServiceServer) GetAllProducts(context.Context, *Empty) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedCatalogServiceServer) GetProductByID(context.Context, *GetProductByIDRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedCatalogServiceServer) mustEmbedUnimplementedCatalogServiceServer() {}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.CatalogService/GetAllProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetAllProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.CatalogService/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProductByID(ctx, req.(*GetProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microservicesshop.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllProducts",
			Handler:    _CatalogService_GetAllProducts_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _CatalogService_GetProductByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

// RecommendationServiceClient is the client API for RecommendationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendationServiceClient interface {
	GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*Products, error)
}

type recommendationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationServiceClient(cc grpc.ClientConnInterface) RecommendationServiceClient {
	return &recommendationServiceClient{cc}
}

func (c *recommendationServiceClient) GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/microservicesshop.RecommendationService/GetRecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceServer is the server API for RecommendationService service.
// All implementations must embed UnimplementedRecommendationServiceServer
// for forward compatibility
type RecommendationServiceServer interface {
	GetRecommendations(context.Context, *GetRecommendationsRequest) (*Products, error)
	mustEmbedUnimplementedRecommendationServiceServer()
}

// UnimplementedRecommendationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecommendationServiceServer struct {
}

func (UnimplementedRecommendationServiceServer) GetRecommendations(context.Context, *GetRecommendationsRequest) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendations not implemented")
}
func (UnimplementedRecommendationServiceServer) mustEmbedUnimplementedRecommendationServiceServer() {}

// UnsafeRecommendationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendationServiceServer will
// result in compilation errors.
type UnsafeRecommendationServiceServer interface {
	mustEmbedUnimplementedRecommendationServiceServer()
}

func RegisterRecommendationServiceServer(s grpc.ServiceRegistrar, srv RecommendationServiceServer) {
	s.RegisterService(&RecommendationService_ServiceDesc, srv)
}

func _RecommendationService_GetRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).GetRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.RecommendationService/GetRecommendations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).GetRecommendations(ctx, req.(*GetRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendationService_ServiceDesc is the grpc.ServiceDesc for RecommendationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microservicesshop.RecommendationService",
	HandlerType: (*RecommendationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecommendations",
			Handler:    _RecommendationService_GetRecommendations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	AddToCart(ctx context.Context, in *CartUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveFromCart(ctx context.Context, in *CartUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveAllFromCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*Empty, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, "/microservicesshop.CartService/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddToCart(ctx context.Context, in *CartUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/microservicesshop.CartService/AddToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveFromCart(ctx context.Context, in *CartUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/microservicesshop.CartService/RemoveFromCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveAllFromCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/microservicesshop.CartService/RemoveAllFromCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	GetCart(context.Context, *CartRequest) (*CartResponse, error)
	AddToCart(context.Context, *CartUpdateRequest) (*Empty, error)
	RemoveFromCart(context.Context, *CartUpdateRequest) (*Empty, error)
	RemoveAllFromCart(context.Context, *CartRequest) (*Empty, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) GetCart(context.Context, *CartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) AddToCart(context.Context, *CartUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveFromCart(context.Context, *CartUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveAllFromCart(context.Context, *CartRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAllFromCart not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.CartService/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.CartService/AddToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddToCart(ctx, req.(*CartUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.CartService/RemoveFromCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveFromCart(ctx, req.(*CartUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveAllFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveAllFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.CartService/RemoveAllFromCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveAllFromCart(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microservicesshop.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "AddToCart",
			Handler:    _CartService_AddToCart_Handler,
		},
		{
			MethodName: "RemoveFromCart",
			Handler:    _CartService_RemoveFromCart_Handler,
		},
		{
			MethodName: "RemoveAllFromCart",
			Handler:    _CartService_RemoveAllFromCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

// PricingServiceClient is the client API for PricingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PricingServiceClient interface {
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type pricingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPricingServiceClient(cc grpc.ClientConnInterface) PricingServiceClient {
	return &pricingServiceClient{cc}
}

func (c *pricingServiceClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/microservicesshop.PricingService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricingServiceServer is the server API for PricingService service.
// All implementations must embed UnimplementedPricingServiceServer
// for forward compatibility
type PricingServiceServer interface {
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	mustEmbedUnimplementedPricingServiceServer()
}

// UnimplementedPricingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPricingServiceServer struct {
}

func (UnimplementedPricingServiceServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedPricingServiceServer) mustEmbedUnimplementedPricingServiceServer() {}

// UnsafePricingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PricingServiceServer will
// result in compilation errors.
type UnsafePricingServiceServer interface {
	mustEmbedUnimplementedPricingServiceServer()
}

func RegisterPricingServiceServer(s grpc.ServiceRegistrar, srv PricingServiceServer) {
	s.RegisterService(&PricingService_ServiceDesc, srv)
}

func _PricingService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.PricingService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PricingService_ServiceDesc is the grpc.ServiceDesc for PricingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PricingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microservicesshop.PricingService",
	HandlerType: (*PricingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _PricingService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}

// PurchaseServiceClient is the client API for PurchaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PurchaseServiceClient interface {
	Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
}

type purchaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPurchaseServiceClient(cc grpc.ClientConnInterface) PurchaseServiceClient {
	return &purchaseServiceClient{cc}
}

func (c *purchaseServiceClient) Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, "/microservicesshop.PurchaseService/Purchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PurchaseServiceServer is the server API for PurchaseService service.
// All implementations must embed UnimplementedPurchaseServiceServer
// for forward compatibility
type PurchaseServiceServer interface {
	Purchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
	mustEmbedUnimplementedPurchaseServiceServer()
}

// UnimplementedPurchaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPurchaseServiceServer struct {
}

func (UnimplementedPurchaseServiceServer) Purchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purchase not implemented")
}
func (UnimplementedPurchaseServiceServer) mustEmbedUnimplementedPurchaseServiceServer() {}

// UnsafePurchaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PurchaseServiceServer will
// result in compilation errors.
type UnsafePurchaseServiceServer interface {
	mustEmbedUnimplementedPurchaseServiceServer()
}

func RegisterPurchaseServiceServer(s grpc.ServiceRegistrar, srv PurchaseServiceServer) {
	s.RegisterService(&PurchaseService_ServiceDesc, srv)
}

func _PurchaseService_Purchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServiceServer).Purchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microservicesshop.PurchaseService/Purchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServiceServer).Purchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PurchaseService_ServiceDesc is the grpc.ServiceDesc for PurchaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PurchaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microservicesshop.PurchaseService",
	HandlerType: (*PurchaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Purchase",
			Handler:    _PurchaseService_Purchase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}
