package ding

import (
	"fmt"
	"net/smtp"
	"strings"
)

const (
	HOST        = "smtp.qq.com"
	SERVER_ADDR = "smtp.qq.com:587"

	USER     = "373045134@qq.com" //发送邮件的邮箱
	PASSWORD = "asuebdqvyqdcbhbh" //发送邮件邮箱的密码
)

type Email struct {
	to      string
	subject string
	msg     string
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}

func SetMail() {
	// 邮箱地址
	UserEmail := "373045134@qq.com"
	// 端口号，:25也行
	Mail_Smtp_Port := ":587"
	//邮箱的授权码，去邮箱自己获取
	Mail_Password := "asuebdqvyqdcbhbh"
	// 此处填写SMTP服务器
	Mail_Smtp_Host := "smtp.qq.com"
	auth := smtp.PlainAuth("", UserEmail, Mail_Password, Mail_Smtp_Host)
	to := []string{"lxu@zhiniuxue.com"}
	nickname := "发送人名称"
	user := UserEmail

	subject := "我是我相信"
	content_type := "Content-Type: text/html; charset=UTF-8"

	body := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>课程表</title>
</head>
<body>
    <table border="1" width="60%" bgcolor="#e9faff" cellpadding="2">
        <caption>课程表</caption>
        <tr align="center">
            <td colspan="2">时间\日期</td>
            <td>一</td>
            <td>二</td>
            <td>三</td>
            <td>四</td>
            <td>五</td>
        </tr>

        <tr align="center">
            <td rowspan="2">上午</td>
            <td>9:30-10:15</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
        </tr>

        <tr align="center">
            <td>10:25-11:10</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
        </tr>

        <tr>
            <td colspan="7">&nbsp;</td>
        </tr>

        <tr align="center">
            <td rowspan="2">下午</td>
            <td>14:30-15:15</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
        </tr>

        <tr align="center">
            <td>15:25-16:10</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
            <td>语文</td>
        </tr>
    </table>
</body>
</html>`
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(Mail_Smtp_Host+Mail_Smtp_Port, auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}

}

func SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	sendTo := strings.Split(email.to, ";")
	done := make(chan error, 1024)

	go func() {
		defer close(done)
		for _, v := range sendTo {

			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg

			err := smtp.SendMail(
				SERVER_ADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)

			fmt.Printf("eeeeE:", err)
			done <- err
		}
	}()
	for i := 0; i < len(sendTo); i++ {
		<-done
	}
	return nil
}

func Send() {
	mycontent := "this is go test email"

	email := NewEmail("lxu@zhiniuxue.com",
		"test golang email", mycontent)
	err := SendEmail(email)
	fmt.Println(err)
}
