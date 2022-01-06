package rpc

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type GrpcManager struct {
	serverNux *runtime.ServeMux
	opt       []grpc.DialOption
	port      string
}

func NewGrpcManager() *GrpcManager {
	manager := &GrpcManager{
		port: ":8080",
	}

	//err := gen.RegisterBlogHandlerFromEndpoint(context.Background(), gwmux, ":9005", opt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//rpcServer := grpc.NewServer(
	//	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//		UnaryGSError500(""),
	//	)),
	//	)
	//gen.RegisterBlogServer(rpcServer, new(servers.BlogImplementServer))
	//
	//http.ListenAndServe(
	//	":9005",
	//	grpcHandlerFunc(rpcServer, gwmux),
	//)
	return manager
}

// NewServeMux 创建grpc-gateway服务，转发到grpc的9005端口
func (manager *GrpcManager) NewServeMux(opts ...runtime.ServeMuxOption) {
	manager.serverNux = runtime.NewServeMux(opts...)
}

// SetDialOption  DialOption configures how we set up the connection.
func (manager *GrpcManager) SetDialOption(opt []grpc.DialOption) {
	manager.opt = opt
}

// Run running server
func (manager *GrpcManager) Run() {
	if manager.serverNux == nil {
		manager.NewServeMux()
	}

	if len(manager.opt) == 0 {
		manager.SetDialOption([]grpc.DialOption{grpc.WithInsecure()})
	}

	fmt.Println(manager.opt)

	fmt.Println(13131)
}
