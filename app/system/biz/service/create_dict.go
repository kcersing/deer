package service

import (
	"context"
	"errors"
	system "gen/kitex_gen/system"
	"system/biz/convert"

	"system/biz/dal/db"
	"system/biz/dal/db/ent/dict"
)

type CreateDictService struct {
	ctx context.Context
} // NewCreateDictService new CreateDictService
func NewCreateDictService(ctx context.Context) *CreateDictService {
	return &CreateDictService{ctx: ctx}
}

// Run create note info
func (s *CreateDictService) Run(req *system.Dict) (resp *system.DictResp, err error) {
	// Finish your business logic.

	dictionaryExist, _ := db.Client.Dict.Query().Where(dict.Name(req.GetName())).Exist(s.ctx)
	if dictionaryExist {
		return nil, errors.New("dict name already exists")
	}
	// create dictionary
	save, err := db.Client.Dict.Create().
		SetTitle(req.GetTitle()).
		SetName(req.GetName()).
		SetDesc(req.GetDesc()).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}
	resp = &system.DictResp{
		Data: convert.EntToDict(save),
	}
	return
}
