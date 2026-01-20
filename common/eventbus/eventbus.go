package eventbus

import (
	"context"
	"fmt"
	"sync"
)

type EventBusInterface interface {

	//SubscribeOnce(eventType string, handler Handler) error
	Subscribe(topic string) <-chan Event
	Unsubscribe(topic string, ch <-chan Event)
	Publish(ctx context.Context, topic string, payload any)

	Close() error
}

// EventChan 事件通道
type (
	EventChan chan *Event
)

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
			fmt.Printf("警告: 主题 %s 的一个订阅者通道已满，丢弃事件。", event.Topic)
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
func (eb *EventBus) Use(m Middleware) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.middlewares = append(eb.middlewares, m)
	eb.rebuildChain()
}

// Publish 发布事件
func (eb *EventBus) Publish(ctx context.Context, event *Event) {
	eb.chain.Handle(ctx, event)
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
	fmt.Printf("已订阅主题: %s\n", topic)
	return ch
}

func (eb *EventBus) safeHandle(topic string, handler Handler, event *Event) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[Panic Recover] Topic: %s, Error: %v\n", topic, r)
		}
	}()
	if err := handler.Handle(context.Background(), event); err != nil {
		// 可以在这里做统一的错误日志
		fmt.Printf("[Error] Handle event failed: %v\n", err)
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
				fmt.Printf("已取消订阅主题: %s\n", topic)
				return
			}
		}
	}

}
