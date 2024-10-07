package config

import (
	"github.com/spf13/viper"
)

type MailProvider struct {
	Host         string
	Port         int
	Username     string
	Email        string
	NameProvider string
	Password     string
}

type MailTemplate struct {
	Id       string
	Template string
	Data     map[string]interface{}
	Subject  string
}

var MailProviders map[string]MailProvider = make(map[string]MailProvider)

var MailTemplates []MailTemplate = []MailTemplate{
	{
		Id:       "REGISTER",
		Template: "src/templates/register.html",
		Data:     map[string]interface{}{"code": "123456", "fullName": "Anh Nguyen"},
		Subject:  "Xác thực tài khoản",
	},
}

func InitMail() {
	var Mails []MailProvider = []MailProvider{}
	err := viper.UnmarshalKey("mailserver", &Mails)
	if err != nil {
		panic(err)
	}
	for _, mail := range Mails {
		MailProviders[mail.Email] = mail
	}
}
