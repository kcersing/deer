package db

import (
	"crm/biz/dal/db/ent"
	"crm/conf"
	"fmt"
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
