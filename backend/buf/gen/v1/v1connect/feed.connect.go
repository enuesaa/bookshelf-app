// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/feed.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// FeedName is the fully-qualified name of the Feed service.
	FeedName = "v1.Feed"
)

// FeedClient is a client for the v1.Feed service.
type FeedClient interface {
	AddFeed(context.Context, *connect_go.Request[v1.AddFeedRequest]) (*connect_go.Response[v1.AddFeedResponse], error)
	GetFeed(context.Context, *connect_go.Request[v1.GetFeedRequest]) (*connect_go.Response[v1.GetFeedResponse], error)
	ListItems(context.Context, *connect_go.Request[v1.ListItemsRequest]) (*connect_go.Response[v1.ListItemsResponse], error)
	GetAppearance(context.Context, *connect_go.Request[v1.GetAppearanceRequest]) (*connect_go.Response[v1.GetAppearanceResponse], error)
	UpdateAppearance(context.Context, *connect_go.Request[v1.UpdateAppearanceRequest]) (*connect_go.Response[v1.UpdateAppearanceResponse], error)
	Fetch(context.Context, *connect_go.Request[v1.FetchRequest]) (*connect_go.Response[v1.FetchResponse], error)
	DeleteFeed(context.Context, *connect_go.Request[v1.DeleteFeedRequest]) (*connect_go.Response[v1.DeleteFeedResponse], error)
}

// NewFeedClient constructs a client for the v1.Feed service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFeedClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) FeedClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &feedClient{
		addFeed: connect_go.NewClient[v1.AddFeedRequest, v1.AddFeedResponse](
			httpClient,
			baseURL+"/v1.Feed/AddFeed",
			opts...,
		),
		getFeed: connect_go.NewClient[v1.GetFeedRequest, v1.GetFeedResponse](
			httpClient,
			baseURL+"/v1.Feed/GetFeed",
			opts...,
		),
		listItems: connect_go.NewClient[v1.ListItemsRequest, v1.ListItemsResponse](
			httpClient,
			baseURL+"/v1.Feed/ListItems",
			opts...,
		),
		getAppearance: connect_go.NewClient[v1.GetAppearanceRequest, v1.GetAppearanceResponse](
			httpClient,
			baseURL+"/v1.Feed/GetAppearance",
			opts...,
		),
		updateAppearance: connect_go.NewClient[v1.UpdateAppearanceRequest, v1.UpdateAppearanceResponse](
			httpClient,
			baseURL+"/v1.Feed/UpdateAppearance",
			opts...,
		),
		fetch: connect_go.NewClient[v1.FetchRequest, v1.FetchResponse](
			httpClient,
			baseURL+"/v1.Feed/Fetch",
			opts...,
		),
		deleteFeed: connect_go.NewClient[v1.DeleteFeedRequest, v1.DeleteFeedResponse](
			httpClient,
			baseURL+"/v1.Feed/DeleteFeed",
			opts...,
		),
	}
}

// feedClient implements FeedClient.
type feedClient struct {
	addFeed          *connect_go.Client[v1.AddFeedRequest, v1.AddFeedResponse]
	getFeed          *connect_go.Client[v1.GetFeedRequest, v1.GetFeedResponse]
	listItems        *connect_go.Client[v1.ListItemsRequest, v1.ListItemsResponse]
	getAppearance    *connect_go.Client[v1.GetAppearanceRequest, v1.GetAppearanceResponse]
	updateAppearance *connect_go.Client[v1.UpdateAppearanceRequest, v1.UpdateAppearanceResponse]
	fetch            *connect_go.Client[v1.FetchRequest, v1.FetchResponse]
	deleteFeed       *connect_go.Client[v1.DeleteFeedRequest, v1.DeleteFeedResponse]
}

// AddFeed calls v1.Feed.AddFeed.
func (c *feedClient) AddFeed(ctx context.Context, req *connect_go.Request[v1.AddFeedRequest]) (*connect_go.Response[v1.AddFeedResponse], error) {
	return c.addFeed.CallUnary(ctx, req)
}

// GetFeed calls v1.Feed.GetFeed.
func (c *feedClient) GetFeed(ctx context.Context, req *connect_go.Request[v1.GetFeedRequest]) (*connect_go.Response[v1.GetFeedResponse], error) {
	return c.getFeed.CallUnary(ctx, req)
}

// ListItems calls v1.Feed.ListItems.
func (c *feedClient) ListItems(ctx context.Context, req *connect_go.Request[v1.ListItemsRequest]) (*connect_go.Response[v1.ListItemsResponse], error) {
	return c.listItems.CallUnary(ctx, req)
}

// GetAppearance calls v1.Feed.GetAppearance.
func (c *feedClient) GetAppearance(ctx context.Context, req *connect_go.Request[v1.GetAppearanceRequest]) (*connect_go.Response[v1.GetAppearanceResponse], error) {
	return c.getAppearance.CallUnary(ctx, req)
}

// UpdateAppearance calls v1.Feed.UpdateAppearance.
func (c *feedClient) UpdateAppearance(ctx context.Context, req *connect_go.Request[v1.UpdateAppearanceRequest]) (*connect_go.Response[v1.UpdateAppearanceResponse], error) {
	return c.updateAppearance.CallUnary(ctx, req)
}

// Fetch calls v1.Feed.Fetch.
func (c *feedClient) Fetch(ctx context.Context, req *connect_go.Request[v1.FetchRequest]) (*connect_go.Response[v1.FetchResponse], error) {
	return c.fetch.CallUnary(ctx, req)
}

// DeleteFeed calls v1.Feed.DeleteFeed.
func (c *feedClient) DeleteFeed(ctx context.Context, req *connect_go.Request[v1.DeleteFeedRequest]) (*connect_go.Response[v1.DeleteFeedResponse], error) {
	return c.deleteFeed.CallUnary(ctx, req)
}

// FeedHandler is an implementation of the v1.Feed service.
type FeedHandler interface {
	AddFeed(context.Context, *connect_go.Request[v1.AddFeedRequest]) (*connect_go.Response[v1.AddFeedResponse], error)
	GetFeed(context.Context, *connect_go.Request[v1.GetFeedRequest]) (*connect_go.Response[v1.GetFeedResponse], error)
	ListItems(context.Context, *connect_go.Request[v1.ListItemsRequest]) (*connect_go.Response[v1.ListItemsResponse], error)
	GetAppearance(context.Context, *connect_go.Request[v1.GetAppearanceRequest]) (*connect_go.Response[v1.GetAppearanceResponse], error)
	UpdateAppearance(context.Context, *connect_go.Request[v1.UpdateAppearanceRequest]) (*connect_go.Response[v1.UpdateAppearanceResponse], error)
	Fetch(context.Context, *connect_go.Request[v1.FetchRequest]) (*connect_go.Response[v1.FetchResponse], error)
	DeleteFeed(context.Context, *connect_go.Request[v1.DeleteFeedRequest]) (*connect_go.Response[v1.DeleteFeedResponse], error)
}

// NewFeedHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFeedHandler(svc FeedHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/v1.Feed/AddFeed", connect_go.NewUnaryHandler(
		"/v1.Feed/AddFeed",
		svc.AddFeed,
		opts...,
	))
	mux.Handle("/v1.Feed/GetFeed", connect_go.NewUnaryHandler(
		"/v1.Feed/GetFeed",
		svc.GetFeed,
		opts...,
	))
	mux.Handle("/v1.Feed/ListItems", connect_go.NewUnaryHandler(
		"/v1.Feed/ListItems",
		svc.ListItems,
		opts...,
	))
	mux.Handle("/v1.Feed/GetAppearance", connect_go.NewUnaryHandler(
		"/v1.Feed/GetAppearance",
		svc.GetAppearance,
		opts...,
	))
	mux.Handle("/v1.Feed/UpdateAppearance", connect_go.NewUnaryHandler(
		"/v1.Feed/UpdateAppearance",
		svc.UpdateAppearance,
		opts...,
	))
	mux.Handle("/v1.Feed/Fetch", connect_go.NewUnaryHandler(
		"/v1.Feed/Fetch",
		svc.Fetch,
		opts...,
	))
	mux.Handle("/v1.Feed/DeleteFeed", connect_go.NewUnaryHandler(
		"/v1.Feed/DeleteFeed",
		svc.DeleteFeed,
		opts...,
	))
	return "/v1.Feed/", mux
}

// UnimplementedFeedHandler returns CodeUnimplemented from all methods.
type UnimplementedFeedHandler struct{}

func (UnimplementedFeedHandler) AddFeed(context.Context, *connect_go.Request[v1.AddFeedRequest]) (*connect_go.Response[v1.AddFeedResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("v1.Feed.AddFeed is not implemented"))
}

func (UnimplementedFeedHandler) GetFeed(context.Context, *connect_go.Request[v1.GetFeedRequest]) (*connect_go.Response[v1.GetFeedResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("v1.Feed.GetFeed is not implemented"))
}

func (UnimplementedFeedHandler) ListItems(context.Context, *connect_go.Request[v1.ListItemsRequest]) (*connect_go.Response[v1.ListItemsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("v1.Feed.ListItems is not implemented"))
}

func (UnimplementedFeedHandler) GetAppearance(context.Context, *connect_go.Request[v1.GetAppearanceRequest]) (*connect_go.Response[v1.GetAppearanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("v1.Feed.GetAppearance is not implemented"))
}

func (UnimplementedFeedHandler) UpdateAppearance(context.Context, *connect_go.Request[v1.UpdateAppearanceRequest]) (*connect_go.Response[v1.UpdateAppearanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("v1.Feed.UpdateAppearance is not implemented"))
}

func (UnimplementedFeedHandler) Fetch(context.Context, *connect_go.Request[v1.FetchRequest]) (*connect_go.Response[v1.FetchResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("v1.Feed.Fetch is not implemented"))
}

func (UnimplementedFeedHandler) DeleteFeed(context.Context, *connect_go.Request[v1.DeleteFeedRequest]) (*connect_go.Response[v1.DeleteFeedResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("v1.Feed.DeleteFeed is not implemented"))
}
