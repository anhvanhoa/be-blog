package services

import (
	"be-blog/src/config"
	"be-blog/src/entities"
)

func CheckUserExist(username string) (bool, error) {
	count, err := config.DB.Model(&entities.User{}).Where("username = ?", username).Count()
	if count == 0 {
		return false, err
	}
	return true, nil
}

func CheckEmailExist(email string) (bool, error) {
	count, err := config.DB.Model(&entities.User{}).Where("email = ?", email).Count()
	if count == 0 {
		return false, err
	}
	return true, nil
}
