package client

import (
	"admin/conf"
)

var serviceResolver = conf.GetConf().Hertz.Resolver
