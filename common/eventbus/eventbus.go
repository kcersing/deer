package eventbus

import "context"

type EventBusInterface interface {
	Subscribe(eventType string, handler Handler) error

	SubscribeOnce(eventType string, handler Handler) error

	Unsubscribe(eventType string, handler Handler) error

	Publish(ctx context.Context, event *Event) error

	Close() error
}
