package main

import (
	"common/mtl"
	"common/serversuite"
	"context"
	"message/biz/events"
	"time"

	message "gen/kitex_gen/message/messageservice"
	"message/biz/dal"
	"message/biz/dal/mq"
	"message/conf"
	"message/rpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudwego/kitex/pkg/klog"
)

func init() {
	dal.Init()
	mq.InitMQ()
}

var serviceName = conf.GetConf().Kitex.Service
var snowflakeNode = conf.GetConf().Kitex.Node
var registry = conf.GetConf().Kitex.Registry

func main() {
	mtl.InitFlightRecorder()
	mtl.InitLog(false)
	mtl.InitTracing(serviceName)
	mtl.InitProvider(serviceName)
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
		if err := events.Shutdown(ctx); err != nil {
			klog.Errorf("Failed to shutdown events: %v", err)
		}
	}()

	if err := svr.Run(); err != nil {
		klog.Fatalf("Server stopped with error: %v", err)
	}
}
