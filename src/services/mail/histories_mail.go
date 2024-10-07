package mail

import (
	"be-blog/src/config"
	"be-blog/src/entities"
)

func HistoriesMail() (mailHistories entities.MailHistory, err error) {
	err = config.DB.Model(&mailHistories).Select()
	if err != nil {
		return mailHistories, err
	}
	return mailHistories, nil
}
