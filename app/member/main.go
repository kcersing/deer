package main

import (
	"common/mtl"
	"common/serversuite"
	"fmt"
	member "gen/kitex_gen/member/memberservice"
	"member/biz/dal"
	"member/conf"
	"member/rpc"
	"os"
	"runtime"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/cobra"
)

var serviceName = conf.GetConf().Kitex.Service
var snowflakeNode = conf.GetConf().Kitex.Node
var registry = conf.GetConf().Kitex.Registry

var (
	version   string // Git提交哈希值
	buildTime string // 构建时间
	commit    string
	goVersion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH) // Go编译器版本
)

var rootCmd = &cobra.Command{
	Use:   "member",
	Short: "member service for deer platform",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 初始化链路追踪 (TracerProvider)
		mtl.InitTracing(serviceName)
		// 2. 初始化日志 (依赖前一步的 Tracer，用于注入 trace_id)
		mtl.InitLog(conf.LogLevel())
		// 3. 初始化指标
		//metricsPort := fmt.Sprintf(":%d", conf.MetricsPort)
		//mtl.InitMetric(serviceName+"-metrics", metricsPort, conf.EtcdAddress)
		metricsPort := fmt.Sprintf(":%d", "9090")
		mtl.InitMetric(serviceName+"-metrics", metricsPort, "127.0.0.1:2379")
		// 4. 初始化飞行记录器
		mtl.InitFlightRecorder()

		dal.Init()
		rpc.Init()

		address := conf.GetConf().Kitex.Address
		opts := serversuite.Option(registry, serviceName, address, snowflakeNode)
		svr := member.NewServer(new(MemberServiceImpl), opts...)

		if err := svr.Run(); err != nil {
			klog.Fatal(err)
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
