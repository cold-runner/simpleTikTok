// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: idl/RelationServer.proto

package RelationService

import (
	context "context"
	BaseResponse "github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	UserService "github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

// 用户关注/取关操作
type RelationActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                                 // 关注者id
	ToUserId   int64 `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`     // 被关注者id
	ActionType int32 `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-关注，2-取消关注
}

func (x *RelationActionRequest) Reset() {
	*x = RelationActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionRequest) ProtoMessage() {}

func (x *RelationActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionRequest.ProtoReflect.Descriptor instead.
func (*RelationActionRequest) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RelationActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *RelationActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResponse.BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *RelationActionResponse) Reset() {
	*x = RelationActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResponse) ProtoMessage() {}

func (x *RelationActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResponse.ProtoReflect.Descriptor instead.
func (*RelationActionResponse) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{2}
}

func (x *RelationActionResponse) GetBaseResp() *BaseResponse.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// 用户关注列表
type RelationFollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                  // 用户id
	ToUid int64 `protobuf:"varint,2,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"` // 被查询用户id
}

func (x *RelationFollowListRequest) Reset() {
	*x = RelationFollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListRequest) ProtoMessage() {}

func (x *RelationFollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowListRequest) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{3}
}

func (x *RelationFollowListRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RelationFollowListRequest) GetToUid() int64 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

type RelationFollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp   *BaseResponse.BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	FollowList []*UserService.User    `protobuf:"bytes,2,rep,name=follow_list,json=followList,proto3" json:"follow_list,omitempty"`
}

func (x *RelationFollowListResponse) Reset() {
	*x = RelationFollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListResponse) ProtoMessage() {}

func (x *RelationFollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowListResponse) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{4}
}

func (x *RelationFollowListResponse) GetBaseResp() *BaseResponse.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *RelationFollowListResponse) GetFollowList() []*UserService.User {
	if x != nil {
		return x.FollowList
	}
	return nil
}

// 用户粉丝列表
type RelationFollowerListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                  // 用户id
	ToUid int64 `protobuf:"varint,2,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"` // 被查询用户id
}

func (x *RelationFollowerListRequest) Reset() {
	*x = RelationFollowerListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowerListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowerListRequest) ProtoMessage() {}

func (x *RelationFollowerListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowerListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowerListRequest) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{5}
}

func (x *RelationFollowerListRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RelationFollowerListRequest) GetToUid() int64 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

type RelationFollowerListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp     *BaseResponse.BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	FollowerList []*UserService.User    `protobuf:"bytes,2,rep,name=follower_list,json=followerList,proto3" json:"follower_list,omitempty"`
}

func (x *RelationFollowerListResponse) Reset() {
	*x = RelationFollowerListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowerListResponse) ProtoMessage() {}

func (x *RelationFollowerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowerListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowerListResponse) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{6}
}

func (x *RelationFollowerListResponse) GetBaseResp() *BaseResponse.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *RelationFollowerListResponse) GetFollowerList() []*UserService.User {
	if x != nil {
		return x.FollowerList
	}
	return nil
}

// 用户关系查询
type RelationQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                             // 用户id
	ToUserId int64 `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"` // 被查询用户id
}

func (x *RelationQueryRequest) Reset() {
	*x = RelationQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationQueryRequest) ProtoMessage() {}

func (x *RelationQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationQueryRequest.ProtoReflect.Descriptor instead.
func (*RelationQueryRequest) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{7}
}

func (x *RelationQueryRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RelationQueryRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

type RelationQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResponse.BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	IsFollow bool                   `protobuf:"varint,2,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"` // true-已关注，false-未关注
}

func (x *RelationQueryResponse) Reset() {
	*x = RelationQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationQueryResponse) ProtoMessage() {}

func (x *RelationQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationQueryResponse.ProtoReflect.Descriptor instead.
func (*RelationQueryResponse) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{8}
}

func (x *RelationQueryResponse) GetBaseResp() *BaseResponse.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *RelationQueryResponse) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

// 用户信息查询关系服务，批量查询带有关注关系的用户信息
type QueryUserInfosWithRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                            // 用户id
	IdList []int64 `protobuf:"varint,2,rep,packed,name=id_list,json=idList,proto3" json:"id_list,omitempty"` // 被查询用户id列表
}

func (x *QueryUserInfosWithRelationRequest) Reset() {
	*x = QueryUserInfosWithRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserInfosWithRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserInfosWithRelationRequest) ProtoMessage() {}

func (x *QueryUserInfosWithRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserInfosWithRelationRequest.ProtoReflect.Descriptor instead.
func (*QueryUserInfosWithRelationRequest) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{9}
}

func (x *QueryUserInfosWithRelationRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *QueryUserInfosWithRelationRequest) GetIdList() []int64 {
	if x != nil {
		return x.IdList
	}
	return nil
}

type QueryUserInfosWithRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfoList []*UserService.User `protobuf:"bytes,1,rep,name=UserInfoList,proto3" json:"UserInfoList,omitempty"`
}

func (x *QueryUserInfosWithRelationResponse) Reset() {
	*x = QueryUserInfosWithRelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_RelationServer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserInfosWithRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserInfosWithRelationResponse) ProtoMessage() {}

func (x *QueryUserInfosWithRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_RelationServer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserInfosWithRelationResponse.ProtoReflect.Descriptor instead.
func (*QueryUserInfosWithRelationResponse) Descriptor() ([]byte, []int) {
	return file_idl_RelationServer_proto_rawDescGZIP(), []int{10}
}

func (x *QueryUserInfosWithRelationResponse) GetUserInfoList() []*UserService.User {
	if x != nil {
		return x.UserInfoList
	}
	return nil
}

var File_idl_RelationServer_proto protoreflect.FileDescriptor

var file_idl_RelationServer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x64, 0x6c, 0x2f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x69, 0x64, 0x6c, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x69, 0x64, 0x6c, 0x2f,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x22, 0x68, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x49, 0x0a, 0x16,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x44, 0x0a, 0x19, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x64, 0x22, 0x7a, 0x0a,
	0x1a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x0b,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0a, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x1b, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69,
	0x64, 0x22, 0x80, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x15,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x22, 0x4e, 0x0a, 0x21, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x22, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xf9, 0x03, 0x0a, 0x0f, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a,
	0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x1a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x57, 0x69, 0x74, 0x68,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x64, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x69, 0x6b, 0x54, 0x6f, 0x6b, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_RelationServer_proto_rawDescOnce sync.Once
	file_idl_RelationServer_proto_rawDescData = file_idl_RelationServer_proto_rawDesc
)

func file_idl_RelationServer_proto_rawDescGZIP() []byte {
	file_idl_RelationServer_proto_rawDescOnce.Do(func() {
		file_idl_RelationServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_RelationServer_proto_rawDescData)
	})
	return file_idl_RelationServer_proto_rawDescData
}

var file_idl_RelationServer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_idl_RelationServer_proto_goTypes = []interface{}{
	(*User)(nil),                               // 0: relation.User
	(*RelationActionRequest)(nil),              // 1: relation.RelationActionRequest
	(*RelationActionResponse)(nil),             // 2: relation.RelationActionResponse
	(*RelationFollowListRequest)(nil),          // 3: relation.RelationFollowListRequest
	(*RelationFollowListResponse)(nil),         // 4: relation.RelationFollowListResponse
	(*RelationFollowerListRequest)(nil),        // 5: relation.RelationFollowerListRequest
	(*RelationFollowerListResponse)(nil),       // 6: relation.RelationFollowerListResponse
	(*RelationQueryRequest)(nil),               // 7: relation.RelationQueryRequest
	(*RelationQueryResponse)(nil),              // 8: relation.RelationQueryResponse
	(*QueryUserInfosWithRelationRequest)(nil),  // 9: relation.QueryUserInfosWithRelationRequest
	(*QueryUserInfosWithRelationResponse)(nil), // 10: relation.QueryUserInfosWithRelationResponse
	(*BaseResponse.BaseResp)(nil),              // 11: baseResp.BaseResp
	(*UserService.User)(nil),                   // 12: user.User
}
var file_idl_RelationServer_proto_depIdxs = []int32{
	11, // 0: relation.RelationActionResponse.base_resp:type_name -> baseResp.BaseResp
	11, // 1: relation.RelationFollowListResponse.base_resp:type_name -> baseResp.BaseResp
	12, // 2: relation.RelationFollowListResponse.follow_list:type_name -> user.User
	11, // 3: relation.RelationFollowerListResponse.base_resp:type_name -> baseResp.BaseResp
	12, // 4: relation.RelationFollowerListResponse.follower_list:type_name -> user.User
	11, // 5: relation.RelationQueryResponse.base_resp:type_name -> baseResp.BaseResp
	12, // 6: relation.QueryUserInfosWithRelationResponse.UserInfoList:type_name -> user.User
	1,  // 7: relation.RelationService.RelationAction:input_type -> relation.RelationActionRequest
	3,  // 8: relation.RelationService.GetFollowList:input_type -> relation.RelationFollowListRequest
	5,  // 9: relation.RelationService.GetFollowerList:input_type -> relation.RelationFollowerListRequest
	7,  // 10: relation.RelationService.QueryRelation:input_type -> relation.RelationQueryRequest
	9,  // 11: relation.RelationService.QueryUserInfosWithRelation:input_type -> relation.QueryUserInfosWithRelationRequest
	2,  // 12: relation.RelationService.RelationAction:output_type -> relation.RelationActionResponse
	4,  // 13: relation.RelationService.GetFollowList:output_type -> relation.RelationFollowListResponse
	6,  // 14: relation.RelationService.GetFollowerList:output_type -> relation.RelationFollowerListResponse
	8,  // 15: relation.RelationService.QueryRelation:output_type -> relation.RelationQueryResponse
	10, // 16: relation.RelationService.QueryUserInfosWithRelation:output_type -> relation.QueryUserInfosWithRelationResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_idl_RelationServer_proto_init() }
func file_idl_RelationServer_proto_init() {
	if File_idl_RelationServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_RelationServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_idl_RelationServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionRequest); i {
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
		file_idl_RelationServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResponse); i {
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
		file_idl_RelationServer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListRequest); i {
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
		file_idl_RelationServer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListResponse); i {
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
		file_idl_RelationServer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowerListRequest); i {
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
		file_idl_RelationServer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowerListResponse); i {
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
		file_idl_RelationServer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationQueryRequest); i {
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
		file_idl_RelationServer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationQueryResponse); i {
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
		file_idl_RelationServer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserInfosWithRelationRequest); i {
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
		file_idl_RelationServer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserInfosWithRelationResponse); i {
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
			RawDescriptor: file_idl_RelationServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_RelationServer_proto_goTypes,
		DependencyIndexes: file_idl_RelationServer_proto_depIdxs,
		MessageInfos:      file_idl_RelationServer_proto_msgTypes,
	}.Build()
	File_idl_RelationServer_proto = out.File
	file_idl_RelationServer_proto_rawDesc = nil
	file_idl_RelationServer_proto_goTypes = nil
	file_idl_RelationServer_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.6.2. DO NOT EDIT.

type RelationService interface {
	RelationAction(ctx context.Context, req *RelationActionRequest) (res *RelationActionResponse, err error)
	GetFollowList(ctx context.Context, req *RelationFollowListRequest) (res *RelationFollowListResponse, err error)
	GetFollowerList(ctx context.Context, req *RelationFollowerListRequest) (res *RelationFollowerListResponse, err error)
	QueryRelation(ctx context.Context, req *RelationQueryRequest) (res *RelationQueryResponse, err error)
	QueryUserInfosWithRelation(ctx context.Context, req *QueryUserInfosWithRelationRequest) (res *QueryUserInfosWithRelationResponse, err error)
}
