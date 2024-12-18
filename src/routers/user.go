package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func UserRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/users"})
	router.Get("", rbac.Allow(rbac.RoleAdmin), true, controllers.GetUsers) // Lấy danh sách user
	router.Get("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.GetUser) // Lấy user
	router.Put("/status", rbac.Allow(rbac.RoleAdmin), true, controllers.UpdateUserStatus) // Cập nhật trạng thái user
	router.Put("/roles", rbac.Allow(rbac.RoleAdmin), true, controllers.UpdateUserRoles) // Cập nhật roles user
}
