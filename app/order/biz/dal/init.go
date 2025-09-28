package dal

import (
	db "deer/app/order/biz/dal/mysql"
	"deer/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
