package routers

import (
	"be-blog/src/libs/rbac"
	"be-blog/src/ws"

	"github.com/kataras/iris/v12"
)

func RegisterRouter(app *iris.Application) {
	ws.RegisterWs(app)
	api := app.Party("/api")
	api.Use(rbac.MiddlewarePermission)
	registerBlogRouter(api)
	registerCommentRouter(api)
	registerCategoryRouter(api)
	registerAuthRouter(api)
	registerAccountRouter(api)
	MailRouter(api)
	TagRouter(api)
	MediaRouter(api)
	UserRouter(api)
	StatisticRouter(api)
}
