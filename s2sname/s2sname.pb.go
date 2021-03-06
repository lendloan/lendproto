// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: s2sname.proto

package s2sname

import (
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

type Const int32

const (
	Const_Const_None    Const = 0
	Const_Const_Expired Const = 5 // 超时时间
)

// Enum value maps for Const.
var (
	Const_name = map[int32]string{
		0: "Const_None",
		5: "Const_Expired",
	}
	Const_value = map[string]int32{
		"Const_None":    0,
		"Const_Expired": 5,
	}
)

func (x Const) Enum() *Const {
	p := new(Const)
	*p = x
	return p
}

func (x Const) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Const) Descriptor() protoreflect.EnumDescriptor {
	return file_s2sname_proto_enumTypes[0].Descriptor()
}

func (Const) Type() protoreflect.EnumType {
	return &file_s2sname_proto_enumTypes[0]
}

func (x Const) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Const.Descriptor instead.
func (Const) EnumDescriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type S2Sname struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host    string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Prority int32  `protobuf:"varint,3,opt,name=prority,proto3" json:"prority,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Expired int64  `protobuf:"varint,5,opt,name=expired,proto3" json:"expired,omitempty"`
}

func (x *S2Sname) Reset() {
	*x = S2Sname{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2Sname) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2Sname) ProtoMessage() {}

func (x *S2Sname) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2Sname.ProtoReflect.Descriptor instead.
func (*S2Sname) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{1}
}

func (x *S2Sname) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *S2Sname) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *S2Sname) GetPrority() int32 {
	if x != nil {
		return x.Prority
	}
	return 0
}

func (x *S2Sname) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S2Sname) GetExpired() int64 {
	if x != nil {
		return x.Expired
	}
	return 0
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	S2S  *S2Sname `protobuf:"bytes,2,opt,name=s2s,proto3" json:"s2s,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterReq) GetS2S() *S2Sname {
	if x != nil {
		return x.S2S
	}
	return nil
}

type RegisterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string       `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
}

func (x *RegisterRes) Reset() {
	*x = RegisterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRes) ProtoMessage() {}

func (x *RegisterRes) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRes.ProtoReflect.Descriptor instead.
func (*RegisterRes) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *RegisterRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	S2S  *S2Sname `protobuf:"bytes,2,opt,name=s2s,proto3" json:"s2s,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateReq) GetS2S() *S2Sname {
	if x != nil {
		return x.S2S
	}
	return nil
}

type UpdateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string       `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
}

func (x *UpdateRes) Reset() {
	*x = UpdateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRes) ProtoMessage() {}

func (x *UpdateRes) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRes.ProtoReflect.Descriptor instead.
func (*UpdateRes) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *UpdateRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

type FetchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string       `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	S2Ss    []*S2Sname   `protobuf:"bytes,3,rep,name=s2ss,proto3" json:"s2ss,omitempty"`
}

func (x *FetchRes) Reset() {
	*x = FetchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRes) ProtoMessage() {}

func (x *FetchRes) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRes.ProtoReflect.Descriptor instead.
func (*FetchRes) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{6}
}

func (x *FetchRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *FetchRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *FetchRes) GetS2Ss() []*S2Sname {
	if x != nil {
		return x.S2Ss
	}
	return nil
}

type HeartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	S2S  *S2Sname `protobuf:"bytes,2,opt,name=s2s,proto3" json:"s2s,omitempty"`
}

func (x *HeartReq) Reset() {
	*x = HeartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartReq) ProtoMessage() {}

func (x *HeartReq) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartReq.ProtoReflect.Descriptor instead.
func (*HeartReq) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{7}
}

func (x *HeartReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HeartReq) GetS2S() *S2Sname {
	if x != nil {
		return x.S2S
	}
	return nil
}

type HeartRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string       `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
}

func (x *HeartRes) Reset() {
	*x = HeartRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartRes) ProtoMessage() {}

func (x *HeartRes) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartRes.ProtoReflect.Descriptor instead.
func (*HeartRes) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{8}
}

func (x *HeartRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *HeartRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

type S2SinfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2SinfoReq) Reset() {
	*x = S2SinfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2SinfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2SinfoReq) ProtoMessage() {}

func (x *S2SinfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2SinfoReq.ProtoReflect.Descriptor instead.
func (*S2SinfoReq) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{9}
}

type S2SinfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code                       `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string                             `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	S2Sinfo map[string]*S2SinfoRes_S2SnameItem `protobuf:"bytes,3,rep,name=s2sinfo,proto3" json:"s2sinfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *S2SinfoRes) Reset() {
	*x = S2SinfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2SinfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2SinfoRes) ProtoMessage() {}

func (x *S2SinfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2SinfoRes.ProtoReflect.Descriptor instead.
func (*S2SinfoRes) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{10}
}

func (x *S2SinfoRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *S2SinfoRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *S2SinfoRes) GetS2Sinfo() map[string]*S2SinfoRes_S2SnameItem {
	if x != nil {
		return x.S2Sinfo
	}
	return nil
}

type S2SinfoRes_S2SnameItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*S2Sname `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *S2SinfoRes_S2SnameItem) Reset() {
	*x = S2SinfoRes_S2SnameItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s2sname_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2SinfoRes_S2SnameItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2SinfoRes_S2SnameItem) ProtoMessage() {}

func (x *S2SinfoRes_S2SnameItem) ProtoReflect() protoreflect.Message {
	mi := &file_s2sname_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2SinfoRes_S2SnameItem.ProtoReflect.Descriptor instead.
func (*S2SinfoRes_S2SnameItem) Descriptor() ([]byte, []int) {
	return file_s2sname_proto_rawDescGZIP(), []int{10, 0}
}

func (x *S2SinfoRes_S2SnameItem) GetItem() []*S2Sname {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_s2sname_proto protoreflect.FileDescriptor

var file_s2sname_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x79,
	0x0a, 0x07, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x22, 0x45, 0x0a, 0x0b, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x03,
	0x73, 0x32, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x32, 0x73, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x03, 0x73, 0x32, 0x73,
	0x22, 0x4e, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67,
	0x22, 0x43, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x03, 0x73, 0x32, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x03, 0x73, 0x32, 0x73, 0x22, 0x4c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x6d, 0x73, 0x67, 0x22, 0x71, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67,
	0x12, 0x24, 0x0a, 0x04, 0x73, 0x32, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x73, 0x32, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x73, 0x32, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x32,
	0x73, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x03, 0x73, 0x32, 0x73, 0x22, 0x4b, 0x0a, 0x08, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x32, 0x73, 0x69, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x9b, 0x02, 0x0a, 0x0a, 0x53, 0x32, 0x73, 0x69, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x6d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x32, 0x73, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65,
	0x2e, 0x53, 0x32, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x2e, 0x53, 0x32, 0x73, 0x69,
	0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x32, 0x73, 0x69, 0x6e, 0x66,
	0x6f, 0x1a, 0x33, 0x0a, 0x0b, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x24, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x5b, 0x0a, 0x0c, 0x53, 0x32, 0x73, 0x69, 0x6e, 0x66,
	0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x53, 0x32, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x2e, 0x53, 0x32, 0x73,
	0x6e, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x2a, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x05, 0x32,
	0xe3, 0x02, 0x0a, 0x0e, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x32,
	0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x32,
	0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x32, 0x73,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x32,
	0x73, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x05, 0x48, 0x65, 0x61, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65,
	0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x32, 0x73, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x07, 0x53, 0x32, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x73, 0x32, 0x73, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x53, 0x32, 0x73, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x73, 0x32, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x32, 0x73, 0x69, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x32, 0x73, 0x6e, 0x61,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s2sname_proto_rawDescOnce sync.Once
	file_s2sname_proto_rawDescData = file_s2sname_proto_rawDesc
)

func file_s2sname_proto_rawDescGZIP() []byte {
	file_s2sname_proto_rawDescOnce.Do(func() {
		file_s2sname_proto_rawDescData = protoimpl.X.CompressGZIP(file_s2sname_proto_rawDescData)
	})
	return file_s2sname_proto_rawDescData
}

var file_s2sname_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_s2sname_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_s2sname_proto_goTypes = []interface{}{
	(Const)(0),                     // 0: s2sname.Const
	(*Request)(nil),                // 1: s2sname.Request
	(*S2Sname)(nil),                // 2: s2sname.S2sname
	(*RegisterReq)(nil),            // 3: s2sname.RegisterReq
	(*RegisterRes)(nil),            // 4: s2sname.RegisterRes
	(*UpdateReq)(nil),              // 5: s2sname.UpdateReq
	(*UpdateRes)(nil),              // 6: s2sname.UpdateRes
	(*FetchRes)(nil),               // 7: s2sname.FetchRes
	(*HeartReq)(nil),               // 8: s2sname.HeartReq
	(*HeartRes)(nil),               // 9: s2sname.HeartRes
	(*S2SinfoReq)(nil),             // 10: s2sname.S2sinfoReq
	(*S2SinfoRes)(nil),             // 11: s2sname.S2sinfoRes
	(*S2SinfoRes_S2SnameItem)(nil), // 12: s2sname.S2sinfoRes.S2snameItem
	nil,                            // 13: s2sname.S2sinfoRes.S2sinfoEntry
	(rescode.Code)(0),              // 14: rescode.Code
}
var file_s2sname_proto_depIdxs = []int32{
	2,  // 0: s2sname.RegisterReq.s2s:type_name -> s2sname.S2sname
	14, // 1: s2sname.RegisterRes.rescode:type_name -> rescode.Code
	2,  // 2: s2sname.UpdateReq.s2s:type_name -> s2sname.S2sname
	14, // 3: s2sname.UpdateRes.rescode:type_name -> rescode.Code
	14, // 4: s2sname.FetchRes.rescode:type_name -> rescode.Code
	2,  // 5: s2sname.FetchRes.s2ss:type_name -> s2sname.S2sname
	2,  // 6: s2sname.HeartReq.s2s:type_name -> s2sname.S2sname
	14, // 7: s2sname.HeartRes.rescode:type_name -> rescode.Code
	14, // 8: s2sname.S2sinfoRes.rescode:type_name -> rescode.Code
	13, // 9: s2sname.S2sinfoRes.s2sinfo:type_name -> s2sname.S2sinfoRes.S2sinfoEntry
	2,  // 10: s2sname.S2sinfoRes.S2snameItem.item:type_name -> s2sname.S2sname
	12, // 11: s2sname.S2sinfoRes.S2sinfoEntry.value:type_name -> s2sname.S2sinfoRes.S2snameItem
	3,  // 12: s2sname.S2snameService.RegisterS2sname:input_type -> s2sname.RegisterReq
	5,  // 13: s2sname.S2snameService.UpdateS2sname:input_type -> s2sname.UpdateReq
	1,  // 14: s2sname.S2snameService.FetchS2sname:input_type -> s2sname.Request
	1,  // 15: s2sname.S2snameService.FetchS2snames:input_type -> s2sname.Request
	8,  // 16: s2sname.S2snameService.Heart:input_type -> s2sname.HeartReq
	10, // 17: s2sname.S2snameService.S2sinfo:input_type -> s2sname.S2sinfoReq
	4,  // 18: s2sname.S2snameService.RegisterS2sname:output_type -> s2sname.RegisterRes
	6,  // 19: s2sname.S2snameService.UpdateS2sname:output_type -> s2sname.UpdateRes
	7,  // 20: s2sname.S2snameService.FetchS2sname:output_type -> s2sname.FetchRes
	7,  // 21: s2sname.S2snameService.FetchS2snames:output_type -> s2sname.FetchRes
	9,  // 22: s2sname.S2snameService.Heart:output_type -> s2sname.HeartRes
	11, // 23: s2sname.S2snameService.S2sinfo:output_type -> s2sname.S2sinfoRes
	18, // [18:24] is the sub-list for method output_type
	12, // [12:18] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_s2sname_proto_init() }
func file_s2sname_proto_init() {
	if File_s2sname_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s2sname_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_s2sname_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2Sname); i {
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
		file_s2sname_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_s2sname_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRes); i {
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
		file_s2sname_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_s2sname_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRes); i {
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
		file_s2sname_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRes); i {
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
		file_s2sname_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartReq); i {
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
		file_s2sname_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartRes); i {
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
		file_s2sname_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2SinfoReq); i {
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
		file_s2sname_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2SinfoRes); i {
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
		file_s2sname_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2SinfoRes_S2SnameItem); i {
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
			RawDescriptor: file_s2sname_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_s2sname_proto_goTypes,
		DependencyIndexes: file_s2sname_proto_depIdxs,
		EnumInfos:         file_s2sname_proto_enumTypes,
		MessageInfos:      file_s2sname_proto_msgTypes,
	}.Build()
	File_s2sname_proto = out.File
	file_s2sname_proto_rawDesc = nil
	file_s2sname_proto_goTypes = nil
	file_s2sname_proto_depIdxs = nil
}
