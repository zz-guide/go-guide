package main

import "fmt"

/**
	结论：1.在Go语言里的 switch 语句中，是不需要使用 break 来退出一个case的。也就是说，case执行完成后，是不会继续向下匹配的。
		 2.可以使用关键字 fallthrough 进行执行下一个case，且fallthrough不会判断下一个case的条件。话句话说，不论下一个case是否被匹配，都会被执行
         3.case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行
		 4.fallthrough只能用在基础类型的case语句,cannot fallthrough in type switch
*/
func main() {
	F2()
}

func fallthroughExample() {
	p := fmt.Println

	a := 1
	switch a {
	case 0:
		p(0)
	case 1:
		p(1)
		fallthrough
	case 2:
		p(2) //因为在上一个case中使用了fallthrough关键字，这个case就算条件不匹配也会被执行
	case 3:
		p(3)
	default:
		p("default")
	}
}

func F1() {
	var x interface{}
	var y = 10
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是：%T", i)
	case int:
		fmt.Printf("x是 int 类型")
	case float64:
		fmt.Printf("x是 float64 类型")
	case func(int) float64:
		fmt.Printf("x是func(int)类型")
	case bool, string:
		fmt.Printf("x是bool或者string类型")
	default:
		fmt.Printf("未知型")
	}
}

func F2() {
	month := 99
	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误...")
	}
}
