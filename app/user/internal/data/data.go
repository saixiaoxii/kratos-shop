package data

import (
	"context"
	"kratos-shop/app/user/internal/conf"
	"kratos-shop/app/user/internal/data/ent"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NeweDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
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
