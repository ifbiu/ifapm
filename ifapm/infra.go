package ifapm

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

type infra struct {
	Db  *sql.DB
	Rdb *redis.Client
}

var Infra = &infra{}

type InfraOption func(*infra)

func InfraMysqlDbOption(connectUrl string) InfraOption {
	return func(i *infra) {
		var err error
		i.Db, err = sql.Open("mysql", connectUrl)
		if err != nil {
			panic(err)
		}
		err = i.Db.Ping()
		if err != nil {
			panic(err)
		}
	}
}

func InfraRdbOption(connectUrl string) InfraOption {
	return func(i *infra) {
		rdb := redis.NewClient(&redis.Options{
			Addr: connectUrl,
			DB:   0,
		})
		res, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
		if res != "PONG" {
			panic("redis ping fail")
		}
		i.Rdb = rdb
	}
}

func (i *infra) Init(options ...InfraOption) {
	for _, option := range options {
		option(i)
	}
}
