package dal

import (
	db "deer/rpc/order/biz/dal/mysql"
	"deer/rpc/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
