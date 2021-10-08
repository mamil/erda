// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: legacy_upstream_lb.proto

package pb

import (
	transport "github.com/erda-project/erda-infra/pkg/transport"
	reflect "reflect"
)

// RegisterUpstreamLbServiceImp legacy_upstream_lb.proto
func RegisterUpstreamLbServiceImp(regester transport.Register, srv UpstreamLbServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterUpstreamLbServiceHandler(regester, UpstreamLbServiceHandler(srv), _ops.HTTP...)
	RegisterUpstreamLbServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.hepa.legacy_upstream_lb.UpstreamLbService",
	)
}

var (
	upstreamLbServiceClientType  = reflect.TypeOf((*UpstreamLbServiceClient)(nil)).Elem()
	upstreamLbServiceServerType  = reflect.TypeOf((*UpstreamLbServiceServer)(nil)).Elem()
	upstreamLbServiceHandlerType = reflect.TypeOf((*UpstreamLbServiceHandler)(nil)).Elem()
)

// UpstreamLbServiceClientType .
func UpstreamLbServiceClientType() reflect.Type { return upstreamLbServiceClientType }

// UpstreamLbServiceServerType .
func UpstreamLbServiceServerType() reflect.Type { return upstreamLbServiceServerType }

// UpstreamLbServiceHandlerType .
func UpstreamLbServiceHandlerType() reflect.Type { return upstreamLbServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		upstreamLbServiceClientType,
		// server types
		upstreamLbServiceServerType,
		// handler types
		upstreamLbServiceHandlerType,
	}
}