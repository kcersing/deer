package main

import (
	"common/mtl"
	"common/serversuite"
	"context"
	"fmt"
	"message/biz/events"
	"runtime"
	"time"

	message "gen/kitex_gen/message/messageservice"
	"message/biz/dal"
	"message/conf"
	"message/rpc"
	"os"
	"os/signal"
	"syscall"

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

	// 启动事件系统
	if err := events.Bootstrap(); err != nil {
		klog.Fatalf("Failed to bootstrap event system: %v", err)
	}

	address := conf.GetConf().Kitex.Address
	opts := serversuite.Option(registry, serviceName, address, snowflakeNode)
	svr := message.NewServer(new(MessageServiceImpl), opts...)

	// 优雅关闭
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		klog.Info("Shutdown signal received")

		// 停止 Kitex 服务
		if err := svr.Stop(); err != nil {
			klog.Errorf("Failed to stop server: %v", err)
		}

		// 停止事件系统
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := events.GetManager().Shutdown(ctx); err != nil {
			klog.Errorf("Failed to shutdown events: %v", err)
		}
	}()

	if err := svr.Run(); err != nil {
		klog.Fatalf("Server stopped with error: %v", err)
	}
}
