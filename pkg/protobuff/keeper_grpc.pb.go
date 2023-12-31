// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: keeper.proto

package protobuff

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
	KeeperService_RegisterUser_FullMethodName  = "/keeper.keeperService/RegisterUser"
	KeeperService_LoginUser_FullMethodName     = "/keeper.keeperService/LoginUser"
	KeeperService_CreateAccount_FullMethodName = "/keeper.keeperService/CreateAccount"
	KeeperService_DeleteAccount_FullMethodName = "/keeper.keeperService/DeleteAccount"
	KeeperService_UpdateAccount_FullMethodName = "/keeper.keeperService/UpdateAccount"
	KeeperService_ListAccount_FullMethodName   = "/keeper.keeperService/ListAccount"
	KeeperService_CreateText_FullMethodName    = "/keeper.keeperService/CreateText"
	KeeperService_DeleteText_FullMethodName    = "/keeper.keeperService/DeleteText"
	KeeperService_UpdateText_FullMethodName    = "/keeper.keeperService/UpdateText"
	KeeperService_ListText_FullMethodName      = "/keeper.keeperService/ListText"
	KeeperService_CreateBinary_FullMethodName  = "/keeper.keeperService/CreateBinary"
	KeeperService_DeleteBinary_FullMethodName  = "/keeper.keeperService/DeleteBinary"
	KeeperService_UpdateBinary_FullMethodName  = "/keeper.keeperService/UpdateBinary"
	KeeperService_ListBinary_FullMethodName    = "/keeper.keeperService/ListBinary"
	KeeperService_CreateCard_FullMethodName    = "/keeper.keeperService/CreateCard"
	KeeperService_DeleteCard_FullMethodName    = "/keeper.keeperService/DeleteCard"
	KeeperService_UpdateCard_FullMethodName    = "/keeper.keeperService/UpdateCard"
	KeeperService_ListCard_FullMethodName      = "/keeper.keeperService/ListCard"
)

// KeeperServiceClient is the client API for KeeperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeeperServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccountResponse, error)
	CreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error)
	DeleteText(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteTextResponse, error)
	UpdateText(ctx context.Context, in *UpdateTextRequest, opts ...grpc.CallOption) (*UpdateTextResponse, error)
	ListText(ctx context.Context, in *ListTextRequest, opts ...grpc.CallOption) (*ListTextResponse, error)
	CreateBinary(ctx context.Context, in *CreateBinaryRequest, opts ...grpc.CallOption) (*CreateBinaryResponse, error)
	DeleteBinary(ctx context.Context, in *DeleteBinaryRequest, opts ...grpc.CallOption) (*DeleteBinaryResponse, error)
	UpdateBinary(ctx context.Context, in *UpdateBinaryRequest, opts ...grpc.CallOption) (*UpdateBinaryResponse, error)
	ListBinary(ctx context.Context, in *ListBinaryRequest, opts ...grpc.CallOption) (*ListBinaryResponse, error)
	CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error)
	DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error)
	UpdateCard(ctx context.Context, in *UpdateCardRequest, opts ...grpc.CallOption) (*UpdateCardResponse, error)
	ListCard(ctx context.Context, in *ListCardRequest, opts ...grpc.CallOption) (*ListCardResponse, error)
}

type keeperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeeperServiceClient(cc grpc.ClientConnInterface) KeeperServiceClient {
	return &keeperServiceClient{cc}
}

func (c *keeperServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, KeeperService_RegisterUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, KeeperService_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, KeeperService_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, KeeperService_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, KeeperService_UpdateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccountResponse, error) {
	out := new(ListAccountResponse)
	err := c.cc.Invoke(ctx, KeeperService_ListAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) CreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error) {
	out := new(CreateTextResponse)
	err := c.cc.Invoke(ctx, KeeperService_CreateText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) DeleteText(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteTextResponse, error) {
	out := new(DeleteTextResponse)
	err := c.cc.Invoke(ctx, KeeperService_DeleteText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) UpdateText(ctx context.Context, in *UpdateTextRequest, opts ...grpc.CallOption) (*UpdateTextResponse, error) {
	out := new(UpdateTextResponse)
	err := c.cc.Invoke(ctx, KeeperService_UpdateText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) ListText(ctx context.Context, in *ListTextRequest, opts ...grpc.CallOption) (*ListTextResponse, error) {
	out := new(ListTextResponse)
	err := c.cc.Invoke(ctx, KeeperService_ListText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) CreateBinary(ctx context.Context, in *CreateBinaryRequest, opts ...grpc.CallOption) (*CreateBinaryResponse, error) {
	out := new(CreateBinaryResponse)
	err := c.cc.Invoke(ctx, KeeperService_CreateBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) DeleteBinary(ctx context.Context, in *DeleteBinaryRequest, opts ...grpc.CallOption) (*DeleteBinaryResponse, error) {
	out := new(DeleteBinaryResponse)
	err := c.cc.Invoke(ctx, KeeperService_DeleteBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) UpdateBinary(ctx context.Context, in *UpdateBinaryRequest, opts ...grpc.CallOption) (*UpdateBinaryResponse, error) {
	out := new(UpdateBinaryResponse)
	err := c.cc.Invoke(ctx, KeeperService_UpdateBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) ListBinary(ctx context.Context, in *ListBinaryRequest, opts ...grpc.CallOption) (*ListBinaryResponse, error) {
	out := new(ListBinaryResponse)
	err := c.cc.Invoke(ctx, KeeperService_ListBinary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*CreateCardResponse, error) {
	out := new(CreateCardResponse)
	err := c.cc.Invoke(ctx, KeeperService_CreateCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error) {
	out := new(DeleteCardResponse)
	err := c.cc.Invoke(ctx, KeeperService_DeleteCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) UpdateCard(ctx context.Context, in *UpdateCardRequest, opts ...grpc.CallOption) (*UpdateCardResponse, error) {
	out := new(UpdateCardResponse)
	err := c.cc.Invoke(ctx, KeeperService_UpdateCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) ListCard(ctx context.Context, in *ListCardRequest, opts ...grpc.CallOption) (*ListCardResponse, error) {
	out := new(ListCardResponse)
	err := c.cc.Invoke(ctx, KeeperService_ListCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeeperServiceServer is the server API for KeeperService service.
// All implementations must embed UnimplementedKeeperServiceServer
// for forward compatibility
type KeeperServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	ListAccount(context.Context, *ListAccountRequest) (*ListAccountResponse, error)
	CreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error)
	DeleteText(context.Context, *DeleteTextRequest) (*DeleteTextResponse, error)
	UpdateText(context.Context, *UpdateTextRequest) (*UpdateTextResponse, error)
	ListText(context.Context, *ListTextRequest) (*ListTextResponse, error)
	CreateBinary(context.Context, *CreateBinaryRequest) (*CreateBinaryResponse, error)
	DeleteBinary(context.Context, *DeleteBinaryRequest) (*DeleteBinaryResponse, error)
	UpdateBinary(context.Context, *UpdateBinaryRequest) (*UpdateBinaryResponse, error)
	ListBinary(context.Context, *ListBinaryRequest) (*ListBinaryResponse, error)
	CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error)
	DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error)
	UpdateCard(context.Context, *UpdateCardRequest) (*UpdateCardResponse, error)
	ListCard(context.Context, *ListCardRequest) (*ListCardResponse, error)
	mustEmbedUnimplementedKeeperServiceServer()
}

// UnimplementedKeeperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeeperServiceServer struct {
}

func (UnimplementedKeeperServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedKeeperServiceServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedKeeperServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedKeeperServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedKeeperServiceServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedKeeperServiceServer) ListAccount(context.Context, *ListAccountRequest) (*ListAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
}
func (UnimplementedKeeperServiceServer) CreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateText not implemented")
}
func (UnimplementedKeeperServiceServer) DeleteText(context.Context, *DeleteTextRequest) (*DeleteTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteText not implemented")
}
func (UnimplementedKeeperServiceServer) UpdateText(context.Context, *UpdateTextRequest) (*UpdateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateText not implemented")
}
func (UnimplementedKeeperServiceServer) ListText(context.Context, *ListTextRequest) (*ListTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListText not implemented")
}
func (UnimplementedKeeperServiceServer) CreateBinary(context.Context, *CreateBinaryRequest) (*CreateBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBinary not implemented")
}
func (UnimplementedKeeperServiceServer) DeleteBinary(context.Context, *DeleteBinaryRequest) (*DeleteBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBinary not implemented")
}
func (UnimplementedKeeperServiceServer) UpdateBinary(context.Context, *UpdateBinaryRequest) (*UpdateBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBinary not implemented")
}
func (UnimplementedKeeperServiceServer) ListBinary(context.Context, *ListBinaryRequest) (*ListBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBinary not implemented")
}
func (UnimplementedKeeperServiceServer) CreateCard(context.Context, *CreateCardRequest) (*CreateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedKeeperServiceServer) DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCard not implemented")
}
func (UnimplementedKeeperServiceServer) UpdateCard(context.Context, *UpdateCardRequest) (*UpdateCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCard not implemented")
}
func (UnimplementedKeeperServiceServer) ListCard(context.Context, *ListCardRequest) (*ListCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCard not implemented")
}
func (UnimplementedKeeperServiceServer) mustEmbedUnimplementedKeeperServiceServer() {}

// UnsafeKeeperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeeperServiceServer will
// result in compilation errors.
type UnsafeKeeperServiceServer interface {
	mustEmbedUnimplementedKeeperServiceServer()
}

func RegisterKeeperServiceServer(s grpc.ServiceRegistrar, srv KeeperServiceServer) {
	s.RegisterService(&KeeperService_ServiceDesc, srv)
}

func _KeeperService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_ListAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).ListAccount(ctx, req.(*ListAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_CreateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).CreateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_CreateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).CreateText(ctx, req.(*CreateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_DeleteText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).DeleteText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_DeleteText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).DeleteText(ctx, req.(*DeleteTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_UpdateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).UpdateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_UpdateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).UpdateText(ctx, req.(*UpdateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_ListText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).ListText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_ListText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).ListText(ctx, req.(*ListTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_CreateBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).CreateBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_CreateBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).CreateBinary(ctx, req.(*CreateBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_DeleteBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).DeleteBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_DeleteBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).DeleteBinary(ctx, req.(*DeleteBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_UpdateBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).UpdateBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_UpdateBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).UpdateBinary(ctx, req.(*UpdateBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_ListBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).ListBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_ListBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).ListBinary(ctx, req.(*ListBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_CreateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).CreateCard(ctx, req.(*CreateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_DeleteCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).DeleteCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_DeleteCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).DeleteCard(ctx, req.(*DeleteCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_UpdateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).UpdateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_UpdateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).UpdateCard(ctx, req.(*UpdateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_ListCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).ListCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_ListCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).ListCard(ctx, req.(*ListCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeeperService_ServiceDesc is the grpc.ServiceDesc for KeeperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeeperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.keeperService",
	HandlerType: (*KeeperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _KeeperService_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _KeeperService_LoginUser_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _KeeperService_CreateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _KeeperService_DeleteAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _KeeperService_UpdateAccount_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _KeeperService_ListAccount_Handler,
		},
		{
			MethodName: "CreateText",
			Handler:    _KeeperService_CreateText_Handler,
		},
		{
			MethodName: "DeleteText",
			Handler:    _KeeperService_DeleteText_Handler,
		},
		{
			MethodName: "UpdateText",
			Handler:    _KeeperService_UpdateText_Handler,
		},
		{
			MethodName: "ListText",
			Handler:    _KeeperService_ListText_Handler,
		},
		{
			MethodName: "CreateBinary",
			Handler:    _KeeperService_CreateBinary_Handler,
		},
		{
			MethodName: "DeleteBinary",
			Handler:    _KeeperService_DeleteBinary_Handler,
		},
		{
			MethodName: "UpdateBinary",
			Handler:    _KeeperService_UpdateBinary_Handler,
		},
		{
			MethodName: "ListBinary",
			Handler:    _KeeperService_ListBinary_Handler,
		},
		{
			MethodName: "CreateCard",
			Handler:    _KeeperService_CreateCard_Handler,
		},
		{
			MethodName: "DeleteCard",
			Handler:    _KeeperService_DeleteCard_Handler,
		},
		{
			MethodName: "UpdateCard",
			Handler:    _KeeperService_UpdateCard_Handler,
		},
		{
			MethodName: "ListCard",
			Handler:    _KeeperService_ListCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper.proto",
}
