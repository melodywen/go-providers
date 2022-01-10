package rpc

import (
	"context"
	runtime2 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/melodywen/go-box/illuminate/contracts/bootstrap"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-box/illuminate/contracts/log"
	"github.com/melodywen/go-providers/illuminate/contracts/rpc"
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

	serverNux *runtime2.ServeMux
	opt       []grpc.DialOption
	endpoint  string
	Ctx       context.Context

	RpcServer *grpc.Server

	registerHandlerFromEndpointSlice []rpc.RegisterHandlerFromEndpointFun
}

func NewGrpcManager(app foundation.ApplicationInterface) *GrpcManager {
	var logger log.LoggerInterface
	var config bootstrap.ConfigureInterface

	manager := &GrpcManager{
		app:      app,
		log:      app.Make(&logger).(log.LoggerInterface),
		config:   app.Make(&config).(bootstrap.ConfigureInterface),
		endpoint: "",
		Ctx:      context.Background(),
	}

	return manager
}

// NewServeMux 创建grpc-gateway服务，转发到grpc的9005端口
func (manager *GrpcManager) NewServeMux(opts ...runtime2.ServeMuxOption) {
	manager.serverNux = runtime2.NewServeMux(opts...)
}

// NewServer creates a gRPC server which has no service registered and has not
// started to accept requests yet.
func (manager *GrpcManager) NewServer(opt ...grpc.ServerOption) {
	manager.RpcServer = grpc.NewServer(opt...)
}

// SetDialOption  DialOption configures how we set up the connection.
func (manager *GrpcManager) SetDialOption(opt []grpc.DialOption) {
	manager.opt = opt
}

// SetEndpoint set entrypoint
func (manager *GrpcManager) SetEndpoint(endpoint string) {
	if endpoint == "" && manager.endpoint == "" {
		endpoint = manager.config.GetString("grpc.server.endpoint")
	}
	if endpoint == "" && manager.endpoint == "" {
		endpoint = ":8080"
	}
	manager.endpoint = endpoint
}

// Init init
func (manager *GrpcManager) Init() {
	manager.SetEndpoint("")
	if manager.serverNux == nil {
		manager.NewServeMux()
	}
	if manager.RpcServer == nil {
		manager.NewServer()
	}

	if len(manager.opt) == 0 {
		manager.SetDialOption([]grpc.DialOption{grpc.WithInsecure()})
	}
	manager.isInit = true
}

// SetRegisterHandlerFromEndpointCallbacks set register handler from endpoint
func (manager *GrpcManager) SetRegisterHandlerFromEndpointCallbacks(callbacks []rpc.RegisterHandlerFromEndpointFun) {
	manager.registerHandlerFromEndpointSlice = append(manager.registerHandlerFromEndpointSlice, callbacks...)
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

	// register grpc handler from endpoint
	for _, fun := range manager.registerHandlerFromEndpointSlice {
		manager.RegisterHandlerFromEndpoint(fun(manager.Ctx, manager.serverNux, manager.endpoint, manager.opt))
		funName := manager.app.AbstractToString(fun)
		manager.log.Info("register handler from endpoint ", map[string]interface{}{"name": funName})
	}

	manager.log.Info("grpc start success", map[string]interface{}{"endpoint": manager.endpoint})

	err := http.ListenAndServe(
		manager.endpoint,
		grpcHandlerFunc(manager.RpcServer, manager.serverNux),
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
