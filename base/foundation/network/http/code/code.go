package main

import (
	"log"
	"net/http"
)

/**
http code码
*/
func main() {
	server()
}

type MyHandler struct {
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method=%s,Header=%s\n", r.Method, r.Header)

	w.Header().Set("Content-Type", "application/json")
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

func StartWith100(w http.ResponseWriter, r *http.Request) {

}

// StartWith200 成功
// 200开头的code码
///**
func StartWith200(w http.ResponseWriter, r *http.Request) {

}

// StartWith300 重定向
// 300开头的code码
///**
func StartWith300(w http.ResponseWriter, r *http.Request) {
	// 301 永久重定向
	// 302 临时重定向
}

// StartWith400
// 400开头的code码 客户端错误
///**
func StartWith400(w http.ResponseWriter, r *http.Request) {
	// 405 请求方法不允许
}

// StartWith500
// 500开头的code码 服务端错误
///**
func StartWith500(w http.ResponseWriter, r *http.Request) {

}
