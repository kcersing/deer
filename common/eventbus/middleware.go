package eventbus

import (
	"github.com/cloudwego/kitex/pkg/klog"
)

// Handler 是订阅者处理函数类型
type Handler func(event Event)

// Middleware 是中间件函数类型
// 它接收下一个要执行的处理函数（next）并返回一个新的处理函数
type Middleware func(next Handler) Handler

// LoggingPlugin 日志记录
func LoggingPlugin() Middleware {
	return func(next Handler) Handler {
		return func(event Event) {
			klog.Infof("[日志插件] 收到事件: Topic=%s \n", event.Topic)
			next(event) // 继续执行下一个中间件或最终处理函数
			klog.Infof("[日志插件] 事件处理完毕: %s \n", event.Topic)
			next(event)
		}
	}
}

// FilterPlugin 消息过滤
func FilterPlugin(filterTopic string) Middleware {
	return func(next Handler) Handler {
		return func(event Event) {
			// 只有当主题不是我们想过滤的主题时才继续
			if event.Topic == filterTopic {
				klog.Warnf("[过滤插件] 过滤掉主题为 '%s' 的事件\n", filterTopic)
				return // 中止执行链，事件不会被分发
			}
			next(event)
		}
	}
}

// TransformPlugin 消息转换
func TransformPlugin() Middleware {
	return func(next Handler) Handler {
		return func(event Event) {
			if event.Topic == "order" {
				// 假设负载是字符串，我们给它添加一个前缀
				if originalPayload, ok := event.Payload.(string); ok {
					event.Payload = "已转换: " + originalPayload
					klog.Warnf("[转换插件] 转换订单事件 Payload\n")
				}
			}
			next(event) // 传递修改后的事件
		}
	}
}
