package service

import (
	"context"
	"gen/kitex_gen/crm"
)

type CreateOpportunitiesService struct {
	ctx context.Context
} // NewCreateOpportunitiesService new CreateOpportunitiesService
func NewCreateOpportunitiesService(ctx context.Context) *CreateOpportunitiesService {
	return &CreateOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *CreateOpportunitiesService) Run(req *crm.CreateOpportunitiesReq) (resp *crm.OpportunitiesResp, err error) {
	// Finish your business logic.

	return
}
