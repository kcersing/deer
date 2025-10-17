package service

import (
	"context"
	system "gen/kitex_gen/system"
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

	return
}
