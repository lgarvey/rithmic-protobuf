// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: request_market_data_update.proto

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
type RequestMarketDataUpdate_UpdateBits int32

const (
	RequestMarketDataUpdate_LAST_TRADE           RequestMarketDataUpdate_UpdateBits = 1
	RequestMarketDataUpdate_BBO                  RequestMarketDataUpdate_UpdateBits = 2
	RequestMarketDataUpdate_ORDER_BOOK           RequestMarketDataUpdate_UpdateBits = 4
	RequestMarketDataUpdate_OPEN                 RequestMarketDataUpdate_UpdateBits = 8
	RequestMarketDataUpdate_OPENING_INDICATOR    RequestMarketDataUpdate_UpdateBits = 16
	RequestMarketDataUpdate_HIGH_LOW             RequestMarketDataUpdate_UpdateBits = 32
	RequestMarketDataUpdate_HIGH_BID_LOW_ASK     RequestMarketDataUpdate_UpdateBits = 64
	RequestMarketDataUpdate_CLOSE                RequestMarketDataUpdate_UpdateBits = 128
	RequestMarketDataUpdate_CLOSING_INDICATOR    RequestMarketDataUpdate_UpdateBits = 256
	RequestMarketDataUpdate_SETTLEMENT           RequestMarketDataUpdate_UpdateBits = 512
	RequestMarketDataUpdate_MARKET_MODE          RequestMarketDataUpdate_UpdateBits = 1024
	RequestMarketDataUpdate_OPEN_INTEREST        RequestMarketDataUpdate_UpdateBits = 2048
	RequestMarketDataUpdate_MARGIN_RATE          RequestMarketDataUpdate_UpdateBits = 4096
	RequestMarketDataUpdate_HIGH_PRICE_LIMIT     RequestMarketDataUpdate_UpdateBits = 8192
	RequestMarketDataUpdate_LOW_PRICE_LIMIT      RequestMarketDataUpdate_UpdateBits = 16384
	RequestMarketDataUpdate_PROJECTED_SETTLEMENT RequestMarketDataUpdate_UpdateBits = 32768
	RequestMarketDataUpdate_ADJUSTED_CLOSE       RequestMarketDataUpdate_UpdateBits = 65536
)

// Enum value maps for RequestMarketDataUpdate_UpdateBits.
var (
	RequestMarketDataUpdate_UpdateBits_name = map[int32]string{
		1:     "LAST_TRADE",
		2:     "BBO",
		4:     "ORDER_BOOK",
		8:     "OPEN",
		16:    "OPENING_INDICATOR",
		32:    "HIGH_LOW",
		64:    "HIGH_BID_LOW_ASK",
		128:   "CLOSE",
		256:   "CLOSING_INDICATOR",
		512:   "SETTLEMENT",
		1024:  "MARKET_MODE",
		2048:  "OPEN_INTEREST",
		4096:  "MARGIN_RATE",
		8192:  "HIGH_PRICE_LIMIT",
		16384: "LOW_PRICE_LIMIT",
		32768: "PROJECTED_SETTLEMENT",
		65536: "ADJUSTED_CLOSE",
	}
	RequestMarketDataUpdate_UpdateBits_value = map[string]int32{
		"LAST_TRADE":           1,
		"BBO":                  2,
		"ORDER_BOOK":           4,
		"OPEN":                 8,
		"OPENING_INDICATOR":    16,
		"HIGH_LOW":             32,
		"HIGH_BID_LOW_ASK":     64,
		"CLOSE":                128,
		"CLOSING_INDICATOR":    256,
		"SETTLEMENT":           512,
		"MARKET_MODE":          1024,
		"OPEN_INTEREST":        2048,
		"MARGIN_RATE":          4096,
		"HIGH_PRICE_LIMIT":     8192,
		"LOW_PRICE_LIMIT":      16384,
		"PROJECTED_SETTLEMENT": 32768,
		"ADJUSTED_CLOSE":       65536,
	}
)

func (x RequestMarketDataUpdate_UpdateBits) Enum() *RequestMarketDataUpdate_UpdateBits {
	p := new(RequestMarketDataUpdate_UpdateBits)
	*p = x
	return p
}

func (x RequestMarketDataUpdate_UpdateBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestMarketDataUpdate_UpdateBits) Descriptor() protoreflect.EnumDescriptor {
	return file_request_market_data_update_proto_enumTypes[0].Descriptor()
}

func (RequestMarketDataUpdate_UpdateBits) Type() protoreflect.EnumType {
	return &file_request_market_data_update_proto_enumTypes[0]
}

func (x RequestMarketDataUpdate_UpdateBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestMarketDataUpdate_UpdateBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestMarketDataUpdate_UpdateBits(num)
	return nil
}

// Deprecated: Use RequestMarketDataUpdate_UpdateBits.Descriptor instead.
func (RequestMarketDataUpdate_UpdateBits) EnumDescriptor() ([]byte, []int) {
	return file_request_market_data_update_proto_rawDescGZIP(), []int{0, 0}
}

type RequestMarketDataUpdate_Request int32

const (
	RequestMarketDataUpdate_SUBSCRIBE   RequestMarketDataUpdate_Request = 1
	RequestMarketDataUpdate_UNSUBSCRIBE RequestMarketDataUpdate_Request = 2
)

// Enum value maps for RequestMarketDataUpdate_Request.
var (
	RequestMarketDataUpdate_Request_name = map[int32]string{
		1: "SUBSCRIBE",
		2: "UNSUBSCRIBE",
	}
	RequestMarketDataUpdate_Request_value = map[string]int32{
		"SUBSCRIBE":   1,
		"UNSUBSCRIBE": 2,
	}
)

func (x RequestMarketDataUpdate_Request) Enum() *RequestMarketDataUpdate_Request {
	p := new(RequestMarketDataUpdate_Request)
	*p = x
	return p
}

func (x RequestMarketDataUpdate_Request) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestMarketDataUpdate_Request) Descriptor() protoreflect.EnumDescriptor {
	return file_request_market_data_update_proto_enumTypes[1].Descriptor()
}

func (RequestMarketDataUpdate_Request) Type() protoreflect.EnumType {
	return &file_request_market_data_update_proto_enumTypes[1]
}

func (x RequestMarketDataUpdate_Request) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestMarketDataUpdate_Request) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestMarketDataUpdate_Request(num)
	return nil
}

// Deprecated: Use RequestMarketDataUpdate_Request.Descriptor instead.
func (RequestMarketDataUpdate_Request) EnumDescriptor() ([]byte, []int) {
	return file_request_market_data_update_proto_rawDescGZIP(), []int{0, 1}
}

type RequestMarketDataUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32                           `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	UserMsg    []string                         `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	Symbol     *string                          `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`
	Exchange   *string                          `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`
	Request    *RequestMarketDataUpdate_Request `protobuf:"varint,100000,opt,name=request,enum=rti.RequestMarketDataUpdate_Request" json:"request,omitempty"`
	UpdateBits *uint32                          `protobuf:"varint,154211,opt,name=update_bits,json=updateBits" json:"update_bits,omitempty"`
}

func (x *RequestMarketDataUpdate) Reset() {
	*x = RequestMarketDataUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_market_data_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMarketDataUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMarketDataUpdate) ProtoMessage() {}

func (x *RequestMarketDataUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_request_market_data_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMarketDataUpdate.ProtoReflect.Descriptor instead.
func (*RequestMarketDataUpdate) Descriptor() ([]byte, []int) {
	return file_request_market_data_update_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMarketDataUpdate) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestMarketDataUpdate) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestMarketDataUpdate) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *RequestMarketDataUpdate) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *RequestMarketDataUpdate) GetRequest() RequestMarketDataUpdate_Request {
	if x != nil && x.Request != nil {
		return *x.Request
	}
	return RequestMarketDataUpdate_SUBSCRIBE
}

func (x *RequestMarketDataUpdate) GetUpdateBits() uint32 {
	if x != nil && x.UpdateBits != nil {
		return *x.UpdateBits
	}
	return 0
}

var File_request_market_data_update_proto protoreflect.FileDescriptor

var file_request_market_data_update_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xeb, 0x04, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0xa0, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0xe3, 0xb4, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x74, 0x73,
	0x22, 0xc7, 0x02, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x74, 0x73, 0x12,
	0x0e, 0x0a, 0x0a, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x45, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x42, 0x42, 0x4f, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e,
	0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e,
	0x44, 0x49, 0x43, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x49, 0x47,
	0x48, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x20, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x49, 0x47, 0x48, 0x5f,
	0x42, 0x49, 0x44, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x53, 0x4b, 0x10, 0x40, 0x12, 0x0a, 0x0a,
	0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x80, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x43, 0x4c, 0x4f,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x44, 0x49, 0x43, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x80,
	0x02, 0x12, 0x0f, 0x0a, 0x0a, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x80, 0x04, 0x12, 0x10, 0x0a, 0x0b, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x10, 0x80, 0x08, 0x12, 0x12, 0x0a, 0x0d, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x80, 0x10, 0x12, 0x10, 0x0a, 0x0b, 0x4d, 0x41, 0x52, 0x47,
	0x49, 0x4e, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x80, 0x20, 0x12, 0x15, 0x0a, 0x10, 0x48, 0x49,
	0x47, 0x48, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x80,
	0x40, 0x12, 0x15, 0x0a, 0x0f, 0x4c, 0x4f, 0x57, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x4c,
	0x49, 0x4d, 0x49, 0x54, 0x10, 0x80, 0x80, 0x01, 0x12, 0x1a, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x4a,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x80, 0x80, 0x02, 0x12, 0x14, 0x0a, 0x0e, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x45, 0x44,
	0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x80, 0x80, 0x04, 0x22, 0x29, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49,
	0x42, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52,
	0x49, 0x42, 0x45, 0x10, 0x02, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_request_market_data_update_proto_rawDescOnce sync.Once
	file_request_market_data_update_proto_rawDescData = file_request_market_data_update_proto_rawDesc
)

func file_request_market_data_update_proto_rawDescGZIP() []byte {
	file_request_market_data_update_proto_rawDescOnce.Do(func() {
		file_request_market_data_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_market_data_update_proto_rawDescData)
	})
	return file_request_market_data_update_proto_rawDescData
}

var file_request_market_data_update_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_request_market_data_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_market_data_update_proto_goTypes = []interface{}{
	(RequestMarketDataUpdate_UpdateBits)(0), // 0: rti.RequestMarketDataUpdate.UpdateBits
	(RequestMarketDataUpdate_Request)(0),    // 1: rti.RequestMarketDataUpdate.Request
	(*RequestMarketDataUpdate)(nil),         // 2: rti.RequestMarketDataUpdate
}
var file_request_market_data_update_proto_depIdxs = []int32{
	1, // 0: rti.RequestMarketDataUpdate.request:type_name -> rti.RequestMarketDataUpdate.Request
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_request_market_data_update_proto_init() }
func file_request_market_data_update_proto_init() {
	if File_request_market_data_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_market_data_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMarketDataUpdate); i {
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
			RawDescriptor: file_request_market_data_update_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_market_data_update_proto_goTypes,
		DependencyIndexes: file_request_market_data_update_proto_depIdxs,
		EnumInfos:         file_request_market_data_update_proto_enumTypes,
		MessageInfos:      file_request_market_data_update_proto_msgTypes,
	}.Build()
	File_request_market_data_update_proto = out.File
	file_request_market_data_update_proto_rawDesc = nil
	file_request_market_data_update_proto_goTypes = nil
	file_request_market_data_update_proto_depIdxs = nil
}
