package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerBlogRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/blogs"})
	router.Get("", "Lấy danh sách bài viết", false, controllers.GetBlogs)
	router.Get("/category/{slug}", "Lấy danh sách bài viết theo danh mục", false, controllers.GetBlogsByCategory)
	router.Get("/{slug}", "Lấy chi tiết bài viết", false, controllers.GetBlogBySlug)
}
