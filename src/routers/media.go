package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func MediaRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/media"})
	router.Get("/images", rbac.Allow(rbac.RoleAdmin), true, controllers.GetImages)    //Lấy danh sách ảnh
	router.Post("/images", rbac.Allow(rbac.RoleAdmin), true, controllers.UploadImage) // Tải ảnh lên
	router.Delete("/images/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.DeleteImage)
	router.Put("/images", rbac.Allow(rbac.RoleAdmin), true, controllers.UpdateImage)
}
