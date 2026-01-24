package events

import (
	"context"

	"common/eventbus"
	"gen/kitex_gen/message"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消息服务事件处理器 ============

// HandleSendUserMessages 发送用户消息处理器
func HandleSendUserMessages(ctx context.Context, event *eventbus.Event) error {
	req, ok := event.Payload.(*message.SendUserMessagesReq)
	if !ok {
		klog.Error("[SendUserMessages] invalid payload type")
		return nil
	}

	klog.Infof("[SendUserMessages] Processing: userID=%s, title=%s",
		req.GetUserId(), req.GetTitle())

	// TODO: 实现业务逻辑
	// 1. 保存消息到数据库
	// 2. 发送通知
	// 3. 记录审计日志

	return nil
}

// HandleSendMemberMessages 发送会员消息处理器
func HandleSendMemberMessages(ctx context.Context, event *eventbus.Event) error {
	klog.Infof("[SendMemberMessages] Event received: %s", event.Id)

	// TODO: 实现业务逻辑
	// 1. 更新消息状态
	// 2. 触发下游服务

	return nil
}
