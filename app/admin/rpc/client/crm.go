package client

import (
	"common/consts"
	"common/rpc/client"
	"sync"

	"gen/kitex_gen/crm/crmservice"

	"github.com/cloudwego/kitex/pkg/klog"
)

var CrmClient crmservice.Client
var CrmOnceClient sync.Once

func InitCrmRpc() {

	CrmOnceClient.Do(func() {

		nr := client.NewNacosResolver(consts.CrmRpcServiceName, consts.CrmRpcServiceName)

		r := client.Resolver{
			R:                nr,
			ServiceName:      consts.CrmRpcServiceName,
			BasicServiceName: consts.AdminServiceName,
			EndpointAddress:  consts.OpenTelemetryAddress,
		}

		c, err := crmservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		CrmClient = c

	})
}
