// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: request_new_order.proto

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

type RequestNewOrder_TransactionType int32

const (
	RequestNewOrder_BUY  RequestNewOrder_TransactionType = 1
	RequestNewOrder_SELL RequestNewOrder_TransactionType = 2
)

// Enum value maps for RequestNewOrder_TransactionType.
var (
	RequestNewOrder_TransactionType_name = map[int32]string{
		1: "BUY",
		2: "SELL",
	}
	RequestNewOrder_TransactionType_value = map[string]int32{
		"BUY":  1,
		"SELL": 2,
	}
)

func (x RequestNewOrder_TransactionType) Enum() *RequestNewOrder_TransactionType {
	p := new(RequestNewOrder_TransactionType)
	*p = x
	return p
}

func (x RequestNewOrder_TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestNewOrder_TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_new_order_proto_enumTypes[0].Descriptor()
}

func (RequestNewOrder_TransactionType) Type() protoreflect.EnumType {
	return &file_request_new_order_proto_enumTypes[0]
}

func (x RequestNewOrder_TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestNewOrder_TransactionType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestNewOrder_TransactionType(num)
	return nil
}

// Deprecated: Use RequestNewOrder_TransactionType.Descriptor instead.
func (RequestNewOrder_TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_request_new_order_proto_rawDescGZIP(), []int{0, 0}
}

type RequestNewOrder_OrderPlacement int32

const (
	RequestNewOrder_MANUAL RequestNewOrder_OrderPlacement = 1
	RequestNewOrder_AUTO   RequestNewOrder_OrderPlacement = 2
)

// Enum value maps for RequestNewOrder_OrderPlacement.
var (
	RequestNewOrder_OrderPlacement_name = map[int32]string{
		1: "MANUAL",
		2: "AUTO",
	}
	RequestNewOrder_OrderPlacement_value = map[string]int32{
		"MANUAL": 1,
		"AUTO":   2,
	}
)

func (x RequestNewOrder_OrderPlacement) Enum() *RequestNewOrder_OrderPlacement {
	p := new(RequestNewOrder_OrderPlacement)
	*p = x
	return p
}

func (x RequestNewOrder_OrderPlacement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestNewOrder_OrderPlacement) Descriptor() protoreflect.EnumDescriptor {
	return file_request_new_order_proto_enumTypes[1].Descriptor()
}

func (RequestNewOrder_OrderPlacement) Type() protoreflect.EnumType {
	return &file_request_new_order_proto_enumTypes[1]
}

func (x RequestNewOrder_OrderPlacement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestNewOrder_OrderPlacement) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestNewOrder_OrderPlacement(num)
	return nil
}

// Deprecated: Use RequestNewOrder_OrderPlacement.Descriptor instead.
func (RequestNewOrder_OrderPlacement) EnumDescriptor() ([]byte, []int) {
	return file_request_new_order_proto_rawDescGZIP(), []int{0, 1}
}

type RequestNewOrder_Duration int32

const (
	RequestNewOrder_DAY RequestNewOrder_Duration = 1
	RequestNewOrder_GTC RequestNewOrder_Duration = 2
	RequestNewOrder_IOC RequestNewOrder_Duration = 3
	RequestNewOrder_FOK RequestNewOrder_Duration = 4
)

// Enum value maps for RequestNewOrder_Duration.
var (
	RequestNewOrder_Duration_name = map[int32]string{
		1: "DAY",
		2: "GTC",
		3: "IOC",
		4: "FOK",
	}
	RequestNewOrder_Duration_value = map[string]int32{
		"DAY": 1,
		"GTC": 2,
		"IOC": 3,
		"FOK": 4,
	}
)

func (x RequestNewOrder_Duration) Enum() *RequestNewOrder_Duration {
	p := new(RequestNewOrder_Duration)
	*p = x
	return p
}

func (x RequestNewOrder_Duration) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestNewOrder_Duration) Descriptor() protoreflect.EnumDescriptor {
	return file_request_new_order_proto_enumTypes[2].Descriptor()
}

func (RequestNewOrder_Duration) Type() protoreflect.EnumType {
	return &file_request_new_order_proto_enumTypes[2]
}

func (x RequestNewOrder_Duration) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestNewOrder_Duration) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestNewOrder_Duration(num)
	return nil
}

// Deprecated: Use RequestNewOrder_Duration.Descriptor instead.
func (RequestNewOrder_Duration) EnumDescriptor() ([]byte, []int) {
	return file_request_new_order_proto_rawDescGZIP(), []int{0, 2}
}

type RequestNewOrder_PriceType int32

const (
	RequestNewOrder_LIMIT             RequestNewOrder_PriceType = 1
	RequestNewOrder_MARKET            RequestNewOrder_PriceType = 2
	RequestNewOrder_STOP_LIMIT        RequestNewOrder_PriceType = 3
	RequestNewOrder_STOP_MARKET       RequestNewOrder_PriceType = 4
	RequestNewOrder_MARKET_IF_TOUCHED RequestNewOrder_PriceType = 5
	RequestNewOrder_LIMIT_IF_TOUCHED  RequestNewOrder_PriceType = 6
)

// Enum value maps for RequestNewOrder_PriceType.
var (
	RequestNewOrder_PriceType_name = map[int32]string{
		1: "LIMIT",
		2: "MARKET",
		3: "STOP_LIMIT",
		4: "STOP_MARKET",
		5: "MARKET_IF_TOUCHED",
		6: "LIMIT_IF_TOUCHED",
	}
	RequestNewOrder_PriceType_value = map[string]int32{
		"LIMIT":             1,
		"MARKET":            2,
		"STOP_LIMIT":        3,
		"STOP_MARKET":       4,
		"MARKET_IF_TOUCHED": 5,
		"LIMIT_IF_TOUCHED":  6,
	}
)

func (x RequestNewOrder_PriceType) Enum() *RequestNewOrder_PriceType {
	p := new(RequestNewOrder_PriceType)
	*p = x
	return p
}

func (x RequestNewOrder_PriceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestNewOrder_PriceType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_new_order_proto_enumTypes[3].Descriptor()
}

func (RequestNewOrder_PriceType) Type() protoreflect.EnumType {
	return &file_request_new_order_proto_enumTypes[3]
}

func (x RequestNewOrder_PriceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestNewOrder_PriceType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestNewOrder_PriceType(num)
	return nil
}

// Deprecated: Use RequestNewOrder_PriceType.Descriptor instead.
func (RequestNewOrder_PriceType) EnumDescriptor() ([]byte, []int) {
	return file_request_new_order_proto_rawDescGZIP(), []int{0, 3}
}

type RequestNewOrder_PriceField int32

const (
	RequestNewOrder_BID_PRICE   RequestNewOrder_PriceField = 1
	RequestNewOrder_OFFER_PRICE RequestNewOrder_PriceField = 2
	RequestNewOrder_TRADE_PRICE RequestNewOrder_PriceField = 3
	RequestNewOrder_LEAN_PRICE  RequestNewOrder_PriceField = 4
)

// Enum value maps for RequestNewOrder_PriceField.
var (
	RequestNewOrder_PriceField_name = map[int32]string{
		1: "BID_PRICE",
		2: "OFFER_PRICE",
		3: "TRADE_PRICE",
		4: "LEAN_PRICE",
	}
	RequestNewOrder_PriceField_value = map[string]int32{
		"BID_PRICE":   1,
		"OFFER_PRICE": 2,
		"TRADE_PRICE": 3,
		"LEAN_PRICE":  4,
	}
)

func (x RequestNewOrder_PriceField) Enum() *RequestNewOrder_PriceField {
	p := new(RequestNewOrder_PriceField)
	*p = x
	return p
}

func (x RequestNewOrder_PriceField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestNewOrder_PriceField) Descriptor() protoreflect.EnumDescriptor {
	return file_request_new_order_proto_enumTypes[4].Descriptor()
}

func (RequestNewOrder_PriceField) Type() protoreflect.EnumType {
	return &file_request_new_order_proto_enumTypes[4]
}

func (x RequestNewOrder_PriceField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestNewOrder_PriceField) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestNewOrder_PriceField(num)
	return nil
}

// Deprecated: Use RequestNewOrder_PriceField.Descriptor instead.
func (RequestNewOrder_PriceField) EnumDescriptor() ([]byte, []int) {
	return file_request_new_order_proto_rawDescGZIP(), []int{0, 4}
}

type RequestNewOrder_Condition int32

const (
	RequestNewOrder_EQUAL_TO              RequestNewOrder_Condition = 1
	RequestNewOrder_NOT_EQUAL_TO          RequestNewOrder_Condition = 2
	RequestNewOrder_GREATER_THAN          RequestNewOrder_Condition = 3
	RequestNewOrder_GREATER_THAN_EQUAL_TO RequestNewOrder_Condition = 4
	RequestNewOrder_LESSER_THAN           RequestNewOrder_Condition = 5
	RequestNewOrder_LESSER_THAN_EQUAL_TO  RequestNewOrder_Condition = 6
)

// Enum value maps for RequestNewOrder_Condition.
var (
	RequestNewOrder_Condition_name = map[int32]string{
		1: "EQUAL_TO",
		2: "NOT_EQUAL_TO",
		3: "GREATER_THAN",
		4: "GREATER_THAN_EQUAL_TO",
		5: "LESSER_THAN",
		6: "LESSER_THAN_EQUAL_TO",
	}
	RequestNewOrder_Condition_value = map[string]int32{
		"EQUAL_TO":              1,
		"NOT_EQUAL_TO":          2,
		"GREATER_THAN":          3,
		"GREATER_THAN_EQUAL_TO": 4,
		"LESSER_THAN":           5,
		"LESSER_THAN_EQUAL_TO":  6,
	}
)

func (x RequestNewOrder_Condition) Enum() *RequestNewOrder_Condition {
	p := new(RequestNewOrder_Condition)
	*p = x
	return p
}

func (x RequestNewOrder_Condition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestNewOrder_Condition) Descriptor() protoreflect.EnumDescriptor {
	return file_request_new_order_proto_enumTypes[5].Descriptor()
}

func (RequestNewOrder_Condition) Type() protoreflect.EnumType {
	return &file_request_new_order_proto_enumTypes[5]
}

func (x RequestNewOrder_Condition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestNewOrder_Condition) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestNewOrder_Condition(num)
	return nil
}

// Deprecated: Use RequestNewOrder_Condition.Descriptor instead.
func (RequestNewOrder_Condition) EnumDescriptor() ([]byte, []int) {
	return file_request_new_order_proto_rawDescGZIP(), []int{0, 5}
}

type RequestNewOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId          *int32                           `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	UserMsg             []string                         `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	UserTag             *string                          `protobuf:"bytes,154119,opt,name=user_tag,json=userTag" json:"user_tag,omitempty"`
	WindowName          *string                          `protobuf:"bytes,154629,opt,name=window_name,json=windowName" json:"window_name,omitempty"`
	FcmId               *string                          `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`
	IbId                *string                          `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`
	AccountId           *string                          `protobuf:"bytes,154008,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Symbol              *string                          `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`
	Exchange            *string                          `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`
	Quantity            *int32                           `protobuf:"varint,112004,opt,name=quantity" json:"quantity,omitempty"`
	Price               *float64                         `protobuf:"fixed64,110306,opt,name=price" json:"price,omitempty"`
	TriggerPrice        *float64                         `protobuf:"fixed64,149247,opt,name=trigger_price,json=triggerPrice" json:"trigger_price,omitempty"`
	TransactionType     *RequestNewOrder_TransactionType `protobuf:"varint,112003,opt,name=transaction_type,json=transactionType,enum=rti.RequestNewOrder_TransactionType" json:"transaction_type,omitempty"`
	Duration            *RequestNewOrder_Duration        `protobuf:"varint,112005,opt,name=duration,enum=rti.RequestNewOrder_Duration" json:"duration,omitempty"`
	PriceType           *RequestNewOrder_PriceType       `protobuf:"varint,112008,opt,name=price_type,json=priceType,enum=rti.RequestNewOrder_PriceType" json:"price_type,omitempty"`
	TradeRoute          *string                          `protobuf:"bytes,112016,opt,name=trade_route,json=tradeRoute" json:"trade_route,omitempty"`
	ManualOrAuto        *RequestNewOrder_OrderPlacement  `protobuf:"varint,154710,opt,name=manual_or_auto,json=manualOrAuto,enum=rti.RequestNewOrder_OrderPlacement" json:"manual_or_auto,omitempty"`
	TrailingStop        *bool                            `protobuf:"varint,157063,opt,name=trailing_stop,json=trailingStop" json:"trailing_stop,omitempty"`
	TrailByTicks        *int32                           `protobuf:"varint,157064,opt,name=trail_by_ticks,json=trailByTicks" json:"trail_by_ticks,omitempty"`
	TrailByPriceId      *int32                           `protobuf:"varint,157065,opt,name=trail_by_price_id,json=trailByPriceId" json:"trail_by_price_id,omitempty"`
	ReleaseAtSsboe      *int32                           `protobuf:"varint,154487,opt,name=release_at_ssboe,json=releaseAtSsboe" json:"release_at_ssboe,omitempty"`
	ReleaseAtUsecs      *int32                           `protobuf:"varint,154549,opt,name=release_at_usecs,json=releaseAtUsecs" json:"release_at_usecs,omitempty"`
	CancelAtSsboe       *int32                           `protobuf:"varint,157085,opt,name=cancel_at_ssboe,json=cancelAtSsboe" json:"cancel_at_ssboe,omitempty"`
	CancelAtUsecs       *int32                           `protobuf:"varint,157086,opt,name=cancel_at_usecs,json=cancelAtUsecs" json:"cancel_at_usecs,omitempty"`
	CancelAfterSecs     *int32                           `protobuf:"varint,154488,opt,name=cancel_after_secs,json=cancelAfterSecs" json:"cancel_after_secs,omitempty"`
	IfTouchedSymbol     *string                          `protobuf:"bytes,154451,opt,name=if_touched_symbol,json=ifTouchedSymbol" json:"if_touched_symbol,omitempty"`
	IfTouchedExchange   *string                          `protobuf:"bytes,154452,opt,name=if_touched_exchange,json=ifTouchedExchange" json:"if_touched_exchange,omitempty"`
	IfTouchedCondition  *RequestNewOrder_Condition       `protobuf:"varint,154453,opt,name=if_touched_condition,json=ifTouchedCondition,enum=rti.RequestNewOrder_Condition" json:"if_touched_condition,omitempty"`
	IfTouchedPriceField *RequestNewOrder_PriceField      `protobuf:"varint,154454,opt,name=if_touched_price_field,json=ifTouchedPriceField,enum=rti.RequestNewOrder_PriceField" json:"if_touched_price_field,omitempty"`
	IfTouchedPrice      *float64                         `protobuf:"fixed64,153632,opt,name=if_touched_price,json=ifTouchedPrice" json:"if_touched_price,omitempty"`
}

func (x *RequestNewOrder) Reset() {
	*x = RequestNewOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_new_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestNewOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestNewOrder) ProtoMessage() {}

func (x *RequestNewOrder) ProtoReflect() protoreflect.Message {
	mi := &file_request_new_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestNewOrder.ProtoReflect.Descriptor instead.
func (*RequestNewOrder) Descriptor() ([]byte, []int) {
	return file_request_new_order_proto_rawDescGZIP(), []int{0}
}

func (x *RequestNewOrder) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestNewOrder) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestNewOrder) GetUserTag() string {
	if x != nil && x.UserTag != nil {
		return *x.UserTag
	}
	return ""
}

func (x *RequestNewOrder) GetWindowName() string {
	if x != nil && x.WindowName != nil {
		return *x.WindowName
	}
	return ""
}

func (x *RequestNewOrder) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *RequestNewOrder) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *RequestNewOrder) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *RequestNewOrder) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *RequestNewOrder) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *RequestNewOrder) GetQuantity() int32 {
	if x != nil && x.Quantity != nil {
		return *x.Quantity
	}
	return 0
}

func (x *RequestNewOrder) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *RequestNewOrder) GetTriggerPrice() float64 {
	if x != nil && x.TriggerPrice != nil {
		return *x.TriggerPrice
	}
	return 0
}

func (x *RequestNewOrder) GetTransactionType() RequestNewOrder_TransactionType {
	if x != nil && x.TransactionType != nil {
		return *x.TransactionType
	}
	return RequestNewOrder_BUY
}

func (x *RequestNewOrder) GetDuration() RequestNewOrder_Duration {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return RequestNewOrder_DAY
}

func (x *RequestNewOrder) GetPriceType() RequestNewOrder_PriceType {
	if x != nil && x.PriceType != nil {
		return *x.PriceType
	}
	return RequestNewOrder_LIMIT
}

func (x *RequestNewOrder) GetTradeRoute() string {
	if x != nil && x.TradeRoute != nil {
		return *x.TradeRoute
	}
	return ""
}

func (x *RequestNewOrder) GetManualOrAuto() RequestNewOrder_OrderPlacement {
	if x != nil && x.ManualOrAuto != nil {
		return *x.ManualOrAuto
	}
	return RequestNewOrder_MANUAL
}

func (x *RequestNewOrder) GetTrailingStop() bool {
	if x != nil && x.TrailingStop != nil {
		return *x.TrailingStop
	}
	return false
}

func (x *RequestNewOrder) GetTrailByTicks() int32 {
	if x != nil && x.TrailByTicks != nil {
		return *x.TrailByTicks
	}
	return 0
}

func (x *RequestNewOrder) GetTrailByPriceId() int32 {
	if x != nil && x.TrailByPriceId != nil {
		return *x.TrailByPriceId
	}
	return 0
}

func (x *RequestNewOrder) GetReleaseAtSsboe() int32 {
	if x != nil && x.ReleaseAtSsboe != nil {
		return *x.ReleaseAtSsboe
	}
	return 0
}

func (x *RequestNewOrder) GetReleaseAtUsecs() int32 {
	if x != nil && x.ReleaseAtUsecs != nil {
		return *x.ReleaseAtUsecs
	}
	return 0
}

func (x *RequestNewOrder) GetCancelAtSsboe() int32 {
	if x != nil && x.CancelAtSsboe != nil {
		return *x.CancelAtSsboe
	}
	return 0
}

func (x *RequestNewOrder) GetCancelAtUsecs() int32 {
	if x != nil && x.CancelAtUsecs != nil {
		return *x.CancelAtUsecs
	}
	return 0
}

func (x *RequestNewOrder) GetCancelAfterSecs() int32 {
	if x != nil && x.CancelAfterSecs != nil {
		return *x.CancelAfterSecs
	}
	return 0
}

func (x *RequestNewOrder) GetIfTouchedSymbol() string {
	if x != nil && x.IfTouchedSymbol != nil {
		return *x.IfTouchedSymbol
	}
	return ""
}

func (x *RequestNewOrder) GetIfTouchedExchange() string {
	if x != nil && x.IfTouchedExchange != nil {
		return *x.IfTouchedExchange
	}
	return ""
}

func (x *RequestNewOrder) GetIfTouchedCondition() RequestNewOrder_Condition {
	if x != nil && x.IfTouchedCondition != nil {
		return *x.IfTouchedCondition
	}
	return RequestNewOrder_EQUAL_TO
}

func (x *RequestNewOrder) GetIfTouchedPriceField() RequestNewOrder_PriceField {
	if x != nil && x.IfTouchedPriceField != nil {
		return *x.IfTouchedPriceField
	}
	return RequestNewOrder_BID_PRICE
}

func (x *RequestNewOrder) GetIfTouchedPrice() float64 {
	if x != nil && x.IfTouchedPrice != nil {
		return *x.IfTouchedPrice
	}
	return 0
}

var File_request_new_order_proto protoreflect.FileDescriptor

var file_request_new_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0x8b,
	0x0e, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d,
	0x73, 0x67, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x87,
	0xb4, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12,
	0x21, 0x0a, 0x0b, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x85,
	0xb8, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x98, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x84, 0xeb, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0xe2, 0xdd, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x0d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0xff, 0x8d, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x83, 0xeb, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x85, 0xeb, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x77,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x88, 0xeb, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x90, 0xeb, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x0e,
	0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x18, 0xd6,
	0xb8, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x6d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x4f, 0x72, 0x41, 0x75, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x0d, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x87, 0xcb, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x6f, 0x70,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x69, 0x63,
	0x6b, 0x73, 0x18, 0x88, 0xcb, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x69,
	0x6c, 0x42, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x69,
	0x6c, 0x5f, 0x62, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x89, 0xcb,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x61, 0x74, 0x5f, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xf7, 0xb6, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x74, 0x53, 0x73, 0x62, 0x6f,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xb5, 0xb7, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x74, 0x55, 0x73, 0x65, 0x63, 0x73, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x61, 0x74, 0x5f, 0x73, 0x73, 0x62, 0x6f, 0x65,
	0x18, 0x9d, 0xcb, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x41, 0x74, 0x53, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x5f, 0x61, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0x9e, 0xcb, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x74, 0x55, 0x73, 0x65, 0x63,
	0x73, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0xf8, 0xb6, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x65, 0x63, 0x73, 0x12,
	0x2c, 0x0a, 0x11, 0x69, 0x66, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0xd3, 0xb6, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x66,
	0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x30, 0x0a,
	0x13, 0x69, 0x66, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0xd4, 0xb6, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x66,
	0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x52, 0x0a, 0x14, 0x69, 0x66, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xd5, 0xb6, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x77,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x69, 0x66, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x16, 0x69, 0x66, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0xd6, 0xb6,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x13, 0x69, 0x66, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x69,
	0x66, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0xa0, 0xb0, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x69, 0x66, 0x54, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55,
	0x59, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x02, 0x22, 0x26, 0x0a,
	0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x41,
	0x55, 0x54, 0x4f, 0x10, 0x02, 0x22, 0x2e, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x54,
	0x43, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x43, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x46, 0x4f, 0x4b, 0x10, 0x04, 0x22, 0x70, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x4f,
	0x50, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x4f,
	0x50, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x41,
	0x52, 0x4b, 0x45, 0x54, 0x5f, 0x49, 0x46, 0x5f, 0x54, 0x4f, 0x55, 0x43, 0x48, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x49, 0x46, 0x5f, 0x54, 0x4f,
	0x55, 0x43, 0x48, 0x45, 0x44, 0x10, 0x06, 0x22, 0x4d, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x49, 0x44, 0x5f, 0x50, 0x52, 0x49,
	0x43, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x5f, 0x50, 0x52,
	0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x50,
	0x52, 0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x41, 0x4e, 0x5f, 0x50,
	0x52, 0x49, 0x43, 0x45, 0x10, 0x04, 0x22, 0x83, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x5f, 0x54, 0x4f,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x5f,
	0x54, 0x4f, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f,
	0x54, 0x48, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x5f, 0x54, 0x4f, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x53, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e,
	0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x45, 0x53, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41,
	0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x5f, 0x54, 0x4f, 0x10, 0x06, 0x42, 0x1e, 0x5a, 0x1c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76,
	0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_request_new_order_proto_rawDescOnce sync.Once
	file_request_new_order_proto_rawDescData = file_request_new_order_proto_rawDesc
)

func file_request_new_order_proto_rawDescGZIP() []byte {
	file_request_new_order_proto_rawDescOnce.Do(func() {
		file_request_new_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_new_order_proto_rawDescData)
	})
	return file_request_new_order_proto_rawDescData
}

var file_request_new_order_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_request_new_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_new_order_proto_goTypes = []interface{}{
	(RequestNewOrder_TransactionType)(0), // 0: rti.RequestNewOrder.TransactionType
	(RequestNewOrder_OrderPlacement)(0),  // 1: rti.RequestNewOrder.OrderPlacement
	(RequestNewOrder_Duration)(0),        // 2: rti.RequestNewOrder.Duration
	(RequestNewOrder_PriceType)(0),       // 3: rti.RequestNewOrder.PriceType
	(RequestNewOrder_PriceField)(0),      // 4: rti.RequestNewOrder.PriceField
	(RequestNewOrder_Condition)(0),       // 5: rti.RequestNewOrder.Condition
	(*RequestNewOrder)(nil),              // 6: rti.RequestNewOrder
}
var file_request_new_order_proto_depIdxs = []int32{
	0, // 0: rti.RequestNewOrder.transaction_type:type_name -> rti.RequestNewOrder.TransactionType
	2, // 1: rti.RequestNewOrder.duration:type_name -> rti.RequestNewOrder.Duration
	3, // 2: rti.RequestNewOrder.price_type:type_name -> rti.RequestNewOrder.PriceType
	1, // 3: rti.RequestNewOrder.manual_or_auto:type_name -> rti.RequestNewOrder.OrderPlacement
	5, // 4: rti.RequestNewOrder.if_touched_condition:type_name -> rti.RequestNewOrder.Condition
	4, // 5: rti.RequestNewOrder.if_touched_price_field:type_name -> rti.RequestNewOrder.PriceField
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_request_new_order_proto_init() }
func file_request_new_order_proto_init() {
	if File_request_new_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_new_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestNewOrder); i {
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
			RawDescriptor: file_request_new_order_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_new_order_proto_goTypes,
		DependencyIndexes: file_request_new_order_proto_depIdxs,
		EnumInfos:         file_request_new_order_proto_enumTypes,
		MessageInfos:      file_request_new_order_proto_msgTypes,
	}.Build()
	File_request_new_order_proto = out.File
	file_request_new_order_proto_rawDesc = nil
	file_request_new_order_proto_goTypes = nil
	file_request_new_order_proto_depIdxs = nil
}
