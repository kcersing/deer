package eventbus

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

// EventChan 事件通道
type (
	EventChan chan *Event
)

var _ Bus = (*EventBus)(nil)
var _ Publish = (*EventBus)(nil)
var _ Subscribe = (*EventBus)(nil)

// EventBus 事件总线
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]EventChan // 订阅者映射：topic → 多个通道
	middlewares []Middleware           // 中间件件
	chain       Handler                // 缓存的中间件
}

// NewEventBus 创建事件总线
func NewEventBus() *EventBus {
	eb := &EventBus{
		subscribers: make(map[string][]EventChan),
		middlewares: []Middleware{},
	}
	eb.rebuildChain() // 初始化链
	return eb
}

// 将事件分发到内存通道
func (eb *EventBus) dispatch(ctx context.Context, event *Event) error {
	eb.mu.RLock()
	subscribers := append([]EventChan{}, eb.subscribers[event.Topic]...)
	defer eb.mu.RUnlock()
	for _, subscriber := range subscribers {
		select {
		case subscriber <- event: //发送事件
		default:
			klog.Infof("警告: 主题 %s 的一个订阅者通道已满，丢弃事件。", event.Topic)
		}
	}
	return nil
}
func (eb *EventBus) rebuildChain() {
	var h Handler = EventHandlerFunc(eb.dispatch)
	for i := len(eb.middlewares) - 1; i >= 0; i-- {
		h = eb.middlewares[i](h)
	}
	eb.chain = h
}

// Use 添加中间件
func (eb *EventBus) Use(mw ...Middleware) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.middlewares = append(eb.middlewares, mw...)
	eb.rebuildChain()
}

// Publish 发布事件到内存总线
func (eb *EventBus) Publish(ctx context.Context, event *Event) {

	if err := eb.chain.Handle(ctx, event); err != nil {
		klog.Infof("[Error] Handle event failed: %v\n", err)
	}
}

// PublishByTopic 按主题发布事件（简化版）
func (eb *EventBus) PublishByTopic(ctx context.Context, topic string, payload any) {
	event := NewEvent(topic, payload)
	eb.Publish(ctx, event)
}

// Subscribe 订阅事件
/**
* 创建通道 → 加入订阅者列表 → 返回通道供外部使用
* 使用示例：
* ch := eventBus.Subscribe("user_registered")
* for event := range ch {
*     fmt.Println(event)
* }
**/
func (eb *EventBus) Subscribe(topic string) EventChan {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	ch := make(EventChan, 100)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	klog.Infof("[Subscribe] 已订阅主题: %s\n", topic)
	return ch
}

func (eb *EventBus) safeHandle(topic string, handler Handler, event *Event) {

	defer func() {
		if r := recover(); r != nil {
			klog.Infof("[Panic Recover] Topic: %s, Error: %v\n", topic, r)
		}
	}()
	if err := handler.Handle(context.Background(), event); err != nil {
		// 可以在这里做统一的错误日志
		klog.Infof("[Error] Handle event failed: %v\n", err)
	}
}

// SubscribeAsync 异步订阅：用户只需提供 handler，无需自己管理协程
func (eb *EventBus) SubscribeAsync(topic string, handler Handler, concurrency int) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	ch := make(EventChan, 100)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)

	for i := 0; i < concurrency; i++ {
		go func() {
			for event := range ch {
				eb.safeHandle(topic, handler, event)
			}
		}()
	}

}

// Unsubscribe 取消订阅
func (eb *EventBus) Unsubscribe(topic string, ch EventChan) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if subscribers, ok := eb.subscribers[topic]; ok {
		for i, subscriber := range subscribers {
			if ch == subscriber {
				eb.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
				//close(ch)
				klog.Infof("已取消订阅主题: %s\n", topic)
				return
			}
		}
	}

}

// ============ 使用消费者池订阅 ============

// SubscribeWithPool 使用消费者池订阅事件
func (eb *EventBus) SubscribeWithPool(topic string, handler Handler, workerNum int32) *ConsumerPool {
	pool := NewConsumerPool(topic, handler, workerNum)
	pool.Start()

	// 创建通道并启动转发 goroutine
	ch := eb.Subscribe(topic)
	go func() {
		defer func() {
			close(ch)
			pool.Stop()
		}()
		for event := range ch {
			pool.Consume(event)
		}
	}()

	return pool
}
func (eb *EventBus) Close() error {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	// 关闭所有订阅通道
	for _, chs := range eb.subscribers {
		for _, ch := range chs {
			close(ch)
		}
	}
	eb.subscribers = make(map[string][]EventChan)
	return nil
}
