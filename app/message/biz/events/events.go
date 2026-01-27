package events

import (
	"common/eventbus"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消息服务事件常量 ============

const (
	// 发送用户消息事件
	EventSendUserMessages = "send_user_messages"

	// 发送会员消息事件
	EventSendMemberMessages = "send_member_messages"
)

// ============ 事件验证 ============

var ValidMessageEvents = map[string]bool{
	EventSendUserMessages:   true,
	EventSendMemberMessages: true,
}

func IsValidMessageEvent(eventType string) bool {
	return ValidMessageEvents[eventType]
}

// ============ 消费者注册 ============
// InitMessageConsumers 初始化消息服务事件消费者
func InitMessageConsumers() error {
	if consumerRegistry == nil {
		consumerRegistry = eventbus.NewConsumerRegistry()
	}

	// 注册处理器
	err := consumerRegistry.RegisterHandler("send_user_messages", eventbus.EventHandlerFunc(HandleSendUserMessages))
	if err != nil {
		return err
	}
	klog.Infof("[InitMessageConsumers] send user messages ok")

	err = consumerRegistry.RegisterHandler("send_member_messages", eventbus.EventHandlerFunc(HandleSendMemberMessages))
	if err != nil {
		return err
	}
	klog.Infof("[InitMessageConsumers] send member messages ok")
	// 注册消费者
	err = consumerRegistry.RegisterConsumer(EventSendUserMessages, "send_user_messages", 10)
	if err != nil {
		return err
	}
	err = consumerRegistry.RegisterConsumer(EventSendMemberMessages, "send_member_messages", 5)
	if err != nil {
		return err
	}

	klog.Infof("[InitMessageConsumers] Event consumers initialized")
	return nil
}

// StartMessageConsumers 启动消息服务消费者
func StartMessageConsumers() error {
	if consumerRegistry == nil {
		return errors.New("consumer registry not initialized")
	}
	klog.Infof("[InitMessageConsumers] Start Message Consumers")
	return consumerRegistry.StartAll(GetGlobalEventBus())
}
