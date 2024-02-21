package main

import (
	"log"
	"path"
)

func main() {
	// path.Base 返回路径中最后一个元素，并且会去掉最后一个斜杠
	str := "./aaa/bbb/ccc/a.txt"
	log.Println(path.Base(str))
}
