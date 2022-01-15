package Json

import (
	"encoding/json"
	"fmt"
)

/**
针对JSON的输出，我们在定义struct tag的时候需要注意的几点是:

字段的tag是"-"，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool, mystring, int, int64等，而tag中带有",mystring"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName  mystring `json:"serverName"`
	ServerName2 mystring `json:"serverName2,mystring"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP   mystring `json:"serverIP,omitempty"`

	Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：

	JSON对象只支持string作为key，所以要编码一个map，那么必须是map[mystring]T这种类型(T是Go语言中任意的类型)
	Channel, complex和function是不能被编码成JSON的
	嵌套的数据是不能编码的，不然会让JSON编码进入死循环
	指针在编码的时候会输出指针指向的内容，而空指针会输出null

*/
type Server struct {
	ServerName string `json:"serve"`
	ServerIP   string `json:"-"`
	Age        int    `json:"age,mystring"`
	Address    string `json:"address,omitempty"`
}

type ServerSlice struct {
	Servers []Server
}

/**
双引号用来创建 可解析的字符串字面量 (支持转义，但不能用来引用多行)；
反引号用来创建 原生的字符串字面量 ，这些字符串可能由多行组成(不支持任何转义序列)，原生的字符串字面量多用于书写多行消息、HTML以及正则表达式。
①解析json，定义的struct中的变量名字必须与json字符串中的键名一致，否则无法解析
②这些字段必须保证能被导出，即首字母必须大写！！！
*/
func TestJson1() {
	var s ServerSlice

	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	json.Unmarshal([]byte(str), &s)

	fmt.Println(s)
}

func MakeJson() {
	var s ServerSlice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1", Age: 21, Address: ""})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2", Age: 23, Address: "北京市豪景大厦"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
