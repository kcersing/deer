package service

import (
	"context"
	system "gen/kitex_gen/system"
	"github.com/pkg/errors"
	"system/biz/dal/db"
	"system/biz/dal/db/ent/dict"
)

type UpdateDictService struct {
	ctx context.Context
} // NewUpdateDictService new UpdateDictService
func NewUpdateDictService(ctx context.Context) *UpdateDictService {
	return &UpdateDictService{ctx: ctx}
}

// Run create note info
func (s *UpdateDictService) Run(req *system.Dict) (resp *system.DictResp, err error) {
	// Finish your business logic.
	dictionaryExist, _ := db.Client.Dict.Query().Where(dict.ID(req.GetId())).Exist(s.ctx)
	if !dictionaryExist {
		return nil, errors.New("The dictionary try to update is not exists")
	}
	// update dictionary
	_, err = db.Client.Dict.UpdateOneID(req.GetId()).
		SetTitle(req.Title).
		SetName(req.Name).
		SetStatus(req.Status).
		SetDescription(req.Description).
		Save(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "update Dictionary failed")
	}
	return
}
