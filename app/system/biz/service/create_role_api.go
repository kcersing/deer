package service

import (
	"context"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/pkg/errors"
	"strconv"
	"system/biz/dal/casbin"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/api"
	"system/biz/dal/db/ent/role"
)

type CreateRoleApiService struct {
	ctx context.Context
} // NewCreateRoleApiService new CreateRoleApiService
func NewCreateRoleApiService(ctx context.Context) *CreateRoleApiService {
	return &CreateRoleApiService{ctx: ctx}
}

// Run create note info
func (s *CreateRoleApiService) Run(req *system.CreateMenuAuthReq) (resp *base.NilResponse, err error) {
	var oldPolicies [][]string
	var roleId = strconv.FormatInt(req.GetRoleId(), 10)
	oldPolicies, _ = casbin.CasbinEnforcer.GetFilteredPolicy(0, roleId)
	if len(oldPolicies) != 0 {
		removeResult, err := casbin.CasbinEnforcer.RemoveFilteredPolicy(0, roleId)
		if err != nil {
			return nil, err
		}
		if !removeResult {
			return nil, errors.New("casbin policies remove failed")
		}
	}

	apiAll, err := db.Client.API.Query().Where(api.IDIn(req.GetIds()...)).All(s.ctx)
	// add new policies
	var policies [][]string
	for _, v := range apiAll {
		policies = append(policies, []string{roleId, v.Path, v.Method})
	}
	addResult, err := casbin.CasbinEnforcer.AddPolicies(policies)
	if err != nil {
		return nil, err
	}
	if !addResult {
		return nil, errors.New("casbin policies add failed")
	}

	jsonBytes, _ := json.Marshal(req.GetIds())
	var intSlice []int64
	err = json.Unmarshal(jsonBytes, &intSlice)
	if err != nil {
		return nil, err
	}

	_, err = db.Client.Role.Update().Where(role.ID(req.GetRoleId())).SetApis(intSlice).Save(s.ctx)
	if err != nil {
		return nil, err
	}
	return
}
