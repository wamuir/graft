// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: tensorflow/core/protobuf/named_tensor.proto

package for_core_protos_go_proto

import (
	tensor_go_proto "github.com/wamuir/graft/tensorflow/core/framework/tensor_go_proto"
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

// A pair of tensor name and tensor values.
type NamedTensorProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the tensor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The client can populate a TensorProto using a tensorflow::Tensor`, or
	// directly using the protobuf field accessors.
	//
	// The client specifies whether the returned tensor values should be
	// filled tensor fields (float_val, int_val, etc.) or encoded in a
	// compact form in tensor.tensor_content.
	Tensor *tensor_go_proto.TensorProto `protobuf:"bytes,2,opt,name=tensor,proto3" json:"tensor,omitempty"`
}

func (x *NamedTensorProto) Reset() {
	*x = NamedTensorProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_named_tensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamedTensorProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedTensorProto) ProtoMessage() {}

func (x *NamedTensorProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_named_tensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedTensorProto.ProtoReflect.Descriptor instead.
func (*NamedTensorProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_named_tensor_proto_rawDescGZIP(), []int{0}
}

func (x *NamedTensorProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamedTensorProto) GetTensor() *tensor_go_proto.TensorProto {
	if x != nil {
		return x.Tensor
	}
	return nil
}

var File_tensorflow_core_protobuf_named_tensor_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_named_tensor_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x64,
	0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x57, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x42, 0x89, 0x01, 0x0a, 0x18, 0x6f,
	0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x11, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_named_tensor_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_named_tensor_proto_rawDescData = file_tensorflow_core_protobuf_named_tensor_proto_rawDesc
)

func file_tensorflow_core_protobuf_named_tensor_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_named_tensor_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_named_tensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_named_tensor_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_named_tensor_proto_rawDescData
}

var file_tensorflow_core_protobuf_named_tensor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_named_tensor_proto_goTypes = []interface{}{
	(*NamedTensorProto)(nil),            // 0: tensorflow.NamedTensorProto
	(*tensor_go_proto.TensorProto)(nil), // 1: tensorflow.TensorProto
}
var file_tensorflow_core_protobuf_named_tensor_proto_depIdxs = []int32{
	1, // 0: tensorflow.NamedTensorProto.tensor:type_name -> tensorflow.TensorProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_named_tensor_proto_init() }
func file_tensorflow_core_protobuf_named_tensor_proto_init() {
	if File_tensorflow_core_protobuf_named_tensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_named_tensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamedTensorProto); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_named_tensor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_named_tensor_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_named_tensor_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_named_tensor_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_named_tensor_proto = out.File
	file_tensorflow_core_protobuf_named_tensor_proto_rawDesc = nil
	file_tensorflow_core_protobuf_named_tensor_proto_goTypes = nil
	file_tensorflow_core_protobuf_named_tensor_proto_depIdxs = nil
}
