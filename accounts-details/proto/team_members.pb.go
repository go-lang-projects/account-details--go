// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: team_members.proto

package proto

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

type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_members_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_team_members_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_team_members_proto_rawDescGZIP(), []int{0}
}

func (x *Details) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Details) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SingleMember *Details   `protobuf:"bytes,2,opt,name=single_member,json=singleMember,proto3" json:"single_member,omitempty"`
	Members      []*Details `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_members_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_team_members_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_team_members_proto_rawDescGZIP(), []int{1}
}

func (x *Team) GetSingleMember() *Details {
	if x != nil {
		return x.SingleMember
	}
	return nil
}

func (x *Team) GetMembers() []*Details {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_team_members_proto protoreflect.FileDescriptor

var file_team_members_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a,
	0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x04,
	0x54, 0x65, 0x61, 0x6d, 0x12, 0x35, 0x0a, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0c, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x65, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_team_members_proto_rawDescOnce sync.Once
	file_team_members_proto_rawDescData = file_team_members_proto_rawDesc
)

func file_team_members_proto_rawDescGZIP() []byte {
	file_team_members_proto_rawDescOnce.Do(func() {
		file_team_members_proto_rawDescData = protoimpl.X.CompressGZIP(file_team_members_proto_rawDescData)
	})
	return file_team_members_proto_rawDescData
}

var file_team_members_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_team_members_proto_goTypes = []interface{}{
	(*Details)(nil), // 0: account.Details
	(*Team)(nil),    // 1: account.Team
}
var file_team_members_proto_depIdxs = []int32{
	0, // 0: account.Team.single_member:type_name -> account.Details
	0, // 1: account.Team.members:type_name -> account.Details
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_team_members_proto_init() }
func file_team_members_proto_init() {
	if File_team_members_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_team_members_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
		file_team_members_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
			RawDescriptor: file_team_members_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_team_members_proto_goTypes,
		DependencyIndexes: file_team_members_proto_depIdxs,
		MessageInfos:      file_team_members_proto_msgTypes,
	}.Build()
	File_team_members_proto = out.File
	file_team_members_proto_rawDesc = nil
	file_team_members_proto_goTypes = nil
	file_team_members_proto_depIdxs = nil
}