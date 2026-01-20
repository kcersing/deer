package main

import (
	"common/mtl"

	"common/serversuite"
	"crm/biz/dal"
	"crm/conf"
	"crm/rpc"
	crm "gen/kitex_gen/crm/crmservice"

	"github.com/cloudwego/kitex/pkg/klog"
)

func init() {
	dal.Init()
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

	svr := crm.NewServer(new(CrmServiceImpl), opts...)
	err := svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
