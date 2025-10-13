package main

import (
	"common/mtl"
	"common/mw"
	"common/pkg/utils"
	user "gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net"
	"strings"
	"user/conf"
	"user/rpc"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load()

	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	//mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	opts := kitexInit()

	rpc.Init()

	svr := user.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
func kitexInit() (opts []server.Option) {

	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}

	r, err := etcd.NewEtcdRegistry([]string{conf.GetConf().Etcd.Address})
	opts = append(opts,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		//server.WithSuite(tracing.NewServerSuite()),
	)

	return
}
