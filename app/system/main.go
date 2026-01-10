package main

import (
	"common/consts"
	"common/mtl"
	"common/mw"
	"common/pkg/utils"
	system "gen/kitex_gen/system/systemservice"

	"net"

	"strings"
	"system/biz/dal"
	"system/conf"
	"system/rpc"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"
)

func init() {
	dal.Init()
}

var serviceName = conf.GetConf().Kitex.Service

func main() {

	mtl.InitFlightRecorder()

	mtl.InitLog(false)

	mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, "9090", "127.0.0.1")

	mtl.InitProvider(serviceName)

	opts := kitexInit()

	rpc.Init()

	svr := system.NewServer(new(SystemServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
func kitexInit() (opts []server.Option) {

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
	retryConfig := retry.NewRetryConfig(
		retry.WithMaxAttemptTimes(10),
		retry.WithObserveDelay(20*time.Second),
		retry.WithRetryDelay(5*time.Second),
	)
	r, err := etcd.NewEtcdRegistryWithRetry([]string{consts.EtcdAddress}, retryConfig)
	r, info := rpc.Registry()
	rpc.Init()
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
