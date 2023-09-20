package main

import (
	"log"
	"reflect"
)

type T struct {
	int
}

func (t T) testT() {
	log.Println("类型 *T 方法集包含全部 receiver T 方法。")
}

func (t T) testQ() {
	log.Println("QQQQQ")
}

func (t *T) testP() {
	log.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

func testT1(t T) {
	log.Println("测试语法糖")
}

/**
结论：
1.T和*T不能包含相同名字的方法
2.类型 *T 方法集包含全部 receiver T + *T 方法。
3.T类型方法生成的时候会生成*T类型的包装方法，为了兼容接口（不会正真调用包装方法，而是做引用处理）
4.T调用*T需要对参数做引用处理。反之需要做解引用处理。
*/
func main() {
	T2()
}

func T2() {
	t1 := reflect.TypeOf(testT1)
	t2 := reflect.TypeOf(T.testT)
	t3 := reflect.TypeOf(T.testQ)
	log.Println("t1 == t2:", t1 == t2) // 函数类型经过反射判断一致，参数类型，返回值相同，类型相同，跟名字无关
	log.Println("t1 == t2:", t1 == t3) // 函数类型经过反射判断一致，参数类型，返回值相同，类型相同
}

func T1() {
	t1 := T{1}
	t1.testP()
	//t1.testT()
	t2 := &T{2}
	log.Printf("t2 is : %v\n", t2)
	t2.testT()
	//t2.testP()
}
