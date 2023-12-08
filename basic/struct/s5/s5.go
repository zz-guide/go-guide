package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	num      int
	language string
}

func (p Gopher) code() {
	p.num++
	fmt.Printf("I am coding %s language, num is %d\n", p.language, p.num)
}

func (p *Gopher) debug() {
	p.num++
	fmt.Printf("I am debuging %s language, num is %d\n", p.language, p.num)
}

func (p Gopher) test1() {
	fmt.Printf("I am test1 %s language, num is %d\n", p.language, p.num)
	p.test()
	fmt.Printf("I am test1 %s language, num is %d\n", p.language, p.num)
	p.num += 4
}

func (p *Gopher) test() {
	p.num += 3
	fmt.Printf("I am test %s language, num is %d\n", p.language, p.num)
}

/*
  1、使用值调用
*/
func main() {
	F4()
}

func F5() {
	var c Gopher = Gopher{1, "Go"}
	//c.code()
	//c.debug()
	//c.test1()
	c.test()
	fmt.Println("cccc:", c.num)
}

func F1() {
	// 2、使用指针调用
	var c *Gopher = &Gopher{1, "Go"}
	c.code()
	c.debug()
	c.code()
}

func F2() {
	// 使用interface作为调用者
	var c coder = &Gopher{1, "Go"}
	c.code()
	c.debug()
	c.code()
}

func F3() {
	//使用interface作为调用者
	//         cannot use Gopher literal (type Gopher) as type coder in assignment:
	//                 Gopher does not implement coder (debug method has pointer receiver)
	//var c coder = Gopher{1, "Go"} //此处报错
	//c.code()
	//c.debug()
	//c.code()
}

type MaxHeap []int

func (h MaxHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func F4() {
	h := MaxHeap{1, 2, 3, 4}
	h.swap(0, 1)
	fmt.Println(h)
}
