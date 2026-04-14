package client

import (
	"common/consts"
	"common/rpc/client"
	"sync"

	"gen/kitex_gen/crm/crmservice"

	"github.com/cloudwego/kitex/pkg/klog"
)

var PaymentClient crmservice.Client
var PaymentOnceClient sync.Once

func InitPaymentRpc() {

	PaymentOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.PaymentRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)

		c, err := crmservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		PaymentClient = c

	})
}
