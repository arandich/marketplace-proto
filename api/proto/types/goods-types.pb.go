// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proto/types/goods-types.proto

package types

import (
	enums "github.com/arandich/marketplace-proto/api/proto/enums"
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

type Good struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string           `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64            `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Status      enums.GoodStatus `protobuf:"varint,6,opt,name=status,proto3,enum=marketplace.GoodStatus" json:"status,omitempty"`
	ImageUrl    string           `protobuf:"bytes,7,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *Good) Reset() {
	*x = Good{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_goods_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Good) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Good) ProtoMessage() {}

func (x *Good) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_goods_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Good.ProtoReflect.Descriptor instead.
func (*Good) Descriptor() ([]byte, []int) {
	return file_proto_types_goods_types_proto_rawDescGZIP(), []int{0}
}

func (x *Good) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Good) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Good) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Good) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Good) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Good) GetStatus() enums.GoodStatus {
	if x != nil {
		return x.Status
	}
	return enums.GoodStatus(0)
}

func (x *Good) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_proto_types_goods_types_proto protoreflect.FileDescriptor

var file_proto_types_goods_types_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x1a, 0x1d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2d,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x04,
	0x47, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x68, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_goods_types_proto_rawDescOnce sync.Once
	file_proto_types_goods_types_proto_rawDescData = file_proto_types_goods_types_proto_rawDesc
)

func file_proto_types_goods_types_proto_rawDescGZIP() []byte {
	file_proto_types_goods_types_proto_rawDescOnce.Do(func() {
		file_proto_types_goods_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_goods_types_proto_rawDescData)
	})
	return file_proto_types_goods_types_proto_rawDescData
}

var file_proto_types_goods_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_types_goods_types_proto_goTypes = []interface{}{
	(*Good)(nil),          // 0: marketplace.Good
	(enums.GoodStatus)(0), // 1: marketplace.GoodStatus
}
var file_proto_types_goods_types_proto_depIdxs = []int32{
	1, // 0: marketplace.Good.status:type_name -> marketplace.GoodStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_types_goods_types_proto_init() }
func file_proto_types_goods_types_proto_init() {
	if File_proto_types_goods_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_goods_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Good); i {
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
			RawDescriptor: file_proto_types_goods_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_goods_types_proto_goTypes,
		DependencyIndexes: file_proto_types_goods_types_proto_depIdxs,
		MessageInfos:      file_proto_types_goods_types_proto_msgTypes,
	}.Build()
	File_proto_types_goods_types_proto = out.File
	file_proto_types_goods_types_proto_rawDesc = nil
	file_proto_types_goods_types_proto_goTypes = nil
	file_proto_types_goods_types_proto_depIdxs = nil
}
