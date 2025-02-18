// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: update_easy_to_borrow_list.proto

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

type UpdateEasyToBorrowList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId   *int32  `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`       // PB_OFFSET + MNM_TEMPLATE_ID
	BrokerDealer *string `protobuf:"bytes,154612,opt,name=broker_dealer,json=brokerDealer" json:"broker_dealer,omitempty"`  // PB_OFFSET + MNM_BROKER_DEALER
	Symbol       *string `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                  // PB_OFFSET + MNM_SYMBOL
	SymbolName   *string `protobuf:"bytes,100003,opt,name=symbol_name,json=symbolName" json:"symbol_name,omitempty"`        // PB_OFFSET + MNM_SYMBOL_NAME
	QtyAvailable *int32  `protobuf:"varint,154613,opt,name=qty_available,json=qtyAvailable" json:"qty_available,omitempty"` // PB_OFFSET + MNM_TOTAL_AVAILABLE_QTY
	QtyNeeded    *int32  `protobuf:"varint,154614,opt,name=qty_needed,json=qtyNeeded" json:"qty_needed,omitempty"`          // PB_OFFSET + MNM_TOTAL_USED_QTY
	Borrowable   *bool   `protobuf:"varint,110353,opt,name=borrowable" json:"borrowable,omitempty"`                         // PB_OFFSET + MNM_SHORT_LIST_INDICATOR
}

func (x *UpdateEasyToBorrowList) Reset() {
	*x = UpdateEasyToBorrowList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_update_easy_to_borrow_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEasyToBorrowList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEasyToBorrowList) ProtoMessage() {}

func (x *UpdateEasyToBorrowList) ProtoReflect() protoreflect.Message {
	mi := &file_update_easy_to_borrow_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEasyToBorrowList.ProtoReflect.Descriptor instead.
func (*UpdateEasyToBorrowList) Descriptor() ([]byte, []int) {
	return file_update_easy_to_borrow_list_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateEasyToBorrowList) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *UpdateEasyToBorrowList) GetBrokerDealer() string {
	if x != nil && x.BrokerDealer != nil {
		return *x.BrokerDealer
	}
	return ""
}

func (x *UpdateEasyToBorrowList) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *UpdateEasyToBorrowList) GetSymbolName() string {
	if x != nil && x.SymbolName != nil {
		return *x.SymbolName
	}
	return ""
}

func (x *UpdateEasyToBorrowList) GetQtyAvailable() int32 {
	if x != nil && x.QtyAvailable != nil {
		return *x.QtyAvailable
	}
	return 0
}

func (x *UpdateEasyToBorrowList) GetQtyNeeded() int32 {
	if x != nil && x.QtyNeeded != nil {
		return *x.QtyNeeded
	}
	return 0
}

func (x *UpdateEasyToBorrowList) GetBorrowable() bool {
	if x != nil && x.Borrowable != nil {
		return *x.Borrowable
	}
	return false
}

var File_update_easy_to_borrow_list_proto protoreflect.FileDescriptor

var file_update_easy_to_borrow_list_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x61, 0x73, 0x79, 0x5f, 0x74, 0x6f,
	0x5f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0x89, 0x02, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x61, 0x73, 0x79, 0x54, 0x6f, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x18, 0xf4, 0xb7, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xa3, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x71, 0x74, 0x79,
	0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0xf5, 0xb7, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x71, 0x74, 0x79, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1f, 0x0a, 0x0a, 0x71, 0x74, 0x79, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0xf6,
	0xb7, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x71, 0x74, 0x79, 0x4e, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x12, 0x20, 0x0a, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x91, 0xde, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f,
}

var (
	file_update_easy_to_borrow_list_proto_rawDescOnce sync.Once
	file_update_easy_to_borrow_list_proto_rawDescData = file_update_easy_to_borrow_list_proto_rawDesc
)

func file_update_easy_to_borrow_list_proto_rawDescGZIP() []byte {
	file_update_easy_to_borrow_list_proto_rawDescOnce.Do(func() {
		file_update_easy_to_borrow_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_update_easy_to_borrow_list_proto_rawDescData)
	})
	return file_update_easy_to_borrow_list_proto_rawDescData
}

var file_update_easy_to_borrow_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_update_easy_to_borrow_list_proto_goTypes = []interface{}{
	(*UpdateEasyToBorrowList)(nil), // 0: rti.UpdateEasyToBorrowList
}
var file_update_easy_to_borrow_list_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_update_easy_to_borrow_list_proto_init() }
func file_update_easy_to_borrow_list_proto_init() {
	if File_update_easy_to_borrow_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_update_easy_to_borrow_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEasyToBorrowList); i {
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
			RawDescriptor: file_update_easy_to_borrow_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_update_easy_to_borrow_list_proto_goTypes,
		DependencyIndexes: file_update_easy_to_borrow_list_proto_depIdxs,
		MessageInfos:      file_update_easy_to_borrow_list_proto_msgTypes,
	}.Build()
	File_update_easy_to_borrow_list_proto = out.File
	file_update_easy_to_borrow_list_proto_rawDesc = nil
	file_update_easy_to_borrow_list_proto_goTypes = nil
	file_update_easy_to_borrow_list_proto_depIdxs = nil
}
