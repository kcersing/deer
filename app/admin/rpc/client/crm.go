package client

import (
	"common/consts"
	"common/rpc"
	"sync"

	"gen/kitex_gen/crm/crmservice"

	"github.com/cloudwego/kitex/pkg/klog"
)

var CrmClient crmservice.Client
var CrmOnceClient sync.Once

func InitCrmRpc() {

	CrmOnceClient.Do(func() {

		nr := rpc.NewNacosResolver("consts.NacosNamespaceId", consts.CrmRpcServiceName)

		r := rpc.Resolver{
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
