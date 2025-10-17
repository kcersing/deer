package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type UpdateDicthtService struct {
	ctx context.Context
} // NewUpdateDicthtService new UpdateDicthtService
func NewUpdateDicthtService(ctx context.Context) *UpdateDicthtService {
	return &UpdateDicthtService{ctx: ctx}
}

// Run create note info
func (s *UpdateDicthtService) Run(req *system.Dictht) (resp *system.DictResp, err error) {
	// Finish your business logic.

	return
}
