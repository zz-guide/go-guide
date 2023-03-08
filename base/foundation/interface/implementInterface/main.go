package main

type Animal interface {
	name() string
}

type Dog struct {
}

func (d Dog) name() string {
	return "Dog"
}

func main() {
	T1()
}

func T1() {
	// 借助编译器保证类型实现某个接口
	var _ Animal = &Dog{}
}
