package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

/**
instance.method(args...) ---> <type>.func(instance, args...)
前者称为 method value，后者 method expression。
两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。
*/
func main() {
	u := User{1, "Tom"}
	u.Test()

	mValue := u.Test
	//instance.method(args...)
	mValue() // 隐式传递 receiver

	//<type>.func(instance, args...)
	mExpression := (*User).Test
	mExpression(&u) // 显式传递 receiver
}
