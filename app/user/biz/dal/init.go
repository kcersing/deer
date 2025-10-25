package dal

import (
	"user/biz/dal/db"
)

func Init() {
	//redis.Init()
	db.InitDB()
}
