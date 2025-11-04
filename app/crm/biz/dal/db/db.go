package db

import (
	"deer/app/crm/biz/dal/mysql/ent"
	"deer/app/crm/conf"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		DB = InItDB(conf.GetConf().MySQL.DSN, true)
	})
}
