// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: inventory_main.proto

package inventory_service

import (
	common "genproto/common"
	order_service "genproto/order_service"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_inventory_main_proto protoreflect.FileDescriptor

var file_inventory_main_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65,
	0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcc, 0x15,
	0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x0c, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0a, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x0e,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12,
	0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x18,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x47, 0x0a, 0x17,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x50, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x1d, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49,
	0x44, 0x12, 0x3f, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x44, 0x12, 0x4c, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x69, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44,
	0x12, 0x2e, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44,
	0x12, 0x31, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x73, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x78, 0x65, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x1a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x45, 0x78, 0x65, 0x6c,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x41, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12,
	0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66,
	0x66, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12,
	0x41, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x6f, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x12, 0x1c, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x54, 0x6f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x44, 0x12, 0x42, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x12, 0x1c, 0x2e, 0x41, 0x64,
	0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x6f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x12, 0x12, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x12, 0x12, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66,
	0x52, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x42, 0x79, 0x69, 0x64, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x4f, 0x66, 0x66, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x39,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x19,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x44, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x1a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x6f, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49,
	0x44, 0x12, 0x37, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x0a, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x40, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x17, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x42, 0x1c, 0x5a, 0x1a,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_inventory_main_proto_goTypes = []interface{}{
	(*CreateImportRequest)(nil),                  // 0: CreateImportRequest
	(*order_service.CreateOrderCopyRequest)(nil), // 1: CreateOrderCopyRequest
	(*FinishImportReq)(nil),                      // 2: FinishImportReq
	(*common.SearchRequest)(nil),                 // 3: SearchRequest
	(*common.RequestID)(nil),                     // 4: RequestID
	(*CreateSupplierRequest)(nil),                // 5: CreateSupplierRequest
	(*UpdateSupplierRequest)(nil),                // 6: UpdateSupplierRequest
	(*CreateSupplierOrderRequest)(nil),           // 7: CreateSupplierOrderRequest
	(*CreateSupplierOrderItemRequest)(nil),       // 8: CreateSupplierOrderItemRequest
	(*common.SupplierOrderSearchRequest)(nil),    // 9: SupplierOrderSearchRequest
	(*GetAllSupplierOrderItemsRequest)(nil),      // 10: GetAllSupplierOrderItemsRequest
	(*FinishSupplierOrderRequest)(nil),           // 11: FinishSupplierOrderRequest
	(*UpdateSupplierOrderRecivedRequest)(nil),    // 12: UpdateSupplierOrderRecivedRequest
	(*common.RequestIDs)(nil),                    // 13: RequestIDs
	(*GetProductHistoryReq)(nil),                 // 14: GetProductHistoryReq
	(*GetWriteOffReq)(nil),                       // 15: GetWriteOffReq
	(*CreateTransferRequest)(nil),                // 16: CreateTransferRequest
	(*GetAllTransferRequest)(nil),                // 17: GetAllTransferRequest
	(*GetAllTransferItemsRequest)(nil),           // 18: GetAllTransferItemsRequest
	(*AddItemToTransferRequest)(nil),             // 19: AddItemToTransferRequest
	(*CreateWriteOffRequest)(nil),                // 20: CreateWriteOffRequest
	(*UpdateWriteOffRequest)(nil),                // 21: UpdateWriteOffRequest
	(*AddProductToWriteOffRequest)(nil),          // 22: AddProductToWriteOffRequest
	(*FinishWriteOffReq)(nil),                    // 23: FinishWriteOffReq
	(*GetAllWriteOffReq)(nil),                    // 24: GetAllWriteOffReq
	(*CreateRepricingRequest)(nil),               // 25: CreateRepricingRequest
	(*GetAllRepricingRequest)(nil),               // 26: GetAllRepricingRequest
	(*GetAllRepricingItemsRequest)(nil),          // 27: GetAllRepricingItemsRequest
	(*UpsertMultiItemRequest)(nil),               // 28: UpsertMultiItemRequest
	(*FinishRepricingRequest)(nil),               // 29: FinishRepricingRequest
	(*BulkChangePriceRequest)(nil),               // 30: BulkChangePriceRequest
	(*common.ItemIdRequest)(nil),                 // 31: ItemIdRequest
	(*common.ResponseID)(nil),                    // 32: ResponseID
	(*GetAllImportRes)(nil),                      // 33: GetAllImportRes
	(*ImportById)(nil),                           // 34: ImportById
	(*GetAllSuppliersResponse)(nil),              // 35: GetAllSuppliersResponse
	(*GetSupplierByIdResponse)(nil),              // 36: GetSupplierByIdResponse
	(*GetAllSupplierOrderResponse)(nil),          // 37: GetAllSupplierOrderResponse
	(*GetAllSupplierOrderItemsResponse)(nil),     // 38: GetAllSupplierOrderItemsResponse
	(*GetSupplierOrderByIdResponse)(nil),         // 39: GetSupplierOrderByIdResponse
	(*common.Empty)(nil),                         // 40: Empty
	(*GetProductHistoryRes)(nil),                 // 41: GetProductHistoryRes
	(*GetAllTransferResponse)(nil),               // 42: GetAllTransferResponse
	(*GetAllTransferItemsResponse)(nil),          // 43: GetAllTransferItemsResponse
	(*GetAllWriteOffRes)(nil),                    // 44: GetAllWriteOffRes
	(*GetWriteOffRes)(nil),                       // 45: GetWriteOffRes
	(*GetWriteOffItemsRes)(nil),                  // 46: GetWriteOffItemsRes
	(*GetRepricingByIDResponse)(nil),             // 47: GetRepricingByIDResponse
	(*GetAllRepricingResponse)(nil),              // 48: GetAllRepricingResponse
	(*GetAllRepricingItemsResponse)(nil),         // 49: GetAllRepricingItemsResponse
}
var file_inventory_main_proto_depIdxs = []int32{
	0,  // 0: InventoryService.CreateImport:input_type -> CreateImportRequest
	1,  // 1: InventoryService.CreateOrder:input_type -> CreateOrderCopyRequest
	2,  // 2: InventoryService.FinishImport:input_type -> FinishImportReq
	3,  // 3: InventoryService.GetAllImport:input_type -> SearchRequest
	4,  // 4: InventoryService.GetImportByID:input_type -> RequestID
	5,  // 5: InventoryService.CreateSupplier:input_type -> CreateSupplierRequest
	3,  // 6: InventoryService.GetAllSupplier:input_type -> SearchRequest
	6,  // 7: InventoryService.UpdateSupplier:input_type -> UpdateSupplierRequest
	4,  // 8: InventoryService.GetSupplierById:input_type -> RequestID
	4,  // 9: InventoryService.Delete:input_type -> RequestID
	7,  // 10: InventoryService.CreateSupplierOrder:input_type -> CreateSupplierOrderRequest
	8,  // 11: InventoryService.UpsertSupplierOrderItem:input_type -> CreateSupplierOrderItemRequest
	9,  // 12: InventoryService.GetAllSupplierOrder:input_type -> SupplierOrderSearchRequest
	10, // 13: InventoryService.GetAllSupplierOrderItems:input_type -> GetAllSupplierOrderItemsRequest
	4,  // 14: InventoryService.GetSupplierOrderById:input_type -> RequestID
	4,  // 15: InventoryService.UpdateSupplierOrderStatus:input_type -> RequestID
	11, // 16: InventoryService.FinishSupplierOrder:input_type -> FinishSupplierOrderRequest
	12, // 17: InventoryService.UpdateSupplierOrderAmount:input_type -> UpdateSupplierOrderRecivedRequest
	4,  // 18: InventoryService.DeleteSupplierOrder:input_type -> RequestID
	4,  // 19: InventoryService.DeleteSupplierOrderItemById:input_type -> RequestID
	13, // 20: InventoryService.DeleteSupplierOrders:input_type -> RequestIDs
	14, // 21: InventoryService.GetProductHistory:input_type -> GetProductHistoryReq
	4,  // 22: InventoryService.CreateDownloadExelTemplate:input_type -> RequestID
	15, // 23: InventoryService.CreateWriteOffExelTemplate:input_type -> GetWriteOffReq
	16, // 24: InventoryService.CreateTransfer:input_type -> CreateTransferRequest
	17, // 25: InventoryService.GetAllTransfer:input_type -> GetAllTransferRequest
	18, // 26: InventoryService.GetAllTransferItems:input_type -> GetAllTransferItemsRequest
	19, // 27: InventoryService.AddItemToTransfer:input_type -> AddItemToTransferRequest
	20, // 28: InventoryService.CreateWriteOff:input_type -> CreateWriteOffRequest
	21, // 29: InventoryService.UpdateToWriteOff:input_type -> UpdateWriteOffRequest
	22, // 30: InventoryService.AddProductToWriteOff:input_type -> AddProductToWriteOffRequest
	22, // 31: InventoryService.DeleteProductWriteOff:input_type -> AddProductToWriteOffRequest
	23, // 32: InventoryService.FinishWriteOff:input_type -> FinishWriteOffReq
	24, // 33: InventoryService.GetAllWriteOff:input_type -> GetAllWriteOffReq
	15, // 34: InventoryService.GetWriteOffByid:input_type -> GetWriteOffReq
	15, // 35: InventoryService.GetWriteOffItems:input_type -> GetWriteOffReq
	4,  // 36: InventoryService.DeleteWriteOff:input_type -> RequestID
	4,  // 37: InventoryService.GetRepricingByID:input_type -> RequestID
	25, // 38: InventoryService.CreateRepricing:input_type -> CreateRepricingRequest
	26, // 39: InventoryService.GetAllRepricing:input_type -> GetAllRepricingRequest
	27, // 40: InventoryService.GetAllRepricingItems:input_type -> GetAllRepricingItemsRequest
	28, // 41: InventoryService.UpsertMultiItemToRepricing:input_type -> UpsertMultiItemRequest
	29, // 42: InventoryService.FinishRepricing:input_type -> FinishRepricingRequest
	4,  // 43: InventoryService.DeleteRepricing:input_type -> RequestID
	30, // 44: InventoryService.RepricingBulkChangePrice:input_type -> BulkChangePriceRequest
	31, // 45: InventoryService.DeleteRepricingItem:input_type -> ItemIdRequest
	32, // 46: InventoryService.CreateImport:output_type -> ResponseID
	32, // 47: InventoryService.CreateOrder:output_type -> ResponseID
	32, // 48: InventoryService.FinishImport:output_type -> ResponseID
	33, // 49: InventoryService.GetAllImport:output_type -> GetAllImportRes
	34, // 50: InventoryService.GetImportByID:output_type -> ImportById
	32, // 51: InventoryService.CreateSupplier:output_type -> ResponseID
	35, // 52: InventoryService.GetAllSupplier:output_type -> GetAllSuppliersResponse
	32, // 53: InventoryService.UpdateSupplier:output_type -> ResponseID
	36, // 54: InventoryService.GetSupplierById:output_type -> GetSupplierByIdResponse
	32, // 55: InventoryService.Delete:output_type -> ResponseID
	32, // 56: InventoryService.CreateSupplierOrder:output_type -> ResponseID
	32, // 57: InventoryService.UpsertSupplierOrderItem:output_type -> ResponseID
	37, // 58: InventoryService.GetAllSupplierOrder:output_type -> GetAllSupplierOrderResponse
	38, // 59: InventoryService.GetAllSupplierOrderItems:output_type -> GetAllSupplierOrderItemsResponse
	39, // 60: InventoryService.GetSupplierOrderById:output_type -> GetSupplierOrderByIdResponse
	32, // 61: InventoryService.UpdateSupplierOrderStatus:output_type -> ResponseID
	32, // 62: InventoryService.FinishSupplierOrder:output_type -> ResponseID
	32, // 63: InventoryService.UpdateSupplierOrderAmount:output_type -> ResponseID
	32, // 64: InventoryService.DeleteSupplierOrder:output_type -> ResponseID
	40, // 65: InventoryService.DeleteSupplierOrderItemById:output_type -> Empty
	40, // 66: InventoryService.DeleteSupplierOrders:output_type -> Empty
	41, // 67: InventoryService.GetProductHistory:output_type -> GetProductHistoryRes
	32, // 68: InventoryService.CreateDownloadExelTemplate:output_type -> ResponseID
	32, // 69: InventoryService.CreateWriteOffExelTemplate:output_type -> ResponseID
	32, // 70: InventoryService.CreateTransfer:output_type -> ResponseID
	42, // 71: InventoryService.GetAllTransfer:output_type -> GetAllTransferResponse
	43, // 72: InventoryService.GetAllTransferItems:output_type -> GetAllTransferItemsResponse
	32, // 73: InventoryService.AddItemToTransfer:output_type -> ResponseID
	32, // 74: InventoryService.CreateWriteOff:output_type -> ResponseID
	32, // 75: InventoryService.UpdateToWriteOff:output_type -> ResponseID
	32, // 76: InventoryService.AddProductToWriteOff:output_type -> ResponseID
	32, // 77: InventoryService.DeleteProductWriteOff:output_type -> ResponseID
	32, // 78: InventoryService.FinishWriteOff:output_type -> ResponseID
	44, // 79: InventoryService.GetAllWriteOff:output_type -> GetAllWriteOffRes
	45, // 80: InventoryService.GetWriteOffByid:output_type -> GetWriteOffRes
	46, // 81: InventoryService.GetWriteOffItems:output_type -> GetWriteOffItemsRes
	32, // 82: InventoryService.DeleteWriteOff:output_type -> ResponseID
	47, // 83: InventoryService.GetRepricingByID:output_type -> GetRepricingByIDResponse
	32, // 84: InventoryService.CreateRepricing:output_type -> ResponseID
	48, // 85: InventoryService.GetAllRepricing:output_type -> GetAllRepricingResponse
	49, // 86: InventoryService.GetAllRepricingItems:output_type -> GetAllRepricingItemsResponse
	32, // 87: InventoryService.UpsertMultiItemToRepricing:output_type -> ResponseID
	32, // 88: InventoryService.FinishRepricing:output_type -> ResponseID
	32, // 89: InventoryService.DeleteRepricing:output_type -> ResponseID
	32, // 90: InventoryService.RepricingBulkChangePrice:output_type -> ResponseID
	32, // 91: InventoryService.DeleteRepricingItem:output_type -> ResponseID
	46, // [46:92] is the sub-list for method output_type
	0,  // [0:46] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_inventory_main_proto_init() }
func file_inventory_main_proto_init() {
	if File_inventory_main_proto != nil {
		return
	}
	file_import_proto_init()
	file_supplier_proto_init()
	file_supplier_order_proto_init()
	file_transfer_proto_init()
	file_write_off_proto_init()
	file_repricing_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventory_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_main_proto_goTypes,
		DependencyIndexes: file_inventory_main_proto_depIdxs,
	}.Build()
	File_inventory_main_proto = out.File
	file_inventory_main_proto_rawDesc = nil
	file_inventory_main_proto_goTypes = nil
	file_inventory_main_proto_depIdxs = nil
}
