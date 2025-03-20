package mail

import (
	"be-blog/src/config"
	"be-blog/src/constants"
	"be-blog/src/models"
	mailService "be-blog/src/services/mail"
	"crypto/tls"
	"fmt"
	"net/mail"

	"github.com/flosch/pongo2/v6"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type PayloadMail struct {
	Tos      []string
	Template string
	Data     map[string]interface{}
	From     string
}

func SendMailFile(payload PayloadMail) {
	local := pongo2.Must(pongo2.FromFile(payload.Template))
	html, err := local.Execute(payload.Data)
	if err != nil {
		panic(err)
	}
	payload.Template = html
	SendMail(payload)
}

func GetTemplate(ID string) config.MailTemplate {
	for _, template := range config.MailTemplates {
		if template.Id == ID {
			return template
		}
	}
	return config.MailTemplate{}
}

func RenderTemplateFile(file string, data map[string]interface{}) (string, error) {
	template, err := pongo2.FromFile(file)
	if err != nil {
		return "", err
	}
	html, err := template.Execute(data)
	if err != nil {
		return "", err
	}
	return html, nil
}
func RenderTemplateString(body string, data map[string]interface{}) (string, error) {
	template, err := pongo2.FromString(body)
	if err != nil {
		return "", err
	}
	html, err := template.Execute(data)
	if err != nil {
		return "", err
	}
	return html, nil
}

func SendMail(payload PayloadMail) {
	id, _ := uuid.NewV7()
	template := GetTemplate(payload.Template)
	mailHistory := models.MailHistory{
		Id:           id.String(),
		MailProvider: payload.From,
		ReceivedMail: "unknown",
		SenderBy:     "system",
		Data:         payload.Data,
		Template:     template.Id,
		Status:       constants.STATUS_MAIL_SUCCESS,
		Error:        "",
	}
	subject, err := RenderTemplateString(template.Subject, payload.Data)
	if err != nil {
		mailHistory.Error = err.Error()
		mailHistory.Status = constants.STATUS_MAIL_FAILED
		err = mailService.InsertMail(mailHistory)
		if err != nil {
			panic(err)
		}
		return
	}
	html, err := RenderTemplateFile(template.Template, payload.Data)
	if err != nil {
		mailHistory.Error = err.Error()
		mailHistory.Status = constants.STATUS_MAIL_FAILED
		err = mailService.InsertMail(mailHistory)
		if err != nil {
			panic(err)
		}
		return
	}
	payload.Template = html
	mailProvider := config.MailProviders[payload.From]
	dialer := gomail.NewDialer(mailProvider.Host, mailProvider.Port, mailProvider.Username, mailProvider.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	from := mail.Address{
		Name:    mailProvider.NameProvider,
		Address: mailProvider.Email,
	}
	m.SetHeader("From", from.String())
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", payload.Template)
	for _, to := range payload.Tos {
		m.SetHeader("To", to)
		mailHistory.ReceivedMail = to
		if err := dialer.DialAndSend(m); err != nil {
			mailHistory.Error = err.Error()
			mailHistory.Status = constants.STATUS_MAIL_FAILED
			err = mailService.InsertMail(mailHistory)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Send mail to: ", to)
	}
	err = mailService.InsertMail(mailHistory)
	if err != nil {
		panic(err)
	}
}
