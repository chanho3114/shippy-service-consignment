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
// 	protoc        v3.12.3
// source: google/ads/googleads/v1/enums/education_placeholder_field.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Possible values for Education placeholder fields.
type EducationPlaceholderFieldEnum_EducationPlaceholderField int32

const (
	// Not specified.
	EducationPlaceholderFieldEnum_UNSPECIFIED EducationPlaceholderFieldEnum_EducationPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	EducationPlaceholderFieldEnum_UNKNOWN EducationPlaceholderFieldEnum_EducationPlaceholderField = 1
	// Data Type: STRING. Required. Combination of PROGRAM ID and LOCATION ID
	// must be unique per offer.
	EducationPlaceholderFieldEnum_PROGRAM_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 2
	// Data Type: STRING. Combination of PROGRAM ID and LOCATION ID must be
	// unique per offer.
	EducationPlaceholderFieldEnum_LOCATION_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with program name to be shown
	// in dynamic ad.
	EducationPlaceholderFieldEnum_PROGRAM_NAME EducationPlaceholderFieldEnum_EducationPlaceholderField = 4
	// Data Type: STRING. Area of study that can be shown in dynamic ad.
	EducationPlaceholderFieldEnum_AREA_OF_STUDY EducationPlaceholderFieldEnum_EducationPlaceholderField = 5
	// Data Type: STRING. Description of program that can be shown in dynamic
	// ad.
	EducationPlaceholderFieldEnum_PROGRAM_DESCRIPTION EducationPlaceholderFieldEnum_EducationPlaceholderField = 6
	// Data Type: STRING. Name of school that can be shown in dynamic ad.
	EducationPlaceholderFieldEnum_SCHOOL_NAME EducationPlaceholderFieldEnum_EducationPlaceholderField = 7
	// Data Type: STRING. Complete school address, including postal code.
	EducationPlaceholderFieldEnum_ADDRESS EducationPlaceholderFieldEnum_EducationPlaceholderField = 8
	// Data Type: URL. Image to be displayed in ads.
	EducationPlaceholderFieldEnum_THUMBNAIL_IMAGE_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 9
	// Data Type: URL. Alternative hosted file of image to be used in the ad.
	EducationPlaceholderFieldEnum_ALTERNATIVE_THUMBNAIL_IMAGE_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 10
	// Data Type: URL_LIST. Required. Final URLs to be used in ad when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific program and its location).
	EducationPlaceholderFieldEnum_FINAL_URLS EducationPlaceholderFieldEnum_EducationPlaceholderField = 11
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	EducationPlaceholderFieldEnum_FINAL_MOBILE_URLS EducationPlaceholderFieldEnum_EducationPlaceholderField = 12
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	EducationPlaceholderFieldEnum_TRACKING_URL EducationPlaceholderFieldEnum_EducationPlaceholderField = 13
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	EducationPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS EducationPlaceholderFieldEnum_EducationPlaceholderField = 14
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	EducationPlaceholderFieldEnum_ANDROID_APP_LINK EducationPlaceholderFieldEnum_EducationPlaceholderField = 15
	// Data Type: STRING_LIST. List of recommended program IDs to show together
	// with this item.
	EducationPlaceholderFieldEnum_SIMILAR_PROGRAM_IDS EducationPlaceholderFieldEnum_EducationPlaceholderField = 16
	// Data Type: STRING. iOS app link.
	EducationPlaceholderFieldEnum_IOS_APP_LINK EducationPlaceholderFieldEnum_EducationPlaceholderField = 17
	// Data Type: INT64. iOS app store ID.
	EducationPlaceholderFieldEnum_IOS_APP_STORE_ID EducationPlaceholderFieldEnum_EducationPlaceholderField = 18
)

// Enum value maps for EducationPlaceholderFieldEnum_EducationPlaceholderField.
var (
	EducationPlaceholderFieldEnum_EducationPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "PROGRAM_ID",
		3:  "LOCATION_ID",
		4:  "PROGRAM_NAME",
		5:  "AREA_OF_STUDY",
		6:  "PROGRAM_DESCRIPTION",
		7:  "SCHOOL_NAME",
		8:  "ADDRESS",
		9:  "THUMBNAIL_IMAGE_URL",
		10: "ALTERNATIVE_THUMBNAIL_IMAGE_URL",
		11: "FINAL_URLS",
		12: "FINAL_MOBILE_URLS",
		13: "TRACKING_URL",
		14: "CONTEXTUAL_KEYWORDS",
		15: "ANDROID_APP_LINK",
		16: "SIMILAR_PROGRAM_IDS",
		17: "IOS_APP_LINK",
		18: "IOS_APP_STORE_ID",
	}
	EducationPlaceholderFieldEnum_EducationPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"UNKNOWN":                         1,
		"PROGRAM_ID":                      2,
		"LOCATION_ID":                     3,
		"PROGRAM_NAME":                    4,
		"AREA_OF_STUDY":                   5,
		"PROGRAM_DESCRIPTION":             6,
		"SCHOOL_NAME":                     7,
		"ADDRESS":                         8,
		"THUMBNAIL_IMAGE_URL":             9,
		"ALTERNATIVE_THUMBNAIL_IMAGE_URL": 10,
		"FINAL_URLS":                      11,
		"FINAL_MOBILE_URLS":               12,
		"TRACKING_URL":                    13,
		"CONTEXTUAL_KEYWORDS":             14,
		"ANDROID_APP_LINK":                15,
		"SIMILAR_PROGRAM_IDS":             16,
		"IOS_APP_LINK":                    17,
		"IOS_APP_STORE_ID":                18,
	}
)

func (x EducationPlaceholderFieldEnum_EducationPlaceholderField) Enum() *EducationPlaceholderFieldEnum_EducationPlaceholderField {
	p := new(EducationPlaceholderFieldEnum_EducationPlaceholderField)
	*p = x
	return p
}

func (x EducationPlaceholderFieldEnum_EducationPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EducationPlaceholderFieldEnum_EducationPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v1_enums_education_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (EducationPlaceholderFieldEnum_EducationPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v1_enums_education_placeholder_field_proto_enumTypes[0]
}

func (x EducationPlaceholderFieldEnum_EducationPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EducationPlaceholderFieldEnum_EducationPlaceholderField.Descriptor instead.
func (EducationPlaceholderFieldEnum_EducationPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Education placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type EducationPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EducationPlaceholderFieldEnum) Reset() {
	*x = EducationPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_enums_education_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EducationPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EducationPlaceholderFieldEnum) ProtoMessage() {}

func (x *EducationPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_enums_education_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EducationPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*EducationPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v1_enums_education_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf,
	0x03, 0x0a, 0x1d, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0x9d, 0x03, 0x0a, 0x19, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x04, 0x12,
	0x11, 0x0a, 0x0d, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x54, 0x55, 0x44, 0x59,
	0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x48, 0x55,
	0x4d, 0x42, 0x4e, 0x41, 0x49, 0x4c, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c,
	0x10, 0x09, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x54, 0x48, 0x55, 0x4d, 0x42, 0x4e, 0x41, 0x49, 0x4c, 0x5f, 0x49, 0x4d, 0x41, 0x47,
	0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c,
	0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4e, 0x41, 0x4c,
	0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0c, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0d,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41, 0x4c, 0x5f, 0x4b,
	0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x0f, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x41, 0x4d, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f,
	0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x12,
	0x42, 0xf3, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1e, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescData = file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v1_enums_education_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v1_enums_education_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v1_enums_education_placeholder_field_proto_goTypes = []interface{}{
	(EducationPlaceholderFieldEnum_EducationPlaceholderField)(0), // 0: google.ads.googleads.v1.enums.EducationPlaceholderFieldEnum.EducationPlaceholderField
	(*EducationPlaceholderFieldEnum)(nil),                        // 1: google.ads.googleads.v1.enums.EducationPlaceholderFieldEnum
}
var file_google_ads_googleads_v1_enums_education_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v1_enums_education_placeholder_field_proto_init() }
func file_google_ads_googleads_v1_enums_education_placeholder_field_proto_init() {
	if File_google_ads_googleads_v1_enums_education_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v1_enums_education_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EducationPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v1_enums_education_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v1_enums_education_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v1_enums_education_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v1_enums_education_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v1_enums_education_placeholder_field_proto = out.File
	file_google_ads_googleads_v1_enums_education_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v1_enums_education_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v1_enums_education_placeholder_field_proto_depIdxs = nil
}