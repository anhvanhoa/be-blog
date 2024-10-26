package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerAccountRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/account"})
	router.Get("/me", rbac.AllowAll(), true, controllers.Login)
}
