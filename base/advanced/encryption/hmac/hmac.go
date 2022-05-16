package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/**
HMAC是密钥相关的哈希运算消息认证码，HMAC运算利用哈希算法，以一个密钥和一个消息为输入，生成一个消息摘要作为输出。

主要用于验证接口签名~
*/
func main() {
	key := "kuteng"
	data := "www.5lmh.com"
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	fmt.Println(hex.EncodeToString(hmac.Sum([]byte(""))))
}
