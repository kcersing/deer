package dal

import (
	"product/biz/dal/db"
	"product/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
