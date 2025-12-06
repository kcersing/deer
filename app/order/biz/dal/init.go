package dal

import (
	"order/biz/dal/db"
	"order/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()

}
