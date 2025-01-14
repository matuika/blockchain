



package v2alpha1

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)


var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_ReflectionService_GetAuthnDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, client ReflectionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAuthnDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetAuthnDescriptor(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ReflectionService_GetAuthnDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, server ReflectionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetAuthnDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetAuthnDescriptor(ctx, &protoReq)
	return msg, metadata, err

}

func request_ReflectionService_GetChainDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, client ReflectionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetChainDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetChainDescriptor(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ReflectionService_GetChainDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, server ReflectionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetChainDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetChainDescriptor(ctx, &protoReq)
	return msg, metadata, err

}

func request_ReflectionService_GetCodecDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, client ReflectionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetCodecDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetCodecDescriptor(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ReflectionService_GetCodecDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, server ReflectionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetCodecDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetCodecDescriptor(ctx, &protoReq)
	return msg, metadata, err

}

func request_ReflectionService_GetConfigurationDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, client ReflectionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetConfigurationDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetConfigurationDescriptor(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ReflectionService_GetConfigurationDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, server ReflectionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetConfigurationDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetConfigurationDescriptor(ctx, &protoReq)
	return msg, metadata, err

}

func request_ReflectionService_GetQueryServicesDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, client ReflectionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetQueryServicesDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetQueryServicesDescriptor(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ReflectionService_GetQueryServicesDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, server ReflectionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetQueryServicesDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetQueryServicesDescriptor(ctx, &protoReq)
	return msg, metadata, err

}

func request_ReflectionService_GetTxDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, client ReflectionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetTxDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetTxDescriptor(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ReflectionService_GetTxDescriptor_0(ctx context.Context, marshaler runtime.Marshaler, server ReflectionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetTxDescriptorRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetTxDescriptor(ctx, &protoReq)
	return msg, metadata, err

}





func RegisterReflectionServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ReflectionServiceServer) error {

	mux.Handle("GET", pattern_ReflectionService_GetAuthnDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ReflectionService_GetAuthnDescriptor_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetAuthnDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetChainDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ReflectionService_GetChainDescriptor_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetChainDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetCodecDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ReflectionService_GetCodecDescriptor_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetCodecDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetConfigurationDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ReflectionService_GetConfigurationDescriptor_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetConfigurationDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetQueryServicesDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ReflectionService_GetQueryServicesDescriptor_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetQueryServicesDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetTxDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ReflectionService_GetTxDescriptor_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetTxDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}



func RegisterReflectionServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterReflectionServiceHandler(ctx, mux, conn)
}



func RegisterReflectionServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterReflectionServiceHandlerClient(ctx, mux, NewReflectionServiceClient(conn))
}






func RegisterReflectionServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ReflectionServiceClient) error {

	mux.Handle("GET", pattern_ReflectionService_GetAuthnDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ReflectionService_GetAuthnDescriptor_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetAuthnDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetChainDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ReflectionService_GetChainDescriptor_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetChainDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetCodecDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ReflectionService_GetCodecDescriptor_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetCodecDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetConfigurationDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ReflectionService_GetConfigurationDescriptor_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetConfigurationDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetQueryServicesDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ReflectionService_GetQueryServicesDescriptor_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetQueryServicesDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ReflectionService_GetTxDescriptor_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ReflectionService_GetTxDescriptor_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ReflectionService_GetTxDescriptor_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ReflectionService_GetAuthnDescriptor_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"cosmos", "base", "reflection", "v1beta1", "app_descriptor", "authn"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_ReflectionService_GetChainDescriptor_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"cosmos", "base", "reflection", "v1beta1", "app_descriptor", "chain"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_ReflectionService_GetCodecDescriptor_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"cosmos", "base", "reflection", "v1beta1", "app_descriptor", "codec"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_ReflectionService_GetConfigurationDescriptor_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"cosmos", "base", "reflection", "v1beta1", "app_descriptor", "configuration"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_ReflectionService_GetQueryServicesDescriptor_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"cosmos", "base", "reflection", "v1beta1", "app_descriptor", "query_services"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_ReflectionService_GetTxDescriptor_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"cosmos", "base", "reflection", "v1beta1", "app_descriptor", "tx_descriptor"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_ReflectionService_GetAuthnDescriptor_0 = runtime.ForwardResponseMessage

	forward_ReflectionService_GetChainDescriptor_0 = runtime.ForwardResponseMessage

	forward_ReflectionService_GetCodecDescriptor_0 = runtime.ForwardResponseMessage

	forward_ReflectionService_GetConfigurationDescriptor_0 = runtime.ForwardResponseMessage

	forward_ReflectionService_GetQueryServicesDescriptor_0 = runtime.ForwardResponseMessage

	forward_ReflectionService_GetTxDescriptor_0 = runtime.ForwardResponseMessage
)
