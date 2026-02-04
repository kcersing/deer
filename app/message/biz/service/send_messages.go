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

type SendMessagesService struct {
	ctx context.Context
}

// NewSendMessagesService new SendMessagesService
func NewSendMessagesService(ctx context.Context) *SendMessagesService {
	return &SendMessagesService{ctx: ctx}
}

// Run create note info
func (s *SendMessagesService) Run(req *message.SendMessagesReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		klog.Error("Tx")
		return nil, err
	}

	save, err := tx.Messages.Create().
		SetCreatedID(req.GetCreatedId()).
		SetContent(req.GetContent()).
		SetTitle(req.GetTitle()).
		SetType(req.GetType()).
		Save(s.ctx)
	if err != nil {
		return nil, rollback(tx, errors.Wrap(err, "create Member Profile failed"))
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	req.Id = save.ID
	if err = events.SendUserMessages(s.ctx, req); err != nil {
		klog.Errorf("Failed to publish SendUserMessages event: %v", err)
	}

	return &base.NilResponse{}, nil
}
