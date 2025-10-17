package db

import (
	"sync"
	"system/biz/dal/db/ent"
	"system/conf"
)

var onceClient sync.Once

var Client *ent.Client

func InitDB() {
	onceClient.Do(func() {
		Client = InItDB(conf.GetConf().MySQL.DSN, true)
	})
}
