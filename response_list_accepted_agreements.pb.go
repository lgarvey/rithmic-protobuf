// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: response_list_accepted_agreements.proto

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

type ResponseListAcceptedAgreements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId                 *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	UserMsg                    []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	RqHandlerRpCode            []string `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"`
	RpCode                     []string `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`
	FcmId                      *string  `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`
	IbId                       *string  `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`
	AgreementAcceptanceSsboe   *int32   `protobuf:"varint,153427,opt,name=agreement_acceptance_ssboe,json=agreementAcceptanceSsboe" json:"agreement_acceptance_ssboe,omitempty"`
	AgreementAcceptanceStatus  *string  `protobuf:"bytes,153426,opt,name=agreement_acceptance_status,json=agreementAcceptanceStatus" json:"agreement_acceptance_status,omitempty"`
	AgreementAcceptanceRequest *string  `protobuf:"bytes,153430,opt,name=agreement_acceptance_request,json=agreementAcceptanceRequest" json:"agreement_acceptance_request,omitempty"`
	AgreementTitle             *string  `protobuf:"bytes,153406,opt,name=agreement_title,json=agreementTitle" json:"agreement_title,omitempty"`
	AgreementId                *string  `protobuf:"bytes,153407,opt,name=agreement_id,json=agreementId" json:"agreement_id,omitempty"`
}

func (x *ResponseListAcceptedAgreements) Reset() {
	*x = ResponseListAcceptedAgreements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_list_accepted_agreements_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseListAcceptedAgreements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseListAcceptedAgreements) ProtoMessage() {}

func (x *ResponseListAcceptedAgreements) ProtoReflect() protoreflect.Message {
	mi := &file_response_list_accepted_agreements_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseListAcceptedAgreements.ProtoReflect.Descriptor instead.
func (*ResponseListAcceptedAgreements) Descriptor() ([]byte, []int) {
	return file_response_list_accepted_agreements_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseListAcceptedAgreements) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseListAcceptedAgreements) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseListAcceptedAgreements) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseListAcceptedAgreements) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseListAcceptedAgreements) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *ResponseListAcceptedAgreements) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *ResponseListAcceptedAgreements) GetAgreementAcceptanceSsboe() int32 {
	if x != nil && x.AgreementAcceptanceSsboe != nil {
		return *x.AgreementAcceptanceSsboe
	}
	return 0
}

func (x *ResponseListAcceptedAgreements) GetAgreementAcceptanceStatus() string {
	if x != nil && x.AgreementAcceptanceStatus != nil {
		return *x.AgreementAcceptanceStatus
	}
	return ""
}

func (x *ResponseListAcceptedAgreements) GetAgreementAcceptanceRequest() string {
	if x != nil && x.AgreementAcceptanceRequest != nil {
		return *x.AgreementAcceptanceRequest
	}
	return ""
}

func (x *ResponseListAcceptedAgreements) GetAgreementTitle() string {
	if x != nil && x.AgreementTitle != nil {
		return *x.AgreementTitle
	}
	return ""
}

func (x *ResponseListAcceptedAgreements) GetAgreementId() string {
	if x != nil && x.AgreementId != nil {
		return *x.AgreementId
	}
	return ""
}

var File_response_list_accepted_agreements_proto protoreflect.FileDescriptor

var file_response_list_accepted_agreements_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xf0,
	0x03, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9e, 0x8d, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x06, 0x66,
	0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66,
	0x63, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x1a, 0x61,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xd3, 0xae, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x18, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x40, 0x0a, 0x1b, 0x61,
	0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0xd2, 0xae, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x19, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a,
	0x1c, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0xd6, 0xae,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x0f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0xbe, 0xae, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x67,
	0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0c,
	0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0xbf, 0xae, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f,
}

var (
	file_response_list_accepted_agreements_proto_rawDescOnce sync.Once
	file_response_list_accepted_agreements_proto_rawDescData = file_response_list_accepted_agreements_proto_rawDesc
)

func file_response_list_accepted_agreements_proto_rawDescGZIP() []byte {
	file_response_list_accepted_agreements_proto_rawDescOnce.Do(func() {
		file_response_list_accepted_agreements_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_list_accepted_agreements_proto_rawDescData)
	})
	return file_response_list_accepted_agreements_proto_rawDescData
}

var file_response_list_accepted_agreements_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_list_accepted_agreements_proto_goTypes = []interface{}{
	(*ResponseListAcceptedAgreements)(nil), // 0: rti.ResponseListAcceptedAgreements
}
var file_response_list_accepted_agreements_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_list_accepted_agreements_proto_init() }
func file_response_list_accepted_agreements_proto_init() {
	if File_response_list_accepted_agreements_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_list_accepted_agreements_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseListAcceptedAgreements); i {
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
			RawDescriptor: file_response_list_accepted_agreements_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_list_accepted_agreements_proto_goTypes,
		DependencyIndexes: file_response_list_accepted_agreements_proto_depIdxs,
		MessageInfos:      file_response_list_accepted_agreements_proto_msgTypes,
	}.Build()
	File_response_list_accepted_agreements_proto = out.File
	file_response_list_accepted_agreements_proto_rawDesc = nil
	file_response_list_accepted_agreements_proto_goTypes = nil
	file_response_list_accepted_agreements_proto_depIdxs = nil
}
