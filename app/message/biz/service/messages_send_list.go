package service

import (
	"context"
	message "gen/kitex_gen/message"
)

type MessagesSendListService struct {
	ctx context.Context
}

// NewMessagesSendListService new MessagesSendListService
func NewMessagesSendListService(ctx context.Context) *MessagesSendListService {
	return &MessagesSendListService{ctx: ctx}
}

// Run create note info
func (s *MessagesSendListService) Run(req *message.MessagesListReq) (resp *message.MessagesSendListResp, err error) {
	// Finish your business logic.

	return
}
