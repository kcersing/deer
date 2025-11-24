package eventbus

import (
	"fmt"
	"testing"
	"time"
)

func TestEventBusMiddleware_Run(t *testing.T) {

	eb := NewEventBus()

	//注册插件（顺序很重要）
	eb.Use(LoggingPlugin())     //日志（最外层）
	eb.Use(FilterPlugin("pay")) // 过滤（在日志和转换之间）
	eb.Use(TransformPlugin())   // 转换（最内层，靠近分发）

	//订阅设置
	ordersCh := make(chan Event, 10)
	payCh := make(chan Event, 10)
	eb.subscribers["order"] = append(eb.subscribers["order"], ordersCh)
	eb.subscribers["pay"] = append(eb.subscribers["pay"], payCh)

	go func() {
		for event := range ordersCh {
			fmt.Printf("[订阅者] 收到最终订单事件: %s \n", event.Payload)
		}
	}()
	go func() {
		for event := range payCh {
			fmt.Printf("[订阅者] 收到最终支付事件: %s \n", event.Payload)
		}
	}()
	time.Sleep(200 * time.Millisecond)

	//发布事件
	fmt.Println("\n--- 发布订单事件 ---")
	eb.Publish("order", Event{Payload: "订单10000"})
	eb.Publish("order", Event{Payload: "订单10001"})
	eb.Publish("order", Event{Payload: "订单10002"})
	fmt.Println("\n--- 发布支付事件 ---")
	eb.Publish("pay", Event{Payload: "支付100001"})
	eb.Publish("pay", Event{Payload: "支付100002"})
	// 保持程序运行以便观察输出
	time.Sleep(10000 * time.Millisecond)
}
