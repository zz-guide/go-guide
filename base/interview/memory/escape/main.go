package main

// go run -gcflags "-m -l" main.go (-m打印逃逸分析信息，-l禁止内联编译)
// go tool compile -S main.go

//Go官网上有一段可以表明分配的规则：
//
//准确来说，你并不需要知道。Golang 中的变量只要被引用就一直会存活，
//存储在堆上还是栈上由内部实现决定而和具体的语法没有关系。知道变量的存储位置确实对程序的效率有帮助。
//如果可能，Golang 编译器会将函数的局部变量分配到函数栈帧（stack frame）上。然而，
//如果编译器不能确保变量在函数 return 之后不再被引用，编译器就会将变量分配到堆上。而且，
//如果一个局部变量非常大，那么它也应该被分配到堆上而不是栈上。当前情况下，如果一个变量被取地址，
//那么它就有可能被分配到堆上。然而，还要对这些变量做逃逸分析，如果函数 return 之后，变量不再被引用，
//则将其分配到栈上。

// 逃逸情况
// 1.make声明slice，map，channel
// 2.interface类型赋值
// 3.可变参数，参数是引用类型等
// 4.fmt,log包打印函数调用
// 5.如果一个局部变量非常大，无法放在栈上，也会被放到堆上。
// 6.如果一个局部变量在函数return之后没有被引用了，就会将其放到栈上。反之则在堆上
// 7.如果局部变量在函数外被引用了，变量会被分配到堆上。

func main() {
	//_ = tArray()
	//_ = tSlice()
	//_ = tMap()
	//_ = tStr()
	//_ = tInt()
	//_ = tAny()
	//_ = tStrPointer()
	//_ = tIntPointer()
	//tChannel()
	tArgs()
}

func tArray() [2]int {

	aArray := [2]int{1, 2}
	//aArray := make([2]int{}, 2)
	return aArray
}

func tSlice() []int {
	// slice不论大小，一律逃逸
	// make([]int, 10000) escapes to heap
	a := make([]int, 1)
	return a
}

func tMap() map[string]string {
	// escapes to heap
	a := make(map[string]string)
	return a
}

func tStr() string {
	a := "hello world"
	return a
}

func tInt() int {
	return 0
}

func tIntPointer() *int {
	aInt := 1
	return &aInt
}

func tStrPointer() *string {
	// moved to heap: a 表示直接分配在堆上
	aStr := "hello world"
	return &aStr
}

func tAny() any {
	// 赋值给interface变量一定会逃逸
	aAny := "hello world"
	return aAny
}

func tChannel() {
	var (
		chInteger   = make(chan *int)
		chMap       = make(chan map[int]int)
		chSlice     = make(chan []int)
		chInterface = make(chan interface{})
		a, b, c, d  = 0, map[int]int{}, []int{}, 32
	)
	chInteger <- &a  // 逃逸
	chMap <- b       // 逃逸
	chSlice <- c     // 逃逸
	chInterface <- d // 逃逸
}

func tArgs() {
	aFunc1 := func(arg ...int) []int {
		//arg2 := arg
		//return arg2
		return nil
	}
	aFunc1(1, 2, 3) //  func literal 函数字面量

	aFunc2 := func(arg []int) []int {
		//arg2 := arg
		//return arg2
		return nil
	}
	aFunc2([]int{1, 2, 3})
}
