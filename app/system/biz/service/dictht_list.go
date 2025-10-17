package service

import (
	"context"
	system "gen/kitex_gen/system"
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

	return
}
