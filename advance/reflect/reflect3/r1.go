package main

import (
	"log"
	"reflect"
)

/**

 */

type People interface {
}

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	//F1()
	F2()
}

func F1() {
	stu := &Student{Id: 1, Name: "许磊"}

	// 值和类型的反射对象
	// 分别获取值和类型的反射对象的原型
	rt := reflect.TypeOf(stu)
	rte := rt.Elem()

	rv := reflect.ValueOf(stu)
	rve := rv.Elem()

	// 动态修改，结论：直接修改反射值的属性会直接作用到变量身上
	rve.Field(0).Set(reflect.ValueOf(12))
	rve.FieldByName("Name").SetString("李四")
	log.Printf("stu:%+v\n", stu)

	// 1. reflect.New() 返回一个新的对应类型的零值指针
	newStu := reflect.New(rte)
	log.Println("newStu:", newStu.Kind())

	// 2.动态修改值
	newStu.Elem().FieldByName("Name").SetString("张三")

	// 3.已知类型的情况下可以使用以下方法来生成对应类型的变量
	//s := newStu.Interface().(*Student)
	//log.Printf("s:%+v\n", s)

	unwrapFields(newStu)
}

func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

func unwrapFields(v reflect.Value) []reflect.Value {
	var fields []reflect.Value
	indirect := reflect.Indirect(v)
	for i := 0; i < indirect.NumField(); i++ {
		// 按照定义的顺序，获取当前结构体的属性，只有Struct才可以调用
		child := indirect.Field(i)
		// 如果是指针的话，先初始化
		if child.Kind() == reflect.Ptr && child.IsNil() {
			baseValueType := Deref(child.Type())
			child.Set(reflect.New(baseValueType))
		}

		// 获取当前属性的原始值
		child = reflect.Indirect(child)
		childType := indirect.Type().Field(i)
		// 如果是匿名结构体，递归处理
		if child.Kind() == reflect.Struct && childType.Anonymous {
			fields = append(fields, unwrapFields(child)...)
		} else {
			fields = append(fields, child)
		}
	}

	return fields
}

func F2() {
	//var stu People
	//var stu func(i int)
	var stu interface {
		Less(i int)
	}
	//stu = &Student{Id: 1, Name: "许磊"}
	//stu = stu.(People)

	// 1.interface类型的变量是可以直接获取内存地址的，struct变量必须加&
	// 2.通过value获取的Type与Typeof()获取的Type本质上是同一个变量
	// 3.什么情况下反射结果是Interface？
	rt := reflect.TypeOf(stu)
	log.Printf("rt:%+v, %p, %p\n", rt.Kind(), rt, &rt)
	rv := reflect.ValueOf(stu)
	log.Printf("rv:%+v, %p\n", rv.Kind(), &rv)
	rvt := rv.Type()
	log.Printf("rvt:%+v, %p, %p\n", rvt.Kind(), rvt, &rvt)
}
