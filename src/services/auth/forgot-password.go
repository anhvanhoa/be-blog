package auth

import (
	"be-blog/src/config"
	"be-blog/src/constants"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/mail"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

func ForgotPassword(email string) error {
	var user entities.Auth
	err := config.DB.Model(&user).Where("status = 'active' AND email = ?", email).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return errors.NewErrorBadRequest("Email không tồn tại")
		}
		return err
	}

	token, err := jwt.CreateForgotPassToken(jwt.PayloadResetPass{
		Email: email,
	}, constants.TEN_MIN)

	if err != nil {
		return err
	}

	go mail.SendMail(mail.PayloadMail{
		Tos:      []string{		email},
		Template: "FORGOT_PASSWORD",
		Data:     map[string]interface{}{"link": viper.GetString("hostClient") + "/forgot-password/" + token},
		From:     viper.GetString("mailSender"),
	})

	return nil
}

func CheckResetPassword(token string) (string, error) {
	payload, err := jwt.ParseTokenPayloadForgotPass(token)
	if err != nil {
		return "", errors.NewErrorBadRequest("Token không hợp lệ")
	}
	return payload.Email, nil
}