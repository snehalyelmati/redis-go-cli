package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v9"
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
	keys := insertRandomData(rdb, ctx, count)

	// get all the key-value pairs
	printExistingData(rdb, ctx)

	// deleting dummy data
	deleteData(rdb, ctx, keys)
}

func getExistingKeysWithPattern(rdb *redis.Client, ctx context.Context, pattern string) []string {
	fmt.Println("Get all existing keys pairs")
	keys := []string{}

	iter := rdb.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}

func printExistingData(rdb *redis.Client, ctx context.Context) {
	fmt.Println("Print all key/value pairs:")

	keys := getExistingKeysWithPattern(rdb, ctx, "*")
	for _, v := range keys {
		fmt.Println(rdb.Get(ctx, v))
	}
	fmt.Println()
}

func insertRandomData(rdb *redis.Client, ctx context.Context, count int) []string {
	fmt.Println("Setting dummy keys...")
	keys := []string{}
	for i := 0; i < count; i++ {
		key, value := "someKey"+strconv.Itoa(i), "someValue"+strconv.Itoa(i)
		err := rdb.Set(ctx, key, value, 0).Err()
		keys = append(keys, key)

		if err != nil {
			fmt.Println("Error while setting")
			panic(err)
		}
	}
	// fmt.Println("Keys of the inserted elements:", keys)
	fmt.Println("Done.")
	fmt.Println()

	return keys
}

func deleteData(rdb *redis.Client, ctx context.Context, keys []string) {
	fmt.Println("Deleting specific keys...")
	err := rdb.Del(ctx, keys...).Err()
	if err != nil {
		fmt.Println("Error while setting")
		panic(err)
	}
	fmt.Println("Done.")
}

func deleteAllData(rdb *redis.Client, ctx context.Context) {
	fmt.Println("Deleting all keys...")

	err := rdb.Del(ctx, getExistingKeysWithPattern(rdb, ctx, "*")...).Err()
	if err != nil {
		fmt.Println("Error while setting")
		panic(err)
	}
	fmt.Println("Done.")
}
