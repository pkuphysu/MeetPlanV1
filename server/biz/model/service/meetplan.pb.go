// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: meetplan.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	model "meetplan/biz/model"
	_ "meetplan/biz/model/common"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_meetplan_proto protoreflect.FileDescriptor

var file_meetplan_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9b, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x4d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0xaa, 0xc8, 0x18, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x57, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6c,
	0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xca, 0xc1, 0x18, 0x11, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x6c, 0x66,
	0xaa, 0xc8, 0x18, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1c, 0xca, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x3a, 0x69, 0x64, 0xaa, 0xc8, 0x18, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x55, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xca,
	0xc1, 0x18, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xaa,
	0xc8, 0x18, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x5b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xd2, 0xc1, 0x18, 0x0c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xaa, 0xc8, 0x18, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xda, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x3a, 0x69, 0x64, 0xaa, 0xc8, 0x18,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xcc, 0x09, 0x0a, 0x08, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0xca, 0xc1, 0x18, 0x14, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x3a,
	0x69, 0x64, 0xaa, 0xc8, 0x18, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x69,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1a,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0xca, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0xaa, 0xc8, 0x18,
	0x08, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x6f, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0xaa, 0xc8,
	0x18, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x73, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0xda, 0xc1, 0x18, 0x14, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f,
	0x3a, 0x69, 0x64, 0xaa, 0xc8, 0x18, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12,
	0x73, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24,
	0xe2, 0xc1, 0x18, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x3a, 0x69, 0x64, 0xaa, 0xc8, 0x18, 0x08, 0x6d, 0x65, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x12, 0x72, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0xe2, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0xaa, 0xc8, 0x18, 0x08,
	0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x5e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0xca, 0xc1, 0x18, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x3a, 0x69, 0x64, 0xaa, 0xc8, 0x18, 0x08,
	0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x5d, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0xaa, 0xc8, 0x18, 0x08, 0x6d,
	0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x63, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xd2,
	0xc1, 0x18, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0xaa, 0xc8, 0x18, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x67, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0xda, 0xc1, 0x18, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x3a, 0x69, 0x64, 0xaa, 0xc8, 0x18, 0x08, 0x6d, 0x65, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x8c, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x24, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x6e, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0xd2,
	0xc1, 0x18, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x6e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0xaa, 0xc8, 0x18, 0x08, 0x6d, 0x65, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x32, 0xdb, 0x05, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x6f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0xca, 0xc1, 0x18, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x6b, 0xaa, 0xc8, 0x18, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x75, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0xd2, 0xc1, 0x18, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x6b, 0xaa, 0xc8, 0x18,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0xca, 0xc1,
	0x18, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0xaa, 0xc8, 0x18, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x7d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0xd2, 0xc1, 0x18,
	0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0xaa, 0xc8, 0x18, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x73, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0xca, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x65, 0xaa, 0xc8, 0x18, 0x06, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x72, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72,
	0x6d, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1e, 0xda, 0xc1, 0x18, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x65, 0xaa, 0xc8, 0x18, 0x06, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x65, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_meetplan_proto_goTypes = []interface{}{
	(*model.LoginRequest)(nil),                   // 0: model.LoginRequest
	(*model.GetSelfRequest)(nil),                 // 1: model.GetSelfRequest
	(*model.GetUserRequest)(nil),                 // 2: model.GetUserRequest
	(*model.ListUserRequest)(nil),                // 3: model.ListUserRequest
	(*model.CreateUserRequest)(nil),              // 4: model.CreateUserRequest
	(*model.UpdateUserRequest)(nil),              // 5: model.UpdateUserRequest
	(*model.GetMeetPlanRequest)(nil),             // 6: model.GetMeetPlanRequest
	(*model.ListMeetPlanRequest)(nil),            // 7: model.ListMeetPlanRequest
	(*model.CreateMeetPlanRequest)(nil),          // 8: model.CreateMeetPlanRequest
	(*model.UpdateMeetPlanRequest)(nil),          // 9: model.UpdateMeetPlanRequest
	(*model.DeleteMeetPlanRequest)(nil),          // 10: model.DeleteMeetPlanRequest
	(*model.DeleteMeetPlansRequest)(nil),         // 11: model.DeleteMeetPlansRequest
	(*model.GetOrderRequest)(nil),                // 12: model.GetOrderRequest
	(*model.ListOrderRequest)(nil),               // 13: model.ListOrderRequest
	(*model.CreateOrderRequest)(nil),             // 14: model.CreateOrderRequest
	(*model.UpdateOrderRequest)(nil),             // 15: model.UpdateOrderRequest
	(*model.CreateMeetPlanAndOrderRequest)(nil),  // 16: model.CreateMeetPlanAndOrderRequest
	(*model.ListFriendLinkRequest)(nil),          // 17: model.ListFriendLinkRequest
	(*model.CreateFriendLinkRequest)(nil),        // 18: model.CreateFriendLinkRequest
	(*model.ListUpdateRecordRequest)(nil),        // 19: model.ListUpdateRecordRequest
	(*model.CreateUpdateRecordRequest)(nil),      // 20: model.CreateUpdateRecordRequest
	(*model.GetTermDateRangeRequest)(nil),        // 21: model.GetTermDateRangeRequest
	(*model.UpdateTermDateRangeRequest)(nil),     // 22: model.UpdateTermDateRangeRequest
	(*model.LoginResponse)(nil),                  // 23: model.LoginResponse
	(*model.GetSelfResponse)(nil),                // 24: model.GetSelfResponse
	(*model.GetUserResponse)(nil),                // 25: model.GetUserResponse
	(*model.ListUserResponse)(nil),               // 26: model.ListUserResponse
	(*model.CreateUserResponse)(nil),             // 27: model.CreateUserResponse
	(*model.UpdateUserResponse)(nil),             // 28: model.UpdateUserResponse
	(*model.GetMeetPlanResponse)(nil),            // 29: model.GetMeetPlanResponse
	(*model.ListMeetPlanResponse)(nil),           // 30: model.ListMeetPlanResponse
	(*model.CreateMeetPlanResponse)(nil),         // 31: model.CreateMeetPlanResponse
	(*model.UpdateMeetPlanResponse)(nil),         // 32: model.UpdateMeetPlanResponse
	(*model.DeleteMeetPlanResponse)(nil),         // 33: model.DeleteMeetPlanResponse
	(*model.DeleteMeetPlansResponse)(nil),        // 34: model.DeleteMeetPlansResponse
	(*model.GetOrderResponse)(nil),               // 35: model.GetOrderResponse
	(*model.ListOrderResponse)(nil),              // 36: model.ListOrderResponse
	(*model.CreateOrderResponse)(nil),            // 37: model.CreateOrderResponse
	(*model.UpdateOrderResponse)(nil),            // 38: model.UpdateOrderResponse
	(*model.CreateMeetPlanAndOrderResponse)(nil), // 39: model.CreateMeetPlanAndOrderResponse
	(*model.ListFriendLinkResponse)(nil),         // 40: model.ListFriendLinkResponse
	(*model.CreateFriendLinkResponse)(nil),       // 41: model.CreateFriendLinkResponse
	(*model.ListUpdateRecordResponse)(nil),       // 42: model.ListUpdateRecordResponse
	(*model.CreateUpdateRecordResponse)(nil),     // 43: model.CreateUpdateRecordResponse
	(*model.GetTermDateRangeResponse)(nil),       // 44: model.GetTermDateRangeResponse
	(*model.UpdateTermDateRangeResponse)(nil),    // 45: model.UpdateTermDateRangeResponse
}
var file_meetplan_proto_depIdxs = []int32{
	0,  // 0: service.User.Login:input_type -> model.LoginRequest
	1,  // 1: service.User.GetSelf:input_type -> model.GetSelfRequest
	2,  // 2: service.User.GetUser:input_type -> model.GetUserRequest
	3,  // 3: service.User.ListUser:input_type -> model.ListUserRequest
	4,  // 4: service.User.CreateUser:input_type -> model.CreateUserRequest
	5,  // 5: service.User.UpdateUser:input_type -> model.UpdateUserRequest
	6,  // 6: service.MeetPlan.GetMeetPlan:input_type -> model.GetMeetPlanRequest
	7,  // 7: service.MeetPlan.ListMeetPlan:input_type -> model.ListMeetPlanRequest
	8,  // 8: service.MeetPlan.CreateMeetPlan:input_type -> model.CreateMeetPlanRequest
	9,  // 9: service.MeetPlan.UpdateMeetPlan:input_type -> model.UpdateMeetPlanRequest
	10, // 10: service.MeetPlan.DeleteMeetPlan:input_type -> model.DeleteMeetPlanRequest
	11, // 11: service.MeetPlan.DeleteMeetPlans:input_type -> model.DeleteMeetPlansRequest
	12, // 12: service.MeetPlan.GetOrder:input_type -> model.GetOrderRequest
	13, // 13: service.MeetPlan.ListOrder:input_type -> model.ListOrderRequest
	14, // 14: service.MeetPlan.CreateOrder:input_type -> model.CreateOrderRequest
	15, // 15: service.MeetPlan.UpdateOrder:input_type -> model.UpdateOrderRequest
	16, // 16: service.MeetPlan.CreateMeetPlanAndOrder:input_type -> model.CreateMeetPlanAndOrderRequest
	17, // 17: service.Option.ListFriendLink:input_type -> model.ListFriendLinkRequest
	18, // 18: service.Option.CreateFriendLink:input_type -> model.CreateFriendLinkRequest
	19, // 19: service.Option.ListUpdateRecord:input_type -> model.ListUpdateRecordRequest
	20, // 20: service.Option.CreateUpdateRecord:input_type -> model.CreateUpdateRecordRequest
	21, // 21: service.Option.GetTermDateRange:input_type -> model.GetTermDateRangeRequest
	22, // 22: service.Option.UpdateTermDateRange:input_type -> model.UpdateTermDateRangeRequest
	23, // 23: service.User.Login:output_type -> model.LoginResponse
	24, // 24: service.User.GetSelf:output_type -> model.GetSelfResponse
	25, // 25: service.User.GetUser:output_type -> model.GetUserResponse
	26, // 26: service.User.ListUser:output_type -> model.ListUserResponse
	27, // 27: service.User.CreateUser:output_type -> model.CreateUserResponse
	28, // 28: service.User.UpdateUser:output_type -> model.UpdateUserResponse
	29, // 29: service.MeetPlan.GetMeetPlan:output_type -> model.GetMeetPlanResponse
	30, // 30: service.MeetPlan.ListMeetPlan:output_type -> model.ListMeetPlanResponse
	31, // 31: service.MeetPlan.CreateMeetPlan:output_type -> model.CreateMeetPlanResponse
	32, // 32: service.MeetPlan.UpdateMeetPlan:output_type -> model.UpdateMeetPlanResponse
	33, // 33: service.MeetPlan.DeleteMeetPlan:output_type -> model.DeleteMeetPlanResponse
	34, // 34: service.MeetPlan.DeleteMeetPlans:output_type -> model.DeleteMeetPlansResponse
	35, // 35: service.MeetPlan.GetOrder:output_type -> model.GetOrderResponse
	36, // 36: service.MeetPlan.ListOrder:output_type -> model.ListOrderResponse
	37, // 37: service.MeetPlan.CreateOrder:output_type -> model.CreateOrderResponse
	38, // 38: service.MeetPlan.UpdateOrder:output_type -> model.UpdateOrderResponse
	39, // 39: service.MeetPlan.CreateMeetPlanAndOrder:output_type -> model.CreateMeetPlanAndOrderResponse
	40, // 40: service.Option.ListFriendLink:output_type -> model.ListFriendLinkResponse
	41, // 41: service.Option.CreateFriendLink:output_type -> model.CreateFriendLinkResponse
	42, // 42: service.Option.ListUpdateRecord:output_type -> model.ListUpdateRecordResponse
	43, // 43: service.Option.CreateUpdateRecord:output_type -> model.CreateUpdateRecordResponse
	44, // 44: service.Option.GetTermDateRange:output_type -> model.GetTermDateRangeResponse
	45, // 45: service.Option.UpdateTermDateRange:output_type -> model.UpdateTermDateRangeResponse
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_meetplan_proto_init() }
func file_meetplan_proto_init() {
	if File_meetplan_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meetplan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_meetplan_proto_goTypes,
		DependencyIndexes: file_meetplan_proto_depIdxs,
	}.Build()
	File_meetplan_proto = out.File
	file_meetplan_proto_rawDesc = nil
	file_meetplan_proto_goTypes = nil
	file_meetplan_proto_depIdxs = nil
}
