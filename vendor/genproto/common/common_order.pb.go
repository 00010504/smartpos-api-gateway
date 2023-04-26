// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: common_order.proto

package common

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

type OrderCopyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string                        `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Items     []*CreateOrderItemCopyRequest `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Request   *Request                      `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	CompanyId string                        `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ShopId    string                        `protobuf:"bytes,5,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	ShopName  string                        `protobuf:"bytes,6,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
}

func (x *OrderCopyRequest) Reset() {
	*x = OrderCopyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCopyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCopyRequest) ProtoMessage() {}

func (x *OrderCopyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCopyRequest.ProtoReflect.Descriptor instead.
func (*OrderCopyRequest) Descriptor() ([]byte, []int) {
	return file_common_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderCopyRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderCopyRequest) GetItems() []*CreateOrderItemCopyRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderCopyRequest) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *OrderCopyRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *OrderCopyRequest) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *OrderCopyRequest) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

type CreateOrderItemCopyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string  `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Value       float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	SupplyPrice float32 `protobuf:"fixed32,3,opt,name=supply_price,json=supplyPrice,proto3" json:"supply_price,omitempty"`
	SalePrice   float32 `protobuf:"fixed32,4,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
	Name        string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateOrderItemCopyRequest) Reset() {
	*x = CreateOrderItemCopyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderItemCopyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderItemCopyRequest) ProtoMessage() {}

func (x *CreateOrderItemCopyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderItemCopyRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderItemCopyRequest) Descriptor() ([]byte, []int) {
	return file_common_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderItemCopyRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateOrderItemCopyRequest) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CreateOrderItemCopyRequest) GetSupplyPrice() float32 {
	if x != nil {
		return x.SupplyPrice
	}
	return 0
}

func (x *CreateOrderItemCopyRequest) GetSalePrice() float32 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *CreateOrderItemCopyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_common_order_proto protoreflect.FileDescriptor

var file_common_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x70,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_order_proto_rawDescOnce sync.Once
	file_common_order_proto_rawDescData = file_common_order_proto_rawDesc
)

func file_common_order_proto_rawDescGZIP() []byte {
	file_common_order_proto_rawDescOnce.Do(func() {
		file_common_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_order_proto_rawDescData)
	})
	return file_common_order_proto_rawDescData
}

var file_common_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_order_proto_goTypes = []interface{}{
	(*OrderCopyRequest)(nil),           // 0: OrderCopyRequest
	(*CreateOrderItemCopyRequest)(nil), // 1: CreateOrderItemCopyRequest
	(*Request)(nil),                    // 2: Request
}
var file_common_order_proto_depIdxs = []int32{
	1, // 0: OrderCopyRequest.items:type_name -> CreateOrderItemCopyRequest
	2, // 1: OrderCopyRequest.request:type_name -> Request
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_order_proto_init() }
func file_common_order_proto_init() {
	if File_common_order_proto != nil {
		return
	}
	file_request_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCopyRequest); i {
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
		file_common_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderItemCopyRequest); i {
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
			RawDescriptor: file_common_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_order_proto_goTypes,
		DependencyIndexes: file_common_order_proto_depIdxs,
		MessageInfos:      file_common_order_proto_msgTypes,
	}.Build()
	File_common_order_proto = out.File
	file_common_order_proto_rawDesc = nil
	file_common_order_proto_goTypes = nil
	file_common_order_proto_depIdxs = nil
}