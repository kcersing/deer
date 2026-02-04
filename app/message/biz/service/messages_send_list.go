package service

import (
	"context"
	Base "gen/kitex_gen/base"
	message "gen/kitex_gen/message"
	"message/biz/convert"
	"message/biz/dal/db"
	"message/biz/dal/db/ent"
	"message/biz/dal/db/ent/messages"
	"message/biz/dal/db/ent/predicate"
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

	var (
		dataResp []*Base.MessagesSend
	)
	var predicates []predicate.MessagesSentRecords
	//if req.GetKeyword() != "" {
	//	predicates = append(predicates, messages.Or(
	//		messages.ContentContains(req.GetKeyword()),
	//	))
	//}

	all, err := db.Client.MessagesSentRecords.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(messages.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range all {

		dataResp = append(dataResp, convert.EntToMessagesSentRecords(v))
	}
	return &message.MessagesSendListResp{
		Data: dataResp,
	}, nil
}
