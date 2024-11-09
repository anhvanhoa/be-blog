package account

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/models"
	"be-blog/src/services"
	"time"
)

func Profile(id string) (profile models.Profile, err error) {
	var account entities.User
	err = config.DB.Model(&account).Where("id = ?", id).Select()
	profile.ID = account.ID
	profile.Username = account.Username
	profile.Email = account.Email
	profile.Avatar = account.Avatar
	profile.Bio = account.Bio
	profile.Gender = account.Gender
	profile.FullName = account.FullName
	profile.Birthday = account.Birthday.Format("2006-01-02")
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func UpdateProfile(body models.ProfileUpdateReq, user *jwt.PayloadUser) error {
	if body.Email != user.Email {
		ok, err := services.CheckEmailExist(body.Email)
		if err != nil {
			return err
		}
		if ok {
			err := errors.NewErrorBadRequest("Email đã tồn tại")
			return err
		}
	}
	if body.Username != user.UserName {
		ok, err := services.CheckUserExist(body.Username)
		if err != nil {
			return err
		}
		if ok {
			err := errors.NewErrorBadRequest("Username đã tồn tại")
			return err
		}
	}
	formatBirthday, err := time.Parse("2006-01-02", body.Birthday)

	if err != nil {
		return err
	}

	profile := entities.UserUpdate{
		ID:        body.ID,
		Username:  body.Username,
		Email:     body.Email,
		FullName:  body.FullName,
		Bio:       body.Bio,
		Gender:    body.Gender,
		Birthday:  formatBirthday,
		UpdatedAt: time.Now(),
	}
	_, err = config.DB.Model(&profile).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
