package service

import (
	"context"
	system "gen/kitex_gen/system"
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

	return
}
