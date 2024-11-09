package account

import (
	"be-blog/src/config"
	"be-blog/src/entities"
)

func DeleteAvatar(id string) error {
	_, err := config.DB.Model(&entities.User{ID: id}).WherePK().Set("avatar = ?", nil).Update()
	if err != nil {
		return err
	}
	return nil
}

func UploadAvatar(id string, avatar string) error {
	_, err := config.DB.Model(&entities.User{ID: id}).WherePK().Set("avatar = ?", avatar).Update()
	if err != nil {
		return err
	}
	return nil
}
