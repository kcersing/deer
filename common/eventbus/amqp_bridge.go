package eventbus

import (
	"common/amqpclt"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// AMQPBridge 是 EventBus 与 AMQP 的双向桥接
// 功能：
// 1. 拦截内存事件总线的事件，同步发布到 RabbitMQ
// 2. 后台监听 RabbitMQ，将消息转发到内存事件总线
type AMQPBridge struct {
	eventBus   *EventBus          // 内存事件总线
	publisher  *amqpclt.Publish   // AMQP 发布者
	subscriber *amqpclt.Subscribe // AMQP 订阅者
	ctx        context.Context    // 生命周期上下文
	cancel     context.CancelFunc // 取消函数
	done       chan struct{}      // 关闭信号
}

// NewAMQPBridge 创建桥接器
func NewAMQPBridge(eventBus *EventBus, publisher *amqpclt.Publish, subscriber *amqpclt.Subscribe) *AMQPBridge {
	ctx, cancel := context.WithCancel(context.Background())
	return &AMQPBridge{
		eventBus:   eventBus,
		publisher:  publisher,
		subscriber: subscriber,
		ctx:        ctx,
		cancel:     cancel,
		done:       make(chan struct{}),
	}
}

// ============ 中间件：发布内存事件到 AMQP ============

// AMQPPublishingMiddleware 中间件：拦截内存事件，同时发布到 RabbitMQ
// 使用方式：eventBus.Use(bridge.AMQPPublishingMiddleware())
func (bridge *AMQPBridge) AMQPPublishingMiddleware() Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			// 1. 先将事件分发到内存订阅者
			if err := next.Handle(ctx, event); err != nil {
				klog.Errorf("[AMQPBridge] dispatch to memory failed: %v", err)
				return err
			}

			// 2. 异步发布到 RabbitMQ（不阻塞内存事件处理）
			go func() {
				msg := amqpclt.Message{
					Event:     event.Topic,
					Payload:   event.Payload,
					Timestamp: time.Now(),
				}
				err := bridge.publisher.Publish(ctx, event.Topic, event.Id, msg)
				if err != nil {
					klog.Errorf("[AMQPBridge] publish to AMQP failed, topic=%s, error=%v", event.Topic, err)
					// 可选：记录失败信息或触发重试机制
				} else {
					klog.Infof("[AMQPBridge] event published to AMQP, topic=%s, eventId=%s", event.Topic, event.Id)
				}
			}()

			return nil
		})
	}
}

// ============ 后台监听：从 AMQP 消费消息并转发到内存总线 ============

// StartListener 启动后台监听 RabbitMQ，将消息转发到内存事件总线
// 需在程序启动时调用，与程序生命周期同步
func (bridge *AMQPBridge) StartListener(ctx context.Context) error {
	go func() {
		defer close(bridge.done)

		msgCh, cleanup, err := bridge.subscriber.Subscribe(bridge.ctx)
		if err != nil {
			klog.Errorf("[AMQPBridge] failed to subscribe AMQP: %v", err)
			return
		}
		defer cleanup()

		klog.Infof("[AMQPBridge] AMQP listener started, waiting for messages...")

		for {
			select {
			case <-bridge.ctx.Done():
				klog.Infof("[AMQPBridge] listener shutdown")
				return

			case msg, ok := <-msgCh:
				if !ok {
					klog.Warn("[AMQPBridge] message channel closed")
					return
				}

				// 将 AMQP 消息转换为内存事件并发布
				event := &Event{
					Id:        msg.CorrelationId,
					Topic:     msg.Event,
					Payload:   msg.Payload,
					Timestamp: msg.Timestamp,
					Source:    "amqp",
					Version:   1,
				}

				// 发布到内存总线（非阻塞）
				bridge.eventBus.Publish(ctx, event)
				klog.Infof("[AMQPBridge] event forwarded to memory bus, topic=%s, eventId=%s", event.Topic, event.Id)
			}
		}
	}()

	return nil
}

// Stop 优雅关闭桥接器
func (bridge *AMQPBridge) Stop() error {
	bridge.cancel()
	<-bridge.done
	klog.Infof("[AMQPBridge] stopped")
	return nil
}

// ============ 便捷方法：批量发布事件到 AMQP ============

// PublishBatchToAMQP 批量发布内存事件到 AMQP
// 用于需要同时发送大量事件的场景（如群发消息）
func (bridge *AMQPBridge) PublishBatchToAMQP(ctx context.Context, events []*Event) ([]int, error) {
	if len(events) == 0 {
		return nil, nil
	}

	// 将内存事件转换为 AMQP 批量消息
	batchMessages := make([]amqpclt.BatchMessage, len(events))
	for i, event := range events {
		batchMessages[i] = amqpclt.BatchMessage{
			RoutingKey:    event.Topic,
			CorrelationId: event.Id,
			Message: amqpclt.Message{
				Event:     event.Topic,
				Payload:   event.Payload,
				Timestamp: event.Timestamp,
			},
		}
	}

	// 批量发布
	failedIndices, err := bridge.publisher.PublishBatch(ctx, batchMessages)
	if err != nil {
		klog.Errorf("[AMQPBridge] batch publish failed: %v", err)
		return failedIndices, err
	}

	if len(failedIndices) > 0 {
		klog.Warnf("[AMQPBridge] batch publish: %d/%d success", len(events)-len(failedIndices), len(events))
	} else {
		klog.Infof("[AMQPBridge] batch publish: all %d events published successfully", len(events))
	}

	return failedIndices, nil
}

// PublishAsyncBatchToAMQP 异步批量发布内存事件到 AMQP
// 返回结果通道，调用者可异步收集发布结果
func (bridge *AMQPBridge) PublishAsyncBatchToAMQP(ctx context.Context, events []*Event) <-chan amqpclt.PublishResult {
	resultCh := make(chan amqpclt.PublishResult, len(events))

	if len(events) == 0 {
		close(resultCh)
		return resultCh
	}

	// 将内存事件转换为 AMQP 批量消息
	batchMessages := make([]amqpclt.BatchMessage, len(events))
	for i, event := range events {
		batchMessages[i] = amqpclt.BatchMessage{
			RoutingKey:    event.Topic,
			CorrelationId: event.Id,
			Message: amqpclt.Message{
				Event:     event.Topic,
				Payload:   event.Payload,
				Timestamp: event.Timestamp,
			},
		}
	}

	// 异步发布
	bridge.publisher.PublishAsync(ctx, batchMessages, resultCh)
	return resultCh
}

// ============ 辅助方法 ============

// GetEventBus 获取内存事件总线
func (bridge *AMQPBridge) GetEventBus() *EventBus {
	return bridge.eventBus
}

// GetPublisher 获取 AMQP 发布者
func (bridge *AMQPBridge) GetPublisher() *amqpclt.Publish {
	return bridge.publisher
}

// GetSubscriber 获取 AMQP 订阅者
func (bridge *AMQPBridge) GetSubscriber() *amqpclt.Subscribe {
	return bridge.subscriber
}
