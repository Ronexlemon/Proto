// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: coffee_proto.proto

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price    int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity int64  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_coffee_proto_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_coffee_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type MenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MenuRequest) Reset() {
	*x = MenuRequest{}
	mi := &file_coffee_proto_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuRequest) ProtoMessage() {}

func (x *MenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuRequest.ProtoReflect.Descriptor instead.
func (*MenuRequest) Descriptor() ([]byte, []int) {
	return file_coffee_proto_proto_rawDescGZIP(), []int{1}
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*Item `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"` //repeated its an List
}

func (x *Menu) Reset() {
	*x = Menu{}
	mi := &file_coffee_proto_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_coffee_proto_proto_rawDescGZIP(), []int{2}
}

func (x *Menu) GetItem() []*Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_coffee_proto_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_proto_msgTypes[3]
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
	return file_coffee_proto_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	mi := &file_coffee_proto_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_coffee_proto_proto_rawDescGZIP(), []int{4}
}

func (x *Receipt) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderStatus) Reset() {
	*x = OrderStatus{}
	mi := &file_coffee_proto_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatus) ProtoMessage() {}

func (x *OrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_proto_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatus.ProtoReflect.Descriptor instead.
func (*OrderStatus) Descriptor() ([]byte, []int) {
	return file_coffee_proto_proto_rawDescGZIP(), []int{5}
}

func (x *OrderStatus) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_coffee_proto_proto protoreflect.FileDescriptor

var file_coffee_proto_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x21, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x19, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x22, 0x24, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x77, 0x0a, 0x09, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x53,
	0x68, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0c,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x4d,
	0x65, 0x6e, 0x75, 0x30, 0x01, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x1a, 0x0c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f,
	0x6e, 0x65, 0x78, 0x4c, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coffee_proto_proto_rawDescOnce sync.Once
	file_coffee_proto_proto_rawDescData = file_coffee_proto_proto_rawDesc
)

func file_coffee_proto_proto_rawDescGZIP() []byte {
	file_coffee_proto_proto_rawDescOnce.Do(func() {
		file_coffee_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_coffee_proto_proto_rawDescData)
	})
	return file_coffee_proto_proto_rawDescData
}

var file_coffee_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_coffee_proto_proto_goTypes = []any{
	(*Item)(nil),        // 0: Item
	(*MenuRequest)(nil), // 1: MenuRequest
	(*Menu)(nil),        // 2: Menu
	(*Order)(nil),       // 3: Order
	(*Receipt)(nil),     // 4: Receipt
	(*OrderStatus)(nil), // 5: OrderStatus
}
var file_coffee_proto_proto_depIdxs = []int32{
	0, // 0: Menu.item:type_name -> Item
	0, // 1: Order.items:type_name -> Item
	1, // 2: CoffeShop.GetMenu:input_type -> MenuRequest
	3, // 3: CoffeShop.PlaceOrder:input_type -> Order
	4, // 4: CoffeShop.GetOrderStatus:input_type -> Receipt
	2, // 5: CoffeShop.GetMenu:output_type -> Menu
	4, // 6: CoffeShop.PlaceOrder:output_type -> Receipt
	5, // 7: CoffeShop.GetOrderStatus:output_type -> OrderStatus
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coffee_proto_proto_init() }
func file_coffee_proto_proto_init() {
	if File_coffee_proto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coffee_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coffee_proto_proto_goTypes,
		DependencyIndexes: file_coffee_proto_proto_depIdxs,
		MessageInfos:      file_coffee_proto_proto_msgTypes,
	}.Build()
	File_coffee_proto_proto = out.File
	file_coffee_proto_proto_rawDesc = nil
	file_coffee_proto_proto_goTypes = nil
	file_coffee_proto_proto_depIdxs = nil
}
