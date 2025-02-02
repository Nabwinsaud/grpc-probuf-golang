// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: proto/payment.proto

package payment_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentMethod int32

const (
	PaymentMethod_STRIPE           PaymentMethod = 0
	PaymentMethod_PAYPAL           PaymentMethod = 1
	PaymentMethod_GOOGLE_PAY       PaymentMethod = 2
	PaymentMethod_CREDIT_CARD      PaymentMethod = 3
	PaymentMethod_CASH_ON_DELIVERY PaymentMethod = 4
)

// Enum value maps for PaymentMethod.
var (
	PaymentMethod_name = map[int32]string{
		0: "STRIPE",
		1: "PAYPAL",
		2: "GOOGLE_PAY",
		3: "CREDIT_CARD",
		4: "CASH_ON_DELIVERY",
	}
	PaymentMethod_value = map[string]int32{
		"STRIPE":           0,
		"PAYPAL":           1,
		"GOOGLE_PAY":       2,
		"CREDIT_CARD":      3,
		"CASH_ON_DELIVERY": 4,
	}
)

func (x PaymentMethod) Enum() *PaymentMethod {
	p := new(PaymentMethod)
	*p = x
	return p
}

func (x PaymentMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_payment_proto_enumTypes[0].Descriptor()
}

func (PaymentMethod) Type() protoreflect.EnumType {
	return &file_proto_payment_proto_enumTypes[0]
}

func (x PaymentMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethod.Descriptor instead.
func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{0}
}

type PaymentDetails struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	StripePaymentId         string                 `protobuf:"bytes,5,opt,name=stripe_payment_id,json=stripePaymentId,proto3" json:"stripe_payment_id,omitempty"`
	PaypalPaymentId         string                 `protobuf:"bytes,6,opt,name=paypal_payment_id,json=paypalPaymentId,proto3" json:"paypal_payment_id,omitempty"`
	GooglePayPaymentId      string                 `protobuf:"bytes,7,opt,name=google_pay_payment_id,json=googlePayPaymentId,proto3" json:"google_pay_payment_id,omitempty"`
	CreditCardPaymentId     string                 `protobuf:"bytes,8,opt,name=credit_card_payment_id,json=creditCardPaymentId,proto3" json:"credit_card_payment_id,omitempty"`
	CashOnDeliveryPaymentId string                 `protobuf:"bytes,9,opt,name=cash_on_delivery_payment_id,json=cashOnDeliveryPaymentId,proto3" json:"cash_on_delivery_payment_id,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *PaymentDetails) Reset() {
	*x = PaymentDetails{}
	mi := &file_proto_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentDetails) ProtoMessage() {}

func (x *PaymentDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentDetails.ProtoReflect.Descriptor instead.
func (*PaymentDetails) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentDetails) GetStripePaymentId() string {
	if x != nil {
		return x.StripePaymentId
	}
	return ""
}

func (x *PaymentDetails) GetPaypalPaymentId() string {
	if x != nil {
		return x.PaypalPaymentId
	}
	return ""
}

func (x *PaymentDetails) GetGooglePayPaymentId() string {
	if x != nil {
		return x.GooglePayPaymentId
	}
	return ""
}

func (x *PaymentDetails) GetCreditCardPaymentId() string {
	if x != nil {
		return x.CreditCardPaymentId
	}
	return ""
}

func (x *PaymentDetails) GetCashOnDeliveryPaymentId() string {
	if x != nil {
		return x.CashOnDeliveryPaymentId
	}
	return ""
}

type Payment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	PaymentMethod PaymentMethod          `protobuf:"varint,3,opt,name=payment_method,json=paymentMethod,proto3,enum=payment.PaymentMethod" json:"payment_method,omitempty"`
	PaymentDate   string                 `protobuf:"bytes,4,opt,name=payment_date,json=paymentDate,proto3" json:"payment_date,omitempty"`
	// oneof is used to define the one field from multiple fields
	PaymentDetails    *PaymentDetails        `protobuf:"bytes,5,opt,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	PaymentLinkExpiry *durationpb.Duration   `protobuf:"bytes,11,opt,name=payment_link_expiry,json=paymentLinkExpiry,proto3" json:"payment_link_expiry,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Payment) Reset() {
	*x = Payment{}
	mi := &file_proto_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Payment) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *Payment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Payment) GetPaymentMethod() PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return PaymentMethod_STRIPE
}

func (x *Payment) GetPaymentDate() string {
	if x != nil {
		return x.PaymentDate
	}
	return ""
}

func (x *Payment) GetPaymentDetails() *PaymentDetails {
	if x != nil {
		return x.PaymentDetails
	}
	return nil
}

func (x *Payment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Payment) GetPaymentLinkExpiry() *durationpb.Duration {
	if x != nil {
		return x.PaymentLinkExpiry
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{2}
}

type PaymentList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payments      []*Payment             `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentList) Reset() {
	*x = PaymentList{}
	mi := &file_proto_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentList) ProtoMessage() {}

func (x *PaymentList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentList.ProtoReflect.Descriptor instead.
func (*PaymentList) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentList) GetPayments() []*Payment {
	if x != nil {
		return x.Payments
	}
	return nil
}

type PaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Payment       *Payment               `protobuf:"bytes,2,opt,name=payment,proto3" json:"payment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	mi := &file_proto_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{4}
}

func (x *PaymentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PaymentResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	PaymentId         string                 `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Status            string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	PaymentMethod     PaymentMethod          `protobuf:"varint,3,opt,name=payment_method,json=paymentMethod,proto3,enum=payment.PaymentMethod" json:"payment_method,omitempty"`
	PaymentDate       string                 `protobuf:"bytes,4,opt,name=payment_date,json=paymentDate,proto3" json:"payment_date,omitempty"`
	PaymentDetails    *PaymentDetails        `protobuf:"bytes,5,opt,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	PaymentLinkExpiry *durationpb.Duration   `protobuf:"bytes,7,opt,name=payment_link_expiry,json=paymentLinkExpiry,proto3" json:"payment_link_expiry,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	mi := &file_proto_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{5}
}

func (x *PaymentRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *PaymentRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PaymentRequest) GetPaymentMethod() PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return PaymentMethod_STRIPE
}

func (x *PaymentRequest) GetPaymentDate() string {
	if x != nil {
		return x.PaymentDate
	}
	return ""
}

func (x *PaymentRequest) GetPaymentDetails() *PaymentDetails {
	if x != nil {
		return x.PaymentDetails
	}
	return nil
}

func (x *PaymentRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PaymentRequest) GetPaymentLinkExpiry() *durationpb.Duration {
	if x != nil {
		return x.PaymentLinkExpiry
	}
	return nil
}

var File_proto_payment_proto protoreflect.FileDescriptor

var file_proto_payment_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8e, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x11, 0x70, 0x61, 0x79, 0x70, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x70, 0x61,
	0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x50, 0x61, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a,
	0x16, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x63, 0x61, 0x73, 0x68, 0x5f, 0x6f, 0x6e, 0x5f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x61, 0x73, 0x68, 0x4f, 0x6e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0xea, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3b, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x57, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xf1, 0x02, 0x0a,
	0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x2a, 0x5e, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x41, 0x59, 0x50, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x4f, 0x4f,
	0x47, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45,
	0x44, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41,
	0x53, 0x48, 0x5f, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x04,
	0x32, 0x91, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_payment_proto_rawDescOnce sync.Once
	file_proto_payment_proto_rawDescData []byte
)

func file_proto_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_payment_proto_rawDesc), len(file_proto_payment_proto_rawDesc)))
	})
	return file_proto_payment_proto_rawDescData
}

var file_proto_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_payment_proto_goTypes = []any{
	(PaymentMethod)(0),            // 0: payment.PaymentMethod
	(*PaymentDetails)(nil),        // 1: payment.PaymentDetails
	(*Payment)(nil),               // 2: payment.Payment
	(*Empty)(nil),                 // 3: payment.Empty
	(*PaymentList)(nil),           // 4: payment.PaymentList
	(*PaymentResponse)(nil),       // 5: payment.PaymentResponse
	(*PaymentRequest)(nil),        // 6: payment.PaymentRequest
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_proto_payment_proto_depIdxs = []int32{
	0,  // 0: payment.Payment.payment_method:type_name -> payment.PaymentMethod
	1,  // 1: payment.Payment.payment_details:type_name -> payment.PaymentDetails
	7,  // 2: payment.Payment.created_at:type_name -> google.protobuf.Timestamp
	8,  // 3: payment.Payment.payment_link_expiry:type_name -> google.protobuf.Duration
	2,  // 4: payment.PaymentList.payments:type_name -> payment.Payment
	2,  // 5: payment.PaymentResponse.payment:type_name -> payment.Payment
	0,  // 6: payment.PaymentRequest.payment_method:type_name -> payment.PaymentMethod
	1,  // 7: payment.PaymentRequest.payment_details:type_name -> payment.PaymentDetails
	7,  // 8: payment.PaymentRequest.created_at:type_name -> google.protobuf.Timestamp
	8,  // 9: payment.PaymentRequest.payment_link_expiry:type_name -> google.protobuf.Duration
	6,  // 10: payment.PaymentService.CreatePayment:input_type -> payment.PaymentRequest
	3,  // 11: payment.PaymentService.GetPayment:input_type -> payment.Empty
	5,  // 12: payment.PaymentService.UpdatePayment:input_type -> payment.PaymentResponse
	6,  // 13: payment.PaymentService.DeletePayment:input_type -> payment.PaymentRequest
	5,  // 14: payment.PaymentService.CreatePayment:output_type -> payment.PaymentResponse
	4,  // 15: payment.PaymentService.GetPayment:output_type -> payment.PaymentList
	5,  // 16: payment.PaymentService.UpdatePayment:output_type -> payment.PaymentResponse
	5,  // 17: payment.PaymentService.DeletePayment:output_type -> payment.PaymentResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_payment_proto_init() }
func file_proto_payment_proto_init() {
	if File_proto_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_payment_proto_rawDesc), len(file_proto_payment_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_proto_depIdxs,
		EnumInfos:         file_proto_payment_proto_enumTypes,
		MessageInfos:      file_proto_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_proto = out.File
	file_proto_payment_proto_goTypes = nil
	file_proto_payment_proto_depIdxs = nil
}
