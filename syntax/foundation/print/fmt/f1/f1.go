package main

import "fmt"

type user struct {
	Name string
	Age  int
}

/**
  结论：1.fmt包是非线程安全的
2.有Print开头的打印函数
3.有Scan开头的格式化输入函数,例如:Scan,Scanf,Scanln
4.有S开头的输出到字符村的函数，例如：Sprint，Sprintln，Sprintf
5.有F开头的输出到文件的函数，例如：Fprint,Fprintf
5.Errorf,功能同 Sprintf，只不过结果字符串被包装成了 error 类型。
6.String() 用于对结构体的标准输出等。
Error() 封装error的方法，可以改一些错误上传到日志系统或者打印Stack。
Format() 对于String()的高级用法，用于多种类型或者格式使用。
GoString() 常用于相对值。
根据顺序调用，Format->GoString->Error->String
7.不能Fatal
注意：通过fmt包打印的参数可能发生逃逸

内置的print/println函数总是写入标准错误。
fmt标准包里的打印函数总是写入标准输出。
log标准包里的打印函数会默认写入标准错误，然而也可以通过log.SetOutput函数来配置。
内置print/println函数的调用不能接受数组和结构体参数。
对于组合类型的参数，内置的print/println函数将输出参数的底层值部的地址，而fmt和log标准库包中的打印函数将输出参数的字面值。目前（Go 1.12），
对于标准编译器，调用内置的print/println函数不会使调用参数引用的值逃逸到堆上，而fmt和log标准库包中的的打印函数将使调用参数引用的值逃逸到堆上。如果一个实参有String() string或Error() string方法，那么fmt和log标准库包里的打印函数在打印参数时会调用这两个方法，而内置的print/println函数则会忽略参数的这些方法。内置的print/println函数不保证在未来的Go版本中继续存在。
*/
func main() {
	Stringer()
}

func PrintF() {
	userInfo := user{
		Name: "Bill",
		Age:  25,
	}
	// 区别 ：“%+v”会以字段键值对的形式key-value格式打印，“%v”只会打印字段值value信息
	// 结构体打印(json格式等...)
	fmt.Printf("%+v\n", userInfo) // {Name:Bill Age:25}
	fmt.Printf("%#v\n", userInfo) // main.user{Name:"Bill", Age:25} #包名也会打印出来
	fmt.Printf("%v\n", userInfo)  // {Bill 25}
}

func Print() {
	userInfo := user{
		Name: "Bill",
		Age:  25,
	}

	// 可以把不同类型的参数拼接在一起当成字符串
	fmt.Print(12, "%s6 你好", userInfo, "\n")
}

func PrintLn() {
	userInfo := user{
		Name: "Bill",
		Age:  25,
	}

	// 可以把不同类型的参数拼接在一起当成字符串
	fmt.Println(12, "%s6 你好", userInfo, "\n", "asd")
}

type Animal struct {
	Name string
	Age  uint
}

func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func (a Animal) GoString() string {
	return "哈哈"
}

func (a Animal) Error() string {
	return "Error"
}

func (t Animal) Format(s fmt.State, c rune) {
	switch c {
	case 'c':
		switch {
		case s.Flag('+'):
			fmt.Printf("我是+c\n")
		default:
			fmt.Fprint(s, "我是c\n")
		}
	default:
		fmt.Print("我是Format")
	}
}

func Stringer() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
	fmt.Printf("%#v", a)
}
