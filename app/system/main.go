package main

import (
	"common/mtl"
	"common/serversuite"
	"fmt"
	system "gen/kitex_gen/system/systemservice"
	"os"
	"runtime"

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

	mtl.InitFlightRecorder()

	mtl.InitLog(false)

	mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, "9090", "127.0.0.1")

	mtl.InitProvider(serviceName)

	rpc.Init()

	address := conf.GetConf().Kitex.Address

	opts := serversuite.Option(registry, serviceName, address, snowflakeNode)

	svr := system.NewServer(new(SystemServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
