// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: tensorflow/tsl/protobuf/histogram.proto

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

// Serialization format for histogram module in
// tsl/lib/histogram/histogram.h
type HistogramProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min        float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	Max        float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	Num        float64 `protobuf:"fixed64,3,opt,name=num,proto3" json:"num,omitempty"`
	Sum        float64 `protobuf:"fixed64,4,opt,name=sum,proto3" json:"sum,omitempty"`
	SumSquares float64 `protobuf:"fixed64,5,opt,name=sum_squares,json=sumSquares,proto3" json:"sum_squares,omitempty"`
	// Parallel arrays encoding the bucket boundaries and the bucket values.
	// bucket(i) is the count for the bucket i.  The range for
	// a bucket is:
	//
	//	i == 0:  -DBL_MAX .. bucket_limit(0)
	//	i != 0:  bucket_limit(i-1) .. bucket_limit(i)
	BucketLimit []float64 `protobuf:"fixed64,6,rep,packed,name=bucket_limit,json=bucketLimit,proto3" json:"bucket_limit,omitempty"`
	Bucket      []float64 `protobuf:"fixed64,7,rep,packed,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *HistogramProto) Reset() {
	*x = HistogramProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_protobuf_histogram_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramProto) ProtoMessage() {}

func (x *HistogramProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_protobuf_histogram_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramProto.ProtoReflect.Descriptor instead.
func (*HistogramProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_protobuf_histogram_proto_rawDescGZIP(), []int{0}
}

func (x *HistogramProto) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *HistogramProto) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *HistogramProto) GetNum() float64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *HistogramProto) GetSum() float64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *HistogramProto) GetSumSquares() float64 {
	if x != nil {
		return x.SumSquares
	}
	return 0
}

func (x *HistogramProto) GetBucketLimit() []float64 {
	if x != nil {
		return x.BucketLimit
	}
	return nil
}

func (x *HistogramProto) GetBucket() []float64 {
	if x != nil {
		return x.Bucket
	}
	return nil
}

var File_tensorflow_tsl_protobuf_histogram_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_protobuf_histogram_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x73, 0x75, 0x6d,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x6d, 0x5f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x75, 0x6d, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x06, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x42, 0x75, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x67, 0x6f, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_protobuf_histogram_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_protobuf_histogram_proto_rawDescData = file_tensorflow_tsl_protobuf_histogram_proto_rawDesc
)

func file_tensorflow_tsl_protobuf_histogram_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_protobuf_histogram_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_protobuf_histogram_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_protobuf_histogram_proto_rawDescData)
	})
	return file_tensorflow_tsl_protobuf_histogram_proto_rawDescData
}

var file_tensorflow_tsl_protobuf_histogram_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_tsl_protobuf_histogram_proto_goTypes = []interface{}{
	(*HistogramProto)(nil), // 0: tensorflow.HistogramProto
}
var file_tensorflow_tsl_protobuf_histogram_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_protobuf_histogram_proto_init() }
func file_tensorflow_tsl_protobuf_histogram_proto_init() {
	if File_tensorflow_tsl_protobuf_histogram_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_protobuf_histogram_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramProto); i {
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
			RawDescriptor: file_tensorflow_tsl_protobuf_histogram_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_protobuf_histogram_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_protobuf_histogram_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_protobuf_histogram_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_protobuf_histogram_proto = out.File
	file_tensorflow_tsl_protobuf_histogram_proto_rawDesc = nil
	file_tensorflow_tsl_protobuf_histogram_proto_goTypes = nil
	file_tensorflow_tsl_protobuf_histogram_proto_depIdxs = nil
}
