package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-11920.c81.us-east-1-2.ec2.redns.redis-cloud.com:11920",
		Username: "default",
		Password: "ugcVvpbaQGdFtQGDMjaJrV7Zjx3paZmx",
	})

	fmt.Println(" Starting publisher...")
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Notification #%d", i)
		err := rdb.Publish(ctx, "notifications", message).Err()
		if err != nil {
			fmt.Println("Error publishing:", err)
		} else {
			fmt.Println(" Published:", message)
		}
		time.Sleep(2 * time.Second)
	}
}


