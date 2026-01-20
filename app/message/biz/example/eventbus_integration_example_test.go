package example

import (
	"common/amqpclt"
	"common/eventbus"
	"context"
	"fmt"
	"message/biz/dal/mq"
	"sync"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// =====================================================================
// 示例 1: 基础集成 - 内存事件 + AMQP 双向同步
// =====================================================================

func TestExampleBasicIntegration_Run(t *testing.T) {
	// 1. 初始化 MQ 连接
	mq.InitMQ()

	// 2. 创建 AMQP 发布者和订阅者
	publisher, err := amqpclt.NewPublisher(mq.Client, "events")
	if err != nil {
		klog.Fatal("cannot create publisher", err)
	}
	defer publisher.Close()

	subscriber, err := amqpclt.NewSubscribe(mq.Client, "events")
	if err != nil {
		klog.Fatal("cannot create subscriber", err)
	}

	// 3. 创建内存事件总线
	eb := eventbus.NewEventBus()

	// 4. 创建桥接器，连接内存总线和 AMQP
	bridge := eventbus.NewAMQPBridge(eb, publisher, subscriber)

	// 5. 添加中间件：拦截内存事件，同时发布到 AMQP
	eb.Use(bridge.AMQPPublishingMiddleware())

	// 6. 启动后台监听：从 AMQP 消费消息并转发到内存总线
	ctx := context.Background()
	bridge.StartListener(ctx)

	// 7. 订阅内存总线的事件（同时会从 AMQP 接收）
	ch := eb.Subscribe("user_registered")

	// 在另一个 goroutine 中处理事件
	go func() {
		for event := range ch {
			klog.Infof("[Handler] Received event: %s, Payload: %v", event.Topic, event.Payload)
		}
	}()

	// 8. 发布内存事件（自动同步到 AMQP）
	event := eventbus.NewEvent("user_registered", map[string]interface{}{
		"user_id": 123,
		"name":    "John Doe",
		"email":   "john@example.com",
	})
	eb.Publish(ctx, event)

	time.Sleep(2 * time.Second)
	bridge.Stop()
}

// =====================================================================
// 示例 2: 群发消息场景 - 批量发布到 AMQP
// =====================================================================

func TestExampleBatchMessagingWithAMQP_Run(t *testing.T) {
	// 初始化
	mq.InitMQ()
	publisher, _ := amqpclt.NewPublisher(mq.Client, "messages")
	defer publisher.Close()
	subscriber, _ := amqpclt.NewSubscribe(mq.Client, "messages")
	eb := eventbus.NewEventBus()
	bridge := eventbus.NewAMQPBridge(eb, publisher, subscriber)

	ctx := context.Background()

	// 模拟接收者列表
	recipients := []string{"user_1", "user_2", "user_3", "user_4", "user_5"}

	// 为每个接收者创建一个事件
	events := make([]*eventbus.Event, len(recipients))
	for i, recipientID := range recipients {
		events[i] = &eventbus.Event{
			Id:    fmt.Sprintf("msg-%d", i),
			Topic: "member_message_send",
			Payload: map[string]interface{}{
				"recipient_id": recipientID,
				"title":        "群发消息",
				"content":      "这是一条群发消息",
				"sent_at":      time.Now().Unix(),
			},
			Timestamp: time.Now(),
			Source:    "service",
		}
	}

	// 批量发布到 AMQP
	failedIndices, err := bridge.PublishBatchToAMQP(ctx, events)
	if err != nil {
		klog.Errorf("batch publish failed: %v", err)
	}

	if len(failedIndices) == 0 {
		klog.Infof("all %d messages published successfully", len(events))
	} else {
		klog.Warnf("%d messages failed: %v", len(failedIndices), failedIndices)
	}
}

// =====================================================================
// 示例 3: 异步群发 - 监听发布结果
// =====================================================================

func TestExampleAsyncBatchMessaging_Run(t *testing.T) {
	// 初始化
	mq.InitMQ()
	publisher, _ := amqpclt.NewPublisher(mq.Client, "messages")
	defer publisher.Close()
	subscriber, _ := amqpclt.NewSubscribe(mq.Client, "messages")
	eb := eventbus.NewEventBus()
	bridge := eventbus.NewAMQPBridge(eb, publisher, subscriber)

	ctx := context.Background()

	// 创建多个事件
	var events []*eventbus.Event
	for i := 0; i < 10; i++ {
		events = append(events, &eventbus.Event{
			Id:    fmt.Sprintf("msg-%d", i),
			Topic: "order_created",
			Payload: map[string]interface{}{
				"order_id": 1000 + i,
				"total":    99.99 + float64(i),
			},
			Timestamp: time.Now(),
		})
	}

	// 异步发布，获取结果通道
	resultCh := bridge.PublishAsyncBatchToAMQP(ctx, events)

	// 收集结果
	successCount := 0
	failedCount := 0

	for result := range resultCh {
		if result.Error != nil {
			failedCount++
			klog.Errorf("message %d publish failed: %v", result.Index, result.Error)
		} else {
			successCount++
			klog.Infof("message %d published successfully", result.Index)
		}
	}

	klog.Infof("publish result: %d success, %d failed", successCount, failedCount)
}

// =====================================================================
// 示例 4: 订阅处理 - 从 AMQP 接收并处理
// =====================================================================

func TestExampleSubscribeAndHandle_Run(t *testing.T) {
	// 初始化
	mq.InitMQ()
	publisher, _ := amqpclt.NewPublisher(mq.Client, "events")
	subscriber, _ := amqpclt.NewSubscribe(mq.Client, "events")
	eb := eventbus.NewEventBus()
	bridge := eventbus.NewAMQPBridge(eb, publisher, subscriber)

	ctx := context.Background()

	// 启动后台监听
	bridge.StartListener(ctx)

	// 定义事件处理器
	userRegisteredHandler := eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
		klog.Infof("[UserRegisteredHandler] Processing user registration: %v", event.Payload)
		// 业务逻辑：发送欢迎邮件、初始化用户数据等
		return nil
	})

	// 异步订阅（使用并发处理）
	eb.SubscribeAsync("user_registered", userRegisteredHandler, 3)

	// 发布事件到内存总线（会自动同步到 AMQP）
	event := eventbus.NewEvent("user_registered", map[string]interface{}{
		"user_id": 123,
		"name":    "Alice",
	})
	eb.Publish(ctx, event)

	time.Sleep(3 * time.Second)
	bridge.Stop()
}

// =====================================================================
// 示例 5: 整合到实际服务 - 群发消息
// =====================================================================

type SendMemberMessagesService struct {
	ctx    context.Context
	bridge *eventbus.AMQPBridge
}

func (s *SendMemberMessagesService) SendMessages(recipients []string, title, content string) error {
	ctx := context.Background()

	// 为每个接收者创建消息事件
	events := make([]*eventbus.Event, len(recipients))
	for i, recipientID := range recipients {
		events[i] = &eventbus.Event{
			Id:    fmt.Sprintf("broadcast-%s-%d", recipientID, time.Now().Unix()),
			Topic: "member_message_broadcast",
			Payload: map[string]interface{}{
				"recipient_id": recipientID,
				"title":        title,
				"content":      content,
				"sent_at":      time.Now().Unix(),
			},
			Timestamp: time.Now(),
		}
	}

	// 异步批量发布
	resultCh := s.bridge.PublishAsyncBatchToAMQP(ctx, events)

	// 异步等待结果
	go func() {
		successCount := 0
		failedCount := 0

		for result := range resultCh {
			if result.Error != nil {
				failedCount++
			} else {
				successCount++
			}
		}

		klog.Infof("[SendMemberMessagesService] Broadcast complete: %d success, %d failed",
			successCount, failedCount)
	}()

	return nil
}

// =====================================================================
// 示例 6: 中间件链 - 多个中间件组合
// =====================================================================

func TestExampleMiddlewareChain_Run(t *testing.T) {
	mq.InitMQ()
	publisher, _ := amqpclt.NewPublisher(mq.Client, "events")
	subscriber, _ := amqpclt.NewSubscribe(mq.Client, "events")
	eb := eventbus.NewEventBus()
	bridge := eventbus.NewAMQPBridge(eb, publisher, subscriber)

	ctx := context.Background()

	// 添加多个中间件
	eb.Use(eventbus.LoggingPlugin())          // 1. 日志中间件
	eb.Use(bridge.AMQPPublishingMiddleware()) // 2. AMQP 发布中间件
	eb.Use(eventbus.TransformPlugin())        // 3. 消息转换中间件

	bridge.StartListener(ctx)

	// 订阅处理
	eb.SubscribeAsync("order_created", eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
		klog.Infof("[OrderHandler] Processing order: %v", event.Payload)
		return nil
	}), 2)

	// 发布事件（会依次经过所有中间件）
	event := eventbus.NewEvent("order_created", map[string]interface{}{
		"order_id": 12345,
		"total":    299.99,
	})
	eb.Publish(ctx, event)

	time.Sleep(2 * time.Second)
	bridge.Stop()
}

// =====================================================================
// 示例 7: 全局桥接器初始化（推荐用于生产环境）
// =====================================================================

var (
	globalEventBus *eventbus.EventBus
	globalBridge   *eventbus.AMQPBridge
	once           sync.Once
)

// InitGlobalEventBus 初始化全局事件总线和桥接器
func InitGlobalEventBus() error {
	var err error
	once.Do(func() {
		// 初始化 MQ
		mq.InitMQ()

		// 创建 AMQP 客户端
		publisher, err := amqpclt.NewPublisher(mq.Client, "deer")
		if err != nil {
			klog.Errorf("failed to create publisher: %v", err)
			return
		}

		subscriber, err := amqpclt.NewSubscribe(mq.Client, "deer")
		if err != nil {
			klog.Errorf("failed to create subscriber: %v", err)
			return
		}

		// 创建事件总线和桥接器
		globalEventBus = eventbus.NewEventBus()
		globalBridge = eventbus.NewAMQPBridge(globalEventBus, publisher, subscriber)

		// 注册中间件
		globalEventBus.Use(eventbus.LoggingPlugin())
		globalEventBus.Use(globalBridge.AMQPPublishingMiddleware())

		// 启动后台监听
		ctx := context.Background()
		globalBridge.StartListener(ctx)

		klog.Infof("[EventBus] Global event bus initialized successfully")
	})

	return err
}

// GetEventBus 获取全局事件总线
func GetEventBus() *eventbus.EventBus {
	return globalEventBus
}

// GetAMQPBridge 获取全局 AMQP 桥接器
func GetAMQPBridge() *eventbus.AMQPBridge {
	return globalBridge
}

// =====================================================================
// 示例使用方式
// =====================================================================

func TestEvenbus_Run(t *testing.T) {
	// 初始化全局事件总线
	if err := InitGlobalEventBus(); err != nil {
		klog.Fatal("failed to initialize event bus", err)
	}

	eb := GetEventBus()
	bridge := GetAMQPBridge()

	ctx := context.Background()

	// 示例：订阅事件
	eb.SubscribeAsync("order_created", eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
		klog.Infof("Order created event received: %v", event.Payload)
		return nil
	}), 2)

	// 示例：发布事件
	event := eventbus.NewEvent("order_created", map[string]interface{}{
		"order_id": 12345,
		"total":    299.99,
	})
	eb.Publish(ctx, event)

	// 示例：批量消息
	recipients := []string{"user_1", "user_2", "user_3"}
	events := make([]*eventbus.Event, len(recipients))
	for i, recipient := range recipients {
		events[i] = &eventbus.Event{
			Id:    fmt.Sprintf("msg-%d", i),
			Topic: "member_message",
			Payload: map[string]interface{}{
				"recipient": recipient,
				"message":   "Hello from event bus",
			},
			Timestamp: time.Now(),
		}
	}

	resultCh := bridge.PublishAsyncBatchToAMQP(ctx, events)
	for result := range resultCh {
		if result.Error != nil {
			klog.Errorf("Publish failed: %v", result.Error)
		}
	}

	time.Sleep(5 * time.Second)
	bridge.Stop()
}
