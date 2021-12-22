package service_mail

import (
	gomail "gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": "xxxxxxxxx@yyyyy.zzz",
		"pass": "123456",
		"host": "smtp.exmail.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "toolKit"))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err

}
