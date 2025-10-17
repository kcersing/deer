package service

import (
	"context"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/api"
)

type UpdateApiService struct {
	ctx context.Context
} // NewUpdateApiService new UpdateApiService
func NewUpdateApiService(ctx context.Context) *UpdateApiService {
	return &UpdateApiService{ctx: ctx}
}

// Run create note info
func (s *UpdateApiService) Run(req *system.UpdateApiReq) (resp *system.ApiResp, err error) {

	_, err = db.Client.API.Update().
		Where(api.IDEQ(req.GetId())).
		SetTitle(req.GetTitle()).
		SetDescription(req.GetDescription()).
		SetMethod(req.GetMethod()).
		SetPath(req.GetPath()).
		SetAPIGroup(req.GetGroup()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	only, err := db.Client.API.Query().
		Where(api.IDEQ(req.GetId())).Only(s.ctx)
	if err != nil {
		return nil, err
	}
	resp.Data = convert.EntToApi(only)

	return
}
