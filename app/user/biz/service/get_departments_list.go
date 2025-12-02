package service

import (
	"context"
	Base "gen/kitex_gen/base"
	"gen/kitex_gen/user"
	"user/biz/convert"
	"user/biz/dal/db"
	"user/biz/dal/db/ent"
	"user/biz/dal/db/ent/department"
	"user/biz/dal/db/ent/predicate"
)

type GetDepartmentsListService struct {
	ctx context.Context
} // NewGetDepartmentsListService new GetDepartmentsListService
func NewGetDepartmentsListService(ctx context.Context) *GetDepartmentsListService {
	return &GetDepartmentsListService{ctx: ctx}
}

// Run create note info
func (s *GetDepartmentsListService) Run(req *user.GetDepartmentsListReq) (resp *user.DepartmentsListResp, err error) {
	// Finish your business logic.
	var (
		dataResp []*Base.Departments
	)
	var predicates []predicate.Department
	if req.GetKeyword() != "" {
		predicates = append(predicates, department.Or(
			department.NameContains(req.GetKeyword()),
		),
		)
	}

	all, err := db.Client.Department.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(department.FieldID)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		return nil, err
	}

	dataResp = convert.FindDepartmentsChildren(all, 0)

	return &user.DepartmentsListResp{
		Data: dataResp,
	}, nil
}
