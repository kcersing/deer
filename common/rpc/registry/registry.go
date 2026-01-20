package registry

import (
	"github.com/cloudwego/kitex/pkg/registry"
)

func Registry(r string) (re registry.Registry) {
	if r == "nacos" {
		re = NewRegisterNacos()
	} else if r == "etcd" {
		re = RegisterEtcd()
	}
	return re
}
