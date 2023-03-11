// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/getCaptcha.proto

package getCaptcha

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

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_getCaptcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_getCaptcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_proto_getCaptcha_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img []byte `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_getCaptcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_getCaptcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_proto_getCaptcha_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetImg() []byte {
	if x != nil {
		return x.Img
	}
	return nil
}

var File_proto_getCaptcha_proto protoreflect.FileDescriptor

var file_proto_getCaptcha_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x65, 0x74, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x32, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x3b, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12,
	0x17, 0x2e, 0x67, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x65, 0x74, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x67, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_getCaptcha_proto_rawDescOnce sync.Once
	file_proto_getCaptcha_proto_rawDescData = file_proto_getCaptcha_proto_rawDesc
)

func file_proto_getCaptcha_proto_rawDescGZIP() []byte {
	file_proto_getCaptcha_proto_rawDescOnce.Do(func() {
		file_proto_getCaptcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_getCaptcha_proto_rawDescData)
	})
	return file_proto_getCaptcha_proto_rawDescData
}

var file_proto_getCaptcha_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_getCaptcha_proto_goTypes = []interface{}{
	(*CallRequest)(nil),  // 0: getCaptcha.CallRequest
	(*CallResponse)(nil), // 1: getCaptcha.CallResponse
}
var file_proto_getCaptcha_proto_depIdxs = []int32{
	0, // 0: getCaptcha.GetCaptcha.Call:input_type -> getCaptcha.CallRequest
	1, // 1: getCaptcha.GetCaptcha.Call:output_type -> getCaptcha.CallResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_getCaptcha_proto_init() }
func file_proto_getCaptcha_proto_init() {
	if File_proto_getCaptcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_getCaptcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_proto_getCaptcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
			RawDescriptor: file_proto_getCaptcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_getCaptcha_proto_goTypes,
		DependencyIndexes: file_proto_getCaptcha_proto_depIdxs,
		MessageInfos:      file_proto_getCaptcha_proto_msgTypes,
	}.Build()
	File_proto_getCaptcha_proto = out.File
	file_proto_getCaptcha_proto_rawDesc = nil
	file_proto_getCaptcha_proto_goTypes = nil
	file_proto_getCaptcha_proto_depIdxs = nil
}
