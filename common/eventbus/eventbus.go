package eventbus

import (
	"context"
	"fmt"
	"sync"
)

type EventBusInterface interface {
	Subscribe(eventType string, handler Handler) error

	//SubscribeOnce(eventType string, handler Handler) error

	Unsubscribe(eventType string, handler Handler) error

	Publish(ctx context.Context, event *Event) error

	Close() error
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
	chain       Handler // 缓存构建好的中间件链
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
func (eb *EventBus) dispatch(ctx context.Context, event *Event) {
	eb.mu.RLock()
	subscribers := append([]EventChan{}, eb.subscribers[event.Topic]...)
	defer eb.mu.RUnlock()
	for _, subscriber := range subscribers {
		select {
		case subscriber <- *event:
		default:
			fmt.Printf("警告: 主题 %s 的一个订阅者通道已满，丢弃事件。", event.Topic)
		}
	}
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
				//close(ch)
				fmt.Printf("已取消订阅主题: %s\n", topic)
				return
			}
		}
	}

}
