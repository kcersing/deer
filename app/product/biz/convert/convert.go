package convert

import (
	"common/pkg/utils"
	"gen/kitex_gen/base"
	"product/biz/dal/db/ent"
	"time"
)

func EntToItem(e *ent.ProductItem) *base.Item {

	mapper := utils.NewCopierMapper[base.Item, ent.ProductItem]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return dto
}
func EntToProduct(e *ent.Product) *base.Product {
	mapper := utils.NewCopierMapper[base.Product, ent.Product]()
	var dto = mapper.ToDTO(e)

	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return dto
}
