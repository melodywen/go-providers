package rpc

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/support"
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
	provider.App.Instance("grpc-manager", NewGrpcManager())
	fmt.Println("grpc register")
}
