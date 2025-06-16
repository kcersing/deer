package domain

import (
	"reflect"
	"sync"
)

type EventBus interface {
	// Publish 发布事件
	Publish(evens ...interface{}) error
	Subscribe(handler EventHandler, events ...interface{})
}

// EventHandler 事件处理器接口
type EventHandler interface {
	Handle(event interface{}) error
}
type MemoryEventBus struct {
	handlers map[reflect.Type][]EventHandler
	sync.RWMutex
}

func NewMemoryEventBus() *MemoryEventBus {
	return &MemoryEventBus{
		handlers: make(map[reflect.Type][]EventHandler),
	}
}

// Publish 发布事件
func (b *MemoryEventBus) Publish(events ...interface{}) error {
	for _, event := range events {
		eventType := reflect.TypeOf(event)
		b.RLock()
		handlers := b.handlers[eventType]
		b.RUnlock()
		for _, handler := range handlers {
			if err := handler.Handle(event); err != nil {
				return err
			}
		}
	}
	return nil
}

// Subscribe 订阅事件
func (b *MemoryEventBus) Subscribe(handler EventHandler, events ...interface{}) {
	b.Lock()
	defer b.Unlock()
	for _, event := range events {
		eventType := reflect.TypeOf(event)
		b.handlers[eventType] = append(b.handlers[eventType], handler)
	}
}
