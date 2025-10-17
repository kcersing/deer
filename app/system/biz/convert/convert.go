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

func EntToMenu(e *ent.Menu) *system.Menu {
	return &system.Menu{
		Id:        e.ID,
		Name:      e.Name,
		ParentId:  e.ParentID,
		Level:     e.Level,
		Path:      e.Path,
		Redirect:  e.Redirect,
		Component: e.Component,
		MenuType:  e.MenuType,
		Hidden:    e.Hidden,
		Sort:      e.Sort,
		Status:    e.Status,
		Url:       e.URL,
		//Children:  nil,
		CreatedAt: e.CreatedAt.Format(time.DateOnly),
		UpdatedAt: e.UpdatedAt.Format(time.DateOnly),
		Title:     e.Title,
		Type:      e.Type,
	}
}

func EntToRole(e *ent.Role) *system.Role {
	return &system.Role{
		Id:            e.ID,
		Name:          e.Name,
		Value:         e.Value,
		DefaultRouter: e.DefaultRouter,
		Remark:        e.Remark,
		//Apis:          e.Apis,
	}
}

func EntToDict(e *ent.Dict) *system.Dict {
	return &system.Dict{
		Id:          0,
		Title:       "",
		Name:        "",
		Status:      0,
		Description: "",
		CreatedAt:   "",
		UpdatedAt:   "",
	}
}
func EntToDictht(e *ent.Dictht) *system.Dictht {
	return &system.Dictht{
		Id:        0,
		Title:     "",
		Key:       "",
		Value:     "",
		Status:    0,
		CreatedAt: "",
		UpdatedAt: "",
		ParentID:  0,
	}
}
