package service

import (
	"context"
	system "gen/kitex_gen/system"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/predicate"
)

type LogListService struct {
	ctx context.Context
} // NewLogListService new LogListService
func NewLogListService(ctx context.Context) *LogListService {
	return &LogListService{ctx: ctx}
}

// Run create note info
func (s *LogListService) Run(req *system.LogListReq) (resp *system.LogListResp, err error) {
	// Finish your business logic.
	var (
		predicates []predicate.Logs
		dataResp   []*system.Log
	)
	apis, err := db.Client.Logs.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return resp, err
	}
	for _, v := range apis {
		dataResp = append(dataResp, convert.EntToLog(v))
	}
	resp.Data = dataResp

	return
}
