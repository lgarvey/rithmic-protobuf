// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: request_account_list.proto

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

type RequestAccountList_UserType int32

const (
	RequestAccountList_USER_TYPE_FCM    RequestAccountList_UserType = 1
	RequestAccountList_USER_TYPE_IB     RequestAccountList_UserType = 2
	RequestAccountList_USER_TYPE_TRADER RequestAccountList_UserType = 3
)

// Enum value maps for RequestAccountList_UserType.
var (
	RequestAccountList_UserType_name = map[int32]string{
		1: "USER_TYPE_FCM",
		2: "USER_TYPE_IB",
		3: "USER_TYPE_TRADER",
	}
	RequestAccountList_UserType_value = map[string]int32{
		"USER_TYPE_FCM":    1,
		"USER_TYPE_IB":     2,
		"USER_TYPE_TRADER": 3,
	}
)

func (x RequestAccountList_UserType) Enum() *RequestAccountList_UserType {
	p := new(RequestAccountList_UserType)
	*p = x
	return p
}

func (x RequestAccountList_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestAccountList_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_account_list_proto_enumTypes[0].Descriptor()
}

func (RequestAccountList_UserType) Type() protoreflect.EnumType {
	return &file_request_account_list_proto_enumTypes[0]
}

func (x RequestAccountList_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestAccountList_UserType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestAccountList_UserType(num)
	return nil
}

// Deprecated: Use RequestAccountList_UserType.Descriptor instead.
func (RequestAccountList_UserType) EnumDescriptor() ([]byte, []int) {
	return file_request_account_list_proto_rawDescGZIP(), []int{0, 0}
}

type RequestAccountList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32                       `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                                // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg    []string                     `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                          // PB_OFFSET + MNM_USER_MSG
	FcmId      *string                      `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`                                                // PB_OFFSET + MNM_FCM_ID
	IbId       *string                      `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`                                                   // PB_OFFSET + MNM_IB_ID
	UserType   *RequestAccountList_UserType `protobuf:"varint,154036,opt,name=user_type,json=userType,enum=rti.RequestAccountList_UserType" json:"user_type,omitempty"` // PB_OFFSET + MNM_USER_TYPE
}

func (x *RequestAccountList) Reset() {
	*x = RequestAccountList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_account_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAccountList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAccountList) ProtoMessage() {}

func (x *RequestAccountList) ProtoReflect() protoreflect.Message {
	mi := &file_request_account_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAccountList.ProtoReflect.Descriptor instead.
func (*RequestAccountList) Descriptor() ([]byte, []int) {
	return file_request_account_list_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAccountList) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestAccountList) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestAccountList) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *RequestAccountList) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *RequestAccountList) GetUserType() RequestAccountList_UserType {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return RequestAccountList_USER_TYPE_FCM
}

var File_request_account_list_proto protoreflect.FileDescriptor

var file_request_account_list_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74,
	0x69, 0x22, 0x8c, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x05, 0x69, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xb4, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x43, 0x4d, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x42, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x45, 0x52, 0x10, 0x03,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_request_account_list_proto_rawDescOnce sync.Once
	file_request_account_list_proto_rawDescData = file_request_account_list_proto_rawDesc
)

func file_request_account_list_proto_rawDescGZIP() []byte {
	file_request_account_list_proto_rawDescOnce.Do(func() {
		file_request_account_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_account_list_proto_rawDescData)
	})
	return file_request_account_list_proto_rawDescData
}

var file_request_account_list_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_request_account_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_account_list_proto_goTypes = []interface{}{
	(RequestAccountList_UserType)(0), // 0: rti.RequestAccountList.UserType
	(*RequestAccountList)(nil),       // 1: rti.RequestAccountList
}
var file_request_account_list_proto_depIdxs = []int32{
	0, // 0: rti.RequestAccountList.user_type:type_name -> rti.RequestAccountList.UserType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_request_account_list_proto_init() }
func file_request_account_list_proto_init() {
	if File_request_account_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_account_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAccountList); i {
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
			RawDescriptor: file_request_account_list_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_account_list_proto_goTypes,
		DependencyIndexes: file_request_account_list_proto_depIdxs,
		EnumInfos:         file_request_account_list_proto_enumTypes,
		MessageInfos:      file_request_account_list_proto_msgTypes,
	}.Build()
	File_request_account_list_proto = out.File
	file_request_account_list_proto_rawDesc = nil
	file_request_account_list_proto_goTypes = nil
	file_request_account_list_proto_depIdxs = nil
}
