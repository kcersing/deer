package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/product/productservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var ProductClient productservice.Client
var productOnceClient sync.Once

func InitProductRpc() {
	productOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.ProductRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)

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
