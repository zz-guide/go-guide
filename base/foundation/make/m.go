package main

import "fmt"

type Student struct {
	Name string
}

//make和new都是golang用来分配内存的內建函数，且在堆上分配内存，make 即分配内存，也初始化内存。new只是将内存清零，并没有初始化内存。
//make返回的还是引用类型本身；而new返回的是指向类型的指针。
//make只能用来分配及初始化类型为slice，map，channel等结构；new可以分配任意类型的数据。
func main() {
	example1()
}

func example2() {
	var a [5]int
	fmt.Printf("a: %p %#v \n", &a, a)

	ap := new([5]int)
	fmt.Printf("ap: %p \n", &ap)
	fmt.Printf("ap: %p \n", ap)
	fmt.Printf("ap: %p %#v \n", &ap, ap)

	ap[1] = 8 // 等同于：(*ap)[1] = 8
	fmt.Printf("ap: %p %#v \n", &ap, ap)
}

func example3() {
	var a *[]int

	fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc042004028 (*[]int)(nil)

	av := new([]int)

	fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[]int(nil)

	*av = append(*av, 1)                 // 切片只能使用append进行添加，否则会报错
	fmt.Printf("av: %p %#v \n", &av, av) //panic: runtime error: index out of range
}

func example4() {
	var m map[string]string

	fmt.Printf("m: %p %#v \n", &m, m) //m: 0xc042068018 map[stringSearch]stringSearch(nil)

	mv := new(map[string]string)

	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000006028 &map[stringSearch]stringSearch(nil)

	(*mv)["a"] = "a"

	fmt.Printf("mv: %p %#v \n", &mv, mv) //这里会报错panic: assignment to entry in nil map

}

func example5() {
	cv := new(chan string)

	fmt.Printf("cv: %p %#v \n", &cv, cv) //cv: 0xc000074018 (*chan stringSearch)(0xc000074020)

	//cv <- "good" //会报 invalid operation: cv <- "good" (send to non-chan type *chan stringSearch)
}

func example1() {
	// make用于slice，map，和channel的初始化
	// new用来分配内存的，等同于&,&不能用于基本类型

	var a *int
	fmt.Println("a:", a) // 指针如果是nil的话不能直接赋值，需要使用new来分配内存
	// 此时应该使用new来分配
	a = new(int) // 指向的是类型的0值
	fmt.Println("a:", *a)
	//*a = 2
	fmt.Println("a的值:", a) // runtime error: invalid memory address or nil pointer dereference
}

func example6() {
	av := make([]int, 5)

	fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000046400 []int{0, 0, 0, 0, 0}

	av[0] = 1

	fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000046400 []int{1, 0, 0, 0, 0}

	mv := make(map[string]string)

	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[stringSearch]stringSearch{}

	mv["m"] = "m"

	fmt.Printf("mv: %p %#v \n", &mv, mv) //mv: 0xc000074020 map[stringSearch]stringSearch{"m":"m"}

	chv := make(chan string)

	fmt.Printf("chv: %p %#v \n", &chv, chv) //chv: 0xc000074028 (chan stringSearch)(0xc00003e060)

	go func(message string) {

		chv <- message // 存消息

	}("Ping!")

	fmt.Println(<-chv) // 取消息 //"Ping!"

	close(chv)
}
