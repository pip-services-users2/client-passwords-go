// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: protos/passwords_v1.proto

package protos

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

// PasswordsClient is the client API for Passwords service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PasswordsClient interface {
	GetPasswordInfo(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordInfoReply, error)
	ValidatePassword(ctx context.Context, in *PasswordValueRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error)
	SetPassword(ctx context.Context, in *PasswordIdAndValueRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error)
	SetTempPassword(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordValueReply, error)
	DeletePassword(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error)
	Authenticate(ctx context.Context, in *PasswordIdAndValueRequest, opts ...grpc.CallOption) (*PasswordAuthenticateReply, error)
	ChangePassword(ctx context.Context, in *PasswordIdAndValuesRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error)
	ValidateCode(ctx context.Context, in *PasswordIdAndCodeRequest, opts ...grpc.CallOption) (*PasswordValidReply, error)
	ResetPassword(ctx context.Context, in *PasswordIdAndCodeAndValueRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error)
	RecoverPassword(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error)
}

type passwordsClient struct {
	cc grpc.ClientConnInterface
}

func NewPasswordsClient(cc grpc.ClientConnInterface) PasswordsClient {
	return &passwordsClient{cc}
}

func (c *passwordsClient) GetPasswordInfo(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordInfoReply, error) {
	out := new(PasswordInfoReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/get_password_info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) ValidatePassword(ctx context.Context, in *PasswordValueRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error) {
	out := new(PasswordEmptyReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/validate_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) SetPassword(ctx context.Context, in *PasswordIdAndValueRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error) {
	out := new(PasswordEmptyReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/set_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) SetTempPassword(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordValueReply, error) {
	out := new(PasswordValueReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/set_temp_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) DeletePassword(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error) {
	out := new(PasswordEmptyReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/delete_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) Authenticate(ctx context.Context, in *PasswordIdAndValueRequest, opts ...grpc.CallOption) (*PasswordAuthenticateReply, error) {
	out := new(PasswordAuthenticateReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) ChangePassword(ctx context.Context, in *PasswordIdAndValuesRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error) {
	out := new(PasswordEmptyReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/change_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) ValidateCode(ctx context.Context, in *PasswordIdAndCodeRequest, opts ...grpc.CallOption) (*PasswordValidReply, error) {
	out := new(PasswordValidReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/validate_code", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) ResetPassword(ctx context.Context, in *PasswordIdAndCodeAndValueRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error) {
	out := new(PasswordEmptyReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/reset_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsClient) RecoverPassword(ctx context.Context, in *PasswordIdRequest, opts ...grpc.CallOption) (*PasswordEmptyReply, error) {
	out := new(PasswordEmptyReply)
	err := c.cc.Invoke(ctx, "/passwords_v1.Passwords/recover_password", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordsServer is the server API for Passwords service.
// All implementations must embed UnimplementedPasswordsServer
// for forward compatibility
type PasswordsServer interface {
	GetPasswordInfo(context.Context, *PasswordIdRequest) (*PasswordInfoReply, error)
	ValidatePassword(context.Context, *PasswordValueRequest) (*PasswordEmptyReply, error)
	SetPassword(context.Context, *PasswordIdAndValueRequest) (*PasswordEmptyReply, error)
	SetTempPassword(context.Context, *PasswordIdRequest) (*PasswordValueReply, error)
	DeletePassword(context.Context, *PasswordIdRequest) (*PasswordEmptyReply, error)
	Authenticate(context.Context, *PasswordIdAndValueRequest) (*PasswordAuthenticateReply, error)
	ChangePassword(context.Context, *PasswordIdAndValuesRequest) (*PasswordEmptyReply, error)
	ValidateCode(context.Context, *PasswordIdAndCodeRequest) (*PasswordValidReply, error)
	ResetPassword(context.Context, *PasswordIdAndCodeAndValueRequest) (*PasswordEmptyReply, error)
	RecoverPassword(context.Context, *PasswordIdRequest) (*PasswordEmptyReply, error)
	mustEmbedUnimplementedPasswordsServer()
}

// UnimplementedPasswordsServer must be embedded to have forward compatible implementations.
type UnimplementedPasswordsServer struct {
}

func (UnimplementedPasswordsServer) GetPasswordInfo(context.Context, *PasswordIdRequest) (*PasswordInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswordInfo not implemented")
}
func (UnimplementedPasswordsServer) ValidatePassword(context.Context, *PasswordValueRequest) (*PasswordEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePassword not implemented")
}
func (UnimplementedPasswordsServer) SetPassword(context.Context, *PasswordIdAndValueRequest) (*PasswordEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedPasswordsServer) SetTempPassword(context.Context, *PasswordIdRequest) (*PasswordValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTempPassword not implemented")
}
func (UnimplementedPasswordsServer) DeletePassword(context.Context, *PasswordIdRequest) (*PasswordEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassword not implemented")
}
func (UnimplementedPasswordsServer) Authenticate(context.Context, *PasswordIdAndValueRequest) (*PasswordAuthenticateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedPasswordsServer) ChangePassword(context.Context, *PasswordIdAndValuesRequest) (*PasswordEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedPasswordsServer) ValidateCode(context.Context, *PasswordIdAndCodeRequest) (*PasswordValidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCode not implemented")
}
func (UnimplementedPasswordsServer) ResetPassword(context.Context, *PasswordIdAndCodeAndValueRequest) (*PasswordEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedPasswordsServer) RecoverPassword(context.Context, *PasswordIdRequest) (*PasswordEmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverPassword not implemented")
}
func (UnimplementedPasswordsServer) mustEmbedUnimplementedPasswordsServer() {}

// UnsafePasswordsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PasswordsServer will
// result in compilation errors.
type UnsafePasswordsServer interface {
	mustEmbedUnimplementedPasswordsServer()
}

func RegisterPasswordsServer(s grpc.ServiceRegistrar, srv PasswordsServer) {
	s.RegisterService(&Passwords_ServiceDesc, srv)
}

func _Passwords_GetPasswordInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).GetPasswordInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/get_password_info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).GetPasswordInfo(ctx, req.(*PasswordIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_ValidatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).ValidatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/validate_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).ValidatePassword(ctx, req.(*PasswordValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdAndValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/set_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).SetPassword(ctx, req.(*PasswordIdAndValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_SetTempPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).SetTempPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/set_temp_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).SetTempPassword(ctx, req.(*PasswordIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/delete_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).DeletePassword(ctx, req.(*PasswordIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdAndValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).Authenticate(ctx, req.(*PasswordIdAndValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdAndValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/change_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).ChangePassword(ctx, req.(*PasswordIdAndValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_ValidateCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdAndCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).ValidateCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/validate_code",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).ValidateCode(ctx, req.(*PasswordIdAndCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdAndCodeAndValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/reset_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).ResetPassword(ctx, req.(*PasswordIdAndCodeAndValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passwords_RecoverPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsServer).RecoverPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwords_v1.Passwords/recover_password",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsServer).RecoverPassword(ctx, req.(*PasswordIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Passwords_ServiceDesc is the grpc.ServiceDesc for Passwords service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Passwords_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "passwords_v1.Passwords",
	HandlerType: (*PasswordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_password_info",
			Handler:    _Passwords_GetPasswordInfo_Handler,
		},
		{
			MethodName: "validate_password",
			Handler:    _Passwords_ValidatePassword_Handler,
		},
		{
			MethodName: "set_password",
			Handler:    _Passwords_SetPassword_Handler,
		},
		{
			MethodName: "set_temp_password",
			Handler:    _Passwords_SetTempPassword_Handler,
		},
		{
			MethodName: "delete_password",
			Handler:    _Passwords_DeletePassword_Handler,
		},
		{
			MethodName: "authenticate",
			Handler:    _Passwords_Authenticate_Handler,
		},
		{
			MethodName: "change_password",
			Handler:    _Passwords_ChangePassword_Handler,
		},
		{
			MethodName: "validate_code",
			Handler:    _Passwords_ValidateCode_Handler,
		},
		{
			MethodName: "reset_password",
			Handler:    _Passwords_ResetPassword_Handler,
		},
		{
			MethodName: "recover_password",
			Handler:    _Passwords_RecoverPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/passwords_v1.proto",
}