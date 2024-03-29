package main

import (
	"context"
	crypto_rand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	math_rand "math/rand"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/snehalyelmati/redis-go-cli/utils"
)

var (
	HOSTNAME = flag.String("hostname", "localhost", "Hostname of the Redis server")
	PORT     = flag.Int("port", 6379, "Port number of the Redis server")
	USERNAME = flag.String("username", "", "Username for authentication")
	PASSWORD = flag.String("password", "", "Password for authentication")

	config            = flag.Bool("config", false, "Boolean flag to print Redis config")
	data              = flag.Bool("data", false, "Boolean flag to print all the data in Redis")
	configWithPattern = flag.String("configWithPattern", "", "To get configuration with pattern")

	testReadWrite    = flag.Bool("testReadWrite", false, "Test the Redis connection by reading and writing sample data (default is 5 records)")
	count            = flag.Int("count", 5, "Number of random records to insert and/or delete (default is 5)")
	insertSampleData = flag.Bool("insertSampleData", false, "Inserts sample data into Redis based on the count parameter (default is 5)")
	deleteAllData    = flag.Bool("deleteAllData", false, "Deletes all data in the Redis instance!")
)

func main() {
	fmt.Println("Redis Go CLI")
	flag.Parse()

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *HOSTNAME + ":" + strconv.Itoa(*PORT),
		Username: *USERNAME,
		Password: *PASSWORD,
		DB:       0,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	} else {
		fmt.Println("Redis connection status: SUCCESS")
	}
	fmt.Println()

	// initialize non-deterministic random number
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Cannot seed math/rand package with cryptographically secure random number generation...")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	// testReadWrite
	if *testReadWrite {
		keys := utils.InsertRandomData(rdb, ctx, *count)
		utils.PrintData(rdb, ctx, keys)
		utils.DeleteData(rdb, ctx, keys)
	}

	// get all config
	if *config {
		utils.PrintConfig(rdb, ctx, "*")
	}

	// get all data
	if *data {
		utils.PrintAllData(rdb, ctx)
	}

	// get config with pattern
	if *configWithPattern != "" {
		utils.PrintConfig(rdb, ctx, *configWithPattern)
	}

	// insert sample data
	if *insertSampleData {
		utils.InsertRandomData(rdb, ctx, *count)
	}

	// delete all existing data
	if *deleteAllData {
		utils.DeleteAllData(rdb, ctx)
	}
}
