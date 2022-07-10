package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func GetExistingKeysWithPattern(rdb *redis.Client, ctx context.Context, pattern string) []string {
	fmt.Println("Get all existing keys: ")
	keys := []string{}

	iter := rdb.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	fmt.Println(keys)
	fmt.Println()

	return keys
}

func PrintAllData(rdb *redis.Client, ctx context.Context) {
	fmt.Println("DATA:")
	keys := GetExistingKeysWithPattern(rdb, ctx, "*")
	PrintData(rdb, ctx, keys)
}

func PrintData(rdb *redis.Client, ctx context.Context, keys []string) {
	if len(keys) > 0 {
		fmt.Println("Key/value pairs:")
		for _, v := range keys {
			data := rdb.Get(ctx, v)
			fmt.Println(v, data.Val())

			if err := data.Err(); err != nil {
				fmt.Println("Error while fetching")
				panic(err)
			}
		}
		fmt.Println()
	} else {
		fmt.Println("No keys exist!")
	}
}

func InsertRandomData(rdb *redis.Client, ctx context.Context, count int) []string {
	fmt.Println("Setting random key/value pairs...")
	keys := []string{}
	for i := 0; i < count; i++ {
		key, value := GenerateRandomString(8), GenerateRandomString(8)
		err := rdb.Set(ctx, key, value, 0).Err()
		keys = append(keys, key)

		if err != nil {
			fmt.Println("Error while inserting")
			panic(err)
		}
	}
	fmt.Println("Keys of the inserted elements:", keys)
	fmt.Println()

	return keys
}

func DeleteData(rdb *redis.Client, ctx context.Context, keys []string) {
	fmt.Println("Deleting data for:", keys)
	err := rdb.Del(ctx, keys...).Err()
	if err != nil {
		fmt.Println("Error while deleting")
		panic(err)
	}
	fmt.Println("Done.")
}

func DeleteAllData(rdb *redis.Client, ctx context.Context) {
	fmt.Println("DELETING ALL KEYS!!!")

	err := rdb.Del(ctx, GetExistingKeysWithPattern(rdb, ctx, "*")...).Err()
	if err != nil {
		fmt.Println("Error while deleting")
		panic(err)
	}
	fmt.Println("Done.")
}

func PrintConfig(rdb *redis.Client, ctx context.Context, pattern string) {
	fmt.Println("CONFIG:")
	fmt.Println(TransformData(rdb.ConfigGet(ctx, pattern).Val()))
	fmt.Println()
}
