package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/order/orderservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var OrderClient orderservice.Client
var orderOnceClient sync.Once

func InitOrderRpc() {

	orderOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.OrderRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)

		c, err := orderservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		OrderClient = c

	})
}
