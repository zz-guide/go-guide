package main

import (
	"context"
	"fmt"
	"go-guide/libother/db/redis/go-redis/common"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	rdb := common.Rdb()
	defer rdb.Close()

	result, err := rdb.BLPop(ctx, 0, "queue").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(result[0], result[1])
	// Output: queue message

	fmt.Println("关闭")
}
