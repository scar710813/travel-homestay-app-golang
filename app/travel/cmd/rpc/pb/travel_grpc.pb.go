// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: travel.proto

package pb

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

const (
	Travel_HomestayDetail_FullMethodName   = "/pb.travel/homestayDetail"
	Travel_AddHomestay_FullMethodName      = "/pb.travel/addHomestay"
	Travel_AddComment_FullMethodName       = "/pb.travel/addComment"
	Travel_WishList_FullMethodName         = "/pb.travel/wishList"
	Travel_AddWishList_FullMethodName      = "/pb.travel/addWishList"
	Travel_RemoveWishList_FullMethodName   = "/pb.travel/removeWishList"
	Travel_HistoryList_FullMethodName      = "/pb.travel/historyList"
	Travel_RemoveHistory_FullMethodName    = "/pb.travel/removeHistory"
	Travel_ClearHistory_FullMethodName     = "/pb.travel/clearHistory"
	Travel_SearchByLocation_FullMethodName = "/pb.travel/searchByLocation"
)

// TravelClient is the client API for Travel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelClient interface {
	// homestayDetail
	HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error)
	AddHomestay(ctx context.Context, in *AddHomestayReq, opts ...grpc.CallOption) (*AddHomestayResp, error)
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
	WishList(ctx context.Context, in *WishListReq, opts ...grpc.CallOption) (*WishListResp, error)
	AddWishList(ctx context.Context, in *AddWishListReq, opts ...grpc.CallOption) (*AddWishListResp, error)
	RemoveWishList(ctx context.Context, in *RemoveWishListReq, opts ...grpc.CallOption) (*RemoveWishListResp, error)
	HistoryList(ctx context.Context, in *HistoryListReq, opts ...grpc.CallOption) (*HistoryListResp, error)
	RemoveHistory(ctx context.Context, in *RemoveHistoryReq, opts ...grpc.CallOption) (*RemoveHistoryResp, error)
	ClearHistory(ctx context.Context, in *ClearHistoryReq, opts ...grpc.CallOption) (*ClearHistoryResp, error)
	SearchByLocation(ctx context.Context, in *SearchByLocationReq, opts ...grpc.CallOption) (*SearchByLocationResp, error)
}

type travelClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelClient(cc grpc.ClientConnInterface) TravelClient {
	return &travelClient{cc}
}

func (c *travelClient) HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error) {
	out := new(HomestayDetailResp)
	err := c.cc.Invoke(ctx, Travel_HomestayDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AddHomestay(ctx context.Context, in *AddHomestayReq, opts ...grpc.CallOption) (*AddHomestayResp, error) {
	out := new(AddHomestayResp)
	err := c.cc.Invoke(ctx, Travel_AddHomestay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, Travel_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) WishList(ctx context.Context, in *WishListReq, opts ...grpc.CallOption) (*WishListResp, error) {
	out := new(WishListResp)
	err := c.cc.Invoke(ctx, Travel_WishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AddWishList(ctx context.Context, in *AddWishListReq, opts ...grpc.CallOption) (*AddWishListResp, error) {
	out := new(AddWishListResp)
	err := c.cc.Invoke(ctx, Travel_AddWishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) RemoveWishList(ctx context.Context, in *RemoveWishListReq, opts ...grpc.CallOption) (*RemoveWishListResp, error) {
	out := new(RemoveWishListResp)
	err := c.cc.Invoke(ctx, Travel_RemoveWishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) HistoryList(ctx context.Context, in *HistoryListReq, opts ...grpc.CallOption) (*HistoryListResp, error) {
	out := new(HistoryListResp)
	err := c.cc.Invoke(ctx, Travel_HistoryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) RemoveHistory(ctx context.Context, in *RemoveHistoryReq, opts ...grpc.CallOption) (*RemoveHistoryResp, error) {
	out := new(RemoveHistoryResp)
	err := c.cc.Invoke(ctx, Travel_RemoveHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) ClearHistory(ctx context.Context, in *ClearHistoryReq, opts ...grpc.CallOption) (*ClearHistoryResp, error) {
	out := new(ClearHistoryResp)
	err := c.cc.Invoke(ctx, Travel_ClearHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) SearchByLocation(ctx context.Context, in *SearchByLocationReq, opts ...grpc.CallOption) (*SearchByLocationResp, error) {
	out := new(SearchByLocationResp)
	err := c.cc.Invoke(ctx, Travel_SearchByLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelServer is the server API for Travel service.
// All implementations must embed UnimplementedTravelServer
// for forward compatibility
type TravelServer interface {
	// homestayDetail
	HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error)
	AddHomestay(context.Context, *AddHomestayReq) (*AddHomestayResp, error)
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	WishList(context.Context, *WishListReq) (*WishListResp, error)
	AddWishList(context.Context, *AddWishListReq) (*AddWishListResp, error)
	RemoveWishList(context.Context, *RemoveWishListReq) (*RemoveWishListResp, error)
	HistoryList(context.Context, *HistoryListReq) (*HistoryListResp, error)
	RemoveHistory(context.Context, *RemoveHistoryReq) (*RemoveHistoryResp, error)
	ClearHistory(context.Context, *ClearHistoryReq) (*ClearHistoryResp, error)
	SearchByLocation(context.Context, *SearchByLocationReq) (*SearchByLocationResp, error)
	mustEmbedUnimplementedTravelServer()
}

// UnimplementedTravelServer must be embedded to have forward compatible implementations.
type UnimplementedTravelServer struct {
}

func (UnimplementedTravelServer) HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomestayDetail not implemented")
}
func (UnimplementedTravelServer) AddHomestay(context.Context, *AddHomestayReq) (*AddHomestayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHomestay not implemented")
}
func (UnimplementedTravelServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedTravelServer) WishList(context.Context, *WishListReq) (*WishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WishList not implemented")
}
func (UnimplementedTravelServer) AddWishList(context.Context, *AddWishListReq) (*AddWishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWishList not implemented")
}
func (UnimplementedTravelServer) RemoveWishList(context.Context, *RemoveWishListReq) (*RemoveWishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWishList not implemented")
}
func (UnimplementedTravelServer) HistoryList(context.Context, *HistoryListReq) (*HistoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryList not implemented")
}
func (UnimplementedTravelServer) RemoveHistory(context.Context, *RemoveHistoryReq) (*RemoveHistoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHistory not implemented")
}
func (UnimplementedTravelServer) ClearHistory(context.Context, *ClearHistoryReq) (*ClearHistoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearHistory not implemented")
}
func (UnimplementedTravelServer) SearchByLocation(context.Context, *SearchByLocationReq) (*SearchByLocationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByLocation not implemented")
}
func (UnimplementedTravelServer) mustEmbedUnimplementedTravelServer() {}

// UnsafeTravelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelServer will
// result in compilation errors.
type UnsafeTravelServer interface {
	mustEmbedUnimplementedTravelServer()
}

func RegisterTravelServer(s grpc.ServiceRegistrar, srv TravelServer) {
	s.RegisterService(&Travel_ServiceDesc, srv)
}

func _Travel_HomestayDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomestayDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).HomestayDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_HomestayDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).HomestayDetail(ctx, req.(*HomestayDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AddHomestay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHomestayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AddHomestay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_AddHomestay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AddHomestay(ctx, req.(*AddHomestayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_WishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).WishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_WishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).WishList(ctx, req.(*WishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AddWishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AddWishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_AddWishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AddWishList(ctx, req.(*AddWishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_RemoveWishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).RemoveWishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_RemoveWishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).RemoveWishList(ctx, req.(*RemoveWishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_HistoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).HistoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_HistoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).HistoryList(ctx, req.(*HistoryListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_RemoveHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).RemoveHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_RemoveHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).RemoveHistory(ctx, req.(*RemoveHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_ClearHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).ClearHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_ClearHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).ClearHistory(ctx, req.(*ClearHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_SearchByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchByLocationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).SearchByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_SearchByLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).SearchByLocation(ctx, req.(*SearchByLocationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Travel_ServiceDesc is the grpc.ServiceDesc for Travel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Travel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.travel",
	HandlerType: (*TravelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "homestayDetail",
			Handler:    _Travel_HomestayDetail_Handler,
		},
		{
			MethodName: "addHomestay",
			Handler:    _Travel_AddHomestay_Handler,
		},
		{
			MethodName: "addComment",
			Handler:    _Travel_AddComment_Handler,
		},
		{
			MethodName: "wishList",
			Handler:    _Travel_WishList_Handler,
		},
		{
			MethodName: "addWishList",
			Handler:    _Travel_AddWishList_Handler,
		},
		{
			MethodName: "removeWishList",
			Handler:    _Travel_RemoveWishList_Handler,
		},
		{
			MethodName: "historyList",
			Handler:    _Travel_HistoryList_Handler,
		},
		{
			MethodName: "removeHistory",
			Handler:    _Travel_RemoveHistory_Handler,
		},
		{
			MethodName: "clearHistory",
			Handler:    _Travel_ClearHistory_Handler,
		},
		{
			MethodName: "searchByLocation",
			Handler:    _Travel_SearchByLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "travel.proto",
}
