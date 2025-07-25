package db

import (
	"kcers-order/biz/dal/mysql"
	"kcers-order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
