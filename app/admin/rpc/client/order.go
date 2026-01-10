package client

import (
	"common/consts"
	"common/rpc"
	"gen/kitex_gen/order/orderservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var OrderClient orderservice.Client
var orderOnceClient sync.Once

func InitOrderRpc() {

	orderOnceClient.Do(func() {

		nr := rpc.NewNacosResolver("consts.NacosNamespaceId", consts.OrderRpcServiceName)

		r := rpc.Resolver{
			R:                nr,
			ServiceName:      consts.OrderRpcServiceName,
			BasicServiceName: consts.AdminServiceName,
			EndpointAddress:  consts.OpenTelemetryAddress,
		}

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
