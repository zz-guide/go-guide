package main

import (
	"log"
)

func main() {

}

//go语言中的源码定义为utf-8文本，不允许其他的表示。
//但是也存在特殊处理，那就是字符串上使用for…range循环。
//range循环迭代时，就会解码一个utf-8编码的rune。

func F1() {
	str := "Hello,世界"
	log.Println("Utf-8遍历")
	for i := 0; i < len(str); i++ {
		ch := str[i]
		log.Println(ch)
	}

	log.Println("Unicode遍历")
	for _, ch1 := range str {
		log.Println(ch1)
	}
}
