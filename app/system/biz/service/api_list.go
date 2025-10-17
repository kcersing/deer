package service

import (
	"context"
	"gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/predicate"
)

type ApiListService struct {
	ctx context.Context
} // NewApiListService new ApiListService
func NewApiListService(ctx context.Context) *ApiListService {
	return &ApiListService{ctx: ctx}
}

// Run create note info
func (s *ApiListService) Run(req *system.ApiListReq) (resp *system.ApiListResp, err error) {
	// Finish your business logic.
	var (
		predicates []predicate.API
		ApiResp    []*system.Api
	)
	apis, err := db.Client.API.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return resp, err
	}
	for _, v := range apis {
		ApiResp = append(ApiResp, convert.EntToApi(v))
	}
	resp.Data = ApiResp
	return
}
