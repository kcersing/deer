package service

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
)

type SendMemberMessagesService struct {
	ctx context.Context
}

// NewSendMemberMessagesService new SendMemberMessagesService
func NewSendMemberMessagesService(ctx context.Context) *SendMemberMessagesService {
	return &SendMemberMessagesService{ctx: ctx}
}

// Run create note info
func (s *SendMemberMessagesService) Run(req *message.SendMemberMessagesReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
