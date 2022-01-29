package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-guide/libother/db/redis/go-redis/common"
)

func main() {
	rdb := common.Rdb()
	defer rdb.Close()

	ctx := context.Background()
	/*	info:=redis.NewStatusCmd("bf.add", "bl", "1")
		_ = GlobalClient.Process(info)
		if err := info.Err(); err != nil {
			print(err)
		}
		info1:=redis.NewStatusCmd("bf.add", "bl", "2")
		_ = GlobalClient.Process(info1)
		if err := info1.Err(); err != nil {
			print(err)
		}
		info3:=redis.NewStatusCmd("bf.add", "bl", "3")
		_ = GlobalClient.Process(info3)
		if err := info3.Err(); err != nil {
			print(err)
		}*/
	info4 := redis.NewIntCmd(ctx, "bf.exists", "bl", "6")
	_ = rdb.Process(ctx, info4)
	if err := info4.Err(); err != nil {
		print(err)
	}
	v, err := info4.Result()
	fmt.Println("err", err)
	fmt.Println("v", v) //存在 v==1  不存在 ==0
	//fmt.Println(GlobalClient.Get("mykey").String())
}
