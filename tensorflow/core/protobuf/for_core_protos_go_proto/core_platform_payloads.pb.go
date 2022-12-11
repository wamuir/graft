// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: tensorflow/core/protobuf/core_platform_payloads.proto

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

type ErrorSourceProto_ErrorSource int32

const (
	ErrorSourceProto_UNKNOWN        ErrorSourceProto_ErrorSource = 0
	ErrorSourceProto_TPU_COMPILE_OP ErrorSourceProto_ErrorSource = 1
	// Old bridge.
	ErrorSourceProto_TF_XLA_BRIDGE ErrorSourceProto_ErrorSource = 2
	// TPUBridge.
	ErrorSourceProto_MLIR_BRIDGE_PHASE_1 ErrorSourceProto_ErrorSource = 3
	// LegalizeToHlo.
	ErrorSourceProto_MLIR_BRIDGE_PHASE_2 ErrorSourceProto_ErrorSource = 4
	// eager::RemoteMgr.
	ErrorSourceProto_EAGER_REMOTE_MGR ErrorSourceProto_ErrorSource = 5
)

// Enum value maps for ErrorSourceProto_ErrorSource.
var (
	ErrorSourceProto_ErrorSource_name = map[int32]string{
		0: "UNKNOWN",
		1: "TPU_COMPILE_OP",
		2: "TF_XLA_BRIDGE",
		3: "MLIR_BRIDGE_PHASE_1",
		4: "MLIR_BRIDGE_PHASE_2",
		5: "EAGER_REMOTE_MGR",
	}
	ErrorSourceProto_ErrorSource_value = map[string]int32{
		"UNKNOWN":             0,
		"TPU_COMPILE_OP":      1,
		"TF_XLA_BRIDGE":       2,
		"MLIR_BRIDGE_PHASE_1": 3,
		"MLIR_BRIDGE_PHASE_2": 4,
		"EAGER_REMOTE_MGR":    5,
	}
)

func (x ErrorSourceProto_ErrorSource) Enum() *ErrorSourceProto_ErrorSource {
	p := new(ErrorSourceProto_ErrorSource)
	*p = x
	return p
}

func (x ErrorSourceProto_ErrorSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorSourceProto_ErrorSource) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_core_platform_payloads_proto_enumTypes[0].Descriptor()
}

func (ErrorSourceProto_ErrorSource) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_core_platform_payloads_proto_enumTypes[0]
}

func (x ErrorSourceProto_ErrorSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorSourceProto_ErrorSource.Descriptor instead.
func (ErrorSourceProto_ErrorSource) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescGZIP(), []int{0, 0}
}

// If included as a payload, this message contains the error source information
// where the error was raised.
// URI: "type.googleapis.com/tensorflow.core.platform.ErrorSourceProto"
type ErrorSourceProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorSource ErrorSourceProto_ErrorSource `protobuf:"varint,1,opt,name=error_source,json=errorSource,proto3,enum=tensorflow.core.platform.ErrorSourceProto_ErrorSource" json:"error_source,omitempty"`
}

func (x *ErrorSourceProto) Reset() {
	*x = ErrorSourceProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_core_platform_payloads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorSourceProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorSourceProto) ProtoMessage() {}

func (x *ErrorSourceProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_core_platform_payloads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorSourceProto.ProtoReflect.Descriptor instead.
func (*ErrorSourceProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorSourceProto) GetErrorSource() ErrorSourceProto_ErrorSource {
	if x != nil {
		return x.ErrorSource
	}
	return ErrorSourceProto_UNKNOWN
}

var File_tensorflow_core_protobuf_core_platform_payloads_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x22, 0xf9, 0x01, 0x0a, 0x10, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x59, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x54, 0x50, 0x55, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x49, 0x4c, 0x45, 0x5f, 0x4f, 0x50,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x46, 0x5f, 0x58, 0x4c, 0x41, 0x5f, 0x42, 0x52, 0x49,
	0x44, 0x47, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x4c, 0x49, 0x52, 0x5f, 0x42, 0x52,
	0x49, 0x44, 0x47, 0x45, 0x5f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x31, 0x10, 0x03, 0x12, 0x17,
	0x0a, 0x13, 0x4d, 0x4c, 0x49, 0x52, 0x5f, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45, 0x5f, 0x50, 0x48,
	0x41, 0x53, 0x45, 0x5f, 0x32, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x41, 0x47, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x4d, 0x47, 0x52, 0x10, 0x05, 0x42, 0x5a, 0x5a,
	0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescData = file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDesc
)

func file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDescData
}

var file_tensorflow_core_protobuf_core_platform_payloads_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_protobuf_core_platform_payloads_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_core_platform_payloads_proto_goTypes = []interface{}{
	(ErrorSourceProto_ErrorSource)(0), // 0: tensorflow.core.platform.ErrorSourceProto.ErrorSource
	(*ErrorSourceProto)(nil),          // 1: tensorflow.core.platform.ErrorSourceProto
}
var file_tensorflow_core_protobuf_core_platform_payloads_proto_depIdxs = []int32{
	0, // 0: tensorflow.core.platform.ErrorSourceProto.error_source:type_name -> tensorflow.core.platform.ErrorSourceProto.ErrorSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_core_platform_payloads_proto_init() }
func file_tensorflow_core_protobuf_core_platform_payloads_proto_init() {
	if File_tensorflow_core_protobuf_core_platform_payloads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_core_platform_payloads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorSourceProto); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_core_platform_payloads_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_core_platform_payloads_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_protobuf_core_platform_payloads_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_protobuf_core_platform_payloads_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_core_platform_payloads_proto = out.File
	file_tensorflow_core_protobuf_core_platform_payloads_proto_rawDesc = nil
	file_tensorflow_core_protobuf_core_platform_payloads_proto_goTypes = nil
	file_tensorflow_core_protobuf_core_platform_payloads_proto_depIdxs = nil
}
