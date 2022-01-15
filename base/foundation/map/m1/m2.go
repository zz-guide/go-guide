package main

import "fmt"

// map定义方法，默认是指针类型的，不能写*
// 结构体的属性可以取地址，map不能对属性取地址，二者有区别
func main() {

}

type Dict map[string]int

func (d Dict) PrintDict() {
	fmt.Println(d)
}

// Add *会报错
func (d Dict) Add(key string, value int) {
	d[key] = value
}
