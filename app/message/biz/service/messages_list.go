package service

import (
	"context"
	message "gen/kitex_gen/message"
)

type MessagesListService struct {
	ctx context.Context
}

// NewMessagesListService new MessagesListService
func NewMessagesListService(ctx context.Context) *MessagesListService {
	return &MessagesListService{ctx: ctx}
}

// Run create note info
func (s *MessagesListService) Run(req *message.MessagesListReq) (resp *message.MessagesListResp, err error) {
	// Finish your business logic.

	return
}
