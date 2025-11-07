package convert

import (
	Member "gen/kitex_gen/member"
	"user/biz/dal/db/ent"
)

func EntToMember(e *ent.Member) *Member.Member {
	return &Member.Member{
		Id:        0,
		Username:  "",
		Password:  "",
		Avatar:    "",
		Mobile:    "",
		Name:      "",
		Status:    0,
		Level:     0,
		Gender:    0,
		Birthday:  "",
		LastAt:    "",
		LastIp:    "",
		CreatedAt: "",
		UpdatedAt: "",
		CreatedId: 0,
	}
}
