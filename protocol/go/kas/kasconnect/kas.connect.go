// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: kas/kas.proto

package kasconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	kas "github.com/opentdf/platform/protocol/go/kas"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AccessServiceName is the fully-qualified name of the AccessService service.
	AccessServiceName = "kas.AccessService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccessServicePublicKeyProcedure is the fully-qualified name of the AccessService's PublicKey RPC.
	AccessServicePublicKeyProcedure = "/kas.AccessService/PublicKey"
	// AccessServiceLegacyPublicKeyProcedure is the fully-qualified name of the AccessService's
	// LegacyPublicKey RPC.
	AccessServiceLegacyPublicKeyProcedure = "/kas.AccessService/LegacyPublicKey"
	// AccessServiceRewrapProcedure is the fully-qualified name of the AccessService's Rewrap RPC.
	AccessServiceRewrapProcedure = "/kas.AccessService/Rewrap"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accessServiceServiceDescriptor               = kas.File_kas_kas_proto.Services().ByName("AccessService")
	accessServicePublicKeyMethodDescriptor       = accessServiceServiceDescriptor.Methods().ByName("PublicKey")
	accessServiceLegacyPublicKeyMethodDescriptor = accessServiceServiceDescriptor.Methods().ByName("LegacyPublicKey")
	accessServiceRewrapMethodDescriptor          = accessServiceServiceDescriptor.Methods().ByName("Rewrap")
)

// AccessServiceClient is a client for the kas.AccessService service.
type AccessServiceClient interface {
	PublicKey(context.Context, *connect.Request[kas.PublicKeyRequest]) (*connect.Response[kas.PublicKeyResponse], error)
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	LegacyPublicKey(context.Context, *connect.Request[kas.LegacyPublicKeyRequest]) (*connect.Response[wrapperspb.StringValue], error)
	Rewrap(context.Context, *connect.Request[kas.RewrapRequest]) (*connect.Response[kas.RewrapResponse], error)
}

// NewAccessServiceClient constructs a client for the kas.AccessService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccessServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccessServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accessServiceClient{
		publicKey: connect.NewClient[kas.PublicKeyRequest, kas.PublicKeyResponse](
			httpClient,
			baseURL+AccessServicePublicKeyProcedure,
			connect.WithSchema(accessServicePublicKeyMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		legacyPublicKey: connect.NewClient[kas.LegacyPublicKeyRequest, wrapperspb.StringValue](
			httpClient,
			baseURL+AccessServiceLegacyPublicKeyProcedure,
			connect.WithSchema(accessServiceLegacyPublicKeyMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		rewrap: connect.NewClient[kas.RewrapRequest, kas.RewrapResponse](
			httpClient,
			baseURL+AccessServiceRewrapProcedure,
			connect.WithSchema(accessServiceRewrapMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accessServiceClient implements AccessServiceClient.
type accessServiceClient struct {
	publicKey       *connect.Client[kas.PublicKeyRequest, kas.PublicKeyResponse]
	legacyPublicKey *connect.Client[kas.LegacyPublicKeyRequest, wrapperspb.StringValue]
	rewrap          *connect.Client[kas.RewrapRequest, kas.RewrapResponse]
}

// PublicKey calls kas.AccessService.PublicKey.
func (c *accessServiceClient) PublicKey(ctx context.Context, req *connect.Request[kas.PublicKeyRequest]) (*connect.Response[kas.PublicKeyResponse], error) {
	return c.publicKey.CallUnary(ctx, req)
}

// LegacyPublicKey calls kas.AccessService.LegacyPublicKey.
func (c *accessServiceClient) LegacyPublicKey(ctx context.Context, req *connect.Request[kas.LegacyPublicKeyRequest]) (*connect.Response[wrapperspb.StringValue], error) {
	return c.legacyPublicKey.CallUnary(ctx, req)
}

// Rewrap calls kas.AccessService.Rewrap.
func (c *accessServiceClient) Rewrap(ctx context.Context, req *connect.Request[kas.RewrapRequest]) (*connect.Response[kas.RewrapResponse], error) {
	return c.rewrap.CallUnary(ctx, req)
}

// AccessServiceHandler is an implementation of the kas.AccessService service.
type AccessServiceHandler interface {
	PublicKey(context.Context, *connect.Request[kas.PublicKeyRequest]) (*connect.Response[kas.PublicKeyResponse], error)
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	LegacyPublicKey(context.Context, *connect.Request[kas.LegacyPublicKeyRequest]) (*connect.Response[wrapperspb.StringValue], error)
	Rewrap(context.Context, *connect.Request[kas.RewrapRequest]) (*connect.Response[kas.RewrapResponse], error)
}

// NewAccessServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccessServiceHandler(svc AccessServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accessServicePublicKeyHandler := connect.NewUnaryHandler(
		AccessServicePublicKeyProcedure,
		svc.PublicKey,
		connect.WithSchema(accessServicePublicKeyMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	accessServiceLegacyPublicKeyHandler := connect.NewUnaryHandler(
		AccessServiceLegacyPublicKeyProcedure,
		svc.LegacyPublicKey,
		connect.WithSchema(accessServiceLegacyPublicKeyMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	accessServiceRewrapHandler := connect.NewUnaryHandler(
		AccessServiceRewrapProcedure,
		svc.Rewrap,
		connect.WithSchema(accessServiceRewrapMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/kas.AccessService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccessServicePublicKeyProcedure:
			accessServicePublicKeyHandler.ServeHTTP(w, r)
		case AccessServiceLegacyPublicKeyProcedure:
			accessServiceLegacyPublicKeyHandler.ServeHTTP(w, r)
		case AccessServiceRewrapProcedure:
			accessServiceRewrapHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccessServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccessServiceHandler struct{}

func (UnimplementedAccessServiceHandler) PublicKey(context.Context, *connect.Request[kas.PublicKeyRequest]) (*connect.Response[kas.PublicKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kas.AccessService.PublicKey is not implemented"))
}

func (UnimplementedAccessServiceHandler) LegacyPublicKey(context.Context, *connect.Request[kas.LegacyPublicKeyRequest]) (*connect.Response[wrapperspb.StringValue], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kas.AccessService.LegacyPublicKey is not implemented"))
}

func (UnimplementedAccessServiceHandler) Rewrap(context.Context, *connect.Request[kas.RewrapRequest]) (*connect.Response[kas.RewrapResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("kas.AccessService.Rewrap is not implemented"))
}
