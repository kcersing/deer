package events

import (
	"context"
	"gen/kitex_gen/message"

	"github.com/cloudwego/kitex/pkg/klog"
)

// SendMessages 发布一个用户消息事件
// 仅发布到本地内存总线，避免 MQ 回流导致本服务重复消费
func SendMessages(ctx context.Context, req *message.SendMessagesReq) error {
	klog.Info("[SendUserMessages] req: %s", req.String())
	GetManager().Publisher.Local(
		ctx,
		EventSendMessages, // 使用在 events.go 中定义的常量
		req,
	)
	return nil

}
