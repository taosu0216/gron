package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"gron/internal/conf"
)

func NewCache(conf *conf.Data) *redis.Client {
	addr := conf.Redis.Addr
	redisConn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Db),
		PoolSize: int(conf.Redis.PoolSize),
	})
	if redisConn == nil {
		log.Fatal("failed to call redis.NewClient")
	}
	return redisConn
}
