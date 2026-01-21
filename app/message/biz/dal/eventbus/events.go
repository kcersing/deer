package eventbus

import (
	"common/amqpclt"
	"common/eventbus"
	"context"
	"message/biz/dal/eventbus/mq"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

// 全局事件总线和桥接器
var (
	globalEventBus *eventbus.EventBus
	globalBridge   *eventbus.AMQPBridge
	once           sync.Once
)

// GetGlobalEventBus 获取全局事件总线
func GetGlobalEventBus() *eventbus.EventBus {
	return globalEventBus
}

// GetGlobalAMQPBridge 获取全局桥接器
func GetGlobalAMQPBridge() *eventbus.AMQPBridge {
	return globalBridge
}

// InitGlobalEventBus 全局初始化（应在应用启动时调用）
func InitGlobalEventBus() error {
	var err error
	once.Do(func() {
		mq.InitMQ()
		// 创建 AMQP 客户端

		publisher, err := amqpclt.NewPublisher(mq.Client, "eventbus")
		if err != nil {
			klog.Errorf("failed to create publisher: %v", err)
			return
		}

		subscriber, err := amqpclt.NewSubscribe(mq.Client, "eventbus")
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
