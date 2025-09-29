package dal

import (
	"member/biz/dal/mysql"
	"member/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
