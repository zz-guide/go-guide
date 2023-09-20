package main

import "log"

/**
结论：定义指针接收者方法和值接受者方法的标准
	1.如果方法是要实现接口中的方法，那么就全定义成指针接受者。如果定义成值接受者，编译期间还会生成包装方法，代码量增多
	2.其他情况如果要修改内部值，定义成指针，其他情况定义成值
*/

type S struct {
	name string
}

func (s S) GetName() string {
	return s.name
}

func (s *S) SetName(name string) {
	s.name = name
}

type TInterface interface {
	GetName() string
	SetName(name string)
}

func main() {
	//ss := S{name: "许磊"}
	//log.Println(ss.GetName())

	// go tool compile -trimpath="`pwd`=>" -l -p gom main.go
	// go tool nm main.o | grep T
	// go tool nm main | grep T

	// 通过命令编译得知生成以下方法
	/**
	2d49 T gom.(*S).GetName
	   2873 T gom.(*S).SetName
	   286d T gom.S.GetName
	   28d2 T gom.main

	*/

	/**
	结论：编译器会为值接收者生成一个指针接受者的同名包装方法，但不是为了指针调用。
	包装方法是为了非空接口类型变量调用方法时使用的。值通常是有大小的，而指针大小是固定的，编一阶段通常无法确定值大小，但是
	指针大小可以确定，所以接口不能直接使用值接收方法。但是最终可执行文件中不一定有这些方法。
	不只是这些包装方法，就连代码中的原始方法也不一定会存在于可执行文件中。
	链接器裁剪的时候会检查用户代码是否会通过反射来调用方法，会的话就把该类型的方法保留下了，只有在明确确认这些方法在运行阶段不会被用到时，才可以安全的裁剪。
	一是代码中存在从目标类型到接口类型的赋值操作，因为运行阶段类型信息萃取始于接口；
	二是代码中调用了MethodByName、Method这些方法。因为代码中有太多灵活的逻辑，编译阶段的分析无法做到尽如人意。
	*/

	/**
	*T包含T的方法集，多出来的就是那些包装方法。
	T与*T不能定义同名方法，会造成编译器冲突，不允许。
	反射也是基于接口实现的，会被链接器认为使用了方法，从而导致不准确的结果
	*/
	var ti TInterface = &S{name: "哈哈"}
	log.Println(ti.GetName())
}
