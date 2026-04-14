package service

import (
	"contents/biz/convert"
	"contents/biz/dal/db"
	"contents/biz/dal/db/ent"
	"contents/biz/dal/db/ent/article"
	"contents/biz/dal/db/ent/predicate"
	"context"
	Base "gen/kitex_gen/base"

	contents "gen/kitex_gen/contents"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
)

type ArticleListService struct {
	ctx context.Context
}

// NewArticleListService new ArticleListService
func NewArticleListService(ctx context.Context) *ArticleListService {
	return &ArticleListService{ctx: ctx}
}

// Run create note info
func (s *ArticleListService) Run(req *contents.ArticleListReq) (resp *contents.ArticleListResp, err error) {
	// Finish your business logic.

	var (
		dataResp []*Base.Article
	)
	var predicates []predicate.Article

	if req.GetKeyword() != "" {
		predicates = append(predicates, article.Or(
			article.TitleContains(req.GetKeyword()),
		))
	}
	if len(req.GetTagId()) > 0 {
		predicates = append(predicates, article.Or(
			func(s *sql.Selector) {
				var t []*sql.Predicate
				for _, v := range req.GetTagId() {
					t = append(t, sqljson.ValueContains(article.FieldTagID, v))
				}
				s.Where(sql.Or(t...))
			}))
	}
	predicates = append(predicates, article.Or(article.Delete(0)))

	all, err := db.Client.Debug().Article.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(article.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range all {

		dataResp = append(dataResp, convert.EntToArticle(v))
	}
	return &contents.ArticleListResp{
		Data: dataResp,
	}, nil
}
