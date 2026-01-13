package main

import (
	"common/mtl"
	"common/serversuite"
	system "gen/kitex_gen/system/systemservice"

	"system/biz/dal"
	"system/conf"
	"system/rpc"

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

	//mtl.InitMetric(serviceName, "9090", "127.0.0.1")

	mtl.InitProvider(serviceName)

	rpc.Init()

	address := conf.GetConf().Kitex.Address

	opts := serversuite.Option(serviceName, address, snowflakeNode)

	svr := system.NewServer(new(SystemServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
