package main

import (
	"common/consts"
	"common/mtl"
	"common/mw"
	"common/pkg/utils"
	"common/serversuite"
	"crm/biz/dal"
	"crm/conf"
	"crm/rpc"
	crm "gen/kitex_gen/crm/crmservice"
	"net"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
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

	svr := crm.NewServer(new(CrmServiceImpl), opts...)
	err := svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
