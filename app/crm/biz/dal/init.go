package dal

import (
	db "deer/app/crm/biz/dal/mysql"
	"deer/app/crm/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
}
