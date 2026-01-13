package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/user/userservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var UserClient userservice.Client
var userOnceClient sync.Once

func InitUserRpc() {
	userOnceClient.Do(func() {

		nr := client.NewNacosResolver(consts.UserRpcServiceName, consts.UserRpcServiceName)

		r := client.Resolver{
			R:                nr,
			ServiceName:      consts.UserRpcServiceName,
			BasicServiceName: consts.AdminServiceName,
			EndpointAddress:  consts.OpenTelemetryAddress,
		}

		c, err := userservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}

		UserClient = c

	})
}
