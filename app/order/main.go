package main

import (
	"common/mtl"
	"common/serversuite"
	"context"
	"fmt"
	order "gen/kitex_gen/order/orderservice"
	"order/biz/dal"
	"order/biz/infras"
	"order/conf"
	"order/rpc"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

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

	//mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	mtl.InitProvider(serviceName)

	rpc.Init()

	address := conf.GetConf().Kitex.Address

	opts := serversuite.Option(registry, serviceName, address, snowflakeNode)

	svr := order.NewServer(new(OrderServiceImpl), opts...)

	// 启动事件系统
	if err := infras.Bootstrap(); err != nil {
		klog.Fatalf("Failed to bootstrap event system: %v", err)
	}

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

		if err := infras.GetManager().Shutdown(ctx); err != nil {
			klog.Errorf("Failed to shutdown events: %v", err)
		}
	}()

	if err := svr.Run(); err != nil {
		klog.Fatalf("Server stopped with error: %v", err)
	}

}
