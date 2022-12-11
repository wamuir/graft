// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: tensorflow/tsl/protobuf/rpc_options.proto

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

// RPC options for distributed runtime.
type RPCOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, always use RPC to contact the session target.
	//
	// If false (the default option), TensorFlow may use an optimized
	// transport for client-master communication that avoids the RPC
	// stack. This option is primarily for used testing the RPC stack.
	UseRpcForInprocessMaster bool `protobuf:"varint,1,opt,name=use_rpc_for_inprocess_master,json=useRpcForInprocessMaster,proto3" json:"use_rpc_for_inprocess_master,omitempty"`
	// The compression algorithm to be used. One of "deflate", "gzip".
	CompressionAlgorithm string `protobuf:"bytes,2,opt,name=compression_algorithm,json=compressionAlgorithm,proto3" json:"compression_algorithm,omitempty"`
	// If compression_algorithm is set, the compression level to be used.
	// From 0 (no compression), up to 3.
	CompressionLevel int32 `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"`
	// Setting cache_rpc_response to true will enable sender side caching of
	// response for RecvTensorAsync and RecvBufAsync to allow receiver to retry
	// requests . This is only necessary when the network fabric is experiencing a
	// significant error rate.  Without it we'll fail a step on an network error,
	// while with it we'll be able to complete long steps (like complex
	// initializations) in the face of some network errors during RecvTensor.
	CacheRpcResponse bool `protobuf:"varint,4,opt,name=cache_rpc_response,json=cacheRpcResponse,proto3" json:"cache_rpc_response,omitempty"`
	// Disables TCP connection sharing when opening a new RPC channel.
	DisableSessionConnectionSharing bool `protobuf:"varint,5,opt,name=disable_session_connection_sharing,json=disableSessionConnectionSharing,proto3" json:"disable_session_connection_sharing,omitempty"`
	// Setting num_channels_per_target > 0 allows uses of multiple channels to
	// communicate to the same target. This can be used to improve the aggregate
	// throughput on high speed links (e.g 100G) where single connection is not
	// sufficient to maximize link utilization. Note that a single RPC only goes
	// on a single channel, this only helps in situations where there are multiple
	// transfers to the same target overlapping in time.
	NumChannelsPerTarget int32 `protobuf:"varint,6,opt,name=num_channels_per_target,json=numChannelsPerTarget,proto3" json:"num_channels_per_target,omitempty"`
}

func (x *RPCOptions) Reset() {
	*x = RPCOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_protobuf_rpc_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCOptions) ProtoMessage() {}

func (x *RPCOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_protobuf_rpc_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCOptions.ProtoReflect.Descriptor instead.
func (*RPCOptions) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescGZIP(), []int{0}
}

func (x *RPCOptions) GetUseRpcForInprocessMaster() bool {
	if x != nil {
		return x.UseRpcForInprocessMaster
	}
	return false
}

func (x *RPCOptions) GetCompressionAlgorithm() string {
	if x != nil {
		return x.CompressionAlgorithm
	}
	return ""
}

func (x *RPCOptions) GetCompressionLevel() int32 {
	if x != nil {
		return x.CompressionLevel
	}
	return 0
}

func (x *RPCOptions) GetCacheRpcResponse() bool {
	if x != nil {
		return x.CacheRpcResponse
	}
	return false
}

func (x *RPCOptions) GetDisableSessionConnectionSharing() bool {
	if x != nil {
		return x.DisableSessionConnectionSharing
	}
	return false
}

func (x *RPCOptions) GetNumChannelsPerTarget() int32 {
	if x != nil {
		return x.NumChannelsPerTarget
	}
	return 0
}

var File_tensorflow_tsl_protobuf_rpc_options_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_protobuf_rpc_options_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xe0, 0x02, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x1c, 0x75, 0x73, 0x65, 0x5f, 0x72, 0x70,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x75, 0x73,
	0x65, 0x52, 0x70, 0x63, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x2b, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x61, 0x63, 0x68, 0x65, 0x52, 0x70, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x22, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x35, 0x0a, 0x17, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x50, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescData = file_tensorflow_tsl_protobuf_rpc_options_proto_rawDesc
)

func file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescData)
	})
	return file_tensorflow_tsl_protobuf_rpc_options_proto_rawDescData
}

var file_tensorflow_tsl_protobuf_rpc_options_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_tsl_protobuf_rpc_options_proto_goTypes = []interface{}{
	(*RPCOptions)(nil), // 0: tensorflow.RPCOptions
}
var file_tensorflow_tsl_protobuf_rpc_options_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_protobuf_rpc_options_proto_init() }
func file_tensorflow_tsl_protobuf_rpc_options_proto_init() {
	if File_tensorflow_tsl_protobuf_rpc_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_protobuf_rpc_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCOptions); i {
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
			RawDescriptor: file_tensorflow_tsl_protobuf_rpc_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_protobuf_rpc_options_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_protobuf_rpc_options_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_protobuf_rpc_options_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_protobuf_rpc_options_proto = out.File
	file_tensorflow_tsl_protobuf_rpc_options_proto_rawDesc = nil
	file_tensorflow_tsl_protobuf_rpc_options_proto_goTypes = nil
	file_tensorflow_tsl_protobuf_rpc_options_proto_depIdxs = nil
}
