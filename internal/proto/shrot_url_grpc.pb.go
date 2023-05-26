// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: shrot_url.proto

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

// URLsClient is the client API for URLs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type URLsClient interface {
	ShortenURL(ctx context.Context, in *ShortenURLRequest, opts ...grpc.CallOption) (*ShortenURLResponse, error)
	DeleteURLs(ctx context.Context, in *DeleteURLsRequest, opts ...grpc.CallOption) (*DeleteURLsResponse, error)
	ReturnFullURL(ctx context.Context, in *ReturnFullURLRequest, opts ...grpc.CallOption) (*ReturnFullURLResponse, error)
	GetFullURL(ctx context.Context, in *GetFullURLRequest, opts ...grpc.CallOption) (*GetFullURLResponse, error)
	GetByUserID(ctx context.Context, in *GetByUserIDRequest, opts ...grpc.CallOption) (*GetByUserIDResponse, error)
	SaveMany(ctx context.Context, in *SaveManyRequest, opts ...grpc.CallOption) (*SaveManyResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	GetAllUsersAndUrls(ctx context.Context, in *GetAllUsersAndUrlsRequest, opts ...grpc.CallOption) (*GetAllUsersAndUrlsResponse, error)
}

type uRLsClient struct {
	cc grpc.ClientConnInterface
}

func NewURLsClient(cc grpc.ClientConnInterface) URLsClient {
	return &uRLsClient{cc}
}

func (c *uRLsClient) ShortenURL(ctx context.Context, in *ShortenURLRequest, opts ...grpc.CallOption) (*ShortenURLResponse, error) {
	out := new(ShortenURLResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/ShortenURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLsClient) DeleteURLs(ctx context.Context, in *DeleteURLsRequest, opts ...grpc.CallOption) (*DeleteURLsResponse, error) {
	out := new(DeleteURLsResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/DeleteURLs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLsClient) ReturnFullURL(ctx context.Context, in *ReturnFullURLRequest, opts ...grpc.CallOption) (*ReturnFullURLResponse, error) {
	out := new(ReturnFullURLResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/ReturnFullURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLsClient) GetFullURL(ctx context.Context, in *GetFullURLRequest, opts ...grpc.CallOption) (*GetFullURLResponse, error) {
	out := new(GetFullURLResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/GetFullURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLsClient) GetByUserID(ctx context.Context, in *GetByUserIDRequest, opts ...grpc.CallOption) (*GetByUserIDResponse, error) {
	out := new(GetByUserIDResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/GetByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLsClient) SaveMany(ctx context.Context, in *SaveManyRequest, opts ...grpc.CallOption) (*SaveManyResponse, error) {
	out := new(SaveManyResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/SaveMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLsClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLsClient) GetAllUsersAndUrls(ctx context.Context, in *GetAllUsersAndUrlsRequest, opts ...grpc.CallOption) (*GetAllUsersAndUrlsResponse, error) {
	out := new(GetAllUsersAndUrlsResponse)
	err := c.cc.Invoke(ctx, "/demo.URLs/GetAllUsersAndUrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLsServer is the server API for URLs service.
// All implementations must embed UnimplementedURLsServer
// for forward compatibility
type URLsServer interface {
	ShortenURL(context.Context, *ShortenURLRequest) (*ShortenURLResponse, error)
	DeleteURLs(context.Context, *DeleteURLsRequest) (*DeleteURLsResponse, error)
	ReturnFullURL(context.Context, *ReturnFullURLRequest) (*ReturnFullURLResponse, error)
	GetFullURL(context.Context, *GetFullURLRequest) (*GetFullURLResponse, error)
	GetByUserID(context.Context, *GetByUserIDRequest) (*GetByUserIDResponse, error)
	SaveMany(context.Context, *SaveManyRequest) (*SaveManyResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	GetAllUsersAndUrls(context.Context, *GetAllUsersAndUrlsRequest) (*GetAllUsersAndUrlsResponse, error)
	mustEmbedUnimplementedURLsServer()
}

// UnimplementedURLsServer must be embedded to have forward compatible implementations.
type UnimplementedURLsServer struct {
}

func (UnimplementedURLsServer) ShortenURL(context.Context, *ShortenURLRequest) (*ShortenURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenURL not implemented")
}
func (UnimplementedURLsServer) DeleteURLs(context.Context, *DeleteURLsRequest) (*DeleteURLsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteURLs not implemented")
}
func (UnimplementedURLsServer) ReturnFullURL(context.Context, *ReturnFullURLRequest) (*ReturnFullURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnFullURL not implemented")
}
func (UnimplementedURLsServer) GetFullURL(context.Context, *GetFullURLRequest) (*GetFullURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullURL not implemented")
}
func (UnimplementedURLsServer) GetByUserID(context.Context, *GetByUserIDRequest) (*GetByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUserID not implemented")
}
func (UnimplementedURLsServer) SaveMany(context.Context, *SaveManyRequest) (*SaveManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMany not implemented")
}
func (UnimplementedURLsServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedURLsServer) GetAllUsersAndUrls(context.Context, *GetAllUsersAndUrlsRequest) (*GetAllUsersAndUrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsersAndUrls not implemented")
}
func (UnimplementedURLsServer) mustEmbedUnimplementedURLsServer() {}

// UnsafeURLsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to URLsServer will
// result in compilation errors.
type UnsafeURLsServer interface {
	mustEmbedUnimplementedURLsServer()
}

func RegisterURLsServer(s grpc.ServiceRegistrar, srv URLsServer) {
	s.RegisterService(&URLs_ServiceDesc, srv)
}

func _URLs_ShortenURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).ShortenURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/ShortenURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).ShortenURL(ctx, req.(*ShortenURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLs_DeleteURLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteURLsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).DeleteURLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/DeleteURLs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).DeleteURLs(ctx, req.(*DeleteURLsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLs_ReturnFullURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnFullURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).ReturnFullURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/ReturnFullURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).ReturnFullURL(ctx, req.(*ReturnFullURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLs_GetFullURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).GetFullURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/GetFullURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).GetFullURL(ctx, req.(*GetFullURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLs_GetByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).GetByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/GetByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).GetByUserID(ctx, req.(*GetByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLs_SaveMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).SaveMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/SaveMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).SaveMany(ctx, req.(*SaveManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLs_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLs_GetAllUsersAndUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersAndUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLsServer).GetAllUsersAndUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.URLs/GetAllUsersAndUrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLsServer).GetAllUsersAndUrls(ctx, req.(*GetAllUsersAndUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// URLs_ServiceDesc is the grpc.ServiceDesc for URLs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var URLs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.URLs",
	HandlerType: (*URLsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShortenURL",
			Handler:    _URLs_ShortenURL_Handler,
		},
		{
			MethodName: "DeleteURLs",
			Handler:    _URLs_DeleteURLs_Handler,
		},
		{
			MethodName: "ReturnFullURL",
			Handler:    _URLs_ReturnFullURL_Handler,
		},
		{
			MethodName: "GetFullURL",
			Handler:    _URLs_GetFullURL_Handler,
		},
		{
			MethodName: "GetByUserID",
			Handler:    _URLs_GetByUserID_Handler,
		},
		{
			MethodName: "SaveMany",
			Handler:    _URLs_SaveMany_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _URLs_Ping_Handler,
		},
		{
			MethodName: "GetAllUsersAndUrls",
			Handler:    _URLs_GetAllUsersAndUrls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shrot_url.proto",
}