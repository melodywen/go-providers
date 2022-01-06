package rpc

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type GrpcManagerInterface interface {
	NewServeMux(opts ...runtime.ServeMuxOption)
	SetDialOption([]grpc.DialOption)
	Run()
}
