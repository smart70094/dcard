package infra

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	client     *redis.Client
	clientOnce sync.Once
)

func GetRedisClient() *redis.Client {
	clientOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			//Addr: "127.0.0.1:6379", // 修改为您的 Redis 服务器地址
			Addr: "10.98.181.177:6379", // 修改为您的 Redis 服务器地址
			DB:   0,                    // 选择用于默认操作的数据库
		})
	})
	return client
}
