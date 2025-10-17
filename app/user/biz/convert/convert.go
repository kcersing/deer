package convert

import (
	User "gen/kitex_gen/user"
	"time"
	"user/biz/dal/db/ent"
)

func EntToUser(e *ent.User) *User.User {
	return &User.User{
		Id:       e.ID,
		Avatar:   e.Avatar,
		Mobile:   e.Mobile,
		Name:     e.Name,
		Status:   e.Status,
		Gender:   e.Gender,
		Birthday: e.Birthday.Format(time.DateOnly),
		LastAt:   e.LastAt.Format(time.DateOnly),
		LastIp:   e.LastIP,
		Detail:   e.Detail,
		//Roles:     e.Roles,
		CreatedAt: e.CreatedAt.Format(time.DateOnly),
		UpdatedAt: e.UpdatedAt.Format(time.DateOnly),
		CreatedId: e.CreatedID,
	}
}
