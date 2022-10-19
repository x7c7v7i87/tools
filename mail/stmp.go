package mail

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

var (
	fromAddress string
	fromName    string
	fromHost    string
	fromPort    int
	fromUname   string
	fromPasswd  string
)

type Email struct {
	to      string "to"
	subject string "subject"
	body    string "msg"
}

func NewEmail(to, subject, body string) *Email {
	return &Email{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (email *Email) Send() (bool, error) {

	m := gomail.NewMessage()

	// 发件人
	m.SetAddressHeader("From", fromAddress, fromName)

	// 收件人
	m.SetHeader("To", m.FormatAddress(email.to, email.to))
	//抄送
	//m.SetHeader("Cc", m.FormatAddress("xxxx@foxmail.com", "收件人"))
	//暗送
	//m.SetHeader("Bcc", m.FormatAddress("xxxx@gmail.com", "收件人"))
	// 主题
	m.SetHeader("Subject", email.subject)

	// 可以放html..还有其他的
	m.SetBody("text/html", email.body)

	//add Attach
	//m.Attach("Attach")

	// send email server ,port,send from id, from passwd
	d := gomail.NewDialer(fromHost, fromPort, fromUname, fromPasswd)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("发送失败", err)
		return false, err
	}

	return true, nil
}
