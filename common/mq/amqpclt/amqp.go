package amqpclt

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher implements an amqp publisher.
type Publisher struct {
	ch       *amqp.Channel // AMQP 通道
	exchange string        // 交换机名称
}

func NewPublisher(conn *amqp.Connection, exchange string) (*Publisher, error) {
	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	// 声明交换机
	if err = declareExchange(ch, exchange); err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}
	return &Publisher{
		ch:       ch,
		exchange: exchange,
	}, nil
}

// Publish publishes a message.
func (p *Publisher) Publish(_ context.Context, s interface{}) error {
	// 将消息结构体序列化为 JSON

	body, err := sonic.Marshal(&s)

	if err != nil {
		return fmt.Errorf("cannot marshal: %v", err)
	}

	return p.ch.Publish(
		p.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			//ContentType:  "application/json",
			Body: body,
			//DeliveryMode: amqp.Persistent, // 消息持久化
		},
	)
}

// Subscriber implements an amqp subscriber.
type Subscriber struct {
	conn     *amqp.Connection // AMQP 连接
	exchange string           // 交换机名称
}

// NewSubscriber creates an amqp subscriber.
func NewSubscriber(conn *amqp.Connection, exchange string) (*Subscriber, error) {

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	defer ch.Close()

	if err = declareExchange(ch, exchange); err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}

	return &Subscriber{
		conn:     conn,
		exchange: exchange,
	}, nil
}

// SubscribeRaw subscribes and returns a channel with raw amqp delivery.
func (s *Subscriber) SubscribeRaw(_ context.Context) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, func() {}, fmt.Errorf("cannot allocate channel: %v", err)
	}

	if err = declareExchange(ch, s.exchange); err != nil {
		return nil, func() {}, fmt.Errorf("cannot declare exchange: %v", err)
	}
	//关闭通道的函数
	closeCh := func() {
		err := ch.Close()
		if err != nil {
			klog.Errorf("cannot close channel %s", err.Error())
		}
	}

	// 声明临时队列
	q, err := ch.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		return nil, closeCh, fmt.Errorf("cannot declare queue: %v", err)
	}

	// 清理函数：删除队列并关闭通道
	cleanUp := func() {
		_, err := ch.QueueDelete(q.Name, false, false, false)
		if err != nil {
			klog.Errorf("cannot delete queue %s : %s", q.Name, err.Error())
		}
		closeCh()
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
func (s *Subscriber) Subscribe(c context.Context) (chan *interface{}, func(), error) {
	msgCh, cleanUp, err := s.SubscribeRaw(c)
	if err != nil {
		return nil, cleanUp, err
	}
	carCh := make(chan *interface{})
	go func() {
		defer close(carCh) // 确保通道关闭，避免接收方阻塞
		for msg := range msgCh {
			var carEn interface{}
			// 反序列化消息
			klog.Infof("msg.Body: ", msg.Body)
			if err := sonic.Unmarshal(msg.Body, &carEn); err != nil {
				klog.Errorf("cannot unmarshal message body: %v", err)
				// 反序列化失败：拒绝消息（不重新入队，避免死循环）
				if nackErr := msg.Nack(false, false); nackErr != nil {
					klog.Errorf("failed to nack message: %v", nackErr)
				}
				continue
			}
			klog.Infof("carEn: ", carEn)
			// 发送到业务通道（阻塞直到业务方接收）
			carCh <- &carEn
			// 业务处理完成：手动确认消息
			if ackErr := msg.Ack(false); ackErr != nil {
				klog.Errorf("failed to ack message: %v", ackErr)
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
