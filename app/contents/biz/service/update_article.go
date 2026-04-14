package service

import (
	"contents/biz/convert"
	"contents/biz/dal/db"
	"context"
	contents "gen/kitex_gen/contents"
)

type UpdateArticleService struct {
	ctx context.Context
}

// NewUpdateArticleService new UpdateArticleService
func NewUpdateArticleService(ctx context.Context) *UpdateArticleService {
	return &UpdateArticleService{ctx: ctx}
}

// Run create note info
func (s *UpdateArticleService) Run(req *contents.UpdateArticleReq) (resp *contents.ArticleResp, err error) {
	// Finish your business logic.

	entity, err := db.Client.Article.UpdateOneID(req.GetId()).
		SetTitle(req.Title).
		SetContent(req.Content).
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
