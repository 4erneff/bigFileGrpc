// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: proto/server.proto

package proto

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

type FileMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FileMetadataRequest) Reset() {
	*x = FileMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadataRequest) ProtoMessage() {}

func (x *FileMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadataRequest.ProtoReflect.Descriptor instead.
func (*FileMetadataRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartChunk int32 `protobuf:"varint,1,opt,name=start_chunk,json=startChunk,proto3" json:"start_chunk,omitempty"` // Starting chunk number for resuming downloads
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *FileRequest) GetStartChunk() int32 {
	if x != nil {
		return x.StartChunk
	}
	return 0
}

type FileMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSize   int64 `protobuf:"varint,1,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`       // Total size of the file in bytes
	TotalChunks int32 `protobuf:"varint,2,opt,name=total_chunks,json=totalChunks,proto3" json:"total_chunks,omitempty"` // Total number of chunks
}

func (x *FileMetadataResponse) Reset() {
	*x = FileMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadataResponse) ProtoMessage() {}

func (x *FileMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadataResponse.ProtoReflect.Descriptor instead.
func (*FileMetadataResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{2}
}

func (x *FileMetadataResponse) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *FileMetadataResponse) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceNumber int32  `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	ChunkData      []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
	TotalSize      int64  `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	Checksum       string `protobuf:"bytes,4,opt,name=checksum,proto3" json:"checksum,omitempty"`
	TotalChunks    int32  `protobuf:"varint,5,opt,name=total_chunks,json=totalChunks,proto3" json:"total_chunks,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *FileChunk) GetSequenceNumber() int32 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *FileChunk) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

func (x *FileChunk) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *FileChunk) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

func (x *FileChunk) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

var File_proto_server_proto protoreflect.FileDescriptor

var file_proto_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x15, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x58, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x32, 0xaa, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x34, 0x65, 0x72, 0x6e, 0x65, 0x66, 0x66, 0x2f, 0x61, 0x6c, 0x63, 0x61, 0x74, 0x72,
	0x61, 0x7a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData = file_proto_server_proto_rawDesc
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_proto_rawDescData)
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_server_proto_goTypes = []any{
	(*FileMetadataRequest)(nil),  // 0: fileservice.FileMetadataRequest
	(*FileRequest)(nil),          // 1: fileservice.FileRequest
	(*FileMetadataResponse)(nil), // 2: fileservice.FileMetadataResponse
	(*FileChunk)(nil),            // 3: fileservice.FileChunk
}
var file_proto_server_proto_depIdxs = []int32{
	0, // 0: fileservice.FileService.GetFileMetadata:input_type -> fileservice.FileMetadataRequest
	1, // 1: fileservice.FileService.GetFileStream:input_type -> fileservice.FileRequest
	2, // 2: fileservice.FileService.GetFileMetadata:output_type -> fileservice.FileMetadataResponse
	3, // 3: fileservice.FileService.GetFileStream:output_type -> fileservice.FileChunk
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FileMetadataRequest); i {
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
		file_proto_server_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FileRequest); i {
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
		file_proto_server_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FileMetadataResponse); i {
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
		file_proto_server_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FileChunk); i {
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
			RawDescriptor: file_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_rawDesc = nil
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}
