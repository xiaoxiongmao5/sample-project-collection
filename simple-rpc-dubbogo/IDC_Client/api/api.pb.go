// EDIT IT, change to your package, service and message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: api.proto

package api

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

type GenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId  string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	AppAge int32  `protobuf:"varint,2,opt,name=appAge,proto3" json:"appAge,omitempty"`
}

func (x *GenReq) Reset() {
	*x = GenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenReq) ProtoMessage() {}

func (x *GenReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenReq.ProtoReflect.Descriptor instead.
func (*GenReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GenReq) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *GenReq) GetAppAge() int32 {
	if x != nil {
		return x.AppAge
	}
	return 0
}

type GenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *GenResp) Reset() {
	*x = GenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenResp) ProtoMessage() {}

func (x *GenResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenResp.ProtoReflect.Descriptor instead.
func (*GenResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GenResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenResp) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type XtestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *XtestReq) Reset() {
	*x = XtestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XtestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XtestReq) ProtoMessage() {}

func (x *XtestReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XtestReq.ProtoReflect.Descriptor instead.
func (*XtestReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *XtestReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type XtestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Sex  string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
}

func (x *XtestResp) Reset() {
	*x = XtestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XtestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XtestResp) ProtoMessage() {}

func (x *XtestResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XtestResp.ProtoReflect.Descriptor instead.
func (*XtestResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *XtestResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XtestResp) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *XtestResp) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0x36, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x61, 0x70, 0x70, 0x41, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x58, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x09, 0x58, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x32, 0x32,
	0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x31, 0x12, 0x24, 0x0a, 0x05,
	0x47, 0x65, 0x74, 0x49, 0x44, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x32, 0x33, 0x0a, 0x05, 0x58, 0x74, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x58, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x58, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_goTypes = []interface{}{
	(*GenReq)(nil),    // 0: api.GenReq
	(*GenResp)(nil),   // 1: api.GenResp
	(*XtestReq)(nil),  // 2: api.XtestReq
	(*XtestResp)(nil), // 3: api.XtestResp
}
var file_api_proto_depIdxs = []int32{
	0, // 0: api.Generator1.GetID:input_type -> api.GenReq
	2, // 1: api.Xtest.GetUser:input_type -> api.XtestReq
	1, // 2: api.Generator1.GetID:output_type -> api.GenResp
	3, // 3: api.Xtest.GetUser:output_type -> api.XtestResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenReq); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenResp); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XtestReq); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XtestResp); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
