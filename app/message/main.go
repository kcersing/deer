package main

import (
	"common/mtl"
	"common/serversuite"
	message "gen/kitex_gen/message/messageservice"
	"message/biz/dal"
	"message/biz/dal/eventbus/mq"
	"message/conf"
	"message/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

func init() {
	dal.Init()
	mq.InitMQ()
}

var serviceName = conf.GetConf().Kitex.Service
var snowflakeNode = conf.GetConf().Kitex.Node
var registry = conf.GetConf().Kitex.Registry

func main() {

	mtl.InitFlightRecorder()

	mtl.InitLog(false)

	mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	mtl.InitProvider(serviceName)

	rpc.Init()

	address := conf.GetConf().Kitex.Address

	opts := serversuite.Option(registry, serviceName, address, snowflakeNode)

	svr := message.NewServer(new(MessageServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
