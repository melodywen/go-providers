package config

import (
	"github.com/melodywen/go-box/illuminate/contracts/support"
	"github.com/melodywen/go-providers/app/providers"
	"github.com/melodywen/go-providers/illuminate/rpc"
)

// EagerServices eager services
var EagerServices []support.ServiceProviderInterface

// DeferServices defer services
var DeferServices map[string]support.ServiceProviderInterface

func init() {
	EagerServices = []support.ServiceProviderInterface{
		rpc.NewGRpcServiceProvider(),
		providers.NewAppServiceProvider(),
		//providers.NewSchoolServiceProvider(),
	}
	DeferServices = map[string]support.ServiceProviderInterface{
		//"teacher": providers.NewTeacherServiceProvider(),
		//"student": providers.NewStudentServiceProvider(),
	}
}
