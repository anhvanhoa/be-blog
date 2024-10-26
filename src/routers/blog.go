package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerBlogRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/blogs"})                     // group
	router.Get("", rbac.Allow(), false, controllers.GetBlogs)                              // Lấy danh sách bài viết
	router.Get("/category/{slug}", rbac.AllowAll(), false, controllers.GetBlogsByCategory) // Lấy danh sách bài viết theo danh mục
	router.Get("/{slug}", rbac.AllowAll(), false, controllers.GetBlogBySlug)               // Lấy chi tiết bài viết

	//	manager blog
	routerManager := rbac.NewRoute(app, &rbac.Config{RelativePath: "/manager/blogs"})     // group
	routerManager.Get("", rbac.Allow(rbac.RoleAdmin), true, controllers.GetManagerBlogs)  // Lấy danh sách bài viết cho admin
	routerManager.Post("", rbac.Allow(rbac.RoleAdmin), true, controllers.CreateBlog)      // Tạo bài viết
	routerManager.Put("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.UpdateBlog)  // Cập nhật bài viết
	routerManager.Get("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.GetBlogById) // Lấy chi tiết bài viết cho admin
	routerManager.Delete("/{id}", rbac.Allow(rbac.RoleAdmin), true, controllers.DeleteBlog) // Xóa bài viết
}
