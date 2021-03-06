// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: friendnode.proto

package friendnode

import (
	common "github.com/lendloan/lendproto/common"
	rescode "github.com/lendloan/lendproto/rescode"
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

type AddFriendsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid1  int64             `protobuf:"varint,1,opt,name=uid1,proto3" json:"uid1,omitempty"`
	Uid2  int64             `protobuf:"varint,2,opt,name=uid2,proto3" json:"uid2,omitempty"`
	Extra map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AddFriendsReq) Reset() {
	*x = AddFriendsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friendnode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendsReq) ProtoMessage() {}

func (x *AddFriendsReq) ProtoReflect() protoreflect.Message {
	mi := &file_friendnode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendsReq.ProtoReflect.Descriptor instead.
func (*AddFriendsReq) Descriptor() ([]byte, []int) {
	return file_friendnode_proto_rawDescGZIP(), []int{0}
}

func (x *AddFriendsReq) GetUid1() int64 {
	if x != nil {
		return x.Uid1
	}
	return 0
}

func (x *AddFriendsReq) GetUid2() int64 {
	if x != nil {
		return x.Uid2
	}
	return 0
}

func (x *AddFriendsReq) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type AddFriendsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string            `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Extra   map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AddFriendsRes) Reset() {
	*x = AddFriendsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friendnode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendsRes) ProtoMessage() {}

func (x *AddFriendsRes) ProtoReflect() protoreflect.Message {
	mi := &file_friendnode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendsRes.ProtoReflect.Descriptor instead.
func (*AddFriendsRes) Descriptor() ([]byte, []int) {
	return file_friendnode_proto_rawDescGZIP(), []int{1}
}

func (x *AddFriendsRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code(0)
}

func (x *AddFriendsRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *AddFriendsRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type FriendsCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64             `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Extra map[string]string `protobuf:"bytes,2,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FriendsCountReq) Reset() {
	*x = FriendsCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friendnode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendsCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendsCountReq) ProtoMessage() {}

func (x *FriendsCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_friendnode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendsCountReq.ProtoReflect.Descriptor instead.
func (*FriendsCountReq) Descriptor() ([]byte, []int) {
	return file_friendnode_proto_rawDescGZIP(), []int{2}
}

func (x *FriendsCountReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FriendsCountReq) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type FriendsCountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string            `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Count   int32             `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Extra   map[string]string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FriendsCountRes) Reset() {
	*x = FriendsCountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friendnode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendsCountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendsCountRes) ProtoMessage() {}

func (x *FriendsCountRes) ProtoReflect() protoreflect.Message {
	mi := &file_friendnode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendsCountRes.ProtoReflect.Descriptor instead.
func (*FriendsCountRes) Descriptor() ([]byte, []int) {
	return file_friendnode_proto_rawDescGZIP(), []int{3}
}

func (x *FriendsCountRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code(0)
}

func (x *FriendsCountRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *FriendsCountRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FriendsCountRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type GetFriendsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64             `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Extra map[string]string `protobuf:"bytes,2,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetFriendsReq) Reset() {
	*x = GetFriendsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friendnode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendsReq) ProtoMessage() {}

func (x *GetFriendsReq) ProtoReflect() protoreflect.Message {
	mi := &file_friendnode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendsReq.ProtoReflect.Descriptor instead.
func (*GetFriendsReq) Descriptor() ([]byte, []int) {
	return file_friendnode_proto_rawDescGZIP(), []int{4}
}

func (x *GetFriendsReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetFriendsReq) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type GetFriendsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code       `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string             `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Friends []*common.UserInfo `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty"`
	Extra   map[string]string  `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetFriendsRes) Reset() {
	*x = GetFriendsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friendnode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendsRes) ProtoMessage() {}

func (x *GetFriendsRes) ProtoReflect() protoreflect.Message {
	mi := &file_friendnode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendsRes.ProtoReflect.Descriptor instead.
func (*GetFriendsRes) Descriptor() ([]byte, []int) {
	return file_friendnode_proto_rawDescGZIP(), []int{5}
}

func (x *GetFriendsRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code(0)
}

func (x *GetFriendsRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *GetFriendsRes) GetFriends() []*common.UserInfo {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *GetFriendsRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_friendnode_proto protoreflect.FileDescriptor

var file_friendnode_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x15,
	0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x69, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x69, 0x64, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x75, 0x69, 0x64, 0x32, 0x12, 0x3a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc6, 0x01, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x12, 0x3a, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x9b, 0x01, 0x0a, 0x0f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xe0, 0x01, 0x0a, 0x0f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x97, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf2,
	0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73,
	0x67, 0x12, 0x2a, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x3a, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0xeb, 0x01, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x0c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1b, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friendnode_proto_rawDescOnce sync.Once
	file_friendnode_proto_rawDescData = file_friendnode_proto_rawDesc
)

func file_friendnode_proto_rawDescGZIP() []byte {
	file_friendnode_proto_rawDescOnce.Do(func() {
		file_friendnode_proto_rawDescData = protoimpl.X.CompressGZIP(file_friendnode_proto_rawDescData)
	})
	return file_friendnode_proto_rawDescData
}

var file_friendnode_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_friendnode_proto_goTypes = []interface{}{
	(*AddFriendsReq)(nil),   // 0: friendnode.AddFriendsReq
	(*AddFriendsRes)(nil),   // 1: friendnode.AddFriendsRes
	(*FriendsCountReq)(nil), // 2: friendnode.FriendsCountReq
	(*FriendsCountRes)(nil), // 3: friendnode.FriendsCountRes
	(*GetFriendsReq)(nil),   // 4: friendnode.GetFriendsReq
	(*GetFriendsRes)(nil),   // 5: friendnode.GetFriendsRes
	nil,                     // 6: friendnode.AddFriendsReq.ExtraEntry
	nil,                     // 7: friendnode.AddFriendsRes.ExtraEntry
	nil,                     // 8: friendnode.FriendsCountReq.ExtraEntry
	nil,                     // 9: friendnode.FriendsCountRes.ExtraEntry
	nil,                     // 10: friendnode.GetFriendsReq.ExtraEntry
	nil,                     // 11: friendnode.GetFriendsRes.ExtraEntry
	(rescode.Code)(0),       // 12: rescode.Code
	(*common.UserInfo)(nil), // 13: common.UserInfo
}
var file_friendnode_proto_depIdxs = []int32{
	6,  // 0: friendnode.AddFriendsReq.extra:type_name -> friendnode.AddFriendsReq.ExtraEntry
	12, // 1: friendnode.AddFriendsRes.rescode:type_name -> rescode.Code
	7,  // 2: friendnode.AddFriendsRes.extra:type_name -> friendnode.AddFriendsRes.ExtraEntry
	8,  // 3: friendnode.FriendsCountReq.extra:type_name -> friendnode.FriendsCountReq.ExtraEntry
	12, // 4: friendnode.FriendsCountRes.rescode:type_name -> rescode.Code
	9,  // 5: friendnode.FriendsCountRes.extra:type_name -> friendnode.FriendsCountRes.ExtraEntry
	10, // 6: friendnode.GetFriendsReq.extra:type_name -> friendnode.GetFriendsReq.ExtraEntry
	12, // 7: friendnode.GetFriendsRes.rescode:type_name -> rescode.Code
	13, // 8: friendnode.GetFriendsRes.friends:type_name -> common.UserInfo
	11, // 9: friendnode.GetFriendsRes.extra:type_name -> friendnode.GetFriendsRes.ExtraEntry
	0,  // 10: friendnode.FriendnodeService.AddFriends:input_type -> friendnode.AddFriendsReq
	2,  // 11: friendnode.FriendnodeService.FriendsCount:input_type -> friendnode.FriendsCountReq
	4,  // 12: friendnode.FriendnodeService.GetFriends:input_type -> friendnode.GetFriendsReq
	1,  // 13: friendnode.FriendnodeService.AddFriends:output_type -> friendnode.AddFriendsRes
	3,  // 14: friendnode.FriendnodeService.FriendsCount:output_type -> friendnode.FriendsCountRes
	5,  // 15: friendnode.FriendnodeService.GetFriends:output_type -> friendnode.GetFriendsRes
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_friendnode_proto_init() }
func file_friendnode_proto_init() {
	if File_friendnode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friendnode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendsReq); i {
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
		file_friendnode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendsRes); i {
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
		file_friendnode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendsCountReq); i {
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
		file_friendnode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendsCountRes); i {
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
		file_friendnode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendsReq); i {
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
		file_friendnode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendsRes); i {
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
			RawDescriptor: file_friendnode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friendnode_proto_goTypes,
		DependencyIndexes: file_friendnode_proto_depIdxs,
		MessageInfos:      file_friendnode_proto_msgTypes,
	}.Build()
	File_friendnode_proto = out.File
	file_friendnode_proto_rawDesc = nil
	file_friendnode_proto_goTypes = nil
	file_friendnode_proto_depIdxs = nil
}
