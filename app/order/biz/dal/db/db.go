package db

import (
	"order/biz/dal/db/ent"
	"order/conf"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var onceClient sync.Once

var Client *ent.Client

func InitDB() {
	onceClient.Do(func() {
		//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
		klog.Info(conf.GetConf().PostgreSQL.DSN)
		Client = InItDB(conf.GetConf().PostgreSQL.DSN, true)

	})
}
