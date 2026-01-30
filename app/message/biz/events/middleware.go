package events

import (
	"common/eventbus"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// AuditLogMiddleware 负责创建审计日志条目，并在事件处理后将其持久化
func AuditLogMiddleware() eventbus.Middleware {
	return func(next eventbus.Handler) eventbus.Handler {
		return eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
			// 1. 创建审计条目并注入 context
			entry := NewAuditEntry()
			ctxWithAudit := WithAuditEntry(ctx, entry)

			// 2. 执行下一个中间件或最终的处理器
			err := next.Handle(ctxWithAudit, event)

			// 3. 处理器执行完毕，填充最终状态并持久化
			if err != nil {
				entry.Status = "Failure"
				entry.Error = err.Error()
			} else {
				entry.Status = "Success"
			}

			// 在这里执行持久化操作
			// TODO: 将 entry 保存到数据库或专门的审计日志服务
			klog.Infof("[Middleware-审计] Audit Log: EventID=%s, Status=%s, Details=%v", event.Id, entry.Status, entry.Details)
			//新创建事件执行入库
			return err
		})
	}
}

// TimingMiddleware 记录每个事件处理的耗时
func TimingMiddleware() eventbus.Middleware {
	return func(next eventbus.Handler) eventbus.Handler {
		return eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) error {
			start := time.Now()
			err := next.Handle(ctx, event)
			klog.Infof("[Middleware-耗时] Event handled: Topic=%s, EventID=%s, Cost=%v", event.Topic, event.Id, time.Since(start))
			return err
		})
	}
}

// RecoverMiddleware 从处理器发生的 panic 中恢复
func RecoverMiddleware() eventbus.Middleware {
	return func(next eventbus.Handler) eventbus.Handler {
		return eventbus.EventHandlerFunc(func(ctx context.Context, event *eventbus.Event) (err error) {
			defer func() {
				if r := recover(); r != nil {
					klog.Errorf("[Middleware-恢复] Panic recovered: Topic=%s, EventID=%s, Error: %v", event.Topic, event.Id, r)
				}
			}()
			return next.Handle(ctx, event)
		})
	}
}
