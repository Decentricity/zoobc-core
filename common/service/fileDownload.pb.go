// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/fileDownload.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	model "github.com/zoobc/zoobc-core/common/model"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("service/fileDownload.proto", fileDescriptor_fde5bac5b80ffd69) }

var fileDescriptor_fde5bac5b80ffd69 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0xcb, 0xcc, 0x49, 0x75, 0xc9, 0x2f, 0xcf, 0xcb, 0xc9, 0x4f, 0x4c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0xca, 0x49, 0x49, 0xe4, 0xe6, 0xa7, 0xa4,
	0xe6, 0x60, 0x51, 0x22, 0x25, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58, 0x90, 0xa9,
	0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x35, 0xaa, 0xe5,
	0x12, 0x76, 0x43, 0xd2, 0x13, 0x0c, 0x31, 0x4e, 0x28, 0x8d, 0x8b, 0x07, 0x59, 0x58, 0x48, 0x4a,
	0x0f, 0x6c, 0xbe, 0x1e, 0xb2, 0x60, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94, 0x34, 0x56,
	0xb9, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x25, 0xd9, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x89, 0x0b,
	0x89, 0xea, 0x97, 0x19, 0x82, 0x5d, 0xa7, 0x8f, 0xac, 0xcc, 0x49, 0x27, 0x4a, 0x2b, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x2a, 0x3f, 0x3f, 0x29, 0x19, 0x42, 0xea,
	0x26, 0xe7, 0x17, 0xa5, 0xea, 0x27, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0xe9, 0x43, 0x3d, 0x99, 0xc4,
	0x06, 0x76, 0xb3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xaf, 0x05, 0xa4, 0x12, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileDownloadServiceClient is the client API for FileDownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileDownloadServiceClient interface {
	FileDownload(ctx context.Context, in *model.FileDownloadRequest, opts ...grpc.CallOption) (*model.FileDownloadResponse, error)
}

type fileDownloadServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileDownloadServiceClient(cc *grpc.ClientConn) FileDownloadServiceClient {
	return &fileDownloadServiceClient{cc}
}

func (c *fileDownloadServiceClient) FileDownload(ctx context.Context, in *model.FileDownloadRequest, opts ...grpc.CallOption) (*model.FileDownloadResponse, error) {
	out := new(model.FileDownloadResponse)
	err := c.cc.Invoke(ctx, "/service.FileDownloadService/FileDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileDownloadServiceServer is the server API for FileDownloadService service.
type FileDownloadServiceServer interface {
	FileDownload(context.Context, *model.FileDownloadRequest) (*model.FileDownloadResponse, error)
}

// UnimplementedFileDownloadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileDownloadServiceServer struct {
}

func (*UnimplementedFileDownloadServiceServer) FileDownload(ctx context.Context, req *model.FileDownloadRequest) (*model.FileDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileDownload not implemented")
}

func RegisterFileDownloadServiceServer(s *grpc.Server, srv FileDownloadServiceServer) {
	s.RegisterService(&_FileDownloadService_serviceDesc, srv)
}

func _FileDownloadService_FileDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.FileDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileDownloadServiceServer).FileDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileDownloadService/FileDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileDownloadServiceServer).FileDownload(ctx, req.(*model.FileDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileDownloadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.FileDownloadService",
	HandlerType: (*FileDownloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FileDownload",
			Handler:    _FileDownloadService_FileDownload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/fileDownload.proto",
}
