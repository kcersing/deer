package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/system/systemservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var SystemClient systemservice.Client
var systemOnceClient sync.Once

func InitSystemRpc() {
	systemOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.SystemRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)

		c, err := systemservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		SystemClient = c

	})
}
