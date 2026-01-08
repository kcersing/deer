package client

import (
	"common/consts"
	"common/mw"
	"sync"

	"gen/kitex_gen/crm/crmservice"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var CrmClient crmservice.Client
var CrmOnceClient sync.Once

func initCrmRpc() {

	CrmOnceClient.Do(func() {

		r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
		if err != nil {
			hlog.Error("NewEtcdResolver err: %s", err)
			return
		}
		c, err := crmservice.NewClient(
			consts.CrmRpcServiceName,
			client.WithResolver(r), // resolver
			client.WithMuxConnection(1),
			client.WithRPCTimeout(3*time.Second),              // rpc timeout
			client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
			client.WithFailureRetry(retry.NewFailurePolicy()), // retry

			client.WithMiddleware(mw.CommonMiddleware),
			client.WithInstanceMW(mw.ClientsMiddleware),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "admin"}),
		)
		if err != nil {
			hlog.Error("NewClient err: %s", err)
			return
		}
		CrmClient = c

	})
}
