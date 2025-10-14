package dal

import (
	"user/biz/dal/db"
	"user/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
