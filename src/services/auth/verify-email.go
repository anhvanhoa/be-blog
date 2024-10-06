package auth

import (
	"be-blog/src/config"
	"be-blog/src/constants"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"time"

	"github.com/go-pg/pg/v10"
)

func CheckVerifyEmail(tokenEmail string) (string, bool) {
	claim, err := jwt.ParseTokenPayloadVerifyEmail(tokenEmail)
	if err != nil {
		return "", false
	}
	if claim.Exp < time.Now().Unix() {
		return "", false
	}

	email := claim.Email
	count, err := config.DB.Model(&entities.Auth{}).Where("email = ? AND status = ?", email, constants.STATUS_INACTIVE).Count()
	if err != nil || count == 0 {
		return "", false
	}
	return email, true
}

func ResendCodeVerifyEmail(email string) error {
	_, err := config.DB.Model(&entities.Auth{}).Where("email = ?", email).Set("code_mail = ?", RandomOTP()).Update()
	if err != nil {
		return err
	}
	return nil
}

func VerifyAccount(mail, code string) error {
	var user *entities.Auth = new(entities.Auth)
	err := config.DB.Model(user).Where("code_mail = ? AND email = ?", code, mail).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return errors.NewErrorBadRequest("Mã xác thực không hợp lệ")
		}
		return err
	}
	if user.Email != mail {
		err := errors.NewErrorBadRequest("Mã xác thực không hợp lệ")
		return err
	}
	_, err = config.DB.Model(&entities.Auth{}).Where("code_mail = ? AND email = ?", code, mail).Set("status = ?", constants.STATUS_ACTIVE).Update()
	if err != nil {
		return err
	}
	return nil
}
