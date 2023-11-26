// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: apis/v1/health.proto

package apisv1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	v1 "github.com/go-zen-chu/connect-go-sample/pkg/apis/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// HealthServiceName is the fully-qualified name of the HealthService service.
	HealthServiceName = "apis.v1.HealthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthServiceHealthProcedure is the fully-qualified name of the HealthService's Health RPC.
	HealthServiceHealthProcedure = "/apis.v1.HealthService/Health"
)

// HealthServiceClient is a client for the apis.v1.HealthService service.
type HealthServiceClient interface {
	Health(context.Context, *connect.Request[v1.HealthRequest]) (*connect.Response[v1.HealthResponse], error)
}

// NewHealthServiceClient constructs a client for the apis.v1.HealthService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) HealthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &healthServiceClient{
		health: connect.NewClient[v1.HealthRequest, v1.HealthResponse](
			httpClient,
			baseURL+HealthServiceHealthProcedure,
			opts...,
		),
	}
}

// healthServiceClient implements HealthServiceClient.
type healthServiceClient struct {
	health *connect.Client[v1.HealthRequest, v1.HealthResponse]
}

// Health calls apis.v1.HealthService.Health.
func (c *healthServiceClient) Health(ctx context.Context, req *connect.Request[v1.HealthRequest]) (*connect.Response[v1.HealthResponse], error) {
	return c.health.CallUnary(ctx, req)
}

// HealthServiceHandler is an implementation of the apis.v1.HealthService service.
type HealthServiceHandler interface {
	Health(context.Context, *connect.Request[v1.HealthRequest]) (*connect.Response[v1.HealthResponse], error)
}

// NewHealthServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthServiceHandler(svc HealthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	healthServiceHealthHandler := connect.NewUnaryHandler(
		HealthServiceHealthProcedure,
		svc.Health,
		opts...,
	)
	return "/apis.v1.HealthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthServiceHealthProcedure:
			healthServiceHealthHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHealthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthServiceHandler struct{}

func (UnimplementedHealthServiceHandler) Health(context.Context, *connect.Request[v1.HealthRequest]) (*connect.Response[v1.HealthResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("apis.v1.HealthService.Health is not implemented"))
}
