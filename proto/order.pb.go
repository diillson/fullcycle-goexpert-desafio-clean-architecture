// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: internal/grpc/proto/order.proto

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

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64   `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount     float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Status     string  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_internal_grpc_proto_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreateOrderRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateOrderRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId int64   `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount     float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status     string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt  string  `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_internal_grpc_proto_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Order) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ListOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_internal_grpc_proto_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_order_proto_rawDescGZIP(), []int{2}
}

type ListOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_internal_grpc_proto_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_proto_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_internal_grpc_proto_order_proto protoreflect.FileDescriptor

var file_internal_grpc_proto_order_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x87, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x32, 0x8d, 0x01, 0x0a, 0x0c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpc_proto_order_proto_rawDescOnce sync.Once
	file_internal_grpc_proto_order_proto_rawDescData = file_internal_grpc_proto_order_proto_rawDesc
)

func file_internal_grpc_proto_order_proto_rawDescGZIP() []byte {
	file_internal_grpc_proto_order_proto_rawDescOnce.Do(func() {
		file_internal_grpc_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_proto_order_proto_rawDescData)
	})
	return file_internal_grpc_proto_order_proto_rawDescData
}

var file_internal_grpc_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_grpc_proto_order_proto_goTypes = []any{
	(*CreateOrderRequest)(nil), // 0: order.CreateOrderRequest
	(*Order)(nil),              // 1: order.Order
	(*ListOrdersRequest)(nil),  // 2: order.ListOrdersRequest
	(*ListOrdersResponse)(nil), // 3: order.ListOrdersResponse
}
var file_internal_grpc_proto_order_proto_depIdxs = []int32{
	1, // 0: order.ListOrdersResponse.orders:type_name -> order.Order
	2, // 1: order.OrderService.ListOrders:input_type -> order.ListOrdersRequest
	0, // 2: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	3, // 3: order.OrderService.ListOrders:output_type -> order.ListOrdersResponse
	1, // 4: order.OrderService.CreateOrder:output_type -> order.Order
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_grpc_proto_order_proto_init() }
func file_internal_grpc_proto_order_proto_init() {
	if File_internal_grpc_proto_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_grpc_proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_proto_order_proto_goTypes,
		DependencyIndexes: file_internal_grpc_proto_order_proto_depIdxs,
		MessageInfos:      file_internal_grpc_proto_order_proto_msgTypes,
	}.Build()
	File_internal_grpc_proto_order_proto = out.File
	file_internal_grpc_proto_order_proto_rawDesc = nil
	file_internal_grpc_proto_order_proto_goTypes = nil
	file_internal_grpc_proto_order_proto_depIdxs = nil
}
