package service

import (
	eventbus2 "common/eventbus"
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/dal/db"
	"message/biz/dal/db/ent/messages"
	"message/biz/events"
)

type SendMemberMessagesService struct {
	ctx context.Context
}

// NewSendMemberMessagesService new SendMemberMessagesService
func NewSendMemberMessagesService(ctx context.Context) *SendMemberMessagesService {
	return &SendMemberMessagesService{ctx: ctx}
}
func (s *SendMemberMessagesService) Run(req *message.SendMemberMessagesReq) (resp *base.NilResponse, err error) {
	_, err = db.Client.Messages.Create().
		SetTitle(req.GetTitle()).
		SetContent(req.GetContent()).
		SetFromUserID(req.GetCreatedId()).
		SetType(messages.Type(req.GetType())).
		SetCreatedID(req.GetCreatedId()).
		Save(s.ctx)

	if err != nil {
		return nil, err
	}
	eb := events.GetGlobalEventBus()
	event := eventbus2.NewEvent("send_member_messages", req)
	eb.Publish(s.ctx, event)
	return
}
