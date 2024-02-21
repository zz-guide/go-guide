package main

import (
	"log"
	"strings"
)

func main() {
	// 需要注意，strings.Join 不会过滤空字符串
	str1 := ""
	str2 := "bbb"
	log.Println(strings.Join([]string{str1, str2}, "/")) // 输出 /bbb
}
