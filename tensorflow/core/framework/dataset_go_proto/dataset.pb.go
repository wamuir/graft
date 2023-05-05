// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.11
// source: tensorflow/core/framework/dataset.proto

package dataset_go_proto

import (
	tensor_go_proto "github.com/wamuir/graft/tensorflow/core/framework/tensor_go_proto"
	tensor_shape_go_proto "github.com/wamuir/graft/tensorflow/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/wamuir/graft/tensorflow/core/framework/types_go_proto"
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

// Metadata describing a compressed component of a dataset element.
type CompressedComponentMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dtype of the component tensor.
	Dtype types_go_proto.DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// The shape of the component tensor.
	TensorShape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=tensor_shape,json=tensorShape,proto3" json:"tensor_shape,omitempty"`
	// The amount of uncompressed tensor data.
	// - For string tensors, there is an element for each string indicating the
	// size of the string.
	// - For all other tensors, there is a single element indicating the size of
	// the tensor.
	UncompressedBytes []uint64 `protobuf:"varint,4,rep,packed,name=uncompressed_bytes,json=uncompressedBytes,proto3" json:"uncompressed_bytes,omitempty"`
}

func (x *CompressedComponentMetadata) Reset() {
	*x = CompressedComponentMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_dataset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressedComponentMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressedComponentMetadata) ProtoMessage() {}

func (x *CompressedComponentMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_dataset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressedComponentMetadata.ProtoReflect.Descriptor instead.
func (*CompressedComponentMetadata) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *CompressedComponentMetadata) GetDtype() types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return types_go_proto.DataType(0)
}

func (x *CompressedComponentMetadata) GetTensorShape() *tensor_shape_go_proto.TensorShapeProto {
	if x != nil {
		return x.TensorShape
	}
	return nil
}

func (x *CompressedComponentMetadata) GetUncompressedBytes() []uint64 {
	if x != nil {
		return x.UncompressedBytes
	}
	return nil
}

type CompressedElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Compressed tensor bytes for all components of the element.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Metadata for the components of the element.
	ComponentMetadata []*CompressedComponentMetadata `protobuf:"bytes,2,rep,name=component_metadata,json=componentMetadata,proto3" json:"component_metadata,omitempty"`
	// Version of the CompressedElement. CompressedElements may be stored on disk
	// and read back by later versions of code, so we store a version number to
	// help readers understand which version they are reading. When you add a new
	// field to this proto, you need to increment kCompressedElementVersion in
	// tensorflow/core/data/compression_utils.cc.
	Version int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CompressedElement) Reset() {
	*x = CompressedElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_dataset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressedElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressedElement) ProtoMessage() {}

func (x *CompressedElement) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_dataset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressedElement.ProtoReflect.Descriptor instead.
func (*CompressedElement) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_dataset_proto_rawDescGZIP(), []int{1}
}

func (x *CompressedElement) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CompressedElement) GetComponentMetadata() []*CompressedComponentMetadata {
	if x != nil {
		return x.ComponentMetadata
	}
	return nil
}

func (x *CompressedElement) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// An uncompressed dataset element.
type UncompressedElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Components []*tensor_go_proto.TensorProto `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *UncompressedElement) Reset() {
	*x = UncompressedElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_dataset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UncompressedElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncompressedElement) ProtoMessage() {}

func (x *UncompressedElement) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_dataset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncompressedElement.ProtoReflect.Descriptor instead.
func (*UncompressedElement) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_dataset_proto_rawDescGZIP(), []int{2}
}

func (x *UncompressedElement) GetComponents() []*tensor_go_proto.TensorProto {
	if x != nil {
		return x.Components
	}
	return nil
}

var File_tensorflow_core_framework_dataset_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_dataset_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x11, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x9e, 0x01, 0x0a, 0x11, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x5b, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x13, 0x55, 0x6e,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x53, 0x5a, 0x4e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_dataset_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_dataset_proto_rawDescData = file_tensorflow_core_framework_dataset_proto_rawDesc
)

func file_tensorflow_core_framework_dataset_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_dataset_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_dataset_proto_rawDescData)
	})
	return file_tensorflow_core_framework_dataset_proto_rawDescData
}

var file_tensorflow_core_framework_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_framework_dataset_proto_goTypes = []interface{}{
	(*CompressedComponentMetadata)(nil),            // 0: tensorflow.data.CompressedComponentMetadata
	(*CompressedElement)(nil),                      // 1: tensorflow.data.CompressedElement
	(*UncompressedElement)(nil),                    // 2: tensorflow.data.UncompressedElement
	(types_go_proto.DataType)(0),                   // 3: tensorflow.DataType
	(*tensor_shape_go_proto.TensorShapeProto)(nil), // 4: tensorflow.TensorShapeProto
	(*tensor_go_proto.TensorProto)(nil),            // 5: tensorflow.TensorProto
}
var file_tensorflow_core_framework_dataset_proto_depIdxs = []int32{
	3, // 0: tensorflow.data.CompressedComponentMetadata.dtype:type_name -> tensorflow.DataType
	4, // 1: tensorflow.data.CompressedComponentMetadata.tensor_shape:type_name -> tensorflow.TensorShapeProto
	0, // 2: tensorflow.data.CompressedElement.component_metadata:type_name -> tensorflow.data.CompressedComponentMetadata
	5, // 3: tensorflow.data.UncompressedElement.components:type_name -> tensorflow.TensorProto
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_dataset_proto_init() }
func file_tensorflow_core_framework_dataset_proto_init() {
	if File_tensorflow_core_framework_dataset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_dataset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressedComponentMetadata); i {
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
		file_tensorflow_core_framework_dataset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressedElement); i {
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
		file_tensorflow_core_framework_dataset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UncompressedElement); i {
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
			RawDescriptor: file_tensorflow_core_framework_dataset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_dataset_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_dataset_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_dataset_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_dataset_proto = out.File
	file_tensorflow_core_framework_dataset_proto_rawDesc = nil
	file_tensorflow_core_framework_dataset_proto_goTypes = nil
	file_tensorflow_core_framework_dataset_proto_depIdxs = nil
}
