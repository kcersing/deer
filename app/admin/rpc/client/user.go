package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/user/userservice"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
)

var UserClient userservice.Client
var userOnceClient sync.Once

func InitUserRpc() {
	userOnceClient.Do(func() {

		r := client.NewResolver(serviceResolver, consts.UserRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)
		hlog.Info(r)
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
