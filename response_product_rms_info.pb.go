// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: response_product_rms_info.proto

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

// below enum is just for reference only, not used in this message
type ResponseProductRmsInfo_PresenceBits int32

const (
	ResponseProductRmsInfo_BUY_LIMIT            ResponseProductRmsInfo_PresenceBits = 1
	ResponseProductRmsInfo_SELL_LIMIT           ResponseProductRmsInfo_PresenceBits = 2
	ResponseProductRmsInfo_LOSS_LIMIT           ResponseProductRmsInfo_PresenceBits = 4
	ResponseProductRmsInfo_MAX_ORDER_QUANTITY   ResponseProductRmsInfo_PresenceBits = 8
	ResponseProductRmsInfo_BUY_MARGIN_RATE      ResponseProductRmsInfo_PresenceBits = 16
	ResponseProductRmsInfo_SELL_MARGIN_RATE     ResponseProductRmsInfo_PresenceBits = 32
	ResponseProductRmsInfo_COMMISSION_FILL_RATE ResponseProductRmsInfo_PresenceBits = 64
)

// Enum value maps for ResponseProductRmsInfo_PresenceBits.
var (
	ResponseProductRmsInfo_PresenceBits_name = map[int32]string{
		1:  "BUY_LIMIT",
		2:  "SELL_LIMIT",
		4:  "LOSS_LIMIT",
		8:  "MAX_ORDER_QUANTITY",
		16: "BUY_MARGIN_RATE",
		32: "SELL_MARGIN_RATE",
		64: "COMMISSION_FILL_RATE",
	}
	ResponseProductRmsInfo_PresenceBits_value = map[string]int32{
		"BUY_LIMIT":            1,
		"SELL_LIMIT":           2,
		"LOSS_LIMIT":           4,
		"MAX_ORDER_QUANTITY":   8,
		"BUY_MARGIN_RATE":      16,
		"SELL_MARGIN_RATE":     32,
		"COMMISSION_FILL_RATE": 64,
	}
)

func (x ResponseProductRmsInfo_PresenceBits) Enum() *ResponseProductRmsInfo_PresenceBits {
	p := new(ResponseProductRmsInfo_PresenceBits)
	*p = x
	return p
}

func (x ResponseProductRmsInfo_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseProductRmsInfo_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_response_product_rms_info_proto_enumTypes[0].Descriptor()
}

func (ResponseProductRmsInfo_PresenceBits) Type() protoreflect.EnumType {
	return &file_response_product_rms_info_proto_enumTypes[0]
}

func (x ResponseProductRmsInfo_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResponseProductRmsInfo_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResponseProductRmsInfo_PresenceBits(num)
	return nil
}

// Deprecated: Use ResponseProductRmsInfo_PresenceBits.Descriptor instead.
func (ResponseProductRmsInfo_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_response_product_rms_info_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseProductRmsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId         *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                            // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg            []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                      // PB_OFFSET + MNM_USER_MSG
	RqHandlerRpCode    []string `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"`          // PB_OFFSET + MNM_REQUEST_HANDLER_RESPONSE_CODE
	RpCode             []string `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                         // PB_OFFSET + MNM_RESPONSE_CODE
	PresenceBits       *uint32  `protobuf:"varint,153622,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"`                      // PB_OFFSET + MNM_FIELD_PRESENCE_BITS
	FcmId              *string  `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`                                            // PB_OFFSET + MNM_FCM_ID
	IbId               *string  `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`                                               // PB_OFFSET + MNM_IB_ID
	AccountId          *string  `protobuf:"bytes,154008,opt,name=account_id,json=accountId" json:"account_id,omitempty"`                                // PB_OFFSET + MNM_ACCOUNT_ID
	ProductCode        *string  `protobuf:"bytes,100749,opt,name=product_code,json=productCode" json:"product_code,omitempty"`                          // PB_OFFSET + MNM_PRODUCT_CODE
	LossLimit          *float64 `protobuf:"fixed64,154019,opt,name=loss_limit,json=lossLimit" json:"loss_limit,omitempty"`                              // PB_OFFSET + MNM_LOSS_LIMIT
	CommissionFillRate *float64 `protobuf:"fixed64,156969,opt,name=commission_fill_rate,json=commissionFillRate" json:"commission_fill_rate,omitempty"` // PB_OFFSET +	MNM_COMMISSION_FILL_RATE
	BuyMarginRate      *float64 `protobuf:"fixed64,157003,opt,name=buy_margin_rate,json=buyMarginRate" json:"buy_margin_rate,omitempty"`                // PB_OFFSET + MNM_BUY_MARGIN_RATE
	SellMarginRate     *float64 `protobuf:"fixed64,157004,opt,name=sell_margin_rate,json=sellMarginRate" json:"sell_margin_rate,omitempty"`             // PB_OFFSET + MNM_SELL_MARGIN_RATE
	BuyLimit           *int32   `protobuf:"varint,154009,opt,name=buy_limit,json=buyLimit" json:"buy_limit,omitempty"`                                  // PB_OFFSET + MNM_BUY_LIMIT
	MaxOrderQuantity   *int32   `protobuf:"varint,110105,opt,name=max_order_quantity,json=maxOrderQuantity" json:"max_order_quantity,omitempty"`        // PB_OFFSET + MNM_MAX_LIMIT_QUAN
	SellLimit          *int32   `protobuf:"varint,154035,opt,name=sell_limit,json=sellLimit" json:"sell_limit,omitempty"`                               // PB_OFFSET + MNM_SELL_LIMIT
}

func (x *ResponseProductRmsInfo) Reset() {
	*x = ResponseProductRmsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_product_rms_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseProductRmsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseProductRmsInfo) ProtoMessage() {}

func (x *ResponseProductRmsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_response_product_rms_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseProductRmsInfo.ProtoReflect.Descriptor instead.
func (*ResponseProductRmsInfo) Descriptor() ([]byte, []int) {
	return file_response_product_rms_info_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseProductRmsInfo) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseProductRmsInfo) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseProductRmsInfo) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseProductRmsInfo) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *ResponseProductRmsInfo) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *ResponseProductRmsInfo) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *ResponseProductRmsInfo) GetProductCode() string {
	if x != nil && x.ProductCode != nil {
		return *x.ProductCode
	}
	return ""
}

func (x *ResponseProductRmsInfo) GetLossLimit() float64 {
	if x != nil && x.LossLimit != nil {
		return *x.LossLimit
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetCommissionFillRate() float64 {
	if x != nil && x.CommissionFillRate != nil {
		return *x.CommissionFillRate
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetBuyMarginRate() float64 {
	if x != nil && x.BuyMarginRate != nil {
		return *x.BuyMarginRate
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetSellMarginRate() float64 {
	if x != nil && x.SellMarginRate != nil {
		return *x.SellMarginRate
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetBuyLimit() int32 {
	if x != nil && x.BuyLimit != nil {
		return *x.BuyLimit
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetMaxOrderQuantity() int32 {
	if x != nil && x.MaxOrderQuantity != nil {
		return *x.MaxOrderQuantity
	}
	return 0
}

func (x *ResponseProductRmsInfo) GetSellLimit() int32 {
	if x != nil && x.SellLimit != nil {
		return *x.SellLimit
	}
	return 0
}

var File_response_product_rms_info_proto protoreflect.FileDescriptor

var file_response_product_rms_info_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x72, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xf7, 0x05, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x6d, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9e, 0x8d, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x96, 0xb0, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x69,
	0x74, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x98, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x8d, 0x93, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6c, 0x6f, 0x73, 0x73,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0xa3, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6c, 0x6f, 0x73, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0xa9, 0xca, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x62, 0x75, 0x79, 0x5f, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0xcb, 0xca, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x62, 0x75, 0x79, 0x4d, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x6c, 0x5f,
	0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0xcc, 0xca, 0x09, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x6c, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x99, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x99, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x6d, 0x61, 0x78, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0xb3, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x42, 0x69, 0x74, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x55, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x4c, 0x4c, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x53, 0x53, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x58, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x42,
	0x55, 0x59, 0x5f, 0x4d, 0x41, 0x52, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x10,
	0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x4c, 0x4c, 0x5f, 0x4d, 0x41, 0x52, 0x47, 0x49, 0x4e, 0x5f,
	0x52, 0x41, 0x54, 0x45, 0x10, 0x20, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4c, 0x4c, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x40,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x67, 0x61, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_response_product_rms_info_proto_rawDescOnce sync.Once
	file_response_product_rms_info_proto_rawDescData = file_response_product_rms_info_proto_rawDesc
)

func file_response_product_rms_info_proto_rawDescGZIP() []byte {
	file_response_product_rms_info_proto_rawDescOnce.Do(func() {
		file_response_product_rms_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_product_rms_info_proto_rawDescData)
	})
	return file_response_product_rms_info_proto_rawDescData
}

var file_response_product_rms_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_product_rms_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_product_rms_info_proto_goTypes = []interface{}{
	(ResponseProductRmsInfo_PresenceBits)(0), // 0: rti.ResponseProductRmsInfo.PresenceBits
	(*ResponseProductRmsInfo)(nil),           // 1: rti.ResponseProductRmsInfo
}
var file_response_product_rms_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_product_rms_info_proto_init() }
func file_response_product_rms_info_proto_init() {
	if File_response_product_rms_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_product_rms_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseProductRmsInfo); i {
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
			RawDescriptor: file_response_product_rms_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_product_rms_info_proto_goTypes,
		DependencyIndexes: file_response_product_rms_info_proto_depIdxs,
		EnumInfos:         file_response_product_rms_info_proto_enumTypes,
		MessageInfos:      file_response_product_rms_info_proto_msgTypes,
	}.Build()
	File_response_product_rms_info_proto = out.File
	file_response_product_rms_info_proto_rawDesc = nil
	file_response_product_rms_info_proto_goTypes = nil
	file_response_product_rms_info_proto_depIdxs = nil
}
