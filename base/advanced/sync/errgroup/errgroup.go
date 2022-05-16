package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

/**
errgroup返回的第一个出错的goroutine抛出的err
*/

func main() {
	T1()
}

func T1() {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		eg.Go(func() error {
			return errors.New("xxxxx")
		})
		return test1(ctx)
	})

	eg.Go(func() error {
		log.Println("asdasdasd")
		return nil
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}

func test1(ctx context.Context) error {
	return errors.New("test2")
}
