package mq

import "context"

type Publisher interface {
	Publish(ctx context.Context, s struct{})
}
type Subscriber interface {
	Subscribe(ctx context.Context) (ch chan struct{}, cleanUp func(), err error)
}
