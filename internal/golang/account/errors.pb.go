// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: protobuf/account/errors.proto

package account

import (
	_ "github.com/classtorch/go-error-generator-examples/internal/errors"
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

type ErrorCode int32

const (
	ErrorCode_SUCCESS       ErrorCode = 0     // 成功
	ErrorCode_UnKnown       ErrorCode = -1    // 账号不存在
	ErrorCode_InValid_Phone ErrorCode = 10001 // 登录失效，请重新登录
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:     "SUCCESS",
		-1:    "UnKnown",
		10001: "InValid_Phone",
	}
	ErrorCode_value = map[string]int32{
		"SUCCESS":       0,
		"UnKnown":       -1,
		"InValid_Phone": 10001,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_account_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_protobuf_account_errors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_account_errors_proto_rawDescGZIP(), []int{0}
}

var File_protobuf_account_errors_proto protoreflect.FileDescriptor

var file_protobuf_account_errors_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x75, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa6, 0x01, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x1a, 0x13, 0xa2, 0x45, 0x06, 0xe6, 0x88, 0x90, 0xe5,
	0x8a, 0x9f, 0xaa, 0x45, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x07,
	0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x1a, 0x1d, 0xa2, 0x45, 0x0c, 0xe6, 0x9c, 0xaa, 0xe7, 0x9f, 0xa5, 0xe9, 0x94, 0x99,
	0xe8, 0xaf, 0xaf, 0xaa, 0x45, 0x0b, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x20, 0x65, 0x72,
	0x72, 0x12, 0x42, 0x0a, 0x0d, 0x49, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x10, 0x91, 0x4e, 0x1a, 0x2e, 0xa2, 0x45, 0x18, 0xe6, 0x89, 0x8b, 0xe6, 0x9c, 0xba,
	0xe5, 0x8f, 0xb7, 0xe6, 0xa0, 0xbc, 0xe5, 0xbc, 0x8f, 0xe4, 0xb8, 0x8d, 0xe6, 0xad, 0xa3, 0xe7,
	0xa1, 0xae, 0xaa, 0x45, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x20, 0x6e, 0x6f, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_account_errors_proto_rawDescOnce sync.Once
	file_protobuf_account_errors_proto_rawDescData = file_protobuf_account_errors_proto_rawDesc
)

func file_protobuf_account_errors_proto_rawDescGZIP() []byte {
	file_protobuf_account_errors_proto_rawDescOnce.Do(func() {
		file_protobuf_account_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_account_errors_proto_rawDescData)
	})
	return file_protobuf_account_errors_proto_rawDescData
}

var file_protobuf_account_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_account_errors_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: uclass.service.account.ErrorCode
}
var file_protobuf_account_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_account_errors_proto_init() }
func file_protobuf_account_errors_proto_init() {
	if File_protobuf_account_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_account_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_account_errors_proto_goTypes,
		DependencyIndexes: file_protobuf_account_errors_proto_depIdxs,
		EnumInfos:         file_protobuf_account_errors_proto_enumTypes,
	}.Build()
	File_protobuf_account_errors_proto = out.File
	file_protobuf_account_errors_proto_rawDesc = nil
	file_protobuf_account_errors_proto_goTypes = nil
	file_protobuf_account_errors_proto_depIdxs = nil
}
