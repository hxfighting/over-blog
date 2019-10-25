package service

import (
	"blog/config"
	"gopkg.in/gomail.v2"
)

type Email struct {
	To          string
	ContentType string
	Body        string
}

var (
	EmailService     = newEmail()
	EmailChan        = make(chan int64, 10)
	ReplyEmailChan   = make(chan int64, 10)
	CommentEmailChan = make(chan int64, 10)
)

func newEmail() *gomail.Dialer {
	host := config.GetConfig("email.host").(string)
	port := int(config.GetConfig("email.port").(int64))
	username := config.GetConfig("email.username").(string)
	pass := config.GetConfig("email.password").(string)
	email := gomail.NewDialer(host, port, username, pass)
	return email
}

/**
发送邮件
*/
func (e *Email) SendEmail() bool {
	m := gomail.NewMessage()
	from := config.GetConfig("email.from").(string)
	if from == "" {
		Log.Error("缺少发送方邮件地址")
		return false
	}
	subject := config.GetConfig("email.subject").(string)
	if subject == "" {
		Log.Error("缺少邮件主题")
		return false
	}
	m.SetHeader("From", from)
	m.SetHeader("To", e.To)
	m.SetHeader("Subject", subject)
	m.SetBody(e.ContentType, e.Body)
	if err := EmailService.DialAndSend(m); err != nil {
		Log.Error(err.Error())
		return false
	} else {
		return true
	}
}
