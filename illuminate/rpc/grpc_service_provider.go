package rpc

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/support"
	"github.com/melodywen/go-providers/illuminate/contracts/rpc"
)

// GRpcServiceProvider grpc struct
type GRpcServiceProvider struct {
	support.ServiceProvider
}

// NewGRpcServiceProvider new grpc service
func NewGRpcServiceProvider() *GRpcServiceProvider {
	return &GRpcServiceProvider{}
}

// Boot Bootstrap any application services.
func (provider *GRpcServiceProvider) Boot() {
	fmt.Println("grpc boot")
}

// Register any application services.
func (provider *GRpcServiceProvider) Register() {
	var grpc rpc.GrpcManagerInterface
	provider.App.Alias("grpc-manager", &grpc)
	provider.App.Instance("grpc-manager", NewGrpcManager(provider.App))
	fmt.Println("grpc register")
}
