package service

import (
	"context"
	system "gen/kitex_gen/system"
)

type CreateDicthtService struct {
	ctx context.Context
} // NewCreateDicthtService new CreateDicthtService
func NewCreateDicthtService(ctx context.Context) *CreateDicthtService {
	return &CreateDicthtService{ctx: ctx}
}

// Run create note info
func (s *CreateDicthtService) Run(req *system.Dictht) (resp *system.DictResp, err error) {
	// Finish your business logic.

	return
}
