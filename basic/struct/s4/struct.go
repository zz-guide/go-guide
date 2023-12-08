package mystruct

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	string
	UserId   int    `json:"user_id" bson:"b_user_id"`
	UserName string `json:"user_name" bson:"b_user_name"`
}

func TestTag() {
	//结论一：打印结构体的值的顺序是根据定义时候来的，不是赋值时的顺序
	//定义方式一
	u := &User{UserId: 1, UserName: "tony", string: "sss"}
	//定义方式二
	u1 := new(User)
	fmt.Println(u)
	fmt.Println(u1)
	j, _ := json.Marshal(u)
	fmt.Println(string(j))

	// 获取tag中的内容
	t := reflect.TypeOf(u)
	fmt.Println(t)
	fmt.Println(t.Elem())
	/*
		type StructField struct {
			Name string //属性名字
			PkgPath string//当属性名字小写的时候，是包名
			Type      Type      // 类型
			Tag       StructTag // field tag string
			Offset    uintptr   // 占用的字节数
			Index     []int     // 第几个位置
			Anonymous bool      // 是不是匿名属性
		}

	*/
	//Go语言中struct的属性可以没有名字而只有类型，使用时类型即为属性名。（因此，一个struct中同一个类型的匿名属性只能有一个）
	field := t.Elem().Field(0)
	fmt.Println(field)
	fmt.Println(field.Tag.Get("json"))
	// 输出：user_id
	fmt.Println(field.Tag.Get("bson"))
}
