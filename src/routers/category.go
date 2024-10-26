package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerCategoryRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/categories"})
	router.Get("", rbac.AllowAll(), false, controllers.GetCategories) //Lấy danh sách danh mục

	routerManager := rbac.NewRoute(app, &rbac.Config{RelativePath: "/manager/categories"})
	routerManager.Get("", rbac.Allow(rbac.RoleAdmin), true, controllers.GetManagerCategories) // Lấy danh sách danh mục
	routerManager.Get("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.GetCategoryById) // Lấy danh mục theo id
	routerManager.Post("", rbac.Allow(rbac.RoleAdmin), true, controllers.CreateCategory)      // Tạo danh mục
	routerManager.Put("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.UpdateCategory)   // Cập nhật danh mục
	routerManager.Delete("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.DeleteCategory) // Xóa danh mục
}
