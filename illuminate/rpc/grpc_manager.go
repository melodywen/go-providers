package rpc

import (
	"context"
	"fmt"
	runtime2 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/melodywen/go-box/illuminate/contracts/bootstrap"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-box/illuminate/contracts/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

type GrpcManager struct {
	app    foundation.ApplicationInterface
	log    log.LoggerInterface
	config bootstrap.ConfigureInterface

	isInit bool

	ServerNux *runtime2.ServeMux
	Opt       []grpc.DialOption
	Endpoint  string
	Ctx       context.Context

	RpcServer *grpc.Server
}

func NewGrpcManager(app foundation.ApplicationInterface) *GrpcManager {
	var logger log.LoggerInterface
	var config bootstrap.ConfigureInterface

	manager := &GrpcManager{
		app:      app,
		log:      app.Make(&logger).(log.LoggerInterface),
		config:   app.Make(&config).(bootstrap.ConfigureInterface),
		Endpoint: "",
		Ctx:      context.Background(),
	}

	return manager
}

// NewServeMux 创建grpc-gateway服务，转发到grpc的9005端口
func (manager *GrpcManager) NewServeMux(opts ...runtime2.ServeMuxOption) {
	manager.ServerNux = runtime2.NewServeMux(opts...)
}

// NewServer creates a gRPC server which has no service registered and has not
// started to accept requests yet.
func (manager *GrpcManager) NewServer(opt ...grpc.ServerOption) {
	manager.RpcServer = grpc.NewServer(opt...)
}

// SetDialOption  DialOption configures how we set up the connection.
func (manager *GrpcManager) SetDialOption(opt []grpc.DialOption) {
	manager.Opt = opt
}

// SetEndpoint set entrypoint
func (manager *GrpcManager) SetEndpoint(endpoint string) {
	if endpoint == "" && manager.Endpoint == "" {
		endpoint = manager.config.GetString("grpc.server.endpoint")
	}
	if endpoint == "" && manager.Endpoint == "" {
		endpoint = ":8080"
	}
	manager.Endpoint = endpoint
}

// Init init
func (manager *GrpcManager) Init() {
	manager.SetEndpoint("")
	if manager.ServerNux == nil {
		manager.NewServeMux()
	}
	if manager.RpcServer == nil {
		manager.NewServer()
	}

	if len(manager.Opt) == 0 {
		manager.SetDialOption([]grpc.DialOption{grpc.WithInsecure()})
	}
	manager.isInit = true
}

// RegisterHandlerFromEndpoint Register register grpc endpoint
func (manager *GrpcManager) RegisterHandlerFromEndpoint(err error) {
	if err != nil {
		manager.log.Fatal("register handler from endpoint", map[string]interface{}{"err": err})
	}
}

// Run running server
func (manager *GrpcManager) Run() {
	if !manager.isInit {
		manager.log.Fatal("grpc manage please init", map[string]interface{}{})
	}
	fmt.Println(13123, manager.Endpoint)

	err := http.ListenAndServe(
		manager.Endpoint,
		grpcHandlerFunc(manager.RpcServer, manager.ServerNux),
	)
	if err != nil {
		manager.log.Fatal("listen server error", map[string]interface{}{"err": err})
	}
}

// grpcHandlerFunc 根据请求头判断是grpc请求还是grpc-gateway请求
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
