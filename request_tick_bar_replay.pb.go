// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: request_tick_bar_replay.proto

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

type RequestTickBarReplay_BarType int32

const (
	RequestTickBarReplay_TICK_BAR   RequestTickBarReplay_BarType = 1
	RequestTickBarReplay_RANGE_BAR  RequestTickBarReplay_BarType = 2
	RequestTickBarReplay_VOLUME_BAR RequestTickBarReplay_BarType = 3
)

// Enum value maps for RequestTickBarReplay_BarType.
var (
	RequestTickBarReplay_BarType_name = map[int32]string{
		1: "TICK_BAR",
		2: "RANGE_BAR",
		3: "VOLUME_BAR",
	}
	RequestTickBarReplay_BarType_value = map[string]int32{
		"TICK_BAR":   1,
		"RANGE_BAR":  2,
		"VOLUME_BAR": 3,
	}
)

func (x RequestTickBarReplay_BarType) Enum() *RequestTickBarReplay_BarType {
	p := new(RequestTickBarReplay_BarType)
	*p = x
	return p
}

func (x RequestTickBarReplay_BarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestTickBarReplay_BarType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_tick_bar_replay_proto_enumTypes[0].Descriptor()
}

func (RequestTickBarReplay_BarType) Type() protoreflect.EnumType {
	return &file_request_tick_bar_replay_proto_enumTypes[0]
}

func (x RequestTickBarReplay_BarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestTickBarReplay_BarType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestTickBarReplay_BarType(num)
	return nil
}

// Deprecated: Use RequestTickBarReplay_BarType.Descriptor instead.
func (RequestTickBarReplay_BarType) EnumDescriptor() ([]byte, []int) {
	return file_request_tick_bar_replay_proto_rawDescGZIP(), []int{0, 0}
}

type RequestTickBarReplay_BarSubType int32

const (
	RequestTickBarReplay_REGULAR RequestTickBarReplay_BarSubType = 1
	RequestTickBarReplay_CUSTOM  RequestTickBarReplay_BarSubType = 2
)

// Enum value maps for RequestTickBarReplay_BarSubType.
var (
	RequestTickBarReplay_BarSubType_name = map[int32]string{
		1: "REGULAR",
		2: "CUSTOM",
	}
	RequestTickBarReplay_BarSubType_value = map[string]int32{
		"REGULAR": 1,
		"CUSTOM":  2,
	}
)

func (x RequestTickBarReplay_BarSubType) Enum() *RequestTickBarReplay_BarSubType {
	p := new(RequestTickBarReplay_BarSubType)
	*p = x
	return p
}

func (x RequestTickBarReplay_BarSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestTickBarReplay_BarSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_tick_bar_replay_proto_enumTypes[1].Descriptor()
}

func (RequestTickBarReplay_BarSubType) Type() protoreflect.EnumType {
	return &file_request_tick_bar_replay_proto_enumTypes[1]
}

func (x RequestTickBarReplay_BarSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestTickBarReplay_BarSubType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestTickBarReplay_BarSubType(num)
	return nil
}

// Deprecated: Use RequestTickBarReplay_BarSubType.Descriptor instead.
func (RequestTickBarReplay_BarSubType) EnumDescriptor() ([]byte, []int) {
	return file_request_tick_bar_replay_proto_rawDescGZIP(), []int{0, 1}
}

type RequestTickBarReplay_Direction int32

const (
	RequestTickBarReplay_FIRST RequestTickBarReplay_Direction = 1
	RequestTickBarReplay_LAST  RequestTickBarReplay_Direction = 2
)

// Enum value maps for RequestTickBarReplay_Direction.
var (
	RequestTickBarReplay_Direction_name = map[int32]string{
		1: "FIRST",
		2: "LAST",
	}
	RequestTickBarReplay_Direction_value = map[string]int32{
		"FIRST": 1,
		"LAST":  2,
	}
)

func (x RequestTickBarReplay_Direction) Enum() *RequestTickBarReplay_Direction {
	p := new(RequestTickBarReplay_Direction)
	*p = x
	return p
}

func (x RequestTickBarReplay_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestTickBarReplay_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_request_tick_bar_replay_proto_enumTypes[2].Descriptor()
}

func (RequestTickBarReplay_Direction) Type() protoreflect.EnumType {
	return &file_request_tick_bar_replay_proto_enumTypes[2]
}

func (x RequestTickBarReplay_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestTickBarReplay_Direction) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestTickBarReplay_Direction(num)
	return nil
}

// Deprecated: Use RequestTickBarReplay_Direction.Descriptor instead.
func (RequestTickBarReplay_Direction) EnumDescriptor() ([]byte, []int) {
	return file_request_tick_bar_replay_proto_rawDescGZIP(), []int{0, 2}
}

type RequestTickBarReplay_TimeOrder int32

const (
	RequestTickBarReplay_FORWARDS  RequestTickBarReplay_TimeOrder = 1
	RequestTickBarReplay_BACKWARDS RequestTickBarReplay_TimeOrder = 2
)

// Enum value maps for RequestTickBarReplay_TimeOrder.
var (
	RequestTickBarReplay_TimeOrder_name = map[int32]string{
		1: "FORWARDS",
		2: "BACKWARDS",
	}
	RequestTickBarReplay_TimeOrder_value = map[string]int32{
		"FORWARDS":  1,
		"BACKWARDS": 2,
	}
)

func (x RequestTickBarReplay_TimeOrder) Enum() *RequestTickBarReplay_TimeOrder {
	p := new(RequestTickBarReplay_TimeOrder)
	*p = x
	return p
}

func (x RequestTickBarReplay_TimeOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestTickBarReplay_TimeOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_request_tick_bar_replay_proto_enumTypes[3].Descriptor()
}

func (RequestTickBarReplay_TimeOrder) Type() protoreflect.EnumType {
	return &file_request_tick_bar_replay_proto_enumTypes[3]
}

func (x RequestTickBarReplay_TimeOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestTickBarReplay_TimeOrder) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestTickBarReplay_TimeOrder(num)
	return nil
}

// Deprecated: Use RequestTickBarReplay_TimeOrder.Descriptor instead.
func (RequestTickBarReplay_TimeOrder) EnumDescriptor() ([]byte, []int) {
	return file_request_tick_bar_replay_proto_rawDescGZIP(), []int{0, 3}
}

type RequestTickBarReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId            *int32                           `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	UserMsg               []string                         `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	Symbol                *string                          `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`
	Exchange              *string                          `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`
	BarType               *RequestTickBarReplay_BarType    `protobuf:"varint,119200,opt,name=bar_type,json=barType,enum=rti.RequestTickBarReplay_BarType" json:"bar_type,omitempty"`
	BarSubType            *RequestTickBarReplay_BarSubType `protobuf:"varint,119208,opt,name=bar_sub_type,json=barSubType,enum=rti.RequestTickBarReplay_BarSubType" json:"bar_sub_type,omitempty"`
	BarTypeSpecifier      *string                          `protobuf:"bytes,148162,opt,name=bar_type_specifier,json=barTypeSpecifier" json:"bar_type_specifier,omitempty"`
	StartIndex            *int32                           `protobuf:"varint,153002,opt,name=start_index,json=startIndex" json:"start_index,omitempty"`
	FinishIndex           *int32                           `protobuf:"varint,153003,opt,name=finish_index,json=finishIndex" json:"finish_index,omitempty"`
	UserMaxCount          *int32                           `protobuf:"varint,154020,opt,name=user_max_count,json=userMaxCount" json:"user_max_count,omitempty"`
	CustomSessionOpenSsm  *int32                           `protobuf:"varint,119209,opt,name=custom_session_open_ssm,json=customSessionOpenSsm" json:"custom_session_open_ssm,omitempty"`
	CustomSessionCloseSsm *int32                           `protobuf:"varint,119210,opt,name=custom_session_close_ssm,json=customSessionCloseSsm" json:"custom_session_close_ssm,omitempty"`
	Direction             *RequestTickBarReplay_Direction  `protobuf:"varint,149253,opt,name=direction,enum=rti.RequestTickBarReplay_Direction" json:"direction,omitempty"`
	TimeOrder             *RequestTickBarReplay_TimeOrder  `protobuf:"varint,149307,opt,name=time_order,json=timeOrder,enum=rti.RequestTickBarReplay_TimeOrder" json:"time_order,omitempty"`
	ResumeBars            *bool                            `protobuf:"varint,153642,opt,name=resume_bars,json=resumeBars" json:"resume_bars,omitempty"`
}

func (x *RequestTickBarReplay) Reset() {
	*x = RequestTickBarReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_tick_bar_replay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTickBarReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTickBarReplay) ProtoMessage() {}

func (x *RequestTickBarReplay) ProtoReflect() protoreflect.Message {
	mi := &file_request_tick_bar_replay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTickBarReplay.ProtoReflect.Descriptor instead.
func (*RequestTickBarReplay) Descriptor() ([]byte, []int) {
	return file_request_tick_bar_replay_proto_rawDescGZIP(), []int{0}
}

func (x *RequestTickBarReplay) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestTickBarReplay) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestTickBarReplay) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *RequestTickBarReplay) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *RequestTickBarReplay) GetBarType() RequestTickBarReplay_BarType {
	if x != nil && x.BarType != nil {
		return *x.BarType
	}
	return RequestTickBarReplay_TICK_BAR
}

func (x *RequestTickBarReplay) GetBarSubType() RequestTickBarReplay_BarSubType {
	if x != nil && x.BarSubType != nil {
		return *x.BarSubType
	}
	return RequestTickBarReplay_REGULAR
}

func (x *RequestTickBarReplay) GetBarTypeSpecifier() string {
	if x != nil && x.BarTypeSpecifier != nil {
		return *x.BarTypeSpecifier
	}
	return ""
}

func (x *RequestTickBarReplay) GetStartIndex() int32 {
	if x != nil && x.StartIndex != nil {
		return *x.StartIndex
	}
	return 0
}

func (x *RequestTickBarReplay) GetFinishIndex() int32 {
	if x != nil && x.FinishIndex != nil {
		return *x.FinishIndex
	}
	return 0
}

func (x *RequestTickBarReplay) GetUserMaxCount() int32 {
	if x != nil && x.UserMaxCount != nil {
		return *x.UserMaxCount
	}
	return 0
}

func (x *RequestTickBarReplay) GetCustomSessionOpenSsm() int32 {
	if x != nil && x.CustomSessionOpenSsm != nil {
		return *x.CustomSessionOpenSsm
	}
	return 0
}

func (x *RequestTickBarReplay) GetCustomSessionCloseSsm() int32 {
	if x != nil && x.CustomSessionCloseSsm != nil {
		return *x.CustomSessionCloseSsm
	}
	return 0
}

func (x *RequestTickBarReplay) GetDirection() RequestTickBarReplay_Direction {
	if x != nil && x.Direction != nil {
		return *x.Direction
	}
	return RequestTickBarReplay_FIRST
}

func (x *RequestTickBarReplay) GetTimeOrder() RequestTickBarReplay_TimeOrder {
	if x != nil && x.TimeOrder != nil {
		return *x.TimeOrder
	}
	return RequestTickBarReplay_FORWARDS
}

func (x *RequestTickBarReplay) GetResumeBars() bool {
	if x != nil && x.ResumeBars != nil {
		return *x.ResumeBars
	}
	return false
}

var File_request_tick_bar_replay_proto protoreflect.FileDescriptor

var file_request_tick_bar_replay_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x62,
	0x61, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x74, 0x69, 0x22, 0x85, 0x07, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x21, 0x0a,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x62, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0xa0, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x74, 0x69, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x79, 0x2e, 0x42, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x62, 0x61,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x62, 0x61, 0x72, 0x5f, 0x73, 0x75, 0x62,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xa8, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x42,
	0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x42, 0x61, 0x72, 0x53, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x62, 0x61, 0x72, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x62, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0xc2, 0x85, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0xaa,
	0xab, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x23, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0xab, 0xab, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xa4, 0xb3, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x37, 0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x73, 0x6d, 0x18, 0xa9, 0xa3, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x73, 0x6d, 0x12, 0x39, 0x0a, 0x18, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x5f, 0x73, 0x73, 0x6d, 0x18, 0xaa, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x73, 0x6d, 0x12, 0x43, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x85, 0x8e, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0xbb, 0x8e, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23,
	0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x42, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x62, 0x61, 0x72, 0x73, 0x18, 0xaa, 0xb0,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x61, 0x72,
	0x73, 0x22, 0x36, 0x0a, 0x07, 0x42, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x54, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x4f, 0x4c,
	0x55, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x03, 0x22, 0x25, 0x0a, 0x0a, 0x42, 0x61, 0x72,
	0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x47, 0x55, 0x4c,
	0x41, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x02,
	0x22, 0x20, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x41, 0x53, 0x54,
	0x10, 0x02, 0x22, 0x28, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x42, 0x41, 0x43, 0x4b, 0x57, 0x41, 0x52, 0x44, 0x53, 0x10, 0x02, 0x42, 0x1e, 0x5a, 0x1c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x61, 0x72, 0x76,
	0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_request_tick_bar_replay_proto_rawDescOnce sync.Once
	file_request_tick_bar_replay_proto_rawDescData = file_request_tick_bar_replay_proto_rawDesc
)

func file_request_tick_bar_replay_proto_rawDescGZIP() []byte {
	file_request_tick_bar_replay_proto_rawDescOnce.Do(func() {
		file_request_tick_bar_replay_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_tick_bar_replay_proto_rawDescData)
	})
	return file_request_tick_bar_replay_proto_rawDescData
}

var file_request_tick_bar_replay_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_request_tick_bar_replay_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_tick_bar_replay_proto_goTypes = []interface{}{
	(RequestTickBarReplay_BarType)(0),    // 0: rti.RequestTickBarReplay.BarType
	(RequestTickBarReplay_BarSubType)(0), // 1: rti.RequestTickBarReplay.BarSubType
	(RequestTickBarReplay_Direction)(0),  // 2: rti.RequestTickBarReplay.Direction
	(RequestTickBarReplay_TimeOrder)(0),  // 3: rti.RequestTickBarReplay.TimeOrder
	(*RequestTickBarReplay)(nil),         // 4: rti.RequestTickBarReplay
}
var file_request_tick_bar_replay_proto_depIdxs = []int32{
	0, // 0: rti.RequestTickBarReplay.bar_type:type_name -> rti.RequestTickBarReplay.BarType
	1, // 1: rti.RequestTickBarReplay.bar_sub_type:type_name -> rti.RequestTickBarReplay.BarSubType
	2, // 2: rti.RequestTickBarReplay.direction:type_name -> rti.RequestTickBarReplay.Direction
	3, // 3: rti.RequestTickBarReplay.time_order:type_name -> rti.RequestTickBarReplay.TimeOrder
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_request_tick_bar_replay_proto_init() }
func file_request_tick_bar_replay_proto_init() {
	if File_request_tick_bar_replay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_tick_bar_replay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTickBarReplay); i {
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
			RawDescriptor: file_request_tick_bar_replay_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_tick_bar_replay_proto_goTypes,
		DependencyIndexes: file_request_tick_bar_replay_proto_depIdxs,
		EnumInfos:         file_request_tick_bar_replay_proto_enumTypes,
		MessageInfos:      file_request_tick_bar_replay_proto_msgTypes,
	}.Build()
	File_request_tick_bar_replay_proto = out.File
	file_request_tick_bar_replay_proto_rawDesc = nil
	file_request_tick_bar_replay_proto_goTypes = nil
	file_request_tick_bar_replay_proto_depIdxs = nil
}
