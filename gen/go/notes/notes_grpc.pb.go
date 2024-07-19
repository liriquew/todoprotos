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
	CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*NoteMeta, error)
	GetNoteByID(ctx context.Context, in *NoteMeta, opts ...grpc.CallOption) (*Note, error)
	UpdateNoteByID(ctx context.Context, in *NoteWithID, opts ...grpc.CallOption) (*NoteMeta, error)
	DeleteNotebyID(ctx context.Context, in *NoteMeta, opts ...grpc.CallOption) (*NoteMeta, error)
}

type notesClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesClient(cc grpc.ClientConnInterface) NotesClient {
	return &notesClient{cc}
}

func (c *notesClient) CreateNote(ctx context.Context, in *Note, opts ...grpc.CallOption) (*NoteMeta, error) {
	out := new(NoteMeta)
	err := c.cc.Invoke(ctx, "/notes.Notes/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) GetNoteByID(ctx context.Context, in *NoteMeta, opts ...grpc.CallOption) (*Note, error) {
	out := new(Note)
	err := c.cc.Invoke(ctx, "/notes.Notes/GetNoteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) UpdateNoteByID(ctx context.Context, in *NoteWithID, opts ...grpc.CallOption) (*NoteMeta, error) {
	out := new(NoteMeta)
	err := c.cc.Invoke(ctx, "/notes.Notes/UpdateNoteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notesClient) DeleteNotebyID(ctx context.Context, in *NoteMeta, opts ...grpc.CallOption) (*NoteMeta, error) {
	out := new(NoteMeta)
	err := c.cc.Invoke(ctx, "/notes.Notes/DeleteNotebyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotesServer is the server API for Notes service.
// All implementations must embed UnimplementedNotesServer
// for forward compatibility
type NotesServer interface {
	CreateNote(context.Context, *Note) (*NoteMeta, error)
	GetNoteByID(context.Context, *NoteMeta) (*Note, error)
	UpdateNoteByID(context.Context, *NoteWithID) (*NoteMeta, error)
	DeleteNotebyID(context.Context, *NoteMeta) (*NoteMeta, error)
	mustEmbedUnimplementedNotesServer()
}

// UnimplementedNotesServer must be embedded to have forward compatible implementations.
type UnimplementedNotesServer struct {
}

func (UnimplementedNotesServer) CreateNote(context.Context, *Note) (*NoteMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNotesServer) GetNoteByID(context.Context, *NoteMeta) (*Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteByID not implemented")
}
func (UnimplementedNotesServer) UpdateNoteByID(context.Context, *NoteWithID) (*NoteMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNoteByID not implemented")
}
func (UnimplementedNotesServer) DeleteNotebyID(context.Context, *NoteMeta) (*NoteMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotebyID not implemented")
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
	in := new(Note)
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
		return srv.(NotesServer).CreateNote(ctx, req.(*Note))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_GetNoteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).GetNoteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/GetNoteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).GetNoteByID(ctx, req.(*NoteMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_UpdateNoteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteWithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).UpdateNoteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/UpdateNoteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).UpdateNoteByID(ctx, req.(*NoteWithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notes_DeleteNotebyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoteMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotesServer).DeleteNotebyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notes.Notes/DeleteNotebyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotesServer).DeleteNotebyID(ctx, req.(*NoteMeta))
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
			MethodName: "GetNoteByID",
			Handler:    _Notes_GetNoteByID_Handler,
		},
		{
			MethodName: "UpdateNoteByID",
			Handler:    _Notes_UpdateNoteByID_Handler,
		},
		{
			MethodName: "DeleteNotebyID",
			Handler:    _Notes_DeleteNotebyID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notes/notes.proto",
}
