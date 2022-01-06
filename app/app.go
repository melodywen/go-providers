package app

import (
	"github.com/melodywen/go-box/illuminate/contracts/http"
	"github.com/melodywen/go-box/illuminate/foundation"
	http2 "github.com/melodywen/go-box/illuminate/foundation/http"
	"github.com/melodywen/go-providers/config"
	"os"
)

type Application struct {
	foundation.Application
}

var App *Application

func init() {
	App = NewApplication()
}

func NewApplication() *Application {
	dir, _ := os.Getwd()
	application := foundation.NewApplication(dir)
	application.Instance("eager-services", config.EagerServices)
	application.Instance("defer-services", config.DeferServices)

	//application.BootstrapOpenListen()
	var httpKernel http.KernelInterface
	application.Singleton(&httpKernel, http2.NewKernel)

	return &Application{
		Application: *application,
	}
}
