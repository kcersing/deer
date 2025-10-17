package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type CreateDictService struct {
	ctx context.Context
} // NewCreateDictService new CreateDictService
func NewCreateDictService(ctx context.Context) *CreateDictService {
	return &CreateDictService{ctx: ctx}
}

// Run create note info
func (s *CreateDictService) Run(req *system.Dict) (resp *system.DictResp, err error) {
	// Finish your business logic.

	return
}
