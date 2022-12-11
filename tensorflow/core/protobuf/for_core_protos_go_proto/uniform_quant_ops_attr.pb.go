// Copyright 2022 The TensorFlow Authors. All Rights Reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: tensorflow/core/util/quantization/uniform_quant_ops_attr.proto

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

// Describes the dimension numbers for Convolution op. Corresponds to
// ::mlir::mhlo::ConvDimensionNumbersAttr.
type UniformQuantizedConvolutionDimensionNumbersAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dimension that represents batch in the input.
	InputBatchDimension int64 `protobuf:"varint,1,opt,name=input_batch_dimension,json=inputBatchDimension,proto3" json:"input_batch_dimension,omitempty"`
	// The dimension that represents features in the input.
	InputFeatureDimension int64 `protobuf:"varint,2,opt,name=input_feature_dimension,json=inputFeatureDimension,proto3" json:"input_feature_dimension,omitempty"`
	// The dimensions that represents spatial dimensions in the input. Length must
	// be rank-2 for the tensor rank for Convolution op.
	InputSpatialDimensions []int64 `protobuf:"varint,3,rep,packed,name=input_spatial_dimensions,json=inputSpatialDimensions,proto3" json:"input_spatial_dimensions,omitempty"`
	// The dimension that represents input features in the kernel (rhs).
	KernelInputFeatureDimension int64 `protobuf:"varint,4,opt,name=kernel_input_feature_dimension,json=kernelInputFeatureDimension,proto3" json:"kernel_input_feature_dimension,omitempty"`
	// The dimension that represents output features in the kernel (rhs).
	KernelOutputFeatureDimension int64 `protobuf:"varint,5,opt,name=kernel_output_feature_dimension,json=kernelOutputFeatureDimension,proto3" json:"kernel_output_feature_dimension,omitempty"`
	// The dimensions that represents spatial dimensions in the kernel (rhs).
	// Length must be rank-2 for the tensor rank for Convolution op.
	KernelSpatialDimensions []int64 `protobuf:"varint,6,rep,packed,name=kernel_spatial_dimensions,json=kernelSpatialDimensions,proto3" json:"kernel_spatial_dimensions,omitempty"`
	// The dimension that represents batch in the output.
	OutputBatchDimension int64 `protobuf:"varint,7,opt,name=output_batch_dimension,json=outputBatchDimension,proto3" json:"output_batch_dimension,omitempty"`
	// The dimension that represents features in the output.
	OutputFeatureDimension int64 `protobuf:"varint,8,opt,name=output_feature_dimension,json=outputFeatureDimension,proto3" json:"output_feature_dimension,omitempty"`
	// The dimensions that represents spatial dimensions in the output. Length
	// must be rank-2 for the tensor rank for Convolution op.
	OutputSpatialDimensions []int64 `protobuf:"varint,9,rep,packed,name=output_spatial_dimensions,json=outputSpatialDimensions,proto3" json:"output_spatial_dimensions,omitempty"`
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) Reset() {
	*x = UniformQuantizedConvolutionDimensionNumbersAttr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniformQuantizedConvolutionDimensionNumbersAttr) ProtoMessage() {}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniformQuantizedConvolutionDimensionNumbersAttr.ProtoReflect.Descriptor instead.
func (*UniformQuantizedConvolutionDimensionNumbersAttr) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescGZIP(), []int{0}
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetInputBatchDimension() int64 {
	if x != nil {
		return x.InputBatchDimension
	}
	return 0
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetInputFeatureDimension() int64 {
	if x != nil {
		return x.InputFeatureDimension
	}
	return 0
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetInputSpatialDimensions() []int64 {
	if x != nil {
		return x.InputSpatialDimensions
	}
	return nil
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetKernelInputFeatureDimension() int64 {
	if x != nil {
		return x.KernelInputFeatureDimension
	}
	return 0
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetKernelOutputFeatureDimension() int64 {
	if x != nil {
		return x.KernelOutputFeatureDimension
	}
	return 0
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetKernelSpatialDimensions() []int64 {
	if x != nil {
		return x.KernelSpatialDimensions
	}
	return nil
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetOutputBatchDimension() int64 {
	if x != nil {
		return x.OutputBatchDimension
	}
	return 0
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetOutputFeatureDimension() int64 {
	if x != nil {
		return x.OutputFeatureDimension
	}
	return 0
}

func (x *UniformQuantizedConvolutionDimensionNumbersAttr) GetOutputSpatialDimensions() []int64 {
	if x != nil {
		return x.OutputSpatialDimensions
	}
	return nil
}

var File_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto protoreflect.FileDescriptor

var file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xcb, 0x04, 0x0a,
	0x2f, 0x55, 0x6e, 0x69, 0x66, 0x6f, 0x72, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x41, 0x74, 0x74, 0x72,
	0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x13, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x16,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x44, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x1f, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x1c, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x19, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x61,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x17, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x53, 0x70, 0x61,
	0x74, 0x69, 0x61, 0x6c, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34,
	0x0a, 0x16, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x19, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x17, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescOnce sync.Once
	file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescData = file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDesc
)

func file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescGZIP() []byte {
	file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescData)
	})
	return file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDescData
}

var file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_goTypes = []interface{}{
	(*UniformQuantizedConvolutionDimensionNumbersAttr)(nil), // 0: tensorflow.UniformQuantizedConvolutionDimensionNumbersAttr
}
var file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_init() }
func file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_init() {
	if File_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniformQuantizedConvolutionDimensionNumbersAttr); i {
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
			RawDescriptor: file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_msgTypes,
	}.Build()
	File_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto = out.File
	file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_rawDesc = nil
	file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_goTypes = nil
	file_tensorflow_core_util_quantization_uniform_quant_ops_attr_proto_depIdxs = nil
}
