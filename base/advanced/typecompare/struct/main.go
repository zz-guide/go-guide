package main

import (
	"log"
	"reflect"
)

/**
结构体是否可以比较
*/

func main() {
	//SameStructCompare()
	DifferentStructCompare()
}

// SameStructCompare 相同结构体比较
func SameStructCompare() {
	type Student struct {
		name   string
		gender *int
		//ch chan bool
		//m map[string]string
		//fn func()
		//address []string
	}

	// 结论1：对于值变量，属性值相等，并且没有指针的引用，那么就是相等的

	// 相等
	//s1 := Student{name: "许磊"}
	//s2 := Student{name: "许磊"}

	// 结论2：当结构体属性含有指针属性时，不相等，除非都是nil
	// 不相等
	//s1 := Student{name: "许磊", gender: new(int)}
	//s2 := Student{name: "许磊", gender: new(int)}
	// 结论：以下这种写法是可以比较的
	gender := new(int)
	s1 := Student{name: "许磊", gender: gender}
	s2 := Student{name: "许磊", gender: gender}

	// 结论3：当结构体含有切片类型的属性时，==会报错，Invalid operation: s1 == s2 (the operator == is not defined on Student)
	//s1 := Student{name: "许磊", address: []string{"北京", "上海"}}
	//s2 := Student{name: "许磊", address: []string{"北京", "上海"}}
	if s1 == s2 {
		log.Println("相等")
	} else {
		log.Println("不相等")
	}
}

// DifferentStructCompare 不同结构体比较
func DifferentStructCompare() {
	type Student struct {
		name   string
		gender *int
	}

	type Teacher struct {
		name   string
		gender *int
	}

	// 结论1：不同的结构体含有相同的属性，无法比较
	// Invalid operation: s1 == t1 (mismatched types Student and Teacher)
	//s1 := Student{name: "许磊"}
	//t1 := Teacher{name: "许磊"}

	// 结论2：可以借助强制转换类型实现比较,相等，前提是2种类型必须是可以比较的，否则不行
	//if s1 == Student(t1) {
	//	log.Println("相等")
	//} else {
	//	log.Println("不相等")
	//}

	// 结论：当结构体含有slice,map,function时是不可以比较的，强行比较会报错
	// 指针引用，new(string)，返回的具体地址不一样，除非说是用一个指针然后分别赋值给2个变量
	// 结论：终极法宝，使用reflect.DeepEqual来进行比较
	// 规则：1.不同类型的变量永远不相等
	// 相同类型的变量，值深度相等，则相等
	// 当数组值对应元素深度相等，则相等
	s1 := Student{name: "许磊"}
	t1 := Teacher{name: "许磊"}
	// 具体查看源码说明
	if reflect.DeepEqual(s1, t1) {
		log.Println("相等")
	} else {
		log.Println("不相等")
	}
}
