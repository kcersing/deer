package service

import (
	"context"
	base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
)

type SmsService struct {
	ctx context.Context
}

// NewSmsService new SmsService
func NewSmsService(ctx context.Context) *SmsService {
	return &SmsService{ctx: ctx}
}

// Run create note info
func (s *SmsService) Run(req *base.IdReq) (resp *message.SmsResp, err error) {
	// Finish your business logic.

	return
}
