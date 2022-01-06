package providers

import (
	"github.com/melodywen/go-box/illuminate/support"
	"github.com/melodywen/go-providers/illuminate/rpc"
	"github.com/melodywen/go-providers/servers/gen/golang"
	"github.com/melodywen/go-providers/servers/implement_servers"
)

type AppServiceProvider struct {
	support.ServiceProvider
}

func NewAppServiceProvider() *AppServiceProvider {
	return &AppServiceProvider{}
}

// Boot Bootstrap any application services.
func (provider *AppServiceProvider) Boot() {
	var manager *rpc.GrpcManager
	manager = provider.App.Make(manager).(*rpc.GrpcManager)
	manager.Init()

	// 注册
	manager.RegisterHandlerFromEndpoint(golang.RegisterBlogHandlerFromEndpoint(manager.Ctx, manager.ServerNux, manager.Endpoint, manager.Opt))
	golang.RegisterBlogServer(manager.RpcServer, implement_servers.NewBlogImplementServer())

}

// Register any application services.
func (provider *AppServiceProvider) Register() {

}
