package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/**
1.context.Background(),context.TODO()
2.context.WithCancel(),context.WithDeadline(),context.WithValue(),context.WithTimeout()
3.emptyCtx
4.context.WithValue()key最好不要设置成string,int等基础类型，应该起个别名，防止与其他context冲突
*/
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	eatNum := chiHanBao /**/ (ctx) /**/
	for n := range eatNum {
		fmt.Println("asdasd")
		if n >= 10 {
			cancel()
			break /**/
		}
	}

	fmt.Println("正在统计结果。。。")
	time.Sleep(1 * time.Second)
}

func chiHanBao(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	n := 0
	// 时间
	t := 0
	go func() {
		for {
			//time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("耗时 %d 秒，吃了 %d 个汉堡 \n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n >= 10 {
					n = 10
				}
				t++
				fmt.Printf("我吃了 %d 个汉堡\n", n)
			}
		}
	}()

	return c
}
