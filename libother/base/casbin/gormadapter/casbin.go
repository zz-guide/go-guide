package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	File()
}

func File() {
	// model可以从conf，代码，字符串生成
	e, err := casbin.NewEnforcer("./conf/rbac/model.conf", "./conf/rbac/policy.csv")
	sub := "alice"
	obj := "book/3"
	act := "read"

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Println("出错了:", err)
		return
	}

	if ok == true {
		fmt.Println("成功")
		return
	} else {
		fmt.Println("不通过")
		return
	}
}

func Db() {
	a, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	e, err := casbin.NewEnforcer("./conf/nouseracl/model.conf", a)

	enforce, err := e.Enforce("data2", "read")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if enforce {
		fmt.Println("通过")
	} else {
		fmt.Println("拒绝")
	}
}
