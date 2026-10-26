package db

import (
	"database/sql"
	"user/conf"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	"log"
	"time"
)

func NewDB(databaseCfg conf.Database) (drv *entsql.Driver) {
	switch databaseCfg.Type {
	case "mysql":
		drv = OpenMySql(databaseCfg.DSN)
	case "postgres":
		drv = OpenPq(databaseCfg.DSN)
	}
	return
}
func OpenMySql(databaseUrl string) *entsql.Driver {
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 从db变量中构造一个ent.Driver对象。
	drv := entsql.OpenDB(dialect.MySQL, db)

	return drv
}

func OpenPq(databaseUrl string) *entsql.Driver {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	return drv
}
