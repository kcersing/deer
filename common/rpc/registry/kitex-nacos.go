package registry

import (
	"common/consts"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/kitex/pkg/registry"

	nacos "github.com/kitex-contrib/registry-nacos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NewRegisterNacos() registry.Registry {
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

	return r

}
