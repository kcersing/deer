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
	"message/biz/dal/db/ent/smssentrecords"
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

	var (
		dataResp []*Base.SmsSend
	)
	var predicates []predicate.SmsSentRecords
	if req.GetMobile() != "" {
		predicates = append(predicates, smssentrecords.Or(
			smssentrecords.MobileContains(req.GetMobile()),
		))
	}

	all, err := db.Client.SmsSentRecords.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(messages.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range all {
		dataResp = append(dataResp, convert.EntToSmsSentRecords(v))
	}
	return &message.SmsSendListResp{
		Data: dataResp,
	}, nil
}
