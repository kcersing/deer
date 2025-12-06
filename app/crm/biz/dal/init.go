package dal

import (
	"crm/biz/dal/db"
	"crm/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()

}
