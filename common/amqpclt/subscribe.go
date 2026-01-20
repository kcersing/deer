package amqpclt

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)
var _ Subscriber = (*Subscribe)(nil)
// Subscribe implements an amqp subscribe.
type Subscribe struct {
	conn     *amqp.Connection // AMQP 连接
	exchange string           // 交换机名称
}

// NewSubscribe creates an amqp subscribe.
func NewSubscribe(conn *amqp.Connection, exchange string) (*Subscribe, error) {
	// 不在此处创建和关闭 channel，实际的订阅会在 SubscribeRaw 中创建独立的 channel
	return &Subscribe{
		conn:     conn,
		exchange: exchange,
	}, nil
}

// SubscribeRaw subscribes and returns a channel with raw amqp delivery.
func (s *Subscribe) SubscribeRaw(_ context.Context) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, func() {}, fmt.Errorf("cannot allocate channel: %v", err)
	}

	if err := ch.Qos(3, 0, false); err != nil {
		return nil, func() {}, fmt.Errorf("cannot set qos: %v", err)
	}

	if err = declareExchange(ch, s.exchange); err != nil {
		return nil, func() {}, fmt.Errorf("cannot declare exchange: %v", err)
	}

	// 声明临时队列：非持久、自动删除、排他队列（随连接关闭自动删除）
	q, err := ch.QueueDeclare(
		"",    //队列名称（让服务器生成随机名）
		false, // durable: 非持久
		true,  // autoDelete: 当所有消费者断开后自动删除
		true,  // exclusive: 排他队列，仅限当前连接使用
		false, // noWait: 阻塞等待服务器响应
		nil,   // 额外参数
	)
	if err != nil {
		return nil, func() {}, fmt.Errorf("cannot declare queue: %v", err)
	}

	// 清理函数：删除队列并关闭通道
	cleanUp := func() {
		_, err := ch.QueueDelete(q.Name, false, false, false)
		if err != nil {
			klog.Errorf("cannot delete queue %s : %s", q.Name, err.Error())
		}
		if err := ch.Close(); err != nil {
			klog.Errorf("cannot close channel %s", err.Error())
		}
	}

	// 将队列绑定到交换机
	err = ch.QueueBind(q.Name, "", s.exchange, false, nil)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot bind: %v", err)
	}

	// 消费队列消息
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot consume queue: %v", err)
	}
	return msgs, cleanUp, nil
}

// Subscribe subscribes and returns a channel with CarEntity data.
func (s *Subscribe) Subscribe(c context.Context) (chan *Message, func(), error) {
	msgCh, cleanUp, err := s.SubscribeRaw(c)
	if err != nil {
		return nil, cleanUp, err
	}
	carCh := make(chan *Message)
	go func() {
		defer close(carCh) // 确保通道关闭，避免接收方阻塞
		for {
			select {
			case <-c.Done():
				return
			case msg, ok := <-msgCh:
				if !ok {
					return
				}
				var carEn Message
				// 反序列化消息
				if err := sonic.Unmarshal(msg.Body, &carEn); err != nil {
					klog.Errorf("cannot unmarshal message body: %v", err)
					// 反序列化失败：拒绝消息（不重新入队，避免死循环）
					if nackErr := msg.Nack(false, false); nackErr != nil {
						klog.Errorf("failed to nack message: %v", nackErr)
					}
					continue
				}

				// 发送到业务通道（阻塞直到业务方接收）
				select {
				case carCh <- &carEn:
					// 业务处理完成：手动确认消息
					if ackErr := msg.Ack(false); ackErr != nil {
						klog.Errorf("failed to ack message: %v", ackErr)
					}
				case <-c.Done():
					// 上下文取消，尝试重新入队或 nack
					if nackErr := msg.Nack(false, true); nackErr != nil {
						klog.Errorf("failed to nack message on context done: %v", nackErr)
					}
					return
				}
			}
		}
	}()
	return carCh, cleanUp, nil
}

// 声明交换机，类型为 fanout（广播模式）
func declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(
		exchange, // 交换机名称
		"fanout", // 类型：fanout 会将消息广播到所有绑定的队列
		true,     // durable: 交换机是否持久化
		false,    // autoDelete: 当没有队列绑定时是否自动删除
		false,    // internal: 是否为内部交换机
		false,    // noWait: 是否非阻塞
		nil,      // 参数
	)
}
