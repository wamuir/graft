// Copyright 2016 The TensorFlow Authors. All Rights Reserved.
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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: tensorflow/core/protobuf/cluster.proto

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

// Defines a single job in a TensorFlow cluster.
type JobDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this job.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Mapping from task ID to "hostname:port" string.
	//
	// If the `name` field contains "worker", and the `tasks` map contains a
	// mapping from 7 to "example.org:2222", then the device prefix
	// "/job:worker/task:7" will be assigned to "example.org:2222".
	Tasks map[int32]string `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *JobDef) Reset() {
	*x = JobDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDef) ProtoMessage() {}

func (x *JobDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDef.ProtoReflect.Descriptor instead.
func (*JobDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *JobDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobDef) GetTasks() map[int32]string {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// Defines a TensorFlow cluster as a set of jobs.
type ClusterDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The jobs that comprise the cluster.
	Job []*JobDef `protobuf:"bytes,1,rep,name=job,proto3" json:"job,omitempty"`
}

func (x *ClusterDef) Reset() {
	*x = ClusterDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDef) ProtoMessage() {}

func (x *ClusterDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDef.ProtoReflect.Descriptor instead.
func (*ClusterDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterDef) GetJob() []*JobDef {
	if x != nil {
		return x.Job
	}
	return nil
}

var File_tensorflow_core_protobuf_cluster_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_cluster_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x22, 0x8b, 0x01, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x66, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x4a, 0x6f, 0x62, 0x44, 0x65, 0x66, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x32, 0x0a, 0x0a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66,
	0x12, 0x24, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65,
	0x66, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x42, 0x87, 0x01, 0x0a, 0x1a, 0x6f, 0x72, 0x67, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_cluster_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_cluster_proto_rawDescData = file_tensorflow_core_protobuf_cluster_proto_rawDesc
)

func file_tensorflow_core_protobuf_cluster_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_cluster_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_cluster_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_cluster_proto_rawDescData
}

var file_tensorflow_core_protobuf_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_protobuf_cluster_proto_goTypes = []interface{}{
	(*JobDef)(nil),     // 0: tensorflow.JobDef
	(*ClusterDef)(nil), // 1: tensorflow.ClusterDef
	nil,                // 2: tensorflow.JobDef.TasksEntry
}
var file_tensorflow_core_protobuf_cluster_proto_depIdxs = []int32{
	2, // 0: tensorflow.JobDef.tasks:type_name -> tensorflow.JobDef.TasksEntry
	0, // 1: tensorflow.ClusterDef.job:type_name -> tensorflow.JobDef
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_cluster_proto_init() }
func file_tensorflow_core_protobuf_cluster_proto_init() {
	if File_tensorflow_core_protobuf_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDef); i {
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
		file_tensorflow_core_protobuf_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDef); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_cluster_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_cluster_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_cluster_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_cluster_proto = out.File
	file_tensorflow_core_protobuf_cluster_proto_rawDesc = nil
	file_tensorflow_core_protobuf_cluster_proto_goTypes = nil
	file_tensorflow_core_protobuf_cluster_proto_depIdxs = nil
}
