package eventbus

import (
	"common/amqpclt"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// PublishScope 定义了事件的发布范围
type PublishScope int

const (
	// ScopeLocal 只发布到本地内存总线
	ScopeLocal PublishScope = 1
	// ScopeDistributed 发布到MQ和本地内存总线
	ScopeDistributed PublishScope = 2
	// ScopeMQOnly 只发布到MQ
	ScopeMQOnly PublishScope = 3
)

// PublishOptions 定义了发布的选项
type PublishOptions struct {
	Scope PublishScope
}

// WithScope 创建一个指定发布范围的选项
func WithScope(scope PublishScope) func(*PublishOptions) {
	return func(o *PublishOptions) {
		o.Scope = scope
	}
}

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

// Publish 是统一的事件发布方法
func (pub *EventPublisher) Publish(ctx context.Context, topic string, payload any, opts ...func(*PublishOptions)) error {
	options := &PublishOptions{
		Scope: ScopeLocal, // 默认为本地发布
	}
	for _, opt := range opts {
		opt(options)
	}

	event := NewEvent(topic, payload)

	var err error
	switch options.Scope {
	case ScopeLocal:
		event.Source = "local"
		pub.memoryBus.Publish(ctx, event)
		klog.Infof("[Publish] Event published to memory bus only, topic=%s, eventId=%s", topic, event.Id)

	case ScopeDistributed:
		if pub.amqpPub == nil {
			klog.Warnf("[Publish] AMQP publisher not configured for distributed scope, falling back to local only")
			return pub.Publish(ctx, topic, payload, WithScope(ScopeLocal))
		}
		event.Source = "service"
		// 1. 异步发送到MQ
		go pub.publishToMQ(ctx, event)
		// 2. 发送到内存总线
		pub.memoryBus.Publish(ctx, event)
		klog.Infof("[Publish] Event published to memory bus and MQ, topic=%s, eventId=%s", topic, event.Id)

	case ScopeMQOnly:
		if pub.amqpPub == nil {

			return fmt.Errorf("[Publish] AMQP publisher not configured, cannot publish to MQ only")
		}
		event.Source = "service"
		err = pub.publishToMQ(ctx, event)

	default:
		err = fmt.Errorf("unknown publish scope: %v", options.Scope)
	}

	return err
}

func (pub *EventPublisher) publishToMQ(ctx context.Context, event *Event) error {
	msg := amqpclt.Message{
		Event:     event.Topic,
		Payload:   event.Payload,
		Timestamp: time.Now(),
	}
	err := pub.amqpPub.Publish(ctx, event.Topic, event.Id, msg)
	if err != nil {
		klog.Errorf("[Publish] publish to MQ failed, topic=%s, error=%v", event.Topic, err)
	} else {
		klog.Infof("[Publish] event published to MQ, topic=%s, eventId=%s", event.Topic, event.Id)
	}
	return err
}

// ============ 方便的简写方法 ============

// Local 短方法：发布到本地内存
func (pub *EventPublisher) Local(ctx context.Context, topic string, payload any) {
	_ = pub.Publish(ctx, topic, payload, WithScope(ScopeLocal))
}

// Distributed 短方法：发布到分布式
func (pub *EventPublisher) Distributed(ctx context.Context, topic string, payload any) error {
	return pub.Publish(ctx, topic, payload, WithScope(ScopeDistributed))
}

// MQOnly 短方法：只发到MQ
func (pub *EventPublisher) MQOnly(ctx context.Context, topic string, payload any) error {
	return pub.Publish(ctx, topic, payload, WithScope(ScopeMQOnly))
}
