// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: response_give_tick_size_type_table.proto

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

// bit constants are defined using enum
type ResponseGiveTickSizeTypeTable_PresenceBits int32

const (
	ResponseGiveTickSizeTypeTable_TICK_SIZE_FIRST_PRICE ResponseGiveTickSizeTypeTable_PresenceBits = 1
	ResponseGiveTickSizeTypeTable_TICK_SIZE_LAST_PRICE  ResponseGiveTickSizeTypeTable_PresenceBits = 2
	ResponseGiveTickSizeTypeTable_TICK_SIZE_FP_OPERATOR ResponseGiveTickSizeTypeTable_PresenceBits = 4
	ResponseGiveTickSizeTypeTable_TICK_SIZE_LP_OPERATOR ResponseGiveTickSizeTypeTable_PresenceBits = 8
)

// Enum value maps for ResponseGiveTickSizeTypeTable_PresenceBits.
var (
	ResponseGiveTickSizeTypeTable_PresenceBits_name = map[int32]string{
		1: "TICK_SIZE_FIRST_PRICE",
		2: "TICK_SIZE_LAST_PRICE",
		4: "TICK_SIZE_FP_OPERATOR",
		8: "TICK_SIZE_LP_OPERATOR",
	}
	ResponseGiveTickSizeTypeTable_PresenceBits_value = map[string]int32{
		"TICK_SIZE_FIRST_PRICE": 1,
		"TICK_SIZE_LAST_PRICE":  2,
		"TICK_SIZE_FP_OPERATOR": 4,
		"TICK_SIZE_LP_OPERATOR": 8,
	}
)

func (x ResponseGiveTickSizeTypeTable_PresenceBits) Enum() *ResponseGiveTickSizeTypeTable_PresenceBits {
	p := new(ResponseGiveTickSizeTypeTable_PresenceBits)
	*p = x
	return p
}

func (x ResponseGiveTickSizeTypeTable_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseGiveTickSizeTypeTable_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_response_give_tick_size_type_table_proto_enumTypes[0].Descriptor()
}

func (ResponseGiveTickSizeTypeTable_PresenceBits) Type() protoreflect.EnumType {
	return &file_response_give_tick_size_type_table_proto_enumTypes[0]
}

func (x ResponseGiveTickSizeTypeTable_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResponseGiveTickSizeTypeTable_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResponseGiveTickSizeTypeTable_PresenceBits(num)
	return nil
}

// Deprecated: Use ResponseGiveTickSizeTypeTable_PresenceBits.Descriptor instead.
func (ResponseGiveTickSizeTypeTable_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_response_give_tick_size_type_table_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseGiveTickSizeTypeTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId         *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                              // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg            []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                        // PB_OFFSET + MNM_USER_MSG
	RqHandlerRpCode    []string `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"`            // PB_OFFSET + MNM_REQUEST_HANDLER_RESPONSE_CODE
	RpCode             []string `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                           // PB_OFFSET + MNM_RESPONSE_CODE
	PresenceBits       *uint32  `protobuf:"varint,149138,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"`                        // PB_OFFSET + MNM_PRICING_INDICATOR
	TickSizeType       *string  `protobuf:"bytes,154167,opt,name=tick_size_type,json=tickSizeType" json:"tick_size_type,omitempty"`                       // PB_OFFSET + MNM_TICK_SIZE_TYPE
	TickSizeFpOperator *string  `protobuf:"bytes,154170,opt,name=tick_size_fp_operator,json=tickSizeFpOperator" json:"tick_size_fp_operator,omitempty"`   // PB_OFFSET + MNM_TICK_SIZE_FP_OPERATOR
	TickSizeLpOperator *string  `protobuf:"bytes,154171,opt,name=tick_size_lp_operator,json=tickSizeLpOperator" json:"tick_size_lp_operator,omitempty"`   // PB_OFFSET + MNM_TICK_SIZE_LP_OPERATOR
	MinFpriceChange    *float64 `protobuf:"fixed64,154387,opt,name=min_fprice_change,json=minFpriceChange" json:"min_fprice_change,omitempty"`            // PB_OFFSET + MNM_MIN_FPRICE_CHANGE
	TickSizeFirstPrice *float64 `protobuf:"fixed64,154168,opt,name=tick_size_first_price,json=tickSizeFirstPrice" json:"tick_size_first_price,omitempty"` // PB_OFFSET + MNM_TICK_SIZE_FIRST_PRICE
	TickSizeLastPrice  *float64 `protobuf:"fixed64,154169,opt,name=tick_size_last_price,json=tickSizeLastPrice" json:"tick_size_last_price,omitempty"`    // PB_OFFSET + MNM_TICK_SIZE_LAST_PRICE
}

func (x *ResponseGiveTickSizeTypeTable) Reset() {
	*x = ResponseGiveTickSizeTypeTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_give_tick_size_type_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGiveTickSizeTypeTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGiveTickSizeTypeTable) ProtoMessage() {}

func (x *ResponseGiveTickSizeTypeTable) ProtoReflect() protoreflect.Message {
	mi := &file_response_give_tick_size_type_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGiveTickSizeTypeTable.ProtoReflect.Descriptor instead.
func (*ResponseGiveTickSizeTypeTable) Descriptor() ([]byte, []int) {
	return file_response_give_tick_size_type_table_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseGiveTickSizeTypeTable) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseGiveTickSizeTypeTable) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseGiveTickSizeTypeTable) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseGiveTickSizeTypeTable) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseGiveTickSizeTypeTable) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *ResponseGiveTickSizeTypeTable) GetTickSizeType() string {
	if x != nil && x.TickSizeType != nil {
		return *x.TickSizeType
	}
	return ""
}

func (x *ResponseGiveTickSizeTypeTable) GetTickSizeFpOperator() string {
	if x != nil && x.TickSizeFpOperator != nil {
		return *x.TickSizeFpOperator
	}
	return ""
}

func (x *ResponseGiveTickSizeTypeTable) GetTickSizeLpOperator() string {
	if x != nil && x.TickSizeLpOperator != nil {
		return *x.TickSizeLpOperator
	}
	return ""
}

func (x *ResponseGiveTickSizeTypeTable) GetMinFpriceChange() float64 {
	if x != nil && x.MinFpriceChange != nil {
		return *x.MinFpriceChange
	}
	return 0
}

func (x *ResponseGiveTickSizeTypeTable) GetTickSizeFirstPrice() float64 {
	if x != nil && x.TickSizeFirstPrice != nil {
		return *x.TickSizeFirstPrice
	}
	return 0
}

func (x *ResponseGiveTickSizeTypeTable) GetTickSizeLastPrice() float64 {
	if x != nil && x.TickSizeLastPrice != nil {
		return *x.TickSizeLastPrice
	}
	return 0
}

var File_response_give_tick_size_type_table_proto protoreflect.FileDescriptor

var file_response_give_tick_size_type_table_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x5f,
	0x74, 0x69, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22,
	0xf3, 0x04, 0x0a, 0x1d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x69, 0x76, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x54, 0x79, 0x70, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9e, 0x8d, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x92, 0x8d, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x69,
	0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0xb7, 0xb4, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69,
	0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x74, 0x69,
	0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x66, 0x70, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0xba, 0xb4, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x69, 0x63,
	0x6b, 0x53, 0x69, 0x7a, 0x65, 0x46, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x33, 0x0a, 0x15, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x70, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0xbb, 0xb4, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x74, 0x69, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x70, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x5f, 0x66, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x93, 0xb6, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x46, 0x70, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xb8, 0xb4, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x12, 0x74, 0x69, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x74, 0x69, 0x63, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0xb9, 0xb4, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x69, 0x63, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x79, 0x0a, 0x0c, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x49,
	0x43, 0x4b, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x50, 0x52,
	0x49, 0x43, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x49, 0x43, 0x4b, 0x5f, 0x53, 0x49,
	0x5a, 0x45, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12,
	0x19, 0x0a, 0x15, 0x54, 0x49, 0x43, 0x4b, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x46, 0x50, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x49,
	0x43, 0x4b, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4c, 0x50, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x4f, 0x52, 0x10, 0x08, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_response_give_tick_size_type_table_proto_rawDescOnce sync.Once
	file_response_give_tick_size_type_table_proto_rawDescData = file_response_give_tick_size_type_table_proto_rawDesc
)

func file_response_give_tick_size_type_table_proto_rawDescGZIP() []byte {
	file_response_give_tick_size_type_table_proto_rawDescOnce.Do(func() {
		file_response_give_tick_size_type_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_give_tick_size_type_table_proto_rawDescData)
	})
	return file_response_give_tick_size_type_table_proto_rawDescData
}

var file_response_give_tick_size_type_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_give_tick_size_type_table_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_give_tick_size_type_table_proto_goTypes = []interface{}{
	(ResponseGiveTickSizeTypeTable_PresenceBits)(0), // 0: rti.ResponseGiveTickSizeTypeTable.PresenceBits
	(*ResponseGiveTickSizeTypeTable)(nil),           // 1: rti.ResponseGiveTickSizeTypeTable
}
var file_response_give_tick_size_type_table_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_give_tick_size_type_table_proto_init() }
func file_response_give_tick_size_type_table_proto_init() {
	if File_response_give_tick_size_type_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_give_tick_size_type_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGiveTickSizeTypeTable); i {
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
			RawDescriptor: file_response_give_tick_size_type_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_give_tick_size_type_table_proto_goTypes,
		DependencyIndexes: file_response_give_tick_size_type_table_proto_depIdxs,
		EnumInfos:         file_response_give_tick_size_type_table_proto_enumTypes,
		MessageInfos:      file_response_give_tick_size_type_table_proto_msgTypes,
	}.Build()
	File_response_give_tick_size_type_table_proto = out.File
	file_response_give_tick_size_type_table_proto_rawDesc = nil
	file_response_give_tick_size_type_table_proto_goTypes = nil
	file_response_give_tick_size_type_table_proto_depIdxs = nil
}
