package main

import (
	"log"
	"path"
)

func main() {
	// path.Join 会过滤空字符串，不需要指定分隔符
	str1 := ""
	str2 := "bbb"
	log.Println(path.Join(str1, str2))
}
