package client

import (
	"payment/conf"
)

var serviceResolver = conf.GetConf().Kitex.Resolver
