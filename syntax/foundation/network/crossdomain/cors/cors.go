package main

import (
	"log"
	"net/http"
)

/**
跨源资源共享（CORS）
工作原理： 服务器 在返回响应报文的时候，在响应头中 设置一个允许的header
res.setHeader(‘Access-Control-Allow-Origin’, ‘*’)

跨域限制访问，其实是浏览器的限制

官方文档：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CORS


*/
func main() {
	server()
}

type MyHandler struct {
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method=%s,Header=%s\n", r.Method, r.Header)
	SetHeader(w, r)
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, err := w.Write(([]byte)("asdasd"))
	if err != nil {
		return
	}
}

func server() {
	MyHandler := &MyHandler{}
	server := &http.Server{
		Addr:    "127.0.0.1:9998",
		Handler: MyHandler,
	}

	_ = server.ListenAndServe()
}

func SetHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
