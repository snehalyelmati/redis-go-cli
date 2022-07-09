package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/snehalyelmati/redis-go-cli/utils"
)

func main() {
	fmt.Println("Starting Redis client...")

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	count := 10

	// set some dummy data
	keys := utils.InsertRandomData(rdb, ctx, count)

	// get all the key-value pairs
	utils.PrintExistingData(rdb, ctx)

	// deleting dummy data
	utils.DeleteData(rdb, ctx, keys)

	// print all existing config
	utils.PrintExistingConfig(rdb, ctx)

	// TODO: use flags to modularize execution
}
