package routers

import (
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12"
)

func RegisterRouter(app *iris.Application) {
	api := app.Party("/api")
	api.Use(rbac.MiddlewarePermission)
	registerBlogRouter(api)
	registerCategoryRouter(api)
	registerAuthRouter(api)
	MailRouter(api)
}
