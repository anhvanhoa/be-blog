package models

type MailHistory struct {
	Id           string
	MailProvider string
	ReceivedMail string
	SenderBy     string
	Data         map[string]interface{}
	Error        string
	Template     string
	Status       string
}
