package events

import (
	"common/eventbus"
	"context"
	"gen/kitex_gen/message"

	"github.com/cloudwego/kitex/pkg/klog"
)

// SendUserMessages 发布一个用户消息事件
// 这是一个分布式事件，它将被发布到内存总线和MQ
func SendUserMessages(ctx context.Context, req *message.SendUserMessagesReq) error {
	klog.Info("[SendUserMessages] req: %s", req.String())
	return GetManager().Publisher.Publish(
		ctx,
		EventSendUserMessages, // 使用在 events.go 中定义的常量
		req,
		eventbus.WithScope(eventbus.ScopeLocal), // 明确发布范围
	)

}

// SendMemberMessages 发布一个会员消息事件
func SendMemberMessages(ctx context.Context, req *message.SendMemberMessagesReq) error {
	return GetManager().Publisher.Publish(
		ctx,
		EventSendMemberMessages,
		req,
		eventbus.WithScope(eventbus.ScopeLocal),
	)
}
