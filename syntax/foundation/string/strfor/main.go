package main

import (
	"log"
)

func main() {
	strFor()
}

func strFor() {
	str := "Hello,世界"
	log.Println("----字符串普通for遍历----")
	for i := 0; i < len(str); i++ {
		ch := str[i]
		log.Println(ch)
	}

	log.Println("----字符串for range 遍历----")
	for _, ch1 := range str {
		log.Println(ch1) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
	}
}
