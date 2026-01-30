package events

import (
	"common/eventbus"
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

// InitMessageConsumers 初始化消息服务的所有事件消费者。
func InitMessageConsumers() error {
	registry := GetManager().Registry
	if registry == nil {
		// 这是一个防御性检查，正常情况下不应该发生，因为 Bootstrap 已经初始化了它。
		klog.Fatal("[Events] Consumer registry not initialized before InitMessageConsumers.")
		return nil
	}

	// 1. 注册处理器
	err := registry.RegisterHandler(
		handlerSendUserMessages,
		eventbus.WrapTyped(eventbus.TypedHandler[*message.SendUserMessagesReq](HandleSendUserMessages)),
	)
	if err != nil {
		klog.Errorf("Failed to register handler '%s': %v", handlerSendUserMessages, err)
		return err
	}

	err = registry.RegisterHandler(
		handlerSendMemberMessages,
		eventbus.WrapTyped(eventbus.TypedHandler[*message.SendMemberMessagesReq](HandleSendMemberMessages)),
	)
	if err != nil {
		klog.Errorf("Failed to register handler '%s': %v", handlerSendMemberMessages, err)
		return err
	}

	// 2. 注册消费者 (将主题与处理器绑定)
	err = registry.RegisterConsumer(EventSendUserMessages, handlerSendUserMessages, 10)
	if err != nil {
		klog.Errorf("Failed to register consumer for event '%s': %v", EventSendUserMessages, err)
		return err
	}

	err = registry.RegisterConsumer(EventSendMemberMessages, handlerSendMemberMessages, 5)
	if err != nil {
		klog.Errorf("Failed to register consumer for event '%s': %v", EventSendMemberMessages, err)
		return err
	}

	klog.Info("[Events] All message consumers have been configured.")
	return nil
}
