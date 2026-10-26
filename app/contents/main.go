package main

import (
	"common/mtl"
	"fmt"
	"gen/kitex_gen/contents/contentsservice"
	"os"
	"runtime"

	"common/serversuite"
	"contents/biz/dal"
	"contents/conf"
	"contents/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
)

func init() {
	dal.Init()
}

var serviceName = conf.GetConf().Kitex.Service
var snowflakeNode = conf.GetConf().Kitex.Node
var registry = conf.GetConf().Kitex.Registry

var (
	version   string // Git提交哈希值
	buildTime string // 构建时间
	commit    string
	goVersion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH) // Go编译器版本
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Commit Hash: %s\n", commit)
		fmt.Printf("UTC Build Time : %s\n", buildTime)
		fmt.Printf("Version : %s\n", version)
		fmt.Printf("Golang Version : %s\n", goVersion)
		return
	}

	// 1. 初始化链路追踪 (TracerProvider)
	mtl.InitTracing(serviceName)
	// 2. 初始化日志 (依赖前一步的 Tracer，用于注入 trace_id)
	mtl.InitLog(conf.LogLevel())
	// 3. 初始化指标 (配置化 Etcd 地址和端口)
	//metricsPort := fmt.Sprintf(":%d", conf.MetricsPort)
	//mtl.InitMetric(serviceName+"-metrics", metricsPort, conf.EtcdAddress)
	metricsPort := fmt.Sprintf(":%d", "9090")
	mtl.InitMetric(serviceName+"-metrics", metricsPort, "127.0.0.1:2379")
	// 4. 初始化飞行记录器 (排查突发性能问题)
	mtl.InitFlightRecorder()

	rpc.Init()

	address := conf.GetConf().Kitex.Address

	opts := serversuite.Option(registry, serviceName, address, snowflakeNode)

	svr := contentsservice.NewServer(new(ContentsServiceImpl), opts...)
	err := svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
