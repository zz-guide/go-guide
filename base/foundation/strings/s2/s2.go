package main

import "fmt"

func main() {
	s1 := "woaini"
	// 字符串通过slice的语法，返回的还是字符串
	fmt.Println("kmp[:]", s1[:], s1[:] == s1)
	fmt.Println("kmp[0]", s1[0])
}

func f1() {
	str := "烫"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}
}

func f2() {
	str := "烫烫烫烫"
	array := []rune(str)
	n := len(array)
	for i := 0; i < n; i++ {
		ch := array[i]     // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch) //unicode 编码转十进制输出
	}
}

func f3() {
	str := "烫烫烫烫"
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune 默认utf-8编码，一个汉字三个字节
	}
}
