// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: friends/friends.proto

package friends

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID      int64 `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	FriendID int64 `protobuf:"varint,2,opt,name=FriendID,proto3" json:"FriendID,omitempty"`
}

func (x *FriendRequest) Reset() {
	*x = FriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_friends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequest) ProtoMessage() {}

func (x *FriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friends_friends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequest.ProtoReflect.Descriptor instead.
func (*FriendRequest) Descriptor() ([]byte, []int) {
	return file_friends_friends_proto_rawDescGZIP(), []int{0}
}

func (x *FriendRequest) GetUID() int64 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *FriendRequest) GetFriendID() int64 {
	if x != nil {
		return x.FriendID
	}
	return 0
}

type FriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FriendResponse) Reset() {
	*x = FriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_friends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendResponse) ProtoMessage() {}

func (x *FriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friends_friends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendResponse.ProtoReflect.Descriptor instead.
func (*FriendResponse) Descriptor() ([]byte, []int) {
	return file_friends_friends_proto_rawDescGZIP(), []int{1}
}

type ListFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID int64 `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
}

func (x *ListFriendRequest) Reset() {
	*x = ListFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_friends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendRequest) ProtoMessage() {}

func (x *ListFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friends_friends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendRequest.ProtoReflect.Descriptor instead.
func (*ListFriendRequest) Descriptor() ([]byte, []int) {
	return file_friends_friends_proto_rawDescGZIP(), []int{2}
}

func (x *ListFriendRequest) GetUID() int64 {
	if x != nil {
		return x.UID
	}
	return 0
}

type ListFriendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendIDs []int64 `protobuf:"varint,1,rep,packed,name=FriendIDs,proto3" json:"FriendIDs,omitempty"`
}

func (x *ListFriendResponse) Reset() {
	*x = ListFriendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_friends_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendResponse) ProtoMessage() {}

func (x *ListFriendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friends_friends_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendResponse.ProtoReflect.Descriptor instead.
func (*ListFriendResponse) Descriptor() ([]byte, []int) {
	return file_friends_friends_proto_rawDescGZIP(), []int{3}
}

func (x *ListFriendResponse) GetFriendIDs() []int64 {
	if x != nil {
		return x.FriendIDs
	}
	return nil
}

var File_friends_friends_proto protoreflect.FileDescriptor

var file_friends_friends_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x22, 0x3d, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x22,
	0x10, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x25, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x09, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x73, 0x32, 0xd0, 0x01, 0x0a,
	0x07, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x3c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x18, 0x5a, 0x16, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x3b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_friends_friends_proto_rawDescOnce sync.Once
	file_friends_friends_proto_rawDescData = file_friends_friends_proto_rawDesc
)

func file_friends_friends_proto_rawDescGZIP() []byte {
	file_friends_friends_proto_rawDescOnce.Do(func() {
		file_friends_friends_proto_rawDescData = protoimpl.X.CompressGZIP(file_friends_friends_proto_rawDescData)
	})
	return file_friends_friends_proto_rawDescData
}

var file_friends_friends_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_friends_friends_proto_goTypes = []interface{}{
	(*FriendRequest)(nil),      // 0: friends.FriendRequest
	(*FriendResponse)(nil),     // 1: friends.FriendResponse
	(*ListFriendRequest)(nil),  // 2: friends.ListFriendRequest
	(*ListFriendResponse)(nil), // 3: friends.ListFriendResponse
}
var file_friends_friends_proto_depIdxs = []int32{
	0, // 0: friends.Friends.AddFriend:input_type -> friends.FriendRequest
	0, // 1: friends.Friends.RemoveFriend:input_type -> friends.FriendRequest
	2, // 2: friends.Friends.ListFriends:input_type -> friends.ListFriendRequest
	1, // 3: friends.Friends.AddFriend:output_type -> friends.FriendResponse
	1, // 4: friends.Friends.RemoveFriend:output_type -> friends.FriendResponse
	3, // 5: friends.Friends.ListFriends:output_type -> friends.ListFriendResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_friends_friends_proto_init() }
func file_friends_friends_proto_init() {
	if File_friends_friends_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friends_friends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friends_friends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friends_friends_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_friends_friends_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFriendResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_friends_friends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friends_friends_proto_goTypes,
		DependencyIndexes: file_friends_friends_proto_depIdxs,
		MessageInfos:      file_friends_friends_proto_msgTypes,
	}.Build()
	File_friends_friends_proto = out.File
	file_friends_friends_proto_rawDesc = nil
	file_friends_friends_proto_goTypes = nil
	file_friends_friends_proto_depIdxs = nil
}
