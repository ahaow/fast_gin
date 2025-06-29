package core

import (
	"fast_gin/config"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis(cfg *config.Config) *redis.Client {

	redisCfg := cfg.Redis

	if redisCfg.Address == "" {
		log.Fatalf("未配置redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.Db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s", err)
	}
	fmt.Println("连接redis成功")
	return client
}
