package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type KV struct {
	Key   string
	Value string
}

var Ctx = context.Background()
var Rdb *redis.Client

func RedisClient() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Yml.Redis.Addr,
		Password: Yml.Redis.Passwd, // no password set ""
		DB:       Yml.Redis.Db,     // default 0
	})

}
