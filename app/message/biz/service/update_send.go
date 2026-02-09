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

	klog.Info(req.String())
	_, err = db.Client.Messages.UpdateOneID(req.GetId()).
		SetContent(req.GetContent()).
		SetStatus(req.GetStatus()).
		SetTitle(req.GetTitle()).
		SetType(req.GetType()).
		Save(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "update message failed")
	}

	return &base.NilResponse{}, nil
}
