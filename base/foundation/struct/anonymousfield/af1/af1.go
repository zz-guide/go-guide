package main

import (
	"fmt"
	"log"
)

// 内嵌结构体，可以是值也可以是指针

type User struct {
	Id   int
	Name string
}

// 对于匿名字段，一种类型只能有一个没名字的，不能有多个，但是有名字的可以有多个

type Manager struct {
	User
	Title string
	Name  string
}

func (user *User) ToString() string {
	return fmt.Sprintf("User: %p, %+v", user, user)
}

/*func (user *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %+v", user, user)
}*/

func main() {
	m := Manager{User{1, "Tom"}, "Administrator", "ss"}
	log.Println(m.ToString())
	//log.Println(m.User.ToString())

	// 对于匿名字段没办法使用指定具体的属性创建,只能使用不指定属性名字的方式创建
	// 若属性名字冲突，优先使用自身的，其次才是内嵌的
	m1 := Manager{User{1, "Tom"}, "Administrator", ""}
	m1.User.Name = "xxx"
	log.Println(m1.ToString())
	log.Println("Name:", m1.Name)
}
