package user

import (
	"be-blog/src/config"
	"be-blog/src/entities"
	"be-blog/src/libs/rbac"
	"be-blog/src/models"

	"github.com/go-pg/pg/v10"
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

func UpdateUserStatus(req models.StatusUserReq) error {
	user := &entities.User{ID: req.ID}
	_, err := config.DB.Model(user).Set("status = ?", req.Status).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserRoles(req models.RolesUserReq) error {
	if len(req.Roles) == 0 {
		req.Roles = []string{rbac.RoleUser}
	}
	user := &entities.User{ID: req.ID}
	_, err := config.DB.Model(user).Set("roles = ?", pg.Array(req.Roles)).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
