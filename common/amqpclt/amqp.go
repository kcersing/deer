package amqpclt

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher 定义发布接口。
type Publisher interface {
	Publish(c context.Context, routingKey, correlationId string, m Message) (err error)

	PublishBatch(ctx context.Context, messages []BatchMessage) ([]int, error)
	PublishAsync(ctx context.Context, messages []BatchMessage, resultCh chan PublishResult)

	Close()
}

// Subscriber 定义订阅接口。
type Subscriber interface {
	SubscribeRaw(c context.Context) (<-chan amqp.Delivery, func(), error)
	Subscribe(c context.Context) (chan *Message, func(), error)
}
type Connection struct {
	Conn   *amqp.Connection
	ChFunc func() error
}

type Message struct {
	Event     string    `json:"event"`
	Payload   any       `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}
