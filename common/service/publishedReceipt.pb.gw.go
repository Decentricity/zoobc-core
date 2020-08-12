// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: service/publishedReceipt.proto

/*
Package service is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package service

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"github.com/zoobc/zoobc-core/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

var (
	filter_PublishedReceiptService_GetPublishedReceipts_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_PublishedReceiptService_GetPublishedReceipts_0(ctx context.Context, marshaler runtime.Marshaler, client PublishedReceiptServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq model.GetPublishedReceiptsRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_PublishedReceiptService_GetPublishedReceipts_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetPublishedReceipts(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_PublishedReceiptService_GetPublishedReceipts_0(ctx context.Context, marshaler runtime.Marshaler, server PublishedReceiptServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq model.GetPublishedReceiptsRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_PublishedReceiptService_GetPublishedReceipts_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetPublishedReceipts(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterPublishedReceiptServiceHandlerServer registers the http handlers for service PublishedReceiptService to "mux".
// UnaryRPC     :call PublishedReceiptServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterPublishedReceiptServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server PublishedReceiptServiceServer) error {

	mux.Handle("GET", pattern_PublishedReceiptService_GetPublishedReceipts_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PublishedReceiptService_GetPublishedReceipts_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PublishedReceiptService_GetPublishedReceipts_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterPublishedReceiptServiceHandlerFromEndpoint is same as RegisterPublishedReceiptServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPublishedReceiptServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPublishedReceiptServiceHandler(ctx, mux, conn)
}

// RegisterPublishedReceiptServiceHandler registers the http handlers for service PublishedReceiptService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPublishedReceiptServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPublishedReceiptServiceHandlerClient(ctx, mux, NewPublishedReceiptServiceClient(conn))
}

// RegisterPublishedReceiptServiceHandlerClient registers the http handlers for service PublishedReceiptService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PublishedReceiptServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PublishedReceiptServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PublishedReceiptServiceClient" to call the correct interceptors.
func RegisterPublishedReceiptServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PublishedReceiptServiceClient) error {

	mux.Handle("GET", pattern_PublishedReceiptService_GetPublishedReceipts_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PublishedReceiptService_GetPublishedReceipts_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PublishedReceiptService_GetPublishedReceipts_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PublishedReceiptService_GetPublishedReceipts_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "publishedReceipt", "GetPublishedReceipts"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_PublishedReceiptService_GetPublishedReceipts_0 = runtime.ForwardResponseMessage
)