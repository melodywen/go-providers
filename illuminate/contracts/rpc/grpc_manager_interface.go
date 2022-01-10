package rpc

import (
	"context"
	runtime2 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type RegisterHandlerFromEndpointFun func(ctx context.Context, mux *runtime2.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type GrpcManagerInterface interface {
	NewServer(opt ...grpc.ServerOption)
	NewServeMux(opts ...runtime2.ServeMuxOption)
	SetDialOption([]grpc.DialOption)
	SetEndpoint(entrypoint string)
	Init()
	SetRegisterHandlerFromEndpointCallbacks(callbacks []RegisterHandlerFromEndpointFun)
	RegisterHandlerFromEndpoint(err error)
	Run()
}
