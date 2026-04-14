package db

import (
	"contents/biz/dal/db/ent"
	"contents/conf"
	"sync"
)

var onceClient sync.Once

var Client *ent.Client

func InitDB() {
	onceClient.Do(func() {
		Client = InItDB(conf.GetConf().PostgreSQL.DSN, true)
	})
}
