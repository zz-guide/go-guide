package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./aaa/bbb/ccc/a.ff"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("MkdirAll创建目录失败：", err)
	} else {
		fmt.Println("MkdirAll创建目录成功")
	}
}
