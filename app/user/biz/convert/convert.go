package convert

import (
	"common/pkg/utils"
	Base "gen/kitex_gen/base"
	"github.com/jinzhu/copier"
	"time"
	"user/biz/dal/db/ent"
)

func EntToUser(e *ent.User) *Base.User {

	if e == nil {
		return nil
	}

	var dto Base.User

	if err := copier.CopyWithOption(&dto, &e, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		//copier.Copy(&dto, e)
		panic(err)
	}
	dto.Birthday = e.Birthday.Format(time.DateOnly)
	dto.LastAt = e.LastAt.Format(time.DateOnly)
	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return &dto
}

func EntToDepartments(e *ent.Department) *Base.Departments {
	mapper := utils.NewCopierMapper[Base.Departments, ent.Department]()
	var dto = mapper.ToDTO(e)
	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return dto
}
func EntToPosition(e *ent.Position) *Base.Positions {

	if e == nil {
		return nil
	}
	var dto Base.Positions
	if err := copier.CopyWithOption(&dto, &e, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		panic(err)
	}
	dto.CreatedAt = e.CreatedAt.Format(time.DateOnly)
	dto.UpdatedAt = e.UpdatedAt.Format(time.DateOnly)
	return &dto
}
