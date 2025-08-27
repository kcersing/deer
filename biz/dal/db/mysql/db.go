package db

import (
	"deer/biz/dal/db/mysql/ent"
	"deer/conf"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		DB = InItDB(conf.GetConf().MySQL.DSN, true)
	})
}
