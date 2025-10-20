package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent"
	"system/biz/dal/db/ent/api"
	"system/biz/dal/db/ent/menu"
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
	var dataResp []*system.Api
	all, err := db.Client.Role.
		Query().
		Where(role.IDIn(req.GetId())).
		QueryAPI().
		Where(api.DisabledEQ(0)).
		Order(ent.Asc(menu.FieldSort)).
		All(s.ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range all {
		dataResp = append(dataResp, convert.EntToApi(v))
	}
	resp.Data = dataResp
	return
}
