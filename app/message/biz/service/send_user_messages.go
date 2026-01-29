package service

import (
	eventbus2 "common/eventbus"
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/dal/db"
	"message/biz/dal/db/ent/messages"

	"message/biz/events"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cockroachdb/errors"
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

	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		klog.Error("Tx")
		return nil, err
	}

	_, err = tx.Messages.Create().
		SetCreatedID(req.GetCreatedId()).
		SetContent(req.GetContent()).
		SetFromUserID(req.GetUserId()).
		SetTitle(req.GetTitle()).
		SetStatus(messages.Status(req.GetStatus().String())).
		SetType(messages.Type(req.GetType().String())).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "create Member Profile failed"))
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	eb := events.GetManager().Bus
	event := eventbus2.NewEvent("send_user_messages", req)
	eb.Publish(s.ctx, event)
	return
}
