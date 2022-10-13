package main

/**
通过动态创建script标签,通过script标签的src,向一个不同源的接口发送一个get请求

浏览器：自定义响应回调函数，使用script标签的src请求
利用浏览器的src属性没有跨域这一限制特点
服务器：接收callback参数，返回函数调用
处理复杂，并且只支持get请求
原因：get请求参数直接在url后面拼接，而post请求参数是放在请求体中

jsonP缺点：
1.只支持GET请求，不支持POST请求。
2.jsonp在调用失败的时候不会返回各种HTTP状态码
3.不安全，返回的js内容可能被劫持
4.只支持跨域HTTP请求这种情况，不能解决不同域的两个页面之间如何进行JavaScript调用的问题
*/
func main() {

}
