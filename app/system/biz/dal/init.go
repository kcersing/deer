package dal

import (
	"system/biz/dal/casbin"
	"system/biz/dal/db"
	"system/biz/dal/redis"
)

func Init() {
	redis.Init()
	db.InitDB()
	casbin.InitCasbin()

}
