package main

import (
	"fmt"
	"log"
)

/**
Functional Options Pattern（函数式选项模式）
Golang Option模式充分利用了golang 的闭包、函数式参数、可变参数这三种特性。根据不同的Option函数实现不同的操作。
*/
func main() {
	T1()
}

func T1() {
	// 一般学生对象
	stu1 := NewStudent(1, "张三")
	log.Println(fmt.Sprintf("id=%d, name=%s, addr=%s, age=%d", stu1.id, stu1.name, stu1.addr, stu1.age))

	// 自定义addr属性
	stu2 := NewStudent(2, "李四", WithAddr("中国"))
	log.Println(fmt.Sprintf("id=%d, name=%s, addr=%s, age=%d", stu2.id, stu2.name, stu2.addr, stu2.age))

	// 自定义addr和age属性
	stu3 := NewStudent(2, "李四", WithAddr("中国"), WithAge(20))
	log.Println(fmt.Sprintf("id=%d, name=%s, addr=%s, age=%d", stu3.id, stu3.name, stu3.addr, stu3.age))
}

func NewStudent(id int64, name string, options ...Option) *Student {
	s := &Student{
		id:   id,
		name: name,
		addr: "unknown",
		age:  18,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

type Student struct {
	id   int64
	name string
	addr string
	age  int
}

type Option func(*Student)

func WithAddr(addr string) Option {
	return func(student *Student) {
		if addr != "" {
			student.addr = addr
		}
	}
}

func WithAge(age int) Option {
	return func(student *Student) {
		if age > 0 {
			student.age = age
		}
	}
}
