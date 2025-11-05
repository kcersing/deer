package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type CreatePropertyService struct {
	ctx context.Context
} // NewCreatePropertyService new CreatePropertyService
func NewCreatePropertyService(ctx context.Context) *CreatePropertyService {
	return &CreatePropertyService{ctx: ctx}
}

// Run create note info
func (s *CreatePropertyService) Run(req *product.CreatePropertyReq) (resp *product.PropertyResp, err error) {
	// Finish your business logic.

	return
}
