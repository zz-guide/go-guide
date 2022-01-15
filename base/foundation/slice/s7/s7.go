package main

// slice定义方法，默认是指针类型的，不能写*
func main() {

}

type Category []int

func (c Category) Add(n int) {
	for i, v := range c {
		c[i] = v + n
	}
}
