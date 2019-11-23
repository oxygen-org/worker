package main

import (
	"fmt"
	"github.com/oxygen-org/worker/config"
	"github.com/oxygen-org/worker/utils"
	"time"

	"github.com/go-redis/redis"
)

var identify = fmt.Sprintf("%s#%s", config.C.General.HostIP,
	config.C.General.HostName)

// BeatPing 上报自己依旧存活
func BeatPing() {
	key := "heart:worker"
	key = fmt.Sprintf("%s:%s", key, identify)
	go func() {
		expire := time.Second * 10
		//注意，只有在整个生命周期都需要才这样使用time.Tick
		limiter := time.Tick(time.Second * 6)
		for {
			utils.RedisC.Set(key, time.Now().UnixNano(), expire)
			<-limiter
		}

		//如有必要停止Tick
		// ticker := time.NewTicker(10 *time.Second)
		// <- ticker.C
		// ticker.Stop()
	}()
}

// RegisterMe 注册自己
func RegisterMe() {
	key := "register:worker"
	timestamp := time.Now().UnixNano()

	mem := redis.Z{
		Score:  float64(timestamp),
		Member: identify,
	}
	utils.RedisC.ZAdd(key, mem)
}
