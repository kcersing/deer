package main

import (
	"contents/biz/service"
	"context"
	base "gen/kitex_gen/base"
	contents "gen/kitex_gen/contents"
)

// ContentsServiceImpl implements the last service interface defined in the IDL.
type ContentsServiceImpl struct{}

// GetArticle implements the ContentsServiceImpl interface.
func (s *ContentsServiceImpl) GetArticle(ctx context.Context, req *base.IdReq) (resp *contents.ArticleResp, err error) {
	resp, err = service.NewGetArticleService(ctx).Run(req)

	return resp, err
}

// CreateArticle implements the ContentsServiceImpl interface.
func (s *ContentsServiceImpl) CreateArticle(ctx context.Context, req *contents.CreateArticleReq) (resp *contents.ArticleResp, err error) {
	resp, err = service.NewCreateArticleService(ctx).Run(req)

	return resp, err
}

// UpdateArticle implements the ContentsServiceImpl interface.
func (s *ContentsServiceImpl) UpdateArticle(ctx context.Context, req *contents.UpdateArticleReq) (resp *contents.ArticleResp, err error) {
	resp, err = service.NewUpdateArticleService(ctx).Run(req)

	return resp, err
}

// DeleteArticle implements the ContentsServiceImpl interface.
func (s *ContentsServiceImpl) DeleteArticle(ctx context.Context, req *base.IdReq) (resp *base.BaseResp, err error) {
	resp, err = service.NewDeleteArticleService(ctx).Run(req)

	return resp, err
}

// ArticleList implements the ContentsServiceImpl interface.
func (s *ContentsServiceImpl) ArticleList(ctx context.Context, req *contents.ArticleListReq) (resp *contents.ArticleListResp, err error) {
	resp, err = service.NewArticleListService(ctx).Run(req)

	return resp, err
}
