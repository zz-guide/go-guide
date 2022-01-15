package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//2006-01-02 15:04为golang语言的诞生时间，如果再golang中需要打印当前时间，就需这个固定的值。
	//可以简单的记做 2006 1 2 3 4 5
	//year      month       day     hour    minute      second
	//2006      01          02      15      04          05
	fmt.Println(now.Format("2006-01-02 15:04"))
}
