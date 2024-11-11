package controllers

import "github.com/kataras/iris/v12"

func Statistic(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "Statistic",
	})
}
