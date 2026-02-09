package events

import (
	"admin/rpc/client"
	"common/eventbus"
	"context"
	"gen/kitex_gen/message"
	"gen/kitex_gen/user"
	"message/biz/dal/db"
	"message/biz/dal/db/ent"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// ============ 消息服务事件处理器 (类型安全 & 可审计版本) ============

// HandleSendMessages 发送消息处理器
func HandleSendMessages(ctx context.Context, req *message.SendMessagesReq, event eventbus.Event) error {
	klog.Infof("[Handler] 处理发送用户消息事件: Title=%s, EventID=%s", req.GetTitle(), event.Id)

	// 从 context 中获取审计条目并填充业务数据
	if entry, ok := GetAuditEntry(ctx); ok {

		entry.Details["Title"] = req.GetTitle()
		entry.Details["ContentLength"] = len(req.GetContent())
		entry.Details["MessageType"] = req.GetType()
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
		SetTitle(req.GetTitle()).
		SetStatus(req.GetStatus()).
		SetType(req.GetType()).
		Save(ctx)

	var createAll []*ent.MessagesSentRecordsCreate
	for _, userID := range resp.Data {
		createAll = append(createAll,
			db.Client.MessagesSentRecords.Create().
				SetMessageID(req.GetId()).
				SetToUserID(userID).
				SetReceivedAt(time.Now()))
	}

	_, err = db.Client.MessagesSentRecords.CreateBulk(createAll...).Save(ctx)
	if err != nil {
		klog.Errorf("[Handler] 保存消息发送记录失败: %v", err)
		return err
	}

	return nil
}
