package ding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SendDingMsg(msg string) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=eac5d5c15a7ef587d1887e8b87e71da28093895f42a369d8208c43f7dad3969b`
	content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"},
		"at": {
        	"isAtAll": false,
		"atMobiles":[
            		"18810951239"
        	],
    	},
	}`

	Secret := "SEC74abf90ef8bda396557af1a6040fe5fcc800fbc1401a76e6358ee6c763417aff"
	//  构建 签名
	//  把timestamp+"\n"+密钥当做签名字符串，使用HmacSHA256算法计算签名，然后进行Base64 encode，最后再把签名参数再进行urlEncode，得到最终的签名（需要使用UTF-8字符集）。
	timeStampNow := time.Now().UnixNano() / 1e6 // 毫秒
	signStr := fmt.Sprintf("%d\n%s", timeStampNow, Secret)

	hash := hmac.New(sha256.New, []byte(Secret))
	hash.Write([]byte(signStr))
	sum := hash.Sum(nil)

	encode := base64.StdEncoding.EncodeToString(sum)
	urlEncode := url.QueryEscape(encode)

	// 构建 请求 url
	webHook = fmt.Sprintf("%s&timestamp=%d&sign=%s", webHook, timeStampNow, urlEncode)

	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		return
	}
	client := &http.Client{Timeout: time.Second * 12}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("asdaa:", err)
		return
	}

	fmt.Printf("res:%v+", resp)

	defer resp.Body.Close()
}
