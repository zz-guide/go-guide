package main

import (
	"log"
)

/**
测试结构体内部各个类型属性的初始化
*/

type Student struct {
	Id         int
	Name       string
	IsBirthday bool
	Names      []string
	Conf       map[string]string
	Ch         chan int
	Log        Log
	LogPointer *Log
}

type Log struct {
}

func main() {
	t()
}

func t() {
	// 若声明的时候没有进行初始化，那么就是nil
	stu := &Student{
		Names: []string{},
		Conf:  map[string]string{},
	}
	log.Println("int:", stu.Id)
	log.Println("string:", stu.Name)
	log.Println("bool:", stu.IsBirthday)
	log.Println("[]string:", stu.Names, stu.Names == nil)
	log.Println("map[string]string:", stu.Conf, stu.Conf == nil)
	log.Println("Channel:", stu.Ch, stu.Ch == nil)
	log.Println("Log struct:", stu.Log)
	log.Println("Log struct pointer:", stu.LogPointer, stu.LogPointer == nil)
}
