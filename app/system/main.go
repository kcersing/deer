package main

import (
	"common/mtl"
	"common/serversuite"
	"fmt"
	"os"
	"runtime"

	system "gen/kitex_gen/system/systemservice"
	"system/biz/dal"
	"system/conf"
	"system/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/cobra"
)

var serviceName = conf.GetConf().Kitex.Service
var snowflakeNode = conf.GetConf().Kitex.Node
var registry = conf.GetConf().Kitex.Registry

var (
	version   string
	buildTime string
	commit    string
	goVersion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

var rootCmd = &cobra.Command{
	Use:   "system",
	Short: "System service for deer platform",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 初始化链路追踪
		mtl.InitTracing(serviceName)
		// 2. 初始化日志
		mtl.InitLog(conf.LogLevel())
		// 3. 初始化指标
		//metricsPort := fmt.Sprintf(":%d", conf.MetricsPort)
		//mtl.InitMetric(serviceName+"-metrics", metricsPort, conf.EtcdAddress)
		metricsPort := fmt.Sprintf(":%d", "9090")
		mtl.InitMetric(serviceName+"-metrics", metricsPort, "127.0.0.1:2379")
		// 4. 初始化飞行记录器
		mtl.InitFlightRecorder()

		// 5. 业务依赖初始化 (放在 Run 里面，避免 --version 时无谓初始化)
		dal.Init()
		rpc.Init()

		address := conf.GetConf().Kitex.Address
		opts := serversuite.Option(registry, serviceName, address, snowflakeNode)
		svr := system.NewServer(new(SystemServiceImpl), opts...)

		if err := svr.Run(); err != nil {
			// 使用 klog.Error 而不是 Fatal，让程序能正常走完 main 函数触发 ShutdownHook
			klog.Error(err)
		}
	},
}

func init() {
	if version == "" {
		version = "unknown"
	}
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(
		fmt.Sprintf("Git Commit Hash: %s\nUTC Build Time : %s\nVersion : %s\nGolang Version : %s\n",
			commit, buildTime, version, goVersion),
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1) // 如果是 cobra 本身解析失败（如参数错误），可以 Exit
	}
}
