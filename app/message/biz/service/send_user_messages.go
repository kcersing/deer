package service

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
)

type SendUserMessagesService struct {
	ctx context.Context
}

// NewSendUserMessagesService new SendUserMessagesService
func NewSendUserMessagesService(ctx context.Context) *SendUserMessagesService {
	return &SendUserMessagesService{ctx: ctx}
}

// Run create note info
func (s *SendUserMessagesService) Run(req *message.SendUserMessagesReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	// eb := events.GetGlobalEventBus()

	// // 发布事件到事件总线
	// // 消费者已在应用启动时注册，无需在此处创建
	// event := eventbus2.NewEvent(event.EventSendUserMessages, req)

	// eb.Publish(s.ctx, event)

	//tx, err := db.Client.Tx(s.ctx)
	//if err != nil {
	//	klog.Error("Tx")
	//	return nil, err
	//}
	//_, err = tx.Messages.Create().
	//	SetCreatedID(req.GetCreatedId()).
	//	SetContent(req.GetContent()).
	//	SetFromUserID(req.GetUserId()).
	//	SetTitle(req.GetTitle()).
	//	SetType(messages.Type(req.GetType())).
	//	Save(s.ctx)
	//if err != nil {
	//	return nil, rollback(tx, errors.Wrap(err, "create Member Profile failed"))
	//}
	//if err = tx.Commit(); err != nil {
	//	return nil, err
	//}

	return
}
