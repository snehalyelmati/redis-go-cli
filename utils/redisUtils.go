package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func GetExistingKeysWithPattern(rdb *redis.Client, ctx context.Context, pattern string) []string {
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

func PrintExistingData(rdb *redis.Client, ctx context.Context) {
	fmt.Println("Print all key/value pairs:")

	keys := GetExistingKeysWithPattern(rdb, ctx, "*")
	for _, v := range keys {
		fmt.Println(rdb.Get(ctx, v))
	}
	fmt.Println()
}

func InsertRandomData(rdb *redis.Client, ctx context.Context, count int) []string {
	fmt.Println("Setting dummy keys...")
	keys := []string{}
	for i := 0; i < count; i++ {
		key, value := GenerateRandomString(8), GenerateRandomString(8)
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

func DeleteData(rdb *redis.Client, ctx context.Context, keys []string) {
	fmt.Println("Deleting specific keys...")
	err := rdb.Del(ctx, keys...).Err()
	if err != nil {
		fmt.Println("Error while setting")
		panic(err)
	}
	fmt.Println("Done.")
}

func DeleteAllData(rdb *redis.Client, ctx context.Context) {
	fmt.Println("Deleting all keys...")

	err := rdb.Del(ctx, GetExistingKeysWithPattern(rdb, ctx, "*")...).Err()
	if err != nil {
		fmt.Println("Error while setting")
		panic(err)
	}
	fmt.Println("Done.")
}
