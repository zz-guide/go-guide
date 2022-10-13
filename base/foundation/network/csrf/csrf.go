package main

/**

CSRF（Cross-site request forgery），中文名称：跨站请求伪造，也被称为：one click attack/session riding，缩写为：CSRF/XSRF。
攻击者可以盗用你的登陆信息，以你的身份模拟发送各种请求。

从上图可以看出，要完成一次CSRF攻击，受害者必须依次完成两个步骤 ：
登录受信任网站A，并在本地生成Cookie 。
在不退出A的情况下，访问危险网站B。
看到这里，读者也许会问：“如果我不满足以上两个条件中的任意一个，就不会受到CSRF的攻击”。是的，确实如此，但你不能保证以下情况不会发生：

防御CSRF：
1.验证 HTTP Referer 字段；
2.在请求地址中添加 token 并验证；
3.在 HTTP 头中自定义属性并验证。

4.正确使用GET,POST和Cookie；
5.在非GET请求中增加伪随机数；
*/
func main() {

}
