// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proto/enums/goods-enums.proto

package enums

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

type GoodStatus int32

const (
	GoodStatus_ACTIVE   GoodStatus = 0 // активен
	GoodStatus_INACTIVE GoodStatus = 1 // снят с продажи
)

// Enum value maps for GoodStatus.
var (
	GoodStatus_name = map[int32]string{
		0: "ACTIVE",
		1: "INACTIVE",
	}
	GoodStatus_value = map[string]int32{
		"ACTIVE":   0,
		"INACTIVE": 1,
	}
)

func (x GoodStatus) Enum() *GoodStatus {
	p := new(GoodStatus)
	*p = x
	return p
}

func (x GoodStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoodStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enums_goods_enums_proto_enumTypes[0].Descriptor()
}

func (GoodStatus) Type() protoreflect.EnumType {
	return &file_proto_enums_goods_enums_proto_enumTypes[0]
}

func (x GoodStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoodStatus.Descriptor instead.
func (GoodStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_enums_goods_enums_proto_rawDescGZIP(), []int{0}
}

var File_proto_enums_goods_enums_proto protoreflect.FileDescriptor

var file_proto_enums_goods_enums_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2d, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2a, 0x26, 0x0a, 0x0a,
	0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x64, 0x69, 0x63, 0x68, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_enums_goods_enums_proto_rawDescOnce sync.Once
	file_proto_enums_goods_enums_proto_rawDescData = file_proto_enums_goods_enums_proto_rawDesc
)

func file_proto_enums_goods_enums_proto_rawDescGZIP() []byte {
	file_proto_enums_goods_enums_proto_rawDescOnce.Do(func() {
		file_proto_enums_goods_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_enums_goods_enums_proto_rawDescData)
	})
	return file_proto_enums_goods_enums_proto_rawDescData
}

var file_proto_enums_goods_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_enums_goods_enums_proto_goTypes = []interface{}{
	(GoodStatus)(0), // 0: marketplace.GoodStatus
}
var file_proto_enums_goods_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_enums_goods_enums_proto_init() }
func file_proto_enums_goods_enums_proto_init() {
	if File_proto_enums_goods_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_enums_goods_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_enums_goods_enums_proto_goTypes,
		DependencyIndexes: file_proto_enums_goods_enums_proto_depIdxs,
		EnumInfos:         file_proto_enums_goods_enums_proto_enumTypes,
	}.Build()
	File_proto_enums_goods_enums_proto = out.File
	file_proto_enums_goods_enums_proto_rawDesc = nil
	file_proto_enums_goods_enums_proto_goTypes = nil
	file_proto_enums_goods_enums_proto_depIdxs = nil
}