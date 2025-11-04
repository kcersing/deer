package service

import (
	"context"
)

type UpdateOpportunitiesService struct {
	ctx context.Context
} // NewUpdateOpportunitiesService new UpdateOpportunitiesService
func NewUpdateOpportunitiesService(ctx context.Context) *UpdateOpportunitiesService {
	return &UpdateOpportunitiesService{ctx: ctx}
}

// Run create note info
func (s *UpdateOpportunitiesService) Run(req *crm.UpdateOpportunitiesReq) (resp *crm.OpportunitiesResp, err error) {
	// Finish your business logic.

	return
}
