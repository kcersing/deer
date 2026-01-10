package registry

import (
	"common/consts"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hertz-contrib/registry/nacos"

	"net"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewRegisterNacos(namespaceId, rpcServiceName string) (registry.Registry, *registry.Info) {

	nacosClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId:         namespaceId,
				TimeoutMs:           5000,
				NotLoadCacheAtStart: true,
				LogDir:              consts.NacosLogDir,
				CacheDir:            consts.NacosCacheDir,
				LogLevel:            consts.NacosLogLevel,
			},
			ServerConfigs: []constant.ServerConfig{
				{
					IpAddr: consts.NacosIpAddr,
					Port:   consts.NacosPort,
				},
			},
		})
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}
	// 注册
	r := nacos.NewNacosRegistry(nacosClient, nacos.WithRegistryGroup(rpcServiceName))

	sf, err := snowflake.NewNode(2)
	if err != nil {
		klog.Fatalf("new snowflake node failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: rpcServiceName,
		Addr:        utils.NewNetAddr(consts.TCP, net.JoinHostPort(consts.NacosIpAddr, strconv.Itoa(consts.NacosPort))),
		Weight:      100,
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
