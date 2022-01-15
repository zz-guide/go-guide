package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	F3()
}

func F1() {
	s := "你好"
	fmt.Printf("s6:%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	fmt.Printf("s6:%#v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func F2() {
	s := "我hello"
	// 结论：对于字符串，len函数返回的是字节长度;utf8.RuneCountInString返回字符串的长度
	fmt.Println("字符串s长度：", len(s), utf8.RuneCountInString(s))

	a := s
	fmt.Println("字符串a长度：", len(a), utf8.RuneCountInString(a), len([]rune(a)))
	// 结论：a[0]表示的是第一个字节
	fmt.Println("字符串a[]：", a[0], a[1], a[2], a[3])
	fmt.Println("字符串a[]：", string(a[0]), string(a[1]), string(a[2]), string(a[3]))

	// 结论：字符串不能通过a[2] = ''的形式修改，但是可以重新赋值
	a = "asdas"
	// a[2] = 's6'
	fmt.Println("字符串a长度：", len(a), utf8.RuneCountInString(a))

	// 结论：字符串遍历出来的也是数字
	for i, i2 := range a {
		fmt.Println("i:", i)
		fmt.Println("i2:", i2)
	}
}

func F3() {
	// 修改字符串
	s := "hello 世界！"
	b := &s
	fmt.Println("字符串s长度：", &s, len(s), (*reflect.StringHeader)(unsafe.Pointer(&s)))
	r := []rune(s)
	r[6] = '中'
	s = string(r)
	fmt.Println("字符串a长度：", &s, len(s), (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Println(&b, b, *b)
}
