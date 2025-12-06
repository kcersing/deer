package dal

import (
	"member/biz/dal/db"
	"member/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
