package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan struct{}), make(chan struct{})

	go func() {
		defer wg.Done()

		<-a

		for i := 0; i < 5; i++ {
			log.Println("i:", i)
		}
	}()

	go func() {
		defer wg.Done()

		<-b

		for i := 0; i < 5; i++ {
			log.Println("i:", i)
			if i == 2 {
				runtime.Gosched()
			}
		}
	}()
}
