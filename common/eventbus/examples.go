package eventbus

import (
	"common/amqpclt"
	"context"
)

// ============ 实际使用示例 ============

// 示例1：Message 服务 - 发送用户消息（本地处理）
func ExampleMessageService() {
	// 初始化内存总线
	eventBus := NewEventBus()

	// 初始化消费者注册表
	registry := NewConsumerRegistry()

	// 注册处理器：处理1 - 保存消息
	registry.RegisterHandler("save_message_handler",
		&SaveMessageHandler{})

	// 注册处理器：处理2 - 发送推送
	registry.RegisterHandler("push_notification_handler",
		&PushNotificationHandler{})

	// 注册消费者
	registry.RegisterConsumer("send_user_messages", "save_message_handler", 5)
	registry.RegisterConsumer("send_user_messages", "push_notification_handler", 10)

	// 启动所有消费者
	registry.StartAll(eventBus)

	// 发送用户消息 - 仅在本服务处理，不需要其他服务
	// PublishLocal: 高速、内存中、不跨服务

	// ❌ 错误方式：创建Publisher但没有使用
	// pub := NewEventPublisher(eventBus, nil)
	// pub.PublishDistributed(ctx, "send_user_messages", payload)  // 不必要的MQ发送

	// ✅ 正确方式：直接发布到内存总线
	ctx := context.Background()
	payload := map[string]interface{}{
		"userId":  1,
		"content": "test message",
	}

	// 由于这个事件只在本服务处理，不需要跨服务，所以用 PublishLocal
	pub := NewEventPublisher(eventBus, nil)
	pub.PublishLocal(ctx, "send_user_messages", payload)

	// 两个处理器会并发处理这个事件
	// SaveMessageHandler: 保存到数据库
	// PushNotificationHandler: 发送推送通知
}

// 示例2：Order 服务 - 订单创建（跨服务）
func ExampleOrderService(amqpPublisher *amqpclt.Publish) {
	// 初始化
	eventBus := NewEventBus()
	publisher := NewEventPublisher(eventBus, amqpPublisher)

	// 注册本服务处理器：创建订单快照
	eventBus.SubscribeWithPool("order.created",
		&CreateOrderSnapshotHandler{}, 10)

	// 启动AMQP监听（用于接收其他服务的事件）
	amqpSub := &amqpclt.Subscribe{} // 假设已初始化
	listener := NewAMQPListener(eventBus, amqpSub)
	listener.StartListener(context.Background())

	// 订单创建
	orderData := map[string]interface{}{
		"order_id": "ORD-001",
		"amount":   999.9,
		"user_id":  123,
	}

	// 发布分布式事件：
	// 1. 同步发布到内存总线 → 本服务立即处理（创建快照）
	// 2. 异步发送到RabbitMQ → 其他服务消费（库存、支付、通知）
	ctx := context.Background()
	publisher.PublishDistributed(ctx, "order.created", orderData)

	// 流程：
	// 时间T0: 本服务收到订单创建请求
	// 时间T0+: 发送到MQ (异步，不阻塞)
	// 时间T0+1: 本服务处理 CreateOrderSnapshotHandler
	// 时间T0+100: 库存服务接收到事件
	// 时间T0+150: 支付服务接收到事件
	// 时间T0+200: 通知服务接收到事件
}

// 示例3：Notification 服务 - 只从MQ消费，不主动发送事件
func ExampleNotificationService(amqpSubscriber *amqpclt.Subscribe) {
	// 初始化
	eventBus := NewEventBus()

	// 从MQ消费邮件发送请求
	listener := NewAMQPListener(eventBus, amqpSubscriber)
	listener.StartListener(context.Background())

	// 订阅来自MQ的邮件事件
	eventBus.SubscribeWithPool("notification.send_email",
		&SendEmailHandler{}, 5)

	// 订阅来自MQ的短信事件
	eventBus.SubscribeWithPool("notification.send_sms",
		&SendSmsHandler{}, 5)

	// 这个服务被动处理来自MQ的请求
	// 不主动发送事件到MQ（只消费）
}

// 示例4：Analytics 服务 - 主动触发分析任务
func ExampleAnalyticsService(amqpPublisher *amqpclt.Publish) {
	eventBus := NewEventBus()
	publisher := NewEventPublisher(eventBus, amqpPublisher)

	ctx := context.Background()

	// 场景1：只有本服务处理 - 使用 PublishLocal
	publisher.PublishLocal(ctx, "internal.cache_warmed", map[string]interface{}{})

	// 场景2：需要其他服务同步处理 - 使用 PublishDistributed
	publisher.PublishDistributed(ctx, "analytics.user_profile_updated", map[string]interface{}{
		"user_id": 123,
		"stats":   "...",
	})
	// → 本服务立即处理 + MQ转发给其他服务

	// 场景3：仅触发其他服务异步处理 - 使用 PublishToMQOnly
	publisher.PublishToMQOnly(ctx, "export.daily_report", map[string]interface{}{
		"date":   "2026-01-27",
		"format": "xlsx",
	})
	// → 仅发送到MQ，本服务不处理
}

// 示例5：使用消费者池处理高吞吐事件
func ExampleHighThroughput() {
	eventBus := NewEventBus()

	// 方式1：使用 SubscribeWithPool - 最简洁
	pool1 := eventBus.SubscribeWithPool("page_view_tracked",
		&TrackPageViewHandler{}, 50) // 50个worker

	defer pool1.Stop()

	// 方式2：使用 ConsumerRegistry - 最灵活
	registry := NewConsumerRegistry()
	registry.RegisterHandler("track_event_handler", &TrackEventHandler{})
	registry.RegisterConsumer("event.tracked", "track_event_handler", 30)
	registry.StartAll(eventBus)

	defer registry.Shutdown(context.Background())
}

// 示例6：中间件链处理
func ExampleMiddleware() {
	eventBus := NewEventBus()

	// 添加中间件
	eventBus.Use(LoggingPlugin())            // 记录所有事件
	eventBus.Use(RecoverPlugin())            // 捕获panic
	eventBus.Use(FilterPlugin("spam_topic")) // 过滤垃圾事件
	eventBus.Use(TransformPlugin())          // 转换事件数据

	// 发布时会依次经过这些中间件
	ctx := context.Background()
	pub := NewEventPublisher(eventBus, nil)
	pub.PublishLocal(ctx, "user.login", map[string]interface{}{
		"user_id": 123,
	})
}

// ============ Handler 示例 ============

type SaveMessageHandler struct{}

func (h *SaveMessageHandler) Handle(ctx context.Context, event *Event) error {
	// 保存消息到数据库
	return nil
}

type PushNotificationHandler struct{}

func (h *PushNotificationHandler) Handle(ctx context.Context, event *Event) error {
	// 发送推送通知
	return nil
}

type CreateOrderSnapshotHandler struct{}

func (h *CreateOrderSnapshotHandler) Handle(ctx context.Context, event *Event) error {
	// 创建订单快照
	return nil
}

type SendEmailHandler struct{}

func (h *SendEmailHandler) Handle(ctx context.Context, event *Event) error {
	// 发送邮件
	return nil
}

type SendSmsHandler struct{}

func (h *SendSmsHandler) Handle(ctx context.Context, event *Event) error {
	// 发送短信
	return nil
}

type TrackPageViewHandler struct{}

func (h *TrackPageViewHandler) Handle(ctx context.Context, event *Event) error {
	// 记录页面访问
	return nil
}

type TrackEventHandler struct{}

func (h *TrackEventHandler) Handle(ctx context.Context, event *Event) error {
	// 记录事件
	return nil
}
