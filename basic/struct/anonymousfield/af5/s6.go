package main

import "log"

// 内嵌接口

type User struct {
	IUser
}

type IUser interface {
	Fn()
}

var _ IUser = (*User)(nil)

func main() {
	T1()
}

func T1() {
	// 1.判断struct是否实现了某个interface
	// var _ interface = (*结构体)(nil) 形式可以判断一个struct是否继承了一个interface
	// 缺点：嵌入匿名interface，也会被编译器认为实现了，但实际上并没有，运行会报错

	var user IUser = User{}
	log.Printf("地址：%p\n", &user)
	user.Fn()
}
