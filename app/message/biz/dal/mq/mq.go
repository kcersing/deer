package mq

import (
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
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
	//c := conf.GetConf().RabbitMq
	//klog.Info(c)
	//client, err := amqp.Dial(fmt.Sprintf(c.Host, c.User, c.Password, c.Host, c.Port))
	client, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", "kcersing", "kcer-913639", "127.0.0.1", 5672))
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	return client
}
