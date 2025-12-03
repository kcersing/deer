package db

import (
	"fmt"
	"member/biz/dal/db/ent"
	"member/conf"
	"os"
	"sync"
)

var onceClient sync.Once

var Client *ent.Client

func InitDB() {
	onceClient.Do(func() {
		Client = InItDB(conf.GetConf().PostgreSQL.DSN, true)

	})
}
