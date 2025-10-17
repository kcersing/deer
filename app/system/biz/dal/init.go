package dal

import (
	"system/biz/dal/db"
	"system/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
