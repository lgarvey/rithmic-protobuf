// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: request_login.proto

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

type RequestLogin_SysInfraType int32

const (
	RequestLogin_TICKER_PLANT     RequestLogin_SysInfraType = 1
	RequestLogin_ORDER_PLANT      RequestLogin_SysInfraType = 2
	RequestLogin_HISTORY_PLANT    RequestLogin_SysInfraType = 3
	RequestLogin_PNL_PLANT        RequestLogin_SysInfraType = 4
	RequestLogin_REPOSITORY_PLANT RequestLogin_SysInfraType = 5
)

// Enum value maps for RequestLogin_SysInfraType.
var (
	RequestLogin_SysInfraType_name = map[int32]string{
		1: "TICKER_PLANT",
		2: "ORDER_PLANT",
		3: "HISTORY_PLANT",
		4: "PNL_PLANT",
		5: "REPOSITORY_PLANT",
	}
	RequestLogin_SysInfraType_value = map[string]int32{
		"TICKER_PLANT":     1,
		"ORDER_PLANT":      2,
		"HISTORY_PLANT":    3,
		"PNL_PLANT":        4,
		"REPOSITORY_PLANT": 5,
	}
)

func (x RequestLogin_SysInfraType) Enum() *RequestLogin_SysInfraType {
	p := new(RequestLogin_SysInfraType)
	*p = x
	return p
}

func (x RequestLogin_SysInfraType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestLogin_SysInfraType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_login_proto_enumTypes[0].Descriptor()
}

func (RequestLogin_SysInfraType) Type() protoreflect.EnumType {
	return &file_request_login_proto_enumTypes[0]
}

func (x RequestLogin_SysInfraType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestLogin_SysInfraType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestLogin_SysInfraType(num)
	return nil
}

// Deprecated: Use RequestLogin_SysInfraType.Descriptor instead.
func (RequestLogin_SysInfraType) EnumDescriptor() ([]byte, []int) {
	return file_request_login_proto_rawDescGZIP(), []int{0, 0}
}

type RequestLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId       *int32                     `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	TemplateVersion  *string                    `protobuf:"bytes,153634,opt,name=template_version,json=templateVersion" json:"template_version,omitempty"`
	UserMsg          []string                   `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	User             *string                    `protobuf:"bytes,131003,opt,name=user" json:"user,omitempty"`
	Password         *string                    `protobuf:"bytes,130004,opt,name=password" json:"password,omitempty"`
	AppName          *string                    `protobuf:"bytes,130002,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	AppVersion       *string                    `protobuf:"bytes,131803,opt,name=app_version,json=appVersion" json:"app_version,omitempty"`
	SystemName       *string                    `protobuf:"bytes,153628,opt,name=system_name,json=systemName" json:"system_name,omitempty"`
	InfraType        *RequestLogin_SysInfraType `protobuf:"varint,153621,opt,name=infra_type,json=infraType,enum=rti.RequestLogin_SysInfraType" json:"infra_type,omitempty"`
	MacAddr          []string                   `protobuf:"bytes,144108,rep,name=mac_addr,json=macAddr" json:"mac_addr,omitempty"`
	OsVersion        *string                    `protobuf:"bytes,144021,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	OsPlatform       *string                    `protobuf:"bytes,144020,opt,name=os_platform,json=osPlatform" json:"os_platform,omitempty"`
	AggregatedQuotes *bool                      `protobuf:"varint,153644,opt,name=aggregated_quotes,json=aggregatedQuotes" json:"aggregated_quotes,omitempty"` // applicable only for TICKER_PLANT infra_type
}

func (x *RequestLogin) Reset() {
	*x = RequestLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogin) ProtoMessage() {}

func (x *RequestLogin) ProtoReflect() protoreflect.Message {
	mi := &file_request_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogin.ProtoReflect.Descriptor instead.
func (*RequestLogin) Descriptor() ([]byte, []int) {
	return file_request_login_proto_rawDescGZIP(), []int{0}
}

func (x *RequestLogin) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestLogin) GetTemplateVersion() string {
	if x != nil && x.TemplateVersion != nil {
		return *x.TemplateVersion
	}
	return ""
}

func (x *RequestLogin) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestLogin) GetUser() string {
	if x != nil && x.User != nil {
		return *x.User
	}
	return ""
}

func (x *RequestLogin) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *RequestLogin) GetAppName() string {
	if x != nil && x.AppName != nil {
		return *x.AppName
	}
	return ""
}

func (x *RequestLogin) GetAppVersion() string {
	if x != nil && x.AppVersion != nil {
		return *x.AppVersion
	}
	return ""
}

func (x *RequestLogin) GetSystemName() string {
	if x != nil && x.SystemName != nil {
		return *x.SystemName
	}
	return ""
}

func (x *RequestLogin) GetInfraType() RequestLogin_SysInfraType {
	if x != nil && x.InfraType != nil {
		return *x.InfraType
	}
	return RequestLogin_TICKER_PLANT
}

func (x *RequestLogin) GetMacAddr() []string {
	if x != nil {
		return x.MacAddr
	}
	return nil
}

func (x *RequestLogin) GetOsVersion() string {
	if x != nil && x.OsVersion != nil {
		return *x.OsVersion
	}
	return ""
}

func (x *RequestLogin) GetOsPlatform() string {
	if x != nil && x.OsPlatform != nil {
		return *x.OsPlatform
	}
	return ""
}

func (x *RequestLogin) GetAggregatedQuotes() bool {
	if x != nil && x.AggregatedQuotes != nil {
		return *x.AggregatedQuotes
	}
	return false
}

var File_request_login_proto protoreflect.FileDescriptor

var file_request_login_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xce, 0x04, 0x0a, 0x0c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0xa2, 0xb0, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0xbb, 0xff, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0xd4, 0xf7, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xd2, 0xf7, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x61, 0x70, 0x70,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xdb, 0x85, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0b,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x9c, 0xb0, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x95, 0xb0,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0xec, 0xe5, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a,
	0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x95, 0xe5, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0b, 0x6f, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x94, 0xe5,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x2d, 0x0a, 0x11, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0xac, 0xb0, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x22, 0x69, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x4e,
	0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x50,
	0x4c, 0x41, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x4e, 0x4c, 0x5f, 0x50, 0x4c,
	0x41, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x4f, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x54, 0x10, 0x05, 0x42, 0x1e, 0x5a, 0x1c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76, 0x65,
	0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_request_login_proto_rawDescOnce sync.Once
	file_request_login_proto_rawDescData = file_request_login_proto_rawDesc
)

func file_request_login_proto_rawDescGZIP() []byte {
	file_request_login_proto_rawDescOnce.Do(func() {
		file_request_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_login_proto_rawDescData)
	})
	return file_request_login_proto_rawDescData
}

var file_request_login_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_request_login_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_login_proto_goTypes = []interface{}{
	(RequestLogin_SysInfraType)(0), // 0: rti.RequestLogin.SysInfraType
	(*RequestLogin)(nil),           // 1: rti.RequestLogin
}
var file_request_login_proto_depIdxs = []int32{
	0, // 0: rti.RequestLogin.infra_type:type_name -> rti.RequestLogin.SysInfraType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_request_login_proto_init() }
func file_request_login_proto_init() {
	if File_request_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogin); i {
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
			RawDescriptor: file_request_login_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_login_proto_goTypes,
		DependencyIndexes: file_request_login_proto_depIdxs,
		EnumInfos:         file_request_login_proto_enumTypes,
		MessageInfos:      file_request_login_proto_msgTypes,
	}.Build()
	File_request_login_proto = out.File
	file_request_login_proto_rawDesc = nil
	file_request_login_proto_goTypes = nil
	file_request_login_proto_depIdxs = nil
}
