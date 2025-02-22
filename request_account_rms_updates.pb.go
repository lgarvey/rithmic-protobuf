// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: request_account_rms_updates.proto

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

type RequestAccountRmsUpdates_UpdateBits int32

const (
	RequestAccountRmsUpdates_AUTO_LIQ_THRESHOLD_CURRENT_VALUE RequestAccountRmsUpdates_UpdateBits = 1
)

// Enum value maps for RequestAccountRmsUpdates_UpdateBits.
var (
	RequestAccountRmsUpdates_UpdateBits_name = map[int32]string{
		1: "AUTO_LIQ_THRESHOLD_CURRENT_VALUE",
	}
	RequestAccountRmsUpdates_UpdateBits_value = map[string]int32{
		"AUTO_LIQ_THRESHOLD_CURRENT_VALUE": 1,
	}
)

func (x RequestAccountRmsUpdates_UpdateBits) Enum() *RequestAccountRmsUpdates_UpdateBits {
	p := new(RequestAccountRmsUpdates_UpdateBits)
	*p = x
	return p
}

func (x RequestAccountRmsUpdates_UpdateBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestAccountRmsUpdates_UpdateBits) Descriptor() protoreflect.EnumDescriptor {
	return file_request_account_rms_updates_proto_enumTypes[0].Descriptor()
}

func (RequestAccountRmsUpdates_UpdateBits) Type() protoreflect.EnumType {
	return &file_request_account_rms_updates_proto_enumTypes[0]
}

func (x RequestAccountRmsUpdates_UpdateBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestAccountRmsUpdates_UpdateBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestAccountRmsUpdates_UpdateBits(num)
	return nil
}

// Deprecated: Use RequestAccountRmsUpdates_UpdateBits.Descriptor instead.
func (RequestAccountRmsUpdates_UpdateBits) EnumDescriptor() ([]byte, []int) {
	return file_request_account_rms_updates_proto_rawDescGZIP(), []int{0, 0}
}

type RequestAccountRmsUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	UserMsg    []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	FcmId      *string  `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`
	IbId       *string  `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`
	AccountId  *string  `protobuf:"bytes,154008,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Request    *string  `protobuf:"bytes,100000,opt,name=request" json:"request,omitempty"` // values can be either 'subscribe' or 'unsubscribe'
	UpdateBits *int32   `protobuf:"varint,154211,opt,name=update_bits,json=updateBits" json:"update_bits,omitempty"`
}

func (x *RequestAccountRmsUpdates) Reset() {
	*x = RequestAccountRmsUpdates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_account_rms_updates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAccountRmsUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAccountRmsUpdates) ProtoMessage() {}

func (x *RequestAccountRmsUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_request_account_rms_updates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAccountRmsUpdates.ProtoReflect.Descriptor instead.
func (*RequestAccountRmsUpdates) Descriptor() ([]byte, []int) {
	return file_request_account_rms_updates_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAccountRmsUpdates) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestAccountRmsUpdates) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestAccountRmsUpdates) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *RequestAccountRmsUpdates) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *RequestAccountRmsUpdates) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *RequestAccountRmsUpdates) GetRequest() string {
	if x != nil && x.Request != nil {
		return *x.Request
	}
	return ""
}

func (x *RequestAccountRmsUpdates) GetUpdateBits() int32 {
	if x != nil && x.UpdateBits != nil {
		return *x.UpdateBits
	}
	return 0
}

var File_request_account_rms_updates_proto protoreflect.FileDescriptor

var file_request_account_rms_updates_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x6d, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0x9e, 0x02, 0x0a, 0x18, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x6d, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x9d, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x05, 0x69, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x62, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x98, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0xa0, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x69, 0x74,
	0x73, 0x18, 0xe3, 0xb4, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x69, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x4c, 0x49, 0x51, 0x5f,
	0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e,
	0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f,
	0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_request_account_rms_updates_proto_rawDescOnce sync.Once
	file_request_account_rms_updates_proto_rawDescData = file_request_account_rms_updates_proto_rawDesc
)

func file_request_account_rms_updates_proto_rawDescGZIP() []byte {
	file_request_account_rms_updates_proto_rawDescOnce.Do(func() {
		file_request_account_rms_updates_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_account_rms_updates_proto_rawDescData)
	})
	return file_request_account_rms_updates_proto_rawDescData
}

var file_request_account_rms_updates_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_request_account_rms_updates_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_account_rms_updates_proto_goTypes = []interface{}{
	(RequestAccountRmsUpdates_UpdateBits)(0), // 0: rti.RequestAccountRmsUpdates.UpdateBits
	(*RequestAccountRmsUpdates)(nil),         // 1: rti.RequestAccountRmsUpdates
}
var file_request_account_rms_updates_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_account_rms_updates_proto_init() }
func file_request_account_rms_updates_proto_init() {
	if File_request_account_rms_updates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_account_rms_updates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAccountRmsUpdates); i {
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
			RawDescriptor: file_request_account_rms_updates_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_account_rms_updates_proto_goTypes,
		DependencyIndexes: file_request_account_rms_updates_proto_depIdxs,
		EnumInfos:         file_request_account_rms_updates_proto_enumTypes,
		MessageInfos:      file_request_account_rms_updates_proto_msgTypes,
	}.Build()
	File_request_account_rms_updates_proto = out.File
	file_request_account_rms_updates_proto_rawDesc = nil
	file_request_account_rms_updates_proto_goTypes = nil
	file_request_account_rms_updates_proto_depIdxs = nil
}
