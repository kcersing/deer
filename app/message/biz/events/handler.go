package events

import (
	"common/eventbus"
	"context"
	"gen/kitex_gen/message"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消息服务事件处理器 (类型安全版本) ============

// HandleSendUserMessages 发送用户消息处理器
func HandleSendUserMessages(ctx context.Context, req *message.SendUserMessagesReq, event eventbus.Event) error {
	klog.Infof("[Handler] 收到发送用户消息事件: Title=%s, EventID=%s", req.GetTitle(), event.Id)

	// TODO: 实现业务逻辑
	// 1. 保存消息到数据库
	// 2. 发送通知
	// 3. 记录审计日志

	return nil
}

// HandleSendMemberMessages 发送会员消息处理器
func HandleSendMemberMessages(ctx context.Context, req *message.SendMemberMessagesReq, event eventbus.Event) error {
	klog.Infof("[Handler] 收到发送会员消息事件: Title=%s, EventID=%s", req.GetTitle(), event.Id)

	// TODO: 实现业务逻辑
	// 1. 更新消息状态
	// 2. 触发下游服务

	return nil
}
