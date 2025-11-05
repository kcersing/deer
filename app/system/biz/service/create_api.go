package service

import (
	"context"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
)

type CreateApiService struct {
	ctx context.Context
} // NewCreateApiService new CreateApiService
func NewCreateApiService(ctx context.Context) *CreateApiService {
	return &CreateApiService{ctx: ctx}
}

// Run create note info
func (s *CreateApiService) Run(req *system.CreateApiReq) (resp *system.ApiResp, err error) {
	save, err := db.Client.API.Create().
		SetTitle(req.GetTitle()).
		SetDesc(req.GetDesc()).
		SetMethod(req.GetMethod()).
		SetPath(req.GetPath()).
		SetAPIGroup(req.GetGroup()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp.Data = convert.EntToApi(save)
	return
}
