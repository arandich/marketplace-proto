// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proto/services/id-service.proto

package services

import (
	enums "github.com/arandich/marketplace-proto/api/proto/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_id_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_id_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_id_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken  string `protobuf:"bytes,1,opt,name=jwt_token,json=jwtToken,proto3" json:"jwt_token,omitempty"`
	ExpiresAt int64  `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_id_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_id_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_id_service_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

func (x *AuthResponse) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type InitHoldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *InitHoldRequest) Reset() {
	*x = InitHoldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_id_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitHoldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitHoldRequest) ProtoMessage() {}

func (x *InitHoldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_id_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitHoldRequest.ProtoReflect.Descriptor instead.
func (*InitHoldRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_id_service_proto_rawDescGZIP(), []int{2}
}

func (x *InitHoldRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type InitHoldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HoldId int64            `protobuf:"varint,1,opt,name=hold_id,json=holdId,proto3" json:"hold_id,omitempty"`
	Status enums.HoldStatus `protobuf:"varint,2,opt,name=status,proto3,enum=marketplace.HoldStatus" json:"status,omitempty"`
}

func (x *InitHoldResponse) Reset() {
	*x = InitHoldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_id_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitHoldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitHoldResponse) ProtoMessage() {}

func (x *InitHoldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_id_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitHoldResponse.ProtoReflect.Descriptor instead.
func (*InitHoldResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_id_service_proto_rawDescGZIP(), []int{3}
}

func (x *InitHoldResponse) GetHoldId() int64 {
	if x != nil {
		return x.HoldId
	}
	return 0
}

func (x *InitHoldResponse) GetStatus() enums.HoldStatus {
	if x != nil {
		return x.Status
	}
	return enums.HoldStatus(0)
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_id_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_id_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_services_id_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_proto_services_id_service_proto protoreflect.FileDescriptor

var file_proto_services_id_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x69, 0x64, 0x2d, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4a, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6a, 0x77, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74,
	0x48, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x48,
	0x6f, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xab, 0x02,
	0x0a, 0x09, 0x49, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10,
	0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x12, 0x6f, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x1c, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x48,
	0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x48, 0x6f, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x2f, 0x69, 0x73, 0x2d,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x58, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x64, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x64, 0x69,
	0x63, 0x68, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_id_service_proto_rawDescOnce sync.Once
	file_proto_services_id_service_proto_rawDescData = file_proto_services_id_service_proto_rawDesc
)

func file_proto_services_id_service_proto_rawDescGZIP() []byte {
	file_proto_services_id_service_proto_rawDescOnce.Do(func() {
		file_proto_services_id_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_id_service_proto_rawDescData)
	})
	return file_proto_services_id_service_proto_rawDescData
}

var file_proto_services_id_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_services_id_service_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),      // 0: marketplace.AuthRequest
	(*AuthResponse)(nil),     // 1: marketplace.AuthResponse
	(*InitHoldRequest)(nil),  // 2: marketplace.InitHoldRequest
	(*InitHoldResponse)(nil), // 3: marketplace.InitHoldResponse
	(*GetUserResponse)(nil),  // 4: marketplace.GetUserResponse
	(enums.HoldStatus)(0),    // 5: marketplace.HoldStatus
	(*emptypb.Empty)(nil),    // 6: google.protobuf.Empty
}
var file_proto_services_id_service_proto_depIdxs = []int32{
	5, // 0: marketplace.InitHoldResponse.status:type_name -> marketplace.HoldStatus
	0, // 1: marketplace.IdService.Auth:input_type -> marketplace.AuthRequest
	2, // 2: marketplace.IdService.InitHold:input_type -> marketplace.InitHoldRequest
	6, // 3: marketplace.IdService.GetUser:input_type -> google.protobuf.Empty
	1, // 4: marketplace.IdService.Auth:output_type -> marketplace.AuthResponse
	3, // 5: marketplace.IdService.InitHold:output_type -> marketplace.InitHoldResponse
	4, // 6: marketplace.IdService.GetUser:output_type -> marketplace.GetUserResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_services_id_service_proto_init() }
func file_proto_services_id_service_proto_init() {
	if File_proto_services_id_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_id_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_proto_services_id_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_proto_services_id_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitHoldRequest); i {
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
		file_proto_services_id_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitHoldResponse); i {
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
		file_proto_services_id_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
			RawDescriptor: file_proto_services_id_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_id_service_proto_goTypes,
		DependencyIndexes: file_proto_services_id_service_proto_depIdxs,
		MessageInfos:      file_proto_services_id_service_proto_msgTypes,
	}.Build()
	File_proto_services_id_service_proto = out.File
	file_proto_services_id_service_proto_rawDesc = nil
	file_proto_services_id_service_proto_goTypes = nil
	file_proto_services_id_service_proto_depIdxs = nil
}
