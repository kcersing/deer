package service

import (
	"context"
	"gen/kitex_gen/base"

	"message/biz/dal/db"
	"message/biz/dal/db/ent/messages"
)

type DeleteMessagesService struct {
	ctx context.Context
}

// NewDeleteMessagesService new DeleteMessagesService
func NewDeleteMessagesService(ctx context.Context) *DeleteMessagesService {
	return &DeleteMessagesService{ctx: ctx}
}

// Run create note info
func (s *DeleteMessagesService) Run(req *base.IdReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	_, err = db.Client.Messages.Delete().Where(messages.IDEQ(req.GetId())).Exec(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
