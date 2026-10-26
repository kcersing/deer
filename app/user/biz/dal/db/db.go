package db

import (
	"common/db"
	"context"
	"os"
	"sync"
	"user/biz/dal/db/ent"
	"user/biz/dal/db/ent/migrate"
	"user/conf"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	migrate2 "ariga.io/atlas/sql/migrate"
)

var onceClient sync.Once

var Client *ent.Client

func NewDBClient() {
	onceClient.Do(func() {
		Client = newDBClient(conf.GetConf().Database)
	})
}

// newDB
func newDBClient(databaseCfg conf.Database) (client *ent.Client) {

	drv := db.NewDB(databaseCfg)
	server.RegisterShutdownHook(func() {
		drv.Close()
	})

	drive := &db.SlowQueryDriver{
		Driver: drv,
		Log: func(ctx context.Context, a ...any) {
			klog.Warn(a)
		},
	}
	client = ent.NewClient(ent.Driver(drive))

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		klog.Fatalf("failed creating schema resources: %v", err)
	}

	// 添加慢查询日志 hook 和 interceptor
	//client.Use(SlowQueryHook())
	//client.Intercept(SlowQueryInterceptor())

	//创建本地迁移目
	migrationsPath := "./biz/dal/db/migrations"
	_, err := os.Stat(migrationsPath)
	if err != nil {
		if !os.IsNotExist(err) {
			klog.Fatalf("无法获取迁移路径的状态: %v", err)
			return nil
		}
		// Create the directory if it doesn't exist.
		err = os.MkdirAll(migrationsPath, os.ModePerm)
		if err != nil {
			klog.Fatalf("创建迁移路径失败: %v", err)
			return nil
		}
	}
	dir, err := migrate2.NewLocalDir(migrationsPath)
	if err != nil {
		klog.Fatalf("创建迁移目录失败: %v", err)
		return nil
	}
	// 编写迁移差异
	err = client.Schema.Diff(context.Background(), schema.WithDir(dir), schema.WithForeignKeys(false))
	if err != nil {
		klog.Fatalf("创建资源失败: %v", err)
		return nil
	}
	// 运行自动迁移工具
	if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
		klog.Fatalf("创建资源失败: %v", err)
		return nil
	}

	return client
}
