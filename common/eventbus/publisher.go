package eventbus

import (
	"common/amqpclt"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 事件发布管理器 - 明确职责分离 ============

// EventPublisher 事件发布管理器 - 负责统一管理事件发布到不同目标
type EventPublisher struct {
	memoryBus *EventBus        // 内存事件总线
	amqpPub   *amqpclt.Publish // AMQP 发布者（可选）
}

// NewEventPublisher 创建事件发布管理器
func NewEventPublisher(memoryBus *EventBus, amqpPub *amqpclt.Publish) *EventPublisher {
	return &EventPublisher{
		memoryBus: memoryBus,
		amqpPub:   amqpPub,
	}
}

// ============ 方式1: 纯内存事件 - 只在内存总线中处理 ============

// PublishLocal 发布本地内存事件
// 用于：单个服务内部的异步处理
// 特点：高速、低延迟、不持久化、不跨服务
// 示例：SendUserMessages 事件在当前服务处理
func (pub *EventPublisher) PublishLocal(ctx context.Context, topic string, payload any) {
	event := NewEvent(topic, payload)
	event.Source = "local" // 标记为本地事件
	pub.memoryBus.Publish(ctx, event)
	klog.Infof("[PublishLocal] Event published to memory bus only, topic=%s, eventId=%s", topic, event.Id)
}

// ============ 方式2: 分布式事件 - 同时发送到MQ和内存 ============

// PublishDistributed 发布分布式事件
// 用于：跨服务通信、需要持久化的事件
// 特点：1. 异步发送到RabbitMQ（持久化）
//  2. 立即发布到内存总线（本服务处理）
//  3. 其他服务从MQ消费
//
// 流程：
//   - 本服务发送 → RabbitMQ （用于其他服务消费）
//   - 本服务发送 → 内存总线 （本服务立即处理）
//
// 示例：OrderCreated 事件需要通知其他服务
func (pub *EventPublisher) PublishDistributed(ctx context.Context, topic string, payload any) error {
	if pub.amqpPub == nil {
		klog.Warnf("[PublishDistributed] AMQP publisher not configured, falling back to local only")
		pub.PublishLocal(ctx, topic, payload)
		return nil
	}

	event := NewEvent(topic, payload)
	event.Source = "service"

	// 1. 异步发送到MQ（不阻塞）
	go func() {
		msg := amqpclt.Message{
			Event:     topic,
			Payload:   payload,
			Timestamp: time.Now(),
		}
		err := pub.amqpPub.Publish(ctx, topic, event.Id, msg)
		if err != nil {
			klog.Errorf("[PublishDistributed] publish to MQ failed, topic=%s, error=%v", topic, err)
			// TODO: 可以触发重试或死信队列机制
		} else {
			klog.Infof("[PublishDistributed] event published to MQ, topic=%s, eventId=%s", topic, event.Id)
		}
	}()

	// 2. 同时发送到内存总线（本服务立即处理）
	pub.memoryBus.Publish(ctx, event)
	klog.Infof("[PublishDistributed] event published to memory bus, topic=%s, eventId=%s", topic, event.Id)

	return nil
}

// ============ 方式3: 只发送到MQ - 由其他服务消费 ============

// PublishToMQOnly 只发送事件到MQ，不在本服务处理
// 用于：触发其他服务处理、数据同步等
// 特点：1. 只发送到MQ，本服务不处理
//  2. 其他服务从MQ消费处理
//
// 示例：发送通知给notification服务
func (pub *EventPublisher) PublishToMQOnly(ctx context.Context, topic string, payload any) error {
	if pub.amqpPub == nil {
		return klog.Errorf("[PublishToMQOnly] AMQP publisher not configured, cannot publish to MQ")
	}

	event := NewEvent(topic, payload)
	event.Source = "service"

	msg := amqpclt.Message{
		Event:     topic,
		Payload:   payload,
		Timestamp: time.Now(),
	}

	err := pub.amqpPub.Publish(ctx, topic, event.Id, msg)
	if err != nil {
		klog.Errorf("[PublishToMQOnly] publish to MQ failed, topic=%s, error=%v", topic, err)
		return err
	}

	klog.Infof("[PublishToMQOnly] event published to MQ only, topic=%s, eventId=%s", topic, event.Id)
	return nil
}

// ============ 方便的简写方法 ============

// PublishLocal 短方法：发布到本地内存
func (pub *EventPublisher) Local(ctx context.Context, topic string, payload any) {
	pub.PublishLocal(ctx, topic, payload)
}

// Distributed 短方法：发布到分布式
func (pub *EventPublisher) Distributed(ctx context.Context, topic string, payload any) error {
	return pub.PublishDistributed(ctx, topic, payload)
}

// MQOnly 短方法：只发到MQ
func (pub *EventPublisher) MQOnly(ctx context.Context, topic string, payload any) error {
	return pub.PublishToMQOnly(ctx, topic, payload)
}
