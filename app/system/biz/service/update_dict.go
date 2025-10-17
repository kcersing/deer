package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type UpdateDictService struct {
	ctx context.Context
} // NewUpdateDictService new UpdateDictService
func NewUpdateDictService(ctx context.Context) *UpdateDictService {
	return &UpdateDictService{ctx: ctx}
}

// Run create note info
func (s *UpdateDictService) Run(req *system.Dict) (resp *system.DictResp, err error) {
	// Finish your business logic.

	return
}
