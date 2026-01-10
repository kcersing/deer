package client

import (
	"common/consts"
	"common/rpc"
	"gen/kitex_gen/product/productservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var ProductClient productservice.Client
var productOnceClient sync.Once

func InitProductRpc() {
	productOnceClient.Do(func() {

		nr := rpc.NewNacosResolver("consts.NacosNamespaceId", consts.ProductRpcServiceName)

		r := rpc.Resolver{
			R:                nr,
			ServiceName:      consts.ProductRpcServiceName,
			BasicServiceName: consts.AdminServiceName,
			EndpointAddress:  consts.OpenTelemetryAddress,
		}

		c, err := productservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		ProductClient = c

	})
}
