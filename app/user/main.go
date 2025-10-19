package main

import (
	"common/consts"
	"common/mtl"
	"common/mw"
	"common/pkg/utils"
	user "gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net"
	"os"
	"path"
	"strings"
	"time"
	"user/biz/dal"
	"user/conf"
	"user/rpc"
)

func init() {
	dal.Init()
}

var serviceName = conf.GetConf().Kitex.Service
 
func main() {
	_ = godotenv.Load()

	logFilePath := consts.LogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format(time.DateOnly) + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}
	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})

	//mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	//mtl.InitProvider(serviceName)

	opts := kitexInit()

	rpc.Init()

	svr := user.NewServer(new(UserServiceImpl), opts...)

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

	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress})
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
