package mq

import (
	"common/amqpclt"
	"common/eventbus"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Client     *amqp.Connection
	onceClient sync.Once
)

// 全局事件总线和桥接器
var (
	globalEventBus *eventbus.EventBus
	globalBridge   *eventbus.AMQPBridge
	once           sync.Once
)

func InitMQ() {
	onceClient.Do(func() {
		//c := conf.GetConf().RabbitMq
		// fmt.Sprintf(c.Host, c.User, c.Password, c.Host, c.Port)
		url := fmt.Sprintf("amqp://%s:%s@%s:%d/", "kcersing", "kcer-913639", "127.0.0.1", 5672)
		Client = initMQ(url)
	})
}

func initMQ(url string) *amqp.Connection {

	dialConfig := amqp.Config{
		Heartbeat: 30 * time.Second, // 设置心跳间隔为30秒
	}
	client, err := amqp.DialConfig(url, dialConfig)
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	// client.IsClosed()
	return client
}

// InitGlobalEventBus 全局初始化（应在应用启动时调用）
func InitGlobalEventBus() error {
	var err error
	once.Do(func() {
		InitMQ()
		// 创建 AMQP 客户端

		publisher, err := amqpclt.NewPublisher(Client, "events")
		if err != nil {
			klog.Errorf("failed to create publisher: %v", err)
			return
		}

		subscriber, err := amqpclt.NewSubscribe(Client, "events")
		if err != nil {
			klog.Errorf("failed to create subscriber: %v", err)
			return
		}

		// 创建事件总线和桥接器
		globalEventBus = eventbus.NewEventBus()
		globalBridge = eventbus.NewAMQPBridge(globalEventBus, publisher, subscriber)

		// 注册中间件：日志 + AMQP 发布
		globalEventBus.Use(eventbus.LoggingPlugin())
		globalEventBus.Use(globalBridge.AMQPPublishingMiddleware())

		// 启动后台监听
		ctx := context.Background()
		globalBridge.StartListener(ctx)

		klog.Infof("[MessageService] Global event bus initialized")
	})

	return err
}

// GetGlobalEventBus 获取全局事件总线
func GetGlobalEventBus() *eventbus.EventBus {
	return globalEventBus
}

// GetGlobalAMQPBridge 获取全局桥接器
func GetGlobalAMQPBridge() *eventbus.AMQPBridge {
	return globalBridge
}
