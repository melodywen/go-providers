package rpc

import (
	runtime2 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type GrpcManagerInterface interface {
	NewServer(opt ...grpc.ServerOption)
	NewServeMux(opts ...runtime2.ServeMuxOption)
	SetDialOption([]grpc.DialOption)
	SetEndpoint(entrypoint string)
	Init()
	RegisterHandlerFromEndpoint(err error)
	Run()
}
