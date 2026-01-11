package registry

import (
	"common/consts"

	"github.com/bwmarrin/snowflake"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils"

	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/registry"
	nacos "github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewRegisterNacos(namespaceId, rpcServiceName string, flakeNode int64) (registry.Registry, *registry.Info) {

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
		klog.Fatalf("new  registry nacos client failed: %s", err.Error())
	}
	// 注册
	r := nacos.NewNacosRegistry(nacosClient, nacos.WithGroup(rpcServiceName))

	sf, err := snowflake.NewNode(flakeNode)
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
