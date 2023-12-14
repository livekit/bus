// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: testutils.proto

package testutils

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

type LaggyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	SentAt int64  `protobuf:"varint,2,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *LaggyMessage) Reset() {
	*x = LaggyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testutils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaggyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaggyMessage) ProtoMessage() {}

func (x *LaggyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testutils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaggyMessage.ProtoReflect.Descriptor instead.
func (*LaggyMessage) Descriptor() ([]byte, []int) {
	return file_testutils_proto_rawDescGZIP(), []int{0}
}

func (x *LaggyMessage) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *LaggyMessage) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

func (x *LaggyMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_testutils_proto protoreflect.FileDescriptor

var file_testutils_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x65, 0x73, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x22, 0x53, 0x0a, 0x0c,
	0x4c, 0x61, 0x67, 0x67, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x73, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testutils_proto_rawDescOnce sync.Once
	file_testutils_proto_rawDescData = file_testutils_proto_rawDesc
)

func file_testutils_proto_rawDescGZIP() []byte {
	file_testutils_proto_rawDescOnce.Do(func() {
		file_testutils_proto_rawDescData = protoimpl.X.CompressGZIP(file_testutils_proto_rawDescData)
	})
	return file_testutils_proto_rawDescData
}

var file_testutils_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_testutils_proto_goTypes = []interface{}{
	(*LaggyMessage)(nil), // 0: testutils.LaggyMessage
}
var file_testutils_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testutils_proto_init() }
func file_testutils_proto_init() {
	if File_testutils_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testutils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaggyMessage); i {
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
			RawDescriptor: file_testutils_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testutils_proto_goTypes,
		DependencyIndexes: file_testutils_proto_depIdxs,
		MessageInfos:      file_testutils_proto_msgTypes,
	}.Build()
	File_testutils_proto = out.File
	file_testutils_proto_rawDesc = nil
	file_testutils_proto_goTypes = nil
	file_testutils_proto_depIdxs = nil
}
