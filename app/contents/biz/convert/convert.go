package convert

import (
	"common/pkg/utils"
	"contents/biz/dal/db/ent"
	"gen/kitex_gen/base"
	"time"
)

func EntToArticle(e *ent.Article) *base.Article {
	if e == nil {
		return nil
	}

	mapper := utils.NewCopierMapper[base.Article, ent.Article]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)

	return dto

}
