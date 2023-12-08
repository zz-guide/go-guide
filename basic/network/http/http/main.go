package main

import (
	"log"
	"net/http"
	"runtime"
	"time"
)

/**

Content-Type: application/x-www-form-urlencoded
普通的表单
形式：id=3&age=22  键值对用&链接
id=%E8%AE%B8%E7%A3%8A&age=22
遇见汉字会被编码
Content-Length代表的是字符的个数或者叫字节数也没错，不仅包含键值对，还包含链接的字符
Content-Type: multipart/form-data; boundary=--------------------------232316194004562890184642
附带文件的表单
形式：
----------------------------232316194004562890184642
Content-Disposition: form-data; name="id"

3
----------------------------232316194004562890184642
Content-Disposition: form-data; name="age"

23
----------------------------232316194004562890184642--
键值对用boundary隔开，Content-length也会变大


Content-Type: text/plain    body=raw类型
原封不动
形式：{"id":"1","age":"33"}

// 没有 Content-Type 表示body类型为none


Content-Type:application/octet-stream,从字面意思得知，只可以上传二进制数据，通常用来上传文件，
binary

*/

/**
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

*/

func main() {
	//DefaultServer()
	//CustomerServer()
	TimeoutServer()
}

type MyHandler struct {
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("example 请求")
	log.Println("Method:", r.Method)
	log.Println("Header:", r.Header)
}

func CustomerServer() {
	MyHandler := &MyHandler{}
	server := &http.Server{
		Addr:    "127.0.0.1:9998",
		Handler: MyHandler,
	}

	_ = server.ListenAndServe()
}

func DefaultServer() {
	// ExampleHandler 这个函数是一个type，这个type实现了ServeHTTP方法，这个方法默认是通过ServeHTTP处理的
	http.HandleFunc("/example", ExampleHandler)
	log.Println("server start 127.0.0.1:9998")
	_ = http.ListenAndServe("127.0.0.1:9998", nil)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	//_ = r.ParseForm() // Form字段包含了URL上的参数以及body
	// 只解析body=form-data形式，不包含URL上的参数，ParseMultipartForm包含了ParseForm的功能
	_ = r.ParseMultipartForm(32 << 20)
	log.Println("example 请求")
	log.Println("Method:", r.Method)
	log.Println("Header:", r.Header)
	log.Println("ContentLength:", r.ContentLength)
	log.Println("Path:", r.URL.Query()) // get参数
	// 只解析body=x-www-fomr-urlencoded形式,但是包含了URL上的参数
	log.Println("Form:", r.Form)
	// 只解析body=x-www-fomr-urlencoded形式, 不包含URL上的参数
	log.Println("PostForm:", r.PostForm)

	if r.MultipartForm != nil {
		log.Println("MultipartForm-Value:", r.MultipartForm.Value)
		log.Println("MultipartForm-File:", r.MultipartForm.File)
	}

	// raw形式的参数都在body中
	body := make([]byte, r.ContentLength)
	// 读取 r 的请求主体，并将具体内容读入 body 中
	r.Body.Read(body)
	log.Println("body:", string(body))
}

func TimeoutServer() {
	log.Println("server start 127.0.0.1:9998")
	_ = http.ListenAndServe("127.0.0.1:9998", http.TimeoutHandler(
		http.HandlerFunc(TimeHandler),
		2*time.Second,
		"超时返回!\n"),
	)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("----进入----")
	go func() {
		for {
			select {
			case <-r.Context().Done():
				log.Println("---退出---")
				return
			}
		}
	}()

	// 此处time sleep已经超时。不应该继续向下执行
	time.Sleep(time.Second * 5)
	// 模拟数据库操作
	log.Println("student表插入一条数据")
	log.Println("goroutinue数量:", runtime.NumGoroutine())
}
