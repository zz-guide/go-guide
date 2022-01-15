package main

import (
	"fmt"
	"sync"
	"time"
)
/**
		结论：不知道为啥panic以后。waitGroup内的defer还会执行
		有了sleep(3)之后就会执行了？
 */
func main() {
	fmt.Println("----开始----")
	var wg sync.WaitGroup
	defer defer3()
	wg.Add(1)
	defer defer1()
	go goroutine(&wg)
	fmt.Println("asdasdasd")
	wg.Wait()
	fmt.Println("----结束----")
}

func defer1() {
	fmt.Println("defer1")
}

func defer2() {
	fmt.Println("defer2")
}

func defer3() {
	fmt.Println("defer3")
}

func goroutine(wg *sync.WaitGroup) {
	fmt.Println("ffffff")
	defer wg.Done()
	defer defer2()
	//time.Sleep(3)
	time.Sleep(3 * time.Second)
	panic("panic occured in goroutine")
}