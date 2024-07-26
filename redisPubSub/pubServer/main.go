package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var ctx = context.Background()
var cnt int

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
	})
	channelName := "this is my channel"
	msg := "this is my message"
	ticker := time.NewTicker(time.Second * 3)
	t := time.NewTimer(time.Second * 60)
	println("begin pub...")
	for {
		select {
		case <-ticker.C:
			rdb.Publish(ctx, channelName, msg+strconv.Itoa(cnt))
			cnt++
		case <-t.C:
			println("超过1min，退出")
			return
		}
	}
}
