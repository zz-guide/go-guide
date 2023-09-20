package eee

import "go-guide/base/foundation/struct/embed/e4/t"

/**
http://golang.org/ref/spec#Method_sets
https://www.goinggo.net/2014/05/methods-interfaces-and-embedded-types.html
*/

type Chicken struct {
	feet int
	//IAnimal
}

func (_cat Chicken) FnInterface() string {
	return "猫->FnInterface"
}

func (_cat Chicken) private() {

}

// 结论：1.接口中如果有未导出的方法，其他包就不能实现该接口，除了直接嵌套的情况。
var _ t.IAnimal = (*Chicken)(nil)
