package client

import (
	"common/consts"
	"common/mw"
	"gen/kitex_gen/product/productservice"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var ProductClient productservice.Client
var productOnceClient sync.Once

func initProductRpc() {
	productOnceClient.Do(func() {
		r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
		if err != nil {
			hlog.Error("NewEtcdResolver err: %s", err)
			return
		}
		c, err := productservice.NewClient(
			consts.ProductRpcServiceName,
			client.WithResolver(r), // resolver
			//client.WithMuxConnection(1),
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
		ProductClient = c
	})
}
