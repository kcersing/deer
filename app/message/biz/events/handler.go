package events

import (
	"common/eventbus"
	"context"
	"gen/kitex_gen/message"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消息服务事件处理器 ============

// HandleSendUserMessages 发送用户消息处理器
func HandleSendUserMessages(ctx context.Context, event *eventbus.Event) error {

	//mapper := utils.NewCopierMapper[message.SendUserMessagesReq, eventbus.Event]()
	//var req = mapper.ToDTO(event)
	klog.Infof("[HandleSendUserMessages] ", event)
	//req, ok := event.Payload.(*message.SendUserMessagesReq)
	//if !ok {
	//	klog.Error("[HandleSendUserMessages] invalid payload type")
	//	return nil
	//}

	//klog.Infof("[HandleSendUserMessages] Processing: title=%s, content=%s", req.GetTitle(), req.GetContent())

	// TODO: 实现业务逻辑
	// 1. 保存消息到数据库
	// 2. 发送通知
	// 3. 记录审计日志

	return nil
}

// HandleSendMemberMessages 发送会员消息处理器
func HandleSendMemberMessages(ctx context.Context, event *eventbus.Event) error {
	klog.Infof("[HandleSendMemberMessages] Event received: %s", event.Id)
	req, ok := event.Payload.(*message.SendMemberMessagesReq)
	if !ok {
		klog.Error("[HandleSendMemberMessages] invalid payload type")
		return nil
	}

	klog.Infof("[HandleSendMemberMessages] Processing: Title=%s, Content=%s", req.GetTitle(), req.GetContent())
	//client.MemberClient.GetMemberIds()
	// TODO: 实现业务逻辑
	// 1. 更新消息状态
	// 2. 触发下游服务

	return nil
}
