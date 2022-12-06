package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

const EmailHost = "smtp.gmail.com"
const EmailUsername = "easyshop011022@gmail.com"
const EmailPassword = "aosvojlffmutdlqo"
const EmailPort = 587

func SendEmailOtp(to string, data interface{}) {
	var err error
	path, _ := os.Getwd()
	template := path + "/template/send-otp.html"
	subject := "Easy Shop Email Registration"
	err = SendEmail(to, subject, data, template)
	if err == nil {
		fmt.Println("send email '" + subject + "' success")
	} else {
		fmt.Println(err)
	}
}

func SendEmail(to string, subject string, data interface{}, templateFile string) error {
	result, err := ParseTemplate(templateFile, data)
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "support@easyshop.jp")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	d := gomail.NewDialer(EmailHost, EmailPort, EmailUsername, EmailPassword)
	err = d.DialAndSend(m)
	return err
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}
