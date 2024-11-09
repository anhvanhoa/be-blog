package account

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/models"

	"github.com/alexedwards/argon2id"
)

func ChangePassword(id string, req models.ChangePasswordReq) error {
	user := entities.User{ID: id}
	err := config.DB.Model(&user).WherePK().Select()
	if err != nil {
		return errors.NewErrorNotFound("Tài khoản không tồn tại")
	}

	ok, err := argon2id.ComparePasswordAndHash(req.Password, user.Password)
	if err != nil {
		return err
	}

	if !ok {
		return errors.NewErrorBadRequest("Mật khẩu không đúng")
	}

	if req.NewPassword != req.ConfirmPassword {
		return errors.NewErrorBadRequest("Mật khẩu xác nhận không trùng khớp")
	}

	hash, err := argon2id.CreateHash(req.NewPassword, argon2id.DefaultParams)

	_, err = config.DB.Model(&user).WherePK().Set("password = ?", hash).Update()
	if err != nil {
		return err
	}
	return nil
}
