package controllers

import (
	"be-blog/src/libs/errors"
	"be-blog/src/libs/logger"
	category_service "be-blog/src/services/category"

	"github.com/kataras/iris/v12"
)

func GetCategories(ctx iris.Context) {
	categories, err := category_service.GetCategories(true)
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    categories,
		Message: "Success",
	})
}
