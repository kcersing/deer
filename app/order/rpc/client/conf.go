package client

import (
	"order/conf"
)

var serviceResolver = conf.GetConf().Kitex.Resolver
