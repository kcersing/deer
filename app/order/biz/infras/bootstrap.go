package infras

import (
	"context"
	"errors"
	"order/biz/dal/mq"
	"order/biz/infras/common"
	"order/biz/infras/events"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kcersing/eventbus"
	"github.com/kcersing/eventbus/ext/amqp"
	"github.com/kcersing/eventbus/ext/amqp/amqpclt"
	"github.com/kcersing/eventbus/manager"
)

var (
	globalManager *manager.Manager
	onceManager   sync.Once
	initErr       error
)

// GetManager 获取全局的 Manager 实例
func GetManager() *manager.Manager {
	return globalManager
}

// Bootstrap 初始化并启动整个事件系统
func Bootstrap() error {
	onceManager.Do(func() { initErr = doBootstrap() })
	return initErr
}

func doBootstrap() (err error) {
	klog.Info("[Events] Initializing event system...")

	if mq.Client == nil || mq.Client.IsClosed() {
		return errors.New("events: amqp connection not initialized")
	}

	// 1. 创建 AMQP 客户端
	pub, err := amqpclt.NewPublisher(mq.Client, "eventbus")
	if err != nil {
		return err
	}
	sub, err := amqpclt.NewSubscribe(mq.Client, "eventbus")
	if err != nil {
		return err
	}

	// 2. 创建核心组件
	bus := eventbus.New()
	registry := manager.NewRegistry()
	globalManager = manager.New(bus, amqp.NewListener(bus, sub), amqp.NewPublisher(bus, pub), registry)

	// 3. 注册消费者（主题 + 处理器 + worker 数）
	registry.Register(
		string(common.Paying),
		eventbus.WrapTyped(eventbus.TypedHandler[*events.PayingOrderEvent](events.HandleOrderPaying)),
		10,
	)

	registry.Register(
		string(common.Paid),
		eventbus.WrapTyped(eventbus.TypedHandler[*events.PaidOrderEvent](events.HandleOrderPaid)),
		10,
	)

	registry.Register(
		string(common.Shipped),
		eventbus.WrapTyped(eventbus.TypedHandler[*events.ShippedOrderEvent](events.HandleOrderShipped)),
		10,
	)

	registry.Register(
		string(common.Refunded),
		eventbus.WrapTyped(eventbus.TypedHandler[*events.RefundedOrderEvent](events.HandleOrderRefunded)),
		10,
	)

	// 4. 启动所有组件
	return globalManager.Start(context.Background())
}
