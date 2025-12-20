package service

import (
	"context"
	"gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/dal/db"

	"github.com/pkg/errors"
)

type CreateLogService struct {
	ctx context.Context
} // NewCreateLogService new CreateLogService
func NewCreateLogService(ctx context.Context) *CreateLogService {
	return &CreateLogService{ctx: ctx}
}

// Run create note info
func (s *CreateLogService) Run(req *system.CreateLogReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.
	_, err = db.Client.Logs.Create().
		SetType(req.GetType()).
		SetMethod(req.GetMethod()).
		SetAPI(req.GetApi()).
		SetSuccess(req.GetSuccess()).
		SetReqContent(req.GetReqContent()).
		SetRespContent(req.GetRespContent()).
		SetIP(req.GetIp()).
		SetUserAgent(req.GetUserAgent()).
		SetIdentity(req.GetIdentity()).
		SetTime(req.GetTime()).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Log failed")
		return nil, err
	}

	return
}
