package events

import (
	"context"
	"errors"
	"gen/kitex_gen/message"
	"message/biz/dal/mq"
	"sync"
	"time"

	"github.com/kcersing/eventbus"
	"github.com/kcersing/eventbus/ext/amqp"
	"github.com/kcersing/eventbus/ext/amqp/amqpclt"
	"github.com/kcersing/eventbus/ext/pool"
	"github.com/kcersing/eventbus/manager"

	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	globalManager *manager.Manager
	once          sync.Once
	initErr       error
)

// GetManager 获取全局的 Manager 实例
func GetManager() *manager.Manager {
	return globalManager
}

// Bootstrap 初始化并启动整个事件系统
func Bootstrap() error {
	once.Do(func() { initErr = doBootstrap() })
	return initErr
}

func doBootstrap() (err error) {

	klog.Info("[Events] Initializing EventManager...")
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

	// 3. 应用中间件
	// 顺序: Recover (最外层) -> Audit -> Timing -> 最终处理器
	// 仅发布链路
	bus.Use(RecoverMiddleware(), AuditLogMiddleware(), TimingMiddleware())
	klog.Info("[Events] Middlewares applied: Recover, Audit, Timing.")

	globalManager = manager.New(bus, amqp.NewListener(bus, sub), amqp.NewPublisher(bus, pub), registry)
	h := eventbus.WrapTyped(eventbus.TypedHandler[*message.SendMessagesReq](HandleSendMessages))
	registry.Register(
		EventSendMessages,
		h,
		10,
		pool.WithMaxRetries(3),
		pool.WithHandlerTimeout(10*time.Second),
		pool.WithDeadLetterFunc(func(e *eventbus.Event) {
			klog.Errorf("[Events] dead letter: topic=%s id=%s", e.Topic, e.Id)
		}),
	)
	return globalManager.Start(context.Background())

}
