package service

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
)

type SendSmsService struct {
	ctx context.Context
}

// NewSendSmsService new SendSmsService
func NewSendSmsService(ctx context.Context) *SendSmsService {
	return &SendSmsService{ctx: ctx}
}

// Run create note info
func (s *SendSmsService) Run(req *message.SendSmsReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	return
}
