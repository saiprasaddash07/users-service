package redis

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/saiprasaddash07/users-service/config"
)

var client *redis.Client
var once sync.Once

func NewRedisClient(clusterAdd string) *redis.Client {
	redisClusterClient := redis.NewClient(&redis.Options{
		Addr:     clusterAdd,
		Password: config.Get().RedisPassword,
	})
	return redisClusterClient
}

func Init() {
	once.Do(func() {
		clusterAdd := config.Get().RedisAddress
		client = NewRedisClient(clusterAdd)
		_, err := client.Ping().Result()
		if err != nil {
			panic(err.Error())
		}
	})
}

func GetClient() *redis.Client {
	return client
}
