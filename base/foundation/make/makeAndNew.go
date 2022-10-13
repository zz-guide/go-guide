package main

import (
	"log"
	"reflect"
)

type Student struct {
	Name string
}

/**
   问题：new和make区别

	1.make固定用于slice,channel,map的创建，且在堆上分配内存。make 即分配内存，也初始化内存
	make返回的还是引用类型本身

	2.new可以分配任意类型的数据。new只是将内存清零，并没有初始化内存。new返回的是指向类型的指针。
	分配可能是栈内存也可能是堆内存。new不能指定初始化的值。

	3.new等价于&Student{}，var t *T,等同于&,&不能用于基本类型
*/

func main() {
	//TNewArray()
	//TNewSlice()
	//TNewMap()
	//TNewChannel()
	TNewOtherType()
}

func TNewArray() {
	// 结论：new(数组)返回一个数组指针，不是nil
	// array
	var a [5]int
	log.Printf("a: %p %#v \n", &a, a) // 0xc000020150 [5]int{0, 0, 0, 0, 0}

	// array
	ap := new([5]int)
	log.Printf("ap: %p %p  %#v %s\n", &ap, ap, ap, reflect.TypeOf(ap)) //  0xc000010030 0xc0000201b0  &[5]int{0, 0, 0, 0, 0} *[5]int

	ap[1] = 8                            // 等同于：(*ap)[1] = 8
	log.Printf("ap: %p %#v \n", &ap, ap) //  0xc000010030 &[5]int{0, 8, 0, 0, 0}
}

func TNewSlice() {
	// 结论： new(slice)返回一个slice指针，指向nil
	// nil slice
	var a *[]int
	log.Printf("a: %p %#v %t\n", &a, a, a == nil) //a: 0xc042004028 (*[]int)(nil) true

	// slice 同样是nil
	ap := new([]int)
	log.Printf("av: %p %#v %t \n", &ap, ap, *ap == nil) //ap: 0xc000074018 &[]int(nil) true

	*ap = append(*ap, 1)                 // 切片只能使用append进行添加，否则会报错
	log.Printf("av: %p %#v \n", &ap, ap) //panic: runtime error: index out of range
}

func TNewMap() {
	// 结论：new(map)返回一个map指针，指向nil
	// nil map
	var m map[string]string
	log.Printf("m: %p %#v %t\n", &m, m, m == nil) //m: 0xc042068018 map[stringSearch]stringSearch(nil) true

	// nil slice
	mp := new(map[string]string)
	log.Printf("mp: %p %#v %t\n", &mp, mp, *mp == nil) //mv: 0xc000006028 &map[stringSearch]stringSearch(nil)

	(*mp)["a"] = "a"
	log.Printf("mp: %p %#v \n", &mp, mp) //这里会报错panic: assignment to entry in nil map

}

func TNewChannel() {
	// nil channel
	var c chan string
	log.Printf("c: %p %#v %t\n", &c, c, c == nil)

	// nil channel
	cv := new(chan string)
	log.Printf("cv: %p %#v %t\n", &cv, cv, *cv == nil) //cv: 0xc000074018 (*chan stringSearch)(0xc000074020)

	//cv <- "good" //会报 invalid operation: cv <- "good" (send to non-chan type *chan stringSearch)
}

func TNewOtherType() {
	// 结论：new(基础类型)
	var a *int
	log.Printf("a: %p %#v %t\n", &a, a, a == nil) //a: 0xc042004028 (*[]int)(nil) true

	// new(int)
	a1 := new(int)
	log.Printf("a1: %p %#v %t %d\n", &a1, a1, a1 == nil, *a1)

	// 可以赋值
	*a1 = 22
}

func TMake() {
	// 结论：make通常用于slice,map,channel创建
	// make 用于切片
	ma := make([]int, 5)
	log.Printf("ma: %p %#v \n", &ma, ma) //av: 0xc000046400 []int{0, 0, 0, 0, 0}

	ma[0] = 1
	log.Printf("ma: %p %#v \n", &ma, ma) //av: 0xc000046400 []int{1, 0, 0, 0, 0}

	// make用于map
	mv := make(map[string]string)
	log.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[stringSearch]stringSearch{}

	mv["m"] = "m"
	log.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[stringSearch]stringSearch{"m":"m"}

	// make用于channel
	chv := make(chan string)
	log.Printf("chv: %p %#v \n", &chv, chv) //chv: 0xc000074028 (chan stringSearch)(0xc00003e060)

	go func(message string) {

		chv <- message // 存消息

	}("Ping!")

	log.Println(<-chv) // 取消息 //"Ping!"
	close(chv)
}
