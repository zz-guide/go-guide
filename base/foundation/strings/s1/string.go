package main

import (
	"fmt"
)

func main() {

}

/**

几种字符串拼接方式性能测试

fmt.Sprintf
string +
strings.Join
bytes.Buffer
strings.Builder
*/

func TestString() {
	//1.string的定义：是一个字节序列（byte数组，可以是空的）, 字符串是值类型的
	//字符串一旦创建就不能改变
	name := "nihao"
	fmt.Println("name的字节数：", len(name))
	fmt.Println(name)
	fmt.Printf("name类型:%T\n", name)
	fmt.Println(name[0]) //返回第1个字节的字节值，当长度超出len(name)是会抛出panic异常，panic: runtime error: index out of range
	//结论：字符串的不可变指的就是不能通过【索引】方式去修改字符串的值，底层是只读的
	//name[0] = '2'//cannot assign to name[0]
	//fmt.Println(name)
	//ASCII字符  一个字符对应一个字节，非ASCII字符的UTF8编码的字符会要两个或多个字节
	//name = "张三"
	fmt.Println(name)
	fmt.Println("name长度：", len(name))
	//2.字符串拼接
	name1 := "李四"
	namejoin := name1 + name
	fmt.Println("name1+name拼接字符串：", namejoin)
	fmt.Printf("name1+name地址：%p\n", &namejoin)
	//3.基于切片方式创建字符串
	name3 := name[1:2]
	fmt.Println(name3) //如果是英文的话，可以取到字符数，但是中文的话不行，是乱码
	//4.字符串的比较
	name4 := "nihao"
	name5 := "nihao"
	fmt.Printf("name4地址:%p\n", &name4)
	fmt.Printf("name5地址:%p\n", &name5)
	name4 = "lisi"
	fmt.Printf("name4地址:%p\n", &name4)
	//结论：重新赋值的话地址不变，只是值发生了改变
	//5.双引号表示的字符串中支持转义字符，只有单行
	//6.原生字符（反引号包括的字符串）不支持转义字符，可以跨多行，退格，换行等都不会被转义
	fmt.Println("name4 == name5:", name4 == name5) //比较的是值
	fmt.Println("name4 < name5:", name4 < name5)   //比较的是值
	//7.Unicode编码对应rune整数类型（即int32），表示字符
	fmt.Println("\u4e16\u754c") //世界
	//8.rune,用单引号包括
	rune1 := 'c'
	fmt.Println("rune1的ASCII十进制表示：", rune1) //英文字母的时候是这样表示
	fmt.Printf("rune1的值：%v\n", rune1)
	fmt.Printf("rune1的类型：%T\n", rune1)
	fmt.Printf("rune1的值：%q\n", rune1)
	fmt.Printf("rune1的值：%#q\n", rune1)
	fmt.Printf("rune1的值：%U\n", rune1)
	fmt.Printf("rune1的值：%#U\n", rune1)
	fmt.Printf("rune1的值：%b\n", rune1)
	fmt.Printf("rune1的值：%#p\n", &rune1)
	fmt.Println("rune1的值：", string(rune1)) //必须转换为字符串才可以打印出原来的值
}

func TestRawString() {
	//1.单行的时候双引号和反引号作用一样
	rawString1 := `<div>nihao</div>`
	rawString2 := "<div>nihao</div>"
	fmt.Println(rawString1)
	fmt.Println(rawString2)
	fmt.Println("##################################")
	//2.只有反引号可以包含多行字符串，并且原格式不变
	rawString3 := `
		<div class="content">
        <div class="img"></div>
        <div class="text">
            <p>尊敬的用户，使用最新版本的谷歌浏览器，</p>
            <p>速度更快，体验更好，安全性更高哦~</p>
        </div>
        <div class="download">
            <a class="link left" href="https://www.baidu.com/link?url=yMvmPqLOG4yGd7H1TdYUIkgFjDk25GWit2FJBJmVrz1AJvSS73pkLRwPm0hw46c9C-9yFySgpJkkkdy5vQ2W5xsv86Qg4PqnC6ZULRjk-yW&wd=&eqid=f47f1bdb00013552000000035a5d6cc1">
                <i class="iconfont icon-socialwindows"></i>
                <span class="des">Windows版</span>
            </a>
            <a class="link right" href="https://www.baidu.com/link?url=ZVES5jBVnyBSpkZMWiYNVFb7x5jLTOqarmUce3marM7caQAqpaWL5OfbILxdtSoB0oul2MC0Zi8Q84HKkIYT2uE0Po_IQWIkeZ9zmx7Q9EC&wd=&eqid=f47f1bdb00013552000000035a5d6cc1">
                <i class="iconfont icon-mac"></i>
                <span class="des">Mac版</span>
            </a>
        </div>
    </div>`
	fmt.Println(rawString3)
	//3.结论：经过测试发现，rune打印出来的值是对应Unicode码对应的10进制整数值，Unicode一般使用\u是十六进制整数表示，并且是多字节的，不论是哪种编码最终同一个字符转换为10进制整数的值都是一样的
	rune1 := '世'
	rune2 := "&#19990"
	rune3 := "\u4e16"
	fmt.Println("rune1的ASCII十进制表示：", rune1)
	fmt.Println("rune2的：", rune2)
	fmt.Println("rune3的：", rune3)
}

func TansformString() {
	//1.string与byte互相转换
	name1 := "xulei"
	nameBytes := []byte(name1) //byte存的是对应ASCII表十进制数值
	//结论一：[120 117 108 101 105]，证明了一个英文字母是一个字符，占一个字节
	fmt.Println("字母string 转 []byte原值:", nameBytes)
	fmt.Printf("120:%x 117:%x 108:%x 101:%x 105:%x\n", 120, 117, 108, 101, 105)

	//2.测试汉字转化为byte
	name2 := "世"
	nameBytes2 := []byte(name2)
	//结论：证明一个汉字是一个字符，占三个字节
	fmt.Println("汉字string 转 []byte原值:", nameBytes2)
	//3.byte转string
	byte1 := []byte{228, 184, 150}
	fmt.Println("byte转string:", string(byte1))
}
