package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"message/biz/dal/db"
	"message/biz/dal/db/ent"
	"message/biz/dal/db/ent/messages"

	message "gen/kitex_gen/message"
	"message/biz/convert"

	"message/biz/dal/db/ent/predicate"

	"github.com/cloudwego/kitex/pkg/klog"
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
	klog.Info("GetMemberListService.Run req: %v", req)
	var (
		dataResp []*Base.Messages
	)
	var predicates []predicate.Messages
	//if req.GetKeyword() != "" {
	//	predicates = append(predicates, messages.Or(
	//		messages.ContentContains(req.GetKeyword()),
	//	))
	//}

	all, err := db.Client.Messages.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(messages.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range all {
		dataResp = append(dataResp, convert.EntToMessages(v))
	}
	return &message.MessagesListResp{
		Data: dataResp,
	}, nil
}
