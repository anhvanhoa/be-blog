package config

import (
	"github.com/spf13/viper"
)

type MailProvider struct {
	Host     string
	Port     int
	Username string
	Email    string
	Password string
}

var Mails []MailProvider = []MailProvider{}

func InitMail() {
	err := viper.UnmarshalKey("mailserver", &Mails)
	if err != nil {
		panic(err)
	}
}
