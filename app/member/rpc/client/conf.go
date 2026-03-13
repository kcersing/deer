package client

import (
	"member/conf"
)

var serviceResolver = conf.GetConf().Kitex.Resolver
