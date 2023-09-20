package main

import (
	"log"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {

}

func strSplice() {
	// go字符串默认是utf-8编码的

	str := "hello world 你好"
	// 字符串通过切片方式截取，返回的还是字符串
	log.Println("str[:]", str[:], str[:] == str)
	log.Println("str[0]", str[0])

	// len 参数是字符串，返回字节数组的长度
	bytesLength := len(str)
	// 依据下标取字符串中的字符，类型为byte,也是uint8
	for i := 0; i < bytesLength; i++ {
		ch := str[i]
		log.Println(i, ch)
	}

	// 字符串转成rune数组之后,再通过len获取长度为字符长度
	runs := []rune(str)
	charLength := len(runs)
	for i := 0; i < charLength; i++ {
		ch := runs[i]      // 依据下标取字符串中的字符，类型为rune
		log.Println(i, ch) //unicode 编码转十进制输出
	}
}

func strImmutable() {
	// string的定义：是一个字节序列（byte数组，可以是空的）, 字符串是值类型的
	// 字符串的不可变指的就是不能通过【索引】方式去修改字符串的值，底层是只读的
	// ASCII字符  一个字符对应一个字节，非ASCII字符的UTF8编码的字符会要两个或多个字节
	// 字符串不可变，并且是UTF8编码的
	s1 := "hello"
	s2 := "hell"

	log.Printf("s1:%#v\n", *(*reflect.StringHeader)(unsafe.Pointer(&s1)))
	log.Printf("s2:%#v\n", *(*reflect.StringHeader)(unsafe.Pointer(&s2)))

	c := []byte(s1)
	c[2] = 'n'
	s1 = string(c)
	log.Printf("s1:%#v\n", *(*reflect.StringHeader)(unsafe.Pointer(&s1)))

	//var sum *int = new(int) //分配空间
	//*sum = 98
	//log.Println(*sum)

	// 字符串的比较
	name4 := "nihao"
	name5 := "nihao"
	log.Printf("name4地址:%p\n", &name4)
	log.Printf("name5地址:%p\n", &name5)
	name4 = "lisi"
	log.Printf("name4地址:%p\n", &name4)
	// 重新赋值的话地址不变，只是值发生了改变
}

func tRune() {
	// 双引号表示的字符串中支持转义字符，只有单行
	// 原生字符（反引号包括的字符串）不支持转义字符，可以跨多行，退格，换行等都不会被转义
	// Unicode编码对应rune整数类型（即int32），表示字符
	log.Println("\u4e16\u754c") //世界
	// rune,用单引号包括
	rune1 := 'c'
	log.Println("rune1的ASCII十进制表示：", rune1) //英文字母的时候是这样表示
	log.Printf("rune1的值：%v\n", rune1)
	log.Printf("rune1的类型：%T\n", rune1)
	log.Printf("rune1的值：%q\n", rune1)
	log.Printf("rune1的值：%#q\n", rune1)
	log.Printf("rune1的值：%U\n", rune1)
	log.Printf("rune1的值：%#U\n", rune1)
	log.Printf("rune1的值：%b\n", rune1)
	log.Printf("rune1的值：%#p\n", &rune1)
	log.Println("rune1的值：", string(rune1)) //必须转换为字符串才可以打印出原来的值
}

func tRawString() {
	// 单行的时候双引号和反引号作用一样
	rawString1 := `<div>nihao</div>`
	rawString2 := "<div>nihao</div>"
	log.Println(rawString1)
	log.Println(rawString2)
	log.Println("##################################")
	// 只有反引号可以包含多行字符串，并且原格式不变
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
	log.Println(rawString3)
	// 经过测试发现，rune打印出来的值是对应Unicode码对应的10进制整数值，Unicode一般使用\u是十六进制整数表示，并且是多字节的，
	// 不论是哪种编码最终同一个字符转换为10进制整数的值都是一样的
	rune1 := '世'
	rune2 := "&#19990"
	rune3 := "\u4e16"
	log.Println("rune1的ASCII十进制表示：", rune1)
	log.Println("rune2的：", rune2)
	log.Println("rune3的：", rune3)
}

func tStrTransform() {
	// string与byte互相转换
	name1 := "xulei"
	nameBytes := []byte(name1) // byte存的是对应ASCII表十进制数值
	// 120 117 108 101 105]，证明了一个英文字母是一个字符，占一个字节
	log.Println("字母string 转 []byte原值:", nameBytes)
	log.Printf("120:%x 117:%x 108:%x 101:%x 105:%x\n", 120, 117, 108, 101, 105)

	// 测试汉字转化为byte
	name2 := "世"
	nameBytes2 := []byte(name2)
	// 证明一个汉字是一个字符，占三个字节
	log.Println("汉字string 转 []byte原值:", nameBytes2)
	// byte转string
	byte1 := []byte{228, 184, 150}
	log.Println("byte转string:", string(byte1))
}

func F1() {
	s := "你好"
	log.Printf("s6:%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	log.Printf("s6:%#v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func F2() {
	s := "我hello"
	// 结论：对于字符串，len函数返回的是字节长度;utf8.RuneCountInString返回字符串的长度
	log.Println("字符串s长度：", len(s), utf8.RuneCountInString(s))

	a := s
	log.Println("字符串a长度：", len(a), utf8.RuneCountInString(a), len([]rune(a)))
	// 结论：a[0]表示的是第一个字节
	log.Println("字符串a[]：", a[0], a[1], a[2], a[3])
	log.Println("字符串a[]：", string(a[0]), string(a[1]), string(a[2]), string(a[3]))

	// 结论：字符串不能通过a[2] = ''的形式修改，但是可以重新赋值
	a = "asdas"
	// a[2] = 's6'
	log.Println("字符串a长度：", len(a), utf8.RuneCountInString(a))

	// 结论：字符串遍历出来的也是数字
	for i, i2 := range a {
		log.Println("i:", i)
		log.Println("i2:", i2)
	}
}

func F3() {
	// 修改字符串
	s := "hello 世界！"
	b := &s
	log.Println("字符串s长度：", &s, len(s), (*reflect.StringHeader)(unsafe.Pointer(&s)))
	r := []rune(s)
	r[6] = '中'
	s = string(r)
	log.Println("字符串a长度：", &s, len(s), (*reflect.StringHeader)(unsafe.Pointer(&s)))
	log.Println(&b, b, *b)
}
