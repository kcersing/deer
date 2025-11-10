package service

import (
	"context"
	Base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/predicate"
)

type DicthtListService struct {
	ctx context.Context
} // NewDicthtListService new DicthtListService
func NewDicthtListService(ctx context.Context) *DicthtListService {
	return &DicthtListService{ctx: ctx}
}

// Run create note info
func (s *DicthtListService) Run(req *system.DicthtListReq) (resp *system.DicthtListResp, err error) {
	// Finish your business logic.
	var (
		predicates []predicate.Dictht
		dataResp   []*Base.Dictht
	)
	all, err := db.Client.Dictht.Query().Where(predicates...).All(s.ctx)
	if err != nil {
		return resp, err
	}
	for _, v := range all {
		dataResp = append(dataResp, convert.EntToDictht(v))
	}
	resp = &system.DicthtListResp{
		Data: dataResp,
	}
	return
}
