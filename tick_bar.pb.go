// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: tick_bar.proto

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

type TickBar_BarType int32

const (
	TickBar_TICK_BAR   TickBar_BarType = 1
	TickBar_RANGE_BAR  TickBar_BarType = 2
	TickBar_VOLUME_BAR TickBar_BarType = 3
)

// Enum value maps for TickBar_BarType.
var (
	TickBar_BarType_name = map[int32]string{
		1: "TICK_BAR",
		2: "RANGE_BAR",
		3: "VOLUME_BAR",
	}
	TickBar_BarType_value = map[string]int32{
		"TICK_BAR":   1,
		"RANGE_BAR":  2,
		"VOLUME_BAR": 3,
	}
)

func (x TickBar_BarType) Enum() *TickBar_BarType {
	p := new(TickBar_BarType)
	*p = x
	return p
}

func (x TickBar_BarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TickBar_BarType) Descriptor() protoreflect.EnumDescriptor {
	return file_tick_bar_proto_enumTypes[0].Descriptor()
}

func (TickBar_BarType) Type() protoreflect.EnumType {
	return &file_tick_bar_proto_enumTypes[0]
}

func (x TickBar_BarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TickBar_BarType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TickBar_BarType(num)
	return nil
}

// Deprecated: Use TickBar_BarType.Descriptor instead.
func (TickBar_BarType) EnumDescriptor() ([]byte, []int) {
	return file_tick_bar_proto_rawDescGZIP(), []int{0, 0}
}

type TickBar_BarSubType int32

const (
	TickBar_REGULAR TickBar_BarSubType = 1
	TickBar_CUSTOM  TickBar_BarSubType = 2
)

// Enum value maps for TickBar_BarSubType.
var (
	TickBar_BarSubType_name = map[int32]string{
		1: "REGULAR",
		2: "CUSTOM",
	}
	TickBar_BarSubType_value = map[string]int32{
		"REGULAR": 1,
		"CUSTOM":  2,
	}
)

func (x TickBar_BarSubType) Enum() *TickBar_BarSubType {
	p := new(TickBar_BarSubType)
	*p = x
	return p
}

func (x TickBar_BarSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TickBar_BarSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_tick_bar_proto_enumTypes[1].Descriptor()
}

func (TickBar_BarSubType) Type() protoreflect.EnumType {
	return &file_tick_bar_proto_enumTypes[1]
}

func (x TickBar_BarSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TickBar_BarSubType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TickBar_BarSubType(num)
	return nil
}

// Deprecated: Use TickBar_BarSubType.Descriptor instead.
func (TickBar_BarSubType) EnumDescriptor() ([]byte, []int) {
	return file_tick_bar_proto_rawDescGZIP(), []int{0, 1}
}

type TickBar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId           *int32              `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                                   // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol               *string             `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                                              // PB_OFFSET + MNM_SYMBOL
	Exchange             *string             `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                                          // PB_OFFSET + MNM_EXCHANGE
	Type                 *TickBar_BarType    `protobuf:"varint,119200,opt,name=type,enum=rti.TickBar_BarType" json:"type,omitempty"`                                        // PB_OFFSET + MNM_DATA_BAR_TYPE
	SubType              *TickBar_BarSubType `protobuf:"varint,119208,opt,name=sub_type,json=subType,enum=rti.TickBar_BarSubType" json:"sub_type,omitempty"`                // PB_OFFSET + MNM_DATA_BAR_SUB_TYPE
	TypeSpecifier        *string             `protobuf:"bytes,148162,opt,name=type_specifier,json=typeSpecifier" json:"type_specifier,omitempty"`                           // PB_OFFSET + MNM_CATEGORY_SPECIFIC_INFO
	NumTrades            *uint64             `protobuf:"varint,119204,opt,name=num_trades,json=numTrades" json:"num_trades,omitempty"`                                      // PB_OFFSET + MNM_DATA_BAR_NUM_TRADES
	Volume               *uint64             `protobuf:"varint,119205,opt,name=volume" json:"volume,omitempty"`                                                             // PB_OFFSET + MNM_DATA_BAR_TRADE_VOLUME
	BidVolume            *uint64             `protobuf:"varint,119213,opt,name=bid_volume,json=bidVolume" json:"bid_volume,omitempty"`                                      // PB_OFFSET + MNM_DATA_BAR_BID_VOLUME
	AskVolume            *uint64             `protobuf:"varint,119214,opt,name=ask_volume,json=askVolume" json:"ask_volume,omitempty"`                                      // PB_OFFSET + MNM_DATA_BAR_ASK_VOLUME
	OpenPrice            *float64            `protobuf:"fixed64,100019,opt,name=open_price,json=openPrice" json:"open_price,omitempty"`                                     // PB_OFFSET + MNM_OPEN_PRICE
	ClosePrice           *float64            `protobuf:"fixed64,100021,opt,name=close_price,json=closePrice" json:"close_price,omitempty"`                                  // PB_OFFSET + MNM_CLOSE_TRADE_PRICE
	HighPrice            *float64            `protobuf:"fixed64,100012,opt,name=high_price,json=highPrice" json:"high_price,omitempty"`                                     // PB_OFFSET + MNM_HIGH_PRICE
	LowPrice             *float64            `protobuf:"fixed64,100013,opt,name=low_price,json=lowPrice" json:"low_price,omitempty"`                                        // PB_OFFSET + MNM_LOW_PRICE
	CustomSessionOpenSsm *int32              `protobuf:"varint,119209,opt,name=custom_session_open_ssm,json=customSessionOpenSsm" json:"custom_session_open_ssm,omitempty"` // PB_OFFSET + MNM_CUSTOM_SESSION_OPEN_SSM
	DataBarSsboe         []int32             `protobuf:"varint,119202,rep,name=data_bar_ssboe,json=dataBarSsboe" json:"data_bar_ssboe,omitempty"`                           // PB_OFFSET + MNM_DATA_BAR_SSBOE
	DataBarUsecs         []int32             `protobuf:"varint,119203,rep,name=data_bar_usecs,json=dataBarUsecs" json:"data_bar_usecs,omitempty"`                           // PB_OFFSTE + MNM_DATA_BAR_USECS
}

func (x *TickBar) Reset() {
	*x = TickBar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tick_bar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickBar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickBar) ProtoMessage() {}

func (x *TickBar) ProtoReflect() protoreflect.Message {
	mi := &file_tick_bar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickBar.ProtoReflect.Descriptor instead.
func (*TickBar) Descriptor() ([]byte, []int) {
	return file_tick_bar_proto_rawDescGZIP(), []int{0}
}

func (x *TickBar) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *TickBar) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *TickBar) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *TickBar) GetType() TickBar_BarType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return TickBar_TICK_BAR
}

func (x *TickBar) GetSubType() TickBar_BarSubType {
	if x != nil && x.SubType != nil {
		return *x.SubType
	}
	return TickBar_REGULAR
}

func (x *TickBar) GetTypeSpecifier() string {
	if x != nil && x.TypeSpecifier != nil {
		return *x.TypeSpecifier
	}
	return ""
}

func (x *TickBar) GetNumTrades() uint64 {
	if x != nil && x.NumTrades != nil {
		return *x.NumTrades
	}
	return 0
}

func (x *TickBar) GetVolume() uint64 {
	if x != nil && x.Volume != nil {
		return *x.Volume
	}
	return 0
}

func (x *TickBar) GetBidVolume() uint64 {
	if x != nil && x.BidVolume != nil {
		return *x.BidVolume
	}
	return 0
}

func (x *TickBar) GetAskVolume() uint64 {
	if x != nil && x.AskVolume != nil {
		return *x.AskVolume
	}
	return 0
}

func (x *TickBar) GetOpenPrice() float64 {
	if x != nil && x.OpenPrice != nil {
		return *x.OpenPrice
	}
	return 0
}

func (x *TickBar) GetClosePrice() float64 {
	if x != nil && x.ClosePrice != nil {
		return *x.ClosePrice
	}
	return 0
}

func (x *TickBar) GetHighPrice() float64 {
	if x != nil && x.HighPrice != nil {
		return *x.HighPrice
	}
	return 0
}

func (x *TickBar) GetLowPrice() float64 {
	if x != nil && x.LowPrice != nil {
		return *x.LowPrice
	}
	return 0
}

func (x *TickBar) GetCustomSessionOpenSsm() int32 {
	if x != nil && x.CustomSessionOpenSsm != nil {
		return *x.CustomSessionOpenSsm
	}
	return 0
}

func (x *TickBar) GetDataBarSsboe() []int32 {
	if x != nil {
		return x.DataBarSsboe
	}
	return nil
}

func (x *TickBar) GetDataBarUsecs() []int32 {
	if x != nil {
		return x.DataBarUsecs
	}
	return nil
}

var File_tick_bar_proto protoreflect.FileDescriptor

var file_tick_bar_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xd8, 0x05, 0x0a, 0x07, 0x54, 0x69, 0x63, 0x6b, 0x42, 0x61,
	0x72, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94,
	0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c,
	0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0xa0, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72,
	0x74, 0x69, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x72, 0x2e, 0x42, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0xa8, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72,
	0x74, 0x69, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x72, 0x2e, 0x42, 0x61, 0x72, 0x53, 0x75,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27,
	0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0xc2, 0x85, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0xa4, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e,
	0x75, 0x6d, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0xa5, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0xad, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x69, 0x64, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0xae, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x73, 0x6b, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0xb3, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0xb5, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xac, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6c, 0x6f, 0x77,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xad, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x17, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f,
	0x73, 0x73, 0x6d, 0x18, 0xa9, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x73,
	0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x61, 0x72, 0x5f, 0x73, 0x73,
	0x62, 0x6f, 0x65, 0x18, 0xa2, 0xa3, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x61, 0x72, 0x53, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x62, 0x61, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xa3, 0xa3, 0x07, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x42, 0x61, 0x72, 0x55, 0x73, 0x65, 0x63,
	0x73, 0x22, 0x36, 0x0a, 0x07, 0x42, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x54, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x4f, 0x4c,
	0x55, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x03, 0x22, 0x25, 0x0a, 0x0a, 0x42, 0x61, 0x72,
	0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x47, 0x55, 0x4c,
	0x41, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x02,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_tick_bar_proto_rawDescOnce sync.Once
	file_tick_bar_proto_rawDescData = file_tick_bar_proto_rawDesc
)

func file_tick_bar_proto_rawDescGZIP() []byte {
	file_tick_bar_proto_rawDescOnce.Do(func() {
		file_tick_bar_proto_rawDescData = protoimpl.X.CompressGZIP(file_tick_bar_proto_rawDescData)
	})
	return file_tick_bar_proto_rawDescData
}

var file_tick_bar_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tick_bar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tick_bar_proto_goTypes = []interface{}{
	(TickBar_BarType)(0),    // 0: rti.TickBar.BarType
	(TickBar_BarSubType)(0), // 1: rti.TickBar.BarSubType
	(*TickBar)(nil),         // 2: rti.TickBar
}
var file_tick_bar_proto_depIdxs = []int32{
	0, // 0: rti.TickBar.type:type_name -> rti.TickBar.BarType
	1, // 1: rti.TickBar.sub_type:type_name -> rti.TickBar.BarSubType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tick_bar_proto_init() }
func file_tick_bar_proto_init() {
	if File_tick_bar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tick_bar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickBar); i {
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
			RawDescriptor: file_tick_bar_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tick_bar_proto_goTypes,
		DependencyIndexes: file_tick_bar_proto_depIdxs,
		EnumInfos:         file_tick_bar_proto_enumTypes,
		MessageInfos:      file_tick_bar_proto_msgTypes,
	}.Build()
	File_tick_bar_proto = out.File
	file_tick_bar_proto_rawDesc = nil
	file_tick_bar_proto_goTypes = nil
	file_tick_bar_proto_depIdxs = nil
}
