package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func hardWork(job interface{}) error {
	log.Println("---开始等待---")
	time.Sleep(time.Second * 10)
	log.Println("ssssssss")
	return nil
}

// done := make(chan error, 1) 有1个缓冲的话，即使没有接收，也不会阻塞
// 如果是无缓冲的话，没有接收就会阻塞
func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	done := make(chan error, 1)
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case p := <-panicChan:
		panic(p)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	const total = 10
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer func() {
				if p := recover(); p != nil {
					fmt.Println("oops, panic")
				}
			}()

			defer wg.Done()
			err := requestWork(context.Background(), "any")
			log.Println("err:", err)
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	time.Sleep(time.Second * 20)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}
