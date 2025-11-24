package eventbus

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMQEventBus struct {
	conn         *amqp.Connection
	ch           *amqp.Channel
	exchangeName string
}

// Close 关闭连接和通道
func (rb *RabbitMQEventBus) Close() {
	if rb.ch != nil {
		rb.ch.Close()
	}
	if rb.conn != nil {
		rb.conn.Close()
	}
}

// Publish 发布事件到 RabbitMQ Exchange
func (rb *RabbitMQEventBus) Publish(topic string, payload any) error {
	event := Event{Topic: topic, Payload: payload}
	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to serialize event: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 发布事件到 RabbitMQ Exchange

	err = rb.ch.PublishWithContext(ctx,
		rb.exchangeName,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			// delivery mode 2: 持久化消息（确保 RabbitMQ 重启后消息不丢失）
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	fmt.Printf(" [x] Sent '%s' to Exchange '%s'\n", body, rb.exchangeName)
	return nil
}

func (rb *RabbitMQEventBus) Subscribe(consumerName string, handler func(event Event)) error {
	//声明一个队列
	q, err := rb.ch.QueueDeclare(
		consumerName, // 队列名称 , 使用消费者名称作为队列名，可以实现负载均衡
		true,         // durable: true 表示队列持久化
		false,        // delete when unused
		false,        // exclusive: false 表示多个消费者可以连接到此队列
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}
	// 绑定队列到 Exchange，使用消费者名称作为路由键
	err = rb.ch.QueueBind(
		q.Name,          // 队列名称
		"",              // routing key (fanout 模式不需要路由键)
		rb.exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}
	// 注册消费者
	msgs, err := rb.ch.Consume(
		q.Name,       // queue
		consumerName, // consumer tag
		false,        // auto-ack: false 表示手动确认消息，实现可靠消费
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}
	//启动一个 goroutine 来持续监听消息
	go func() {
		for d := range msgs {
			var event Event
			if err := json.Unmarshal(d.Body, &event); err != nil {
				fmt.Errorf("error unmarshalling message: %s", err)
				d.Ack(false) // 无法解析的消息直接确认，避免阻塞
				continue
			}
			// 调用订阅者的处理函数
			handler(event)
			// 手动确认（ACK）：告诉 RabbitMQ 消息已成功处理
			d.Ack(false)
		}
	}()
	fmt.Printf(" [*] Waiting for messages in queue '%s'. To exit press CTRL+C\n", q.Name)
	return nil
}

func NewRabbitMQBus(amqpUrl, exchangeName string) (*RabbitMQEventBus, error) {

	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	// 声明一个 Fanout Exchange，用于 Publish/Subscribe 模式
	// durable: true 表示 Exchange 重启后依然存在
	err = ch.ExchangeDeclare(
		exchangeName,        // exchange name
		amqp.ExchangeFanout, // exchange type: fanout
		true,                // durable: true 表示 Exchange 重启后依然存在
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare an exchange: %w", err)
	}
	return &RabbitMQEventBus{
		conn:         conn,
		ch:           ch,
		exchangeName: exchangeName,
	}, nil
}
