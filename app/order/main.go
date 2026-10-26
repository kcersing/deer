package main

import (
	"common/mtl"
	"common/serversuite"
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	order "gen/kitex_gen/order/orderservice"
	"order/biz/dal"
	"order/biz/infras"
	"order/conf"
	"order/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server" // ⚠️ 引入 server 包用于注册钩子
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
	Use:   "order",
	Short: "Order service for deer platform",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 初始化链路追踪
		mtl.InitTracing(serviceName)
		// 2. 初始化日志 (依赖前一步的 Tracer)
		mtl.InitLog(conf.LogLevel())
		// 3. 初始化指标
		//metricsPort := fmt.Sprintf(":%d", conf.MetricsPort)
		//mtl.InitMetric(serviceName+"-metrics", metricsPort, conf.EtcdAddress)
		metricsPort := fmt.Sprintf(":%d", "9090")
		mtl.InitMetric(serviceName+"-metrics", metricsPort, "127.0.0.1:2379")
		// 4. 初始化飞行记录器
		mtl.InitFlightRecorder()

		// 5. 业务依赖初始化
		dal.Init()
		rpc.Init()

		// 6. 启动事件系统 (必须在 Kitex Server 启动前就绪)
		if err := infras.Bootstrap(); err != nil {
			klog.Fatalf("Failed to bootstrap event system: %v", err)
		}

		// 7. 注册事件系统的优雅关闭钩子
		// 交给 Kitex 底层监听信号，收到信号后会自动按顺序执行这些 Hook
		server.RegisterShutdownHook(func() {
			klog.Info("Executing ShutdownHooks: closing event system...")
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			if err := infras.GetManager().Shutdown(ctx); err != nil {
				klog.Errorf("Failed to shutdown events: %v", err)
			}
		})

		// 8. 启动 Kitex 服务
		address := conf.GetConf().Kitex.Address
		opts := serversuite.Option(registry, serviceName, address, snowflakeNode)
		svr := order.NewServer(new(OrderServiceImpl), opts...)

		if err := svr.Run(); err != nil {
			// 使用 klog.Error，让程序自然走完，触发上面的 ShutdownHook 和日志 flush
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
		os.Exit(1)
	}
}
