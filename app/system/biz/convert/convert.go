package convert

import (
	"gen/kitex_gen/base"
	"gen/kitex_gen/system"
	"strconv"
	"system/biz/dal/db/ent"
	"time"
)

func EntToApi(e *ent.API) *system.Api {
	return &system.Api{
		Id:        e.ID,
		CreatedAt: e.CreatedAt.Format(time.DateOnly),
		UpdatedAt: e.UpdatedAt.Format(time.DateOnly),
		Path:      e.Path,
		Desc:      e.Desc,
		Group:     e.APIGroup,
		Method:    e.Method,
		Title:     e.Title,
	}
}
func EntToLog(e *ent.Logs) *system.Log {
	return &system.Log{
		Type:        e.Type,
		Method:      e.Method,
		Api:         e.API,
		Success:     false,
		ReqContent:  e.ReqContent,
		RespContent: e.RespContent,
		Ip:          e.IP,
		UserAgent:   e.UserAgent,
		Operatorsr:  e.Operatorsr,
		Time:        e.Time,
		CreatedAt:   e.CreatedAt.Format(time.DateTime),
		UpdatedAt:   "",
		Identity:    0,
		Id:          0,
	}
}

func EntToMenu(e *ent.Menu) *system.Menu {
	return &system.Menu{
		Id:        e.ID,
		Icon:      e.Icon,
		Name:      e.Name,
		ParentId:  e.ParentID,
		Level:     e.Level,
		Path:      e.Path,
		Redirect:  e.Redirect,
		Component: e.Component,
		MenuType:  e.MenuType,
		Status:    e.Status,
		CreatedAt: e.CreatedAt.Format(time.DateOnly),
		UpdatedAt: e.UpdatedAt.Format(time.DateOnly),
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
		Id:        0,
		Title:     "",
		Name:      "",
		Status:    0,
		Desc:      "",
		CreatedAt: "",
		UpdatedAt: "",
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

func FindMenuChildren(data []*ent.Menu, parentID int64) []*system.Menu {
	if data == nil {
		return nil
	}
	var result []*system.Menu
	for _, v := range data {
		// discard the parent menu, only find the children menu
		if v.ParentID == parentID && v.ID != parentID {
			m := EntToMenu(v)
			m.Children = FindMenuChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}

func FindMenuTreeChildren(data []*ent.Menu, parentID int64) []*base.Tree {
	if data == nil {
		return nil
	}
	var result []*base.Tree
	for _, v := range data {
		if v.ParentID == parentID && v.ID != parentID {
			var m = new(base.Tree)
			m.Title = v.Name
			m.Value = strconv.FormatInt(v.ID, 10)
			m.Key = strconv.FormatInt(v.ID, 10)
			m.Children = FindMenuTreeChildren(data, v.ID)
			result = append(result, m)
		}
	}
	return result
}
