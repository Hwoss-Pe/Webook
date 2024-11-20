// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/proto/user/v1/user.proto

package userv1

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

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	FindOrCreate(ctx context.Context, in *FindOrCreateRequest, opts ...grpc.CallOption) (*FindOrCreateResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	UpdateNonSensitiveInfo(ctx context.Context, in *UpdateNonSensitiveInfoRequest, opts ...grpc.CallOption) (*UpdateNonSensitiveInfoResponse, error)
	FindOrCreateByWechat(ctx context.Context, in *FindOrCreateByWechatRequest, opts ...grpc.CallOption) (*FindOrCreateByWechatResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UsersService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) FindOrCreate(ctx context.Context, in *FindOrCreateRequest, opts ...grpc.CallOption) (*FindOrCreateResponse, error) {
	out := new(FindOrCreateResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UsersService/FindOrCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UsersService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UsersService/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) UpdateNonSensitiveInfo(ctx context.Context, in *UpdateNonSensitiveInfoRequest, opts ...grpc.CallOption) (*UpdateNonSensitiveInfoResponse, error) {
	out := new(UpdateNonSensitiveInfoResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UsersService/UpdateNonSensitiveInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) FindOrCreateByWechat(ctx context.Context, in *FindOrCreateByWechatRequest, opts ...grpc.CallOption) (*FindOrCreateByWechatResponse, error) {
	out := new(FindOrCreateByWechatResponse)
	err := c.cc.Invoke(ctx, "/user.v1.UsersService/FindOrCreateByWechat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	FindOrCreate(context.Context, *FindOrCreateRequest) (*FindOrCreateResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Profile(context.Context, *ProfileRequest) (*ProfileResponse, error)
	UpdateNonSensitiveInfo(context.Context, *UpdateNonSensitiveInfoRequest) (*UpdateNonSensitiveInfoResponse, error)
	FindOrCreateByWechat(context.Context, *FindOrCreateByWechatRequest) (*FindOrCreateByWechatResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUsersServiceServer) FindOrCreate(context.Context, *FindOrCreateRequest) (*FindOrCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrCreate not implemented")
}
func (UnimplementedUsersServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsersServiceServer) Profile(context.Context, *ProfileRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedUsersServiceServer) UpdateNonSensitiveInfo(context.Context, *UpdateNonSensitiveInfoRequest) (*UpdateNonSensitiveInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNonSensitiveInfo not implemented")
}
func (UnimplementedUsersServiceServer) FindOrCreateByWechat(context.Context, *FindOrCreateByWechatRequest) (*FindOrCreateByWechatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrCreateByWechat not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UsersService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_FindOrCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).FindOrCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UsersService/FindOrCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).FindOrCreate(ctx, req.(*FindOrCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UsersService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UsersService/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Profile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_UpdateNonSensitiveInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNonSensitiveInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).UpdateNonSensitiveInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UsersService/UpdateNonSensitiveInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).UpdateNonSensitiveInfo(ctx, req.(*UpdateNonSensitiveInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_FindOrCreateByWechat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrCreateByWechatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).FindOrCreateByWechat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UsersService/FindOrCreateByWechat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).FindOrCreateByWechat(ctx, req.(*FindOrCreateByWechatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _UsersService_Signup_Handler,
		},
		{
			MethodName: "FindOrCreate",
			Handler:    _UsersService_FindOrCreate_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UsersService_Login_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _UsersService_Profile_Handler,
		},
		{
			MethodName: "UpdateNonSensitiveInfo",
			Handler:    _UsersService_UpdateNonSensitiveInfo_Handler,
		},
		{
			MethodName: "FindOrCreateByWechat",
			Handler:    _UsersService_FindOrCreateByWechat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/user/v1/user.proto",
}
