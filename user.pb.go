// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: user.proto

package grpc

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

// 列挙型
type Sample_EnumData int32

const (
	Sample_UNKNOWN Sample_EnumData = 0
	Sample_DATA1   Sample_EnumData = 1
	Sample_DATA2   Sample_EnumData = 2
	Sample_DATA3   Sample_EnumData = 3
)

// Enum value maps for Sample_EnumData.
var (
	Sample_EnumData_name = map[int32]string{
		0: "UNKNOWN",
		1: "DATA1",
		2: "DATA2",
		3: "DATA3",
	}
	Sample_EnumData_value = map[string]int32{
		"UNKNOWN": 0,
		"DATA1":   1,
		"DATA2":   2,
		"DATA3":   3,
	}
)

func (x Sample_EnumData) Enum() *Sample_EnumData {
	p := new(Sample_EnumData)
	*p = x
	return p
}

func (x Sample_EnumData) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sample_EnumData) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Sample_EnumData) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Sample_EnumData) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sample_EnumData.Descriptor instead.
func (Sample_EnumData) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2, 0}
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age   int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserResponse) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数値、文字列、真偽値、バイト配列がある
	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsBool   bool   `protobuf:"varint,3,opt,name=isBool,proto3" json:"isBool,omitempty"`
	ByteData []byte `protobuf:"bytes,4,opt,name=byteData,proto3" json:"byteData,omitempty"`
	// 廃止予定かつ非推奨であるフィールドを明示する
	//
	// Deprecated: Do not use.
	DuplicatedField string `protobuf:"bytes,5,opt,name=duplicated_field,json=duplicatedField,proto3" json:"duplicated_field,omitempty"`
	// リスト（配列）
	List []string `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	// マップ（連想配列）
	Map map[string]string `protobuf:"bytes,7,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 複数の中からどれかひとつ
	//
	// Types that are assignable to Message:
	//
	//	*Sample_One
	//	*Sample_Other
	Message isSample_Message `protobuf_oneof:"message"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *Sample) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sample) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sample) GetIsBool() bool {
	if x != nil {
		return x.IsBool
	}
	return false
}

func (x *Sample) GetByteData() []byte {
	if x != nil {
		return x.ByteData
	}
	return nil
}

// Deprecated: Do not use.
func (x *Sample) GetDuplicatedField() string {
	if x != nil {
		return x.DuplicatedField
	}
	return ""
}

func (x *Sample) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Sample) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

func (m *Sample) GetMessage() isSample_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *Sample) GetOne() string {
	if x, ok := x.GetMessage().(*Sample_One); ok {
		return x.One
	}
	return ""
}

func (x *Sample) GetOther() string {
	if x, ok := x.GetMessage().(*Sample_Other); ok {
		return x.Other
	}
	return ""
}

type isSample_Message interface {
	isSample_Message()
}

type Sample_One struct {
	One string `protobuf:"bytes,8,opt,name=one,proto3,oneof"`
}

type Sample_Other struct {
	Other string `protobuf:"bytes,9,opt,name=other,proto3,oneof"`
}

func (*Sample_One) isSample_Message() {}

func (*Sample_Other) isSample_Message() {}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5a, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x87, 0x03,
	0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x79, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2d, 0x0a, 0x10, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f,
	0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x03,
	0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x38, 0x0a, 0x08, 0x45, 0x6e, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41, 0x54,
	0x41, 0x31, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41, 0x54, 0x41, 0x32, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x44, 0x41, 0x54, 0x41, 0x33, 0x10, 0x03, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x0b, 0x10,
	0x0c, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0f, 0x32, 0x91, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x16, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x16,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12,
	0x4a, 0x0a, 0x1d, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []interface{}{
	(Sample_EnumData)(0), // 0: user.Sample.EnumData
	(*UserRequest)(nil),  // 1: user.UserRequest
	(*UserResponse)(nil), // 2: user.UserResponse
	(*Sample)(nil),       // 3: user.Sample
	nil,                  // 4: user.Sample.MapEntry
}
var file_user_proto_depIdxs = []int32{
	4, // 0: user.Sample.map:type_name -> user.Sample.MapEntry
	1, // 1: user.UserService.GetUser:input_type -> user.UserRequest
	1, // 2: user.UserService.ServerStreamingGetUser:input_type -> user.UserRequest
	1, // 3: user.UserService.ClientStreamingGetUser:input_type -> user.UserRequest
	1, // 4: user.UserService.BidirectionalStreamingGetUser:input_type -> user.UserRequest
	2, // 5: user.UserService.GetUser:output_type -> user.UserResponse
	2, // 6: user.UserService.ServerStreamingGetUser:output_type -> user.UserResponse
	2, // 7: user.UserService.ClientStreamingGetUser:output_type -> user.UserResponse
	2, // 8: user.UserService.BidirectionalStreamingGetUser:output_type -> user.UserResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
	file_user_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Sample_One)(nil),
		(*Sample_Other)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
