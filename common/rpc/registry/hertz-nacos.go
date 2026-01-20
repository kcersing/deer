package registry

import (
	"common/consts"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/registry/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NewHertzRegisterNacos(serviceName string, flakeNode int64, port int) registry.Registry {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(consts.NacosIpAddr, consts.NacosPort),
	}

	cc := constant.ClientConfig{

		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		LogLevel:            consts.NacosLogLevel,
		CacheDir:            consts.NacosCacheDir,
		NamespaceId:         "public",
	}
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		hlog.Fatalf("new nacos client failed: %s", err.Error())
	}

	r := nacos.NewNacosRegistry(cli, nacos.WithRegistryGroup("facade"))

	return r

}
