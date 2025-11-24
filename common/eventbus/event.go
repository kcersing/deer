package eventbus

import (
	"fmt"
	"sync"
)

// Event 事件
type Event struct {
	Topic   string
	Payload any // 事件负载
}

// EventChan 事件通道
type (
	EventChan chan Event
)

// EventBus 事件总线
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]EventChan
	middlewares []Middleware
}

// NewEventBus 创建事件总线
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]EventChan),
		middlewares: []Middleware{},
	}
}

// Publish 发布事件
func (eb *EventBus) Publish(topic string, event Event) {

	events := Event{Topic: topic, Payload: event}

	// 将事件分发到内存通道
	finalHandler := func(e Event) {
		eb.mu.RLock()
		subscribers := append([]EventChan{}, eb.subscribers[topic]...)
		defer eb.mu.RUnlock()
		//go func() {
		for _, subscriber := range subscribers {
			select {
			case subscriber <- event:
			default:
				fmt.Printf("警告: 主题 %s 的一个订阅者通道已满，丢弃事件。\n", topic)
			}
		}
		//}()
	}
	// 从最后一个中间件开始，向前构建执行链（责任链模式）
	chain := finalHandler
	for i := len(eb.middlewares) - 1; i >= 0; i-- {
		chain = eb.middlewares[i](chain)
	}
	// 执行整个链
	chain(events)
}

// Subscribe 订阅事件
func (eb *EventBus) Subscribe(topic string) EventChan {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	ch := make(EventChan, 100)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	fmt.Printf("已订阅主题: %s\n", topic)
	return ch
}

// Unsubscribe 取消订阅
func (eb *EventBus) Unsubscribe(topic string, ch EventChan) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	if subscribers, ok := eb.subscribers[topic]; ok {
		for i, subscriber := range subscribers {
			if ch == subscriber {
				eb.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
				close(ch)
				fmt.Printf("已取消订阅主题: %s\n", topic)

				return
			}
		}
	}

}

// Use 添加中间件
func (eb *EventBus) Use(m Middleware) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.middlewares = append(eb.middlewares, m)
}
