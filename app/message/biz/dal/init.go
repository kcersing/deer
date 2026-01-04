package dal

import (
	"message/biz/dal/db"
	"message/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
