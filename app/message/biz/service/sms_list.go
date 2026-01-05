package service

import (
	"context"
	message "gen/kitex_gen/message"
)

type SmsListService struct {
	ctx context.Context
}

// NewSmsListService new SmsListService
func NewSmsListService(ctx context.Context) *SmsListService {
	return &SmsListService{ctx: ctx}
}

// Run create note info
func (s *SmsListService) Run(req *message.SendSmsListReq) (resp *message.SendSmsListResp, err error) {
	// Finish your business logic.

	return
}
