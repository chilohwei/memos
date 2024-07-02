// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: store/inbox.proto

package store

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

type InboxMessage_Type int32

const (
	InboxMessage_TYPE_UNSPECIFIED InboxMessage_Type = 0
	InboxMessage_MEMO_COMMENT     InboxMessage_Type = 1
	InboxMessage_VERSION_UPDATE   InboxMessage_Type = 2
)

// Enum value maps for InboxMessage_Type.
var (
	InboxMessage_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "MEMO_COMMENT",
		2: "VERSION_UPDATE",
	}
	InboxMessage_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"MEMO_COMMENT":     1,
		"VERSION_UPDATE":   2,
	}
)

func (x InboxMessage_Type) Enum() *InboxMessage_Type {
	p := new(InboxMessage_Type)
	*p = x
	return p
}

func (x InboxMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InboxMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_inbox_proto_enumTypes[0].Descriptor()
}

func (InboxMessage_Type) Type() protoreflect.EnumType {
	return &file_store_inbox_proto_enumTypes[0]
}

func (x InboxMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InboxMessage_Type.Descriptor instead.
func (InboxMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_inbox_proto_rawDescGZIP(), []int{0, 0}
}

type InboxMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       InboxMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=memos.store.InboxMessage_Type" json:"type,omitempty"`
	ActivityId *int32            `protobuf:"varint,2,opt,name=activity_id,json=activityId,proto3,oneof" json:"activity_id,omitempty"`
}

func (x *InboxMessage) Reset() {
	*x = InboxMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_inbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InboxMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InboxMessage) ProtoMessage() {}

func (x *InboxMessage) ProtoReflect() protoreflect.Message {
	mi := &file_store_inbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InboxMessage.ProtoReflect.Descriptor instead.
func (*InboxMessage) Descriptor() ([]byte, []int) {
	return file_store_inbox_proto_rawDescGZIP(), []int{0}
}

func (x *InboxMessage) GetType() InboxMessage_Type {
	if x != nil {
		return x.Type
	}
	return InboxMessage_TYPE_UNSPECIFIED
}

func (x *InboxMessage) GetActivityId() int32 {
	if x != nil && x.ActivityId != nil {
		return *x.ActivityId
	}
	return 0
}

var File_store_inbox_proto protoreflect.FileDescriptor

var file_store_inbox_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e,
	0x62, 0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x22, 0x42, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x4d,
	0x4f, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x42,
	0x95, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x42, 0x0a, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73,
	0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x4d,
	0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02,
	0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73,
	0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_inbox_proto_rawDescOnce sync.Once
	file_store_inbox_proto_rawDescData = file_store_inbox_proto_rawDesc
)

func file_store_inbox_proto_rawDescGZIP() []byte {
	file_store_inbox_proto_rawDescOnce.Do(func() {
		file_store_inbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_inbox_proto_rawDescData)
	})
	return file_store_inbox_proto_rawDescData
}

var file_store_inbox_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_inbox_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_inbox_proto_goTypes = []any{
	(InboxMessage_Type)(0), // 0: memos.store.InboxMessage.Type
	(*InboxMessage)(nil),   // 1: memos.store.InboxMessage
}
var file_store_inbox_proto_depIdxs = []int32{
	0, // 0: memos.store.InboxMessage.type:type_name -> memos.store.InboxMessage.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_store_inbox_proto_init() }
func file_store_inbox_proto_init() {
	if File_store_inbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_inbox_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InboxMessage); i {
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
	file_store_inbox_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_inbox_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_inbox_proto_goTypes,
		DependencyIndexes: file_store_inbox_proto_depIdxs,
		EnumInfos:         file_store_inbox_proto_enumTypes,
		MessageInfos:      file_store_inbox_proto_msgTypes,
	}.Build()
	File_store_inbox_proto = out.File
	file_store_inbox_proto_rawDesc = nil
	file_store_inbox_proto_goTypes = nil
	file_store_inbox_proto_depIdxs = nil
}
