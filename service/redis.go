package service

import (
	"fmt"
	"os"

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
	addr := os.Getenv("REDIS_URI")
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func CloseRedisClient(client *redis.Client) {
	errClose := client.Close()
	if errClose != nil {
		fmt.Println("Error close redis client:", errClose.Error())
	}
}
