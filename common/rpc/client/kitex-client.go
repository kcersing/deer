package client

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

	nacos "github.com/kitex-contrib/registry-nacos/v2/resolver"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func GetResolver(r, serviceName string) (re discovery.Resolver) {

	if r == "nacos" {
		re = NewNacosResolver(serviceName)
	} else if r == "etcd" {
		re = NewEtcdResolver(consts.EtcdAddress)
	}
	return re
}

func NewEtcdResolver(etcdAddress string) discovery.Resolver {
	r, err := etcd.NewEtcdResolver([]string{etcdAddress})
	if err != nil {
		hlog.Error("NewEtcdResolver err: %s", err)
	}
	return r
}
func NewNacosResolver(serviceName string) discovery.Resolver {
	// the nacos server config
	sc := []constant.ServerConfig{

		*constant.NewServerConfig(consts.NacosIpAddr, consts.NacosPort,
			func(sc *constant.ServerConfig) {
				// it is not recommended to modify GRPC-Port unless necessary,default by server port +1000
				sc.GrpcPort = 9848
			}),
	}
	// the nacos client config
	cc := &constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		LogLevel:            consts.NacosLogLevel,
		CacheDir:            consts.NacosCacheDir,
		NamespaceId:         "public",
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}

	r := nacos.NewNacosResolver(
		namingClient,
		nacos.WithGroup(serviceName),
	)
	hlog.Info(r.Name())
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
