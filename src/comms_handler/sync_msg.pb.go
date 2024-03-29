// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: sync_msg.proto

package comms_handler

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

type SyncMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SyncMessage) Reset() {
	*x = SyncMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessage) ProtoMessage() {}

func (x *SyncMessage) ProtoReflect() protoreflect.Message {
	mi := &file_sync_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessage.ProtoReflect.Descriptor instead.
func (*SyncMessage) Descriptor() ([]byte, []int) {
	return file_sync_msg_proto_rawDescGZIP(), []int{0}
}

func (x *SyncMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SyncMessage) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

var File_sync_msg_proto protoreflect.FileDescriptor

var file_sync_msg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x33, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x73,
	0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_msg_proto_rawDescOnce sync.Once
	file_sync_msg_proto_rawDescData = file_sync_msg_proto_rawDesc
)

func file_sync_msg_proto_rawDescGZIP() []byte {
	file_sync_msg_proto_rawDescOnce.Do(func() {
		file_sync_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_msg_proto_rawDescData)
	})
	return file_sync_msg_proto_rawDescData
}

var file_sync_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sync_msg_proto_goTypes = []interface{}{
	(*SyncMessage)(nil), // 0: SyncMessage
}
var file_sync_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sync_msg_proto_init() }
func file_sync_msg_proto_init() {
	if File_sync_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMessage); i {
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
			RawDescriptor: file_sync_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_msg_proto_goTypes,
		DependencyIndexes: file_sync_msg_proto_depIdxs,
		MessageInfos:      file_sync_msg_proto_msgTypes,
	}.Build()
	File_sync_msg_proto = out.File
	file_sync_msg_proto_rawDesc = nil
	file_sync_msg_proto_goTypes = nil
	file_sync_msg_proto_depIdxs = nil
}
