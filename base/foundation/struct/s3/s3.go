package main

import "log"

type Student struct {
	name string
	age  int
}

// 方法体
func (s Student) GetName() string {
	log.Printf("GetName s地址=%p\n", &s)
	return s.name
}

// 接收者
func (s *Student) AddAge(i int) {
	log.Printf("AddAge s地址=%p\n", &s)
	s.age += i
}

func main() {
	s := Student{name: "许磊", age: 27}
	log.Printf("s地址=%p,+%v\n", &s, s)
	s.AddAge(3)
	s.GetName()

	s1 := &Student{name: "李四", age: 23}
	s1.AddAge(2)
	s1.GetName()

}
