package client

import (
	"admin/conf"
	"common/consts"
	"common/rpc/client"
	"sync"

	"gen/kitex_gen/crm/crmservice"

	"github.com/cloudwego/kitex/pkg/klog"
)

var CrmClient crmservice.Client
var CrmOnceClient sync.Once

var serviceResolver = conf.GetConf().Hertz.Resolver

func InitCrmRpc() {

	CrmOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.CrmRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)

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
