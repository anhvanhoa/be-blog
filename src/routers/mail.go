package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func MailRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/mail"})
	router.Get("/histories", "Lịch sử gửi mail", true, controllers.HistoriesMail)
}
