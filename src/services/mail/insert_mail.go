package mail

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/models"
)

func InsertMail(body models.MailHistory) (err error) {
	var mailHistory entities.MailHistory = entities.MailHistory{
		Id:           body.Id,
		MailProvider: body.MailProvider,
		ReceivedMail: body.ReceivedMail,
		SenderBy:     body.SenderBy,
		Data:         body.Data,
		Error:        body.Error,
		Template:     body.Template,
		Status:       body.Status,
	}
	_, err = config.DB.Model(&mailHistory).Insert()
	if err != nil {
		return err
	}
	return nil
}
