package client

import (
	"common/consts"
	"common/mw"
	"gen/kitex_gen/system/systemservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var SystemClient systemservice.Client
var SystemOnceClient sync.Once

func initSystemRpc() {
	SystemOnceClient.Do(func() {
		r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
		if err != nil {
			panic(err)
		}
		c, err := systemservice.NewClient(
			consts.SystemRpcServiceName,
			client.WithResolver(r), // resolver
			client.WithMuxConnection(1),
			client.WithRPCTimeout(3*time.Second),              // rpc timeout
			client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
			client.WithFailureRetry(retry.NewFailurePolicy()), // retry

			client.WithMiddleware(mw.CommonMiddleware),
			client.WithInstanceMW(mw.ClientsMiddleware),
			client.WithSuite(tracing.NewClientSuite()),
			client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "facade"}),
		)
		if err != nil {
			panic(err)
		}
		SystemClient = c
	})
}
