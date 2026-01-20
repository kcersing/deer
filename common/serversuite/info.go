package serversuite

import (
	"common/consts"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
)

func GetInfo(serviceName, address string, flakeNode int64) (rr *registry.Info) {
	sf, err := snowflake.NewNode(flakeNode)
	if err != nil {
		klog.Fatalf("new snowflake node failed: %s", err.Error())
		return nil
	}

	info := &registry.Info{
		ServiceName: serviceName,
		Addr:        utils.NewNetAddr(consts.TCP, address),
		Weight:      100,
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return info
}
