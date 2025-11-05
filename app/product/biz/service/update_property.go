package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type UpdatePropertyService struct {
	ctx context.Context
} // NewUpdatePropertyService new UpdatePropertyService
func NewUpdatePropertyService(ctx context.Context) *UpdatePropertyService {
	return &UpdatePropertyService{ctx: ctx}
}

// Run create note info
func (s *UpdatePropertyService) Run(req *product.UpdatePropertyReq) (resp *product.PropertyResp, err error) {
	// Finish your business logic.

	return
}
