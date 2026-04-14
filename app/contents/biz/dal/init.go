package dal

import (
	"contents/biz/dal/db"
	"contents/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()

}
