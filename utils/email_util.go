package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

const EmailHost = "smtp.hostinger.com"
const EmailUsername = "no-reply@easyshop-jp.com"
const EmailPassword = "No-reply1"
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

func SendEmailNotifOrder(toAdmin, toAdmin2, toCustomer string, data interface{}, trxno string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"order-notif.html",
		"[Easyshop]: New order #"+trxno,
		data,
	)
}

func SendEmailNotifApprove(toAdmin, toAdmin2, toCustomer string, data interface{}, trxno string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"order-approved-notif.html",
		"[Easyshop] Payment request for your order #"+trxno,
		data,
	)
}

func SendEmailNotifCanceled(toAdmin, toAdmin2, toCustomer string, data interface{}, trxno string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"order-canceled-notif.html",
		"[Easyshop] Cancellation of your order #"+trxno,
		data,
	)
}

func SendEmailNotifReqOrder(toAdmin, toAdmin2, toCustomer string, data interface{}, trxdate string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"req-order-received.html",
		"[Easyshop]: Request Order Received #"+trxdate,
		data,
	)
}

func SendEmailNotifReqOrderApproved(toAdmin, toAdmin2, toCustomer string, data interface{}, trxdate string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"req-order-approved.html",
		"[Easyshop]: Request Order Approved #"+trxdate,
		data,
	)
}

func SendEmailNotifReqOrderRejected(toAdmin, toAdmin2, toCustomer string, data interface{}, trxdate string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"req-order-rejected.html",
		"[Easyshop]: Request Order Rejected #"+trxdate,
		data,
	)
}

func SendEmailNotifPartnershipRequest(toAdmin, toAdmin2, toCustomer string, data interface{}, trxdate string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"partnership-request.html",
		"[Easyshop] Partnership request received #"+trxdate,
		data,
	)
}

func SendEmailNotifPartnershipApproved(toAdmin, toAdmin2, toCustomer string, data interface{}, trxdate string) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"partnership-approved.html",
		"[Easyshop] Partnership request approved #"+trxdate,
		data,
	)
}

func SendEmailNotifPartnershipReferral(toAdmin, toAdmin2, toCustomer string, data interface{}) {
	SendEmailNotif(
		toAdmin,
		toAdmin2,
		toCustomer,
		"partnership-referral.html",
		"[Easyshop] Partnership referral code generated",
		data,
	)
}

func SendEmailNotif(toAdmin, toAdmin2, toCustomer, templateName, subject string, data interface{}) {
	path, _ := os.Getwd()
	mode := os.Getenv("MODE")

	if mode == "DEV" {
		subject = "[DEV] " + subject
	}
	template := path + "/template/" + templateName
	err := SendEmail(toAdmin, subject, data, template)
	err2 := SendEmail(toAdmin2, subject, data, template)
	err3 := SendEmail(toCustomer, subject, data, template)
	if err == nil {
		fmt.Println("send email '" + subject + toAdmin + "' success")
	} else {
		fmt.Println(err)
	}
	if err2 == nil {
		fmt.Println("send email '" + subject + toAdmin2 + "' success")
	} else {
		fmt.Println(err2)
	}
	if err3 == nil {
		fmt.Println("send email '" + subject + toCustomer + "' success")
	} else {
		fmt.Println(err3)
	}
}

func SendEmail(to string, subject string, data interface{}, templateFile string) error {
	result, err := ParseTemplate(templateFile, data)
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@easyshop-jp.com")
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
