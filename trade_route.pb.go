// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: trade_route.proto

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

type TradeRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32  `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"` // PB_OFFSET + MNM_TEMPLATE_ID
	FcmId      *string `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`                 // PB_OFFSET + MNM_FCM_ID
	IbId       *string `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`                    // PB_OFFSET + MNM_IB_ID
	Exchange   *string `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                        // PB_OFFSET + MNM_EXCHANGE
	TradeRoute *string `protobuf:"bytes,112016,opt,name=trade_route,json=tradeRoute" json:"trade_route,omitempty"`  // PB_OFFSET + MNM_TRADE_ROUTE
	Status     *string `protobuf:"bytes,131407,opt,name=status" json:"status,omitempty"`                            // PB_OFFSET + MNM_SERVICE_STATE
	IsDefault  *bool   `protobuf:"varint,154689,opt,name=is_default,json=isDefault" json:"is_default,omitempty"`    // PB_OFFSET + MNM_DEFAULT_ROUTE
}

func (x *TradeRoute) Reset() {
	*x = TradeRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeRoute) ProtoMessage() {}

func (x *TradeRoute) ProtoReflect() protoreflect.Message {
	mi := &file_trade_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeRoute.ProtoReflect.Descriptor instead.
func (*TradeRoute) Descriptor() ([]byte, []int) {
	return file_trade_route_proto_rawDescGZIP(), []int{0}
}

func (x *TradeRoute) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *TradeRoute) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *TradeRoute) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *TradeRoute) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *TradeRoute) GetTradeRoute() string {
	if x != nil && x.TradeRoute != nil {
		return *x.TradeRoute
	}
	return ""
}

func (x *TradeRoute) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *TradeRoute) GetIsDefault() bool {
	if x != nil && x.IsDefault != nil {
		return *x.IsDefault
	}
	return false
}

var File_trade_route_proto protoreflect.FileDescriptor

var file_trade_route_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xdb, 0x01, 0x0a, 0x0a, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x06, 0x66, 0x63,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63,
	0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x90, 0xeb, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0xcf, 0x82, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0xc1, 0xb8, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_trade_route_proto_rawDescOnce sync.Once
	file_trade_route_proto_rawDescData = file_trade_route_proto_rawDesc
)

func file_trade_route_proto_rawDescGZIP() []byte {
	file_trade_route_proto_rawDescOnce.Do(func() {
		file_trade_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_route_proto_rawDescData)
	})
	return file_trade_route_proto_rawDescData
}

var file_trade_route_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trade_route_proto_goTypes = []interface{}{
	(*TradeRoute)(nil), // 0: rti.TradeRoute
}
var file_trade_route_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trade_route_proto_init() }
func file_trade_route_proto_init() {
	if File_trade_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trade_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeRoute); i {
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
			RawDescriptor: file_trade_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_route_proto_goTypes,
		DependencyIndexes: file_trade_route_proto_depIdxs,
		MessageInfos:      file_trade_route_proto_msgTypes,
	}.Build()
	File_trade_route_proto = out.File
	file_trade_route_proto_rawDesc = nil
	file_trade_route_proto_goTypes = nil
	file_trade_route_proto_depIdxs = nil
}
