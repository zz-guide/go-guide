package route

import (
	"net/http"
	"src/sql"
)

//定义公共的header头，用于返回json
var jsonContentType = []string{
	"application/json; charset=utf-8",
}

//设置允许跨域
func writeContentType(w http.ResponseWriter, value []string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

/**
提供http服务
*/
func HttpRun(writer http.ResponseWriter, req *http.Request) {
	//设置允许跨域和，响应头json格式
	writeContentType(writer, jsonContentType)
	writer.Write(sql.StartServer())
}

/**
本地客户端跑
*/
func Run() {
	sql.StartClient()
}
