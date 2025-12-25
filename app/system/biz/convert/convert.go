package convert

import (
	"gen/kitex_gen/base"
	"strconv"
	"system/biz/dal/db/ent"
	"time"
)

func EntToApi(e *ent.API) *base.Api {
	return &base.Api{
		Id:        e.ID,
		CreatedAt: e.CreatedAt.Format(time.DateOnly),
		UpdatedAt: e.UpdatedAt.Format(time.DateOnly),
		Path:      e.Path,
		Desc:      e.Desc,
		Group:     e.Group,
		Method:    e.Method,
		Title:     e.Title,
		CreatedId: e.CreatedID,
	}
}
func EntToLog(e *ent.Logs) *base.Log {
	return &base.Log{
		Type:        e.Type,
		Method:      e.Method,
		Api:         e.API,
		Success:     e.Success,
		ReqContent:  e.ReqContent,
		RespContent: e.RespContent,
		Ip:          e.IP,
		UserAgent:   e.UserAgent,
		Time:        e.Time,
		CreatedAt:   e.CreatedAt.Format(time.DateTime),
		Identity:    e.Identity,
		Id:          e.ID,
	}
}

func EntToMenu(e *ent.Menu) *base.Menu {
	return &base.Menu{
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
		CreatedId: e.CreatedID,
	}
}

func EntToRole(e *ent.Role) *base.Role {
	return &base.Role{
		Id:        e.ID,
		Name:      e.Name,
		Code:      e.Code,
		Desc:      e.Desc,
		CreatedId: e.CreatedID,
		OrderNo:   e.OrderNo,
		Apis:      e.Apis,
		Menus:     e.Menus,
		Status:    e.Status,
	}
}

func EntToDict(e *ent.Dict) *base.Dict {
	return &base.Dict{
		Id:        e.ID,
		Title:     e.Title,
		Status:    e.Status,
		Desc:      e.Desc,
		Code:      e.Code,
		CreatedId: e.CreatedID,
		//CreatedName: e.CreatedName,
		CreatedAt: e.CreatedAt.Format(time.DateOnly),
		UpdatedAt: e.UpdatedAt.Format(time.DateOnly),
	}
}
func EntToDictht(e *ent.Dictht) *base.Dictht {
	return &base.Dictht{
		Id:        e.ID,
		Title:     e.Title,
		Value:     e.Value,
		Status:    e.Status,
		CreatedAt: e.CreatedAt.Format(time.DateOnly),
		UpdatedAt: e.UpdatedAt.Format(time.DateOnly),
		DictId:    e.DictID,
		CreatedId: e.CreatedID,
	}
}

func FindMenuChildren(data []*ent.Menu, parentID int64) []*base.Menu {
	if data == nil {
		return nil
	}
	var result []*base.Menu
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
