// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: spacemesh/v1/gateway.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_spacemesh_v1_gateway_proto protoreflect.FileDescriptor

var file_spacemesh_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6a, 0x0a, 0x0e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x65, 0x74, 0x12,
	0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_spacemesh_v1_gateway_proto_goTypes = []interface{}{
	(*BroadcastPoetRequest)(nil),  // 0: spacemesh.v1.BroadcastPoetRequest
	(*BroadcastPoetResponse)(nil), // 1: spacemesh.v1.BroadcastPoetResponse
}
var file_spacemesh_v1_gateway_proto_depIdxs = []int32{
	0, // 0: spacemesh.v1.GatewayService.BroadcastPoet:input_type -> spacemesh.v1.BroadcastPoetRequest
	1, // 1: spacemesh.v1.GatewayService.BroadcastPoet:output_type -> spacemesh.v1.BroadcastPoetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_gateway_proto_init() }
func file_spacemesh_v1_gateway_proto_init() {
	if File_spacemesh_v1_gateway_proto != nil {
		return
	}
	file_spacemesh_v1_gateway_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_gateway_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_gateway_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_gateway_proto = out.File
	file_spacemesh_v1_gateway_proto_rawDesc = nil
	file_spacemesh_v1_gateway_proto_goTypes = nil
	file_spacemesh_v1_gateway_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	// Submit a poet data packet to the network to broadcast
	BroadcastPoet(ctx context.Context, in *BroadcastPoetRequest, opts ...grpc.CallOption) (*BroadcastPoetResponse, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) BroadcastPoet(ctx context.Context, in *BroadcastPoetRequest, opts ...grpc.CallOption) (*BroadcastPoetResponse, error) {
	out := new(BroadcastPoetResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.GatewayService/BroadcastPoet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	// Submit a poet data packet to the network to broadcast
	BroadcastPoet(context.Context, *BroadcastPoetRequest) (*BroadcastPoetResponse, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) BroadcastPoet(context.Context, *BroadcastPoetRequest) (*BroadcastPoetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastPoet not implemented")
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_BroadcastPoet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastPoetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).BroadcastPoet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.GatewayService/BroadcastPoet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).BroadcastPoet(ctx, req.(*BroadcastPoetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v1.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastPoet",
			Handler:    _GatewayService_BroadcastPoet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v1/gateway.proto",
}
