// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/monitoring/v3/snooze_service.proto

package monitoringpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The message definition for creating a `Snooze`. Users must provide the body
// of the `Snooze` to be created but must omit the `Snooze` field, `name`.
type CreateSnoozeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The
	// [project](https://cloud.google.com/monitoring/api/v3#project_name) in which
	// a `Snooze` should be created. The format is:
	//
	//	projects/[PROJECT_ID_OR_NUMBER]
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The `Snooze` to create. Omit the `name` field, as it will be
	// filled in by the API.
	Snooze *Snooze `protobuf:"bytes,2,opt,name=snooze,proto3" json:"snooze,omitempty"`
}

func (x *CreateSnoozeRequest) Reset() {
	*x = CreateSnoozeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnoozeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnoozeRequest) ProtoMessage() {}

func (x *CreateSnoozeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnoozeRequest.ProtoReflect.Descriptor instead.
func (*CreateSnoozeRequest) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_snooze_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSnoozeRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateSnoozeRequest) GetSnooze() *Snooze {
	if x != nil {
		return x.Snooze
	}
	return nil
}

// The message definition for listing `Snooze`s associated with the given
// `parent`, satisfying the optional `filter`.
type ListSnoozesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The
	// [project](https://cloud.google.com/monitoring/api/v3#project_name) whose
	// `Snooze`s should be listed. The format is:
	//
	//	projects/[PROJECT_ID_OR_NUMBER]
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Optional filter to restrict results to the given criteria. The
	// following fields are supported.
	//
	//   - `interval.start_time`
	//   - `interval.end_time`
	//
	// For example:
	//
	//	```
	//	interval.start_time > "2022-03-11T00:00:00-08:00" AND
	//	    interval.end_time < "2022-03-12T00:00:00-08:00"
	//	```
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. The maximum number of results to return for a single query. The
	// server may further constrain the maximum number of results returned in a
	// single page. The value should be in the range [1, 1000]. If the value given
	// is outside this range, the server will decide the number of results to be
	// returned.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The `next_page_token` from a previous call to
	// `ListSnoozesRequest` to get the next page of results.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListSnoozesRequest) Reset() {
	*x = ListSnoozesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnoozesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnoozesRequest) ProtoMessage() {}

func (x *ListSnoozesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnoozesRequest.ProtoReflect.Descriptor instead.
func (*ListSnoozesRequest) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_snooze_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSnoozesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListSnoozesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListSnoozesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSnoozesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The results of a successful `ListSnoozes` call, containing the matching
// `Snooze`s.
type ListSnoozesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `Snooze`s matching this list call.
	Snoozes []*Snooze `protobuf:"bytes,1,rep,name=snoozes,proto3" json:"snoozes,omitempty"`
	// Page token for repeated calls to `ListSnoozes`, to fetch additional pages
	// of results. If this is empty or missing, there are no more pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSnoozesResponse) Reset() {
	*x = ListSnoozesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnoozesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnoozesResponse) ProtoMessage() {}

func (x *ListSnoozesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnoozesResponse.ProtoReflect.Descriptor instead.
func (*ListSnoozesResponse) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_snooze_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSnoozesResponse) GetSnoozes() []*Snooze {
	if x != nil {
		return x.Snoozes
	}
	return nil
}

func (x *ListSnoozesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The message definition for retrieving a `Snooze`. Users must specify the
// field, `name`, which identifies the `Snooze`.
type GetSnoozeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the `Snooze` to retrieve. The format is:
	//
	//	projects/[PROJECT_ID_OR_NUMBER]/snoozes/[SNOOZE_ID]
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSnoozeRequest) Reset() {
	*x = GetSnoozeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnoozeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnoozeRequest) ProtoMessage() {}

func (x *GetSnoozeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnoozeRequest.ProtoReflect.Descriptor instead.
func (*GetSnoozeRequest) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_snooze_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetSnoozeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The message definition for updating a `Snooze`. The field, `snooze.name`
// identifies the `Snooze` to be updated. The remainder of `snooze` gives the
// content the `Snooze` in question will be assigned.
//
// What fields can be updated depends on the start time and end time of the
// `Snooze`.
//
//   - end time is in the past: These `Snooze`s are considered
//     read-only and cannot be updated.
//   - start time is in the past and end time is in the future: `display_name`
//     and `interval.end_time` can be updated.
//   - start time is in the future: `display_name`, `interval.start_time` and
//     `interval.end_time` can be updated.
type UpdateSnoozeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The `Snooze` to update. Must have the name field present.
	Snooze *Snooze `protobuf:"bytes,1,opt,name=snooze,proto3" json:"snooze,omitempty"`
	// Required. The fields to update.
	//
	// For each field listed in `update_mask`:
	//
	//   - If the `Snooze` object supplied in the `UpdateSnoozeRequest` has a
	//     value for that field, the value of the field in the existing `Snooze`
	//     will be set to the value of the field in the supplied `Snooze`.
	//   - If the field does not have a value in the supplied `Snooze`, the field
	//     in the existing `Snooze` is set to its default value.
	//
	// Fields not listed retain their existing value.
	//
	// The following are the field names that are accepted in `update_mask`:
	//
	//   - `display_name`
	//   - `interval.start_time`
	//   - `interval.end_time`
	//
	// That said, the start time and end time of the `Snooze` determines which
	// fields can legally be updated. Before attempting an update, users should
	// consult the documentation for `UpdateSnoozeRequest`, which talks about
	// which fields can be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateSnoozeRequest) Reset() {
	*x = UpdateSnoozeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSnoozeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnoozeRequest) ProtoMessage() {}

func (x *UpdateSnoozeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_v3_snooze_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnoozeRequest.ProtoReflect.Descriptor instead.
func (*UpdateSnoozeRequest) Descriptor() ([]byte, []int) {
	return file_google_monitoring_v3_snooze_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSnoozeRequest) GetSnooze() *Snooze {
	if x != nil {
		return x.Snooze
	}
	return nil
}

func (x *UpdateSnoozeRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_google_monitoring_v3_snooze_service_proto protoreflect.FileDescriptor

var file_google_monitoring_v3_snooze_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x6e, 0x6f, 0x6f, 0x7a,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x28, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x12, 0x20, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6e, 0x6f, 0x6f, 0x7a,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x22, 0xb9,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x12, 0x20, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x75, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65,
	0x52, 0x07, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e,
	0x6f, 0x6f, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73,
	0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06,
	0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0x98, 0x06, 0x0a, 0x0d, 0x53, 0x6e, 0x6f,
	0x6f, 0x7a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x12, 0x29, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6e,
	0x6f, 0x6f, 0x7a, 0x65, 0x22, 0x3f, 0xda, 0x41, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c,
	0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x06, 0x73, 0x6e,
	0x6f, 0x6f, 0x7a, 0x65, 0x22, 0x1f, 0x2f, 0x76, 0x33, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x6e,
	0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x12, 0x94, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e,
	0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x6f, 0x6f, 0x7a,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0xda, 0x41, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x33,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x12, 0x81, 0x01, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x33, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65,
	0x22, 0x2e, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12,
	0x1f, 0x2f, 0x76, 0x33, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0xa4, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x6f, 0x6f, 0x7a,
	0x65, 0x12, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x22, 0x4b, 0xda, 0x41, 0x12, 0x73,
	0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x06, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x32,
	0x26, 0x2f, 0x76, 0x33, 0x2f, 0x7b, 0x73, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x6e, 0x6f,
	0x6f, 0x7a, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0xa9, 0x01, 0xca, 0x41, 0x19, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x89, 0x01, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x72,
	0x65, 0x61, 0x64, 0x42, 0xcd, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x42, 0x12, 0x53, 0x6e, 0x6f, 0x6f, 0x7a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x76, 0x32, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x5c, 0x56, 0x33, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_v3_snooze_service_proto_rawDescOnce sync.Once
	file_google_monitoring_v3_snooze_service_proto_rawDescData = file_google_monitoring_v3_snooze_service_proto_rawDesc
)

func file_google_monitoring_v3_snooze_service_proto_rawDescGZIP() []byte {
	file_google_monitoring_v3_snooze_service_proto_rawDescOnce.Do(func() {
		file_google_monitoring_v3_snooze_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_v3_snooze_service_proto_rawDescData)
	})
	return file_google_monitoring_v3_snooze_service_proto_rawDescData
}

var file_google_monitoring_v3_snooze_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_monitoring_v3_snooze_service_proto_goTypes = []interface{}{
	(*CreateSnoozeRequest)(nil),   // 0: google.monitoring.v3.CreateSnoozeRequest
	(*ListSnoozesRequest)(nil),    // 1: google.monitoring.v3.ListSnoozesRequest
	(*ListSnoozesResponse)(nil),   // 2: google.monitoring.v3.ListSnoozesResponse
	(*GetSnoozeRequest)(nil),      // 3: google.monitoring.v3.GetSnoozeRequest
	(*UpdateSnoozeRequest)(nil),   // 4: google.monitoring.v3.UpdateSnoozeRequest
	(*Snooze)(nil),                // 5: google.monitoring.v3.Snooze
	(*fieldmaskpb.FieldMask)(nil), // 6: google.protobuf.FieldMask
}
var file_google_monitoring_v3_snooze_service_proto_depIdxs = []int32{
	5, // 0: google.monitoring.v3.CreateSnoozeRequest.snooze:type_name -> google.monitoring.v3.Snooze
	5, // 1: google.monitoring.v3.ListSnoozesResponse.snoozes:type_name -> google.monitoring.v3.Snooze
	5, // 2: google.monitoring.v3.UpdateSnoozeRequest.snooze:type_name -> google.monitoring.v3.Snooze
	6, // 3: google.monitoring.v3.UpdateSnoozeRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: google.monitoring.v3.SnoozeService.CreateSnooze:input_type -> google.monitoring.v3.CreateSnoozeRequest
	1, // 5: google.monitoring.v3.SnoozeService.ListSnoozes:input_type -> google.monitoring.v3.ListSnoozesRequest
	3, // 6: google.monitoring.v3.SnoozeService.GetSnooze:input_type -> google.monitoring.v3.GetSnoozeRequest
	4, // 7: google.monitoring.v3.SnoozeService.UpdateSnooze:input_type -> google.monitoring.v3.UpdateSnoozeRequest
	5, // 8: google.monitoring.v3.SnoozeService.CreateSnooze:output_type -> google.monitoring.v3.Snooze
	2, // 9: google.monitoring.v3.SnoozeService.ListSnoozes:output_type -> google.monitoring.v3.ListSnoozesResponse
	5, // 10: google.monitoring.v3.SnoozeService.GetSnooze:output_type -> google.monitoring.v3.Snooze
	5, // 11: google.monitoring.v3.SnoozeService.UpdateSnooze:output_type -> google.monitoring.v3.Snooze
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_monitoring_v3_snooze_service_proto_init() }
func file_google_monitoring_v3_snooze_service_proto_init() {
	if File_google_monitoring_v3_snooze_service_proto != nil {
		return
	}
	file_google_monitoring_v3_snooze_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_v3_snooze_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnoozeRequest); i {
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
		file_google_monitoring_v3_snooze_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnoozesRequest); i {
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
		file_google_monitoring_v3_snooze_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnoozesResponse); i {
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
		file_google_monitoring_v3_snooze_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnoozeRequest); i {
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
		file_google_monitoring_v3_snooze_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSnoozeRequest); i {
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
			RawDescriptor: file_google_monitoring_v3_snooze_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_monitoring_v3_snooze_service_proto_goTypes,
		DependencyIndexes: file_google_monitoring_v3_snooze_service_proto_depIdxs,
		MessageInfos:      file_google_monitoring_v3_snooze_service_proto_msgTypes,
	}.Build()
	File_google_monitoring_v3_snooze_service_proto = out.File
	file_google_monitoring_v3_snooze_service_proto_rawDesc = nil
	file_google_monitoring_v3_snooze_service_proto_goTypes = nil
	file_google_monitoring_v3_snooze_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SnoozeServiceClient is the client API for SnoozeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SnoozeServiceClient interface {
	// Creates a `Snooze` that will prevent alerts, which match the provided
	// criteria, from being opened. The `Snooze` applies for a specific time
	// interval.
	CreateSnooze(ctx context.Context, in *CreateSnoozeRequest, opts ...grpc.CallOption) (*Snooze, error)
	// Lists the `Snooze`s associated with a project. Can optionally pass in
	// `filter`, which specifies predicates to match `Snooze`s.
	ListSnoozes(ctx context.Context, in *ListSnoozesRequest, opts ...grpc.CallOption) (*ListSnoozesResponse, error)
	// Retrieves a `Snooze` by `name`.
	GetSnooze(ctx context.Context, in *GetSnoozeRequest, opts ...grpc.CallOption) (*Snooze, error)
	// Updates a `Snooze`, identified by its `name`, with the parameters in the
	// given `Snooze` object.
	UpdateSnooze(ctx context.Context, in *UpdateSnoozeRequest, opts ...grpc.CallOption) (*Snooze, error)
}

type snoozeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnoozeServiceClient(cc grpc.ClientConnInterface) SnoozeServiceClient {
	return &snoozeServiceClient{cc}
}

func (c *snoozeServiceClient) CreateSnooze(ctx context.Context, in *CreateSnoozeRequest, opts ...grpc.CallOption) (*Snooze, error) {
	out := new(Snooze)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.SnoozeService/CreateSnooze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snoozeServiceClient) ListSnoozes(ctx context.Context, in *ListSnoozesRequest, opts ...grpc.CallOption) (*ListSnoozesResponse, error) {
	out := new(ListSnoozesResponse)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.SnoozeService/ListSnoozes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snoozeServiceClient) GetSnooze(ctx context.Context, in *GetSnoozeRequest, opts ...grpc.CallOption) (*Snooze, error) {
	out := new(Snooze)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.SnoozeService/GetSnooze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snoozeServiceClient) UpdateSnooze(ctx context.Context, in *UpdateSnoozeRequest, opts ...grpc.CallOption) (*Snooze, error) {
	out := new(Snooze)
	err := c.cc.Invoke(ctx, "/google.monitoring.v3.SnoozeService/UpdateSnooze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnoozeServiceServer is the server API for SnoozeService service.
type SnoozeServiceServer interface {
	// Creates a `Snooze` that will prevent alerts, which match the provided
	// criteria, from being opened. The `Snooze` applies for a specific time
	// interval.
	CreateSnooze(context.Context, *CreateSnoozeRequest) (*Snooze, error)
	// Lists the `Snooze`s associated with a project. Can optionally pass in
	// `filter`, which specifies predicates to match `Snooze`s.
	ListSnoozes(context.Context, *ListSnoozesRequest) (*ListSnoozesResponse, error)
	// Retrieves a `Snooze` by `name`.
	GetSnooze(context.Context, *GetSnoozeRequest) (*Snooze, error)
	// Updates a `Snooze`, identified by its `name`, with the parameters in the
	// given `Snooze` object.
	UpdateSnooze(context.Context, *UpdateSnoozeRequest) (*Snooze, error)
}

// UnimplementedSnoozeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSnoozeServiceServer struct {
}

func (*UnimplementedSnoozeServiceServer) CreateSnooze(context.Context, *CreateSnoozeRequest) (*Snooze, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnooze not implemented")
}
func (*UnimplementedSnoozeServiceServer) ListSnoozes(context.Context, *ListSnoozesRequest) (*ListSnoozesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnoozes not implemented")
}
func (*UnimplementedSnoozeServiceServer) GetSnooze(context.Context, *GetSnoozeRequest) (*Snooze, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnooze not implemented")
}
func (*UnimplementedSnoozeServiceServer) UpdateSnooze(context.Context, *UpdateSnoozeRequest) (*Snooze, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSnooze not implemented")
}

func RegisterSnoozeServiceServer(s *grpc.Server, srv SnoozeServiceServer) {
	s.RegisterService(&_SnoozeService_serviceDesc, srv)
}

func _SnoozeService_CreateSnooze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnoozeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnoozeServiceServer).CreateSnooze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.SnoozeService/CreateSnooze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnoozeServiceServer).CreateSnooze(ctx, req.(*CreateSnoozeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnoozeService_ListSnoozes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnoozesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnoozeServiceServer).ListSnoozes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.SnoozeService/ListSnoozes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnoozeServiceServer).ListSnoozes(ctx, req.(*ListSnoozesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnoozeService_GetSnooze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnoozeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnoozeServiceServer).GetSnooze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.SnoozeService/GetSnooze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnoozeServiceServer).GetSnooze(ctx, req.(*GetSnoozeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnoozeService_UpdateSnooze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnoozeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnoozeServiceServer).UpdateSnooze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.SnoozeService/UpdateSnooze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnoozeServiceServer).UpdateSnooze(ctx, req.(*UpdateSnoozeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SnoozeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.SnoozeService",
	HandlerType: (*SnoozeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSnooze",
			Handler:    _SnoozeService_CreateSnooze_Handler,
		},
		{
			MethodName: "ListSnoozes",
			Handler:    _SnoozeService_ListSnoozes_Handler,
		},
		{
			MethodName: "GetSnooze",
			Handler:    _SnoozeService_GetSnooze_Handler,
		},
		{
			MethodName: "UpdateSnooze",
			Handler:    _SnoozeService_UpdateSnooze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/monitoring/v3/snooze_service.proto",
}
