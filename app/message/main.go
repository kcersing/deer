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
	events.InitGlobalEventBus()

}

var serviceName = conf.GetConf().Kitex.Service
var snowflakeNode = conf.GetConf().Kitex.Node
var registry = conf.GetConf().Kitex.Registry

func main() {

	mtl.InitFlightRecorder()

	mtl.InitLog(false)

	mtl.InitTracing(serviceName)

	//mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])

	mtl.InitProvider(serviceName)

	rpc.Init()

	address := conf.GetConf().Kitex.Address

	opts := serversuite.Option(registry, serviceName, address, snowflakeNode)

	svr := message.NewServer(new(MessageServiceImpl), opts...)

	if err := events.Bootstrap(); err != nil {
		klog.Fatal(err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	if err := svr.Run(); err != nil {
		klog.Fatal(err)
	}

	go func() {

		<-sigChan
		klog.Infof("shutdown signal received")
		if err := svr.Stop(); err != nil {
			klog.Errorf("failed to stop server：%v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := events.Shutdown(ctx); err != nil {
			klog.Infof("failed to shutdown events: %v", err)
		}
	}()

}
