// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: gRPC/interface.proto

package auction

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

type Outcomes int32

const (
	Outcomes_FAIL    Outcomes = 0
	Outcomes_SUCCESS Outcomes = 1
)

// Enum value maps for Outcomes.
var (
	Outcomes_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
	}
	Outcomes_value = map[string]int32{
		"FAIL":    0,
		"SUCCESS": 1,
	}
)

func (x Outcomes) Enum() *Outcomes {
	p := new(Outcomes)
	*p = x
	return p
}

func (x Outcomes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Outcomes) Descriptor() protoreflect.EnumDescriptor {
	return file_gRPC_interface_proto_enumTypes[0].Descriptor()
}

func (Outcomes) Type() protoreflect.EnumType {
	return &file_gRPC_interface_proto_enumTypes[0]
}

func (x Outcomes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Outcomes.Descriptor instead.
func (Outcomes) EnumDescriptor() ([]byte, []int) {
	return file_gRPC_interface_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_gRPC_interface_proto_rawDescGZIP(), []int{0}
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Bid  int32  `protobuf:"varint,2,opt,name=bid,proto3" json:"bid,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_gRPC_interface_proto_rawDescGZIP(), []int{1}
}

func (x *BidRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *BidRequest) GetBid() int32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

type ClientReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ClientReply) Reset() {
	*x = ClientReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReply) ProtoMessage() {}

func (x *ClientReply) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReply.ProtoReflect.Descriptor instead.
func (*ClientReply) Descriptor() ([]byte, []int) {
	return file_gRPC_interface_proto_rawDescGZIP(), []int{2}
}

func (x *ClientReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OutcomeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outcome Outcomes `protobuf:"varint,1,opt,name=outcome,proto3,enum=auction.Outcomes" json:"outcome,omitempty"`
}

func (x *OutcomeReply) Reset() {
	*x = OutcomeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutcomeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutcomeReply) ProtoMessage() {}

func (x *OutcomeReply) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutcomeReply.ProtoReflect.Descriptor instead.
func (*OutcomeReply) Descriptor() ([]byte, []int) {
	return file_gRPC_interface_proto_rawDescGZIP(), []int{3}
}

func (x *OutcomeReply) GetOutcome() Outcomes {
	if x != nil {
		return x.Outcome
	}
	return Outcomes_FAIL
}

type ResultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBid int32  `protobuf:"varint,1,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	User       string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	TimeLeft   int32  `protobuf:"varint,3,opt,name=timeLeft,proto3" json:"timeLeft,omitempty"`
}

func (x *ResultReply) Reset() {
	*x = ResultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_interface_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultReply) ProtoMessage() {}

func (x *ResultReply) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_interface_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultReply.ProtoReflect.Descriptor instead.
func (*ResultReply) Descriptor() ([]byte, []int) {
	return file_gRPC_interface_proto_rawDescGZIP(), []int{4}
}

func (x *ResultReply) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *ResultReply) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ResultReply) GetTimeLeft() int32 {
	if x != nil {
		return x.TimeLeft
	}
	return 0
}

var File_gRPC_interface_proto protoreflect.FileDescriptor

var file_gRPC_interface_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0a, 0x42, 0x69,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x69, 0x64, 0x22, 0x27,
	0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x52, 0x07, 0x6f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c,
	0x65, 0x66, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c,
	0x65, 0x66, 0x74, 0x2a, 0x21, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x32, 0xd3, 0x02, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x65, 0x64, 0x65, 0x72, 0x69, 0x6b, 0x67, 0x61, 0x6e, 0x74, 0x72,
	0x69, 0x69, 0x73, 0x2f, 0x44, 0x49, 0x53, 0x59, 0x53, 0x2d, 0x45, 0x58, 0x41, 0x4d, 0x32, 0x30,
	0x32, 0x31, 0x3b, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gRPC_interface_proto_rawDescOnce sync.Once
	file_gRPC_interface_proto_rawDescData = file_gRPC_interface_proto_rawDesc
)

func file_gRPC_interface_proto_rawDescGZIP() []byte {
	file_gRPC_interface_proto_rawDescOnce.Do(func() {
		file_gRPC_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_interface_proto_rawDescData)
	})
	return file_gRPC_interface_proto_rawDescData
}

var file_gRPC_interface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gRPC_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gRPC_interface_proto_goTypes = []interface{}{
	(Outcomes)(0),        // 0: auction.outcomes
	(*Request)(nil),      // 1: auction.Request
	(*BidRequest)(nil),   // 2: auction.BidRequest
	(*ClientReply)(nil),  // 3: auction.ClientReply
	(*OutcomeReply)(nil), // 4: auction.OutcomeReply
	(*ResultReply)(nil),  // 5: auction.ResultReply
}
var file_gRPC_interface_proto_depIdxs = []int32{
	0, // 0: auction.OutcomeReply.outcome:type_name -> auction.outcomes
	2, // 1: auction.Auction.bid:input_type -> auction.BidRequest
	1, // 2: auction.Auction.result:input_type -> auction.Request
	1, // 3: auction.Auction.reset:input_type -> auction.Request
	2, // 4: auction.Auction.serverBid:input_type -> auction.BidRequest
	1, // 5: auction.Auction.serverResult:input_type -> auction.Request
	1, // 6: auction.Auction.ServerReset:input_type -> auction.Request
	3, // 7: auction.Auction.bid:output_type -> auction.ClientReply
	3, // 8: auction.Auction.result:output_type -> auction.ClientReply
	3, // 9: auction.Auction.reset:output_type -> auction.ClientReply
	4, // 10: auction.Auction.serverBid:output_type -> auction.OutcomeReply
	5, // 11: auction.Auction.serverResult:output_type -> auction.ResultReply
	4, // 12: auction.Auction.ServerReset:output_type -> auction.OutcomeReply
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gRPC_interface_proto_init() }
func file_gRPC_interface_proto_init() {
	if File_gRPC_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPC_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_gRPC_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidRequest); i {
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
		file_gRPC_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReply); i {
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
		file_gRPC_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutcomeReply); i {
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
		file_gRPC_interface_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultReply); i {
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
			RawDescriptor: file_gRPC_interface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_interface_proto_goTypes,
		DependencyIndexes: file_gRPC_interface_proto_depIdxs,
		EnumInfos:         file_gRPC_interface_proto_enumTypes,
		MessageInfos:      file_gRPC_interface_proto_msgTypes,
	}.Build()
	File_gRPC_interface_proto = out.File
	file_gRPC_interface_proto_rawDesc = nil
	file_gRPC_interface_proto_goTypes = nil
	file_gRPC_interface_proto_depIdxs = nil
}
