package main

import (
	"fmt"
	"regexp"
)

func main() {

}

func F1() {
	matched, _ := regexp.Match("b", []byte("hello golang"))
	fmt.Println(matched) // false
}

func F2() {
	matched, _ := regexp.MatchString("b", "hello golang")
	fmt.Println(matched)
}

func F3() {
	re := regexp.MustCompile("a")
	match := re.Find([]byte("hello golang"))
	fmt.Println(string(match)) // a
}

// F4 当n < 0时，返回所有匹配个数；
//当n >= 0且n <= 总匹配个数，返回n个结果；
//当n > 总匹配个数，返回所有结果。
func F4() {
	re := regexp.MustCompile("l[a-z]")
	match := re.FindAll([]byte("hello world, hello golang"), -1)
	for _, m := range match {
		fmt.Println(string(m))
	}
	// ll
	// ld
	// ll
	// la

}

func F5() {
	re := regexp.MustCompile("l[a-z]")
	match := re.FindString("hello world, hello golang")
	fmt.Println(match) // ll
}

func F6() {
	re := regexp.MustCompile("l[a-z]")
	match := re.FindAllString("hello world, hello golang", -1)
	for _, m := range match {
		fmt.Println(string(m))
	}
	// ll
	// ld
	// ll
	// la
}

func F7() {
	re := regexp.MustCompile("l[a-z]")
	match := re.FindIndex([]byte("hello world, hello golang"))
	fmt.Println(match) // [2 4]
}

func F8() {
	re := regexp.MustCompile("l[a-z]")
	match := re.FindAllIndex([]byte("hello world, hello golang"), -1)
	for _, m := range match {
		fmt.Println(m)
	}
	// [2 4] [9 11] [15 17] [21 23]
}

func F9() {
	re := regexp.MustCompile(`hello`)
	match := re.Match([]byte("hello everyone"))
	fmt.Println(match) // true
}

func F10() {
	re := regexp.MustCompile(`hello`)
	match := re.MatchString("hello everyone")
	fmt.Println(match) // true
}

func F11() {
	re := regexp.MustCompile(`hello`)
	match := re.ReplaceAll([]byte("hello everyone"), []byte("hi!"))
	fmt.Println(string(match)) // hi! everyone
}

func F12() {
	re := regexp.MustCompile(`hello`)
	match := re.ReplaceAllString("hello everyone", "hi!")
	fmt.Println(match) // hi! everyone
}

func F13() {
	re := regexp.MustCompile(`a`)
	s := re.Split("abacadaeafff", -1)
	fmt.Println(s) // ["", "b", "c", "d", "e", "fff"]
}
