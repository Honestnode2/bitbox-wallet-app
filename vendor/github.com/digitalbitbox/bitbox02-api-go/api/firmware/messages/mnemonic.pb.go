// Copyright 2019 Shift Cryptosecurity AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: mnemonic.proto

package messages

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

type ShowMnemonicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShowMnemonicRequest) Reset() {
	*x = ShowMnemonicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mnemonic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowMnemonicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowMnemonicRequest) ProtoMessage() {}

func (x *ShowMnemonicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mnemonic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowMnemonicRequest.ProtoReflect.Descriptor instead.
func (*ShowMnemonicRequest) Descriptor() ([]byte, []int) {
	return file_mnemonic_proto_rawDescGZIP(), []int{0}
}

type RestoreFromMnemonicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp      uint32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimezoneOffset int32  `protobuf:"varint,2,opt,name=timezone_offset,json=timezoneOffset,proto3" json:"timezone_offset,omitempty"`
}

func (x *RestoreFromMnemonicRequest) Reset() {
	*x = RestoreFromMnemonicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mnemonic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreFromMnemonicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreFromMnemonicRequest) ProtoMessage() {}

func (x *RestoreFromMnemonicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mnemonic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreFromMnemonicRequest.ProtoReflect.Descriptor instead.
func (*RestoreFromMnemonicRequest) Descriptor() ([]byte, []int) {
	return file_mnemonic_proto_rawDescGZIP(), []int{1}
}

func (x *RestoreFromMnemonicRequest) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RestoreFromMnemonicRequest) GetTimezoneOffset() int32 {
	if x != nil {
		return x.TimezoneOffset
	}
	return 0
}

type SetMnemonicPassphraseEnabledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *SetMnemonicPassphraseEnabledRequest) Reset() {
	*x = SetMnemonicPassphraseEnabledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mnemonic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMnemonicPassphraseEnabledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMnemonicPassphraseEnabledRequest) ProtoMessage() {}

func (x *SetMnemonicPassphraseEnabledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mnemonic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMnemonicPassphraseEnabledRequest.ProtoReflect.Descriptor instead.
func (*SetMnemonicPassphraseEnabledRequest) Descriptor() ([]byte, []int) {
	return file_mnemonic_proto_rawDescGZIP(), []int{2}
}

func (x *SetMnemonicPassphraseEnabledRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_mnemonic_proto protoreflect.FileDescriptor

var file_mnemonic_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x73, 0x68, 0x69, 0x66, 0x74, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x62, 0x69,
	0x74, 0x62, 0x6f, 0x78, 0x30, 0x32, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x68, 0x6f, 0x77, 0x4d, 0x6e,
	0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a,
	0x1a, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x3f, 0x0a, 0x23, 0x53, 0x65, 0x74, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69,
	0x63, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mnemonic_proto_rawDescOnce sync.Once
	file_mnemonic_proto_rawDescData = file_mnemonic_proto_rawDesc
)

func file_mnemonic_proto_rawDescGZIP() []byte {
	file_mnemonic_proto_rawDescOnce.Do(func() {
		file_mnemonic_proto_rawDescData = protoimpl.X.CompressGZIP(file_mnemonic_proto_rawDescData)
	})
	return file_mnemonic_proto_rawDescData
}

var file_mnemonic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mnemonic_proto_goTypes = []interface{}{
	(*ShowMnemonicRequest)(nil),                 // 0: shiftcrypto.bitbox02.ShowMnemonicRequest
	(*RestoreFromMnemonicRequest)(nil),          // 1: shiftcrypto.bitbox02.RestoreFromMnemonicRequest
	(*SetMnemonicPassphraseEnabledRequest)(nil), // 2: shiftcrypto.bitbox02.SetMnemonicPassphraseEnabledRequest
}
var file_mnemonic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mnemonic_proto_init() }
func file_mnemonic_proto_init() {
	if File_mnemonic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mnemonic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowMnemonicRequest); i {
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
		file_mnemonic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreFromMnemonicRequest); i {
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
		file_mnemonic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMnemonicPassphraseEnabledRequest); i {
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
			RawDescriptor: file_mnemonic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mnemonic_proto_goTypes,
		DependencyIndexes: file_mnemonic_proto_depIdxs,
		MessageInfos:      file_mnemonic_proto_msgTypes,
	}.Build()
	File_mnemonic_proto = out.File
	file_mnemonic_proto_rawDesc = nil
	file_mnemonic_proto_goTypes = nil
	file_mnemonic_proto_depIdxs = nil
}
