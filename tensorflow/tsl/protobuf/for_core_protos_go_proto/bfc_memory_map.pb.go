// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: tensorflow/tsl/protobuf/bfc_memory_map.proto

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

// Some of the data from AllocatorStats
type MemAllocatorStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumAllocs           int64   `protobuf:"varint,1,opt,name=num_allocs,json=numAllocs,proto3" json:"num_allocs,omitempty"`
	BytesInUse          int64   `protobuf:"varint,2,opt,name=bytes_in_use,json=bytesInUse,proto3" json:"bytes_in_use,omitempty"`
	PeakBytesInUse      int64   `protobuf:"varint,3,opt,name=peak_bytes_in_use,json=peakBytesInUse,proto3" json:"peak_bytes_in_use,omitempty"`
	LargestAllocSize    int64   `protobuf:"varint,4,opt,name=largest_alloc_size,json=largestAllocSize,proto3" json:"largest_alloc_size,omitempty"`
	FragmentationMetric float32 `protobuf:"fixed32,5,opt,name=fragmentation_metric,json=fragmentationMetric,proto3" json:"fragmentation_metric,omitempty"`
}

func (x *MemAllocatorStats) Reset() {
	*x = MemAllocatorStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemAllocatorStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemAllocatorStats) ProtoMessage() {}

func (x *MemAllocatorStats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemAllocatorStats.ProtoReflect.Descriptor instead.
func (*MemAllocatorStats) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescGZIP(), []int{0}
}

func (x *MemAllocatorStats) GetNumAllocs() int64 {
	if x != nil {
		return x.NumAllocs
	}
	return 0
}

func (x *MemAllocatorStats) GetBytesInUse() int64 {
	if x != nil {
		return x.BytesInUse
	}
	return 0
}

func (x *MemAllocatorStats) GetPeakBytesInUse() int64 {
	if x != nil {
		return x.PeakBytesInUse
	}
	return 0
}

func (x *MemAllocatorStats) GetLargestAllocSize() int64 {
	if x != nil {
		return x.LargestAllocSize
	}
	return 0
}

func (x *MemAllocatorStats) GetFragmentationMetric() float32 {
	if x != nil {
		return x.FragmentationMetric
	}
	return 0
}

type MemChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       uint64 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	Size          int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	RequestedSize int64  `protobuf:"varint,3,opt,name=requested_size,json=requestedSize,proto3" json:"requested_size,omitempty"`
	Bin           int32  `protobuf:"varint,4,opt,name=bin,proto3" json:"bin,omitempty"`
	OpName        string `protobuf:"bytes,5,opt,name=op_name,json=opName,proto3" json:"op_name,omitempty"`
	FreedAtCount  uint64 `protobuf:"varint,6,opt,name=freed_at_count,json=freedAtCount,proto3" json:"freed_at_count,omitempty"`
	ActionCount   uint64 `protobuf:"varint,7,opt,name=action_count,json=actionCount,proto3" json:"action_count,omitempty"`
	InUse         bool   `protobuf:"varint,8,opt,name=in_use,json=inUse,proto3" json:"in_use,omitempty"`
	StepId        uint64 `protobuf:"varint,9,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
}

func (x *MemChunk) Reset() {
	*x = MemChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemChunk) ProtoMessage() {}

func (x *MemChunk) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemChunk.ProtoReflect.Descriptor instead.
func (*MemChunk) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescGZIP(), []int{1}
}

func (x *MemChunk) GetAddress() uint64 {
	if x != nil {
		return x.Address
	}
	return 0
}

func (x *MemChunk) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MemChunk) GetRequestedSize() int64 {
	if x != nil {
		return x.RequestedSize
	}
	return 0
}

func (x *MemChunk) GetBin() int32 {
	if x != nil {
		return x.Bin
	}
	return 0
}

func (x *MemChunk) GetOpName() string {
	if x != nil {
		return x.OpName
	}
	return ""
}

func (x *MemChunk) GetFreedAtCount() uint64 {
	if x != nil {
		return x.FreedAtCount
	}
	return 0
}

func (x *MemChunk) GetActionCount() uint64 {
	if x != nil {
		return x.ActionCount
	}
	return 0
}

func (x *MemChunk) GetInUse() bool {
	if x != nil {
		return x.InUse
	}
	return false
}

func (x *MemChunk) GetStepId() uint64 {
	if x != nil {
		return x.StepId
	}
	return 0
}

type BinSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bin              int32 `protobuf:"varint,1,opt,name=bin,proto3" json:"bin,omitempty"`
	TotalBytesInUse  int64 `protobuf:"varint,2,opt,name=total_bytes_in_use,json=totalBytesInUse,proto3" json:"total_bytes_in_use,omitempty"`
	TotalBytesInBin  int64 `protobuf:"varint,3,opt,name=total_bytes_in_bin,json=totalBytesInBin,proto3" json:"total_bytes_in_bin,omitempty"`
	TotalChunksInUse int64 `protobuf:"varint,4,opt,name=total_chunks_in_use,json=totalChunksInUse,proto3" json:"total_chunks_in_use,omitempty"`
	TotalChunksInBin int64 `protobuf:"varint,5,opt,name=total_chunks_in_bin,json=totalChunksInBin,proto3" json:"total_chunks_in_bin,omitempty"`
}

func (x *BinSummary) Reset() {
	*x = BinSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinSummary) ProtoMessage() {}

func (x *BinSummary) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinSummary.ProtoReflect.Descriptor instead.
func (*BinSummary) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescGZIP(), []int{2}
}

func (x *BinSummary) GetBin() int32 {
	if x != nil {
		return x.Bin
	}
	return 0
}

func (x *BinSummary) GetTotalBytesInUse() int64 {
	if x != nil {
		return x.TotalBytesInUse
	}
	return 0
}

func (x *BinSummary) GetTotalBytesInBin() int64 {
	if x != nil {
		return x.TotalBytesInBin
	}
	return 0
}

func (x *BinSummary) GetTotalChunksInUse() int64 {
	if x != nil {
		return x.TotalChunksInUse
	}
	return 0
}

func (x *BinSummary) GetTotalChunksInBin() int64 {
	if x != nil {
		return x.TotalChunksInBin
	}
	return 0
}

type SnapShot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionCount uint64 `protobuf:"varint,1,opt,name=action_count,json=actionCount,proto3" json:"action_count,omitempty"`
	Size        int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SnapShot) Reset() {
	*x = SnapShot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapShot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapShot) ProtoMessage() {}

func (x *SnapShot) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapShot.ProtoReflect.Descriptor instead.
func (*SnapShot) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescGZIP(), []int{3}
}

func (x *SnapShot) GetActionCount() uint64 {
	if x != nil {
		return x.ActionCount
	}
	return 0
}

func (x *SnapShot) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type MemoryDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllocatorName string             `protobuf:"bytes,1,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	BinSummary    []*BinSummary      `protobuf:"bytes,2,rep,name=bin_summary,json=binSummary,proto3" json:"bin_summary,omitempty"`
	Chunk         []*MemChunk        `protobuf:"bytes,3,rep,name=chunk,proto3" json:"chunk,omitempty"`
	SnapShot      []*SnapShot        `protobuf:"bytes,4,rep,name=snap_shot,json=snapShot,proto3" json:"snap_shot,omitempty"`
	Stats         *MemAllocatorStats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *MemoryDump) Reset() {
	*x = MemoryDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryDump) ProtoMessage() {}

func (x *MemoryDump) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryDump.ProtoReflect.Descriptor instead.
func (*MemoryDump) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescGZIP(), []int{4}
}

func (x *MemoryDump) GetAllocatorName() string {
	if x != nil {
		return x.AllocatorName
	}
	return ""
}

func (x *MemoryDump) GetBinSummary() []*BinSummary {
	if x != nil {
		return x.BinSummary
	}
	return nil
}

func (x *MemoryDump) GetChunk() []*MemChunk {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *MemoryDump) GetSnapShot() []*SnapShot {
	if x != nil {
		return x.SnapShot
	}
	return nil
}

func (x *MemoryDump) GetStats() *MemAllocatorStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

var File_tensorflow_tsl_protobuf_bfc_memory_map_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x66, 0x63, 0x5f, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xe0, 0x01, 0x0a, 0x11, 0x4d,
	0x65, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x73, 0x12,
	0x20, 0x0a, 0x0c, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x55, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x11, 0x70, 0x65, 0x61, 0x6b, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x65,
	0x61, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x6c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x83, 0x02,
	0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x69,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x72,
	0x65, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x64, 0x41, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74,
	0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x65,
	0x70, 0x49, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x62, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x55, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x42, 0x69, 0x6e, 0x12, 0x2d,
	0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x69,
	0x6e, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x5f, 0x69, 0x6e,
	0x5f, 0x62, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x42, 0x69, 0x6e, 0x22, 0x41, 0x0a, 0x08,
	0x53, 0x6e, 0x61, 0x70, 0x53, 0x68, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x80, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x42, 0x69, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x65, 0x6d, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x6e,
	0x61, 0x70, 0x5f, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x53,
	0x68, 0x6f, 0x74, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x53, 0x68, 0x6f, 0x74, 0x12, 0x33, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x65, 0x6d, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescData = file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDesc
)

func file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescData)
	})
	return file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDescData
}

var file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_tsl_protobuf_bfc_memory_map_proto_goTypes = []interface{}{
	(*MemAllocatorStats)(nil), // 0: tensorflow.MemAllocatorStats
	(*MemChunk)(nil),          // 1: tensorflow.MemChunk
	(*BinSummary)(nil),        // 2: tensorflow.BinSummary
	(*SnapShot)(nil),          // 3: tensorflow.SnapShot
	(*MemoryDump)(nil),        // 4: tensorflow.MemoryDump
}
var file_tensorflow_tsl_protobuf_bfc_memory_map_proto_depIdxs = []int32{
	2, // 0: tensorflow.MemoryDump.bin_summary:type_name -> tensorflow.BinSummary
	1, // 1: tensorflow.MemoryDump.chunk:type_name -> tensorflow.MemChunk
	3, // 2: tensorflow.MemoryDump.snap_shot:type_name -> tensorflow.SnapShot
	0, // 3: tensorflow.MemoryDump.stats:type_name -> tensorflow.MemAllocatorStats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_protobuf_bfc_memory_map_proto_init() }
func file_tensorflow_tsl_protobuf_bfc_memory_map_proto_init() {
	if File_tensorflow_tsl_protobuf_bfc_memory_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemAllocatorStats); i {
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
		file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemChunk); i {
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
		file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinSummary); i {
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
		file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapShot); i {
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
		file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryDump); i {
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
			RawDescriptor: file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_protobuf_bfc_memory_map_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_protobuf_bfc_memory_map_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_protobuf_bfc_memory_map_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_protobuf_bfc_memory_map_proto = out.File
	file_tensorflow_tsl_protobuf_bfc_memory_map_proto_rawDesc = nil
	file_tensorflow_tsl_protobuf_bfc_memory_map_proto_goTypes = nil
	file_tensorflow_tsl_protobuf_bfc_memory_map_proto_depIdxs = nil
}
