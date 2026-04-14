package service

import (
	"contents/biz/convert"
	"contents/biz/dal/db"
	"context"
	contents "gen/kitex_gen/contents"
)

type CreateArticleService struct {
	ctx context.Context
}

// NewCreateArticleService new CreateArticleService
func NewCreateArticleService(ctx context.Context) *CreateArticleService {
	return &CreateArticleService{ctx: ctx}
}

// Run create note info
func (s *CreateArticleService) Run(req *contents.CreateArticleReq) (resp *contents.ArticleResp, err error) {
	// Finish your business logic.
	entity, err := db.Client.Article.Create().
		SetTitle(req.Title).
		SetContent(req.Content).
		SetCreatedID(req.CreatedId).
		SetPic(req.Pic).
		SetTagID(req.TagId).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	dataResp := convert.EntToArticle(entity)
	resp = &contents.ArticleResp{
		Data: dataResp,
	}
	return
}
