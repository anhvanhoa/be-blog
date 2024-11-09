package auth

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/models"

	"github.com/alexedwards/argon2id"
)

func ResetPassword(body models.ResetPasswordReq) error {
	payload, err := jwt.ParseTokenPayloadForgotPass(body.T)
	if err != nil {
		return err
	}
	if body.Password != body.ConfirmPassword {
		return errors.NewErrorBadRequest("Mật khẩu không khớp")
	}

	hash, err := argon2id.CreateHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		return err
	}
	_, err = config.DB.Model(&entities.UserResetPass{
		Email:    payload.Email,
		Password: hash,
	}).Where("email = ?", payload.Email).Update("password", hash)

	if err != nil {
		return err
	}
	return nil
}
