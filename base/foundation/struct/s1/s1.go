package main

import "fmt"

/**
结论：2个结构体建议使用reflect.DeepEqual 判断是否相等

1.相同当结构体的属性不含引用属性，且每个属性的值都相等时，== 判定为true
2.当结构体的属性含引用属性，且每个属性的值都相等时，== 判定为false
3.而指针引用，其虽然都是 new(stringSearch)，从表象来看是一个东西，但其具体返回的地址是不一样的。同一个引用可以比较
4.当结构体的属性含切片属性时
5.不同结构体，即使属性一样，也不能相互比较

*/
func main() {
	//F1()
	//F2()
	//F3()
	F4()
}

func F1() {
	// 结论：当结构体的属性不含引用属性，且每个属性的值都相等时，== 判定为true
	type S1 struct {
		Name string
		Age  int
	}

	v1 := S1{Name: "许磊", Age: 1}
	v2 := S1{Name: "许磊", Age: 1}
	if v1 == v2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
}

func F2() {
	// 结论：当结构体的属性含引用属性，且每个属性的值都相等时，== 判定为false
	type S1 struct {
		Name string
		Age  *int
	}

	v1 := S1{Name: "许磊", Age: new(int)}
	v2 := S1{Name: "许磊", Age: new(int)}

	// 而指针引用，其虽然都是 new(stringSearch)，从表象来看是一个东西，但其具体返回的地址是不一样的。同一个引用可以比较
	//age := new(int)
	//v1 := S1{Name: "许磊", Age: age}
	//v2 := S1{Name: "许磊", Age: age}

	if v1 == v2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
}

func F3() {
	// 结论：当结构体的属性含切片属性时，语法报错，invalid operation: v1 == v2 (struct containing []stringSearch cannot be compared)
	// 当其基本类型包含：slice、map、function 时，是不能比较的。若强行比较，就会导致出现例子中的直接报错的情况
	// IDE也会提示报错
	/*type S1 struct {
		Name stringSearch
		Age []stringSearch
	}

	v1 := S1{Name: "许磊", Age:[]stringSearch{"1"}}
	v2 := S1{Name: "许磊", Age:[]stringSearch{"1"}}
	if v1 == v2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}*/
}

func F4() {
	// 结论：不同结构体，即使属性一样，也不能相互比较，invalid operation: v1 == v2 (mismatched types S1 and B1)
	// 如果可以比较，则可以通过强转类型进行比较
	// 强转以后是相等的
	type S1 struct {
		Name string
	}

	type B1 struct {
		Name string
	}

	v1 := S1{Name: "许磊"}
	v2 := B1{Name: "许磊"}
	//if v1 == v2 {
	if v1 == S1(v2) {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
}
