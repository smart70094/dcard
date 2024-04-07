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
			Addr:     "localhost:6379", // 修改为您的 Redis 服务器地址
			Password: "123456",         // 如果您的 Redis 服务器有密码，请在此处填写密码
			DB:       0,                // 选择用于默认操作的数据库
		})
	})
	return client
}
