package convert

import (
	"gen/kitex_gen/system"
	"system/biz/dal/db/ent"
	"time"
)

func EntToApi(e *ent.API) *system.Api {
	return &system.Api{
		Id:          e.ID,
		CreatedAt:   e.CreatedAt.Format(time.DateOnly),
		UpdatedAt:   e.UpdatedAt.Format(time.DateOnly),
		Path:        e.Path,
		Description: e.Description,
		Group:       e.APIGroup,
		Method:      e.Method,
		Title:       e.Title,
	}
}
