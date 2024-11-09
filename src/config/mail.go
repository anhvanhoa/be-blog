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
	{
		Id:       "COMMENT",
		Template: "src/templates/comment.html",
		Data:     map[string]interface{}{"nameBlog": "Bài viết mới", "fullName": "Người bình luận", "content": "Nội dung bình luận", "link": "http://localhost:3000/blog/1"},
		Subject:  "Bình luận mới: {{nameBlog}}",
	},
	{
		Id:       "COMMENT_REPLY",
		Template: "src/templates/comment_reply.html",
		Data:     map[string]interface{}{"nameBlog": "Trả lời bình luận bài viết", "fullName": "Người bình luận", "content": "Nội dung bình luận", "link": "http://localhost:3000/blog/1"},
		Subject:  "Trả lời bình luận bài viết",
	},
	{
		Id:       "FORGOT_PASSWORD",
		Template: "src/templates/forgot_password.html",
		Data:     map[string]interface{}{"link": viper.GetString("hostClient") + "/forgot-password/token"},
		Subject:  "Yêu cầu đổi mật khẩu",
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
