package db

import (
	"sync"
	"user/biz/dal/db/ent"
)

var onceClient sync.Once

var Client *ent.Client

func InitDB() {
	onceClient.Do(func() {
		//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
		//Client = InItDB(conf.GetConf().MySQL.DSN, true)
		Client = InItDB("root:root@tcp(127.0.0.1:3306)/deer_users?charset=utf8mb4&parseTime=True&loc=Local", true)

	})
}
