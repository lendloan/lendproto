// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: codenode.proto

package codenode

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

type CodeType int32

const (
	CodeType_UID     CodeType = 0
	CodeType_ACCOUNT CodeType = 1
	CodeType_MOBILE  CodeType = 2
	CodeType_EMAIL   CodeType = 3
)

// Enum value maps for CodeType.
var (
	CodeType_name = map[int32]string{
		0: "UID",
		1: "ACCOUNT",
		2: "MOBILE",
		3: "EMAIL",
	}
	CodeType_value = map[string]int32{
		"UID":     0,
		"ACCOUNT": 1,
		"MOBILE":  2,
		"EMAIL":   3,
	}
)

func (x CodeType) Enum() *CodeType {
	p := new(CodeType)
	*p = x
	return p
}

func (x CodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_codenode_proto_enumTypes[0].Descriptor()
}

func (CodeType) Type() protoreflect.EnumType {
	return &file_codenode_proto_enumTypes[0]
}

func (x CodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeType.Descriptor instead.
func (CodeType) EnumDescriptor() ([]byte, []int) {
	return file_codenode_proto_rawDescGZIP(), []int{0}
}

// 请求验证码的操作类型
type OperateType int32

const (
	OperateType_NONE     OperateType = 0
	OperateType_REGISTER OperateType = 101 // 注册请求
	OperateType_LOGIN    OperateType = 102 // 登录请求
)

// Enum value maps for OperateType.
var (
	OperateType_name = map[int32]string{
		0:   "NONE",
		101: "REGISTER",
		102: "LOGIN",
	}
	OperateType_value = map[string]int32{
		"NONE":     0,
		"REGISTER": 101,
		"LOGIN":    102,
	}
)

func (x OperateType) Enum() *OperateType {
	p := new(OperateType)
	*p = x
	return p
}

func (x OperateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperateType) Descriptor() protoreflect.EnumDescriptor {
	return file_codenode_proto_enumTypes[1].Descriptor()
}

func (OperateType) Type() protoreflect.EnumType {
	return &file_codenode_proto_enumTypes[1]
}

func (x OperateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperateType.Descriptor instead.
func (OperateType) EnumDescriptor() ([]byte, []int) {
	return file_codenode_proto_rawDescGZIP(), []int{1}
}

type CodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth   *common.Authorize `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Desc   string            `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Type   CodeType          `protobuf:"varint,3,opt,name=type,proto3,enum=codenode.CodeType" json:"type,omitempty"`
	Optype OperateType       `protobuf:"varint,4,opt,name=optype,proto3,enum=codenode.OperateType" json:"optype,omitempty"`
	Extra  map[string]string `protobuf:"bytes,5,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CodeReq) Reset() {
	*x = CodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codenode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeReq) ProtoMessage() {}

func (x *CodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_codenode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeReq.ProtoReflect.Descriptor instead.
func (*CodeReq) Descriptor() ([]byte, []int) {
	return file_codenode_proto_rawDescGZIP(), []int{0}
}

func (x *CodeReq) GetAuth() *common.Authorize {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *CodeReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CodeReq) GetType() CodeType {
	if x != nil {
		return x.Type
	}
	return CodeType_UID
}

func (x *CodeReq) GetOptype() OperateType {
	if x != nil {
		return x.Optype
	}
	return OperateType_NONE
}

func (x *CodeReq) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type CodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string            `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Code    string            `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Extra   map[string]string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CodeRes) Reset() {
	*x = CodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codenode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRes) ProtoMessage() {}

func (x *CodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_codenode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRes.ProtoReflect.Descriptor instead.
func (*CodeRes) Descriptor() ([]byte, []int) {
	return file_codenode_proto_rawDescGZIP(), []int{1}
}

func (x *CodeRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *CodeRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *CodeRes) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CodeRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type VerifyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth   *common.Authorize `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Desc   string            `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Code   string            `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Optype OperateType       `protobuf:"varint,4,opt,name=optype,proto3,enum=codenode.OperateType" json:"optype,omitempty"`
	Extra  map[string]string `protobuf:"bytes,5,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VerifyCodeReq) Reset() {
	*x = VerifyCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codenode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeReq) ProtoMessage() {}

func (x *VerifyCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_codenode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeReq.ProtoReflect.Descriptor instead.
func (*VerifyCodeReq) Descriptor() ([]byte, []int) {
	return file_codenode_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyCodeReq) GetAuth() *common.Authorize {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *VerifyCodeReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *VerifyCodeReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyCodeReq) GetOptype() OperateType {
	if x != nil {
		return x.Optype
	}
	return OperateType_NONE
}

func (x *VerifyCodeReq) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type VerifyCodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rescode rescode.Code      `protobuf:"varint,1,opt,name=rescode,proto3,enum=rescode.Code" json:"rescode,omitempty"`
	Resmsg  string            `protobuf:"bytes,2,opt,name=resmsg,proto3" json:"resmsg,omitempty"`
	Extra   map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VerifyCodeRes) Reset() {
	*x = VerifyCodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codenode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeRes) ProtoMessage() {}

func (x *VerifyCodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_codenode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeRes.ProtoReflect.Descriptor instead.
func (*VerifyCodeRes) Descriptor() ([]byte, []int) {
	return file_codenode_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyCodeRes) GetRescode() rescode.Code {
	if x != nil {
		return x.Rescode
	}
	return rescode.Code_code_SUCCESS
}

func (x *VerifyCodeRes) GetResmsg() string {
	if x != nil {
		return x.Resmsg
	}
	return ""
}

func (x *VerifyCodeRes) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_codenode_proto protoreflect.FileDescriptor

var file_codenode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x26, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x70,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xcc, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x27,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x81, 0x02, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x38, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc4, 0x01, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x37, 0x0a, 0x08,
	0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0x03, 0x2a, 0x30, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x66, 0x32, 0x87, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65,
	0x6e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codenode_proto_rawDescOnce sync.Once
	file_codenode_proto_rawDescData = file_codenode_proto_rawDesc
)

func file_codenode_proto_rawDescGZIP() []byte {
	file_codenode_proto_rawDescOnce.Do(func() {
		file_codenode_proto_rawDescData = protoimpl.X.CompressGZIP(file_codenode_proto_rawDescData)
	})
	return file_codenode_proto_rawDescData
}

var file_codenode_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_codenode_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_codenode_proto_goTypes = []interface{}{
	(CodeType)(0),            // 0: codenode.CodeType
	(OperateType)(0),         // 1: codenode.OperateType
	(*CodeReq)(nil),          // 2: codenode.CodeReq
	(*CodeRes)(nil),          // 3: codenode.CodeRes
	(*VerifyCodeReq)(nil),    // 4: codenode.VerifyCodeReq
	(*VerifyCodeRes)(nil),    // 5: codenode.VerifyCodeRes
	nil,                      // 6: codenode.CodeReq.ExtraEntry
	nil,                      // 7: codenode.CodeRes.ExtraEntry
	nil,                      // 8: codenode.VerifyCodeReq.ExtraEntry
	nil,                      // 9: codenode.VerifyCodeRes.ExtraEntry
	(*common.Authorize)(nil), // 10: common.Authorize
	(rescode.Code)(0),        // 11: rescode.Code
}
var file_codenode_proto_depIdxs = []int32{
	10, // 0: codenode.CodeReq.auth:type_name -> common.Authorize
	0,  // 1: codenode.CodeReq.type:type_name -> codenode.CodeType
	1,  // 2: codenode.CodeReq.optype:type_name -> codenode.OperateType
	6,  // 3: codenode.CodeReq.extra:type_name -> codenode.CodeReq.ExtraEntry
	11, // 4: codenode.CodeRes.rescode:type_name -> rescode.Code
	7,  // 5: codenode.CodeRes.extra:type_name -> codenode.CodeRes.ExtraEntry
	10, // 6: codenode.VerifyCodeReq.auth:type_name -> common.Authorize
	1,  // 7: codenode.VerifyCodeReq.optype:type_name -> codenode.OperateType
	8,  // 8: codenode.VerifyCodeReq.extra:type_name -> codenode.VerifyCodeReq.ExtraEntry
	11, // 9: codenode.VerifyCodeRes.rescode:type_name -> rescode.Code
	9,  // 10: codenode.VerifyCodeRes.extra:type_name -> codenode.VerifyCodeRes.ExtraEntry
	2,  // 11: codenode.CodenodeService.SendCode:input_type -> codenode.CodeReq
	4,  // 12: codenode.CodenodeService.VerifyCode:input_type -> codenode.VerifyCodeReq
	3,  // 13: codenode.CodenodeService.SendCode:output_type -> codenode.CodeRes
	5,  // 14: codenode.CodenodeService.VerifyCode:output_type -> codenode.VerifyCodeRes
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_codenode_proto_init() }
func file_codenode_proto_init() {
	if File_codenode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codenode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeReq); i {
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
		file_codenode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRes); i {
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
		file_codenode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeReq); i {
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
		file_codenode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeRes); i {
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
			RawDescriptor: file_codenode_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codenode_proto_goTypes,
		DependencyIndexes: file_codenode_proto_depIdxs,
		EnumInfos:         file_codenode_proto_enumTypes,
		MessageInfos:      file_codenode_proto_msgTypes,
	}.Build()
	File_codenode_proto = out.File
	file_codenode_proto_rawDesc = nil
	file_codenode_proto_goTypes = nil
	file_codenode_proto_depIdxs = nil
}
