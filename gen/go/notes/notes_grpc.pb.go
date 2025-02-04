// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: notes/notes.proto

package notes

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

// NotesClient is the client API for Notes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotesClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*NoteResponse, error)
	GetNote(ctx context.Context, in *NoteIDRequest, opts ...grpc.CallOption) (*Note, error)
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*NoteResponse, error)
	DeleteNote(ctx context.Context, in *NoteIDRequest, opts ...grpc.CallOption) (*NoteResponse, error)
	ListUserNotesID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*NoteIDList, error)
	ListUserNotes(ctx context.Context, in *NoteIDList, opts ...grpc.CallOption) (*NoteList, error)
	ListUsersNotes(ctx context.Context, in *UsersNotesRequest, opts ...grpc.CallOption) (*UsersNotesList, error)
}

type notesClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesClient(cc grpc.ClientConnInterface) NotesClient {
	return &notesClient{cc}
}

func (c *notesClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*NoteResponse, error) {
	out := new(NoteResponse)
	err := c.cc.Invoke(ctx, "/notes.Notes/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) GetNote(ctx context.Context, in *NoteIDRequest, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/GetNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*NoteResponse, error) {
	out := new(NoteResponse)
	err := c.cc.Invoke(ctx, "/notes.Notes/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) DeleteNote(ctx context.Context, in *NoteIDRequest, opts ...grpc.CallOption) (*NoteResponse, error) {
	out := new(NoteResponse)
	err := c.cc.Invoke(ctx, "/notes.Notes/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) ListUserNotesID(ctx context.Context, in *UserIDRequest, opts ...grpc.CallOption) (*NoteIDList, error) {
	out := new(NoteIDList)
	err := c.cc.Invoke(ctx, "/notes.Notes/ListUserNotesID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) ListUserNotes(ctx context.Context, in *NoteIDList, opts ...grpc.CallOption) (*NoteList, error) {
	out := new(NoteList)
	err := c.cc.Invoke(ctx, "/notes.Notes/ListUserNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) ListUsersNotes(ctx context.Context, in *UsersNotesRequest, opts ...grpc.CallOption) (*UsersNotesList, error) {
	out := new(UsersNotesList)
	err := c.cc.Invoke(ctx, "/notes.Notes/ListUsersNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServer is the server API for Notes service.
// All implementations must embed UnimplementedNotesServer
// for forward compatibility
type NotesServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*NoteResponse, error)
	GetNote(context.Context, *NoteIDRequest) (*Note, error)
	UpdateNote(context.Context, *UpdateNoteRequest) (*NoteResponse, error)
	DeleteNote(context.Context, *NoteIDRequest) (*NoteResponse, error)
	ListUserNotesID(context.Context, *UserIDRequest) (*NoteIDList, error)
	ListUserNotes(context.Context, *NoteIDList) (*NoteList, error)
	ListUsersNotes(context.Context, *UsersNotesRequest) (*UsersNotesList, error)
	mustEmbedUnimplementedNotesServer()
}

// UnimplementedNotesServer must be embedded to have forward compatible implementations.
type UnimplementedNotesServer struct {
}

func (UnimplementedNotesServer) CreateNote(context.Context, *CreateNoteRequest) (*NoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNotesServer) GetNote(context.Context, *NoteIDRequest) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNote not implemented")
}
func (UnimplementedNotesServer) UpdateNote(context.Context, *UpdateNoteRequest) (*NoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedNotesServer) DeleteNote(context.Context, *NoteIDRequest) (*NoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedNotesServer) ListUserNotesID(context.Context, *UserIDRequest) (*NoteIDList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserNotesID not implemented")
}
func (UnimplementedNotesServer) ListUserNotes(context.Context, *NoteIDList) (*NoteList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserNotes not implemented")
}
func (UnimplementedNotesServer) ListUsersNotes(context.Context, *UsersNotesRequest) (*UsersNotesList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersNotes not implemented")
}
func (UnimplementedNotesServer) mustEmbedUnimplementedNotesServer() {}

// UnsafeNotesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotesServer will
// result in compilation errors.
type UnsafeNotesServer interface {
	mustEmbedUnimplementedNotesServer()
}

func RegisterNotesServer(s grpc.ServiceRegistrar, srv NotesServer) {
	s.RegisterService(&Notes_ServiceDesc, srv)
}

func _Notes_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_GetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).GetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/GetNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).GetNote(ctx, req.(*NoteIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).UpdateNote(ctx, req.(*UpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).DeleteNote(ctx, req.(*NoteIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_ListUserNotesID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).ListUserNotesID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/ListUserNotesID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).ListUserNotesID(ctx, req.(*UserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_ListUserNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteIDList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).ListUserNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/ListUserNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).ListUserNotes(ctx, req.(*NoteIDList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_ListUsersNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).ListUsersNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/ListUsersNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).ListUsersNotes(ctx, req.(*UsersNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notes_ServiceDesc is the grpc.ServiceDesc for Notes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notes.Notes",
	HandlerType: (*NotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _Notes_CreateNote_Handler,
		},
		{
			MethodName: "GetNote",
			Handler:    _Notes_GetNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _Notes_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _Notes_DeleteNote_Handler,
		},
		{
			MethodName: "ListUserNotesID",
			Handler:    _Notes_ListUserNotesID_Handler,
		},
		{
			MethodName: "ListUserNotes",
			Handler:    _Notes_ListUserNotes_Handler,
		},
		{
			MethodName: "ListUsersNotes",
			Handler:    _Notes_ListUsersNotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notes/notes.proto",
}
