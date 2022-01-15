package main

import (
	"fmt"
	"sync"
)

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

func main() {
	a:=1
	_ = a
	fmt.Println("------开始-------", a)
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
	fmt.Println("------结束-------")
}
