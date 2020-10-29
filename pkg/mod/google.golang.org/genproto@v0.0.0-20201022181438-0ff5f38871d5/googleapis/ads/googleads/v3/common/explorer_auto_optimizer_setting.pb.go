// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/ads/googleads/v3/common/explorer_auto_optimizer_setting.proto

package common

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Settings for the Display Campaign Optimizer, initially named "Explorer".
// Learn more about
// [automatic targeting](https://support.google.com/google-ads/answer/190596).
type ExplorerAutoOptimizerSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether the optimizer is turned on.
	OptIn *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
}

func (x *ExplorerAutoOptimizerSetting) Reset() {
	*x = ExplorerAutoOptimizerSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerAutoOptimizerSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerAutoOptimizerSetting) ProtoMessage() {}

func (x *ExplorerAutoOptimizerSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerAutoOptimizerSetting.ProtoReflect.Descriptor instead.
func (*ExplorerAutoOptimizerSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescGZIP(), []int{0}
}

func (x *ExplorerAutoOptimizerSetting) GetOptIn() *wrapperspb.BoolValue {
	if x != nil {
		return x.OptIn
	}
	return nil
}

var File_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x1c, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x6f, 0x70, 0x74, 0x49, 0x6e, 0x42, 0xfc, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x21,
	0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x7a, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x33, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x33, 0x3a, 0x3a,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescData = file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDesc
)

func file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDescData
}

var file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_goTypes = []interface{}{
	(*ExplorerAutoOptimizerSetting)(nil), // 0: google.ads.googleads.v3.common.ExplorerAutoOptimizerSetting
	(*wrapperspb.BoolValue)(nil),         // 1: google.protobuf.BoolValue
}
var file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v3.common.ExplorerAutoOptimizerSetting.opt_in:type_name -> google.protobuf.BoolValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_init() }
func file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_init() {
	if File_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplorerAutoOptimizerSetting); i {
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
			RawDescriptor: file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto = out.File
	file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_rawDesc = nil
	file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_goTypes = nil
	file_google_ads_googleads_v3_common_explorer_auto_optimizer_setting_proto_depIdxs = nil
}
