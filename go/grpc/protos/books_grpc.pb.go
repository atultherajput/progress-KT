// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: books.proto

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

// BooksClient is the client API for Books service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksClient interface {
	//CRUD for Book service
	GetAllBooks(ctx context.Context, in *NoParamRequest, opts ...grpc.CallOption) (*BooksResponse, error)
	GetBook(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*BookRequestResponse, error)
	AddBook(ctx context.Context, in *BookRequestResponse, opts ...grpc.CallOption) (*TextResponse, error)
	DeleteBook(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*TextResponse, error)
	UpdateBook(ctx context.Context, in *BookUpdateRequest, opts ...grpc.CallOption) (*TextResponse, error)
}

type booksClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksClient(cc grpc.ClientConnInterface) BooksClient {
	return &booksClient{cc}
}

func (c *booksClient) GetAllBooks(ctx context.Context, in *NoParamRequest, opts ...grpc.CallOption) (*BooksResponse, error) {
	out := new(BooksResponse)
	err := c.cc.Invoke(ctx, "/protos.Books/GetAllBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksClient) GetBook(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*BookRequestResponse, error) {
	out := new(BookRequestResponse)
	err := c.cc.Invoke(ctx, "/protos.Books/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksClient) AddBook(ctx context.Context, in *BookRequestResponse, opts ...grpc.CallOption) (*TextResponse, error) {
	out := new(TextResponse)
	err := c.cc.Invoke(ctx, "/protos.Books/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksClient) DeleteBook(ctx context.Context, in *BookIdRequest, opts ...grpc.CallOption) (*TextResponse, error) {
	out := new(TextResponse)
	err := c.cc.Invoke(ctx, "/protos.Books/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksClient) UpdateBook(ctx context.Context, in *BookUpdateRequest, opts ...grpc.CallOption) (*TextResponse, error) {
	out := new(TextResponse)
	err := c.cc.Invoke(ctx, "/protos.Books/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServer is the server API for Books service.
// All implementations must embed UnimplementedBooksServer
// for forward compatibility
type BooksServer interface {
	//CRUD for Book service
	GetAllBooks(context.Context, *NoParamRequest) (*BooksResponse, error)
	GetBook(context.Context, *BookIdRequest) (*BookRequestResponse, error)
	AddBook(context.Context, *BookRequestResponse) (*TextResponse, error)
	DeleteBook(context.Context, *BookIdRequest) (*TextResponse, error)
	UpdateBook(context.Context, *BookUpdateRequest) (*TextResponse, error)
	mustEmbedUnimplementedBooksServer()
}

// UnimplementedBooksServer must be embedded to have forward compatible implementations.
type UnimplementedBooksServer struct {
}

func (UnimplementedBooksServer) GetAllBooks(context.Context, *NoParamRequest) (*BooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedBooksServer) GetBook(context.Context, *BookIdRequest) (*BookRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBooksServer) AddBook(context.Context, *BookRequestResponse) (*TextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedBooksServer) DeleteBook(context.Context, *BookIdRequest) (*TextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBooksServer) UpdateBook(context.Context, *BookUpdateRequest) (*TextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBooksServer) mustEmbedUnimplementedBooksServer() {}

// UnsafeBooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServer will
// result in compilation errors.
type UnsafeBooksServer interface {
	mustEmbedUnimplementedBooksServer()
}

func RegisterBooksServer(s grpc.ServiceRegistrar, srv BooksServer) {
	s.RegisterService(&Books_ServiceDesc, srv)
}

func _Books_GetAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).GetAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Books/GetAllBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).GetAllBooks(ctx, req.(*NoParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Books_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Books/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).GetBook(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Books_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequestResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Books/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).AddBook(ctx, req.(*BookRequestResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Books_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Books/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).DeleteBook(ctx, req.(*BookIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Books_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Books/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServer).UpdateBook(ctx, req.(*BookUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Books_ServiceDesc is the grpc.ServiceDesc for Books service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Books_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Books",
	HandlerType: (*BooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllBooks",
			Handler:    _Books_GetAllBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Books_GetBook_Handler,
		},
		{
			MethodName: "AddBook",
			Handler:    _Books_AddBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Books_DeleteBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _Books_UpdateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books.proto",
}
