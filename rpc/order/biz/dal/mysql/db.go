package db

import (
	"deer/rpc/order/biz/dal/mysql/ent"
	"deer/rpc/order/conf"
	"sync"
)

var onceClient sync.Once

var DB *ent.Client

func InitDB() {
	onceClient.Do(func() {
		DB = InItDB(conf.GetConf().MySQL.DSN, true)
	})
}
