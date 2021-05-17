// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: github.com/Ayna5/bannersRotation/api/banners-rotation/banners_rotation.proto

package banners_rotation_pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AddBannerToSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId uint64 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   uint64 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *AddBannerToSlotRequest) Reset() {
	*x = AddBannerToSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerToSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerToSlotRequest) ProtoMessage() {}

func (x *AddBannerToSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerToSlotRequest.ProtoReflect.Descriptor instead.
func (*AddBannerToSlotRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{0}
}

func (x *AddBannerToSlotRequest) GetBannerId() uint64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *AddBannerToSlotRequest) GetSlotId() uint64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type AddBannerToSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddBannerToSlotResponse) Reset() {
	*x = AddBannerToSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerToSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerToSlotResponse) ProtoMessage() {}

func (x *AddBannerToSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerToSlotResponse.ProtoReflect.Descriptor instead.
func (*AddBannerToSlotResponse) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{1}
}

type RemoveBannerFromSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId uint64 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId   uint64 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
}

func (x *RemoveBannerFromSlotRequest) Reset() {
	*x = RemoveBannerFromSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBannerFromSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBannerFromSlotRequest) ProtoMessage() {}

func (x *RemoveBannerFromSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBannerFromSlotRequest.ProtoReflect.Descriptor instead.
func (*RemoveBannerFromSlotRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveBannerFromSlotRequest) GetBannerId() uint64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *RemoveBannerFromSlotRequest) GetSlotId() uint64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

type RemoveBannerFromSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveBannerFromSlotResponse) Reset() {
	*x = RemoveBannerFromSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveBannerFromSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveBannerFromSlotResponse) ProtoMessage() {}

func (x *RemoveBannerFromSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveBannerFromSlotResponse.ProtoReflect.Descriptor instead.
func (*RemoveBannerFromSlotResponse) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{3}
}

type AddClickForBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId    uint64 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	SlotId      uint64 `protobuf:"varint,2,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	UserGroupId uint64 `protobuf:"varint,3,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
}

func (x *AddClickForBannerRequest) Reset() {
	*x = AddClickForBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClickForBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClickForBannerRequest) ProtoMessage() {}

func (x *AddClickForBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClickForBannerRequest.ProtoReflect.Descriptor instead.
func (*AddClickForBannerRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{4}
}

func (x *AddClickForBannerRequest) GetBannerId() uint64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

func (x *AddClickForBannerRequest) GetSlotId() uint64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *AddClickForBannerRequest) GetUserGroupId() uint64 {
	if x != nil {
		return x.UserGroupId
	}
	return 0
}

type AddClickForBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddClickForBannerResponse) Reset() {
	*x = AddClickForBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClickForBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClickForBannerResponse) ProtoMessage() {}

func (x *AddClickForBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClickForBannerResponse.ProtoReflect.Descriptor instead.
func (*AddClickForBannerResponse) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{5}
}

type GetBannerForSlotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId      uint64 `protobuf:"varint,1,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	UserGroupId uint64 `protobuf:"varint,2,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
}

func (x *GetBannerForSlotRequest) Reset() {
	*x = GetBannerForSlotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBannerForSlotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBannerForSlotRequest) ProtoMessage() {}

func (x *GetBannerForSlotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBannerForSlotRequest.ProtoReflect.Descriptor instead.
func (*GetBannerForSlotRequest) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{6}
}

func (x *GetBannerForSlotRequest) GetSlotId() uint64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *GetBannerForSlotRequest) GetUserGroupId() uint64 {
	if x != nil {
		return x.UserGroupId
	}
	return 0
}

type GetBannerForSlotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId uint64 `protobuf:"varint,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
}

func (x *GetBannerForSlotResponse) Reset() {
	*x = GetBannerForSlotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBannerForSlotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBannerForSlotResponse) ProtoMessage() {}

func (x *GetBannerForSlotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBannerForSlotResponse.ProtoReflect.Descriptor instead.
func (*GetBannerForSlotResponse) Descriptor() ([]byte, []int) {
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP(), []int{7}
}

func (x *GetBannerForSlotResponse) GetBannerId() uint64 {
	if x != nil {
		return x.BannerId
	}
	return 0
}

var File_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto protoreflect.FileDescriptor

var file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x79, 0x6e,
	0x61, 0x35, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x5f,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e,
	0x0a, 0x16, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x6f, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x19,
	0x0a, 0x17, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x6f, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x1b, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x1e,
	0x0a, 0x1c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x74,
	0x0a, 0x18, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b,
	0x46, 0x6f, 0x72, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x56, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f,
	0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x32, 0xd1, 0x03, 0x0a, 0x0f, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x68, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x54, 0x6f, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x28, 0x2e, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x6f, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x54, 0x6f, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x77, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x2d, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6c, 0x6f, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x11, 0x41, 0x64, 0x64,
	0x43, 0x6c, 0x69, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2a,
	0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64,
	0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x46, 0x6f, 0x72, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x29, 0x2e,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x73, 0x2e, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescOnce sync.Once
	file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescData = file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDesc
)

func file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescGZIP() []byte {
	file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescOnce.Do(func() {
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescData)
	})
	return file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDescData
}

var file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_goTypes = []interface{}{
	(*AddBannerToSlotRequest)(nil),       // 0: banners.rotation.AddBannerToSlotRequest
	(*AddBannerToSlotResponse)(nil),      // 1: banners.rotation.AddBannerToSlotResponse
	(*RemoveBannerFromSlotRequest)(nil),  // 2: banners.rotation.RemoveBannerFromSlotRequest
	(*RemoveBannerFromSlotResponse)(nil), // 3: banners.rotation.RemoveBannerFromSlotResponse
	(*AddClickForBannerRequest)(nil),     // 4: banners.rotation.AddClickForBannerRequest
	(*AddClickForBannerResponse)(nil),    // 5: banners.rotation.AddClickForBannerResponse
	(*GetBannerForSlotRequest)(nil),      // 6: banners.rotation.GetBannerForSlotRequest
	(*GetBannerForSlotResponse)(nil),     // 7: banners.rotation.GetBannerForSlotResponse
}
var file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_depIdxs = []int32{
	0, // 0: banners.rotation.BannersRotation.AddBannerToSlot:input_type -> banners.rotation.AddBannerToSlotRequest
	2, // 1: banners.rotation.BannersRotation.RemoveBannerFromSlot:input_type -> banners.rotation.RemoveBannerFromSlotRequest
	4, // 2: banners.rotation.BannersRotation.AddClickForBanner:input_type -> banners.rotation.AddClickForBannerRequest
	6, // 3: banners.rotation.BannersRotation.GetBannerForSlot:input_type -> banners.rotation.GetBannerForSlotRequest
	1, // 4: banners.rotation.BannersRotation.AddBannerToSlot:output_type -> banners.rotation.AddBannerToSlotResponse
	3, // 5: banners.rotation.BannersRotation.RemoveBannerFromSlot:output_type -> banners.rotation.RemoveBannerFromSlotResponse
	5, // 6: banners.rotation.BannersRotation.AddClickForBanner:output_type -> banners.rotation.AddClickForBannerResponse
	7, // 7: banners.rotation.BannersRotation.GetBannerForSlot:output_type -> banners.rotation.GetBannerForSlotResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_init() }
func file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_init() {
	if File_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerToSlotRequest); i {
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
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerToSlotResponse); i {
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
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBannerFromSlotRequest); i {
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
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveBannerFromSlotResponse); i {
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
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClickForBannerRequest); i {
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
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClickForBannerResponse); i {
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
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBannerForSlotRequest); i {
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
		file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBannerForSlotResponse); i {
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
			RawDescriptor: file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_goTypes,
		DependencyIndexes: file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_depIdxs,
		MessageInfos:      file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_msgTypes,
	}.Build()
	File_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto = out.File
	file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_rawDesc = nil
	file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_goTypes = nil
	file_github_com_Ayna5_bannersRotation_api_banners_rotation_banners_rotation_proto_depIdxs = nil
}