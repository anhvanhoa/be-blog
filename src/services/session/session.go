package session

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/models"
)

func CreateSession(body models.Session) (err error) {
	var session entities.Session = entities.Session{
		UserID:    body.UserID,
		ExpiredAt: body.ExpiredAt,
		Token:     body.Token,
		IP:        body.IP,
	}
	_, err = config.DB.Model(&session).Insert()
	if err != nil {
		return err
	}
	return nil
}

func GetSessionByToken(token string) (session entities.Session, err error) {
	err = config.DB.Model(&session).Where("token = ?", token).Select()
	if err != nil {
		return session, err
	}
	return session, nil
}

func DeleteSessionByToken(token string) (err error) {
	session, err := GetSessionByToken(token)
	if err != nil {
		return err
	}
	_, err = config.DB.Model(&session).Where("token = ?", token).Delete()
	if err != nil {
		return err
	}
	return nil
}
