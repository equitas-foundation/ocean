// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: ocean/v1/account.proto

package oceanv1

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	// CreateAccountBIP44 creates a new BIP44 account.
	CreateAccountBIP44(ctx context.Context, in *CreateAccountBIP44Request, opts ...grpc.CallOption) (*CreateAccountBIP44Response, error)
	// CreateAccountMultiSig creates a new multisig account.
	CreateAccountMultiSig(ctx context.Context, in *CreateAccountMultiSigRequest, opts ...grpc.CallOption) (*CreateAccountMultiSigResponse, error)
	// CreateAccountCustom creates a new custom account for which loading a template.
	CreateAccountCustom(ctx context.Context, in *CreateAccountCustomRequest, opts ...grpc.CallOption) (*CreateAccountCustomResponse, error)
	// SetAccountTemplate sets the template for the account used to generate new addresses.
	SetAccountTemplate(ctx context.Context, in *SetAccountTemplateRequest, opts ...grpc.CallOption) (*SetAccountTemplateResponse, error)
	// DeriveAddresses generates new address(es) for the account.
	DeriveAddresses(ctx context.Context, in *DeriveAddressesRequest, opts ...grpc.CallOption) (*DeriveAddressesResponse, error)
	// DeriveChangeAddresses generates new change address(es) for the account.
	DeriveChangeAddresses(ctx context.Context, in *DeriveChangeAddressesRequest, opts ...grpc.CallOption) (*DeriveChangeAddressesResponse, error)
	// ListAddresses returns all derived addresses for the account.
	ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error)
	// Balance returns the balance for the account, or for specific list of
	// account's addresses.
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	// ListUtxos returns the utxos for the account, or specific list of
	// account's addresses.
	ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error)
	// DeleteAccount deletes an existing account. The operation is allowed only
	// if the account has zero balance.
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccountBIP44(ctx context.Context, in *CreateAccountBIP44Request, opts ...grpc.CallOption) (*CreateAccountBIP44Response, error) {
	out := new(CreateAccountBIP44Response)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/CreateAccountBIP44", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateAccountMultiSig(ctx context.Context, in *CreateAccountMultiSigRequest, opts ...grpc.CallOption) (*CreateAccountMultiSigResponse, error) {
	out := new(CreateAccountMultiSigResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/CreateAccountMultiSig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateAccountCustom(ctx context.Context, in *CreateAccountCustomRequest, opts ...grpc.CallOption) (*CreateAccountCustomResponse, error) {
	out := new(CreateAccountCustomResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/CreateAccountCustom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SetAccountTemplate(ctx context.Context, in *SetAccountTemplateRequest, opts ...grpc.CallOption) (*SetAccountTemplateResponse, error) {
	out := new(SetAccountTemplateResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/SetAccountTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeriveAddresses(ctx context.Context, in *DeriveAddressesRequest, opts ...grpc.CallOption) (*DeriveAddressesResponse, error) {
	out := new(DeriveAddressesResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/DeriveAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeriveChangeAddresses(ctx context.Context, in *DeriveChangeAddressesRequest, opts ...grpc.CallOption) (*DeriveChangeAddressesResponse, error) {
	out := new(DeriveChangeAddressesResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/DeriveChangeAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error) {
	out := new(ListAddressesResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/ListAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error) {
	out := new(ListUtxosResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/ListUtxos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1.AccountService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations should embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	// CreateAccountBIP44 creates a new BIP44 account.
	CreateAccountBIP44(context.Context, *CreateAccountBIP44Request) (*CreateAccountBIP44Response, error)
	// CreateAccountMultiSig creates a new multisig account.
	CreateAccountMultiSig(context.Context, *CreateAccountMultiSigRequest) (*CreateAccountMultiSigResponse, error)
	// CreateAccountCustom creates a new custom account for which loading a template.
	CreateAccountCustom(context.Context, *CreateAccountCustomRequest) (*CreateAccountCustomResponse, error)
	// SetAccountTemplate sets the template for the account used to generate new addresses.
	SetAccountTemplate(context.Context, *SetAccountTemplateRequest) (*SetAccountTemplateResponse, error)
	// DeriveAddresses generates new address(es) for the account.
	DeriveAddresses(context.Context, *DeriveAddressesRequest) (*DeriveAddressesResponse, error)
	// DeriveChangeAddresses generates new change address(es) for the account.
	DeriveChangeAddresses(context.Context, *DeriveChangeAddressesRequest) (*DeriveChangeAddressesResponse, error)
	// ListAddresses returns all derived addresses for the account.
	ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error)
	// Balance returns the balance for the account, or for specific list of
	// account's addresses.
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	// ListUtxos returns the utxos for the account, or specific list of
	// account's addresses.
	ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error)
	// DeleteAccount deletes an existing account. The operation is allowed only
	// if the account has zero balance.
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
}

// UnimplementedAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccountBIP44(context.Context, *CreateAccountBIP44Request) (*CreateAccountBIP44Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountBIP44 not implemented")
}
func (UnimplementedAccountServiceServer) CreateAccountMultiSig(context.Context, *CreateAccountMultiSigRequest) (*CreateAccountMultiSigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountMultiSig not implemented")
}
func (UnimplementedAccountServiceServer) CreateAccountCustom(context.Context, *CreateAccountCustomRequest) (*CreateAccountCustomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountCustom not implemented")
}
func (UnimplementedAccountServiceServer) SetAccountTemplate(context.Context, *SetAccountTemplateRequest) (*SetAccountTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccountTemplate not implemented")
}
func (UnimplementedAccountServiceServer) DeriveAddresses(context.Context, *DeriveAddressesRequest) (*DeriveAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeriveAddresses not implemented")
}
func (UnimplementedAccountServiceServer) DeriveChangeAddresses(context.Context, *DeriveChangeAddressesRequest) (*DeriveChangeAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeriveChangeAddresses not implemented")
}
func (UnimplementedAccountServiceServer) ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddresses not implemented")
}
func (UnimplementedAccountServiceServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedAccountServiceServer) ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUtxos not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccountBIP44_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountBIP44Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccountBIP44(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/CreateAccountBIP44",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccountBIP44(ctx, req.(*CreateAccountBIP44Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateAccountMultiSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountMultiSigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccountMultiSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/CreateAccountMultiSig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccountMultiSig(ctx, req.(*CreateAccountMultiSigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateAccountCustom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountCustomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccountCustom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/CreateAccountCustom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccountCustom(ctx, req.(*CreateAccountCustomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SetAccountTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAccountTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SetAccountTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/SetAccountTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SetAccountTemplate(ctx, req.(*SetAccountTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeriveAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeriveAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeriveAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/DeriveAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeriveAddresses(ctx, req.(*DeriveAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeriveChangeAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeriveChangeAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeriveChangeAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/DeriveChangeAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeriveChangeAddresses(ctx, req.(*DeriveChangeAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ListAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ListAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/ListAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ListAddresses(ctx, req.(*ListAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ListUtxos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUtxosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ListUtxos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/ListUtxos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ListUtxos(ctx, req.(*ListUtxosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1.AccountService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocean.v1.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccountBIP44",
			Handler:    _AccountService_CreateAccountBIP44_Handler,
		},
		{
			MethodName: "CreateAccountMultiSig",
			Handler:    _AccountService_CreateAccountMultiSig_Handler,
		},
		{
			MethodName: "CreateAccountCustom",
			Handler:    _AccountService_CreateAccountCustom_Handler,
		},
		{
			MethodName: "SetAccountTemplate",
			Handler:    _AccountService_SetAccountTemplate_Handler,
		},
		{
			MethodName: "DeriveAddresses",
			Handler:    _AccountService_DeriveAddresses_Handler,
		},
		{
			MethodName: "DeriveChangeAddresses",
			Handler:    _AccountService_DeriveChangeAddresses_Handler,
		},
		{
			MethodName: "ListAddresses",
			Handler:    _AccountService_ListAddresses_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _AccountService_Balance_Handler,
		},
		{
			MethodName: "ListUtxos",
			Handler:    _AccountService_ListUtxos_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocean/v1/account.proto",
}
