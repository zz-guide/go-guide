package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	//ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	chiHanBao(ctx)
	fmt.Println("asdasdas")
	defer cancel()
}

func chiHanBao(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop \n")
			return
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡\n", n)
		}
		time.Sleep(1 * time.Second)
	}
}
