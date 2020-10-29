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
// source: google/ads/googleads/v1/services/ad_group_feed_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [AdGroupFeedService.GetAdGroupFeed][google.ads.googleads.v1.services.AdGroupFeedService.GetAdGroupFeed].
type GetAdGroupFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the ad group feed to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetAdGroupFeedRequest) Reset() {
	*x = GetAdGroupFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdGroupFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdGroupFeedRequest) ProtoMessage() {}

func (x *GetAdGroupFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdGroupFeedRequest.ProtoReflect.Descriptor instead.
func (*GetAdGroupFeedRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAdGroupFeedRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [AdGroupFeedService.MutateAdGroupFeeds][google.ads.googleads.v1.services.AdGroupFeedService.MutateAdGroupFeeds].
type MutateAdGroupFeedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose ad group feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ad group feeds.
	Operations []*AdGroupFeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateAdGroupFeedsRequest) Reset() {
	*x = MutateAdGroupFeedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupFeedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupFeedsRequest) ProtoMessage() {}

func (x *MutateAdGroupFeedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupFeedsRequest.ProtoReflect.Descriptor instead.
func (*MutateAdGroupFeedsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateAdGroupFeedsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAdGroupFeedsRequest) GetOperations() []*AdGroupFeedOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAdGroupFeedsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAdGroupFeedsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group feed.
type AdGroupFeedOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*AdGroupFeedOperation_Create
	//	*AdGroupFeedOperation_Update
	//	*AdGroupFeedOperation_Remove
	Operation isAdGroupFeedOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AdGroupFeedOperation) Reset() {
	*x = AdGroupFeedOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupFeedOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupFeedOperation) ProtoMessage() {}

func (x *AdGroupFeedOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupFeedOperation.ProtoReflect.Descriptor instead.
func (*AdGroupFeedOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescGZIP(), []int{2}
}

func (x *AdGroupFeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *AdGroupFeedOperation) GetOperation() isAdGroupFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AdGroupFeedOperation) GetCreate() *resources.AdGroupFeed {
	if x, ok := x.GetOperation().(*AdGroupFeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AdGroupFeedOperation) GetUpdate() *resources.AdGroupFeed {
	if x, ok := x.GetOperation().(*AdGroupFeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *AdGroupFeedOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AdGroupFeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAdGroupFeedOperation_Operation interface {
	isAdGroupFeedOperation_Operation()
}

type AdGroupFeedOperation_Create struct {
	// Create operation: No resource name is expected for the new ad group feed.
	Create *resources.AdGroupFeed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupFeedOperation_Update struct {
	// Update operation: The ad group feed is expected to have a valid resource
	// name.
	Update *resources.AdGroupFeed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupFeedOperation_Remove struct {
	// Remove operation: A resource name for the removed ad group feed is
	// expected, in this format:
	//
	// `customers/{customer_id}/adGroupFeeds/{ad_group_id}~{feed_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupFeedOperation_Create) isAdGroupFeedOperation_Operation() {}

func (*AdGroupFeedOperation_Update) isAdGroupFeedOperation_Operation() {}

func (*AdGroupFeedOperation_Remove) isAdGroupFeedOperation_Operation() {}

// Response message for an ad group feed mutate.
type MutateAdGroupFeedsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateAdGroupFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateAdGroupFeedsResponse) Reset() {
	*x = MutateAdGroupFeedsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupFeedsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupFeedsResponse) ProtoMessage() {}

func (x *MutateAdGroupFeedsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupFeedsResponse.ProtoReflect.Descriptor instead.
func (*MutateAdGroupFeedsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAdGroupFeedsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateAdGroupFeedsResponse) GetResults() []*MutateAdGroupFeedResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the ad group feed mutate.
type MutateAdGroupFeedResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateAdGroupFeedResult) Reset() {
	*x = MutateAdGroupFeedResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAdGroupFeedResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupFeedResult) ProtoMessage() {}

func (x *MutateAdGroupFeedResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupFeedResult.ProtoReflect.Descriptor instead.
func (*MutateAdGroupFeedResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateAdGroupFeedResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v1_services_ad_group_feed_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xec, 0x01, 0x0a, 0x19, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x5b, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0x8e, 0x02, 0x0a, 0x14, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x48, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x48, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xb9, 0x01, 0x0a, 0x1a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x53, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x3e, 0x0a,
	0x17, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xde, 0x03,
	0x0a, 0x12, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xc1, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64,
	0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x46, 0x65, 0x65, 0x64, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xe6, 0x01, 0x0a, 0x12, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12,
	0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x36, 0x22, 0x31, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x65, 0x65, 0x64, 0x73, 0x3a, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x1b, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0xfe,
	0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x17, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescData = file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDesc
)

func file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDescData
}

var file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v1_services_ad_group_feed_service_proto_goTypes = []interface{}{
	(*GetAdGroupFeedRequest)(nil),      // 0: google.ads.googleads.v1.services.GetAdGroupFeedRequest
	(*MutateAdGroupFeedsRequest)(nil),  // 1: google.ads.googleads.v1.services.MutateAdGroupFeedsRequest
	(*AdGroupFeedOperation)(nil),       // 2: google.ads.googleads.v1.services.AdGroupFeedOperation
	(*MutateAdGroupFeedsResponse)(nil), // 3: google.ads.googleads.v1.services.MutateAdGroupFeedsResponse
	(*MutateAdGroupFeedResult)(nil),    // 4: google.ads.googleads.v1.services.MutateAdGroupFeedResult
	(*field_mask.FieldMask)(nil),       // 5: google.protobuf.FieldMask
	(*resources.AdGroupFeed)(nil),      // 6: google.ads.googleads.v1.resources.AdGroupFeed
	(*status.Status)(nil),              // 7: google.rpc.Status
}
var file_google_ads_googleads_v1_services_ad_group_feed_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v1.services.MutateAdGroupFeedsRequest.operations:type_name -> google.ads.googleads.v1.services.AdGroupFeedOperation
	5, // 1: google.ads.googleads.v1.services.AdGroupFeedOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.ads.googleads.v1.services.AdGroupFeedOperation.create:type_name -> google.ads.googleads.v1.resources.AdGroupFeed
	6, // 3: google.ads.googleads.v1.services.AdGroupFeedOperation.update:type_name -> google.ads.googleads.v1.resources.AdGroupFeed
	7, // 4: google.ads.googleads.v1.services.MutateAdGroupFeedsResponse.partial_failure_error:type_name -> google.rpc.Status
	4, // 5: google.ads.googleads.v1.services.MutateAdGroupFeedsResponse.results:type_name -> google.ads.googleads.v1.services.MutateAdGroupFeedResult
	0, // 6: google.ads.googleads.v1.services.AdGroupFeedService.GetAdGroupFeed:input_type -> google.ads.googleads.v1.services.GetAdGroupFeedRequest
	1, // 7: google.ads.googleads.v1.services.AdGroupFeedService.MutateAdGroupFeeds:input_type -> google.ads.googleads.v1.services.MutateAdGroupFeedsRequest
	6, // 8: google.ads.googleads.v1.services.AdGroupFeedService.GetAdGroupFeed:output_type -> google.ads.googleads.v1.resources.AdGroupFeed
	3, // 9: google.ads.googleads.v1.services.AdGroupFeedService.MutateAdGroupFeeds:output_type -> google.ads.googleads.v1.services.MutateAdGroupFeedsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v1_services_ad_group_feed_service_proto_init() }
func file_google_ads_googleads_v1_services_ad_group_feed_service_proto_init() {
	if File_google_ads_googleads_v1_services_ad_group_feed_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdGroupFeedRequest); i {
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
		file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupFeedsRequest); i {
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
		file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupFeedOperation); i {
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
		file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupFeedsResponse); i {
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
		file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAdGroupFeedResult); i {
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
	file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AdGroupFeedOperation_Create)(nil),
		(*AdGroupFeedOperation_Update)(nil),
		(*AdGroupFeedOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v1_services_ad_group_feed_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v1_services_ad_group_feed_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v1_services_ad_group_feed_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v1_services_ad_group_feed_service_proto = out.File
	file_google_ads_googleads_v1_services_ad_group_feed_service_proto_rawDesc = nil
	file_google_ads_googleads_v1_services_ad_group_feed_service_proto_goTypes = nil
	file_google_ads_googleads_v1_services_ad_group_feed_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupFeedServiceClient is the client API for AdGroupFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupFeedServiceClient interface {
	// Returns the requested ad group feed in full detail.
	GetAdGroupFeed(ctx context.Context, in *GetAdGroupFeedRequest, opts ...grpc.CallOption) (*resources.AdGroupFeed, error)
	// Creates, updates, or removes ad group feeds. Operation statuses are
	// returned.
	MutateAdGroupFeeds(ctx context.Context, in *MutateAdGroupFeedsRequest, opts ...grpc.CallOption) (*MutateAdGroupFeedsResponse, error)
}

type adGroupFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupFeedServiceClient(cc grpc.ClientConnInterface) AdGroupFeedServiceClient {
	return &adGroupFeedServiceClient{cc}
}

func (c *adGroupFeedServiceClient) GetAdGroupFeed(ctx context.Context, in *GetAdGroupFeedRequest, opts ...grpc.CallOption) (*resources.AdGroupFeed, error) {
	out := new(resources.AdGroupFeed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupFeedService/GetAdGroupFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupFeedServiceClient) MutateAdGroupFeeds(ctx context.Context, in *MutateAdGroupFeedsRequest, opts ...grpc.CallOption) (*MutateAdGroupFeedsResponse, error) {
	out := new(MutateAdGroupFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupFeedService/MutateAdGroupFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupFeedServiceServer is the server API for AdGroupFeedService service.
type AdGroupFeedServiceServer interface {
	// Returns the requested ad group feed in full detail.
	GetAdGroupFeed(context.Context, *GetAdGroupFeedRequest) (*resources.AdGroupFeed, error)
	// Creates, updates, or removes ad group feeds. Operation statuses are
	// returned.
	MutateAdGroupFeeds(context.Context, *MutateAdGroupFeedsRequest) (*MutateAdGroupFeedsResponse, error)
}

// UnimplementedAdGroupFeedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupFeedServiceServer struct {
}

func (*UnimplementedAdGroupFeedServiceServer) GetAdGroupFeed(context.Context, *GetAdGroupFeedRequest) (*resources.AdGroupFeed, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupFeed not implemented")
}
func (*UnimplementedAdGroupFeedServiceServer) MutateAdGroupFeeds(context.Context, *MutateAdGroupFeedsRequest) (*MutateAdGroupFeedsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupFeeds not implemented")
}

func RegisterAdGroupFeedServiceServer(s *grpc.Server, srv AdGroupFeedServiceServer) {
	s.RegisterService(&_AdGroupFeedService_serviceDesc, srv)
}

func _AdGroupFeedService_GetAdGroupFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupFeedServiceServer).GetAdGroupFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupFeedService/GetAdGroupFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupFeedServiceServer).GetAdGroupFeed(ctx, req.(*GetAdGroupFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupFeedService_MutateAdGroupFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupFeedServiceServer).MutateAdGroupFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupFeedService/MutateAdGroupFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupFeedServiceServer).MutateAdGroupFeeds(ctx, req.(*MutateAdGroupFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupFeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdGroupFeedService",
	HandlerType: (*AdGroupFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupFeed",
			Handler:    _AdGroupFeedService_GetAdGroupFeed_Handler,
		},
		{
			MethodName: "MutateAdGroupFeeds",
			Handler:    _AdGroupFeedService_MutateAdGroupFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_group_feed_service.proto",
}
