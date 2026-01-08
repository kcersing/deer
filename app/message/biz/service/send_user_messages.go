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

	return
}
