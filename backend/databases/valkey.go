package databases

import (
	"context"
	"ecommerce-project/config"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitValkey() {
	host := fmt.Sprintf("%s:%s", config.Cfg.ValkeyHost, config.Cfg.ValkeyPort)
	log.Println(host)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})

	if _, err := RedisClient.Ping(context.TODO()).Result(); err != nil {
		log.Fatalf("Valkey connection failed: %v", err)
	}
}
