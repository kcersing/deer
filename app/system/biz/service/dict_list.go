package service

import (
	"context"
	"gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/predicate"
)

type DictListService struct {
	ctx context.Context
} // NewDictListService new DictListService
func NewDictListService(ctx context.Context) *DictListService {
	return &DictListService{ctx: ctx}
}

// Run create note info
func (s *DictListService) Run(req *system.DictListReq) (resp *system.DictListResp, err error) {
	// Finish your business logic.
	var (
		predicates []predicate.Dict
		dataResp   []*system.Dict
	)
	apis, err := db.Client.Dict.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return resp, err
	}
	for _, v := range apis {
		dataResp = append(dataResp, convert.EntToDict(v))
	}
	resp.Data = dataResp
	return
}
