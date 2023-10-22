package dto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/spf13/viper"
)

var SMTP_AUTH smtp.Auth
var PATH_TEMPLATE_VERIFICATION_MAIL = "docs/template"

type RequestEmail struct {
	to      []string
	subject string
	body    string
}

func NewRequestEmail(to []string, subject, body string) *RequestEmail {
	return &RequestEmail{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *RequestEmail) SendEmail() (bool, error) {
	addr := fmt.Sprintf("%s:%s", viper.GetString("SMTP_HOST"), viper.GetString("SMTP_PORT"))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from := fmt.Sprintf("From: %s\n", viper.GetString("APP_EMAIL"))
	to := fmt.Sprintf("To: %s\n", r.to[0])
	subject := fmt.Sprintf("Subject: %s\n", r.subject)
	msg := []byte(from + to + subject + mime + "\n" + r.body)

	if err := smtp.SendMail(addr, SMTP_AUTH, viper.GetString("APP_EMAIL"), r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *RequestEmail) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
