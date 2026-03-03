package infras

import (
	"common/amqpclt"
	"common/eventbus"
	"context"

	"order/biz/dal/mq"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	// 事件主题
	EventSendOrder = "send_order"
	// 处理器名称
	handlerSendOrder = "send_order_handler"
)

var (
	globalManager *eventbus.EventManager
	onceManager   sync.Once
)

// GetManager 获取全局的 EventManager 实例
func GetManager() *eventbus.EventManager {
	return globalManager
}

// Bootstrap 初始化并启动整个事件系统
func Bootstrap() (err error) {
	onceManager.Do(func() {
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
		//bus.Use(RecoverMiddleware(), AuditLogMiddleware(), TimingMiddleware())
		klog.Info("[Events] Middlewares applied: Recover, Audit, Timing.")

		globalManager = eventbus.NewEventManager(bus, bridge, registry, eventPublisher)
		// 4. 注册所有消费者
		// 4.1. 注册处理器
		//err := globalManager.Registry.RegisterHandler(
		//	EventSendOrder,
		//	eventbus.WrapTyped(eventbus.TypedHandler[*order.SendOrderReq](HandleSendOrder)),
		//)
		if err != nil {
			klog.Errorf("Failed to register handler '%s': %v", EventSendOrder, err)

		}
		// 4.2. 注册消费者 (将主题与处理器绑定)
		err = globalManager.Registry.RegisterConsumer(EventSendOrder, handlerSendOrder, 10)
		if err != nil {
			klog.Errorf("Failed to register consumer for event '%s': %v", handlerSendOrder, err)
		}

		// 5. 启动所有组件
		if err = globalManager.Start(context.Background()); err != nil {
			return
		}

	})
	return err
}
