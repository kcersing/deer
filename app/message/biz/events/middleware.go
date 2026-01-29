package events

import (
	"common/eventbus"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// TimingMiddleware 记录每个事件处理的耗时
func TimingMiddleware() eventbus.Middleware {
	return func(next eventbus.Handler) eventbus.Handler {
		return eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
			start := time.Now()
			
			// 执行下一个中间件或最终的处理器，并捕获其错误
			err := next.Handle(ctx, event)
			
			klog.Infof("[Middleware-Timing] Event handled: Topic=%s, EventID=%s, Cost=%v", event.Topic, event.Id, time.Since(start))
			
			// 返回从下游收到的错误
			return err
		})
	}
}

// RecoverMiddleware 从处理器发生的 panic 中恢复，防止整个服务崩溃
func RecoverMiddleware() eventbus.Middleware {
	return func(next eventbus.Handler) eventbus.Handler {
		return eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) (err error) {
			defer func() {
				if r := recover(); r != nil {
					klog.Errorf("[Middleware-Recover] Panic recovered: Topic=%s, EventID=%s, Error: %v", event.Topic, event.Id, r)
					// 可以选择将 panic 转换为一个错误返回
					// err = fmt.Errorf("panic recovered: %v", r)
				}
			}()
			
			// 执行下一个中间件或处理器
			return next.Handle(ctx, event)
		})
	}
}
