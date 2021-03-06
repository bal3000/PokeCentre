// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pokemon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PokemonServiceClient is the client API for PokemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokemonServiceClient interface {
	GetAllPokemon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PokemonList, error)
	GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*Pokemon, error)
	GetPokemonByType(ctx context.Context, in *GetPokemonByTypeRequest, opts ...grpc.CallOption) (*PokemonList, error)
	AddPokemon(ctx context.Context, in *AddPokemonRequest, opts ...grpc.CallOption) (*Pokemon, error)
	UpdatePokemon(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeletePokemon(ctx context.Context, in *DeletePokemonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pokemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonServiceClient(cc grpc.ClientConnInterface) PokemonServiceClient {
	return &pokemonServiceClient{cc}
}

func (c *pokemonServiceClient) GetAllPokemon(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PokemonList, error) {
	out := new(PokemonList)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/GetAllPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/GetPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) GetPokemonByType(ctx context.Context, in *GetPokemonByTypeRequest, opts ...grpc.CallOption) (*PokemonList, error) {
	out := new(PokemonList)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/GetPokemonByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) AddPokemon(ctx context.Context, in *AddPokemonRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/AddPokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) UpdatePokemon(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/UpdatePokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) DeletePokemon(ctx context.Context, in *DeletePokemonRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pokemon.PokemonService/DeletePokemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonServiceServer is the server API for PokemonService service.
// All implementations must embed UnimplementedPokemonServiceServer
// for forward compatibility
type PokemonServiceServer interface {
	GetAllPokemon(context.Context, *emptypb.Empty) (*PokemonList, error)
	GetPokemon(context.Context, *GetPokemonRequest) (*Pokemon, error)
	GetPokemonByType(context.Context, *GetPokemonByTypeRequest) (*PokemonList, error)
	AddPokemon(context.Context, *AddPokemonRequest) (*Pokemon, error)
	UpdatePokemon(context.Context, *Pokemon) (*emptypb.Empty, error)
	DeletePokemon(context.Context, *DeletePokemonRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPokemonServiceServer()
}

// UnimplementedPokemonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPokemonServiceServer struct {
}

func (UnimplementedPokemonServiceServer) GetAllPokemon(context.Context, *emptypb.Empty) (*PokemonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPokemon not implemented")
}
func (UnimplementedPokemonServiceServer) GetPokemon(context.Context, *GetPokemonRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (UnimplementedPokemonServiceServer) GetPokemonByType(context.Context, *GetPokemonByTypeRequest) (*PokemonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemonByType not implemented")
}
func (UnimplementedPokemonServiceServer) AddPokemon(context.Context, *AddPokemonRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPokemon not implemented")
}
func (UnimplementedPokemonServiceServer) UpdatePokemon(context.Context, *Pokemon) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) DeletePokemon(context.Context, *DeletePokemonRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) mustEmbedUnimplementedPokemonServiceServer() {}

// UnsafePokemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokemonServiceServer will
// result in compilation errors.
type UnsafePokemonServiceServer interface {
	mustEmbedUnimplementedPokemonServiceServer()
}

func RegisterPokemonServiceServer(s grpc.ServiceRegistrar, srv PokemonServiceServer) {
	s.RegisterService(&PokemonService_ServiceDesc, srv)
}

func _PokemonService_GetAllPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).GetAllPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/GetAllPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).GetAllPokemon(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/GetPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).GetPokemon(ctx, req.(*GetPokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_GetPokemonByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).GetPokemonByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/GetPokemonByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).GetPokemonByType(ctx, req.(*GetPokemonByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_AddPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).AddPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/AddPokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).AddPokemon(ctx, req.(*AddPokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_UpdatePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pokemon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).UpdatePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/UpdatePokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).UpdatePokemon(ctx, req.(*Pokemon))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_DeletePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).DeletePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.PokemonService/DeletePokemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).DeletePokemon(ctx, req.(*DeletePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PokemonService_ServiceDesc is the grpc.ServiceDesc for PokemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.PokemonService",
	HandlerType: (*PokemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPokemon",
			Handler:    _PokemonService_GetAllPokemon_Handler,
		},
		{
			MethodName: "GetPokemon",
			Handler:    _PokemonService_GetPokemon_Handler,
		},
		{
			MethodName: "GetPokemonByType",
			Handler:    _PokemonService_GetPokemonByType_Handler,
		},
		{
			MethodName: "AddPokemon",
			Handler:    _PokemonService_AddPokemon_Handler,
		},
		{
			MethodName: "UpdatePokemon",
			Handler:    _PokemonService_UpdatePokemon_Handler,
		},
		{
			MethodName: "DeletePokemon",
			Handler:    _PokemonService_DeletePokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon.proto",
}
