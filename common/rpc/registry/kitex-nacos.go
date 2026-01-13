package registry

import (
	"common/consts"
	utils2 "common/pkg/utils"

	"net"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"

	nacos "github.com/kitex-contrib/registry-nacos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NewRegisterNacos(rpcServiceName string, flakeNode int64, port string) (registry.Registry, *registry.Info) {
	sc := []constant.ServerConfig{
		{
			IpAddr: consts.NacosIpAddr,
			Port:   consts.NacosPort,
		},
	}
	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		CacheDir:            consts.NacosCacheDir,
		LogLevel:            consts.NacosLogLevel,
	}
	nacosClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		})

	if err != nil {
		klog.Fatalf("new  registry nacos client failed: %s", err.Error())
	}
	// 注册
	r := nacos.NewNacosRegistry(nacosClient, nacos.WithGroup("api"))

	sf, err := snowflake.NewNode(flakeNode)
	if err != nil {
		klog.Fatalf("new snowflake node failed: %s", err.Error())
	}

	localIp := utils2.MustGetLocalIPv4()
	info := &registry.Info{
		ServiceName: rpcServiceName,
		Addr:        utils.NewNetAddr(consts.TCP, net.JoinHostPort(localIp, port)),
		Weight:      100,
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}

	return r, info

}
