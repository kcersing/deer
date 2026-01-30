package events

import (
	"admin/rpc/client"
	"common/eventbus"
	"context"
	"gen/kitex_gen/message"
	"gen/kitex_gen/user"
	"message/biz/dal/db"
	"message/biz/dal/db/ent"
	"message/biz/dal/db/ent/messages"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消息服务事件处理器 (类型安全 & 可审计版本) ============

// HandleSendUserMessages 发送用户消息处理器
func HandleSendUserMessages(ctx context.Context, req *message.SendUserMessagesReq, event eventbus.Event) error {
	klog.Infof("[Handler] 处理发送用户消息事件: Title=%s, EventID=%s", req.GetTitle(), event.Id)

	// 从 context 中获取审计条目并填充业务数据
	if entry, ok := GetAuditEntry(ctx); ok {
		entry.Details["UserID"] = req.GetUserId()
		entry.Details["Title"] = req.GetTitle()
		entry.Details["ContentLength"] = len(req.GetContent())
		entry.Details["MessageType"] = req.GetType().String()
	}

	// TODO: 实现核心业务逻辑
	// 1. 保存消息到数据库
	// 2. 发送通知

	klog.Info(event)
	resp, err := client.UserClient.GetUserIds(ctx, &user.GetUserListReq{
		Tags: req.TagId,
	})
	if err != nil {
		klog.Errorf("[Handler] 获取用户ID失败: %v", err)
		return err
	}

	// 检查是否有用户ID
	if len(resp.Data) == 0 {
		klog.Warnf("[Handler] 没有用户符合标签: %v", req.TagId)
		return nil
	}

	_, err = db.Client.Messages.Create().
		SetCreatedID(req.GetCreatedId()).
		SetContent(req.GetContent()).
		SetFromUserID(req.GetUserId()).
		SetTitle(req.GetTitle()).
		SetStatus(messages.Status(req.GetStatus().String())).
		SetType(messages.Type(req.GetType().String())).
		Save(ctx)

	var createAll []*ent.MessagesSentRecordsCreate
	for _, userID := range resp.Data {
		createAll = append(createAll,
			db.Client.MessagesSentRecords.Create().
				SetMessageID(req.GetUserId()).SetToUserID(userID).
				SetReceivedAt(time.Now()))
	}

	_, err = db.Client.MessagesSentRecords.CreateBulk(createAll...).Save(ctx)
	if err != nil {
		klog.Errorf("[Handler] 保存消息发送记录失败: %v", err)
		return err
	}

	return nil
}

// HandleSendMemberMessages 发送会员消息处理器
func HandleSendMemberMessages(ctx context.Context, req *message.SendMemberMessagesReq, event eventbus.Event) error {
	klog.Infof("[Handler] 处理发送会员消息事件: Title=%s, EventID=%s", req.GetTitle(), event.Id)

	if entry, ok := GetAuditEntry(ctx); ok {
		entry.Details["Title"] = req.GetTitle()
		entry.Details["MessageType"] = req.GetType().String()
	}

	// TODO: 实现核心业务逻辑

	return nil
}
