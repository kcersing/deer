package eventbus

import (
	"context"
	"fmt"
)

type Handler interface {
	Handle(ctx context.Context, event *Event) error
}

// EventHandler 是订阅者处理函数类型
type EventHandler func(ctx context.Context, event *Event) error

// EventHandlerFunc 定义函数类型，它适配了 Handler 接口
type EventHandlerFunc func(ctx context.Context, event *Event) error

// Handle 实现 Handler 接口
func (f EventHandlerFunc) Handle(ctx context.Context, event *Event) error {
	return f(ctx, event)
}

// TypedHandler 是一个泛型函数类型
type TypedHandler[T any] func(ctx context.Context, payload T, event Event) error

// WrapTyped 把泛型处理函数转换为标准 Handler 接口
// T: 你期望的 Payload 类型
func WrapTyped[T any](handler TypedHandler[T]) Handler {
	return EventHandlerFunc(func(ctx context.Context, event *Event) error {
		// 1 尝试断言
		typedPayload, ok := event.Payload.(T)
		if !ok {
			// 如果类型不匹配，可以选择报错或忽略
			fmt.Errorf("type mismatch: expected %T, got %T", new(T), event.Payload)
		}
		// 2. 调用具体的业务函数
		return handler(ctx, typedPayload, *event)
	})
}
