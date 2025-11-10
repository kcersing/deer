package service

import (
	"context"
	"gen/kitex_gen/base"
	system "gen/kitex_gen/system"
	"github.com/pkg/errors"
	"system/biz/convert"
	"system/biz/dal/db"
	"system/biz/dal/db/ent"
	"system/biz/dal/db/ent/dictht"
)

type UpdateDicthtService struct {
	ctx context.Context
} // NewUpdateDicthtService new UpdateDicthtService
func NewUpdateDicthtService(ctx context.Context) *UpdateDicthtService {
	return &UpdateDicthtService{ctx: ctx}
}

// Run create note info
func (s *UpdateDicthtService) Run(req *base.Dictht) (resp *system.DicthtResp, err error) {
	// Finish your business logic.

	_, err = db.Client.Dictht.Query().
		Where(dictht.ID(req.GetId())).
		WithDict(func(q *ent.DictQuery) {
		}).Only(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query DictionaryDetail failed")
	}
	// update dictionary detail
	save, err := db.Client.Dictht.UpdateOneID(req.GetId()).
		SetTitle(req.Title).
		SetKey(req.Key).
		SetValue(req.Value).
		SetStatus(req.Status).
		Save(s.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "update DictionaryDetail failed")
	}
	resp = &system.DicthtResp{
		Data: convert.EntToDictht(save),
	}
	return
}
