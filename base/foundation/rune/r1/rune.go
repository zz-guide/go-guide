package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	F2()
}

func F1() {
	// rune 是专门处理中文等复合字符的类型，是int32的别名, byte 是uint8的别名
	var ch = "我是"
	var ch1 rune
	fmt.Println("ch1:", ch1)
	fmt.Println("length1:", len(ch))                    // 默认是UTF8字符，所以是6个字节
	fmt.Println("length2:", len([]rune(ch)))            // 返回字符数，2
	fmt.Println("length3:", utf8.RuneCountInString(ch)) // 返回字符数，2
	fmt.Println("length4:", utf8.RuneCount([]byte(ch))) // 返回字符数，2
}

func F2() {
	r := 'c'
	res := r - 'a'
	log.Println("res:", res)
}
