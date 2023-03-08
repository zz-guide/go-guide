package main

import (
	"fmt"
	"time"
)

/**
2006-01-02 15:04为golang语言的诞生时间
可以简单的记做 2006 1 2 3 4 5
year      month       day     hour    minute      second
2006      01          02      15      04          05
*/

func main() {
	now := time.Now()
	fmt.Println(now.Year())   // 年
	fmt.Println(now.Month())  // 月
	fmt.Println(now.Day())    // 日
	fmt.Println(now.Hour())   // 时
	fmt.Println(now.Minute()) // 分
	fmt.Println(now.Second()) // 秒

	fmt.Println(now.Format("2006-01-02 15:04"))
}
