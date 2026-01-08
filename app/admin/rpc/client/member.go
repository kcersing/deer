package client

import (
	"common/consts"
	"common/mw"
	"sync"

	"gen/kitex_gen/member/memberservice"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"

	nacos "github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var MemberClient memberservice.Client
var memberOnceClient sync.Once

func initMemberRpc() {
	memberOnceClient.Do(func() {

		r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
		if err != nil {
			hlog.Error("NewEtcdResolver err: %s", err)
			return
		}
		c, err := memberservice.NewClient(
			consts.MemberRpcServiceName,
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

		MemberClient = c

	})
}
func initMemberNocosRpc() {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				TimeoutMs:           5000,
				NotLoadCacheAtStart: true,
				LogDir:              "",
				LogLevel:            "",
				CacheDir:            "",
				NamespaceId:         "",
			},
			ServerConfigs: []constant.ServerConfig{
				{
					IpAddr: "101.126.9.226",
					Port:   8848,
				},
			},
		},
	)
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}
	r := nacos.NewNacosResolver(
		namingClient,
		nacos.WithGroup(consts.MemberRpcServiceName),
	)
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.MemberRpcServiceName),
		provider.WithExportEndpoint(consts.OtelExporterEndpoint),
		provider.WithInsecure(),
	)
	c, err := memberservice.NewClient(
		consts.MemberRpcServiceName,
		client.WithResolver(r),                                     // resolver
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		//client.WithMuxConnection(1),                       // multiplexing
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry

		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientsMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "admin"}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	MemberClient = c
	return
}
