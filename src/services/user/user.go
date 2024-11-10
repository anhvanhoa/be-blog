package user

import (
	"be-blog/src/config"
	"be-blog/src/entities"
)

func GetUsers() ([]entities.User, error) {
	users := []entities.User{}
	err := config.DB.Model(&users).ColumnExpr("id, username, email, full_name, roles, status").Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id string) (*entities.User, error) {
	user := &entities.User{ID: id}
	err := config.DB.Model(user).WherePK().ExcludeColumn("password").Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}
