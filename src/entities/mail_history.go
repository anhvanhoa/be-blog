package entities

type MailHistory struct {
	tableName    struct{}               `pg:"mail_histories"`
	Id           string                 `pg:"id,pk"`
	MailProvider string                 `pg:"mail_provider"`
	ReceivedMail string                 `pg:"received_mail"`
	SenderBy     string                 `pg:"sender_by"`
	Data         map[string]interface{} `pg:"data"`
	Template     string                 `pg:"template"`
	Error        string                 `pg:"error"`
	Status       string                 `pg:"status"`
	CreatedAt    string                 `pg:"created_at"`
}
