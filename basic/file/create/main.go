package main

import (
	"log"
	"os"
)

/*
*
1. 路径中的目录必须存在，os.Create 不会自动创建目录
*/
func main() {
	file, err := os.Create("./aaa/c.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	log.Println("创建文件成功")
}
