package command

import (
	"fmt"
	"mysqldoc/config"
	"mysqldoc/route"
	"net/http"
	"os"
)

var Model int

const MODEL_SEVER = 1
const MODEL_CLIENT = 2

func Start() {
	//支持两种模式：model-one:http服务；model-two：本地json文件
	//fmt.Printf("请选择模式：1-服务模式；2-本地模式")
	//fmt.Scanln(&Model)

	Model = 2
	if Model == MODEL_SEVER {
		StdInput()
		RunServer()
		fmt.Println("请求地址默认是http://localhost:3000")
	} else if Model == MODEL_CLIENT {
		StdInput()
		RunClient()
	} else {
		fmt.Println("请输入正确的模式编号")
		os.Exit(1)
	}
}

func StdInput() {
	fmt.Printf("默认参数为127.0.0.1 root 123456 bingo" +
		"\n请依次正确输入连接mysql的host,username,password,port,database以空格分隔：")
	fmt.Scanln(&config.Host, &config.Username, &config.Password, &config.Port, &config.Database)

	fmt.Println("连接主机：", config.Host)
	fmt.Println("用户名：", config.Username)
	fmt.Println("密码：", config.Password)
	fmt.Println("端口：", config.Port)
	fmt.Println("数据库：", config.Database)
}

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", route.HttpRun)
	http.ListenAndServe(":3000", mux)
}

func RunClient() {
	route.Run()
}
