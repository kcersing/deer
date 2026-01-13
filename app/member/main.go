package main

import (
	"common/mtl"
	"common/serversuite"
	member "gen/kitex_gen/member/memberservice"
	"member/biz/dal"
	"member/conf"
	"member/rpc"

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

	svr := member.NewServer(new(MemberServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
