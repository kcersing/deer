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
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.ProductResp, err error) {
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

// CreateItem implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) CreateItem(ctx context.Context, req *product.CreateItemReq) (resp *product.ItemResp, err error) {
	resp, err = service.NewCreateItemService(ctx).Run(req)

	return resp, err
}

// UpdateItem implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateItem(ctx context.Context, req *product.UpdateItemReq) (resp *product.ItemResp, err error) {
	resp, err = service.NewUpdateItemService(ctx).Run(req)

	return resp, err
}

// DeleteItem implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteItem(ctx context.Context, req *base.IdReq) (resp *base.NilResponse, err error) {
	resp, err = service.NewDeleteItemService(ctx).Run(req)

	return resp, err
}

// ItemList implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ItemList(ctx context.Context, req *product.ItemListReq) (resp *product.ItemListResp, err error) {
	resp, err = service.NewItemListService(ctx).Run(req)

	return resp, err
}

// GetItem implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetItem(ctx context.Context, req *base.IdReq) (resp *product.ItemResp, err error) {
	resp, err = service.NewGetItemService(ctx).Run(req)

	return resp, err
}
