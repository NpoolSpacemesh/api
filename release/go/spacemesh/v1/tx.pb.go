// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: spacemesh/v1/tx.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_spacemesh_v1_tx_proto protoreflect.FileDescriptor

var file_spacemesh_v1_tx_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xda, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x64, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_spacemesh_v1_tx_proto_goTypes = []interface{}{
	(*SubmitTransactionRequest)(nil),        // 0: spacemesh.v1.SubmitTransactionRequest
	(*TransactionsStateRequest)(nil),        // 1: spacemesh.v1.TransactionsStateRequest
	(*TransactionsStateStreamRequest)(nil),  // 2: spacemesh.v1.TransactionsStateStreamRequest
	(*SubmitTransactionResponse)(nil),       // 3: spacemesh.v1.SubmitTransactionResponse
	(*TransactionsStateResponse)(nil),       // 4: spacemesh.v1.TransactionsStateResponse
	(*TransactionsStateStreamResponse)(nil), // 5: spacemesh.v1.TransactionsStateStreamResponse
}
var file_spacemesh_v1_tx_proto_depIdxs = []int32{
	0, // 0: spacemesh.v1.TransactionService.SubmitTransaction:input_type -> spacemesh.v1.SubmitTransactionRequest
	1, // 1: spacemesh.v1.TransactionService.TransactionsState:input_type -> spacemesh.v1.TransactionsStateRequest
	2, // 2: spacemesh.v1.TransactionService.TransactionsStateStream:input_type -> spacemesh.v1.TransactionsStateStreamRequest
	3, // 3: spacemesh.v1.TransactionService.SubmitTransaction:output_type -> spacemesh.v1.SubmitTransactionResponse
	4, // 4: spacemesh.v1.TransactionService.TransactionsState:output_type -> spacemesh.v1.TransactionsStateResponse
	5, // 5: spacemesh.v1.TransactionService.TransactionsStateStream:output_type -> spacemesh.v1.TransactionsStateStreamResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_tx_proto_init() }
func file_spacemesh_v1_tx_proto_init() {
	if File_spacemesh_v1_tx_proto != nil {
		return
	}
	file_spacemesh_v1_tx_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_tx_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_tx_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_tx_proto = out.File
	file_spacemesh_v1_tx_proto_rawDesc = nil
	file_spacemesh_v1_tx_proto_goTypes = nil
	file_spacemesh_v1_tx_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	// Submit a new tx to the node for processing. The response
	// TransactionState message contains both the txid of the new tx, as well
	// as whether or not it was admitted into the mempool.
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	// Returns current tx state for one or more txs which indicates if a tx is
	// on the mesh, on its way to the mesh or was rejected and will never get
	// to the mesh
	TransactionsState(ctx context.Context, in *TransactionsStateRequest, opts ...grpc.CallOption) (*TransactionsStateResponse, error)
	// Returns tx state for one or more txs every time the tx state changes for
	// one of these txs
	TransactionsStateStream(ctx context.Context, in *TransactionsStateStreamRequest, opts ...grpc.CallOption) (TransactionService_TransactionsStateStreamClient, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.TransactionService/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) TransactionsState(ctx context.Context, in *TransactionsStateRequest, opts ...grpc.CallOption) (*TransactionsStateResponse, error) {
	out := new(TransactionsStateResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.TransactionService/TransactionsState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) TransactionsStateStream(ctx context.Context, in *TransactionsStateStreamRequest, opts ...grpc.CallOption) (TransactionService_TransactionsStateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TransactionService_serviceDesc.Streams[0], "/spacemesh.v1.TransactionService/TransactionsStateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionServiceTransactionsStateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransactionService_TransactionsStateStreamClient interface {
	Recv() (*TransactionsStateStreamResponse, error)
	grpc.ClientStream
}

type transactionServiceTransactionsStateStreamClient struct {
	grpc.ClientStream
}

func (x *transactionServiceTransactionsStateStreamClient) Recv() (*TransactionsStateStreamResponse, error) {
	m := new(TransactionsStateStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	// Submit a new tx to the node for processing. The response
	// TransactionState message contains both the txid of the new tx, as well
	// as whether or not it was admitted into the mempool.
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	// Returns current tx state for one or more txs which indicates if a tx is
	// on the mesh, on its way to the mesh or was rejected and will never get
	// to the mesh
	TransactionsState(context.Context, *TransactionsStateRequest) (*TransactionsStateResponse, error)
	// Returns tx state for one or more txs every time the tx state changes for
	// one of these txs
	TransactionsStateStream(*TransactionsStateStreamRequest, TransactionService_TransactionsStateStreamServer) error
}

// UnimplementedTransactionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (*UnimplementedTransactionServiceServer) SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (*UnimplementedTransactionServiceServer) TransactionsState(context.Context, *TransactionsStateRequest) (*TransactionsStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionsState not implemented")
}
func (*UnimplementedTransactionServiceServer) TransactionsStateStream(*TransactionsStateStreamRequest, TransactionService_TransactionsStateStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TransactionsStateStream not implemented")
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.TransactionService/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_TransactionsState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionsStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).TransactionsState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.TransactionService/TransactionsState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).TransactionsState(ctx, req.(*TransactionsStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_TransactionsStateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionsStateStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionServiceServer).TransactionsStateStream(m, &transactionServiceTransactionsStateStreamServer{stream})
}

type TransactionService_TransactionsStateStreamServer interface {
	Send(*TransactionsStateStreamResponse) error
	grpc.ServerStream
}

type transactionServiceTransactionsStateStreamServer struct {
	grpc.ServerStream
}

func (x *transactionServiceTransactionsStateStreamServer) Send(m *TransactionsStateStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v1.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _TransactionService_SubmitTransaction_Handler,
		},
		{
			MethodName: "TransactionsState",
			Handler:    _TransactionService_TransactionsState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransactionsStateStream",
			Handler:       _TransactionService_TransactionsStateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v1/tx.proto",
}
