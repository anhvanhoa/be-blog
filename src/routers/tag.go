package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func TagRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/tags"})
	router.Get("", rbac.AllowAll(), true, controllers.GetTags)                       // Lấy danh sách tag
	router.Post("", rbac.Allow(rbac.RoleAdmin), true, controllers.CreateTag)         // Tạo tag
	router.Put("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.UpdateTag)     // Sửa tag
	router.Get("/{id:string}", rbac.Allow(rbac.RoleAdmin), true, controllers.GetTag) // Lấy tag theo id
	router.Delete("/{id:string}", rbac.Allow(rbac.RoleAdmin), true, controllers.DeleteTag)
}
