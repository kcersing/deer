package mq

import (
	"fmt"
	"message/conf"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Client     *amqp.Connection
	onceClient sync.Once
)

func InitMQ() {
	onceClient.Do(func() {
		c := conf.GetConf().RabbitMq
		url := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.User, c.Password, c.Host, c.Port)
		Client = initMQ(url)
	})
}

func initMQ(url string) *amqp.Connection {

	dialConfig := amqp.Config{
		Heartbeat: 30 * time.Second, // 设置心跳间隔为30秒
	}
	client, err := amqp.DialConfig(url, dialConfig)
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	// client.IsClosed()
	return client
}
