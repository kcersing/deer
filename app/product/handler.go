package main

import (
	"context"
	base "gen/kitex_gen/base"
	product "gen/kitex_gen/product"
	"product/biz/service"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// CreateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (resp *product.ProductResp, err error) {
	resp, err = service.NewCreateProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *product.EditProductReq) (resp *product.ProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// OnlineProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) OnlineProduct(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewOnlineProductService(ctx).Run(req)

	return resp, err
}

// OfflineProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) OfflineProduct(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewOfflineProductService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *base.IdReq) (resp *product.ProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) SearchProduct(ctx context.Context, req *product.SearchProductReq) (resp *product.ProductListResp, err error) {
	resp, err = service.NewSearchProductService(ctx).Run(req)

	return resp, err
}

// ProductList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ProductList(ctx context.Context, req *product.ListReq) (resp *product.ProductListResp, err error) {
	resp, err = service.NewProductListService(ctx).Run(req)

	return resp, err
}

// DecrStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DecrStock(ctx context.Context, req *product.DecrStockReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDecrStockService(ctx).Run(req)

	return resp, err
}

// DecrStockRevert implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DecrStockRevert(ctx context.Context, req *product.DecrStockReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDecrStockRevertService(ctx).Run(req)

	return resp, err
}

// CreateProperty implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateProperty(ctx context.Context, req *product.CreatePropertyReq) (resp *product.PropertyResp, err error) {
	resp, err = service.NewCreatePropertyService(ctx).Run(req)

	return resp, err
}

// UpdateProperty implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateProperty(ctx context.Context, req *product.UpdatePropertyReq) (resp *product.PropertyResp, err error) {
	resp, err = service.NewUpdatePropertyService(ctx).Run(req)

	return resp, err
}

// DeleteProperty implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteProperty(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeletePropertyService(ctx).Run(req)

	return resp, err
}

// PropertyList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) PropertyList(ctx context.Context, req *product.PropertyListReq) (resp *product.PropertyListResp, err error) {
	resp, err = service.NewPropertyListService(ctx).Run(req)

	return resp, err
}
