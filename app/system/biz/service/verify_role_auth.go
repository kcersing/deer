package service

import (
	"context"
	"fmt"
	base "gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"github.com/cloudwego/kitex/pkg/klog"
	"system/biz/dal/casbin"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/role"
	"time"
)

type VerifyRoleAuthService struct {
	ctx context.Context
} // NewVerifyRoleAuthService new VerifyRoleAuthService
func NewVerifyRoleAuthService(ctx context.Context) *VerifyRoleAuthService {
	return &VerifyRoleAuthService{ctx: ctx}
}

// Run create note info
func (s *VerifyRoleAuthService) Run(req *system.VerifyRoleAuthReq) (resp *system.VerifyRoleAuthResp, err error) {
	// Finish your business logic.

	//existToken := rpc.NewToken(ctx, c).IsExistByUserId(int64(id))
	//if !existToken {
	//	return false
	//}
	//check the role status
	roleInfo, err := db.Client.Role.Query().Where(role.IDEQ(req.GetRoleId())).Only(s.ctx)
	// if the role is not exist or the role is not active, return false
	if err != nil {
		resp = &system.VerifyRoleAuthResp{BaseResp: &base.BaseResp{
			Code:    1,
			Message: "role is not exist",
			Time:    time.Now().Format(time.DateTime),
		},
		}
		return
	}

	if roleInfo.Status != 1 {
		resp = &system.VerifyRoleAuthResp{BaseResp: &base.BaseResp{
			Code:    1,
			Message: "role cache is not a valid *ent.Role or the role is not active",
			Time:    time.Now().Format(time.DateTime),
		},
		}
		return
	}

	sub := req.GetRoleId()
	obj := req.GetObj()
	act := req.GetAct()
	//check the permission
	pass, err := casbin.CasbinEnforcer.Enforce(sub, obj, act)
	if err != nil {
		resp = &system.VerifyRoleAuthResp{BaseResp: &base.BaseResp{
			Code:    1,
			Message: fmt.Sprint("casbin err,  role id: ", sub, " path: ", obj, " method: ", act, " pass: ", pass, " err: ", err.Error()),
			Time:    time.Now().Format(time.DateTime),
		},
		}
		return
	}
	if !pass {
		klog.Info("casbin forbid role id: ", sub, " path: ", obj, " method: ", act, " pass: ", pass)
	}
	klog.Info("casbin allow role id: ", sub, " path: ", obj, " method: ", act, " pass: ", pass)
	resp = &system.VerifyRoleAuthResp{BaseResp: &base.BaseResp{
		Code: 0,
		Time: time.Now().Format(time.DateTime),
	},
	}
	return
}
