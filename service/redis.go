package service

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ping Redis server to check the connection
	_, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("Error redis connection ", err.Error())
	}
}

func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
