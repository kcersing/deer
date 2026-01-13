package service

import (
	"common/mq/amqpclt"
	"context"
	"fmt"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/dal/mq"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type SendMemberMessagesService struct {
	ctx context.Context
}

// NewSendMemberMessagesService new SendMemberMessagesService
func NewSendMemberMessagesService(ctx context.Context) *SendMemberMessagesService {
	return &SendMemberMessagesService{ctx: ctx}
}

// Run create note info
func (s *SendMemberMessagesService) Run(req *message.SendMemberMessagesReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	mq.InitMQ()
	m, err := amqpclt.NewPublisher(mq.Client, "deer")
	if err != nil {
		klog.Fatal("cannot create publisher", err)
	}
	subscriber, err := amqpclt.NewSubscriber(mq.Client, "deer")
	if err != nil {
		klog.Fatal("cannot create subscriber", err.Error())
	}

	go func() {
		for i := 0; i < 50; i++ {
			time.Sleep(1 * time.Second)
			var al = strconv.Itoa(i + 100000)

			err = m.Publish(s.ctx, al)
			if err != nil {
				if err != nil {
					klog.Fatal("11111111", err.Error())
				}
			}
		}

	}()
	func() {
		msgs, cleanUp, err := subscriber.Subscribe(s.ctx)
		defer cleanUp()
		if err != nil {
			klog.Fatal("11111111", err.Error())
		}
		for ms := range msgs {
			time.Sleep(1 * time.Second)
			klog.Info(*ms)
			fmt.Print(*ms)

		}

	}()
	return
}
