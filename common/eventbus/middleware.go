package eventbus

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// Middleware 中间件类型定义
type Middleware func(next Handler) Handler

// LoggingPlugin 日志记录
func LoggingPlugin() Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			start := time.Now()
			next.Handle(ctx, event) // 继续执行下一个中间件或最终处理函数
			klog.Infof("[日志插件] Done: Topic=%s cost=%v", event.Topic, time.Since(start))
			return nil
		})
	}
}

// FilterPlugin 消息过滤
func FilterPlugin(filterTopic string) Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			// 只有当主题不是我们想过滤的主题时才继续
			if event.Topic == filterTopic {
				klog.Warnf("[过滤插件] 过滤掉主题为 '%s' 的事件\n", filterTopic)
				return nil // 中止执行链，事件不会被分发
			}
			return next.Handle(ctx, event)
		})
	}
}

// TransformPlugin 消息转换
func TransformPlugin() Middleware {
	return func(next Handler) Handler {
		return EventHandlerFunc(func(ctx context.Context, event *Event) error {
			if event.Topic == "order" {
				// 假设负载是字符串，我们给它添加一个前缀
				if originalPayload, ok := event.Payload.(string); ok {
					event.Payload = "已转换: " + originalPayload
					klog.Warnf("[转换插件] 转换订单事件 Payload\n")
				}
			}
			return next.Handle(ctx, event) // 传递修改后的事件

		})
	}
}
