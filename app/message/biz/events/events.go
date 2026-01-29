package events

import (
	"common/eventbus"
	"errors"
	"gen/kitex_gen/message"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消息服务事件常量 ============

const (
	// 事件主题
	EventSendUserMessages   = "send_user_messages"
	EventSendMemberMessages = "send_member_messages"

	// 处理器名称
	handlerSendUserMessages   = "send_user_messages_handler"
	handlerSendMemberMessages = "send_member_messages_handler"
)

// ============ 消费者注册 (现代化版本) ============

var consumerRegistry *eventbus.ConsumerRegistry

// InitMessageConsumers 初始化消息服务事件消费者
func InitMessageConsumers() error {
	if consumerRegistry == nil {
		consumerRegistry = eventbus.NewConsumerRegistry()
	}

	// 注册类型安全的处理器
	err := consumerRegistry.RegisterHandler(
		handlerSendUserMessages,
		eventbus.WrapTyped(eventbus.TypedHandler[*message.SendUserMessagesReq](HandleSendUserMessages)),
	)
	if err != nil {
		return err
	}

	err = consumerRegistry.RegisterHandler(
		handlerSendMemberMessages,
		eventbus.WrapTyped(eventbus.TypedHandler[*message.SendMemberMessagesReq](HandleSendMemberMessages)),
	)
	if err != nil {
		return err
	}

	// 注册消费者 (可以附带池配置)
	err = consumerRegistry.RegisterConsumer(EventSendUserMessages, handlerSendUserMessages, 10)
	if err != nil {
		return err
	}
	err = consumerRegistry.RegisterConsumer(EventSendMemberMessages, handlerSendMemberMessages, 5)
	if err != nil {
		return err
	}

	klog.Infof("[Events] Message consumers initialized")
	return nil
}

// StartMessageConsumers 启动消息服务消费者
func StartMessageConsumers() error {
	if consumerRegistry == nil {
		return errors.New("consumer registry not initialized")
	}
	klog.Infof("[Events] Starting message consumers...")
	return consumerRegistry.StartAll(GetManager().Bus)
}
