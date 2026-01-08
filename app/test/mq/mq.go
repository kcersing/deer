package main

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

var (
	Client     *amqp.Connection
	onceClient sync.Once
)

func InitMQ() {
	onceClient.Do(func() {
		Client = initMQ()
	})
}

func initMQ() *amqp.Connection {

	client, err := amqp.Dial("amqp://kcersing:kcer-913639@127.0.0.1:5672/")
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	return client
}

// Publisher 定义发布接口。
type Publisher interface {
	Publish(c context.Context, a any) (err error)
}

// Subscriber 定义订阅接口。
type Subscriber interface {
	Subscribe(c context.Context) (ch chan *any, cleanUp func(), err error)
}
