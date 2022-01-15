package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	howClose()
}

func sendClose() {
	/**
		向一个已经关闭的channel发送数据会panic
		panic: send on closed channel
	*/
	c := make(chan int, 3)
	close(c)
	c <- 1
}

func readClose() {
	/**
			读已经关闭的通道，如果有值就一直读，没值就返回类型的零值
	 */

	fmt.Println("以下是数值的 chan")
	ci := make(chan byte, 0)
	go func() {
		ci <- 255		// constant 256 overflows byte,溢出了
		fmt.Println("ssssss")
		close(ci)
	}()

	num, ok := <-ci
	fmt.Printf("读chan的协程结束，num=%+v, ok=%v\n", num, ok)
	num1, ok1 := <-ci
	fmt.Printf("再读chan的协程结束，num=%+v, ok=%v\n", num1, ok1)
	num2, ok2 := <-ci
	fmt.Printf("再再读chan的协程结束，num=%+v, ok=%v\n", num2, ok2)


	fmt.Println("以下是字符串的chan")
	cs := make(chan string, 3)
	cs <- "aaa"
	close(cs)

	str, ok := <-cs
	fmt.Printf("读chan的协程结束,num=%+v, ok=%v\n", str, ok)
	str1, okl := <-cs
	fmt.Printf("再读chan的协程结束，num=%+v, ok=%v\n", str1, okl)
	str2, ok2 := <-cs
	fmt.Printf("再再读chan的协程结束，num=%+v, ok=%v\n", str2, ok2)


	fmt.Println("以下是结构体的chan")
	type MyStruct struct {
		Name string
	}
	cstruct := make(chan MyStruct, 3)
	cstruct <- MyStruct{Name: "haha"}
	close(cstruct)

	stru, ok := <-cstruct
	fmt.Printf("读chan的协程结束，num=%v, ok=%v\n", stru, ok)
	stru1, ok1 := <-cstruct
	fmt.Printf("再读chan的协程结束，num=%v, ok=%v\n", stru1, ok1)
	stru2, ok2 := <-cstruct
	fmt.Printf("再再读chan的协程结束，num=%v, ok=%v\n", stru2, ok2)
}


func howClose(){
	/**
	   链接：https://www.zhihu.com/question/450188866/answer/1790314327
	 */

	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.TODO())

	// 专门关闭的协程
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("xxxxxxx")
		cancel()
		// ... 某种条件下，关闭 channel
		close(c)
	}()

	// senders（写端）
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, id int) {
			select {
			case <-ctx.Done():
				fmt.Printf("关闭")
				return
			case c <- id: // 入队
				// ...
			}
		}(ctx, i)
	}

	// receivers（读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v := <-c
			fmt.Println("aaa:", v)
		}()
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
	time.Sleep(10 * time.Second)
}