package convert

import (
	"common/pkg/utils"

	Base "gen/kitex_gen/base"
	"time"
	"user/biz/dal/db/ent"

	"github.com/jinzhu/copier"
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
	dto.PositionsId = *e.PositionID
	dto.DepartmentsId = *e.DepartmentID
	dto.Password = ""

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
func FindDepartmentsChildren(data []*ent.Department, parentID int64) []*Base.Departments {
	if data == nil {
		return nil
	}
	var result []*Base.Departments
	for _, v := range data {
		// discard the parent menu, only find the children menu
		if v.ParentID == parentID && v.ID != parentID {
			m := EntToDepartments(v)
			m.Children = FindDepartmentsChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}

func FindPositionChildren(data []*ent.Position, parentID int64) []*Base.Positions {
	if data == nil {
		return nil
	}
	var result []*Base.Positions
	for _, v := range data {
		// discard the parent menu, only find the children menu
		if v.ParentID == parentID && v.ID != parentID {
			m := EntToPosition(v)
			m.Children = FindPositionChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}
