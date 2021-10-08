// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: global.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/global/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// GlobalService global.proto
	GlobalService() pb.GlobalServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		globalService: pb.NewGlobalServiceClient(cc),
	}
}

type serviceClients struct {
	globalService pb.GlobalServiceClient
}

func (c *serviceClients) GlobalService() pb.GlobalServiceClient {
	return c.globalService
}

type globalServiceWrapper struct {
	client pb.GlobalServiceClient
	opts   []grpc1.CallOption
}

func (s *globalServiceWrapper) GetHealth(ctx context.Context, req *pb.GetHealthRequest) (*pb.GetHealthResponse, error) {
	return s.client.GetHealth(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *globalServiceWrapper) GetTenantGroup(ctx context.Context, req *pb.GetTenantGroupRequest) (*pb.GetTenantGroupResponse, error) {
	return s.client.GetTenantGroup(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *globalServiceWrapper) CreateTenant(ctx context.Context, req *pb.CreateTenantRequest) (*pb.CreateTenantResponse, error) {
	return s.client.CreateTenant(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *globalServiceWrapper) GetFeatures(ctx context.Context, req *pb.GetFeaturesRequest) (*pb.GetFeaturesResponse, error) {
	return s.client.GetFeatures(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}