package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func StatisticRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/statistic"})
	router.Get("/", rbac.Allow(rbac.RoleAdmin), true, controllers.Statistic) //Thông kê chung

}
