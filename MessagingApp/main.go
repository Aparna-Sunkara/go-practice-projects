package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-11920.c81.us-east-1-2.ec2.redns.redis-cloud.com:11920",
		Username: "default",
		Password: "ugcVvpbaQGdFtQGDMjaJrV7Zjx3paZmx",
		DB:       0,
	})

	rdb.Set(ctx, "foo", "bar", 0)
	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	
	err = rdb.Set(ctx, "user:1", "John", 0).Err()
	if err != nil {
		panic(err)
	}

	
	val, err := rdb.Get(ctx, "user:1").Result()
	if err == nil {
		fmt.Println("Cached Value:", val)
	} else {
		fmt.Println("Cache miss!")
	}

	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("key%d", i)
		rdb.Set(ctx, key, i, 0)
	}
	fmt.Println("Inserted 30 keys. Redis will auto-evict using LRU.")
}
