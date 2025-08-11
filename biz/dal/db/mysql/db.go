package db

import (
	"kcers-order/biz/dal/db/mysql/ent"
	"kcers-order/conf"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		DB = InItDB(conf.GetConf().MySQL.DSN, true)
	})
}
