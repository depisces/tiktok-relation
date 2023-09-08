// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: friends.proto

package service

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

type DouyinRelationFriendListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // 用户鉴权token
}

func (x *DouyinRelationFriendListRequest) Reset() {
	*x = DouyinRelationFriendListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinRelationFriendListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinRelationFriendListRequest) ProtoMessage() {}

func (x *DouyinRelationFriendListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinRelationFriendListRequest.ProtoReflect.Descriptor instead.
func (*DouyinRelationFriendListRequest) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinRelationFriendListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DouyinRelationFriendListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DouyinRelationFriendListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	UserList   []*FriendUser `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`        // 用户列表
}

func (x *DouyinRelationFriendListResponse) Reset() {
	*x = DouyinRelationFriendListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinRelationFriendListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinRelationFriendListResponse) ProtoMessage() {}

func (x *DouyinRelationFriendListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinRelationFriendListResponse.ProtoReflect.Descriptor instead.
func (*DouyinRelationFriendListResponse) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinRelationFriendListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinRelationFriendListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DouyinRelationFriendListResponse) GetUserList() []*FriendUser {
	if x != nil {
		return x.UserList
	}
	return nil
}

type FriendUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // 和该好友的最新聊天消息
	MsgType int64  `protobuf:"varint,3,opt,name=msgType,proto3" json:"msgType,omitempty"` // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func (x *FriendUser) Reset() {
	*x = FriendUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friends_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendUser) ProtoMessage() {}

func (x *FriendUser) ProtoReflect() protoreflect.Message {
	mi := &file_friends_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendUser.ProtoReflect.Descriptor instead.
func (*FriendUser) Descriptor() ([]byte, []int) {
	return file_friends_proto_rawDescGZIP(), []int{2}
}

func (x *FriendUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *FriendUser) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FriendUser) GetMsgType() int64 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

var File_friends_proto protoreflect.FileDescriptor

var file_friends_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x23, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x24, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x2f, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friends_proto_rawDescOnce sync.Once
	file_friends_proto_rawDescData = file_friends_proto_rawDesc
)

func file_friends_proto_rawDescGZIP() []byte {
	file_friends_proto_rawDescOnce.Do(func() {
		file_friends_proto_rawDescData = protoimpl.X.CompressGZIP(file_friends_proto_rawDescData)
	})
	return file_friends_proto_rawDescData
}

var file_friends_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_friends_proto_goTypes = []interface{}{
	(*DouyinRelationFriendListRequest)(nil),  // 0: protos.douyin_relation_friend_list_request
	(*DouyinRelationFriendListResponse)(nil), // 1: protos.douyin_relation_friend_list_response
	(*FriendUser)(nil),                       // 2: protos.FriendUser
	(*User)(nil),                             // 3: protos.User
}
var file_friends_proto_depIdxs = []int32{
	2, // 0: protos.douyin_relation_friend_list_response.user_list:type_name -> protos.FriendUser
	3, // 1: protos.FriendUser.user:type_name -> protos.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_friends_proto_init() }
func file_friends_proto_init() {
	if File_friends_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_friends_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinRelationFriendListRequest); i {
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
		file_friends_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinRelationFriendListResponse); i {
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
		file_friends_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendUser); i {
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
			RawDescriptor: file_friends_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_friends_proto_goTypes,
		DependencyIndexes: file_friends_proto_depIdxs,
		MessageInfos:      file_friends_proto_msgTypes,
	}.Build()
	File_friends_proto = out.File
	file_friends_proto_rawDesc = nil
	file_friends_proto_goTypes = nil
	file_friends_proto_depIdxs = nil
}