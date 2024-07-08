package data

import (
	"context"
	slog "log"
	"os"
	"time"
	"user/internal/conf"
	"user/internal/data/ent"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB, NeweDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	edb *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, edb *ent.Client, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{edb: edb, rdb: rdb}, cleanup, nil
}

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {
	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			//IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, // 表名是否加 s
		},
	})

	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}

func NeweDB(c *conf.Data) *ent.Client {
    client, err := ent.Open(c.Database.Driver, c.Database.Source)
    if err != nil {
        log.Errorf("failed opening connection to %s: %v", c.Database.Driver, err)
        panic("failed to connect database")
    }

    // 执行数据库模式创建或迁移
    ctx := context.Background()
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    return client
}