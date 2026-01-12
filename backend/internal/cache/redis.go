package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var redisClient *redis.Client

func init() {
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
}

func InitializeRedis() {
	// 初始化redis连接池
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func CloseRedis() {
	if err := redisClient.Close(); err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func Set(key string, value any, expiration time.Duration) error {
	return redisClient.Set(context.Background(), key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return redisClient.Get(context.Background(), key).Result()
}

func Del(key string) error {
	return redisClient.Del(context.Background(), key).Err()
}
