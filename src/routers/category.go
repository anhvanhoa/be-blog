package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerCategoryRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/categories"})
	router.Get("", "Lấy danh sách danh mục", true, controllers.GetCategories)
}
