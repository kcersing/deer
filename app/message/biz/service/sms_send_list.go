package service

import (
	"context"
	message "gen/kitex_gen/message"
)

type SmsSendListService struct {
	ctx context.Context
}

// NewSmsSendListService new SmsSendListService
func NewSmsSendListService(ctx context.Context) *SmsSendListService {
	return &SmsSendListService{ctx: ctx}
}

// Run create note info
func (s *SmsSendListService) Run(req *message.SmsSendListReq) (resp *message.SmsSendListResp, err error) {
	// Finish your business logic.

	return
}
