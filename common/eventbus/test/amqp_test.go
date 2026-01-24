package test

import (
	"common/eventbus"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestEventBusAmqp_Run(t *testing.T) {
	// 确保您本地的 RabbitMQ 服务正在运行
	amqpURL := "amqp://guest:guest@localhost:5672/"
	exchangeName := "event_bus_exchange"

	bus, err := eventbus.NewRabbitMQBus(amqpURL, exchangeName)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer bus.Close()

	// 订阅者 A
	bus.Subscribe("subscriber_A_queue", func(event eventbus.Event) {
		fmt.Printf("[A] Received event: Topic=%s, Payload=%v\n", event.Topic, event.Payload)
	})

	// 订阅者 B (会收到同样的 Fanout 消息)
	bus.Subscribe("subscriber_B_queue", func(event eventbus.Event) {
		fmt.Printf("[B] Received event: Topic=%s, Payload=%v\n", event.Topic, event.Payload)
	})

	// 等待消费者启动并绑定队列
	time.Sleep(1 * time.Second)

	// 发布事件
	bus.Publish("orders", "订单 1001 已创建")
	bus.Publish("updates", map[string]interface{}{"version": "1.2", "status": "released"})

	fmt.Println("Events published. Waiting for consumers...")

	// 保持 main goroutine 运行以监听消息
	select {}
}
