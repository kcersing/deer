package events

import (
	"common/amqpclt"
	"common/eventbus"
	"context"
	"message/biz/dal/mq"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

// 全局事件总线和桥接器
var (
	globalEventBus   *eventbus.EventBus
	globalBridge     *eventbus.AMQPBridge
	consumerRegistry *eventbus.ConsumerRegistry
	once             sync.Once
	stopChan         chan struct{}
)

// GetGlobalEventBus 获取全局事件总线
func GetGlobalEventBus() *eventbus.EventBus { return globalEventBus }

// GetGlobalAMQPBridge 获取全局桥接器
func GetGlobalAMQPBridge() *eventbus.AMQPBridge { return globalBridge }

// GetConsumerRegistry 获取消费者注册表
func GetConsumerRegistry() *eventbus.ConsumerRegistry { return consumerRegistry }

// InitGlobalEventBus 全局初始化（应在应用启动时调用）
func InitGlobalEventBus() error {
	var err error
	once.Do(func() {

		// 创建 AMQP 客户端
		publisher, e := amqpclt.NewPublisher(mq.Client, "eventbus")
		if e != nil {
			klog.Errorf("[InitGlobalEventBus] failed to create publisher: %v", e)
			return
		}

		subscriber, e := amqpclt.NewSubscribe(mq.Client, "eventbus")
		if e != nil {
			klog.Errorf("[InitGlobalEventBus] failed to create subscriber: %v", e)
			return
		}

		// 创建事件总线和桥接器
		globalEventBus = eventbus.NewEventBus()
		globalBridge = eventbus.NewAMQPBridge(globalEventBus, publisher, subscriber)

		// 注册中间件：恢复、日志、AMQP 发布
		globalEventBus.Use(eventbus.RecoverPlugin())
		globalEventBus.Use(eventbus.LoggingPlugin())

		// 启动后台监听
		ctx := context.Background()
		globalBridge.StartListener(ctx)

		klog.Infof("[InitGlobalEventBus] Global event bus initialized")
	})
	return err
}

// Bootstrap 初始化并启动事件系统，返回 cleanup 函数以便优雅关闭
func Bootstrap() error {

	var err error

	if err := InitGlobalEventBus(); err != nil {
		return err
	}
	if consumerRegistry == nil {
		consumerRegistry = eventbus.NewConsumerRegistry()
	}

	if err := InitMessageConsumers(); err != nil {
		klog.Infof("[Bootstrap] Err", err)
		return err
	}

	if err := StartMessageConsumers(); err != nil {
		return err
	}

	stopChan = make(chan struct{})

	return err

}

func Shutdown(ctx context.Context) error {

	if stopChan != nil {
		close(stopChan)
	}
	if consumerRegistry != nil {
		if err := consumerRegistry.Shutdown(ctx); err != nil {
			return err
		}
	}
	if globalBridge != nil {
		globalBridge.Stop()
	}

	klog.Infof("[Shutdown] Shutdown completed")
	return nil
}
