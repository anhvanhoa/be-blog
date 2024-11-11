package auth

import (
	"be-blog/src/config"
	"be-blog/src/constants"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/models"
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/go-pg/pg/v10"
)

func Login(body models.LoginReq) (*models.LoginRes, string, error) {
	var user entities.Auth
	err := config.DB.Model(&user).Where("username = ? AND status = ?", strings.ToLower(body.Username), constants.STATUS_ACTIVE).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, "", errors.NewErrorBadRequest("Tài khoản không chính xác")
		}
		return nil, "", err
	}
	ok, err := argon2id.ComparePasswordAndHash(body.Password, user.Password)
	if err != nil {
		return nil, "", err
	}
	if !ok {
		return nil, "", errors.NewErrorBadRequest("Mật khẩu không chính xác")
	}
	token, err := jwt.CreateTokenUser(jwt.PayloadUser{
		ID:       user.ID,
		FullName: user.FullName,
		UserName: user.Username,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Roles:    strings.Join(user.Roles, ","),
	}, constants.SIX_MONTH)

	if err != nil {
		return nil, "", err
	}

	return &models.LoginRes{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}, token, nil
}
