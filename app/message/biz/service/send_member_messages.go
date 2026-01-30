package service

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/dal/db"
	"message/biz/dal/db/ent/messages"
	"message/biz/events"

	"github.com/cloudwego/kitex/pkg/klog"
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
	if err = events.SendMemberMessages(s.ctx, req); err != nil {
		klog.Errorf("Failed to publish SendMemberMessages event: %v", err)
	}
	return
}
