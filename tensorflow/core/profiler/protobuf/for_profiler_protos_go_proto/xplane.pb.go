// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.7
// source: tensorflow/core/profiler/protobuf/xplane.proto

package for_profiler_protos_go_proto

import (
	for_profiler_protos_go_proto "github.com/wamuir/graft/tensorflow/tsl/profiler/protobuf/for_profiler_protos_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorflow/tsl/profiler/protobuf/xplane.proto.

type XSpace = for_profiler_protos_go_proto.XSpace
type XPlane = for_profiler_protos_go_proto.XPlane
type XLine = for_profiler_protos_go_proto.XLine
type XEvent = for_profiler_protos_go_proto.XEvent
type XEvent_OffsetPs = for_profiler_protos_go_proto.XEvent_OffsetPs
type XEvent_NumOccurrences = for_profiler_protos_go_proto.XEvent_NumOccurrences
type XStat = for_profiler_protos_go_proto.XStat
type XStat_DoubleValue = for_profiler_protos_go_proto.XStat_DoubleValue
type XStat_Uint64Value = for_profiler_protos_go_proto.XStat_Uint64Value
type XStat_Int64Value = for_profiler_protos_go_proto.XStat_Int64Value
type XStat_StrValue = for_profiler_protos_go_proto.XStat_StrValue
type XStat_BytesValue = for_profiler_protos_go_proto.XStat_BytesValue
type XStat_RefValue = for_profiler_protos_go_proto.XStat_RefValue
type XEventMetadata = for_profiler_protos_go_proto.XEventMetadata
type XStatMetadata = for_profiler_protos_go_proto.XStatMetadata

var File_tensorflow_core_profiler_protobuf_xplane_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_protobuf_xplane_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2d, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x78, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x64, 0x5a, 0x62, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tensorflow_core_profiler_protobuf_xplane_proto_goTypes = []interface{}{}
var file_tensorflow_core_profiler_protobuf_xplane_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_protobuf_xplane_proto_init() }
func file_tensorflow_core_profiler_protobuf_xplane_proto_init() {
	if File_tensorflow_core_profiler_protobuf_xplane_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_profiler_protobuf_xplane_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_protobuf_xplane_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_protobuf_xplane_proto_depIdxs,
	}.Build()
	File_tensorflow_core_profiler_protobuf_xplane_proto = out.File
	file_tensorflow_core_profiler_protobuf_xplane_proto_rawDesc = nil
	file_tensorflow_core_profiler_protobuf_xplane_proto_goTypes = nil
	file_tensorflow_core_profiler_protobuf_xplane_proto_depIdxs = nil
}
