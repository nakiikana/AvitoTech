package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Couldn't ping: %v", err)
		return
	}
	fmt.Println(ping)
}
