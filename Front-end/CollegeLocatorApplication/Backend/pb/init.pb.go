// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: init.proto

package pb

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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_init_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_init_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image        string  `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Address      string  `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	CollegeId    string  `protobuf:"bytes,5,opt,name=collegeId,proto3" json:"collegeId,omitempty"`
	MaxRangeFees float32 `protobuf:"fixed32,6,opt,name=maxRangeFees,proto3" json:"maxRangeFees,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_init_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_init_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *InitResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InitResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InitResponse) GetCollegeId() string {
	if x != nil {
		return x.CollegeId
	}
	return ""
}

func (x *InitResponse) GetMaxRangeFees() float32 {
	if x != nil {
		return x.MaxRangeFees
	}
	return 0
}

var File_init_proto protoreflect.FileDescriptor

var file_init_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0b, 0x49, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x94, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65,
	0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x46,
	0x65, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x46, 0x65, 0x65, 0x73, 0x42, 0x3f, 0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x67, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_init_proto_rawDescOnce sync.Once
	file_init_proto_rawDescData = file_init_proto_rawDesc
)

func file_init_proto_rawDescGZIP() []byte {
	file_init_proto_rawDescOnce.Do(func() {
		file_init_proto_rawDescData = protoimpl.X.CompressGZIP(file_init_proto_rawDescData)
	})
	return file_init_proto_rawDescData
}

var file_init_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_init_proto_goTypes = []interface{}{
	(*InitRequest)(nil),  // 0: InitRequest
	(*InitResponse)(nil), // 1: InitResponse
	(*Location)(nil),     // 2: Location
}
var file_init_proto_depIdxs = []int32{
	2, // 0: InitRequest.location:type_name -> Location
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_init_proto_init() }
func file_init_proto_init() {
	if File_init_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_init_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_init_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
			RawDescriptor: file_init_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_init_proto_goTypes,
		DependencyIndexes: file_init_proto_depIdxs,
		MessageInfos:      file_init_proto_msgTypes,
	}.Build()
	File_init_proto = out.File
	file_init_proto_rawDesc = nil
	file_init_proto_goTypes = nil
	file_init_proto_depIdxs = nil
}
