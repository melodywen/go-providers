package rpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	runtime2 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-providers/servers/gen/golang"
	"github.com/melodywen/go-providers/servers/implement_servers"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

type GrpcManager struct {
	serverNux *runtime.ServeMux
	opt       []grpc.DialOption
	port      string
	ctx       context.Context
}

func NewGrpcManager(app foundation.ApplicationInterface) *GrpcManager {
	manager := &GrpcManager{
		port: ":8080",
		ctx:  context.Background(),
	}

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
	gwmux := runtime2.NewServeMux()
	opt := []grpc.DialOption{grpc.WithInsecure()}

	fmt.Println(13123)
	err := golang.RegisterBlogHandlerFromEndpoint(context.Background(), gwmux, ":9005", opt)
	if err != nil {
		fmt.Println(err)
	}
	rpcServer := grpc.NewServer()
	golang.RegisterBlogServer(rpcServer, implement_servers.NewBlogImplementServer())

	http.ListenAndServe(
		":9005",
		grpcHandlerFunc(rpcServer, gwmux),
	)
}
// grpcHandlerFunc 根据请求头判断是grpc请求还是grpc-gateway请求
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("--------------")
		fmt.Println(1111)
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			fmt.Println("grpc")
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
			fmt.Println("http")
		}
		fmt.Println(3333)
	}), &http2.Server{})
}
