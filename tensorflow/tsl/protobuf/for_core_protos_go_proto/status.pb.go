// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: tsl/protobuf/status.proto

package for_core_protos_go_proto

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

// Wire-format for Status.
// Next tag: 3
type StatusProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status code as defined in tensorflow/tsl/protobuf/error_codes.proto.
	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=tensorflow.error.Code" json:"code,omitempty"`
	// Detail error message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusProto) Reset() {
	*x = StatusProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tsl_protobuf_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusProto) ProtoMessage() {}

func (x *StatusProto) ProtoReflect() protoreflect.Message {
	mi := &file_tsl_protobuf_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusProto.ProtoReflect.Descriptor instead.
func (*StatusProto) Descriptor() ([]byte, []int) {
	return file_tsl_protobuf_status_proto_rawDescGZIP(), []int{0}
}

func (x *StatusProto) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

func (x *StatusProto) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_tsl_protobuf_status_proto protoreflect.FileDescriptor

var file_tsl_protobuf_status_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x1e, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x0a, 0x18,
	0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tsl_protobuf_status_proto_rawDescOnce sync.Once
	file_tsl_protobuf_status_proto_rawDescData = file_tsl_protobuf_status_proto_rawDesc
)

func file_tsl_protobuf_status_proto_rawDescGZIP() []byte {
	file_tsl_protobuf_status_proto_rawDescOnce.Do(func() {
		file_tsl_protobuf_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_tsl_protobuf_status_proto_rawDescData)
	})
	return file_tsl_protobuf_status_proto_rawDescData
}

var file_tsl_protobuf_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tsl_protobuf_status_proto_goTypes = []interface{}{
	(*StatusProto)(nil), // 0: tensorflow.StatusProto
	(Code)(0),           // 1: tensorflow.error.Code
}
var file_tsl_protobuf_status_proto_depIdxs = []int32{
	1, // 0: tensorflow.StatusProto.code:type_name -> tensorflow.error.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tsl_protobuf_status_proto_init() }
func file_tsl_protobuf_status_proto_init() {
	if File_tsl_protobuf_status_proto != nil {
		return
	}
	file_tsl_protobuf_error_codes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tsl_protobuf_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusProto); i {
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
			RawDescriptor: file_tsl_protobuf_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tsl_protobuf_status_proto_goTypes,
		DependencyIndexes: file_tsl_protobuf_status_proto_depIdxs,
		MessageInfos:      file_tsl_protobuf_status_proto_msgTypes,
	}.Build()
	File_tsl_protobuf_status_proto = out.File
	file_tsl_protobuf_status_proto_rawDesc = nil
	file_tsl_protobuf_status_proto_goTypes = nil
	file_tsl_protobuf_status_proto_depIdxs = nil
}
