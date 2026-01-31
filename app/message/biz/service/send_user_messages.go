package service

import (
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/message"
	"message/biz/dal/db"
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
		SetType(req.GetType()).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "create Member Profile failed"))
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	if err = events.SendUserMessages(s.ctx, req); err != nil {
		klog.Errorf("Failed to publish SendUserMessages event: %v", err)
	}

	return &base.NilResponse{}, nil
}
