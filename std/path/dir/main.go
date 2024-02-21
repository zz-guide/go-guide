package main

import (
	"log"
	"path"
)

func main() {
	// path.Dir 返回去掉最后一个元素，剩余部分，并且自动去掉末尾斜杠
	str := "./aaa/bbb/ccc/e.txt"
	log.Println(path.Dir(str))
}
