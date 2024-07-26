package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	channelName := "this is my channel"

	chs := rdb.Subscribe(ctx, channelName).Channel()
	println("begin to sub")
	for msg := range chs {
		log.Printf("Received message: %s from channel: %s", msg.Payload, msg.Channel)
	}
}
