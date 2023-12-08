package main

import (
	"log"
)

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id   int
	addr string
	//同名字段
	name string
}

func main() {
	var s Student
	// 给自己字段赋值了
	s.name = "5lmh"
	log.Println(s)

	// 若给父类同名字段赋值，如下
	s.Person.name = "枯藤"
	log.Println(s)
}
