package utils

import (
	"fmt"
	"log"
	"github.com/oxygen-org/worker/config"
	"github.com/go-redis/redis"
)

var RedisC *redis.Client

func init() {
	c := config.C.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.PW, // no password set
		DB:       c.DB,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(pong, err)
	RedisC = client
}
