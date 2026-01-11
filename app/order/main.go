package main

import (
	"common/consts"
	"common/mtl"
	"common/mw"
	"common/pkg/utils"
	order "gen/kitex_gen/order/orderservice"
	"net"
	"order/biz/dal"
	"order/conf"
	"order/rpc"
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

func main() {

	mtl.InitFlightRecorder()

	mtl.InitLog(false)

	mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	mtl.InitProvider(serviceName)

	rpc.Init()

	opts := kitexInit()

	svr := order.NewServer(new(OrderServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
func kitexInit() (opts []server.Option) {

	r, info := rpc.Registry()

	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, address)
	if err != nil {
		panic(err)
	}

	opts = append(opts,
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		//server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),

		server.WithRegistry(r),
		server.WithRegistryInfo(info),
	)

	return
}
