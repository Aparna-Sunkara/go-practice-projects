package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-11920.c81.us-east-1-2.ec2.redns.redis-cloud.com:11920",
		Username: "default",
		Password: "ugcVvpbaQGdFtQGDMjaJrV7Zjx3paZmx",
	})

	sub := rdb.Subscribe(ctx, "notifications")

	fmt.Println("Subscribed to 'notifications' channel...")
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("Error receiving message:", err)
			continue
		}
		fmt.Printf("Received message on channel %s: %s\n", msg.Channel, msg.Payload)
	}
}
