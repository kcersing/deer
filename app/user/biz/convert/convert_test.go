package convert

import (
	"testing"
	"time"
	"user/biz/dal/db/ent"
)

func TestEntToUser_Run(t *testing.T) {

	s := EntToUser(&ent.User{
		ID:           1,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		Delete:       1,
		CreatedID:    1,
		Status:       1,
		Username:     "test",
		Password:     "",
		Avatar:       "/1.jpg",
		Mobile:       "13999999999",
		Name:         "test",
		Gender:       1,
		Birthday:     time.Time{},
		DepartmentID: nil,
		PositionID:   nil,
		LastAt:       time.Time{},
		LastIP:       "127.0.0.1",
		Detail:       "简介",
	})

	t.Logf("resp: %v", s)

}
