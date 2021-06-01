// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: pkg/rpc/gitops/gitops.proto

package gitops

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

type DeploymentType int32

const (
	DeploymentType_kustomize DeploymentType = 0
	DeploymentType_helm      DeploymentType = 1
)

// Enum value maps for DeploymentType.
var (
	DeploymentType_name = map[int32]string{
		0: "kustomize",
		1: "helm",
	}
	DeploymentType_value = map[string]int32{
		"kustomize": 0,
		"helm":      1,
	}
)

func (x DeploymentType) Enum() *DeploymentType {
	p := new(DeploymentType)
	*p = x
	return p
}

func (x DeploymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_rpc_gitops_gitops_proto_enumTypes[0].Descriptor()
}

func (DeploymentType) Type() protoreflect.EnumType {
	return &file_pkg_rpc_gitops_gitops_proto_enumTypes[0]
}

func (x DeploymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentType.Descriptor instead.
func (DeploymentType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_rpc_gitops_gitops_proto_rawDescGZIP(), []int{0}
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_gitops_gitops_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_gitops_gitops_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRes) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

//
// Represents a Weave GitOps application
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_gitops_gitops_proto_rawDescGZIP(), []int{2}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//
// ListApplication request payload
type AddApplicationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner          string         `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name           string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url            string         `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Path           string         `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Branch         string         `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	DeploymentType DeploymentType `protobuf:"varint,6,opt,name=deployment_type,json=deploymentType,proto3,enum=gitops.DeploymentType" json:"deployment_type,omitempty"`
	PrivateKey     string         `protobuf:"bytes,7,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	DryRun         bool           `protobuf:"varint,8,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	Private        bool           `protobuf:"varint,9,opt,name=private,proto3" json:"private,omitempty"`
	Namespace      string         `protobuf:"bytes,10,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Dir            string         `protobuf:"bytes,11,opt,name=dir,proto3" json:"dir,omitempty"`
}

func (x *AddApplicationReq) Reset() {
	*x = AddApplicationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddApplicationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddApplicationReq) ProtoMessage() {}

func (x *AddApplicationReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddApplicationReq.ProtoReflect.Descriptor instead.
func (*AddApplicationReq) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_gitops_gitops_proto_rawDescGZIP(), []int{3}
}

func (x *AddApplicationReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AddApplicationReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddApplicationReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddApplicationReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AddApplicationReq) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *AddApplicationReq) GetDeploymentType() DeploymentType {
	if x != nil {
		return x.DeploymentType
	}
	return DeploymentType_kustomize
}

func (x *AddApplicationReq) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *AddApplicationReq) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *AddApplicationReq) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *AddApplicationReq) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AddApplicationReq) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

//
// ListApplication response payload
type AddApplicationRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *AddApplicationRes) Reset() {
	*x = AddApplicationRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddApplicationRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddApplicationRes) ProtoMessage() {}

func (x *AddApplicationRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_gitops_gitops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddApplicationRes.ProtoReflect.Descriptor instead.
func (*AddApplicationRes) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_gitops_gitops_proto_rawDescGZIP(), []int{4}
}

func (x *AddApplicationRes) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_pkg_rpc_gitops_gitops_proto protoreflect.FileDescriptor

var file_pkg_rpc_gitops_gitops_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73,
	0x2f, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67,
	0x69, 0x74, 0x6f, 0x70, 0x73, 0x22, 0x20, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2d, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x21, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x3f, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79,
	0x5f, 0x72, 0x75, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52,
	0x75, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x22, 0x4a, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x29, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x6b, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x68, 0x65, 0x6c,
	0x6d, 0x10, 0x01, 0x32, 0x7d, 0x0a, 0x06, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x12, 0x2b, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70,
	0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x67,
	0x69, 0x74, 0x6f, 0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x42, 0x1e, 0x42, 0x0c, 0x47, 0x69, 0x74, 0x4f, 0x70, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x5a, 0x0e, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x69, 0x74, 0x6f,
	0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_gitops_gitops_proto_rawDescOnce sync.Once
	file_pkg_rpc_gitops_gitops_proto_rawDescData = file_pkg_rpc_gitops_gitops_proto_rawDesc
)

func file_pkg_rpc_gitops_gitops_proto_rawDescGZIP() []byte {
	file_pkg_rpc_gitops_gitops_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_gitops_gitops_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_gitops_gitops_proto_rawDescData)
	})
	return file_pkg_rpc_gitops_gitops_proto_rawDescData
}

var file_pkg_rpc_gitops_gitops_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_rpc_gitops_gitops_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_rpc_gitops_gitops_proto_goTypes = []interface{}{
	(DeploymentType)(0),       // 0: gitops.DeploymentType
	(*LoginReq)(nil),          // 1: gitops.LoginReq
	(*LoginRes)(nil),          // 2: gitops.LoginRes
	(*Application)(nil),       // 3: gitops.Application
	(*AddApplicationReq)(nil), // 4: gitops.AddApplicationReq
	(*AddApplicationRes)(nil), // 5: gitops.AddApplicationRes
}
var file_pkg_rpc_gitops_gitops_proto_depIdxs = []int32{
	0, // 0: gitops.AddApplicationReq.deployment_type:type_name -> gitops.DeploymentType
	3, // 1: gitops.AddApplicationRes.application:type_name -> gitops.Application
	1, // 2: gitops.GitOps.Login:input_type -> gitops.LoginReq
	4, // 3: gitops.GitOps.AddApplication:input_type -> gitops.AddApplicationReq
	2, // 4: gitops.GitOps.Login:output_type -> gitops.LoginRes
	5, // 5: gitops.GitOps.AddApplication:output_type -> gitops.AddApplicationRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_rpc_gitops_gitops_proto_init() }
func file_pkg_rpc_gitops_gitops_proto_init() {
	if File_pkg_rpc_gitops_gitops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_gitops_gitops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_pkg_rpc_gitops_gitops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
		file_pkg_rpc_gitops_gitops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_pkg_rpc_gitops_gitops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddApplicationReq); i {
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
		file_pkg_rpc_gitops_gitops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddApplicationRes); i {
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
			RawDescriptor: file_pkg_rpc_gitops_gitops_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rpc_gitops_gitops_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_gitops_gitops_proto_depIdxs,
		EnumInfos:         file_pkg_rpc_gitops_gitops_proto_enumTypes,
		MessageInfos:      file_pkg_rpc_gitops_gitops_proto_msgTypes,
	}.Build()
	File_pkg_rpc_gitops_gitops_proto = out.File
	file_pkg_rpc_gitops_gitops_proto_rawDesc = nil
	file_pkg_rpc_gitops_gitops_proto_goTypes = nil
	file_pkg_rpc_gitops_gitops_proto_depIdxs = nil
}