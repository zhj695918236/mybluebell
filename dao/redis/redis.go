package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"mybluebell/setting"
)

var (
	Client *redis.Client
	Nil    = redis.Nil
)

// Init初始化redis连接

func Init(cfg *setting.RedisConfig) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = Client.Close()
}
