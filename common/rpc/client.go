package rpc

import (
	"common/consts"
	"common/mw"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"

	nacos "github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Resolver struct {
	R                discovery.Resolver
	ServiceName      string
	BasicServiceName string
	EndpointAddress  string
}

func NewEtcdResolver(etcdAddress string) discovery.Resolver {
	r, err := etcd.NewEtcdResolver([]string{etcdAddress})
	if err != nil {
		hlog.Error("NewEtcdResolver err: %s", err)
	}
	return r
}
func NewNacosResolver(namespaceId string, serviceName string) discovery.Resolver {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				TimeoutMs:           5000,
				NotLoadCacheAtStart: true,
				LogDir:              consts.NacosLogDir,
				LogLevel:            consts.NacosLogLevel,
				CacheDir:            consts.NacosCacheDir,
				NamespaceId:         namespaceId,
			},
			ServerConfigs: []constant.ServerConfig{
				{
					IpAddr: consts.NacosIpAddr,
					Port:   consts.NacosPort,
				},
			},
		},
	)
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}
	r := nacos.NewNacosResolver(
		namingClient,
		nacos.WithGroup(serviceName),
	)
	return r
}

func (r Resolver) NewOpenTelemetryProvider() {
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(r.ServiceName),
		provider.WithExportEndpoint(r.EndpointAddress),
		provider.WithInsecure(),
	)
}
func (r Resolver) Options() []client.Option {
	return []client.Option{
		client.WithResolver(r.R),                                   // resolver
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		//client.WithMuxConnection(1),                       // multiplexing
		client.WithRPCTimeout(3 * time.Second),            // rpc timeout
		client.WithConnectTimeout(50 * time.Millisecond),  // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry

		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientsMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: r.BasicServiceName}),
	}
}
