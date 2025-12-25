package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent"
	"system/biz/dal/db/ent/api"
	"system/biz/dal/db/ent/role"
)

type GetRoleApiService struct {
	ctx context.Context
} // NewGetRoleApiService new GetRoleApiService
func NewGetRoleApiService(ctx context.Context) *GetRoleApiService {
	return &GetRoleApiService{ctx: ctx}
}

// Run create note info
func (s *GetRoleApiService) Run(req *base.IdReq) (resp *system.ApiListResp, err error) {
	var dataResp []*base.Api
	f, err := db.Client.Role.
		Query().
		Where(role.IDIn(req.GetId())).First(s.ctx)
	if err != nil {
		return nil, err
	}
	apis, err := db.Client.API.
		Query().
		Where(api.IDIn(f.Apis...)).
		Order(ent.Asc(api.FieldID)).
		All(s.ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range apis {
		dataResp = append(dataResp, convert.EntToApi(v))
	}
	resp = &system.ApiListResp{
		Data: dataResp,
	}
	return
}
