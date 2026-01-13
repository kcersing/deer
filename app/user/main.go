package main

import (
	"common/mtl"
	"common/serversuite"
	user "gen/kitex_gen/user/userservice"
	"user/biz/dal"
	"user/conf"
	"user/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

func init() {
	dal.Init()
}

var serviceName = conf.GetConf().Kitex.Service
var snowflakeNode = conf.GetConf().Kitex.Node

func main() {

	mtl.InitFlightRecorder()

	mtl.InitLog(false)

	mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	mtl.InitProvider(serviceName)

	rpc.Init()

	address := conf.GetConf().Kitex.Address

	opts := serversuite.Option(serviceName, address, snowflakeNode)

	svr := user.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
