package events

import (
	"common/amqpclt"
	"common/eventbus"
	"context"
	"message/biz/dal/mq"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

// EventManager 统一管理事件总线、桥接器和消费者注册表
type EventManager struct {
	Bus       *eventbus.EventBus
	Bridge    *eventbus.AMQPListener
	Registry  *eventbus.ConsumerRegistry
	Publisher *eventbus.EventPublisher
}

var (
	globalManager *EventManager
	once          sync.Once
)

// GetManager 获取全局的 EventManager 实例
func GetManager() *EventManager {
	return globalManager
}

// Bootstrap 初始化并启动整个事件系统
func Bootstrap() (err error) {
	once.Do(func() {
		klog.Info("[Events] Initializing EventManager...")

		// 1. 创建 AMQP 客户端
		publisher, e := amqpclt.NewPublisher(mq.Client, "eventbus")
		if e != nil {
			err = e
			return
		}
		subscriber, e := amqpclt.NewSubscribe(mq.Client, "eventbus")
		if e != nil {
			err = e
			return
		}

		// 2. 创建核心组件
		bus := eventbus.NewEventBus()
		bridge := eventbus.NewAMQPListener(bus, subscriber)
		registry := eventbus.NewConsumerRegistry()
		eventPublisher := eventbus.NewEventPublisher(bus, publisher)

		// 3. 应用中间件
		// 顺序: Recover (最外层) -> Audit -> Timing -> 最终处理器
		bus.Use(RecoverMiddleware(), AuditLogMiddleware(), TimingMiddleware())
		klog.Info("[Events] Middlewares applied: Recover, Audit, Timing.")

		globalManager = &EventManager{
			Bus:       bus,
			Bridge:   bridge,
			Registry:  registry,
			Publisher: eventPublisher,
		}

		// 4. 注册所有消费者
		if err = InitMessageConsumers(); err != nil {
			return
		}

		// 5. 启动所有组件
		if err = bridge.StartListener(context.Background()); err != nil {
			return
		}
		if err = registry.StartAll(bus); err != nil {
			return
		}

		klog.Info("[Events] EventManager started successfully.")
	})
	return err
}

// Shutdown 优雅地关闭整个事件系统
func Shutdown(ctx context.Context) error {
	if globalManager == nil {
		klog.Warn("[Events] EventManager not initialized, skipping shutdown.")
		return nil
	}

	klog.Info("[Events] Shutting down EventManager...")

	if err := globalManager.Bridge.Stop(); err != nil {
		klog.Errorf("[Events] Failed to stop AMQP listener: %v", err)
	}

	if err := globalManager.Registry.Shutdown(ctx); err != nil {
		klog.Errorf("[Events] Failed to shutdown consumer registry: %v", err)
	}

	klog.Info("[Events] EventManager shut down complete.")
	return nil
}
