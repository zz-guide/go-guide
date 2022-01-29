package main

import (
	"context"
	"go-guide/libother/db/redis/go-redis/common"
	"strconv"
)

func main() {
	ctx := context.Background()

	rdb := common.Rdb()
	defer rdb.Close()

	for i := 1; i < 1001; i++ {
		err := rdb.Publish(ctx, "name", "hello"+strconv.Itoa(i)).Err()
		if err != nil {
			panic(err)
		}
	}
}
