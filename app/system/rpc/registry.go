package rpc

import (
	"common/consts"
	"common/rpc/registry"

	registry2 "github.com/cloudwego/kitex/pkg/registry"
)

func Registry(address string) (registry2.Registry, *registry2.Info) {
	return registry.NewRegisterNacos(consts.SystemRpcServiceName, 2, address)
}
