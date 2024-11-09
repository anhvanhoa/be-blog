package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerAccountRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/account"})
	router.Get("/me", rbac.AllowAll(), true, controllers.Me)
	router.Get("/profile", rbac.AllowAll(), true, controllers.Profile)
	router.Put("/profile", rbac.AllowAll(), true, controllers.UpdateProfile)
	router.Post("/avatar", rbac.AllowAll(), true, controllers.UploadAvatar)
	router.Delete("/avatar", rbac.AllowAll(), true, controllers.DeleteAvatar)
	router.Put("/change-password", rbac.AllowAll(), true, controllers.ChangePassword)
}
