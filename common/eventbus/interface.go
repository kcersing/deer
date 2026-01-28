package eventbus

import "context"

type Subscribe interface {
	Subscribe(topic string) EventChan
	SubscribeAsync(topic string, handler Handler, concurrency int)
	SubscribeWithPool(topic string, handler Handler, workerNum int32) *ConsumerPool
	Unsubscribe(topic string, ch EventChan)
}
type Publish interface {
	Publish(ctx context.Context, event *Event)
}
type Bus interface {
	Use(mw ...Middleware)
	Close() error
}
