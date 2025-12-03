package db

import (
	"deer/app/order/biz/dal/mysql/ent"
	"deer/app/order/conf"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		Client = InItDB(conf.GetConf().PostgreSQL.DSN, true)

	})
}
