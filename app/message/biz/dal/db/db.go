package db

import (
	"sync"
	"message/biz/dal/db/ent"
	"message/conf"
)

var onceClient sync.Once

var Client *ent.Client

func InitDB() {
	onceClient.Do(func() {
		Client = InItDB(conf.GetConf().PostgreSQL.DSN, true)
	})
}
