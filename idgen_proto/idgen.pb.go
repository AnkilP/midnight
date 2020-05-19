// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.6.1
// source: idgen.proto

package idgen_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Idrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp    *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ComputerName string               `protobuf:"bytes,2,opt,name=computer_name,json=computerName,proto3" json:"computer_name,omitempty"`
}

func (x *Idrequest) Reset() {
	*x = Idrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idgen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Idrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idrequest) ProtoMessage() {}

func (x *Idrequest) ProtoReflect() protoreflect.Message {
	mi := &file_idgen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idrequest.ProtoReflect.Descriptor instead.
func (*Idrequest) Descriptor() ([]byte, []int) {
	return file_idgen_proto_rawDescGZIP(), []int{0}
}

func (x *Idrequest) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Idrequest) GetComputerName() string {
	if x != nil {
		return x.ComputerName
	}
	return ""
}

type Idresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
}

func (x *Idresponse) Reset() {
	*x = Idresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idgen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Idresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idresponse) ProtoMessage() {}

func (x *Idresponse) ProtoReflect() protoreflect.Message {
	mi := &file_idgen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idresponse.ProtoReflect.Descriptor instead.
func (*Idresponse) Descriptor() ([]byte, []int) {
	return file_idgen_proto_rawDescGZIP(), []int{1}
}

func (x *Idresponse) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

var File_idgen_proto protoreflect.FileDescriptor

var file_idgen_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69,
	0x64, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x09, 0x69,
	0x64, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x69, 0x64, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x32, 0x4e, 0x0a, 0x09, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x69, 0x64, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x69, 0x64, 0x67, 0x65, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x69, 0x64, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69,
	0x64, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_idgen_proto_rawDescOnce sync.Once
	file_idgen_proto_rawDescData = file_idgen_proto_rawDesc
)

func file_idgen_proto_rawDescGZIP() []byte {
	file_idgen_proto_rawDescOnce.Do(func() {
		file_idgen_proto_rawDescData = protoimpl.X.CompressGZIP(file_idgen_proto_rawDescData)
	})
	return file_idgen_proto_rawDescData
}

var file_idgen_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_idgen_proto_goTypes = []interface{}{
	(*Idrequest)(nil),           // 0: idgen_proto.idrequest
	(*Idresponse)(nil),          // 1: idgen_proto.idresponse
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_idgen_proto_depIdxs = []int32{
	2, // 0: idgen_proto.idrequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: idgen_proto.Broadcast.idrequestserve:input_type -> idgen_proto.idrequest
	1, // 2: idgen_proto.Broadcast.idrequestserve:output_type -> idgen_proto.idresponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idgen_proto_init() }
func file_idgen_proto_init() {
	if File_idgen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idgen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Idrequest); i {
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
		file_idgen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Idresponse); i {
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
			RawDescriptor: file_idgen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idgen_proto_goTypes,
		DependencyIndexes: file_idgen_proto_depIdxs,
		MessageInfos:      file_idgen_proto_msgTypes,
	}.Build()
	File_idgen_proto = out.File
	file_idgen_proto_rawDesc = nil
	file_idgen_proto_goTypes = nil
	file_idgen_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BroadcastClient is the client API for Broadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BroadcastClient interface {
	Idrequestserve(ctx context.Context, in *Idrequest, opts ...grpc.CallOption) (*Idresponse, error)
}

type broadcastClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastClient(cc grpc.ClientConnInterface) BroadcastClient {
	return &broadcastClient{cc}
}

func (c *broadcastClient) Idrequestserve(ctx context.Context, in *Idrequest, opts ...grpc.CallOption) (*Idresponse, error) {
	out := new(Idresponse)
	err := c.cc.Invoke(ctx, "/idgen_proto.Broadcast/idrequestserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServer is the server API for Broadcast service.
type BroadcastServer interface {
	Idrequestserve(context.Context, *Idrequest) (*Idresponse, error)
}

// UnimplementedBroadcastServer can be embedded to have forward compatible implementations.
type UnimplementedBroadcastServer struct {
}

func (*UnimplementedBroadcastServer) Idrequestserve(context.Context, *Idrequest) (*Idresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Idrequestserve not implemented")
}

func RegisterBroadcastServer(s *grpc.Server, srv BroadcastServer) {
	s.RegisterService(&_Broadcast_serviceDesc, srv)
}

func _Broadcast_Idrequestserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Idrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).Idrequestserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idgen_proto.Broadcast/Idrequestserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).Idrequestserve(ctx, req.(*Idrequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idgen_proto.Broadcast",
	HandlerType: (*BroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "idrequestserve",
			Handler:    _Broadcast_Idrequestserve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idgen.proto",
}
