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
// source: google/cloud/dialogflow/v2beta1/environment.proto

package dialogflow

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Represents an environment state. When an environment is pointed to a new
// agent version, the environment is temporarily set to the `LOADING` state.
// During that time, the environment keeps on serving the previous version of
// the agent. After the new agent version is done loading, the environment is
// set back to the `RUNNING` state.
type Environment_State int32

const (
	// Not specified. This value is not used.
	Environment_STATE_UNSPECIFIED Environment_State = 0
	// Stopped.
	Environment_STOPPED Environment_State = 1
	// Loading.
	Environment_LOADING Environment_State = 2
	// Running.
	Environment_RUNNING Environment_State = 3
)

// Enum value maps for Environment_State.
var (
	Environment_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STOPPED",
		2: "LOADING",
		3: "RUNNING",
	}
	Environment_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STOPPED":           1,
		"LOADING":           2,
		"RUNNING":           3,
	}
)

func (x Environment_State) Enum() *Environment_State {
	p := new(Environment_State)
	*p = x
	return p
}

func (x Environment_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Environment_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_v2beta1_environment_proto_enumTypes[0].Descriptor()
}

func (Environment_State) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_v2beta1_environment_proto_enumTypes[0]
}

func (x Environment_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Environment_State.Descriptor instead.
func (Environment_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescGZIP(), []int{0, 0}
}

// You can create multiple versions of your agent and publish them to separate
// environments.
//
// When you edit an agent, you are editing the draft agent. At any point, you
// can save the draft agent as an agent version, which is an immutable snapshot
// of your agent.
//
// When you save the draft agent, it is published to the default environment.
// When you create agent versions, you can publish them to custom environments.
// You can create a variety of custom environments for:
//
// - testing
// - development
// - production
// - etc.
//
// For more information, see the [versions and environments
// guide](https://cloud.google.com/dialogflow/docs/agents-versions).
type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The unique identifier of this agent environment.
	// Supported formats:
	// - `projects/<Project Number / ID>/agent/environments/<Environment ID>`
	// - `projects/<Project Number / ID>/locations/<Location
	//   ID>/agent/environments/<Environment ID>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The developer-provided description for this environment.
	// The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. The agent version loaded into this environment.
	// Supported formats:
	// - `projects/<Project ID>/agent/versions/<Version ID>`
	// - `projects/<Project ID>/locations/<Location ID>/agent/versions/<Version
	//   ID>`
	AgentVersion string `protobuf:"bytes,3,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
	// Output only. The state of this environment. This field is read-only, i.e., it cannot be
	// set by create and update methods.
	State Environment_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.dialogflow.v2beta1.Environment_State" json:"state,omitempty"`
	// Output only. The last update time of this environment. This field is read-only, i.e., it
	// cannot be set by create and update methods.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescGZIP(), []int{0}
}

func (x *Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Environment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Environment) GetAgentVersion() string {
	if x != nil {
		return x.AgentVersion
	}
	return ""
}

func (x *Environment) GetState() Environment_State {
	if x != nil {
		return x.State
	}
	return Environment_STATE_UNSPECIFIED
}

func (x *Environment) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// The request message for [Environments.ListEnvironments][google.cloud.dialogflow.v2beta1.Environments.ListEnvironments].
type ListEnvironmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The agent to list all environments from.
	// Format:
	// - `projects/<Project Number / ID>/agent`
	// - `projects/<Project Number / ID>/locations/<Location ID>/agent`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By default 100 and
	// at most 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListEnvironmentsRequest) Reset() {
	*x = ListEnvironmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnvironmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvironmentsRequest) ProtoMessage() {}

func (x *ListEnvironmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvironmentsRequest.ProtoReflect.Descriptor instead.
func (*ListEnvironmentsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescGZIP(), []int{1}
}

func (x *ListEnvironmentsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListEnvironmentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListEnvironmentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response message for [Environments.ListEnvironments][google.cloud.dialogflow.v2beta1.Environments.ListEnvironments].
type ListEnvironmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of agent environments. There will be a maximum number of items
	// returned based on the page_size field in the request.
	Environments []*Environment `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListEnvironmentsResponse) Reset() {
	*x = ListEnvironmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnvironmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvironmentsResponse) ProtoMessage() {}

func (x *ListEnvironmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvironmentsResponse.ProtoReflect.Descriptor instead.
func (*ListEnvironmentsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescGZIP(), []int{2}
}

func (x *ListEnvironmentsResponse) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

func (x *ListEnvironmentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_dialogflow_v2beta1_environment_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_v2beta1_environment_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x03, 0x0a, 0x0b,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0d, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x3a, 0xaa, 0x01,
	0xea, 0x41, 0xa6, 0x01, 0x0a, 0x25, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x7b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d,
	0x12, 0x48, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x22, 0xa6, 0x01, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27, 0x12, 0x25,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x93, 0x03, 0x0a, 0x0c, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x88, 0x02, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x70, 0x12, 0x2f, 0x2f,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x7d, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5a, 0x3d,
	0x12, 0x3b, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x7d,
	0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0xda, 0x41, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x1a, 0x78, 0xca, 0x41, 0x19, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x59, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x42, 0xae, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescData = file_google_cloud_dialogflow_v2beta1_environment_proto_rawDesc
)

func file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_v2beta1_environment_proto_rawDescData
}

var file_google_cloud_dialogflow_v2beta1_environment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_dialogflow_v2beta1_environment_proto_goTypes = []interface{}{
	(Environment_State)(0),           // 0: google.cloud.dialogflow.v2beta1.Environment.State
	(*Environment)(nil),              // 1: google.cloud.dialogflow.v2beta1.Environment
	(*ListEnvironmentsRequest)(nil),  // 2: google.cloud.dialogflow.v2beta1.ListEnvironmentsRequest
	(*ListEnvironmentsResponse)(nil), // 3: google.cloud.dialogflow.v2beta1.ListEnvironmentsResponse
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_google_cloud_dialogflow_v2beta1_environment_proto_depIdxs = []int32{
	0, // 0: google.cloud.dialogflow.v2beta1.Environment.state:type_name -> google.cloud.dialogflow.v2beta1.Environment.State
	4, // 1: google.cloud.dialogflow.v2beta1.Environment.update_time:type_name -> google.protobuf.Timestamp
	1, // 2: google.cloud.dialogflow.v2beta1.ListEnvironmentsResponse.environments:type_name -> google.cloud.dialogflow.v2beta1.Environment
	2, // 3: google.cloud.dialogflow.v2beta1.Environments.ListEnvironments:input_type -> google.cloud.dialogflow.v2beta1.ListEnvironmentsRequest
	3, // 4: google.cloud.dialogflow.v2beta1.Environments.ListEnvironments:output_type -> google.cloud.dialogflow.v2beta1.ListEnvironmentsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_v2beta1_environment_proto_init() }
func file_google_cloud_dialogflow_v2beta1_environment_proto_init() {
	if File_google_cloud_dialogflow_v2beta1_environment_proto != nil {
		return
	}
	file_google_cloud_dialogflow_v2beta1_audio_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnvironmentsRequest); i {
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
		file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnvironmentsResponse); i {
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
			RawDescriptor: file_google_cloud_dialogflow_v2beta1_environment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_dialogflow_v2beta1_environment_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_v2beta1_environment_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_v2beta1_environment_proto_enumTypes,
		MessageInfos:      file_google_cloud_dialogflow_v2beta1_environment_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_v2beta1_environment_proto = out.File
	file_google_cloud_dialogflow_v2beta1_environment_proto_rawDesc = nil
	file_google_cloud_dialogflow_v2beta1_environment_proto_goTypes = nil
	file_google_cloud_dialogflow_v2beta1_environment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EnvironmentsClient is the client API for Environments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnvironmentsClient interface {
	// Returns the list of all non-draft environments of the specified agent.
	ListEnvironments(ctx context.Context, in *ListEnvironmentsRequest, opts ...grpc.CallOption) (*ListEnvironmentsResponse, error)
}

type environmentsClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentsClient(cc grpc.ClientConnInterface) EnvironmentsClient {
	return &environmentsClient{cc}
}

func (c *environmentsClient) ListEnvironments(ctx context.Context, in *ListEnvironmentsRequest, opts ...grpc.CallOption) (*ListEnvironmentsResponse, error) {
	out := new(ListEnvironmentsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Environments/ListEnvironments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentsServer is the server API for Environments service.
type EnvironmentsServer interface {
	// Returns the list of all non-draft environments of the specified agent.
	ListEnvironments(context.Context, *ListEnvironmentsRequest) (*ListEnvironmentsResponse, error)
}

// UnimplementedEnvironmentsServer can be embedded to have forward compatible implementations.
type UnimplementedEnvironmentsServer struct {
}

func (*UnimplementedEnvironmentsServer) ListEnvironments(context.Context, *ListEnvironmentsRequest) (*ListEnvironmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnvironments not implemented")
}

func RegisterEnvironmentsServer(s *grpc.Server, srv EnvironmentsServer) {
	s.RegisterService(&_Environments_serviceDesc, srv)
}

func _Environments_ListEnvironments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnvironmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).ListEnvironments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Environments/ListEnvironments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).ListEnvironments(ctx, req.(*ListEnvironmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Environments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2beta1.Environments",
	HandlerType: (*EnvironmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEnvironments",
			Handler:    _Environments_ListEnvironments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2beta1/environment.proto",
}
