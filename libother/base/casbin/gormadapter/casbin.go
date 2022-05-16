package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	File()
	//Db()
}

/**
文档：https://casbin.org/docs/zh-CN/syntax-for-models

1.数据库定义好casbin_rule
2.项目中定义好model
3.当有实际请求进来的时候，回去user_id,role,path,method等信息组装request_definition然后进行判断
4.casbin_rule可以再进程内部进行缓存，或者缓存到redis


*/

func File() {
	// model可以从conf，代码，字符串生成
	//e, err := casbin.NewEnforcer("./conf/rbac/model.conf", "./conf/rbac/policy.csv")
	e, err := casbin.NewEnforcer("./conf/rbac/rbac_with_resource_roles_model.conf", "./conf/rbac/policy.csv")
	// 1.角色权限可以有继承关系，比如，数学老师继承老师的权限，
	// 2.资源也可以有继承关系
	// https://github.com/casbin/casbin/tree/master/examples 所有的例子
	// casbin只负责校验权限，不负责任验证用户的有效性
	sub := "1"
	obj := "/student/search"
	act := "PUT"

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Println("出错了:", err)
		return
	}

	if ok == true {
		log.Println("通过")
		return
	} else {
		log.Println("拒绝")
		return
	}
}

func Db() {
	adapter, err := gormadapter.NewAdapter("mysql", "root:xl123456?@tcp(47.105.50.31:3306)/", false)
	e, err := casbin.NewEnforcer("./conf/rbac/model.conf", adapter)
	if err != nil {
		log.Println("err:", err)
		return
	}

	_ = e.LoadPolicy()
	sub := "1"
	obj := "/student/search"
	act := "POST"
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Println("err:", err)
		return
	}

	if ok {
		log.Println("通过")
	} else {
		log.Println("拒绝")
	}

	// 修改数据库中的权限
	//Enforcer.AddPolicy("student","/ginFrameWork/resource1","POST")//sub,obj,act
	//Enforcer.SavePolicy()
}
