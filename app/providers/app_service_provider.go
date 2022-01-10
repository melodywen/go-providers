package providers

import (
	"context"
	"fmt"
	"github.com/melodywen/go-box/illuminate/support"
	rpc2 "github.com/melodywen/go-providers/illuminate/contracts/rpc"
	"github.com/melodywen/go-providers/illuminate/rpc"
	"github.com/melodywen/go-providers/servers/gen/golang"
	"github.com/melodywen/go-providers/servers/implement_servers"
	"google.golang.org/grpc"
)

type AppServiceProvider struct {
	support.ServiceProvider
}

func NewAppServiceProvider() *AppServiceProvider {
	return &AppServiceProvider{}
}

func testMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)

	fmt.Println(resp)
	return resp, err
}

// Boot Bootstrap any application services.
func (provider *AppServiceProvider) Boot() {
	var manager *rpc.GrpcManager
	manager = provider.App.Make(manager).(*rpc.GrpcManager)

	manager.NewServer(grpc.UnaryInterceptor(testMiddleware))
	manager.Init()

	// 注册 http
	manager.SetRegisterHandlerFromEndpointCallbacks([]rpc2.RegisterHandlerFromEndpointFun{
		golang.RegisterBlogHandlerFromEndpoint,
	})

	//  注册 grpc
	golang.RegisterBlogServer(manager.RpcServer, implement_servers.NewBlogImplementServer())
}

// Register any application services.
func (provider *AppServiceProvider) Register() {

}
