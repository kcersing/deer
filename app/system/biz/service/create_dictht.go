package service

import (
	"context"
	"fmt"
	system "gen/kitex_gen/system"
	"github.com/pkg/errors"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/dict"
	"system/biz/dal/db/ent/dictht"
)

type CreateDicthtService struct {
	ctx context.Context
} // NewCreateDicthtService new CreateDicthtService
func NewCreateDicthtService(ctx context.Context) *CreateDicthtService {
	return &CreateDicthtService{ctx: ctx}
}

// Run create note info
func (s *CreateDicthtService) Run(req *system.Dictht) (resp *system.DictResp, err error) {
	// Finish your business logic.
	exist, err := db.Client.Dictht.Query().
		Where(dictht.Key(req.GetTitle())).
		Where(dictht.HasDictWith(dict.ID(req.GetParentID()))).
		Exist(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "查询字典详细信息失败")
	}
	if exist {
		return nil, errors.New("字典详细信息已存在")
	}

	// find dictionary by id
	dict, err := db.Client.Dict.Query().Where(dict.ID(req.GetParentID())).Only(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "查询字典信息失败")
	}
	if dict == nil {
		return nil, errors.New(fmt.Sprintf("找不到词典，请检查词典ID, %d", req.ParentID))
	}

	// create dictionary detail
	save := db.Client.Dictht.Create().
		SetDict(dict). // set parent dictionary
		SetTitle(req.GetTitle()).
		SetKey(req.GetKey()).
		SetValue(req.GetValue())

	//save.SetCreatedID(userToken.ID)

	_, err = save.Save(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "创建字典信息失败")
	}

	return
}
