package mq

import "context"

type Msg struct {
	Type string
	Data interface{}
}

// Publisher 定义发布接口。
type Publisher interface {
	Publish(c context.Context, a any) (err error)
}

// Subscriber 定义订阅接口。
type Subscriber interface {
	Subscribe(c context.Context) (ch chan *any, cleanUp func(), err error)
}
