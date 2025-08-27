package dal

import (
	db "deer/biz/dal/db/mysql"
	"deer/biz/dal/db/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
