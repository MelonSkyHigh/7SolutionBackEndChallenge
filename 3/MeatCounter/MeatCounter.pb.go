// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0--rc3
// source: MeatCounter/MeatCounter.proto

package MeatCounter

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MeatCounter_MeatCounter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_MeatCounter_MeatCounter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_MeatCounter_MeatCounter_proto_rawDescGZIP(), []int{0}
}

type MeatRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeatCount string `protobuf:"bytes,2,opt,name=MeatCount,proto3" json:"MeatCount,omitempty"`
}

func (x *MeatRes) Reset() {
	*x = MeatRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MeatCounter_MeatCounter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeatRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeatRes) ProtoMessage() {}

func (x *MeatRes) ProtoReflect() protoreflect.Message {
	mi := &file_MeatCounter_MeatCounter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeatRes.ProtoReflect.Descriptor instead.
func (*MeatRes) Descriptor() ([]byte, []int) {
	return file_MeatCounter_MeatCounter_proto_rawDescGZIP(), []int{1}
}

func (x *MeatRes) GetMeatCount() string {
	if x != nil {
		return x.MeatCount
	}
	return ""
}

var File_MeatCounter_MeatCounter_proto protoreflect.FileDescriptor

var file_MeatCounter_MeatCounter_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x4d, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x4d, 0x65,
	0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x4d, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x07, 0x4d, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x46,
	0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a,
	0x09, 0x4d, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x4d, 0x65, 0x61,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14,
	0x2e, 0x4d, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x4d, 0x65, 0x61, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MeatCounter_MeatCounter_proto_rawDescOnce sync.Once
	file_MeatCounter_MeatCounter_proto_rawDescData = file_MeatCounter_MeatCounter_proto_rawDesc
)

func file_MeatCounter_MeatCounter_proto_rawDescGZIP() []byte {
	file_MeatCounter_MeatCounter_proto_rawDescOnce.Do(func() {
		file_MeatCounter_MeatCounter_proto_rawDescData = protoimpl.X.CompressGZIP(file_MeatCounter_MeatCounter_proto_rawDescData)
	})
	return file_MeatCounter_MeatCounter_proto_rawDescData
}

var file_MeatCounter_MeatCounter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_MeatCounter_MeatCounter_proto_goTypes = []interface{}{
	(*Empty)(nil),   // 0: MeatCounter.Empty
	(*MeatRes)(nil), // 1: MeatCounter.MeatRes
}
var file_MeatCounter_MeatCounter_proto_depIdxs = []int32{
	0, // 0: MeatCounter.MeatCounter.MeatCount:input_type -> MeatCounter.Empty
	1, // 1: MeatCounter.MeatCounter.MeatCount:output_type -> MeatCounter.MeatRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MeatCounter_MeatCounter_proto_init() }
func file_MeatCounter_MeatCounter_proto_init() {
	if File_MeatCounter_MeatCounter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MeatCounter_MeatCounter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_MeatCounter_MeatCounter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeatRes); i {
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
			RawDescriptor: file_MeatCounter_MeatCounter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MeatCounter_MeatCounter_proto_goTypes,
		DependencyIndexes: file_MeatCounter_MeatCounter_proto_depIdxs,
		MessageInfos:      file_MeatCounter_MeatCounter_proto_msgTypes,
	}.Build()
	File_MeatCounter_MeatCounter_proto = out.File
	file_MeatCounter_MeatCounter_proto_rawDesc = nil
	file_MeatCounter_MeatCounter_proto_goTypes = nil
	file_MeatCounter_MeatCounter_proto_depIdxs = nil
}
