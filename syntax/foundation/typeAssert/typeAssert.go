package main

import "fmt"

/**
	类型断言：即Comma-ok断言

写法为value, ok := em.(T)   如果确保em 是同类型的时候可以直接使用value:=em.(T)一般用于switch语句中下面将会讲解到

em代表要判断的变量
T代表被判断的类型
value代表返回的值
ok代表是否为改类型
类型断言应该一看就懂 在这里就不再介绍了主要是介绍我自己碰到的几个问题

*/
func main() {
	F3()
}

func F1() {
	var s interface{}
	s = "sss"
	if v, ok := s.(string); ok {
		fmt.Println(v)
	}
}

func F2() {
	s := "sss"
	if v, ok := interface{}(s).(string); ok {
		fmt.Println(v)
	}
}

func F3() {
	type Element interface{}
	var e Element = 100
	switch value := e.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("stringSearch", value)
	default:
		fmt.Println("unknown", value)
	}
}
