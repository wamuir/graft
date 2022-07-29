// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: tensorflow/core/protobuf/verifier_config.proto

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

type VerifierConfig_Toggle int32

const (
	VerifierConfig_DEFAULT VerifierConfig_Toggle = 0
	VerifierConfig_ON      VerifierConfig_Toggle = 1
	VerifierConfig_OFF     VerifierConfig_Toggle = 2
)

// Enum value maps for VerifierConfig_Toggle.
var (
	VerifierConfig_Toggle_name = map[int32]string{
		0: "DEFAULT",
		1: "ON",
		2: "OFF",
	}
	VerifierConfig_Toggle_value = map[string]int32{
		"DEFAULT": 0,
		"ON":      1,
		"OFF":     2,
	}
)

func (x VerifierConfig_Toggle) Enum() *VerifierConfig_Toggle {
	p := new(VerifierConfig_Toggle)
	*p = x
	return p
}

func (x VerifierConfig_Toggle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerifierConfig_Toggle) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_verifier_config_proto_enumTypes[0].Descriptor()
}

func (VerifierConfig_Toggle) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_verifier_config_proto_enumTypes[0]
}

func (x VerifierConfig_Toggle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerifierConfig_Toggle.Descriptor instead.
func (VerifierConfig_Toggle) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_verifier_config_proto_rawDescGZIP(), []int{0, 0}
}

// The config for graph verifiers.
type VerifierConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deadline for completion of all verification i.e. all the Toggle ON
	// verifiers must complete execution within this time.
	VerificationTimeoutInMs int64 `protobuf:"varint,1,opt,name=verification_timeout_in_ms,json=verificationTimeoutInMs,proto3" json:"verification_timeout_in_ms,omitempty"`
	// Perform structural validation on a tensorflow graph. Default is OFF.
	StructureVerifier VerifierConfig_Toggle `protobuf:"varint,2,opt,name=structure_verifier,json=structureVerifier,proto3,enum=tensorflow.VerifierConfig_Toggle" json:"structure_verifier,omitempty"`
}

func (x *VerifierConfig) Reset() {
	*x = VerifierConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_verifier_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifierConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifierConfig) ProtoMessage() {}

func (x *VerifierConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_verifier_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifierConfig.ProtoReflect.Descriptor instead.
func (*VerifierConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_verifier_config_proto_rawDescGZIP(), []int{0}
}

func (x *VerifierConfig) GetVerificationTimeoutInMs() int64 {
	if x != nil {
		return x.VerificationTimeoutInMs
	}
	return 0
}

func (x *VerifierConfig) GetStructureVerifier() VerifierConfig_Toggle {
	if x != nil {
		return x.StructureVerifier
	}
	return VerifierConfig_DEFAULT
}

var File_tensorflow_core_protobuf_verifier_config_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_verifier_config_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xc7, 0x01, 0x0a,
	0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3b, 0x0a, 0x1a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x17, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49, 0x6e, 0x4d, 0x73, 0x12, 0x50, 0x0a, 0x12,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x52, 0x11, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x26,
	0x0a, 0x06, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x42, 0x8c, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x42, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_verifier_config_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_verifier_config_proto_rawDescData = file_tensorflow_core_protobuf_verifier_config_proto_rawDesc
)

func file_tensorflow_core_protobuf_verifier_config_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_verifier_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_verifier_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_verifier_config_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_verifier_config_proto_rawDescData
}

var file_tensorflow_core_protobuf_verifier_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_protobuf_verifier_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_verifier_config_proto_goTypes = []interface{}{
	(VerifierConfig_Toggle)(0), // 0: tensorflow.VerifierConfig.Toggle
	(*VerifierConfig)(nil),     // 1: tensorflow.VerifierConfig
}
var file_tensorflow_core_protobuf_verifier_config_proto_depIdxs = []int32{
	0, // 0: tensorflow.VerifierConfig.structure_verifier:type_name -> tensorflow.VerifierConfig.Toggle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_verifier_config_proto_init() }
func file_tensorflow_core_protobuf_verifier_config_proto_init() {
	if File_tensorflow_core_protobuf_verifier_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_verifier_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifierConfig); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_verifier_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_verifier_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_verifier_config_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_protobuf_verifier_config_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_protobuf_verifier_config_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_verifier_config_proto = out.File
	file_tensorflow_core_protobuf_verifier_config_proto_rawDesc = nil
	file_tensorflow_core_protobuf_verifier_config_proto_goTypes = nil
	file_tensorflow_core_protobuf_verifier_config_proto_depIdxs = nil
}
