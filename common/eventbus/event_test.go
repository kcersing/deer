package eventbus

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestEventBus_Run(t *testing.T) {
	eb := NewEventBus()
	var wg sync.WaitGroup
	//// 订阅“test”主题活动
	subscribe := eb.Subscribe("test")
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range subscribe {
			fmt.Printf("[订阅者 A]订阅主题: %s \n", event.Payload)
		}
		fmt.Println("[订阅者 A] 通道已关闭，停止接收。")
	}()

	subscribeB := eb.Subscribe("order")
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range subscribeB {
			fmt.Printf("[订阅者 B]订阅主题: %s \n", event.Payload)
		}
		fmt.Println("[订阅者 B] 通道已关闭，停止接收。")
	}()
	subscribeC := eb.Subscribe("pay")
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range subscribeC {
			fmt.Printf("[订阅者 C]订阅主题: %s \n", event.Payload)
		}
		fmt.Println("[订阅者 C] 通道已关闭，停止接收。")
	}()
	// 定义具体的处理函数，注意第二个参数已经是具体的 *OrderPayload 类型了！
	orderHandler := func(ctx context.Context, order *OrderPayload, e eventbus.Event) error {
		// 这里不需要写: order := e.Payload.(*OrderPayload)
		// 直接使用 IDE 的自动补全
		fmt.Printf("处理订单: ID=%s 金额=%.2f\n", order.OrderId, order.Amount)
		return nil
	}

	// 注册时使用 WrapTyped
	eb.SubscribeAsync("topic_order", eventbus.WrapTyped(orderHandler), 1)

	// 等待所有订阅 goroutine 启动
	time.Sleep(100 * time.Millisecond)

	e := NewEvent("test", map[string]any{
		"postId": 1,
		"title":  "Welcome to Leapcell",
		"author": "Leapcell",
	})
	eb.Publish(context.Background(), e)

	eb.Publish(context.Background(), NewEvent("order", "订单10000"))
	eb.Publish(context.Background(), NewEvent("pay", "支付100001"))
	// 暂无订阅者的主题
	eb.Publish(context.Background(), NewEvent("msg", "消息"))
	fmt.Println("--- 事件发布完成 ---")

	time.Sleep(time.Second * 3)
	// 取消订阅“test”主题活动
	eb.Unsubscribe("test", subscribe)
	// 再次发布，只有 B 会收到
	fmt.Println("\n--- 再次发布订单事件（A已取消订阅） ---")
	eb.Publish(context.Background(), NewEvent("order", "订单发货"))
	time.Sleep(500 * time.Millisecond)
}
