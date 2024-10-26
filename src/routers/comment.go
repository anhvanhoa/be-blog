package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerCommentRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/comments"})
	router.Post("", rbac.AllowAll(), true, controllers.CommentPost)             //Tạo bình luận
	router.Get("/{slug}", rbac.AllowAll(), true, controllers.GetCommentsByBlog) //Lấy danh sách bình luận
}
