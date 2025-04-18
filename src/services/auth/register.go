package auth

import (
	"be-blog/src/config"
	"be-blog/src/constants"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/rbac"
	"be-blog/src/mail"
	"be-blog/src/models"
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
)

func RandomOTP() string {
	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	return totp.Now()
}

func Register(body models.RegisterReq) (*models.RegisterRes, error) {
	hash, err := argon2id.CreateHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		return nil, err
	}
	Id, _ := uuid.NewV7()
	newUser := entities.Auth{
		ID:       Id.String(),
		Username: strings.ToLower(body.Username),
		Email:    strings.ToLower(body.Email),
		Password: hash,
		CodeMail: RandomOTP(),
		Roles:    []string{rbac.RoleUser},
		Status:   constants.STATUS_INACTIVE,
	}
	_, err = config.DB.Model(&newUser).Insert()
	if err != nil {
		if pgErr, ok := err.(pg.Error); ok {
			code := pgErr.Field('C')
			if code == errors.DUPLICATE_KEY {
				return nil, errors.NewError(err).SetMessage("Tên đăng nhập hoặc email đã tồn tại")
			}
		}
		return nil, err
	}
	token, err := jwt.CreateTokenVerifyEmail(jwt.PayloadVerify{
		Email: newUser.Email,
	}, constants.ONE_HOUR)
	if err != nil {
		return nil, err
	}
	go mail.SendMail(mail.PayloadMail{
		Tos:      []string{newUser.Email},
		Template: "REGISTER",
		Data:     map[string]interface{}{"fullName": newUser.FullName, "code": newUser.CodeMail},
		From:     viper.GetString("mailSender"),
	})
	result := models.RegisterRes{
		ID:         newUser.ID,
		FullName:   newUser.FullName,
		UserName:   newUser.Username,
		Email:      newUser.Email,
		EmailToken: token,
	}
	return &result, nil
}
