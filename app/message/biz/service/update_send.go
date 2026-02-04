package service

import (
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/message"
	"message/biz/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cockroachdb/errors"
)

type UpdateSendService struct {
	ctx context.Context
}

// NewUpdateSendService new UpdateSendService
func NewUpdateSendService(ctx context.Context) *UpdateSendService {
	return &UpdateSendService{ctx: ctx}
}

// Run create note info
func (s *UpdateSendService) Run(req *message.SendMessagesReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	tx, err := db.Client.Tx(s.ctx)
	if err != nil {
		klog.Error("Tx")
		return nil, err
	}

	_, err = tx.Messages.UpdateOneID(req.Id).
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

	return &base.NilResponse{}, nil
}
