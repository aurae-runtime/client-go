// -------------------------------------------------------------------------- *\
//             Apache 2.0 License Copyright © 2022 The Aurae Authors          *
//                                                                            *
//                +--------------------------------------------+              *
//                |   █████╗ ██╗   ██╗██████╗  █████╗ ███████╗ |              *
//                |  ██╔══██╗██║   ██║██╔══██╗██╔══██╗██╔════╝ |              *
//                |  ███████║██║   ██║██████╔╝███████║█████╗   |              *
//                |  ██╔══██║██║   ██║██╔══██╗██╔══██║██╔══╝   |              *
//                |  ██║  ██║╚██████╔╝██║  ██║██║  ██║███████╗ |              *
//                |  ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ |              *
//                +--------------------------------------------+              *
//                                                                            *
//                         Distributed Systems Runtime                        *
//                                                                            *
// -------------------------------------------------------------------------- *
//                                                                            *
//   Licensed under the Apache License, Version 2.0 (the "License");          *
//   you may not use this file except in compliance with the License.         *
//   You may obtain a copy of the License at                                  *
//                                                                            *
//       http://www.apache.org/licenses/LICENSE-2.0                           *
//                                                                            *
//   Unless required by applicable law or agreed to in writing, software      *
//   distributed under the License is distributed on an "AS IS" BASIS,        *
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
//   See the License for the specific language governing permissions and      *
//   limitations under the License.                                           *
//                                                                            *
//\* --------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: observe.proto

package observe

import (
	meta "github.com/aurae-runtime/client-go/pkg/api/v0/meta"
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

type LogChannelType int32

const (
	LogChannelType_CHANNEL_STDOUT LogChannelType = 0
	LogChannelType_CHANNEL_STDERR LogChannelType = 1
)

// Enum value maps for LogChannelType.
var (
	LogChannelType_name = map[int32]string{
		0: "CHANNEL_STDOUT",
		1: "CHANNEL_STDERR",
	}
	LogChannelType_value = map[string]int32{
		"CHANNEL_STDOUT": 0,
		"CHANNEL_STDERR": 1,
	}
)

func (x LogChannelType) Enum() *LogChannelType {
	p := new(LogChannelType)
	*p = x
	return p
}

func (x LogChannelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogChannelType) Descriptor() protoreflect.EnumDescriptor {
	return file_observe_proto_enumTypes[0].Descriptor()
}

func (LogChannelType) Type() protoreflect.EnumType {
	return &file_observe_proto_enumTypes[0]
}

func (x LogChannelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogChannelType.Descriptor instead.
func (LogChannelType) EnumDescriptor() ([]byte, []int) {
	return file_observe_proto_rawDescGZIP(), []int{0}
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *meta.AuraeMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_observe_proto_rawDescGZIP(), []int{0}
}

func (x *StatusRequest) GetMeta() *meta.AuraeMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *meta.AuraeMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_observe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_observe_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetMeta() *meta.AuraeMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GetAuraeDaemonLogStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuraeDaemonLogStreamRequest) Reset() {
	*x = GetAuraeDaemonLogStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuraeDaemonLogStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuraeDaemonLogStreamRequest) ProtoMessage() {}

func (x *GetAuraeDaemonLogStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuraeDaemonLogStreamRequest.ProtoReflect.Descriptor instead.
func (*GetAuraeDaemonLogStreamRequest) Descriptor() ([]byte, []int) {
	return file_observe_proto_rawDescGZIP(), []int{2}
}

// TODO: not implemented
type GetSubProcessStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelType LogChannelType `protobuf:"varint,1,opt,name=channel_type,json=channelType,proto3,enum=observe.LogChannelType" json:"channel_type,omitempty"`
	ProcessId   int64          `protobuf:"varint,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
}

func (x *GetSubProcessStreamRequest) Reset() {
	*x = GetSubProcessStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubProcessStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubProcessStreamRequest) ProtoMessage() {}

func (x *GetSubProcessStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubProcessStreamRequest.ProtoReflect.Descriptor instead.
func (*GetSubProcessStreamRequest) Descriptor() ([]byte, []int) {
	return file_observe_proto_rawDescGZIP(), []int{3}
}

func (x *GetSubProcessStreamRequest) GetChannelType() LogChannelType {
	if x != nil {
		return x.ChannelType
	}
	return LogChannelType_CHANNEL_STDOUT
}

func (x *GetSubProcessStreamRequest) GetProcessId() int64 {
	if x != nil {
		return x.ProcessId
	}
	return 0
}

type LogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel   string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Line      string `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LogItem) Reset() {
	*x = LogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observe_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogItem) ProtoMessage() {}

func (x *LogItem) ProtoReflect() protoreflect.Message {
	mi := &file_observe_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogItem.ProtoReflect.Descriptor instead.
func (*LogItem) Descriptor() ([]byte, []int) {
	return file_observe_proto_rawDescGZIP(), []int{4}
}

func (x *LogItem) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *LogItem) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

func (x *LogItem) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_observe_proto protoreflect.FileDescriptor

var file_observe_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x1a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x75, 0x72, 0x61, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x0e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x41, 0x75, 0x72, 0x61, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x72, 0x61, 0x65, 0x44, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x07,
	0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2a, 0x38, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x53, 0x54, 0x44, 0x4f, 0x55, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x44, 0x45, 0x52, 0x52, 0x10, 0x01, 0x32, 0xf2, 0x01,
	0x0a, 0x07, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x75, 0x72,
	0x61, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x27, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x72, 0x61, 0x65, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x50, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x72, 0x61, 0x65, 0x2d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x30, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_observe_proto_rawDescOnce sync.Once
	file_observe_proto_rawDescData = file_observe_proto_rawDesc
)

func file_observe_proto_rawDescGZIP() []byte {
	file_observe_proto_rawDescOnce.Do(func() {
		file_observe_proto_rawDescData = protoimpl.X.CompressGZIP(file_observe_proto_rawDescData)
	})
	return file_observe_proto_rawDescData
}

var file_observe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_observe_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_observe_proto_goTypes = []interface{}{
	(LogChannelType)(0),                    // 0: observe.LogChannelType
	(*StatusRequest)(nil),                  // 1: observe.StatusRequest
	(*StatusResponse)(nil),                 // 2: observe.StatusResponse
	(*GetAuraeDaemonLogStreamRequest)(nil), // 3: observe.GetAuraeDaemonLogStreamRequest
	(*GetSubProcessStreamRequest)(nil),     // 4: observe.GetSubProcessStreamRequest
	(*LogItem)(nil),                        // 5: observe.LogItem
	(*meta.AuraeMeta)(nil),                 // 6: meta.AuraeMeta
}
var file_observe_proto_depIdxs = []int32{
	6, // 0: observe.StatusRequest.meta:type_name -> meta.AuraeMeta
	6, // 1: observe.StatusResponse.meta:type_name -> meta.AuraeMeta
	0, // 2: observe.GetSubProcessStreamRequest.channel_type:type_name -> observe.LogChannelType
	1, // 3: observe.Observe.Status:input_type -> observe.StatusRequest
	3, // 4: observe.Observe.GetAuraeDaemonLogStream:input_type -> observe.GetAuraeDaemonLogStreamRequest
	4, // 5: observe.Observe.GetSubProcessStream:input_type -> observe.GetSubProcessStreamRequest
	2, // 6: observe.Observe.Status:output_type -> observe.StatusResponse
	5, // 7: observe.Observe.GetAuraeDaemonLogStream:output_type -> observe.LogItem
	5, // 8: observe.Observe.GetSubProcessStream:output_type -> observe.LogItem
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_observe_proto_init() }
func file_observe_proto_init() {
	if File_observe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_observe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_observe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_observe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuraeDaemonLogStreamRequest); i {
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
		file_observe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubProcessStreamRequest); i {
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
		file_observe_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogItem); i {
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
			RawDescriptor: file_observe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_observe_proto_goTypes,
		DependencyIndexes: file_observe_proto_depIdxs,
		EnumInfos:         file_observe_proto_enumTypes,
		MessageInfos:      file_observe_proto_msgTypes,
	}.Build()
	File_observe_proto = out.File
	file_observe_proto_rawDesc = nil
	file_observe_proto_goTypes = nil
	file_observe_proto_depIdxs = nil
}
