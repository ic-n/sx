// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commitreveal/v1/api.proto

package commitrevealv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commitreveal_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commitreveal_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_commitreveal_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *HealthResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commitreveal_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commitreveal_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_commitreveal_v1_api_proto_rawDescGZIP(), []int{1}
}

var File_commitreveal_v1_api_proto protoreflect.FileDescriptor

var file_commitreveal_v1_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x0f, 0x0a, 0x0d,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x74, 0x0a,
	0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x42, 0xb3, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x2d, 0x6e, 0x2f, 0x73, 0x78, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72, 0x65, 0x76,
	0x65, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72,
	0x65, 0x76, 0x65, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x72,
	0x65, 0x76, 0x65, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commitreveal_v1_api_proto_rawDescOnce sync.Once
	file_commitreveal_v1_api_proto_rawDescData = file_commitreveal_v1_api_proto_rawDesc
)

func file_commitreveal_v1_api_proto_rawDescGZIP() []byte {
	file_commitreveal_v1_api_proto_rawDescOnce.Do(func() {
		file_commitreveal_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_commitreveal_v1_api_proto_rawDescData)
	})
	return file_commitreveal_v1_api_proto_rawDescData
}

var file_commitreveal_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commitreveal_v1_api_proto_goTypes = []interface{}{
	(*HealthResponse)(nil), // 0: commitreveal.v1.HealthResponse
	(*HealthRequest)(nil),  // 1: commitreveal.v1.HealthRequest
}
var file_commitreveal_v1_api_proto_depIdxs = []int32{
	1, // 0: commitreveal.v1.CommitRevealService.Health:input_type -> commitreveal.v1.HealthRequest
	0, // 1: commitreveal.v1.CommitRevealService.Health:output_type -> commitreveal.v1.HealthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commitreveal_v1_api_proto_init() }
func file_commitreveal_v1_api_proto_init() {
	if File_commitreveal_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commitreveal_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
		file_commitreveal_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
			RawDescriptor: file_commitreveal_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commitreveal_v1_api_proto_goTypes,
		DependencyIndexes: file_commitreveal_v1_api_proto_depIdxs,
		MessageInfos:      file_commitreveal_v1_api_proto_msgTypes,
	}.Build()
	File_commitreveal_v1_api_proto = out.File
	file_commitreveal_v1_api_proto_rawDesc = nil
	file_commitreveal_v1_api_proto_goTypes = nil
	file_commitreveal_v1_api_proto_depIdxs = nil
}
