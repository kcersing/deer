package db

import (
	"fmt"
	"os"
	"sync"
	"user/biz/dal/db/ent"
	"user/conf"
)

var onceClient sync.Once

var Client *ent.Client

func InitDB() {
	onceClient.Do(func() {
		dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
		Client = InItDB(dsn, true)
	})
}
