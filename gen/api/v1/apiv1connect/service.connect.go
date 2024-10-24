// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1/service.proto

package apiv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/bit-bom/minefield/gen/api/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// QueryServiceName is the fully-qualified name of the QueryService service.
	QueryServiceName = "api.v1.QueryService"
	// CacheServiceName is the fully-qualified name of the CacheService service.
	CacheServiceName = "api.v1.CacheService"
	// LeaderboardServiceName is the fully-qualified name of the LeaderboardService service.
	LeaderboardServiceName = "api.v1.LeaderboardService"
	// GraphServiceName is the fully-qualified name of the GraphService service.
	GraphServiceName = "api.v1.GraphService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QueryServiceQueryProcedure is the fully-qualified name of the QueryService's Query RPC.
	QueryServiceQueryProcedure = "/api.v1.QueryService/Query"
	// CacheServiceCacheProcedure is the fully-qualified name of the CacheService's Cache RPC.
	CacheServiceCacheProcedure = "/api.v1.CacheService/Cache"
	// CacheServiceClearProcedure is the fully-qualified name of the CacheService's Clear RPC.
	CacheServiceClearProcedure = "/api.v1.CacheService/Clear"
	// LeaderboardServiceCustomLeaderboardProcedure is the fully-qualified name of the
	// LeaderboardService's CustomLeaderboard RPC.
	LeaderboardServiceCustomLeaderboardProcedure = "/api.v1.LeaderboardService/CustomLeaderboard"
	// LeaderboardServiceAllKeysProcedure is the fully-qualified name of the LeaderboardService's
	// AllKeys RPC.
	LeaderboardServiceAllKeysProcedure = "/api.v1.LeaderboardService/AllKeys"
	// GraphServiceGetNodeProcedure is the fully-qualified name of the GraphService's GetNode RPC.
	GraphServiceGetNodeProcedure = "/api.v1.GraphService/GetNode"
	// GraphServiceGetNodesByGlobProcedure is the fully-qualified name of the GraphService's
	// GetNodesByGlob RPC.
	GraphServiceGetNodesByGlobProcedure = "/api.v1.GraphService/GetNodesByGlob"
	// GraphServiceGetNodeByNameProcedure is the fully-qualified name of the GraphService's
	// GetNodeByName RPC.
	GraphServiceGetNodeByNameProcedure = "/api.v1.GraphService/GetNodeByName"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	queryServiceServiceDescriptor                       = v1.File_api_v1_service_proto.Services().ByName("QueryService")
	queryServiceQueryMethodDescriptor                   = queryServiceServiceDescriptor.Methods().ByName("Query")
	cacheServiceServiceDescriptor                       = v1.File_api_v1_service_proto.Services().ByName("CacheService")
	cacheServiceCacheMethodDescriptor                   = cacheServiceServiceDescriptor.Methods().ByName("Cache")
	cacheServiceClearMethodDescriptor                   = cacheServiceServiceDescriptor.Methods().ByName("Clear")
	leaderboardServiceServiceDescriptor                 = v1.File_api_v1_service_proto.Services().ByName("LeaderboardService")
	leaderboardServiceCustomLeaderboardMethodDescriptor = leaderboardServiceServiceDescriptor.Methods().ByName("CustomLeaderboard")
	leaderboardServiceAllKeysMethodDescriptor           = leaderboardServiceServiceDescriptor.Methods().ByName("AllKeys")
	graphServiceServiceDescriptor                       = v1.File_api_v1_service_proto.Services().ByName("GraphService")
	graphServiceGetNodeMethodDescriptor                 = graphServiceServiceDescriptor.Methods().ByName("GetNode")
	graphServiceGetNodesByGlobMethodDescriptor          = graphServiceServiceDescriptor.Methods().ByName("GetNodesByGlob")
	graphServiceGetNodeByNameMethodDescriptor           = graphServiceServiceDescriptor.Methods().ByName("GetNodeByName")
)

// QueryServiceClient is a client for the api.v1.QueryService service.
type QueryServiceClient interface {
	Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error)
}

// NewQueryServiceClient constructs a client for the api.v1.QueryService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQueryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) QueryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &queryServiceClient{
		query: connect.NewClient[v1.QueryRequest, v1.QueryResponse](
			httpClient,
			baseURL+QueryServiceQueryProcedure,
			connect.WithSchema(queryServiceQueryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// queryServiceClient implements QueryServiceClient.
type queryServiceClient struct {
	query *connect.Client[v1.QueryRequest, v1.QueryResponse]
}

// Query calls api.v1.QueryService.Query.
func (c *queryServiceClient) Query(ctx context.Context, req *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	return c.query.CallUnary(ctx, req)
}

// QueryServiceHandler is an implementation of the api.v1.QueryService service.
type QueryServiceHandler interface {
	Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error)
}

// NewQueryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueryServiceHandler(svc QueryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queryServiceQueryHandler := connect.NewUnaryHandler(
		QueryServiceQueryProcedure,
		svc.Query,
		connect.WithSchema(queryServiceQueryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.QueryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueryServiceQueryProcedure:
			queryServiceQueryHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQueryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQueryServiceHandler struct{}

func (UnimplementedQueryServiceHandler) Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.QueryService.Query is not implemented"))
}

// CacheServiceClient is a client for the api.v1.CacheService service.
type CacheServiceClient interface {
	Cache(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error)
	Clear(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error)
}

// NewCacheServiceClient constructs a client for the api.v1.CacheService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCacheServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CacheServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &cacheServiceClient{
		cache: connect.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+CacheServiceCacheProcedure,
			connect.WithSchema(cacheServiceCacheMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		clear: connect.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+CacheServiceClearProcedure,
			connect.WithSchema(cacheServiceClearMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// cacheServiceClient implements CacheServiceClient.
type cacheServiceClient struct {
	cache *connect.Client[emptypb.Empty, emptypb.Empty]
	clear *connect.Client[emptypb.Empty, emptypb.Empty]
}

// Cache calls api.v1.CacheService.Cache.
func (c *cacheServiceClient) Cache(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	return c.cache.CallUnary(ctx, req)
}

// Clear calls api.v1.CacheService.Clear.
func (c *cacheServiceClient) Clear(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	return c.clear.CallUnary(ctx, req)
}

// CacheServiceHandler is an implementation of the api.v1.CacheService service.
type CacheServiceHandler interface {
	Cache(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error)
	Clear(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error)
}

// NewCacheServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCacheServiceHandler(svc CacheServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	cacheServiceCacheHandler := connect.NewUnaryHandler(
		CacheServiceCacheProcedure,
		svc.Cache,
		connect.WithSchema(cacheServiceCacheMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	cacheServiceClearHandler := connect.NewUnaryHandler(
		CacheServiceClearProcedure,
		svc.Clear,
		connect.WithSchema(cacheServiceClearMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.CacheService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CacheServiceCacheProcedure:
			cacheServiceCacheHandler.ServeHTTP(w, r)
		case CacheServiceClearProcedure:
			cacheServiceClearHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCacheServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCacheServiceHandler struct{}

func (UnimplementedCacheServiceHandler) Cache(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.CacheService.Cache is not implemented"))
}

func (UnimplementedCacheServiceHandler) Clear(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.CacheService.Clear is not implemented"))
}

// LeaderboardServiceClient is a client for the api.v1.LeaderboardService service.
type LeaderboardServiceClient interface {
	CustomLeaderboard(context.Context, *connect.Request[v1.CustomLeaderboardRequest]) (*connect.Response[v1.CustomLeaderboardResponse], error)
	AllKeys(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.AllKeysResponse], error)
}

// NewLeaderboardServiceClient constructs a client for the api.v1.LeaderboardService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLeaderboardServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LeaderboardServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &leaderboardServiceClient{
		customLeaderboard: connect.NewClient[v1.CustomLeaderboardRequest, v1.CustomLeaderboardResponse](
			httpClient,
			baseURL+LeaderboardServiceCustomLeaderboardProcedure,
			connect.WithSchema(leaderboardServiceCustomLeaderboardMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		allKeys: connect.NewClient[emptypb.Empty, v1.AllKeysResponse](
			httpClient,
			baseURL+LeaderboardServiceAllKeysProcedure,
			connect.WithSchema(leaderboardServiceAllKeysMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// leaderboardServiceClient implements LeaderboardServiceClient.
type leaderboardServiceClient struct {
	customLeaderboard *connect.Client[v1.CustomLeaderboardRequest, v1.CustomLeaderboardResponse]
	allKeys           *connect.Client[emptypb.Empty, v1.AllKeysResponse]
}

// CustomLeaderboard calls api.v1.LeaderboardService.CustomLeaderboard.
func (c *leaderboardServiceClient) CustomLeaderboard(ctx context.Context, req *connect.Request[v1.CustomLeaderboardRequest]) (*connect.Response[v1.CustomLeaderboardResponse], error) {
	return c.customLeaderboard.CallUnary(ctx, req)
}

// AllKeys calls api.v1.LeaderboardService.AllKeys.
func (c *leaderboardServiceClient) AllKeys(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.AllKeysResponse], error) {
	return c.allKeys.CallUnary(ctx, req)
}

// LeaderboardServiceHandler is an implementation of the api.v1.LeaderboardService service.
type LeaderboardServiceHandler interface {
	CustomLeaderboard(context.Context, *connect.Request[v1.CustomLeaderboardRequest]) (*connect.Response[v1.CustomLeaderboardResponse], error)
	AllKeys(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.AllKeysResponse], error)
}

// NewLeaderboardServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLeaderboardServiceHandler(svc LeaderboardServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	leaderboardServiceCustomLeaderboardHandler := connect.NewUnaryHandler(
		LeaderboardServiceCustomLeaderboardProcedure,
		svc.CustomLeaderboard,
		connect.WithSchema(leaderboardServiceCustomLeaderboardMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	leaderboardServiceAllKeysHandler := connect.NewUnaryHandler(
		LeaderboardServiceAllKeysProcedure,
		svc.AllKeys,
		connect.WithSchema(leaderboardServiceAllKeysMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.LeaderboardService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LeaderboardServiceCustomLeaderboardProcedure:
			leaderboardServiceCustomLeaderboardHandler.ServeHTTP(w, r)
		case LeaderboardServiceAllKeysProcedure:
			leaderboardServiceAllKeysHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLeaderboardServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLeaderboardServiceHandler struct{}

func (UnimplementedLeaderboardServiceHandler) CustomLeaderboard(context.Context, *connect.Request[v1.CustomLeaderboardRequest]) (*connect.Response[v1.CustomLeaderboardResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.LeaderboardService.CustomLeaderboard is not implemented"))
}

func (UnimplementedLeaderboardServiceHandler) AllKeys(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.AllKeysResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.LeaderboardService.AllKeys is not implemented"))
}

// GraphServiceClient is a client for the api.v1.GraphService service.
type GraphServiceClient interface {
	GetNode(context.Context, *connect.Request[v1.GetNodeRequest]) (*connect.Response[v1.GetNodeResponse], error)
	GetNodesByGlob(context.Context, *connect.Request[v1.GetNodesByGlobRequest]) (*connect.Response[v1.GetNodesByGlobResponse], error)
	GetNodeByName(context.Context, *connect.Request[v1.GetNodeByNameRequest]) (*connect.Response[v1.GetNodeByNameResponse], error)
}

// NewGraphServiceClient constructs a client for the api.v1.GraphService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGraphServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GraphServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &graphServiceClient{
		getNode: connect.NewClient[v1.GetNodeRequest, v1.GetNodeResponse](
			httpClient,
			baseURL+GraphServiceGetNodeProcedure,
			connect.WithSchema(graphServiceGetNodeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getNodesByGlob: connect.NewClient[v1.GetNodesByGlobRequest, v1.GetNodesByGlobResponse](
			httpClient,
			baseURL+GraphServiceGetNodesByGlobProcedure,
			connect.WithSchema(graphServiceGetNodesByGlobMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getNodeByName: connect.NewClient[v1.GetNodeByNameRequest, v1.GetNodeByNameResponse](
			httpClient,
			baseURL+GraphServiceGetNodeByNameProcedure,
			connect.WithSchema(graphServiceGetNodeByNameMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// graphServiceClient implements GraphServiceClient.
type graphServiceClient struct {
	getNode        *connect.Client[v1.GetNodeRequest, v1.GetNodeResponse]
	getNodesByGlob *connect.Client[v1.GetNodesByGlobRequest, v1.GetNodesByGlobResponse]
	getNodeByName  *connect.Client[v1.GetNodeByNameRequest, v1.GetNodeByNameResponse]
}

// GetNode calls api.v1.GraphService.GetNode.
func (c *graphServiceClient) GetNode(ctx context.Context, req *connect.Request[v1.GetNodeRequest]) (*connect.Response[v1.GetNodeResponse], error) {
	return c.getNode.CallUnary(ctx, req)
}

// GetNodesByGlob calls api.v1.GraphService.GetNodesByGlob.
func (c *graphServiceClient) GetNodesByGlob(ctx context.Context, req *connect.Request[v1.GetNodesByGlobRequest]) (*connect.Response[v1.GetNodesByGlobResponse], error) {
	return c.getNodesByGlob.CallUnary(ctx, req)
}

// GetNodeByName calls api.v1.GraphService.GetNodeByName.
func (c *graphServiceClient) GetNodeByName(ctx context.Context, req *connect.Request[v1.GetNodeByNameRequest]) (*connect.Response[v1.GetNodeByNameResponse], error) {
	return c.getNodeByName.CallUnary(ctx, req)
}

// GraphServiceHandler is an implementation of the api.v1.GraphService service.
type GraphServiceHandler interface {
	GetNode(context.Context, *connect.Request[v1.GetNodeRequest]) (*connect.Response[v1.GetNodeResponse], error)
	GetNodesByGlob(context.Context, *connect.Request[v1.GetNodesByGlobRequest]) (*connect.Response[v1.GetNodesByGlobResponse], error)
	GetNodeByName(context.Context, *connect.Request[v1.GetNodeByNameRequest]) (*connect.Response[v1.GetNodeByNameResponse], error)
}

// NewGraphServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGraphServiceHandler(svc GraphServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	graphServiceGetNodeHandler := connect.NewUnaryHandler(
		GraphServiceGetNodeProcedure,
		svc.GetNode,
		connect.WithSchema(graphServiceGetNodeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	graphServiceGetNodesByGlobHandler := connect.NewUnaryHandler(
		GraphServiceGetNodesByGlobProcedure,
		svc.GetNodesByGlob,
		connect.WithSchema(graphServiceGetNodesByGlobMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	graphServiceGetNodeByNameHandler := connect.NewUnaryHandler(
		GraphServiceGetNodeByNameProcedure,
		svc.GetNodeByName,
		connect.WithSchema(graphServiceGetNodeByNameMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.v1.GraphService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GraphServiceGetNodeProcedure:
			graphServiceGetNodeHandler.ServeHTTP(w, r)
		case GraphServiceGetNodesByGlobProcedure:
			graphServiceGetNodesByGlobHandler.ServeHTTP(w, r)
		case GraphServiceGetNodeByNameProcedure:
			graphServiceGetNodeByNameHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGraphServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGraphServiceHandler struct{}

func (UnimplementedGraphServiceHandler) GetNode(context.Context, *connect.Request[v1.GetNodeRequest]) (*connect.Response[v1.GetNodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.GraphService.GetNode is not implemented"))
}

func (UnimplementedGraphServiceHandler) GetNodesByGlob(context.Context, *connect.Request[v1.GetNodesByGlobRequest]) (*connect.Response[v1.GetNodesByGlobResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.GraphService.GetNodesByGlob is not implemented"))
}

func (UnimplementedGraphServiceHandler) GetNodeByName(context.Context, *connect.Request[v1.GetNodeByNameRequest]) (*connect.Response[v1.GetNodeByNameResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.v1.GraphService.GetNodeByName is not implemented"))
}
