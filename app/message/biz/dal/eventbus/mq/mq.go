package mq

import (
	"common/amqpclt"
	"common/eventbus"
	"context"
	"fmt"
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
		//c := conf.GetConf().RabbitMq
		// fmt.Sprintf(c.Host, c.User, c.Password, c.Host, c.Port)
		url := fmt.Sprintf("amqp://%s:%s@%s:%d/", "kcersing", "kcer-913639", "127.0.0.1", 5672)
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
