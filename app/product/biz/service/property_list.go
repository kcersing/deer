package service

import (
	"context"
	product "gen/kitex_gen/product"
)

type PropertyListService struct {
	ctx context.Context
} // NewPropertyListService new PropertyListService
func NewPropertyListService(ctx context.Context) *PropertyListService {
	return &PropertyListService{ctx: ctx}
}

// Run create note info
func (s *PropertyListService) Run(req *product.PropertyListReq) (resp *product.PropertyListResp, err error) {
	// Finish your business logic.

	return
}
