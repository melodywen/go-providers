package main

import (
	"github.com/melodywen/go-box/illuminate/contracts/http"
	"github.com/melodywen/go-providers/app"
	"github.com/melodywen/go-providers/illuminate/contracts/rpc"
	"github.com/sirupsen/logrus"
)

func main() {
	var httpKernel http.KernelInterface
	var ok bool

	k := app.App.Make(&httpKernel)
	if httpKernel, ok = k.(http.KernelInterface); !ok {
		logrus.Panicln("获取 http kernel 失败")
	}
	httpKernel.Handle()

	var grpc rpc.GrpcManagerInterface
	grpc = app.App.Make("grpc-manager").(rpc.GrpcManagerInterface)
	grpc.Run()
}
