package dal

import (
	db "kcers-order/biz/dal/db/mysql"
	"kcers-order/biz/dal/db/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
