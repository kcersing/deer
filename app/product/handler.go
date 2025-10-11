package main

import (
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// GetProductInfo implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductInfo(ctx context.Context, req *base.IdReq) (resp *product.ProductResp, err error) {
	// TODO: Your code here...
	return
}
