package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
退出goroutine的2种方式：
1.return结束当前函数,并返回指定值
2.runtime.Goexit结束当前goroutine,其他的goroutine不受影响,主程序也一样继续运行
*/
func main() {
	fmt.Println("----开始-----:")
	printGNum()

	go func() {
		fmt.Println("go1 runtime.Goexit")
		printGNum()
		time.Sleep(time.Second * 3)
		runtime.Goexit()
	}()

	go func() {
		fmt.Println("go2 return")
		printGNum()
		time.Sleep(time.Second * 6)

		return
	}()

	time.Sleep(time.Second * 15)
	fmt.Println("----结束-----")
	printGNum()
}

func printGNum() {
	fmt.Println("g的数量:", runtime.NumGoroutine())
}
